package geometry

import (
	"math"
	"github.com/go-errors/errors"
)

var InvalidCircle = errors.New("The radius of a circle cannot be zero.")

type Circle struct {
	center Vector
	radius float64
}

// NewCircle creates a new circle with the given center and radius.
func NewCircle(center Vector, radius float64) Circle {
	if radius == 0 {
		panic(InvalidCircle)
	}
	return Circle{center, math.Abs(radius)}
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

// Centroid returns the center of the circle.
func (c Circle) Centroid() Vector {
	return c.center
}

// Contains returns whether or not a given point is on or within the circle.
func (c Circle) Contains(point Vector) bool {
	return c.center.DistanceSquared(point) <= c.radius*c.radius
}

func (c Circle) Rotate(angle float64) Shape {
	if angle == 0 {
		return c
	}

	cos := math.Cos(angle)
	sin := math.Sin(angle)

	center := &c.center
	center.X, center.Y = center.X*cos - center.Y*sin, center.X*sin - center.Y*cos

	return c
}

func (c Circle) RotateAboutPoint(angle float64, point Vector) Shape {
	if point.X == 0 && point.Y == 0 {
		return c.Rotate(angle)
	}

	cos := math.Cos(angle)
	sin := math.Sin(angle)

	center := &c.center

	dx := center.X - point.X
	dy := center.Y - point.Y

	center.X = point.X + (dx*cos - dy*sin)
	center.Y = point.Y + (dx*sin + dy*cos)

	return c
}

func (c Circle) Scale(scaleFactor float64) Shape {
	c.center = c.center.Multiply(scaleFactor)
	return c
}

func (c Circle) Translate(v Vector) Shape {
	c.center = c.center.Add(v)
	return c
}
