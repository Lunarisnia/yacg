package object

import "github.com/lunarisnia/yacg/internal/types"

type Object interface {
	Intersect(r types.Ray, hitRecord *types.HitRecord) bool
}
