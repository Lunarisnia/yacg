package main

import (
	"fmt"

	"github.com/lunarisnia/yacg/internal/color"
	"github.com/lunarisnia/yacg/internal/ppm"
	"github.com/lunarisnia/yacg/internal/screen"
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

// TODO: Do a writing on summary of this code
// NOTE: Focal Length is the distance between the eye to the viewport/canvas
// TODO: Learn Hadamard Product
func main() {
	newPPM := ppm.NewPPM()
	screenWidth := 800
	screenHeight := screen.CalculateScreenHeight(float64(screenWidth), screen.SixteenByNine)
	newPPM.InitPPM(&ppm.PPMHeader{
		Width:  screenWidth,
		Height: screenHeight,
	})

	vWidth := float64(3.5)
	vHeight := vWidth / (float64(screenWidth) / float64(screenHeight))
	focalLength := float64(1.0)
	cameraOrigin := types.Vector3f{}

	viewportU := types.Vector3f{X: vWidth, Y: 0, Z: 0}
	viewportV := types.Vector3f{X: 0, Y: -vHeight, Z: 0}

	upperLeftCoordinate := vector.SubtractVector(vector.SubtractVector(vector.SubtractVector(
		cameraOrigin,
		types.Vector3f{X: 0, Y: 0, Z: focalLength},
	), vector.DivideScalar(viewportU, 2)), vector.DivideScalar(viewportV, 2))

	pixelDeltaU := vector.DivideScalar(viewportU, float64(screenWidth))
	pixelDeltaV := vector.DivideScalar(viewportV, float64(screenHeight))
	pixel00 := vector.AddVector(
		vector.AddScalar(upperLeftCoordinate, 0.5),
		vector.MultiplyVector(pixelDeltaU, pixelDeltaV),
	)
	fmt.Println(pixel00)

	for i := range screenHeight {
		for j := range screenWidth {
			normalizedY := float64(i) / float64(screenHeight)
			normalizedX := float64(j) / float64(screenWidth)
			err := newPPM.DrawPixel(&color.RGB{
				Green: int(normalizedY * 256.0),
				Blue:  int(normalizedX * 256.0),
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
