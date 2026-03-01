package pin

import (
	"test/internal/model"
	"time"
)

type ClockIn_R struct {
	TaskId        int       `json:"task_id"`
	StartTs      *time.Time `json:"start_ts"`
}

type ClockOut_R struct {
	EndTs        *time.Time `json:"end_ts"`
}

type ShiftOverview struct {
	Shift model.Shift `json:"shift"`
	Workspace model.Workspace `json:"workspace"`
	Company model.Company `json:"company"`
	Location model.Location `json:"location"`
	Task model.Task `json:"task"`
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
	Workspace model.Workspace
	Company model.Company
	Employment model.Employment
	Contract model.Contract
}
