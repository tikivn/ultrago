package collection

func IndexInt(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []int, t int) bool {
	return IndexInt(vs, t) >= 0
}
