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

func Raycast(r types.Ray, tMin float64, tMax float64, objects []object.Object) *color.RGB {
	closestObject := tMax // Usually start with positive Infinite
	hitRecord := types.HitRecord{}
	hitSomething := false
	// We scan all the objects first and update the hit record
	// After this loop is done we can be sure that hitRecord is filled with the closest object to the camera
	for _, o := range objects {
		if o.Intersect(r, tMin, closestObject, &hitRecord) {
			// The closest object T position so far (for use in calculating ray at T)
			closestObject = hitRecord.T
			hitSomething = true
		}
	}
	if hitSomething {
		return &color.RGB{
			Red:   int(0.5 * ((hitRecord.Normal.X + 1.0) * 255.0)),
			Green: int(0.5 * ((hitRecord.Normal.Y + 1.0) * 255.0)),
			Blue:  int(0.5 * ((hitRecord.Normal.Z + 1.0) * 255.0)),
		}
	}

	return &color.RGB{
		Red:   15,
		Blue:  15,
		Green: 15,
	}
}
