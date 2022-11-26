package web

type JobListCreateRequest struct {
	Description string `json:"description"`
	Location    string `json:"location"`
	FullTime    string `json:"full_time"`
	Page        string `json:"page"`
}
