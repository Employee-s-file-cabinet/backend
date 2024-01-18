package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
	"github.com/Employee-s-file-cabinet/backend/pkg/repoerr"
)

const listContractsQuery = `SELECT 
contracts.id as id, number, contract_type, work_types.title as work_type, probation_period, date_begin, date_end
FROM contracts
JOIN work_types ON contracts.work_type_id = work_types.id
WHERE user_id = @user_id`

func (s *storage) ListContracts(ctx context.Context, userID uint64) ([]model.Contract, error) {
	const op = "postrgresql user storage: list contracts"

	rows, err := s.DB.Query(ctx, listContractsQuery, pgx.NamedArgs{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	trs, err := pgx.CollectRows[contract](rows, pgx.RowToStructByNameLax[contract])
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	contracts := make([]model.Contract, len(trs))
	for i, tr := range trs {
		contracts[i] = convertContractToModelContract(tr)
	}

	return contracts, nil
}

func (s *storage) GetContract(ctx context.Context, userID, contractID uint64) (*model.Contract, error) {
	const op = "postrgresql user storage: get contract"

	//стр-ра
	rows, err := s.DB.Query(ctx,
		`SELECT 
		id, number, contract_type, work_type_id, probation_period, date_begin, date_end
		FROM contracts
		WHERE id = @contract_id AND user_id = @user_id`,
		pgx.NamedArgs{
			"contract_id": contractID,
			"user_id":     userID,
		})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	ed, err := pgx.CollectExactlyOneRow[contract](rows, pgx.RowToStructByNameLax[contract])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("%s: %w", op, repoerr.ErrRecordNotFound)
	}

	med := convertContractToModelContract(ed)
	return &med, nil
}

func (s *storage) AddContract(ctx context.Context, userID uint64, tr model.Contract) (uint64, error) {
	const op = "postrgresql user storage: add contract"

	row := s.DB.QueryRow(ctx, `INSERT INTO contracts
		("user_id", "number", "contract_type", "work_type_id", "probation_period", "date_begin", "date_end")
		VALUES (@user_id, @number, @contract_type, @work_type_id, @probation_period, @date_begin, @date_end)
		RETURNING "id"`,
		pgx.NamedArgs{
			"user_id":          userID,
			"number":           tr.Number,
			"contract_type":    tr.ContractType,
			"work_type_id":     tr.WorkTypeID,
			"probation_period": tr.ProbationPeriod,
			"date_begin":       tr.DateBegin,
			"date_end":         tr.DateEnd,
		})

	if err := row.Scan(&tr.ID); err != nil {
		if strings.Contains(err.Error(), "23") && // Integrity Constraint Violation
			strings.Contains(err.Error(), "user_id") {
			return 0, fmt.Errorf("%s: the user does not exist: %w", op, repoerr.ErrRecordNotFound)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return tr.ID, nil
}
