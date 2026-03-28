package model

type QuizQuestion struct {
	ID      string   `json:"id"`
	Prompt  string   `json:"prompt"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}

type LessonStep struct {
	ID      string `json:"id"`
	Order   int    `json:"order"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Example string `json:"example"`
	Tip     string `json:"tip"`
}

type LessonOutcome struct {
	Text string `json:"text"`
}

type Lesson struct {
	ID               string          `json:"id"`
	TopicID          string          `json:"topic_id"`
	Title            string          `json:"title"`
	Summary          string          `json:"summary"`
	Content          []string        `json:"content"`
	EstimatedMinutes int             `json:"estimated_minutes"`
	Quiz             []QuizQuestion  `json:"quiz"`
	Steps            []LessonStep    `json:"steps"`
	Outcomes         []LessonOutcome `json:"outcomes"`
}
