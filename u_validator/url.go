package u_validator

import (
	"fmt"
	"net/url"
)

func VerifyURL(str string) (*url.URL, error) {
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return nil, err
	}
	if u != nil {
		if u.Scheme == "" {
			return nil, fmt.Errorf("parse failed: missing schema")
		}
		if u.Host == "" {
			return nil, fmt.Errorf("parse failed: missing host")
		}
	}
	return u, nil
}
