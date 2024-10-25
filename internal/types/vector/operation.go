package vector

import (
	"math"

	"github.com/lunarisnia/yacg/internal/types"
)

func LengthSquared(v types.Vector3f) float64 {
	return float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func Length(v types.Vector3f) float64 {
	return math.Sqrt(LengthSquared(v))
}

func MultiplyScalar(a types.Vector3f, b float64) types.Vector3f {
	return types.Vector3f{
		X: a.X * b,
		Y: a.Y * b,
		Z: a.Z * b,
	}
}

func DivideScalar(a types.Vector3f, b float64) types.Vector3f {
	return MultiplyScalar(a, 1/b)
}

func UnitVector(v types.Vector3f) types.Vector3f {
	return DivideScalar(v, Length(v))
}
