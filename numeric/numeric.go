package numeric

import (
	"math"
)

func FormatFloat(v float64, decimals int) float64 {
	return math.Round(v*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
