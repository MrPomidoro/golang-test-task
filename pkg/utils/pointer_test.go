package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "int value",
			input:    42,
			expected: 42,
		},
		{
			name:     "string value",
			input:    "test",
			expected: "test",
		},
		{
			name:     "bool value",
			input:    true,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Pointer(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
