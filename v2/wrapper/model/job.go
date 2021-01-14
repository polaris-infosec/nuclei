package model

import "time"

type Target struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	URL         string `json:"url"`
	Description string `json:"description"`
}

type Vulnerability struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Description string `json:"description"`
	Name        string `json:"name"`
	Matched     string `json:"matched"`
	Template    string `json:"template"`
	Type        string `json:"type"`
	Host        string `json:"host"`
	Severity    string `json:"severity"`

	JobID string `json:"jobId"`
}

type Job struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Description string  `json:"description"`
	TargetID    string  `json:"targetId"`
	Target      *Target `json:"target,omitempty"`
}
