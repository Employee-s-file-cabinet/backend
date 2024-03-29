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

const listEducationsQuery = `SELECT
id, document_number, title_of_program,
title_of_institution, year_of_end, year_of_begin,
(SELECT COUNT(*)>0 FROM scans WHERE scans.document_id=educations.id AND scans.type='Документ об образовании') AS has_scan
FROM educations
WHERE user_id = @user_id`

func (s *storage) ListEducations(ctx context.Context, userID uint64) ([]model.Education, error) {
	const op = "postgresql user storage: list educations"

	rows, err := s.DB.Query(ctx, listEducationsQuery, pgx.NamedArgs{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	eds, err := pgx.CollectRows[education](rows, pgx.RowToStructByNameLax[education])
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	educations := make([]model.Education, len(eds))
	for i, ed := range eds {
		educations[i] = convertEducationToModelEducation(ed)
	}

	return educations, nil
}

func (s *storage) GetEducation(ctx context.Context, userID, educationID uint64) (*model.Education, error) {
	const op = "postgresql user storage: get education"

	rows, err := s.DB.Query(ctx, `SELECT
		id, document_number, title_of_program,
		title_of_institution, year_of_end, year_of_begin,
		(SELECT COUNT(*)>0 FROM scans WHERE user_id=@user_id AND scans.document_id=educations.id AND scans.type='Документ об образовании') AS has_scan
		FROM educations
		WHERE id = @education_id AND user_id = @user_id`,
		pgx.NamedArgs{
			"education_id": educationID,
			"user_id":      userID,
		})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	ed, err := pgx.CollectExactlyOneRow[education](rows, pgx.RowToStructByNameLax[education])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrRecordNotFound
	}

	med := convertEducationToModelEducation(ed)
	return &med, nil
}

func (s *storage) AddEducation(ctx context.Context, userID uint64, ed model.Education) (uint64, error) {
	const op = "postgresql user storage: add education"

	row := s.DB.QueryRow(ctx, `INSERT INTO educations
		("user_id", "document_number", "title_of_program", 
		"title_of_institution", "year_of_end", "year_of_begin") 
		VALUES (@user_id, @number, @program, 
		@issued_institution, @date_to, @date_from)
		RETURNING "id"`,
		pgx.NamedArgs{
			"user_id":            userID,
			"number":             ed.Number,
			"program":            ed.Program,
			"issued_institution": ed.IssuedInstitution,
			"date_to":            ed.DateTo,
			"date_from":          ed.DateFrom,
		})

	if err := row.Scan(&ed.ID); err != nil {
		if strings.Contains(err.Error(), "23") && // Integrity Constraint Violation
			strings.Contains(err.Error(), "user_id") {
			return 0, fmt.Errorf("the user does not exist: %w", repoerr.ErrConflict)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return ed.ID, nil
}

func (s *storage) UpdateEducation(ctx context.Context, userID uint64, ed model.Education) error {
	const op = "postrgresql user storage: update education"

	tag, err := s.DB.Exec(ctx, `UPDATE educations
		SET document_number = @number, 
		title_of_program = @program, 
		title_of_institution = @issued_institution, 
		year_of_end = @date_to, 
		year_of_begin = @date_from
		WHERE id=@id AND user_id=@user_id`,
		pgx.NamedArgs{
			"user_id":            userID,
			"id":                 ed.ID,
			"number":             ed.Number,
			"program":            ed.Program,
			"issued_institution": ed.IssuedInstitution,
			"date_to":            ed.DateTo,
			"date_from":          ed.DateFrom,
		})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if tag.RowsAffected() == 0 { // it's ok for pgx
		return repoerr.ErrRecordNotAffected
	}
	return nil
}
