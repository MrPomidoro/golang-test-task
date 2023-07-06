package validate

import "regexp"

func IsEmail(email string) bool {
	match, err := regexp.Match(`^(?i)[a-z0-9._%+-]+@(?:[a-z0-9-]+\.)+[a-z]{2,}$`, []byte(email))
	if err != nil {
		return false
	}
	return match
}
