package main

import (
	"github.com/lunarisnia/yacg/internal/color"
	"github.com/lunarisnia/yacg/internal/geometry"
	"github.com/lunarisnia/yacg/internal/ppm"
	"github.com/lunarisnia/yacg/internal/screen"
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/ray"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

// TODO: Do a writing on summary of this code
// NOTE: Focal Length is the distance between the eye to the viewport/canvas
// TODO: Learn Hadamard Product
func main() {
	newPPM := ppm.NewPPM()
	screenWidth := 400
	screenHeight := screen.CalculateScreenHeight(float64(screenWidth), screen.SixteenByNine)
	newPPM.InitPPM(&ppm.PPMHeader{
		Width:  screenWidth,
		Height: screenHeight,
	})

	vWidth := float64(3.5)
	vHeight := vWidth / (float64(screenWidth) / float64(screenHeight))
	// vHeight := float64(2.0)
	// vWidth := vHeight * (float64(screenWidth) / float64(screenHeight))
	focalLength := float64(1.0)
	cameraOrigin := types.Vector3f{X: 0, Y: 0, Z: 0}

	viewportU := types.Vector3f{X: vWidth, Y: 0, Z: 0}
	viewportV := types.Vector3f{X: 0, Y: -vHeight, Z: 0}

	upperLeftCoordinate := vector.SubtractVector(cameraOrigin, types.Vector3f{X: 0, Y: 0, Z: focalLength})
	upperLeftCoordinate = vector.SubtractVector(upperLeftCoordinate, vector.DivideScalar(viewportU, float64(2.0)))
	upperLeftCoordinate = vector.SubtractVector(upperLeftCoordinate, vector.DivideScalar(viewportV, float64(2.0)))

	pixelDeltaU := vector.DivideScalar(viewportU, float64(screenWidth))
	pixelDeltaV := vector.DivideScalar(viewportV, float64(screenHeight))

	pixel00Inset := vector.AddVector(pixelDeltaU, pixelDeltaV)
	pixel00 := vector.AddVector(upperLeftCoordinate, vector.MultiplyScalar(pixel00Inset, 0.5))

	for i := range screenHeight {
		for j := range screenWidth {
			pixelCenter := vector.AddVector(
				pixel00,
				vector.AddVector(
					vector.MultiplyScalar(pixelDeltaV, float64(i)),
					vector.MultiplyScalar(pixelDeltaU, float64(j)),
				),
			)
			rayDirection := vector.SubtractVector(pixelCenter, cameraOrigin)
			r := types.Ray{
				Origin:    pixelCenter,
				Direction: vector.UnitVector(rayDirection),
			}
			s := geometry.Sphere{
				Center: types.Vector3f{
					X: 0,
					Y: 0,
					Z: -10,
				},
				Radius: 1,
			}
			if t := s.Intersect(r); t > 0.0 {
				// TODO: Refactor this to separate the color of the sphere
				// colorVector := ray.At(r, t)
				newPPM.DrawPixel(&color.RGB{Red: 255, Green: 255, Blue: 255})
			} else {
				colorVector := ray.Raycast(r)
				newPPM.DrawPixel(vector.ToColor(colorVector))
			}
			// fmt.Println("Center Point: ", pixelCenter)
			// fmt.Println("Ray Point: ", r)
			// colorVector := ray.Raycast(r)
			// newPPM.DrawPixel(vector.ToColor(colorVector))
		}
	}
}
