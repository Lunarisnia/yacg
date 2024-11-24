package material

import (
	"math"
	"math/rand/v2"

	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

type Dielectric struct {
	Albedo          types.Vector3f
	RefractiveIndex float64
}

func (d Dielectric) Scatter(hitRecord *types.HitRecord) (types.Vector3f, types.Ray, bool) {
	c := d.Albedo
	c.X /= 256
	c.Y /= 256
	c.Z /= 256
	var trueRefractiveIndex float64
	if hitRecord.FrontFace {
		trueRefractiveIndex = 1.0 / d.RefractiveIndex
	} else {
		trueRefractiveIndex = d.RefractiveIndex
	}

	// Total internal reflection
	cosTheta := math.Min(
		vector.DotProduct(
			vector.InverseVector(hitRecord.IncidentalRay.Direction),
			hitRecord.Normal,
		),
		1.0,
	)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := trueRefractiveIndex*sinTheta > 1.0
	var rayDirection types.Vector3f

	if cannotRefract || d.reflectance(cosTheta, trueRefractiveIndex) > rand.Float64() {
		rayDirection = vector.Reflect(hitRecord.IncidentalRay.Direction, hitRecord.Normal)
	} else {
		rayDirection = vector.Refract(hitRecord.IncidentalRay.Direction, hitRecord.Normal, trueRefractiveIndex)
	}

	return c, types.Ray{
		Origin:    hitRecord.HitPoint,
		Direction: rayDirection,
	}, true
}

//	static double reflectance(double cosine, double refraction_index) {
//	        // Use Schlick's approximation for reflectance.
//	        auto r0 = (1 - refraction_index) / (1 + refraction_index);
//	        r0 = r0*r0;
//	        return r0 + (1-r0)*std::pow((1 - cosine),5);
//	    }
func (d Dielectric) reflectance(cosine float64, refractiveIndex float64) float64 {
	// Schlick's Approximation
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
