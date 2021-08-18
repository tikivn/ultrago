package placeholder

func Render(text string, params map[string]interface{}) (string, error) {
	template, err := new(text)
	if err != nil {
		return "", err
	}
	return template.Render(params)
}

func RenderWithDefault(text string, params map[string]interface{}, defaultText string) string {
	value, err := Render(text, params)
	if err != nil {
		return defaultText
	}
	return value
}
