package manage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"test/internal/model"
)

func PatchWorkspace(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch WorkspacePatch,
) (*model.Workspace, error) {
	query := "UPDATE workspace SET "
	args := []any{}
	i := 1

	if patch.Name != nil {
		query += fmt.Sprintf("name = $%d,", i)
		args = append(args, *patch.Name)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, name
	`, i)
	args = append(args, id)

	workspace := model.Workspace{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&workspace.Id,
		&workspace.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("workspace not found")
		}
		return nil, fmt.Errorf("PatchWorkspace: %w", err)
	}

	return &workspace, nil
}

func PatchCompany(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch CompanyPatch,
) (*model.Company, error) {
	query := "UPDATE company SET "
	args := []any{}
	i := 1

	if patch.Name != nil {
		query += fmt.Sprintf("name = $%d,", i)
		args = append(args, *patch.Name)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, name, workspace_id
	`, i)
	args = append(args, id)

	company := model.Company{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&company.Id,
		&company.Name,
		&company.WorkspaceId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("company not found")
		}
		return nil, fmt.Errorf("PatchCompany: %w", err)
	}

	return &company, nil
}

func PatchLocation(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch LocationPatch,
) (*model.Location, error) {
	query := "UPDATE location SET "
	args := []any{}
	i := 1

	if patch.Name != nil {
		query += fmt.Sprintf("name = $%d,", i)
		args = append(args, *patch.Name)
		i++
	}
	if patch.Address != nil {
		query += fmt.Sprintf("address = $%d,", i)
		args = append(args, *patch.Address)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, name, address, workspace_id
	`, i)
	args = append(args, id)

	location := model.Location{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&location.Id,
		&location.Name,
		&location.Address,
		&location.WorkspaceId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("location not found")
		}
		return nil, fmt.Errorf("PatchLocation: %w", err)
	}

	return &location, nil
}

func PatchTask(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch TaskPatch,
) (*model.Task, error) {
	query := "UPDATE task SET "
	args := []any{}
	i := 1

	if patch.Name != nil {
		query += fmt.Sprintf("name = $%d,", i)
		args = append(args, *patch.Name)
		i++
	}
	if patch.Description != nil {
		query += fmt.Sprintf("description = $%d,", i)
		args = append(args, *patch.Description)
		i++
	}
	if patch.IsCompleted != nil {
		query += fmt.Sprintf("is_completed = $%d,", i)
		args = append(args, *patch.IsCompleted)
		i++
	}
	if patch.LocationId != nil {
		query += fmt.Sprintf("location_id = $%d,", i)
		args = append(args, *patch.LocationId)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, name, description, location_id, company_id, is_completed
	`, i)
	args = append(args, id)

	task := model.Task{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&task.Id,
		&task.Name,
		&task.Description,
		&task.LocationId,
		&task.CompanyId,
		&task.IsCompleted,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task not found")
		}
		return nil, fmt.Errorf("PatchTask: %w", err)
	}

	return &task, nil
}

func PatchEmployment(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch EmploymentPatch,
) (*model.Employment, error) {
	query := "UPDATE employment SET "
	args := []any{}
	i := 1

	if patch.Role != nil {
		query += fmt.Sprintf("role = $%d,", i)
		args = append(args, *patch.Role)
		i++
	}
	if patch.ContractId != nil {
		query += fmt.Sprintf("contract_id = $%d,", i)
		args = append(args, *patch.ContractId)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, profile_id, company_id, contract_id, role, start_date, end_date
	`, i)
	args = append(args, id)

	employment := model.Employment{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&employment.Id,
		&employment.ProfileId,
		&employment.CompanyId,
		&employment.ContractId,
		&employment.Role,
		&employment.StartDate,
		&employment.EndDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("employment not found")
		}
		return nil, fmt.Errorf("PatchEmployment: %w", err)
	}

	return &employment, nil
}

func PatchContract(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch ContractPatch,
) (*model.Contract, error) {
	query := "UPDATE contract SET "
	args := []any{}
	i := 1

	if patch.HourlyRate != nil {
		query += fmt.Sprintf("hourly_rate = $%d,", i)
		args = append(args, *patch.HourlyRate)
		i++
	}
	if patch.UnpaidLunchMinutes != nil {
		query += fmt.Sprintf("unpaid_lunch_minutes = $%d,", i)
		args = append(args, *patch.UnpaidLunchMinutes)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, hourly_rate, unpaid_lunch_minutes
	`, i)
	args = append(args, id)

	contract := model.Contract{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&contract.Id,
		&contract.HourlyRate,
		&contract.UnpaidLunchMinutes,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("contract not found")
		}
		return nil, fmt.Errorf("PatchContract: %w", err)
	}

	return &contract, nil
}

func PatchProfile(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch ProfilePatch,
) (*model.Profile, error) {
	query := "UPDATE profile SET "
	args := []any{}
	i := 1

	if patch.FirstName != nil {
		query += fmt.Sprintf("first_name = $%d,", i)
		args = append(args, *patch.FirstName)
		i++
	}
	if patch.LastName != nil {
		query += fmt.Sprintf("last_name = $%d,", i)
		args = append(args, *patch.LastName)
		i++
	}
	if patch.KT != nil {
		query += fmt.Sprintf("kt = $%d,", i)
		args = append(args, *patch.KT)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, kt, first_name, last_name
	`, i)
	args = append(args, id)

	profile := model.Profile{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&profile.ID,
		&profile.KT,
		&profile.FirstName,
		&profile.LastName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("profile not found")
		}
		return nil, fmt.Errorf("PatchProfile: %w", err)
	}

	return &profile, nil
}

func PatchShift(
	ctx context.Context,
	db *sql.DB,
	id int,
	patch ShiftPatch,
) (*model.Shift, error) {
	query := "UPDATE shift SET "
	args := []any{}
	i := 1

	if patch.TaskId != nil {
		query += fmt.Sprintf("task_id = $%d,", i)
		args = append(args, *patch.TaskId)
		i++
	}
	if patch.StartTs != nil {
		query += fmt.Sprintf("start_ts = $%d,", i)
		args = append(args, *patch.StartTs)
		i++
	}
	if patch.EndTs != nil {
		query += fmt.Sprintf("end_ts = $%d,", i)
		args = append(args, *patch.EndTs)
		i++
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(` 
		WHERE id = $%d
		RETURNING id, profile_id, task_id, start_ts, end_ts
	`, i)
	args = append(args, id)

	shift := model.Shift{}
	err := db.QueryRowContext(ctx, query, args...).Scan(
		&shift.Id,
		&shift.ProfileId,
		&shift.TaskId,
		&shift.StartTs,
		&shift.EndTs,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("shift not found")
		}
		return nil, fmt.Errorf("PatchShift: %w", err)
	}

	return &shift, nil
}
