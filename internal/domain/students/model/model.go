package model

import (
	"time"
)

type DomainStudent struct {
	ID          string     `json:"id"`
	FullName    string     `json:"full_name"`
	GroupNum    string     `json:"group_num"`
	Email       string     `json:"email"`
	Username    string     `json:"username"`
	VerifyEmail bool       `json:"verify_email"`
	CreateAT    time.Time  `json:"create_at"`
	UpdateAT    *time.Time `json:"update_at"`
}

// NewStudent create new student
func NewStudent(id, fullName string, groupNum string, email string, username string) *DomainStudent {
	return &DomainStudent{
		ID:       id,
		FullName: fullName,
		GroupNum: groupNum,
		Email:    email,
		Username: username,
	}
}
