package jfloat

import "math"

func Decimal2Number(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}