package pin

import "time"

type ClockIn_R struct {
	TaskId        int       `json:"task_id"`
	StartTs      *time.Time `json:"start_ts"`
}

type ClockOut_R struct {
	EndTs        *time.Time `json:"end_ts"`
}

type Shift struct {
	Id        int       `json:"id"`
	ProfileId int       `json:"profile_id"`
	TaskId    int       `json:"task_id"`
	StartTs   time.Time `json:"start_ts"`
	EndTs    *time.Time `json:"end_ts"`
}

type Company struct {
	Id           int    `json:"id"`
	WorkspaceId *int    `json:"workspace_id"`
	Name 		 string `json:"name"`
}

type Location struct {
	Id           int    `json:"id"`
	WorkspaceId *int    `json:"workspace_id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
}

type Workspace struct {
	Id int      `json:"id"`
	Name string `json:"name"`
}

type Role string
const (
	RoleAdmin   Role = "admin"
	RoleManager Role = "manager"
	RoleWorker  Role = "worker"
)

type Employment struct {
	Id         int        `json:"id"`
	ProfileId  int        `json:"profile_id"`
	CompanyId  int        `json:"company_id"`
	ContractId int        `json:"contract_id"`
	Role       Role       `json:"role"`
	StartDate  time.Time  `json:"start_date"`
	EndDate    time.Time  `json:"end_date"`
}

type Contract struct {
	Id                 int `json:"id"`
	HourlyRate         int `json:"hourly_rate"`
	UnpaidLunchMinutes int `json:"unpaid_lunch_minutes"`
}

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
	CompanyId   int    `json:"company_id"`
	LocationId  int    `json:"location_id"`
}

type ShiftOverview struct {
	Shift Shift `json:"shift"`
	Workspace string `json:"workspace"`
	Company string `json:"company"`
	Location string `json:"location"`
	Task string `json:"task"`
}

type ShiftHistoryFilter struct {
	WorkspaceId *int       `json:"workspace_id"`
	CompanyId   *int       `json:"company_id"`
	LocationId  *int       `json:"location_id"`
	TaskId      *int       `json:"task_id"`
	StartTs     *time.Time `json:"start_ts"`
	EndTs       *time.Time `json:"end_ts"`
}

type EmploymentDetailed struct {
	Workspace Workspace
	Company Company
	Employment Employment
	Contract Contract
}
