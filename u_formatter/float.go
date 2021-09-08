package u_formatter

import (
	"math"
)

func FormatFloat(v float64, decimals int) float64 {
	return math.Round(v*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func Float(v *float64) float64 {
	return FloatDefault(v, 0)
}

func FloatDefault(v *float64, dv float64) float64 {
	if v == nil {
		return dv
	}
	return *v
}
