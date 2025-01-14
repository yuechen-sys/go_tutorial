package math

import "github.com/yuechen-sys/go_tutorial/math/simplemath"

type Point struct {
	x, y float64 // y coordinate
}

// NewPoint creates a new Point
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance calculates the distance between two points
func (p Point) Distance(q Point) float64 {
	return simplemath.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}
