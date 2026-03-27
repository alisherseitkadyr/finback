package model

import "time"

type User struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Email            string             `json:"email"`
	ExperienceLevel  string             `json:"experience_level"`
	Goals            []string           `json:"goals"`
	PreferredTopics  []string           `json:"preferred_topics"`
	Assessment       map[string]float64 `json:"assessment"`
	CompletedLessons []string           `json:"completed_lessons"`
	Progress         map[string]float64 `json:"progress"`
	TotalPoints      int                `json:"total_points"`
	CreatedAt        time.Time          `json:"created_at"`
}
