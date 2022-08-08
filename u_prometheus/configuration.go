package u_prometheus

import (
	"net/http"
	"regexp"
)

func NewDefaultIncomingHttpConfig() *HttpConfig {
	return &HttpConfig{
		PathCleanUpMap: map[*regexp.Regexp]string{
			regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})\\/"): "/<id>/",
			regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})"):    "/<id>",
		},
		PathIgnoredMap: map[string]bool{
			"/":            true,
			"/healthcheck": true,
			"/heartbeat":   true,
			"/metrics":     true,
		},
		StatusIgnoredMap: make(map[int]bool, 0),
		MethodIgnoreMap: map[string]bool{
			http.MethodOptions: true,
		},
	}
}

func NewDefaultOutgoingHttpConfig() *HttpConfig {
	return &HttpConfig{
		PathCleanUpMap: map[*regexp.Regexp]string{
			regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})\\/"): "/<id>/",
			regexp.MustCompile("\\/([0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12})"):    "/<id>",
		},
		PathIgnoredMap:   make(map[string]bool, 0),
		StatusIgnoredMap: make(map[int]bool, 0),
		MethodIgnoreMap:  make(map[string]bool, 0),
	}
}

func NewHttpConfig(
	pathCleanUpMap map[*regexp.Regexp]string,
	pathIgnoredMap map[string]bool,
	statusIgnoredMap map[int]bool,
	methodIgnoredMap map[string]bool,
) *HttpConfig {
	return &HttpConfig{
		PathCleanUpMap:   pathCleanUpMap,
		PathIgnoredMap:   pathIgnoredMap,
		StatusIgnoredMap: statusIgnoredMap,
		MethodIgnoreMap:  methodIgnoredMap,
	}
}

func NewEmptyHttpConfig() *HttpConfig {
	return &HttpConfig{
		PathCleanUpMap:   make(map[*regexp.Regexp]string, 0),
		PathIgnoredMap:   make(map[string]bool, 0),
		StatusIgnoredMap: make(map[int]bool, 0),
		MethodIgnoreMap:  make(map[string]bool, 0),
	}
}

type HttpConfig struct {
	PathCleanUpMap   map[*regexp.Regexp]string
	PathIgnoredMap   map[string]bool
	StatusIgnoredMap map[int]bool
	MethodIgnoreMap  map[string]bool
}

func (c *HttpConfig) WithHttpConfig(conf HttpConfig) *HttpConfig {
	c.WithPathCleanUp(conf.PathCleanUpMap)
	c.WithPathIgnored(conf.PathIgnoredMap)
	c.WithStatusIgnored(conf.StatusIgnoredMap)
	c.WithMethodIgnored(conf.MethodIgnoreMap)
	return c
}

func (c *HttpConfig) WithPathCleanUp(pathCleanUpMap map[*regexp.Regexp]string) *HttpConfig {
	if c.PathCleanUpMap == nil {
		c.PathCleanUpMap = make(map[*regexp.Regexp]string, 0)
	}
	for key, value := range pathCleanUpMap {
		c.PathCleanUpMap[key] = value
	}
	return c
}

func (c *HttpConfig) WithPathIgnored(pathIgnoredMap map[string]bool) *HttpConfig {
	if c.PathIgnoredMap == nil {
		c.PathIgnoredMap = make(map[string]bool, 0)
	}
	for key, value := range pathIgnoredMap {
		c.PathIgnoredMap[key] = value
	}
	return c
}

func (c *HttpConfig) WithStatusIgnored(statusIgnoredMap map[int]bool) *HttpConfig {
	if c.StatusIgnoredMap == nil {
		c.StatusIgnoredMap = make(map[int]bool, 0)
	}
	for key, value := range statusIgnoredMap {
		c.StatusIgnoredMap[key] = value
	}
	return c
}

func (c *HttpConfig) WithMethodIgnored(methodIgnoredMap map[string]bool) *HttpConfig {
	if c.MethodIgnoreMap == nil {
		c.MethodIgnoreMap = make(map[string]bool, 0)
	}
	for key, value := range methodIgnoredMap {
		c.MethodIgnoreMap[key] = value
	}
	return c
}
