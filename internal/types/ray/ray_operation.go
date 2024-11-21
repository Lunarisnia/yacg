package ray

import (
	"github.com/lunarisnia/yacg/internal/color"
	"github.com/lunarisnia/yacg/internal/geometry"
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

// At r = o + (t*d)
func At(r types.Ray, t float64) types.Vector3f {
	return vector.AddVector(r.Origin, vector.MultiplyScalar(r.Direction, t))
}

func Raycast(r types.Ray) *color.RGB {
	s := geometry.Sphere{
		Center: types.Vector3f{
			X: 0,
			Y: 0,
			Z: -2,
		},
		Radius: 0.5,
	}
	if t := s.Intersect(r); t > 0.0 {
		r := At(r, t)
		normal := vector.SubtractVector(r, s.Center)
		normal = vector.UnitVector(normal)
		return &color.RGB{
			Red:   int(0.5 * ((normal.X + 1.0) * 255.0)),
			Green: int(0.5 * ((normal.Y + 1.0) * 255.0)),
			Blue:  int(0.5 * ((normal.Z + 1.0) * 255.0)),
		}
	}

	return &color.RGB{
		Red:   15,
		Blue:  15,
		Green: 15,
	}
}
