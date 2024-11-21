package geometry

import (
	"math"

	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

type Sphere struct {
	Center types.Vector3f
	Radius float64
}

func (s Sphere) Intersect(r types.Ray) float64 {
	center := vector.SubtractVector(s.Center, r.Origin)
	a := vector.DotProduct(r.Direction, r.Direction)
	b := vector.DotProduct(r.Direction, center) * -2.0
	c := vector.DotProduct(center, center) - (s.Radius * s.Radius)
	// This tell us how many intersection are there
	discriminant := b*b - 4*a*c
	if discriminant < 0.0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (float64(2.0) * a)
	}
}
