package model

import (
	"github.com/google/uuid"
	"time"
)

type Student struct {
	ID          uuid.UUID  `json:"id"`
	FullName    string     `json:"full_name"`
	GroupNum    string     `json:"group_num"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	VerifyEmail bool       `json:"verify_email"`
	CreateAT    time.Time  `json:"create_at"`
	UpdateAT    *time.Time `json:"update_at"`
}

// NewStudent create new student
func NewStudent(id uuid.UUID, fullName, groupNum, email, username string, verifyEmail bool, CreateAT time.Time, UpdateAT *time.Time) Student {
	return Student{
		ID:          id,
		FullName:    fullName,
		GroupNum:    groupNum,
		Email:       email,
		Username:    username,
		VerifyEmail: verifyEmail,
		CreateAT:    CreateAT,
		UpdateAT:    UpdateAT,
	}
}

type CreateStudent struct {
	ID          string     `json:"id"`
	FullName    string     `json:"full_name"`
	GroupNum    string     `json:"group_num"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	VerifyEmail bool       `json:"verify_email"`
	CreateAT    time.Time  `json:"create_at"`
	UpdateAT    *time.Time `json:"update_at"`
}

func NewCreateStudent(id, fullName, groupNum, email, username string, verifyEmail bool, CreateAT time.Time, UpdateAT *time.Time) CreateStudent {
	return CreateStudent{
		ID:          id,
		FullName:    fullName,
		GroupNum:    groupNum,
		Email:       email,
		Username:    username,
		VerifyEmail: verifyEmail,
		CreateAT:    CreateAT,
		UpdateAT:    UpdateAT,
	}
}
