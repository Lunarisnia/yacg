package main

import (
	"fmt"

	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/tmatrices"
)

func main() {
	f := types.TMatrice{
		[4]float64{1, 8, 9, 1},
		[4]float64{2, 7, 10, 1},
		[4]float64{3, 6, 11, 1},
		[4]float64{4, 5, 12, 1},
	}
	// fmt.Println(f)
	fmt.Println(tmatrices.MultiplyVector(f, types.Vector3f{X: 1, Y: 1, Z: 1}))
	// newPPM := ppm.NewPPM()
	// newPPM.DrawCubeCorner()
}
