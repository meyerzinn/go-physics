package geometry

import (
	"errors"
	"math"

	"github.com/bradfitz/slice"
)

type Polygon struct {
	vertices []Vector
}

var InvalidPolygon = errors.New("A polygon must have at least 3 vertices!")

func NewPolygon(vertices []Vector) Polygon {
	if len(vertices) < 3 {
		panic(InvalidPolygon)
	}
	sorted := clockwiseSort(vertices)
	return Polygon{sorted}
}

func (p Polygon) Area() float64 {
	var area float64 = 0
	var j = len(p.vertices) - 1
	for i := 0; i < len(p.vertices); i ++ {
		area += (p.vertices[j].X - p.vertices[i].X) * (p.vertices[j].Y + p.vertices[i].Y)
		j = i
	}
	return math.Abs(area) / 2
}

func (p Polygon) Centroid() Vector {
	area := p.Area()
	centre := Vector{0, 0}
	var cross float64
	var temp Vector
	var j int

	for i := 0; i < len(p.vertices); i++ {
		j = (i + 1) % len(p.vertices)
		cross = p.vertices[i].CrossProduct(p.vertices[j])
		temp = p.vertices[i].Add(p.vertices[j]).Multiply(cross)
		centre = centre.Add(temp)
	}

	return centre.Divide(6 * area)
}

func (p Polygon) Translate(v Vector) Shape {
	for i := 0; i < len(p.vertices); i++ {
		p.vertices[i] = p.vertices[i].Add(v)
	}
	return p
}

func (p Polygon) Rotate(angle float64) Shape {
	if angle == 0 {
		return p
	}

	cos := math.Cos(angle)
	sin := math.Sin(angle)

	for i := 0; i < len(p.vertices); i++ {
		vertex := &p.vertices[i]
		vertex.X, vertex.Y = vertex.X*cos - vertex.Y*sin, vertex.X*sin - vertex.Y*cos
	}

	return p
}

func (p Polygon) RotateAboutPoint(angle float64, point Vector) Shape {
	if angle == 0 {
		return p
	}
	if point.X == 0 && point.Y == 0 {
		return p.Rotate(angle)
	}

	cos := math.Cos(angle)
	sin := math.Sin(angle)

	for i := 0; i < len(p.vertices); i++ {
		vector := &p.vertices[i]
		dx := vector.X - point.X
		dy := vector.Y - point.Y
		vector.X = point.X + (dx*cos - dy*sin)
		vector.Y = point.Y + (dx*sin + dy*cos)
	}
	return p
}

func (p Polygon) Contains(point Vector) bool {
	for i := 0; i < len(p.vertices); i++ {
		vertex := p.vertices[i]
		nextVertex := p.vertices[(i + 1) % len(p.vertices)]
		if (point.X-vertex.X)*(nextVertex.Y-vertex.Y)+(point.Y-vertex.Y)*(vertex.X-nextVertex.X) > 0 {
			return false
		}
	}
	return true
}

func (p Polygon) Scale(scaleFactor float64) Shape {
	var point = p.Centroid()

	return p.ScaleAboutPoint(scaleFactor, point)
}

func (p Polygon) ScaleAboutPoint(scaleFactor float64, point Vector) Shape {
	if scaleFactor == 1 {
		return p
	}

	for i := 0; i < len(p.vertices); i++ {
		vertex := &p.vertices[i]
		delta := vertex.Subtract(point)

		vertex.X = point.X + delta.X*scaleFactor
		vertex.Y = point.Y + delta.Y*scaleFactor
	}

	return p
}

func mean(vertices []Vector) Vector {
	var average = Vector{0, 0}

	for i := 0; i < len(vertices); i++ {
		average.X += vertices[i].X
		average.Y += vertices[i].Y
	}
	return average.Divide(float64(len(vertices)))
}

func clockwiseSort(vertices []Vector) []Vector {
	var centre = mean(vertices)

	slice.Sort(vertices, func(i, j int) bool {
		if centre.Angle(vertices[i])-centre.Angle(vertices[i]) > 0 {
			return true
		}
		return false
	})

	return vertices
}

func (p Polygon) IsConvex() (bool) {

	flag := byte(0)
	n := len(p.vertices)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		k := (i + 2) % n
		z := (p.vertices[j].X - p.vertices[i].X) * (p.vertices[k].Y - p.vertices[j].Y)
		z -= (p.vertices[j].Y - p.vertices[i].Y) * (p.vertices[k].X - p.vertices[j].X)

		if z < 0 {
			flag |= 1
		} else if z > 0 {
			flag |= 2
		}

		if flag == 3 {
			return false
		}
	}
	return true
}
