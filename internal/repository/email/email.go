package email

import (
	"context"
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/pkg/common/logging"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

type Email struct {
	storage *storage.Storage
}

func NewEmail(storage *storage.Storage) *Email {
	return &Email{storage: storage}
}

func (r *Email) GetEmail(ctx context.Context, username string) (string, error) {
	var email string

	if err := r.storage.Client.QueryRow(ctx, "SELECT email FROM students WHERE username = $1", username).Scan(&email); err != nil {
		err = postgresql.ErrCreateQuery(err)
		return email, err
	}

	logging.WithFields(ctx, logging.StringField("email", email)).Info("get email")

	return email, nil
}
