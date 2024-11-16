package screen

import "math"

type AspectRatio float64

const SixteenByNine AspectRatio = 16.0 / 9.0

func (a AspectRatio) Value() float64 {
	return float64(a)
}

func CalculateScreenHeight(width float64, ratio AspectRatio) int {
	return int(math.Round(width / ratio.Value()))
}
