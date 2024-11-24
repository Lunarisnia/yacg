package types

type Material interface {
	Scatter(hitRecord *HitRecord) (Vector3f, Ray)
}
