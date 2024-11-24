package material

import (
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

type Specular struct {
	Albedo types.Vector3f
	Fuzzy  float64
}

func (s Specular) Scatter(hitRecord *types.HitRecord) (types.Vector3f, types.Ray, bool) {
	reflection := vector.Reflect(hitRecord.IncidentalRay.Direction, hitRecord.Normal)
	// log.Println("REF: ", vector.Length(reflection))
	reflection = vector.AddVector(
		reflection,
		vector.MultiplyScalar(vector.RandomUnitVector(), s.Fuzzy),
	)

	c := s.Albedo
	c.X /= 256
	c.Y /= 256
	c.Z /= 256

	sameDirection := vector.DotProduct(hitRecord.Normal, reflection) > 0.0

	return c, types.Ray{
		Origin:    hitRecord.HitPoint,
		Direction: reflection,
	}, sameDirection
}
