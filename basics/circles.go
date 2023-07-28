package basics

import "math"

func CircleCoords(windowWidth, windowHeight, originX, originY, radius int) []float32 {
	twicePi := 2 * math.Pi
	triangleAmount := 64

	radiusX := float64(radius) / float64(windowWidth)
	radiusY := float64(radius) / float64(windowHeight)

	originXf32 := XtoFloat(originX, windowWidth)
	originYf32 := YtoFloat(originY, windowHeight)

	vertices := make([]float32, 0)
	// vertices = append(vertices, originX, originY, 0)
	for i := 0; i < triangleAmount; i++ {
		x := originXf32 + float32(radiusX*math.Cos(float64(i)*twicePi/float64(triangleAmount)))
		y := originYf32 + float32(radiusY*math.Sin(float64(i)*twicePi/float64(triangleAmount)))
		vertices = append(vertices, x, y, 0)
	}

	return vertices
}
