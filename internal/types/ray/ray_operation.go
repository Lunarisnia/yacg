package ray

import (
	"github.com/lunarisnia/yacg/internal/color"
	"github.com/lunarisnia/yacg/internal/geometry/object"
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

// At r = o + (t*d)
func At(r types.Ray, t float64) types.Vector3f {
	return vector.AddVector(r.Origin, vector.MultiplyScalar(r.Direction, t))
}

func Raycast(r types.Ray, objects []object.Object) *color.RGB {
	for _, o := range objects {
		hitRecord := types.HitRecord{}
		if o.Intersect(r, &hitRecord) {
			return &color.RGB{
				Red:   int(0.5 * ((hitRecord.Normal.X + 1.0) * 255.0)),
				Green: int(0.5 * ((hitRecord.Normal.Y + 1.0) * 255.0)),
				Blue:  int(0.5 * ((hitRecord.Normal.Z + 1.0) * 255.0)),
			}
		}
	}

	return &color.RGB{
		Red:   15,
		Blue:  15,
		Green: 15,
	}
}
