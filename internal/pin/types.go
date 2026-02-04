package pin

import "time"

type ClockIn_R struct {
	//UserId       *int       `json:"user_id"`
	//CompanyId    *int       `json:"company_id"`
	EmploymentId *int       `json:"employment_id"`
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
	StartTs time.Time `json:"start_ts"`
	EndTs *time.Time `json:"end_ts"`
}
