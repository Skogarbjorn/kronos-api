package manage

type Role string
const (
	RoleManager  Role = "manager"
	RoleEmployee Role = "employee"
)

type Workspace struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type WorkspaceCreate struct {
	Name string `json:"name"`
}

type Company struct {
	Id   		int `json:"id"`
	Name        string `json:"name"`
	WorkspaceId int `json:"workspace_id"`
}

type CompanyCreate struct {
	Name        string `json:"name"`
	WorkspaceId int `json:"workspace_id"`
}

type Location struct {
	Id   		int `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	WorkspaceId int `json:"workspace_id"`
}

type LocationCreate struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	WorkspaceId int `json:"workspace_id"`
}

type Task struct {
	Id   		int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	LocationId  int `json:"location_id"`
	CompanyId   int `json:"company_id"`
}

type TaskCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	LocationId  int `json:"location_id"`
	CompanyId   int `json:"company_id"`
}

type Employment struct {
	Id   		int `json:"id"`
	ProfileId      int `json:"profile_id"`
	CompanyId   int `json:"company_id"`
	ContractId  int `json:"contract_id"`
	Role 	    Role `json:"role"`
}

type EmploymentCreate struct {
	ProfileId 	   int `json:"profile_id"`
	CompanyId  int `json:"company_id"`
	ContractId int `json:"contract_id"`
	Role 	   Role `json:"role"`
}

type Contract struct {
	Id          	   int `json:"id"`
	HourlyRate         int `json:"hourly_rate"`
	UnpaidLunchMinutes int `json:"unpaid_lunch_minutes"`
}

type ContractCreate struct {
	HourlyRate         int `json:"hourly_rate"`
	UnpaidLunchMinutes int `json:"unpaid_lunch_minutes"`
}
