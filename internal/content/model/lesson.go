package model

type QuizQuestion struct {
	ID      string   `json:"id"`
	Prompt  string   `json:"prompt"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}

type Lesson struct {
	ID               string         `json:"id"`
	TopicID          string         `json:"topic_id"`
	Title            string         `json:"title"`
	Summary          string         `json:"summary"`
	Content          []string       `json:"content"`
	EstimatedMinutes int            `json:"estimated_minutes"`
	Quiz             []QuizQuestion `json:"quiz"`
}
