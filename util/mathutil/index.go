package mathutil

import (
	"fmt"
	"math"
)

func Round(x, by float64) float64 {
	roundBy := math.Pow(10, by)
	return math.Round(x*roundBy) / roundBy
}

func RoundAndString(x, by float64) string {
	return fmt.Sprint(Round(x, by))
}
