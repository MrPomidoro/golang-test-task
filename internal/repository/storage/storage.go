package storage

import (
	"github.com/Masterminds/squirrel"
	"github.com/golang-test-task/pkg/databases/postgresql"
)

type Storage struct {
	client       postgresql.Client
	queryBuilder squirrel.StatementBuilderType
}

func NewStorage(client postgresql.Client, queryBuilder squirrel.StatementBuilderType) *Storage {
	return &Storage{
		client:       client,
		queryBuilder: queryBuilder,
	}
}
