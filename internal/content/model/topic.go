package model

type Topic struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"`
	Tags        []string `json:"tags"`
}
