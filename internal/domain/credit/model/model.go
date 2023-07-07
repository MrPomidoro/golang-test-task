package model

import (
	"github.com/google/uuid"
	"time"
)

type Credit struct {
	ID          uuid.UUID  `json:"id"`
	StudentID   uuid.UUID  `json:"student_id"`
	CreditLimit int        `json:"credit_limit"`
	CreatedAT   time.Time  `json:"create_at"`
	UpdateAT    *time.Time `json:"update_at"`
}

func NewCreditModel(id, studentID uuid.UUID, creditLimit int, createAT time.Time, updateAT *time.Time) Credit {
	return Credit{
		ID:          id,
		StudentID:   studentID,
		CreditLimit: creditLimit,
		CreatedAT:   createAT,
		UpdateAT:    updateAT,
	}
}
