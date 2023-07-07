package email

import (
	"context"
	"fmt"
	"github.com/golang-test-task/internal/repository"
	"github.com/golang-test-task/internal/repository/storage"
	"github.com/golang-test-task/pkg/common/logging"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

type Email struct {
	storage storage.Storage
}

func NewEmail() *Email {
	return &Email{}
}

func (r *Email) GetEmail(ctx context.Context, username string) (string, error) {
	var email string

	sql, args, err := r.storage.QueryBuilder.
		Select(repository.StudentTable).Columns("email").Where("username = $1", username).ToSql()

	if err != nil {
		err = postgresql.ErrCreateQuery(err)
		return email, err
	}

	logging.WithFields(ctx, logging.StringField("sql", sql)).Info("get email")

	logging.L(ctx).Info(fmt.Sprintf("args %v", args))

	if err := r.storage.Client.QueryRow(ctx, sql, args...).Scan(&email); err != nil {
		err = postgresql.ErrCreateQuery(err)
		return email, err
	}

	return email, nil
}
