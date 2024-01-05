package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/henvic/pgq"
	"github.com/jackc/pgx/v5"

	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
	"github.com/Employee-s-file-cabinet/backend/pkg/repoerr"
)

func (s *storage) ListPassports(ctx context.Context, userID uint64) ([]model.Passport, error) {
	const op = "postrgresql user storage: list passports"

	rows, err := s.DB.Query(ctx, `SELECT 
	id, number, type, issued_date, issued_by, 
	(SELECT COUNT(*) FROM visas WHERE visas.passport_id = passports.id) AS visas_count 
	FROM passports
	WHERE passports.user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	psps, err := pgx.CollectRows[passport](rows, pgx.RowToStructByNameLax[passport])
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	passport := make([]model.Passport, len(psps))
	for i, ed := range psps {
		passport[i] = convertPassportToModelPassport(ed)
	}

	return passport, nil
}

func (s *storage) GetPassport(ctx context.Context, passportID uint64) (*model.Passport, error) {
	const op = "postrgresql user storage: get passport"

	rows, err := s.DB.Query(ctx,
		`SELECT id, number, type, issued_date, issued_by,
		(SELECT COUNT(*) FROM visas WHERE visas.passport_id = $1) AS visas_count 
		FROM passports
		WHERE id = $1`, passportID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	p, err := pgx.CollectExactlyOneRow[passport](rows, pgx.RowToStructByNameLax[passport])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("%s: %w", op, repoerr.ErrRecordNotFound)
	}

	med := convertPassportToModelPassport(p)
	return &med, nil
}

func (s *storage) AddPassport(ctx context.Context, userID uint64, mp model.Passport) (uint64, error) {
	const op = "postrgresql user storage: add passport"

	p := convertModelPassportToPassport(mp)

	qb := pgq.Insert("passports").
		Columns("user_id", "number", "type", "issued_date", "issued_by").
		Values(userID, p.Number, p.Type, p.IssuedDate, p.IssuedBy).
		Returning("id")
	query, args, err := qb.SQL()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	if err := s.DB.QueryRow(ctx, query, args...).Scan(&p.ID); err != nil {
		if strings.Contains(err.Error(), "23") && // Integrity Constraint Violation
			strings.Contains(err.Error(), "user_id") {
			return 0, fmt.Errorf("%s: the user does not exist: %w", op, repoerr.ErrRecordNotFound)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return p.ID, nil
}
