package placeholder

import (
	"github.com/flosch/pongo2"
)

func new(v string) (*template, error) {
	templated, err := pongo2.FromString(v)
	if err != nil {
		return nil, err
	}
	return &template{
		raw:      v,
		rendered: templated,
	}, nil
}

type template struct {
	raw      string
	rendered *pongo2.Template
}

// Render renders the real value given map of data
func (t *template) Render(ctx map[string]interface{}) (string, error) {
	return t.rendered.Execute(ctx)
}

// RenderOrEmpty renders the real value, fallback to empty string when error occurs
func (t *template) RenderOrEmpty(ctx map[string]interface{}) string {
	v, err := t.rendered.Execute(ctx)
	if err != nil {
		return ""
	}
	return v
}

// MarshalText encodes the wrapped Template into a textual form.
//
// This makes it encodable as JSON, YAML, XML, and more.
func (t *template) MarshalText() ([]byte, error) {
	return []byte(t.raw), nil
}

// UnmarshalText decodes text and replaces the wrapped Template with it.
//
// This makes it decodable from JSON, YAML, XML, and more.
func (t *template) UnmarshalText(b []byte) error {
	templated, err := pongo2.FromBytes(b)
	if err != nil {
		return err
	}
	t.raw = string(b)
	t.rendered = templated
	return nil
}
