package forecast

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat"
)

func Forecast(x float64, knownY []float64, knownX []float64) (y float64, err error) {
	if len(knownX) != len(knownY) {
		return 0, fmt.Errorf("knownY and knownY must be in same length")
	}

	a, b := calculateWeights(knownY, knownX)
	y = a + b*x
	if math.IsNaN(y) || math.IsInf(y, 0) {
		y = 0
	}
	return
}

func calculateWeights(knownY []float64, knownX []float64) (a float64, b float64) {
	var (
		sxy float64
		sx2 float64
	)

	meanX := stat.Mean(knownX, nil)
	meanY := stat.Mean(knownY, nil)

	for i, x := range knownX {
		y := knownY[i]
		xc := x - meanX
		yc := y - meanY
		sxy += xc * yc
		sx2 += math.Pow(xc, 2)
	}

	b = sxy / sx2
	a = meanY - b*meanX
	return a, b
}
