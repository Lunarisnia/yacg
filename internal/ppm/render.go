package ppm

import (
	"errors"
	"fmt"
	"math"

	"github.com/lunarisnia/yacg/color"
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
	DrawRemotePixel(c *color.RGB, x int, y int, remoteHandler func(x int, y int, c *color.RGB)) error
	// DebugImage will print out a 3x2 image with some random color just for debugging ppm
	DebugImage()
	DrawCubeCorner()
	RenderWithCustomCoordinate()
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
	if header.MaxColorValue == 0 {
		header.MaxColorValue = 255
	}
	p.imageSize = header.Width * header.Height
	fmt.Println("P3")
	fmt.Printf("%v %v\n", header.Width, header.Height)
	fmt.Println(header.MaxColorValue)
	return nil
}

func (p *PPMImpl) DrawRemotePixel(c *color.RGB, x int, y int, remoteHandler func(x int, y int, c *color.RGB)) error {
	if p.pixelIndex >= p.imageSize {
		return errors.New("pixel exceeded image size. please resize your image")
	}
	if c == nil {
		c = &color.RGB{} // Default to black
	}
	if c.Red > 255 {
		c.Red = 255
	}
	if c.Blue > 255 {
		c.Blue = 255
	}
	if c.Green > 255 {
		c.Green = 255
	}
	fmt.Printf("%v %v %v\n", c.Red, c.Green, c.Blue)
	p.pixelIndex++
	remoteHandler(x, y, c)
	return nil
}

func (p *PPMImpl) DrawPixel(c *color.RGB) error {
	if p.pixelIndex >= p.imageSize {
		return errors.New("pixel exceeded image size. please resize your image")
	}
	if c == nil {
		c = &color.RGB{} // Default to black
	}
	if c.Red > 255 {
		c.Red = 255
	}
	if c.Blue > 255 {
		c.Blue = 255
	}
	if c.Green > 255 {
		c.Green = 255
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

// DrawCubeCorner this use rasterization technique
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
		// This results in screenspace
		pPrimeX := c.X / c.Z * -1
		pPrimeX *= vector.Length(canvasPosition)
		pPrimeY := c.Y / c.Z * -1
		pPrimeY *= vector.Length(canvasPosition)

		// fmt.Printf("X: %v, Y: %v\n", pPrimeX, pPrimeY)

		// Convert to normalized device coordinate
		normalizedPrimeX := (float64(screenWidth)/2 + pPrimeX) / float64(screenWidth)
		normalizedPrimeY := (float64(screenHeight)/2 + pPrimeY) / float64(screenHeight)
		normedCubeProjection = append(normedCubeProjection, types.Vector3f{
			X: normalizedPrimeX,
			Y: normalizedPrimeY,
			Z: c.Z,
		})
		// fmt.Printf("X: %v, Y: %v\n", normalizedPrimeX, normalizedPrimeY)
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
				// Convert to rasterspace
				if i == int(math.Round(cp.X*float64(screenWidth))) &&
					j == int(math.Round(cp.Y*float64(screenHeight))) {
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
