package formSantanizer

import "regexp"

func IsValidUsername(username string) bool {
	// Check if the username contains only letters and numbers
	match, _ := regexp.MatchString("/^(?=[a-zA-Z0-9._]{6,25}$)(?!.*[_.]{2})[^_.].*[^_.]$/", username)
	if !match {
		return false
	}
	// If all checks pass, the username is valid
	return true
}

func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	return regex.MatchString(email)
}
func IsPasswordValid(password string) bool {
	if len(password) >= 6 {
		return true
	} else {
		return false
	}
}
