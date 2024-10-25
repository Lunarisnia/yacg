package ppm

import (
	"errors"
	"fmt"
	"math"

	"github.com/lunarisnia/yacg/internal/color"
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

// P3
// # "P3" means this is a RGB color image in ASCII
// # "3 2" is the width and height of the image in pixels
// # "255" is the maximum value for each color
// # This, up through the "255" line below are the header.
// # Everything after that is the image data: RGB triplets.
// # In order: red, green, blue, yellow, white, and black.
// 3 2
// 255
// 255   0   0
//
//	0 255   0
//	0   0 255
//
// 255 255   0
// 255 255 255
//
//	0   0   0
const (
	MaxColorValue = 255
)

type PPMHeader struct {
	Width         int
	Height        int
	MaxColorValue int
}

type PPM interface {
	InitPPM(header *PPMHeader) error
	DrawPixel(c *color.RGB) error
	// DebugImage will print out a 3x2 image with some random color just for debugging ppm
	DebugImage()
	DrawCubeCorner()
}

type PPMImpl struct {
	header     *PPMHeader
	pixelIndex int
	imageSize  int
}

func NewPPM() PPM {
	return &PPMImpl{}
}

func (p *PPMImpl) InitPPM(header *PPMHeader) error {
	if header == nil {
		return errors.New("header is required")
	}
	p.imageSize = header.Width * header.Height
	fmt.Println("P3")
	fmt.Printf("%v %v\n", header.Width, header.Height)
	fmt.Println(header.MaxColorValue)
	return nil
}

func (p *PPMImpl) DrawPixel(c *color.RGB) error {
	if p.pixelIndex >= p.imageSize {
		return errors.New("pixel exceeded image size. please resize your image")
	}
	if c == nil {
		c = &color.RGB{} // Default to black
	}
	fmt.Printf("%v %v %v\n", c.Red, c.Green, c.Blue)
	p.pixelIndex++
	return nil
}

func (p *PPMImpl) DebugImage() {
	err := p.InitPPM(&PPMHeader{
		Width:         3,
		Height:        2,
		MaxColorValue: MaxColorValue,
	})
	if err != nil {
		panic(err)
	}
	err = p.DrawPixel(&color.RGB{
		Red:   255,
		Green: 0,
		Blue:  0,
	})
	if err != nil {
		panic(err)
	}
	err = p.DrawPixel(&color.RGB{
		Red:   0,
		Green: 255,
		Blue:  0,
	})
	if err != nil {
		panic(err)
	}
	err = p.DrawPixel(&color.RGB{
		Red:   0,
		Green: 0,
		Blue:  255,
	})
	if err != nil {
		panic(err)
	}
	err = p.DrawPixel(&color.RGB{
		Red:   0,
		Green: 255,
		Blue:  255,
	})
	if err != nil {
		panic(err)
	}
	err = p.DrawPixel(&color.RGB{
		Red:   255,
		Green: 255,
		Blue:  255,
	})
	if err != nil {
		panic(err)
	}
	err = p.DrawPixel(&color.RGB{
		Red:   100,
		Green: 100,
		Blue:  100,
	})
	if err != nil {
		panic(err)
	}
}

func (p *PPMImpl) DrawCubeCorner() {
	screenWidth := 20
	screenHeight := 20

	canvasPosition := types.Vector3f{
		X: 0.0,
		Y: 0.0,
		Z: -1.0,
	}

	cube := []types.Vector3f{
		{X: 5, Y: -5, Z: -5},
		{X: 5, Y: -5, Z: -3},
		{X: 5, Y: 5, Z: -5},
		{X: 5, Y: 5, Z: -3},
		{X: -5, Y: -5, Z: -5},
		{X: -5, Y: -5, Z: -3},
		{X: -5, Y: 5, Z: -5},
		{X: -5, Y: 5, Z: -3},
	}
	normedCubeProjection := make([]types.Vector3f, 0)
	for _, c := range cube {
		pPrimeX := c.X / c.Z * -1
		pPrimeX *= vector.Length(canvasPosition)

		pPrimeY := c.Y / c.Z * -1
		pPrimeY *= vector.Length(canvasPosition)
		// fmt.Printf("X: %v, Y: %v\n", pPrimeX, pPrimeY)

		normalizedPrimeX := (float64(screenWidth)/2 + pPrimeX) / float64(screenWidth)
		normalizedPrimeY := (float64(screenHeight)/2 + pPrimeY) / float64(screenHeight)
		normedCubeProjection = append(normedCubeProjection, types.Vector3f{
			X: normalizedPrimeX,
			Y: normalizedPrimeY,
			Z: c.Z,
		})
		// fmt.Printf("X: %v, Y: %v\n", normalizedPrimeX*20, normalizedPrimeY*20)
	}

	p.InitPPM(&PPMHeader{
		Width:         screenWidth,
		Height:        screenHeight,
		MaxColorValue: MaxColorValue,
	})
	for j := range screenHeight {
		for i := range screenWidth {
			c := &color.RGB{
				Red:   0,
				Green: 0,
				Blue:  0,
			}
			for _, cp := range normedCubeProjection {
				if i == int(math.Round(cp.X*float64(screenWidth))) && j == int(math.Round(cp.Y*float64(screenHeight))) {
					c.Red = 255
				}
			}
			err := p.DrawPixel(c)
			if err != nil {
				panic(err)
			}
		}
	}
}
