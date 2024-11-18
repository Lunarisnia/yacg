package ray

import (
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

func At(r types.Ray, t float64) types.Vector3f {
	return vector.AddVector(r.Origin, vector.MultiplyScalar(r.Direction, t))
}

func Raycast(r types.Ray) types.Vector3f {
	return types.Vector3f{
		X: 100,
		Y: 100,
		Z: 20,
	}
}
