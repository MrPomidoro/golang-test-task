package students

import (
	"database/sql"
	"github.com/golang-test-task/internal/domain/students/model"
	"github.com/golang-test-task/pkg/utils"
	"github.com/google/uuid"
	"time"
)

type RepositoryStudent struct {
	ID          uuid.UUID    `json:"id"`
	FullName    string       `json:"full_name"`
	GroupNum    string       `json:"group_num"`
	Email       string       `json:"email"`
	Username    string       `json:"username"`
	VerifyEmail bool         `json:"verify_email"`
	CreateAT    time.Time    `json:"create_at"`
	UpdateAT    sql.NullTime `json:"update_at"`
}

func (s *RepositoryStudent) ToDomain() model.Student {
	var updateAT *time.Time
	if s.UpdateAT.Valid {
		updateAT = utils.Pointer(s.UpdateAT.Time)
	}

	return model.NewStudent(s.ID, s.FullName, s.GroupNum, s.Email, s.Username, s.VerifyEmail, s.CreateAT, updateAT)
}
