package model

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Description string     `json:"description"`
	Cost        int        `json:"cost"`
	CreateAT    time.Time  `json:"create_at"`
	UpdateAT    *time.Time `json:"update_at"`
}

func NewTask(id uuid.UUID, description string, cost int, createAT time.Time, updateAT *time.Time) Task {
	return Task{
		ID:          id,
		Description: description,
		Cost:        cost,
		CreateAT:    createAT,
		UpdateAT:    updateAT,
	}
}
