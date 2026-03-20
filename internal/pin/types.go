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

type EditRequest_R struct {
	ShiftId  int       `json:"shift_id"`
	TaskId  *int       `json:"task_id"`
	StartTs *time.Time `json:"start_ts"`
	EndTs   *time.Time `json:"end_ts"`
	Reason   string    `json:"reason"`
}

type ShiftOverview struct {
	Shift    model.Shift    `json:"shift"`
	Location model.Location `json:"location"`
	Task     model.Task     `json:"task"`
}

type ShiftHistoryResponse struct {
	Shifts   []model.Shift     `json:"shifts"`
	Metadata   HistoryMetadata `json:"metadata"`
}

type HistoryMetadata struct {
	Month     int  `json:"month"`
	Year      int  `json:"year"`
	HasMore   bool `json:"has_more"`
	NextMonth int  `json:"next_month"`
	NextYear  int  `json:"next_year"`
}

type EmploymentDetailed struct {
	Workspace  model.Workspace
	Company    model.Company
	Employment model.Employment
	Contract   model.Contract
}
