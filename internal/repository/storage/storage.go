package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

type Storage struct {
	client       postgresql.Client
	queryBuilder sq.StatementBuilderType
}

func NewStorage(client postgresql.Client) *Storage {
	return &Storage{
		client:       client,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
