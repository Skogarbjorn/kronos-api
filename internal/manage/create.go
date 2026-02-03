package manage

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateWorkspace(
	ctx context.Context,
	db *sql.DB,
	input WorkspaceCreate,
) (*Workspace, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateWorkspace: begin tx: %w", err)
	}
	defer tx.Rollback()

	var workspace Workspace
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO workspace (name)
		VALUES ($1)
		RETURNING id, name
		`,
		input.Name,
	).Scan(
		workspace.Id,
		workspace.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateWorkspace: db insert: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateWorkspace: db commit: %w", err)
	}

	return &workspace, nil
}

func CreateCompany(
	ctx context.Context,
	db *sql.DB,
	input CompanyCreate,
) (*Company, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateCompany: begin tx: %w", err)
	}
	defer tx.Rollback()

	var company Company
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO company (name, workspace_id)
		VALUES ($1, $2)
		RETURNING id, name, workspace_id
		`,
		input.Name,
		input.WorkspaceId,
	).Scan(
		company.Id,
		company.Name,
		company.WorkspaceId,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateCompany: db insert: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateCompany: db commit: %w", err)
	}

	return &company, nil
}

func CreateLocation(
	ctx context.Context,
	db *sql.DB,
	input LocationCreate,
) (*Location, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateLocation: begin tx: %w", err)
	}
	defer tx.Rollback()

	var location Location
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO location (name, address, workspace_id)
		VALUES ($1, $2, $3)
		RETURNING id, name, address, workspace_id
		`,
		input.Name,
		input.Address,
		input.WorkspaceId,
	).Scan(
		location.Id,
		location.Name,
		location.Address,
		location.WorkspaceId,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateLocation: db insert: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateLocation: db commit: %w", err)
	}

	return &location, nil
}

func CreateTask(
	ctx context.Context,
	db *sql.DB,
	input TaskCreate,
) (*Task, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateTask: begin tx: %w", err)
	}
	defer tx.Rollback()

	var task Task
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO task (name, description, is_completed, location_id, company_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, description, is_completed, location_id, company_id
		`,
		input.Name,
		input.Description,
		input.IsCompleted,
		input.LocationId,
		input.CompanyId,
	).Scan(
		task.Id,
		task.Name,
		task.Description,
		task.IsCompleted,
		task.LocationId,
		task.CompanyId,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateTask: db insert: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateTask: db commit: %w", err)
	}

	return &task, nil
}

func CreateEmployment(
	ctx context.Context,
	db *sql.DB,
	input EmploymentCreate,
) (*Employment, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateEmployment: begin tx: %w", err)
	}
	defer tx.Rollback()

	var employment Employment
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO employment (user_id, company_id, contract_id, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, company_id, contract_id, role
		`,
		input.UserId,
		input.CompanyId,
		input.ContractId,
		input.Role,
	).Scan(
		employment.Id,
		employment.UserId,
		employment.CompanyId,
		employment.ContractId,
		employment.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateEmployment: db insert: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateEmployment: db commit: %w", err)
	}

	return &employment, nil
}

func CreateContract(
	ctx context.Context,
	db *sql.DB,
	input ContractCreate,
) (*Contract, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("CreateContract: begin tx: %w", err)
	}
	defer tx.Rollback()

	var contract Contract
	err = tx.QueryRowContext(
		ctx,
		`
		INSERT INTO contract (hourly_rate, unpaid_lunch_minutes)
		VALUES ($1, $2)
		RETURNING id, hourly_rate, unpaid_lunch_minutes
		`,
		input.HourlyRate,
		input.UnpaidLunchMinutes,
	).Scan(
		contract.Id,
		contract.HourlyRate,
		contract.UnpaidLunchMinutes,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateContract: db insert: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("CreateContract: db commit: %w", err)
	}

	return &contract, nil
}
