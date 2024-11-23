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

func Raycast(
	r types.Ray,
	depth int,
	maxDepth int,
	tMin float64,
	tMax float64,
	objects []object.Object,
) *color.RGB {
	if depth >= maxDepth {
		return &color.RGB{}
	}
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
		// Send rays randomly to simulate a diffuse surface
		bounceDirection := vector.RandomUnitVector()
		facingOut := vector.DotProduct(hitRecord.Normal, bounceDirection) > 0.0
		if !facingOut {
			bounceDirection = vector.InverseVector(bounceDirection)
		}
		// True Lambertian Reflection
		bounceDirection = vector.UnitVector(vector.AddVector(hitRecord.Normal, bounceDirection))

		// Trace where the ray go recursively or just return the background if it does not intersect with anything
		c := Raycast(types.Ray{
			Origin:    hitRecord.HitPoint,
			Direction: bounceDirection,
		}, depth+1, maxDepth, tMin, tMax, objects)
		c.Red = int(0.5 * float64(c.Red))
		c.Green = int(0.5 * float64(c.Green))
		c.Blue = int(0.5 * float64(c.Blue))
		return c
	}

	a := 0.5 * (r.Direction.Y + 1.0)
	startColor := vector.MultiplyScalar(types.Vector3f{X: 255, Y: 255, Z: 255}, (1.0 - a))
	endColor := vector.MultiplyScalar(types.Vector3f{X: 128, Y: 179, Z: 255}, a)
	resultingColor := vector.AddVector(startColor, endColor)
	return vector.ToColor(resultingColor)
}
