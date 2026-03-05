package pin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test/internal/auth"
	"test/internal/model"
	"time"
)

func ClockIn(
	ctx context.Context,
	db *sql.DB,
	input ClockIn_R,
) (*model.Shift, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ClockIn: begin tx: %w", err)
	}
	defer tx.Rollback()

	if input.StartTs == nil {
		now := time.Now()
		input.StartTs = &now
	}

	var shift model.Shift
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO shift (profile_id, task_id, start_ts)
		VALUES ($1, $2, $3)
		RETURNING id, profile_id, task_id, start_ts
		`,
		profile_id,
		input.TaskId,
		input.StartTs,
	).Scan(
		&shift.Id,
		&shift.ProfileId,
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
) (*model.Shift, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ClockIn: begin tx: %w", err)
	}
	defer tx.Rollback()

	if input.EndTs == nil {
		now := time.Now()
		input.EndTs = &now
	}

	var shift model.Shift
	err = tx.QueryRowContext(
		ctx,
		`
		UPDATE shift 
		SET end_ts = $1
		WHERE end_ts IS NULL AND profile_id = $2
		RETURNING id, profile_id, task_id, start_ts, end_ts
		`,
		input.EndTs,
		profile_id,
	).Scan(
		&shift.Id,
		&shift.ProfileId,
		&shift.TaskId,
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

func SyncShift(
	ctx context.Context,
	db *sql.DB,
	input SyncShift_R,
) (*model.Shift, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ClockIn: begin tx: %w", err)
	}
	defer tx.Rollback()

	var idToInsert interface{}
	if input.RemoteId != nil {
		idToInsert = *input.RemoteId
	} else {
		idToInsert = nil
	}

	var shift model.Shift
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO shift (id, profile_id, task_id, start_ts, end_ts)
		VALUES (COALESCE($1, nextval(pg_get_serial_sequence('shift', 'id'))), $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE SET
			task_id = EXCLUDED.task_id
			start_ts = EXCLUDED.start_ts
			end_ts = EXCLUDED.end_ts
			profile_id = EXCLUDED.profile_id
		RETURNING id, profile_id, task_id, start_ts, end_ts
		`,
		idToInsert,
		profile_id,
		input.TaskId,
		input.StartTs,
		input.EndTs,
	).Scan(
		&shift.Id,
		&shift.ProfileId,
		&shift.TaskId,
		&shift.StartTs,
		&shift.EndTs,
	)
	if err != nil {
		return nil, translateDBError(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("SyncShift: db commit: %w", err)
	}

	return &shift, nil
}

func GetShiftOverview(
	ctx context.Context,
	db *sql.DB,
) (*ShiftOverview, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	var shiftOverview ShiftOverview
	err := db.QueryRowContext(
		ctx,
		`
		SELECT 
			s.id, s.profile_id, s.task_id, s.start_ts,
			l.id, l.workspace_id, l.name, l.address,
			t.id, t.location_id, t.company_id, t.name, t.description, t.is_completed
		FROM shift s
		JOIN task t ON t.id = s.task_id
		JOIN location l ON l.id = t.location_id
		WHERE s.profile_id = $1
		AND s.end_ts IS NULL
		`,
		profile_id,
	).Scan(
		&shiftOverview.Shift.Id,
		&shiftOverview.Shift.ProfileId,
		&shiftOverview.Shift.TaskId,
		&shiftOverview.Shift.StartTs,
		&shiftOverview.Location.Id,
		&shiftOverview.Location.WorkspaceId,
		&shiftOverview.Location.Name,
		&shiftOverview.Location.Address,
		&shiftOverview.Task.Id,
		&shiftOverview.Task.LocationId,
		&shiftOverview.Task.CompanyId,
		&shiftOverview.Task.Name,
		&shiftOverview.Task.Description,
		&shiftOverview.Task.IsCompleted,
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
) (*[]model.Shift, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	shifts := []model.Shift{}
	rows, err := db.Query(
		`
		SELECT s.id, s.profile_id, s.task_id, s.start_ts, s.end_ts
		FROM shift s
		WHERE s.profile_id = $1
		`,
		profile_id,
	)
	if err != nil {
		return nil, fmt.Errorf("GetShiftHistory: db select: %w", err)
	}

	for rows.Next() {
		var shift model.Shift
		err = rows.Scan(
			&shift.Id,
			&shift.ProfileId,
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

func GetLocations(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Location, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	locations := []model.Location{}
	rows, err := db.Query(
		`
		SELECT l.id, l.name, l.address, l.workspace_id
		FROM location l
		JOIN workspace w ON w.id = l.workspace_id
		JOIN company c ON w.id = c.workspace_id
		JOIN employment e ON c.id = e.company_id
		WHERE e.profile_id = $1
		`,
		profile_id,
	)
	if err != nil {
		return nil, fmt.Errorf("GetLocations: db select: %w", err)
	}

	for rows.Next() {
		var location model.Location
		err = rows.Scan(
			&location.Id,
			&location.Name,
			&location.Address,
			&location.WorkspaceId,
		)
		if err != nil {
			return nil, fmt.Errorf("GetLocations: db scan: %w", err)
		}

		locations = append(locations, location)
	}

	return &locations, nil
}

func GetTasks(
	ctx context.Context,
	db *sql.DB,
	location_id *int,
) (*[]model.Task, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	tasks := []model.Task{}
	rows, err := db.Query(
		`
		SELECT t.id, t.name, t.description, t.is_completed, t.location_id, t.company_id
		FROM task t
		JOIN company c ON c.id = t.company_id
		JOIN employment e ON c.id = e.company_id
		WHERE e.profile_id = $1
		AND ($2::int IS NULL OR location_id = $2)
		`,
		profile_id,
		location_id,
	)
	if err != nil {
		return nil, fmt.Errorf("GetTasks: db select: %w", err)
	}

	for rows.Next() {
		var task model.Task
		err = rows.Scan(
			&task.Id,
			&task.Name,
			&task.Description,
			&task.IsCompleted,
			&task.LocationId,
			&task.CompanyId,
		)
		if err != nil {
			return nil, fmt.Errorf("GetTasks: db scan: %w", err)
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func GetEmploymentsDetailed(
	ctx context.Context,
	db *sql.DB,
) (*[]EmploymentDetailed, error) {
	claims := ctx.Value(auth.ClaimsKey).(*auth.Claims)
	profile_id := claims.ProfileID

	employments := []EmploymentDetailed{}
	rows, err := db.Query(
		`
		SELECT 
			w.id,
			w.name,
			c.id,
			c.name,
			c.workspace_id,
			e.id,
			e.profile_id,
			e.company_id,
			e.contract_id,
			e.role,
			e.start_date,
			e.end_date,
			ct.id,
			ct.hourly_rate,
			ct.unpaid_lunch_minutes
		FROM employment e
		JOIN contract ct ON ct.id = e.contract_id
		JOIN company c ON c.id = e.company_id
		JOIN workspace w ON w.id = c.workspace_id
		WHERE e.profile_id = $1
		`,
		profile_id,
	)
	if err != nil {
		return nil, fmt.Errorf("GetEmploymentsDetailed: db select: %w", err)
	}

	for rows.Next() {
		var e model.Employment
		var c model.Company
		var w model.Workspace
		var ct model.Contract
		err = rows.Scan(
			&w.Id,
			&w.Name,
			&c.Id,
			&c.Name,
			&c.WorkspaceId,
			&e.Id,
			&e.ProfileId,
			&e.CompanyId,
			&e.ContractId,
			&e.Role,
			&e.StartDate,
			&e.EndDate,
			&ct.Id,
			&ct.HourlyRate,
			&ct.UnpaidLunchMinutes,
		)
		if err != nil {
			return nil, fmt.Errorf("GetEmploymentsDetailed: db scan: %w", err)
		}

		detailed := EmploymentDetailed{
			w,
			c,
			e,
			ct,
		}

		employments = append(employments, detailed)
	}

	return &employments, nil
}
