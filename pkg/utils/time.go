package utils

import "time"

func GetMoscowTime(timeString string) (time.Time, error) {
	loc, err := getLocation("Europe/Moscow")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}, err
	}

	// Convert the time to Moscow timezone
	return t.In(loc), nil
}

func getLocation(location string) (*time.Location, error) {
	return time.LoadLocation(location)
}
