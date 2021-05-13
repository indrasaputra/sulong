package entity

import "time"

// TaniFund defines the TaniFund's project API structure.
type TaniFund struct {
	Data Data `json:"data"`
}

// Data defines the outermost JSON.
type Data struct {
	Items []Project `json:"items"`
}

// Project defines the detail of TaniFund's project.
type Project struct {
	ID                string        `json:"id"`
	Title             string        `json:"title"`
	ShortDescription  string        `json:"shortDescription"`
	PricePerUnit      int           `json:"pricePerUnit"`
	StartAt           time.Time     `json:"startAt"`
	EndAt             time.Time     `json:"endAt"`
	PublishedAt       time.Time     `json:"publishedAt"`
	CutoffAt          time.Time     `json:"cutoffAt"`
	InterestPessimist int           `json:"interestPessimist"`
	InterestTarget    int           `json:"interestTarget"`
	InterestOptimist  int           `json:"interestOptimist"`
	Grade             string        `json:"grade"`
	URLSlug           string        `json:"urlSlug"`
	IsHidden          int           `json:"isHidden"`
	Projectstatus     ProjectStatus `json:"projectStatus"`
	ReturnPeriod      ReturnPeriod  `json:"returnPeriod"`
}

// ProjectStatus defines project's status.
type ProjectStatus struct {
	Description string `json:"description"`
}

// ReturnPeriod defines return period.
type ReturnPeriod struct {
	Description string `json:"description"`
}
