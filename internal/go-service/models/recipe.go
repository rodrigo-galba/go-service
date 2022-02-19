package models

import "time"

type Recipe struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"` // backtick annotation
	Tags        []string  `json:"tags"`
	Ingredients []string  `json:"ingredients"`
	PublishedAt time.Time `json:"publishedAt"`
}
