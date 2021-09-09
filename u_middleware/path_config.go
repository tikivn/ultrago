package u_middleware

import (
	"regexp"
)

func NewDefaultPathConfig() PathConfig {
	return &defaultPathConfig{}
}

type defaultPathConfig struct{}

func (defaultPathConfig) PathCleanUp() map[*regexp.Regexp]string {
	return map[*regexp.Regexp]string{
		regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})\\/"): "/<id>/",
		regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})"):    "/<id>",
	}
}

func (defaultPathConfig) PathIgnored() map[string]bool {
	return map[string]bool{
		"/":            true,
		"/healthcheck": true,
		"/heartbeat":   true,
		"/metrics":     true,
	}
}
