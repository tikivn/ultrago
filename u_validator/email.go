package u_validator

import "regexp"

func VerifyEmail(e string) bool {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9_\.\+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-\.]+`)
	return emailRegex.MatchString(e)
}
