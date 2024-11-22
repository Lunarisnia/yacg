package types

type HitRecord struct {
	ObjectName string
	HitPoint   Vector3f
	Normal     Vector3f
	T          float64
	FrontFace  bool
}
