package manage

type Project struct {
	Name string `json:"name"`
}

type CompanyCreate struct {
	Name      string `json:"name"`
	ProjectId int `json:"project_id"`
}

type LocationCreate struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	ProjectId int `json:"project_id"`
}

type TaskCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	LocationId  int `json:"location_id"`
	CompanyId   int `json:"company_id"`
}
