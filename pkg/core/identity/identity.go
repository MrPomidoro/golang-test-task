package identity

import (
	"github.com/google/uuid"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateUUIDv4() uuid.UUID {
	return uuid.New()
}
