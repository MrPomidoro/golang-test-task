package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		want    string
		wantErr error
	}{
		{
			name:    "valid email",
			email:   "Test.Email@example.com",
			want:    "test.email@example.com",
			wantErr: nil,
		},
		{
			name:    "valid email, already lowercase",
			email:   "test.email@example.com",
			want:    "test.email@example.com",
			wantErr: nil,
		},
		{
			name:    "invalid email, no @",
			email:   "test.email.com",
			want:    "",
			wantErr: invalidEmail,
		},
		{
			name:    "invalid email, two @",
			email:   "test@Email@example.com",
			want:    "",
			wantErr: invalidEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateEmail(tt.email)
			if err != nil {
				assert.Error(t, err, "invalid email")
			} else {
				assert.NoError(t, tt.wantErr, "invalid email")
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
