package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/Employee-s-file-cabinet/backend/internal/service/auth/model"
	"github.com/Employee-s-file-cabinet/backend/pkg/repoerr"
)

const (
	selectAuthnData = `
select users.id as user_id, role_id, password_hash
from users
join authorizations a on users.id = a.user_id
where work_email=$1;`

	duplicateKeyErrorCode = "23505"
)

func (s *storage) Get(ctx context.Context, login string) (model.AuthnDAO, error) {
	rows, err := s.DB.Query(ctx, selectAuthnData, login)
	if err != nil {
		return model.AuthnDAO{}, err
	}
	authnData, err := pgx.CollectExactlyOneRow[model.AuthnDAO](rows, pgx.RowToStructByNameLax[model.AuthnDAO])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return authnData, repoerr.ErrRecordNotFound
		}
		return authnData, err
	}
	return authnData, nil
}

func (s *storage) Add(ctx context.Context, userID uint64, pswHash string) (uint64, error) {
	const op = "postrgresql auth storage: add authorization"

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `INSERT INTO authorizations
	(user_id, role_id, password_hash)
	VALUES
	(@user_id, @role_id, @password_hash)
	RETURNING id`,
		pgx.NamedArgs{
			"user_id":       userID,
			"role_id":       4, // TODO: желательно убрать константу
			"password_hash": pswHash,
		})
	var id uint64
	if err := row.Scan(&id); err != nil {
		if strings.Contains(err.Error(), "23") { // Integrity Constraint Violation
			if strings.Contains(err.Error(), "user_id") {
				return 0, fmt.Errorf("the user does not exist: %w", repoerr.ErrConflict)
			}
			if strings.Contains(err.Error(), "role_id") {
				return 0, fmt.Errorf("the role does not exist: %w", repoerr.ErrConflict)
			}
			if strings.Contains(err.Error(), duplicateKeyErrorCode) {
				return 0, fmt.Errorf("auth for the user already exists: %w", repoerr.ErrConflict)
			}
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	tag, err := tx.Exec(ctx, `INSERT INTO public.policies (ptype, v0, v1)
	VALUES ('g', @user_id, @role_id)`,
		pgx.NamedArgs{
			"user_id": userID,
			"role_id": 4, // TODO: желательно убрать константу
		})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	if tag.RowsAffected() == 0 {
		return 0, fmt.Errorf("%s: %w", op, repoerr.ErrRecordNotAffected)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
