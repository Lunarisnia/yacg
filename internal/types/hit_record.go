package types

type HitRecord struct {
	IncidentalRay Ray
	ObjectName    string
	HitPoint      Vector3f
	Normal        Vector3f
	T             float64
	FrontFace     bool
	Material      Material
}
