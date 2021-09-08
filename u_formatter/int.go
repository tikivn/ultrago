package u_formatter

func Int(v *int64) int64 {
	return IntDefault(v, 0)
}

func IntDefault(v *int64, dv int64) int64 {
	if v == nil {
		return dv
	}
	return *v
}
