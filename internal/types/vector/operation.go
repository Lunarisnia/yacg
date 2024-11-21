package vector

import (
	"math"

	"github.com/lunarisnia/yacg/internal/color"
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
	return MultiplyScalar(a, float64(1.0)/b)
}

func UnitVector(v types.Vector3f) types.Vector3f {
	return DivideScalar(v, Length(v))
}

func DotProduct(a types.Vector3f, b types.Vector3f) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func AddVector(a types.Vector3f, b types.Vector3f) types.Vector3f {
	return types.Vector3f{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func SubtractVector(a types.Vector3f, b types.Vector3f) types.Vector3f {
	return types.Vector3f{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// MultiplyVector produce a Hadamard Product
func MultiplyVector(a types.Vector3f, b types.Vector3f) types.Vector3f {
	return types.Vector3f{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func ToColor(a types.Vector3f) *color.RGB {
	return &color.RGB{
		Red:   int(a.X),
		Green: int(a.Y),
		Blue:  int(a.Z),
	}
}

// NOTE: YOU CAN'T SUBTRACT SCALAR TO A VECTOR YOU DUMBO
func SubtractScalar(a types.Vector3f, b float64) types.Vector3f {
	return types.Vector3f{
		X: a.X - b,
		Y: a.Y - b,
		Z: a.Z - b,
	}
}

// NOTE: YOU CAN'T ADD SCALAR TO A VECTOR YOU IDIOT
func AddScalar(a types.Vector3f, b float64) types.Vector3f {
	return types.Vector3f{
		X: a.X + b,
		Y: a.Y + b,
		Z: a.Z + b,
	}
}
