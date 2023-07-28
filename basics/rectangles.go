package basics

type RectSpecs struct {
	Width   int
	Height  int
	OriginX int
	OriginY int
}

// the output of this is good for gl.DrawArrays
func RectangleToCoords(windowWidth, windowHeight int, rectSpec RectSpecs) []float32 {

	point1X := XtoFloat(rectSpec.OriginX, windowWidth)
	point1Y := YtoFloat(rectSpec.OriginY, windowHeight)

	point2X := XtoFloat(rectSpec.OriginX+rectSpec.Width, windowWidth)
	point2Y := YtoFloat(rectSpec.OriginY+rectSpec.Height, windowHeight)

	retFloat32 := []float32{
		// first triangle
		point1X, point1Y, 0,
		point1X, point2Y, 0,
		point2X, point2Y, 0,

		// second triangle
		point1X, point1Y, 0,
		point2X, point1Y, 0,
		point2X, point2Y, 0,
	}

	return retFloat32
}

// the outputs of this is good for gl.DrawElements
func RectangleToCoords2(windowWidth, windowHeight int, rectSpec RectSpecs) ([]float32, []uint32) {

	point1X := XtoFloat(rectSpec.OriginX, windowWidth)
	point1Y := YtoFloat(rectSpec.OriginY, windowHeight)

	point2X := XtoFloat(rectSpec.OriginX+rectSpec.Width, windowWidth)
	point2Y := YtoFloat(rectSpec.OriginY+rectSpec.Height, windowHeight)

	// retFloat32 := []float32{
	// 	// first triangle
	// 	point1X, point1Y, 0,
	// 	point1X, point2Y, 0,
	// 	point2X, point2Y, 0,

	// 	// second triangle
	// 	point1X, point1Y, 0,
	// 	point2X, point1Y, 0,
	// 	point2X, point2Y, 0,
	// }

	retVertices := []float32{
		point1X, point1Y, 0,
		point1X, point2Y, 0,
		point2X, point2Y, 0,
		point2X, point1Y, 0,
	}

	retIndices := []uint32{
		0, 1, 2,
		0, 2, 3,
	}

	return retVertices, retIndices
}

// the outputs of this is good for gl.DrawElements
func ImageCoordinates(windowWidth, windowHeight int, rectSpec RectSpecs) ([]float32, []uint32) {
	tmpVertices, indices := RectangleToCoords2(windowWidth, windowHeight, rectSpec)
	v1 := tmpVertices
	// inject texture coordinates
	vertices := []float32{
		v1[0], v1[1], v1[2], // vertices position
		1.0, 0.0, // texture coordinates

		v1[3], v1[4], v1[5],
		1.0, 1.0,

		v1[6], v1[7], v1[8],
		0.0, 1.0,

		v1[9], v1[10], v1[11],
		0.0, 0.0,
	}

	return vertices, indices
}
