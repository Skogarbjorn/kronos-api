package pin

import "time"

type ClockIn_R struct {
	//UserId       *int       `json:"user_id"`
	//CompanyId    *int       `json:"company_id"`
	EmploymentId *int       `json:"employment_id"`
	TaskId        int       `json:"task_id"`
	StartTs      *time.Time `json:"start_ts"`
}

type ClockOut_R struct {
	//UserId       *int       `json:"user_id"`
	//CompanyId    *int       `json:"company_id"`
	EmploymentId *int       `json:"employment_id"`
	EndTs        *time.Time `json:"end_ts"`
}

type Shift struct {
	Id int `json:"id"`
	EmploymentId int `json:"employment_id"`
	TaskId int `json:"task_id"`
	StartTs time.Time `json:"start_ts"`
	EndTs *time.Time `json:"end_ts"`
}

type Company struct {
	Id int `json:"id"`
	Name int `json:"name"`
}

type Location struct {
	Id int `json:"id"`
	Name int `json:"name"`
	Address int `json:"address"`
}

type Workspace struct {
	Id int `json:"id"`
	Name int `json:"name"`
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
