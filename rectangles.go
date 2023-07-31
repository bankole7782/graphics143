// You must have initialized glfw and gl before using any function here function
package graphics143

import (
	"image"
	"unsafe"

	"github.com/bankole7782/graphics143/basics"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type BorderSide int

const (
	TOP BorderSide = iota
	LEFT
	BOTTOM
	RIGHT
)

func DrawRectangle(windowWidth, windowHeight int, hexColor string, rectSpecs basics.RectSpecs) {
	fragmentShaderSource, _ := basics.GetRectColorShader(hexColor)
	rectProgram, shader1, shader2 := basics.MakeProgram(basics.BasicVertexShaderSource, fragmentShaderSource)
	rectVertices := basics.RectangleToCoords(windowWidth, windowHeight, rectSpecs)
	rectVAO := basics.MakeBasicVao(rectVertices)

	gl.UseProgram(rectProgram)
	gl.BindVertexArray(rectVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(rectVertices)/3))

	gl.DeleteProgram(rectProgram)
	gl.DeleteShader(shader1)
	gl.DeleteShader(shader2)
	gl.DeleteVertexArrays(1, &rectVAO)
	gl.BindVertexArray(0)
}

func DrawRectangleGradient(windowWidth, windowHeight int, hexColor1, hexColor2 string, directionX bool, rectSpecs basics.RectSpecs) {
	fragmentShaderSource, _ := basics.GetRectGradientShader(hexColor1, hexColor2, directionX, windowWidth, windowHeight, rectSpecs)
	rectProgram, shader1, shader2 := basics.MakeProgram(basics.BasicVertexShaderSource, fragmentShaderSource)
	rectVertices := basics.RectangleToCoords(windowWidth, windowHeight, rectSpecs)
	rectVAO := basics.MakeBasicVao(rectVertices)

	gl.UseProgram(rectProgram)
	gl.BindVertexArray(rectVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(rectVertices)/3))

	gl.DeleteProgram(rectProgram)
	gl.DeleteShader(shader1)
	gl.DeleteShader(shader2)
	gl.DeleteVertexArrays(1, &rectVAO)
	gl.BindVertexArray(0)
}

func GetBorderSideRectangle(rectSpec basics.RectSpecs, borderSide BorderSide, borderDepth int) basics.RectSpecs {
	if borderSide == LEFT {
		return basics.RectSpecs{Width: borderDepth, Height: rectSpec.Height, OriginX: rectSpec.OriginX, OriginY: rectSpec.OriginY}
	} else if borderSide == TOP {
		return basics.RectSpecs{Width: rectSpec.Width, Height: borderDepth, OriginX: rectSpec.OriginX, OriginY: rectSpec.OriginY}
	} else if borderSide == RIGHT {
		return basics.RectSpecs{Width: borderDepth, Height: rectSpec.Height, OriginX: rectSpec.OriginX + rectSpec.Width - borderDepth, OriginY: rectSpec.OriginY}
	} else {
		return basics.RectSpecs{Width: rectSpec.Width, Height: borderDepth, OriginX: rectSpec.OriginX, OriginY: rectSpec.OriginY + rectSpec.Height - borderDepth}
	}
}

func GetInsetRectangle(rectSpec basics.RectSpecs, borderDepth int) basics.RectSpecs {
	return basics.RectSpecs{
		Width:   rectSpec.Width - 2*borderDepth,
		Height:  rectSpec.Height - 2*borderDepth,
		OriginX: rectSpec.OriginX + borderDepth,
		OriginY: rectSpec.OriginY + borderDepth,
	}
}

func DrawImage(windowWidth, windowHeight int, img image.Image, imageRectSpecs basics.RectSpecs) {
	imgProgram, shader1, shader2 := basics.MakeProgram(basics.TextureVertexShaderSrc, basics.TextureFragmentShaderSrc)
	vertices, indices := basics.ImageCoordinates(windowWidth, windowHeight, imageRectSpecs)

	VAO := basics.MakeImageVAO(vertices, indices)
	texture0, err := basics.NewTexture(img, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		panic(err.Error())
	}

	// draw vertices
	gl.UseProgram(imgProgram)
	// set texture0 to uniform0 in the fragment shader
	texture0.Bind(gl.TEXTURE0)
	uniform1 := basics.GetUniformLocation(imgProgram, "ourTexture0")
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
