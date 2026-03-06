package pin

import (
	"test/internal/model"
	"time"
)

type ClockIn_R struct { // have to add lat long
	TaskId        int       `json:"task_id"`
	StartTs      *time.Time `json:"start_ts"`
}

type ClockOut_R struct { // have to add lat long
	EndTs        *time.Time `json:"end_ts"`
}

type SyncShift_R struct {
	RemoteId   *int      `json:"remote_id"`
	LocalId     int       `json:"local_id"`
	ProfileId   int       `json:"profile_id"`
	TaskId      int       `json:"task_id"`
	StartTs     time.Time `json:"start_ts"`
	EndTs      *time.Time `json:"end_ts"`
	SLatitude  *float64   `json:"s_latitude"`
	SLongitude *float64   `json:"s_longitude"`
	ELatitude  *float64   `json:"e_latitude"`
	ELongitude *float64   `json:"e_longitude"`
}

type ShiftOverview struct {
	Shift model.Shift `json:"shift"`
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
