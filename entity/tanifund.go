package entity

import "time"

// TaniFund defines the TaniFund's project API structure.
type TaniFund struct {
	Data *Data `json:"data"`
}

// Data defines the outermost JSON.
type Data struct {
	Items []*Project `json:"items"`
}

// Project defines the detail of TaniFund's project.
type Project struct {
	ID               string         `json:"id"`
	Title            string         `json:"title"`
	PricePerUnit     int            `json:"pricePerUnit"`
	MaxUnit          int            `json:"maxUnit"`
	StartAt          time.Time      `json:"startAt"`
	EndAt            time.Time      `json:"endAt"`
	PublishedAt      time.Time      `json:"publishedAt"`
	InterestTarget   int            `json:"interestTarget"`
	URLSlug          string         `json:"urlSlug"`
	ProjectStatus    *ProjectStatus `json:"projectStatus"`
	HumanPublishedAt time.Time      `json:"humanPublishedAt"`
	ProjectLink      string         `json:"projectLink"`
	TargetFund       string         `json:"targetFund"`
	Tenor            int            `json:"tenor"`
}

// ProjectStatus defines project's status.
type ProjectStatus struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
