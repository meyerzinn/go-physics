package geometry

type Shape interface {
	Area() float64
	Centroid() Vector
	Translate(Vector) Shape
	Rotate(float64) Shape
	RotateAboutPoint(float64, Vector) Shape
	Contains(Vector) bool
	Scale(float64) Shape
}
