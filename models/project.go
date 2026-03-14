package models

import (
	"time"
)

type ProjectType string

const (
	UIUX   ProjectType = "UI/UX"
	Web    ProjectType = "Web"
	CLI    ProjectType = "CLI"
	GUI    ProjectType = "GUI"
	Script ProjectType = "Script"
	Core   ProjectType = "Core"
)

type Project struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Type        ProjectType `json:"type"`
	Deadline    time.Time   `json:"deadline"`
	CreatedAt   time.Time   `json:"created_at"`
	IsCompleted bool        `json:"is_completed"`
}

func (p *Project) Overdue() bool {
	return time.Now().After(p.Deadline) && !p.IsCompleted
}

func (p *Project) IsNew() bool {
	return time.Since(p.CreatedAt) < 24*time.Hour
}

// Health returns a value from 0 to 1 based on time remaining vs total time.
// Since we don't have a start date, we'll assume a 7-day default window if created recently,
// or just measure against the deadline.
func (p *Project) Health() float64 {
	if p.IsCompleted {
		return 1.0
	}
	remaining := p.Deadline.Sub(time.Now())
	if remaining < 0 {
		return 0.0
	}
	// Simple heuristic: if deadline is > 7 days away, health is 1.0. 
	// As it gets closer, it drops.
	score := remaining.Hours() / (24 * 7) 
	if score > 1.0 {
		return 1.0
	}
	return score
}
