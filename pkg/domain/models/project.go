package models

type Project struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	StartedDate  string `json:"started_date"`
	FinishedDate string `json:"finished_date"`
	Status       string `json:"status"`
	TeamLeadID   string `json:"teamlead_id"`
	Attachment   string `json:"attachment"`
}
