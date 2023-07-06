package identity

import (
	"github.com/google/uuid"
	"testing"
)

func TestGenerateUUIDv4String(t *testing.T) {
	g := NewGenerator()

	// генерируем две строки UUID
	uuid1 := g.GenerateUUIDv4String()
	uuid2 := g.GenerateUUIDv4String()

	if uuid1 == "" {
		t.Errorf("GenerateUUIDv4String() returned empty string")
	}
	if uuid2 == "" {
		t.Errorf("GenerateUUIDv4String() returned empty string")
	}

	_, err := uuid.Parse(uuid1)
	if err != nil {
		t.Errorf("GenerateUUIDv4String() returned invalid UUID: %s", uuid1)
	}
	_, err = uuid.Parse(uuid2)
	if err != nil {
		t.Errorf("GenerateUUIDv4String() returned invalid UUID: %s", uuid2)
	}
	if uuid1 == uuid2 {
		t.Errorf("GenerateUUIDv4String() returned the same UUID twice: %s", uuid1)
	}
}
