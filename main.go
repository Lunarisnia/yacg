package main

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"os"

	"github.com/lunarisnia/yacg/internal/geometry"
	"github.com/lunarisnia/yacg/internal/geometry/object"
	"github.com/lunarisnia/yacg/internal/material"
	"github.com/lunarisnia/yacg/internal/ppm"
	"github.com/lunarisnia/yacg/internal/screen"
	"github.com/lunarisnia/yacg/internal/trigonometry"
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
// NOTE: this is missing a depth of field by choice
// NOTE: Focal Length is the distance between the eye to the viewport/canvas
func main() {
	newPPM := ppm.NewPPM()
	screenWidth := 640
	screenHeight := screen.CalculateScreenHeight(float64(screenWidth), screen.SixteenByNine)
	newPPM.InitPPM(&ppm.PPMHeader{
		Width:  screenWidth,
		Height: screenHeight,
	})

	// Camera Settings
	vFOV := float64(90)
	cameraLookFrom := types.Vector3f{X: -2, Y: 2, Z: 1}
	cameraLookAt := types.Vector3f{X: 0, Y: 0, Z: -1}
	cameraUp := types.Vector3f{X: 0, Y: 1, Z: 0}

	focalLength := vector.Length(vector.SubtractVector(cameraLookFrom, cameraLookAt))
	theta := trigonometry.Deg2Rad(vFOV)
	h := math.Tan(theta / 2.0)
	vWidth := float64(3.5) * h * focalLength
	vHeight := vWidth / (float64(screenWidth) / float64(screenHeight))
	// vHeight := float64(2.0) * h * focalLength
	// vWidth := vHeight * (float64(screenWidth) / float64(screenHeight))
	cameraOrigin := cameraLookFrom

	w := vector.UnitVector(vector.SubtractVector(cameraLookFrom, cameraLookAt))
	u := vector.UnitVector(vector.CrossProduct(cameraUp, w))
	v := vector.CrossProduct(w, u)

	viewportU := vector.MultiplyScalar(u, vWidth)
	viewportV := vector.MultiplyScalar(vector.InverseVector(v), vHeight)

	upperLeftCoordinate := vector.SubtractVector(
		cameraOrigin,
		vector.MultiplyScalar(w, focalLength),
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
	objects = debuggingObjects()
	objects = append(objects, geometry.Sphere{
		Name: "Ground",
		Center: types.Vector3f{
			X: 0,
			Y: -100.5,
			Z: -1,
		},
		Radius: 100,
		Material: material.Diffuse{
			Albedo: types.Vector3f{
				X: 15,
				Y: 15,
				Z: 15,
			},
		},
	})

	maxDepth := 10
	counter := 0
	samplesPerPixel := 100
	pixelSampleScale := float64(1.0) / float64(samplesPerPixel)

	antiAliasing := true

	for i := range screenHeight {
		for j := range screenWidth {
			log.Printf("Rendering: %v of %v pixels\n", counter+1, screenHeight*screenWidth)
			colorVector := types.Vector3f{}
			// Anti-aliasing
			if antiAliasing {
				for range samplesPerPixel {
					sampleSquare := vector.RandomN(-0.5, 0.5)
					pixelCenter := vector.AddVector(
						pixel00,
						vector.AddVector(
							vector.MultiplyScalar(pixelDeltaV, float64(i)+sampleSquare.Y),
							vector.MultiplyScalar(pixelDeltaU, float64(j)+sampleSquare.X),
						),
					)
					rayDirection := vector.SubtractVector(pixelCenter, cameraOrigin)
					// Originates from the eye point moving towards the pixelCenter
					r := types.Ray{
						Origin:    cameraOrigin,
						Direction: vector.UnitVector(rayDirection),
					}
					// fmt.Println("Center Point: ", pixelCenter)
					// fmt.Println("Ray Point: ", r)
					sampleColor := vector.ToVector(
						ray.Raycast(r, 0, maxDepth, 0.001, math.Inf(1), objects),
					)

					colorVector = vector.AddVector(colorVector, sampleColor)
				}
				colorVector = vector.MultiplyScalar(colorVector, pixelSampleScale)
			} else {
				pixelCenter := vector.AddVector(
					pixel00,
					vector.AddVector(
						vector.MultiplyScalar(pixelDeltaV, float64(i)),
						vector.MultiplyScalar(pixelDeltaU, float64(j)),
					),
				)
				rayDirection := vector.SubtractVector(pixelCenter, cameraOrigin)
				r := types.Ray{
					Origin:    cameraOrigin,
					Direction: vector.UnitVector(rayDirection),
				}
				colorVector = vector.ToVector(ray.Raycast(r, 0, maxDepth, 0.001, math.Inf(1), objects))
			}

			// Conversion from linear color space to gamma space
			colorVector.X /= 256
			colorVector.Y /= 256
			colorVector.Z /= 256
			colorVector.X = math.Sqrt(colorVector.X) * 256
			colorVector.Y = math.Sqrt(colorVector.Y) * 256
			colorVector.Z = math.Sqrt(colorVector.Z) * 256

			newPPM.DrawPixel(vector.ToColor(colorVector))
			counter++
		}
	}
}

func debuggingObjects() []object.Object {
	objects := make([]object.Object, 0)
	objects = append(objects, &geometry.Sphere{
		Name: "Diffuse Outside Camera",
		Center: types.Vector3f{
			X: -1,
			Y: 0,
			Z: 1,
		},
		Radius: 0.5,
		Material: material.Diffuse{
			Albedo: types.Vector3f{
				X: 0,
				Y: 0,
				Z: 255,
			},
		},
	})
	objects = append(objects, &geometry.Sphere{
		Name: "Specular To The Left",
		Center: types.Vector3f{
			X: -2,
			Y: 0,
			Z: -2,
		},
		Radius: 0.5,
		Material: material.Specular{
			Albedo: types.Vector3f{
				X: 100,
				Y: 100,
				Z: 10,
			},
			Fuzzy: 0.1,
		},
	})
	objects = append(objects, &geometry.Sphere{
		Name: "Glass Sphere",
		Center: types.Vector3f{
			X: -0.95,
			Y: 0,
			Z: -1.35,
		},
		Radius: 0.5,
		Material: material.Dielectric{
			Albedo: types.Vector3f{
				X: 255,
				Y: 255,
				Z: 255,
			},
			// Refractive index of a glass
			RefractiveIndex: 1.50,
		},
	})
	objects = append(objects, &geometry.Sphere{
		Name: "Inner Glass Sphere",
		Center: types.Vector3f{
			X: -0.95,
			Y: 0,
			Z: -1.35,
		},
		Radius: 0.3,
		Material: material.Dielectric{
			Albedo: types.Vector3f{
				X: 255,
				Y: 255,
				Z: 255,
			},
			// Refractive index of air / refractive index of water
			RefractiveIndex: 1.00 / 1.50,
		},
	})
	objects = append(objects, geometry.Sphere{
		Name: "Middle",
		Center: types.Vector3f{
			X: 0,
			Y: 0,
			Z: -1,
		},
		Radius: 0.5,
		Material: material.Diffuse{
			Albedo: types.Vector3f{
				X: 128,
				Y: 0,
				Z: 0,
			},
		},
	})
	objects = append(objects, geometry.Sphere{
		Name: "Specular To The Right",
		Center: types.Vector3f{
			X: 1.25,
			Y: 0,
			Z: -1.05,
		},
		Radius: 0.5,
		Material: material.Specular{
			Albedo: types.Vector3f{
				X: 50,
				Y: 50,
				Z: 50,
			},
			Fuzzy: 0.001,
		},
	})
	return objects
}
