package email

import "regexp"

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9_\.\+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-\.]+`)
	return emailRegex.MatchString(e)
}
