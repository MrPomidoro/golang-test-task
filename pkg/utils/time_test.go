package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetMoscowTime(t *testing.T) {
	// Get Moscow location for test
	moscowLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatalf("Failed to load Moscow location: %v", err)
	}

	// Define test cases
	testCases := []struct {
		name         string
		timeString   string
		expectedTime time.Time
		expectedErr  bool
	}{
		{
			name:         "Valid time",
			timeString:   "2023-07-06T12:00:00Z",                             // 12:00 in UTC
			expectedTime: time.Date(2023, 7, 6, 15, 0, 0, 0, moscowLocation), // should be 15:00 in Moscow
			expectedErr:  false,
		},
		{
			name:         "Invalid time",
			timeString:   "Invalid time string",
			expectedTime: time.Time{},
			expectedErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualTime, err := GetMoscowTime(tc.timeString)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// Compare the actual time in Moscow with the expected time
				assert.Equal(t, tc.expectedTime, actualTime)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	testCases := []struct {
		name        string
		location    string
		expectedErr bool
	}{
		{
			name:        "Valid location",
			location:    "Europe/Moscow",
			expectedErr: false,
		},
		{
			name:        "Invalid location",
			location:    "Invalid location",
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := getLocation(tc.location)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
