package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmail(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "valid email",
			email:    "test.email+alex@leetcode.com",
			expected: true,
		},
		{
			name:     "valid email 2",
			email:    "email@domain.com",
			expected: true,
		},
		{
			name:     "invalid email, no @",
			email:    "test.email.leet.alex.com",
			expected: false,
		},
		{
			name:     "invalid email, two @",
			email:    "test@e.mail@leetcode.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, IsEmail(tc.email))
		})
	}
}
