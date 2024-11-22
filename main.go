package main

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"os"

	"github.com/lunarisnia/yacg/internal/geometry"
	"github.com/lunarisnia/yacg/internal/geometry/object"
	"github.com/lunarisnia/yacg/internal/ppm"
	"github.com/lunarisnia/yacg/internal/screen"
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/ray"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

func init() {
	// NOTE: use log pkg for debugging and progress bar without affecting the rendering
	args := os.Args
	if len(args) > 1 && args[1] == "debug" {
		dummy := bufio.NewWriter(bytes.NewBuffer([]byte{}))
		log.SetOutput(dummy)
	} else {
		log.SetOutput(os.Stderr)
	}
}

// TODO: Do another writing on summary of this code
// NOTE: Focal Length is the distance between the eye to the viewport/canvas
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
	// vHeight := float64(2.0)
	// vWidth := vHeight * (float64(screenWidth) / float64(screenHeight))
	focalLength := float64(1.0)
	cameraOrigin := types.Vector3f{X: 0, Y: 0, Z: 0}

	viewportU := types.Vector3f{X: vWidth, Y: 0, Z: 0}
	viewportV := types.Vector3f{X: 0, Y: -vHeight, Z: 0}

	upperLeftCoordinate := vector.SubtractVector(
		cameraOrigin,
		types.Vector3f{X: 0, Y: 0, Z: focalLength},
	)
	upperLeftCoordinate = vector.SubtractVector(
		upperLeftCoordinate,
		vector.DivideScalar(viewportU, float64(2.0)),
	)
	upperLeftCoordinate = vector.SubtractVector(
		upperLeftCoordinate,
		vector.DivideScalar(viewportV, float64(2.0)),
	)

	pixelDeltaU := vector.DivideScalar(viewportU, float64(screenWidth))
	pixelDeltaV := vector.DivideScalar(viewportV, float64(screenHeight))

	pixel00Inset := vector.AddVector(pixelDeltaU, pixelDeltaV)
	pixel00 := vector.AddVector(
		upperLeftCoordinate,
		vector.MultiplyScalar(pixel00Inset, float64(0.5)),
	)

	objects := make([]object.Object, 0)
	sphere01 := geometry.Sphere{
		Name: "Sphere 01",
		Center: types.Vector3f{
			X: 0,
			Y: 0,
			Z: -2,
		},
		Radius: 0.5,
	}
	sphere02 := geometry.Sphere{
		Name: "Sphere 02",
		Center: types.Vector3f{
			X: 0.5,
			Y: 0,
			Z: -4.5,
		},
		Radius: 1.5,
	}
	objects = append(objects, &sphere02)
	objects = append(objects, &sphere01)

	counter := 0
	for i := range screenHeight {
		for j := range screenWidth {
			log.Printf("Rendering: %v of %v frames\n", counter+1, screenHeight*screenWidth)
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
			// fmt.Println("Center Point: ", pixelCenter)
			// fmt.Println("Ray Point: ", r)
			colorVector := ray.Raycast(r, 0.0, math.Inf(1), objects)
			newPPM.DrawPixel(colorVector)
			counter++
		}
	}
}
