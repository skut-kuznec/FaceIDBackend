package domain

type FaceDescriptor [128]float32

type Image struct {
	ID             uint64
	Path           string
	FaceDescriptor FaceDescriptor
}
