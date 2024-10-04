package graphics143

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// DrawImage draws an image to a GLFW window.
// You must have initialized glfw and gl before using this function
// and created a window
func DrawImage(windowWidth, windowHeight int, img image.Image, imageRect Rect) {
	imgProgram, shader1, shader2 := makeProgram(textureVertexShaderSrc, textureFragmentShaderSrc)
	vertices, indices := imageCoordinates(windowWidth, windowHeight, imageRect)

	VAO := makeImageVAO(vertices, indices)
	texture0, err := newTexture(img, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		panic(err.Error())
	}

	// draw vertices
	gl.UseProgram(imgProgram)
	// set texture0 to uniform0 in the fragment shader
	texture0.Bind(gl.TEXTURE0)
	uniform1 := getUniformLocation(imgProgram, "ourTexture0")
	texture0.SetUniform(uniform1)

	gl.BindVertexArray(VAO)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

	// free up memory
	gl.BindVertexArray(0)
	gl.DeleteVertexArrays(1, &VAO)

	texture0.UnBind()
	texture0.Delete()

	gl.DeleteProgram(imgProgram)
	gl.DeleteShader(shader1)
	gl.DeleteShader(shader2)

}

// useful for mouse events
func InRect(aRect Rect, xPos, yPos int) bool {
	if (xPos > aRect.OriginX) && (xPos < aRect.Width+aRect.OriginX) &&
		(yPos > aRect.OriginY) && (yPos < aRect.Height+aRect.OriginY) {
		return true
	}

	return false
}

type Rect struct {
	Width   int
	Height  int
	OriginX int
	OriginY int
}

func NewRect(x, y, w, h int) Rect {
	return Rect{Width: w, Height: h, OriginX: x, OriginY: y}
}

// the outputs of this is good for gl.DrawElements
func rectangleToCoords2(windowWidth, windowHeight int, aRect Rect) ([]float32, []uint32) {

	point1X := xToFloat(aRect.OriginX, windowWidth)
	point1Y := yToFloat(aRect.OriginY, windowHeight)

	point2X := xToFloat(aRect.OriginX+aRect.Width, windowWidth)
	point2Y := yToFloat(aRect.OriginY+aRect.Height, windowHeight)

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
func imageCoordinates(windowWidth, windowHeight int, aRect Rect) ([]float32, []uint32) {
	tmpVertices, indices := rectangleToCoords2(windowWidth, windowHeight, aRect)
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
