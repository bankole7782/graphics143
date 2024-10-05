package graphics143

import (
	"math"
)

type Circle struct {
	TopX   int
	TopY   int
	Radius int
}

func NewCircle(x, y, r int) Circle {
	return Circle{TopX: x, TopY: y, Radius: r}
}

func InCircle(aCircle Circle, xPos, yPos int) bool {
	centerX, centerY := aCircle.CenterCoordinates()
	tmpX := math.Pow(float64(xPos)-centerX, 2.0)
	tmpY := math.Pow(float64(yPos)-centerY, 2.0)
	tmpR := math.Pow(float64(aCircle.Radius), 2.0)

	return tmpX+tmpY <= tmpR
}

func (aCircle Circle) CenterCoordinates() (float64, float64) {
	angleInRadians1 := float64(0) * (math.Pi / 180)
	x1 := float64(aCircle.Radius) * math.Sin(angleInRadians1)
	y1 := float64(aCircle.Radius) * math.Cos(angleInRadians1)

	angleInRadians2 := float64(180) * (math.Pi / 180)
	x2 := float64(aCircle.Radius) * math.Sin(angleInRadians2)
	y2 := float64(aCircle.Radius) * math.Sin(angleInRadians2)

	x3 := (x1 + x2) / 2
	y3 := (y1 + y2) / 2

	return float64(aCircle.TopX) + x3, float64(aCircle.TopY) + y3
}
