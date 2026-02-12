package pin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test/internal/auth"
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
		INSERT INTO shift (employment_id, task_id, start_ts)
		VALUES ($1, $2, $3)
		RETURNING id, employment_id, task_id, start_ts
		`,
		input.EmploymentId,
		input.TaskId,
		input.StartTs,
	).Scan(
		&shift.Id,
		&shift.EmploymentId,
		&shift.TaskId,
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

func GetShiftOverview(
	ctx context.Context,
	db *sql.DB,
) (*ShiftOverview, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	user_id := claims.UserID

	var shiftOverview ShiftOverview
	err := db.QueryRowContext(
		ctx,
		`
		SELECT s.id, s.employment_id, s.task_id, s.start_ts, w.name, c.name, l.name, t.name 
		FROM shift s
		JOIN employment e ON e.id = s.employment_id
		JOIN task t ON t.id = s.task_id
		JOIN location l ON l.id = t.location_id
		JOIN company c ON c.id = t.company_id
		JOIN workspace w ON w.id = c.workspace_id
		WHERE e.user_id = $1
		AND s.end_ts IS NULL
		`,
		user_id,
	).Scan(
		&shiftOverview.Shift.Id,
		&shiftOverview.Shift.EmploymentId,
		&shiftOverview.Shift.TaskId,
		&shiftOverview.Shift.StartTs,
		&shiftOverview.Workspace,
		&shiftOverview.Company,
		&shiftOverview.Location,
		&shiftOverview.Task,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotClockedIn
		}
		return nil, fmt.Errorf("GetShiftOverview: db select: %w", err)
	}

	return &shiftOverview, nil
}

func GetShiftHistory(
	ctx context.Context,
	db *sql.DB,
) (*[]Shift, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	user_id := claims.UserID

	var shifts []Shift
	rows, err := db.Query(
		`
		SELECT s.id, s.employment_id, s.task_id, s.start_ts, s.end_ts
		FROM shift s
		JOIN employment e ON e.id = s.employment_id
		WHERE e.user_id = $1
		`,
		user_id,
	)
	if err != nil {
		return nil, fmt.Errorf("GetShiftHistory: db select: %w", err)
	}

	for rows.Next() {
		var shift Shift
		err = rows.Scan(
			&shift.Id,
			&shift.EmploymentId,
			&shift.TaskId,
			&shift.StartTs,
			&shift.EndTs,
		)
		if err != nil {
			return nil, fmt.Errorf("GetShiftHistory: db scan: %w", err)
		}

		shifts = append(shifts, shift)
	}

	return &shifts, nil
}
