package model

import (
	"time"
)

type Task struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Cost        int        `json:"cost"`
	CreateAT    time.Time  `json:"create_at"`
	UpdateAT    *time.Time `json:"update_at"`
}

func NewTask(id string, description string, cost int, createAT time.Time) Task {
	return Task{
		ID:          id,
		Description: description,
		Cost:        cost,
		CreateAT:    createAT,
		UpdateAT:    nil,
	}
}
