package manage

import (
	"test/internal/model"
	"time"
)

type WorkspaceCreate struct {
	Name string `json:"name"`
}

type CompanyCreate struct {
	Name        string `json:"name"`
	WorkspaceId int `json:"workspace_id"`
}

type LocationCreate struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	WorkspaceId int `json:"workspace_id"`
}

type TaskCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	LocationId  int `json:"location_id"`
	CompanyId   int `json:"company_id"`
}

type EmploymentCreate struct {
	ProfileId 	   int `json:"profile_id"`
	CompanyId  int `json:"company_id"`
	ContractId int `json:"contract_id"`
	Role 	   model.Role `json:"role"`
}

type ContractCreate struct {
	HourlyRate         int `json:"hourly_rate"`
	UnpaidLunchMinutes int `json:"unpaid_lunch_minutes"`
}

type WorkspacePatch struct {
	Name *string `json:"name"`
}

type CompanyPatch struct {
	Name *string `json:"name"`
}

type LocationPatch struct {
    Name    *string `json:"name"`
    Address *string `json:"address"`
}

type TaskPatch struct {
    Name        *string `json:"name"`
    Description *string `json:"description"`
    IsCompleted *bool `json:"is_completed"`
    LocationId  *int `json:"location_id"`
}

type EmploymentPatch struct {
    ContractId *int `json:"contract_id"`
    Role       *model.Role `json:"role"`
}

type ContractPatch struct {
    HourlyRate         *int `json:"hourly_rate"`
    UnpaidLunchMinutes *int `json:"unpaid_lunch_minutes"`
}

type ShiftPatch struct {
	StartTs *time.Time `json:"start_ts"`
	EndTs   *time.Time `json:"end_ts"`
	TaskId  *int       `json:"task_id"`
}

type ProfilePatch struct {
	KT        *string `json:"kt"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
