package dto

import (
	"github.com/golang-test-task/pkg/utils"
	"time"
)

type CreateStudentDTO struct {
	FullName string `json:"full_name"`
	GroupNum string `json:"group_num"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewCreateDTO(fullName, groupNum, email, username string) CreateStudentDTO {
	return CreateStudentDTO{
		FullName: fullName,
		GroupNum: groupNum,
		Email:    email,
		Username: username,
	}
}

type UpdateStudentDTO struct {
	FullName    string    `json:"full_name"`
	GroupNum    string    `json:"group_num"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	VerifyEmail bool      `json:"verify_email"`
	UpdateAT    time.Time `json:"update_at"`
}

func NewUpdateDTO(fullName, groupNum, email, username string, verifyEmail bool, updateAT string) (UpdateStudentDTO, error) {
	upt, err := utils.GetMoscowTime(updateAT)
	if err != nil {
		return UpdateStudentDTO{}, err
	}

	return UpdateStudentDTO{
		FullName:    fullName,
		GroupNum:    groupNum,
		Email:       email,
		Username:    username,
		VerifyEmail: verifyEmail,
		UpdateAT:    upt,
	}, nil
}
