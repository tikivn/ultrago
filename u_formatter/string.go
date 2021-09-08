package u_formatter

func String(v *string) string {
	return StringDefault(v, "")
}

func StringDefault(v *string, dv string) string {
	if v == nil {
		return dv
	}
	return *v
}
