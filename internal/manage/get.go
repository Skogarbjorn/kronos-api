package manage

import (
	"context"
	"database/sql"
	"fmt"
	"test/internal/model"
)

func GetWorkspaces(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Workspace, error) {
	workspaces := []model.Workspace{}
	rows, err := db.Query(
		`
		SELECT id, name
		FROM workspace
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetWorkspaces: db select: %w", err)
	}

	for rows.Next() {
		var workspace model.Workspace
		err = rows.Scan(
			&workspace.Id,
			&workspace.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("GetWorkspaces: db scan: %w", err)
		}

		workspaces = append(workspaces, workspace)
	}

	return &workspaces, nil
}

func GetCompanies(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Company, error) {
	companies := []model.Company{}
	rows, err := db.Query(
		`
		SELECT id, name, workspace_id
		FROM company
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetCompanies: db select: %w", err)
	}

	for rows.Next() {
		var company model.Company
		err = rows.Scan(
			&company.Id,
			&company.Name,
			&company.WorkspaceId,
		)
		if err != nil {
			return nil, fmt.Errorf("GetCompanies: db scan: %w", err)
		}

		companies = append(companies, company)
	}

	return &companies, nil
}

func GetLocations(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Location, error) {
	locations := []model.Location{}
	rows, err := db.Query(
		`
		SELECT id, name, address, workspace_id
		FROM location
		`,
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
) (*[]model.Task, error) {
	tasks := []model.Task{}
	rows, err := db.Query(
		`
		SELECT id, location_id, company_id, name, description, is_completed
		FROM task
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetTasks: db select: %w", err)
	}

	for rows.Next() {
		var task model.Task
		err = rows.Scan(
			&task.Id,
			&task.LocationId,
			&task.CompanyId,
			&task.Name,
			&task.Description,
			&task.IsCompleted,
		)
		if err != nil {
			return nil, fmt.Errorf("GetTasks: db scan: %w", err)
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func GetProfiles(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Profile, error) {
	profiles := []model.Profile{}
	rows, err := db.Query(
		`
		SELECT id, kt, first_name, last_name
		FROM profile
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetProfiles: db select: %w", err)
	}

	for rows.Next() {
		var profile model.Profile
		err = rows.Scan(
			&profile.ID,
			&profile.KT,
			&profile.FirstName,
			&profile.LastName,
		)
		if err != nil {
			return nil, fmt.Errorf("GetProfiles: db scan: %w", err)
		}

		profiles = append(profiles, profile)
	}

	return &profiles, nil
}

func GetEmployments(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Employment, error) {
	employments := []model.Employment{}
	rows, err := db.Query(
		`
		SELECT id, profile_id, company_id, contract_id, role, start_date, end_date
		FROM employment
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetEmployments: db select: %w", err)
	}

	for rows.Next() {
		var employment model.Employment
		err = rows.Scan(
			&employment.Id,
			&employment.ProfileId,
			&employment.CompanyId,
			&employment.ContractId,
			&employment.Role,
			&employment.StartDate,
			&employment.EndDate,
		)
		if err != nil {
			return nil, fmt.Errorf("GetEmployments: db scan: %w", err)
		}

		employments = append(employments, employment)
	}

	return &employments, nil
}

func GetContracts(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Contract, error) {
	contracts := []model.Contract{}
	rows, err := db.Query(
		`
		SELECT id, hourly_rate, unpaid_lunch_minutes
		FROM contract
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetContracts: db select: %w", err)
	}

	for rows.Next() {
		var contract model.Contract
		err = rows.Scan(
			&contract.Id,
			&contract.HourlyRate,
			&contract.UnpaidLunchMinutes,
		)
		if err != nil {
			return nil, fmt.Errorf("GetContracts: db scan: %w", err)
		}

		contracts = append(contracts, contract)
	}

	return &contracts, nil
}

func GetShifts(
	ctx context.Context,
	db *sql.DB,
) (*[]model.Shift, error) {
	shifts := []model.Shift{}
	rows, err := db.Query(
		`
		SELECT id, profile_id, task_id, start_ts, end_ts
		FROM shift
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("GetShifts: db select: %w", err)
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
			return nil, fmt.Errorf("GetShifts: db scan: %w", err)
		}

		shifts = append(shifts, shift)
	}

	return &shifts, nil
}
