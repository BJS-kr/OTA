package collector

import (
	"time"

	"github.com/google/uuid"
)

type CollectionRun struct {
	ID           uuid.UUID
	StartedAt    time.Time
	CompletedAt  *time.Time
	Status       string
	ErrorMessage *string
	RawResponse  *string
}

type ContextItem struct {
	ID              uuid.UUID
	CollectionRunID uuid.UUID
	Category        string   `json:"category"`
	Rank            int      `json:"rank"`
	Topic           string   `json:"topic"`
	Summary         string   `json:"summary"`
	Sources         []string `json:"sources"`
}

type CollectionResult struct {
	Run   CollectionRun
	Items []ContextItem
}
