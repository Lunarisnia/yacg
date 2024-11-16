package ppm

import (
	"github.com/lunarisnia/yacg/internal/color"
)

func (p *PPMImpl) RenderWithCustomCoordinate() {
	screenWidth := 100
	screenHeight := 100

	// In world space (origin: {0, 0, 0})
	// cube := []types.Vector3f{
	// 	{X: 10, Y: 10, Z: -5},
	// 	{X: 10, Y: 10, Z: -10},
	// 	{X: 10, Y: -10, Z: -5},
	// 	{X: 10, Y: -10, Z: -10},
	// 	{X: -10, Y: 10, Z: -5},
	// 	{X: -10, Y: 10, Z: -10},
	// 	{X: -10, Y: -10, Z: -5},
	// 	{X: -10, Y: -10, Z: -10},
	// }

	err := p.InitPPM(&PPMHeader{
		Width:  screenWidth,
		Height: screenHeight,
	})
	if err != nil {
		panic(err)
	}

	for range screenHeight {
		for range screenWidth {
			err := p.DrawPixel(&color.RGB{
				Red: 255,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}
