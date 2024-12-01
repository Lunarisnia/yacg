package vector

import (
	"math"
	"math/rand"

	"github.com/lunarisnia/yacg/color"
	"github.com/lunarisnia/yacg/internal/types"
)

func LengthSquared(v types.Vector3f) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
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

func InverseVector(a types.Vector3f) types.Vector3f {
	return types.Vector3f{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
	}
}

func Random() types.Vector3f {
	return types.Vector3f{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64(),
	}
}

func RandomN(min float64, max float64) types.Vector3f {
	return types.Vector3f{
		X: min + rand.Float64()*(max-min),
		Y: min + rand.Float64()*(max-min),
		Z: min + rand.Float64()*(max-min),
	}
}

// RandomUnitVector https://github.com/RayTracing/raytracing.github.io/discussions/1369 for why does it has to be on the sphere hemisphere
func RandomUnitVector() types.Vector3f {
	for {
		randomVector := RandomN(-1.0, 1.0)
		length := LengthSquared(randomVector)
		if length <= 1 {
			return UnitVector(randomVector)
		}
	}
}

func ToVector(c *color.RGB) types.Vector3f {
	return types.Vector3f{
		X: float64(c.Red),
		Y: float64(c.Green),
		Z: float64(c.Blue),
	}
}

// Reflect is a formula to reflect ray given a direction and a normal
func Reflect(a types.Vector3f, n types.Vector3f) types.Vector3f {
	b := 2.0 * DotProduct(a, n)
	return SubtractVector(a, MultiplyScalar(n, b))
}

func Refract(uv types.Vector3f, n types.Vector3f, refractiveIndex float64) types.Vector3f {
	cosTheta := math.Min(DotProduct(InverseVector(uv), n), 1.0)
	refractiveOutPerpendicular := MultiplyScalar(
		AddVector(MultiplyScalar(n, cosTheta), uv),
		refractiveIndex,
	)
	refractiveOutParallel := MultiplyScalar(
		n,
		-math.Sqrt(math.Abs(1.0-LengthSquared(refractiveOutPerpendicular))),
	)

	return AddVector(refractiveOutPerpendicular, refractiveOutParallel)
}

func CrossProduct(a types.Vector3f, b types.Vector3f) types.Vector3f {
	return types.Vector3f{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// NOTE: YOU CAN'T SUBTRACT SCALAR TO A VECTOR YOU DUMBO
// func SubtractScalar(a types.Vector3f, b float64) types.Vector3f {
// 	return types.Vector3f{
// 		X: a.X - b,
// 		Y: a.Y - b,
// 		Z: a.Z - b,
// 	}
// }

// NOTE: YOU CAN'T ADD SCALAR TO A VECTOR YOU IDIOT
// func AddScalar(a types.Vector3f, b float64) types.Vector3f {
// 	return types.Vector3f{
// 		X: a.X + b,
// 		Y: a.Y + b,
// 		Z: a.Z + b,
// 	}
// }
