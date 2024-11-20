package geometry

import (
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

type Sphere struct {
	Center types.Vector3f
	Radius float64
}

func (s Sphere) Intersect(r types.Ray) bool {
	center := vector.SubtractVector(s.Center, r.Direction)
	a := vector.DotProduct(r.Direction, r.Direction)
	b := vector.DotProduct(r.Direction, center) * -2.0
	c := vector.DotProduct(center, center) - (s.Radius * s.Radius)
	// This tell us how many intersection are there
	discriminant := b*b - 4*a*c
	return discriminant >= 0
}
