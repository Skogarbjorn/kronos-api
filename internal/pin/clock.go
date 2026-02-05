package pin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func ClockIn(
	ctx context.Context,
	db *sql.DB,
	input ClockIn_R,
) (*Shift, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ClockIn: begin tx: %w", err)
	}
	defer tx.Rollback()

	if input.StartTs == nil {
		now := time.Now()
		input.StartTs = &now
	}

	var shift Shift
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO shift (employment_id, start_ts)
		VALUES ($1, $2)
		RETURNING id, employment_id, start_ts
		`,
		input.EmploymentId,
		input.StartTs,
	).Scan(
		&shift.Id,
		&shift.EmploymentId,
		&shift.StartTs,
	)
	if err != nil {
		return nil, translateDBError(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("ClockIn: db commit: %w", err)
	}

	return &shift, nil
}

func ClockOut(
	ctx context.Context,
	db *sql.DB,
	input ClockOut_R,
) (*Shift, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ClockIn: begin tx: %w", err)
	}
	defer tx.Rollback()

	if input.EndTs == nil {
		now := time.Now()
		input.EndTs = &now
	}

	var shift Shift
	err = tx.QueryRowContext(
		ctx,
		`
		UPDATE shift 
		SET end_ts = $1
		WHERE end_ts IS NULL AND employment_id = $2
		RETURNING id, employment_id, start_ts, end_ts
		`,
		input.EndTs,
		input.EmploymentId,
	).Scan(
		&shift.Id,
		&shift.EmploymentId,
		&shift.StartTs,
		&shift.EndTs,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotClockedIn
		}
		return nil, fmt.Errorf("ClockOut: db update: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("ClockOut: db commit: %w", err)
	}

	return &shift, nil
}
