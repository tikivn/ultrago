package u_middleware

import (
	"regexp"
)

type PathConfig interface {
	PathCleanUp() map[*regexp.Regexp]string
	PathIgnored() map[string]bool
}
