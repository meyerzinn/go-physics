package geometry

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func New(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v *Vector) Hashcode() (hash uint64) {
	x, y := uint64(v.X), uint64(v.Y)
	hash = x + y
	return
}

func (v *Vector) Equals(oi interface{}) (equals bool) {
	o, equals := oi.(*Vector)
	if !equals {
		var ov Vector
		ov, equals = oi.(Vector)
		equals = equals && v.EqualsVector(ov)
		return
	}
	equals = v.EqualsVector(*o)
	return
}

func (v Vector) EqualsVector(q Vector) bool { return v.X == q.X && v.Y == q.Y }

func (v Vector) Clone() Vector {
	return New(v.X, v.Y)
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v Vector) MagnitudeSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y)
}

func (v Vector) RotateLeft() Vector {
	return Vector{
		X: -v.Y,
		Y: v.X,
	}
}

func (v Vector) RotateRight() Vector {
	return Vector{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vector) Rotate(angle float64) Vector {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	return Vector{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin - v.Y*cos,
	}
}

func (v Vector) RotateAboutPoint(angle float64, point Vector) Vector {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x := point.X + ((v.X-point.X)*cos - (v.Y-point.Y)*sin)
	y := point.Y + ((v.X-point.X)*sin + (v.Y-point.Y)*cos)
	return Vector{
		X: x,
		Y: y,
	}
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector{
			X: 0,
			Y: 0,
		}
	}
	return Vector{
		X: v.X / magnitude,
		Y: v.Y / magnitude,
	}
}

func (v Vector) DotProduct(o Vector) float64 {
	return (v.X * o.X) + (v.Y * o.Y)
}

func (v Vector) CrossProduct(o Vector) float64 {
	return (v.X * o.Y) - (v.Y * o.X)
}

func (v Vector) Add(o Vector) Vector {
	return Vector{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vector) Subtract(o Vector) Vector {
	return Vector{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vector) Multiply(s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vector) Perpendicular(negate bool) Vector {
	var n float64
	if negate {
		n = -1
	} else {
		n = 1
	}
	return Vector{
		X: n * -v.Y,
		Y: n * -v.X,
	}
}

func (v Vector) Negative() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vector) Angle(o Vector) float64 {
	return math.Atan2(o.Y-v.Y, o.X-v.X)
}

func (v Vector) Divide(s float64) Vector {
	return Vector{
		X: v.X / s,
		Y: v.Y / s,
	}
}

func (v Vector) Distance(o Vector) float64 {
	return math.Sqrt(v.DistanceSquared(o))
}

func (v Vector) DistanceSquared(o Vector) float64 {
	return (o.X-v.X)*(o.X-v.X) + (o.Y-v.Y)*(o.Y-v.Y)
}
