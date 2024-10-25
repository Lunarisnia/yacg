package ppm

import (
	"errors"
	"fmt"

	"github.com/lunarisnia/yacg/internal/color"
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
