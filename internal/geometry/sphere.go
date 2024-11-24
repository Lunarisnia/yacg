package geometry

import (
	"math"

	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/hitrecord"
	"github.com/lunarisnia/yacg/internal/types/ray"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

type Sphere struct {
	Name     string
	Center   types.Vector3f
	Radius   float64
	Material types.Material
}

func (s Sphere) GetName() string {
	return s.Name
}

func (s Sphere) Intersect(
	r types.Ray,
	tMin float64,
	tMax float64,
	hitRecord *types.HitRecord,
) bool {
	center := vector.SubtractVector(s.Center, r.Origin)
	a := vector.DotProduct(r.Direction, r.Direction)
	b := vector.DotProduct(r.Direction, center) * -2.0
	c := vector.DotProduct(center, center) - (s.Radius * s.Radius)
	// This tell us how many intersection are there
	discriminant := b*b - 4*a*c
	if discriminant < 0.0 {
		return false
	}

	root := (-b - math.Sqrt(discriminant)) / (float64(2.0) * a)
	if root <= tMin || root >= tMax {
		root = (-b + math.Sqrt(discriminant)) / (float64(2.0) * a)
		if root <= tMin || root >= tMax {
			return false
		}
	}

	hitRecord.T = root
	hitRecord.HitPoint = ray.At(r, hitRecord.T)
	hitRecord.Normal = vector.UnitVector(vector.SubtractVector(hitRecord.HitPoint, s.Center))
	hitRecord.ObjectName = s.GetName()
	hitrecord.SetFaceNormal(hitRecord, r, hitRecord.Normal)
	hitRecord.Material = s.Material

	return true
}
