package point

import (
	"image"

	"github.com/afarbos/aoc/pkg/mathematic"
)

// Point helper struct on top of image.Point.
type Point struct {
	image.Point
}

// Distance return the distance of the image.Point compare to 0,0.
func Distance(pt image.Point) int {
	return mathematic.AbsInt(pt.X) + mathematic.AbsInt(pt.Y)
}

// New return a new Point.
func New() *Point {
	return new(Point)
}

// GetDistance return the distance of the image.Point compare to 0,0.
func (pt *Point) GetDistance() int {
	return Distance(pt.Point)
}
