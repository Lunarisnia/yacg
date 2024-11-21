package hitrecord

import (
	"github.com/lunarisnia/yacg/internal/types"
	"github.com/lunarisnia/yacg/internal/types/vector"
)

func SetFaceNormal(h *types.HitRecord, r types.Ray, outwardNormal types.Vector3f) {
	// If the ray is coming from the outside the normal would be opposite and the dot product is negative
	h.FrontFace = vector.DotProduct(r.Direction, outwardNormal) < 0.0
	// if the ray is coming from the backface we inverse the normal to face inward so that it's always opposite
	if !h.FrontFace {
		h.Normal = vector.InverseVector(outwardNormal)
	}
}
