package object

import "github.com/lunarisnia/yacg/internal/types"

type Object interface {
	Intersect(r types.Ray, tMin float64, tMax float64, hitRecord *types.HitRecord) bool
	GetName() string
}
