package tmatrices

import (
	"github.com/lunarisnia/yacg/internal/types"
)

func MultiplyVector(a types.TMatrice, b types.Vector3f) types.Vector3f {
	// NOTE: Equivalent to below approach
	// v := []float64{b.X, b.Y, b.Z}
	// res := []float64{0, 0, 0}
	// for range 4 * 4 {
	// 	for i := range v {
	// 		res[i] = v[0]*a[0][i] + v[1]*a[1][i] + v[2]*a[2][i] + a[3][i]
	// 	}
	// }
	// fmt.Println(res)

	// 00, 01, 02, 03
	// 10, 11, 12, 13
	// 20, 21, 22, 23
	// 30, 31, 32, 33
	x := b.X*a[0][0] + b.Y*a[1][0] + b.Z*a[2][0] + a[3][0]
	y := b.X*a[0][1] + b.Y*a[1][1] + b.Z*a[2][1] + a[3][1]
	z := b.X*a[0][2] + b.Y*a[1][2] + b.Z*a[2][2] + a[3][2]
	return types.Vector3f{
		X: x,
		Y: y,
		Z: z,
	}
}
