package screen

import (
	"math"

	"github.com/lunarisnia/yacg/internal/trigonometry"
)

func CalculateCanvasSize(angleOfView float64, zNear float64) float64 {
	return 2.0 * math.Tan(trigonometry.Deg2Rad(angleOfView)/2) * zNear
}
