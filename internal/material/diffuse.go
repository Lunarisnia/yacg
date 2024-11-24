package material

import (
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

type Diffuse struct {
	Albedo types.Vector3f
}

func (d Diffuse) Scatter(hitRecord *types.HitRecord) (types.Vector3f, types.Ray, bool) {
	bounceDirection := vector.RandomUnitVector()
	facingOut := vector.DotProduct(hitRecord.Normal, bounceDirection) > 0.0
	if !facingOut {
		bounceDirection = vector.InverseVector(bounceDirection)
	}
	// True Lambertian Reflection
	bounceDirection = vector.UnitVector(vector.AddVector(hitRecord.Normal, bounceDirection))
	c := d.Albedo
	c.X /= 256
	c.Y /= 256
	c.Z /= 256

	return c, types.Ray{
		Origin:    hitRecord.HitPoint,
		Direction: bounceDirection,
	}, true
}
