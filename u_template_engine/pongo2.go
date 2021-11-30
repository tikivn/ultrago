package u_template_engine

import (
	"regexp"

	"github.com/flosch/pongo2"
)

func Parse(template string, placeholders map[string]interface{}) (string, error) {
	if len(placeholders) == 0 {
		return template, nil
	}

	regex, err := regexp.Compile("\\{\\{[^}]+\\}\\}")
	if err != nil {
		return "", err
	}
	for i := 0; i <= 10; i++ { // allow up to 10 level nested
		if hasPlaceHolder := regex.MatchString(template); !hasPlaceHolder {
			return template, nil
		}

		engine, err := pongo2.FromString(template)
		if err != nil {
			return "", err
		}

		template, err = engine.Execute(placeholders)
		if err != nil {
			return "", err
		}
	}
	return template, nil
}
