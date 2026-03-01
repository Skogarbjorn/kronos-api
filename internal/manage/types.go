package manage

import "test/internal/model"

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
