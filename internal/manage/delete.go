package manage

import (
	"context"
	"database/sql"
	"fmt"
)

func DeleteWorkspace(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM workspace WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteWorkspace: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteWorkspace: rows affected: %w", err)
	}

	return rows, nil
}

func DeleteCompany(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM company WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteCompany: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteCompany: rows affected: %w", err)
	}

	return rows, nil
}

func DeleteLocation(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM location WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteLocation: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteLocation: rows affected: %w", err)
	}

	return rows, nil
}

func DeleteTask(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM task WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteTask: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteTask: rows affected: %w", err)
	}

	return rows, nil
}

func DeleteProfile(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM profile WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteProfile: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteProfile: rows affected: %w", err)
	}

	return rows, nil
}

func DeleteContract(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM contract WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteContract: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteContract: rows affected: %w", err)
	}

	return rows, nil
}

func DeleteShift(
	ctx context.Context,
	db *sql.DB,
	id int,
) (int64, error) {
	result, err := db.ExecContext(
		ctx,
		`
		DELETE FROM shift WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("DeleteShift: db delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteShift: rows affected: %w", err)
	}

	return rows, nil
}
