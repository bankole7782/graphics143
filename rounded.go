package graphics143

import (
	"github.com/bankole7782/graphics143/basics"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func DrawCircle(windowWidth, windowHeight int, hexColor string, radius, originX, originY int) {
	pointVertices := []float32{basics.XtoFloat(originX+radius, windowWidth), basics.YtoFloat(originY+radius, windowHeight), 0}
	pointVAO := basics.MakeBasicVao(pointVertices)

	pointFragmentSource, _ := basics.GetPointShader(hexColor)
	pointProgram, shader1, shader2 := basics.MakeProgram(basics.BasicVertexShaderSource, pointFragmentSource)
	gl.PointSize(float32(radius) * 2)

	gl.UseProgram(pointProgram)
	gl.BindVertexArray(pointVAO)
	gl.DrawArrays(gl.POINTS, 0, int32(len(pointVertices)/3))

	// free memory
	gl.DeleteProgram(pointProgram)
	gl.DeleteShader(shader1)
	gl.DeleteShader(shader2)

	gl.BindVertexArray(0)
}

func DrawRoundedRectangle(windowWidth, windowHeight int, hexColor string, rectSpecs basics.RectSpecs, borderRadius int) {
	if borderRadius > rectSpecs.Height/2 {
		borderRadius = rectSpecs.Height / 2
	}

	mainRectSpecs := basics.RectSpecs{
		Width: rectSpecs.Width - 2*borderRadius, Height: rectSpecs.Height,
		OriginX: rectSpecs.OriginX + borderRadius, OriginY: rectSpecs.OriginY,
	}

	leftRectSpecs := basics.RectSpecs{
		Width: borderRadius, Height: rectSpecs.Height - 2*borderRadius,
		OriginX: rectSpecs.OriginX, OriginY: rectSpecs.OriginY + borderRadius,
	}

	rightRectSpecs := basics.RectSpecs{
		Width: borderRadius, Height: rectSpecs.Height - 2*borderRadius,
		OriginX: rectSpecs.OriginX + mainRectSpecs.Width + borderRadius, OriginY: rectSpecs.OriginY + borderRadius,
	}

	DrawRectangle(windowWidth, windowHeight, hexColor, mainRectSpecs)
	DrawRectangle(windowWidth, windowHeight, hexColor, leftRectSpecs)
	DrawRectangle(windowWidth, windowHeight, hexColor, rightRectSpecs)

	// left top circle
	DrawCircle(windowWidth, windowHeight, hexColor, borderRadius, rectSpecs.OriginX, rectSpecs.OriginY)
	// right top circle
	DrawCircle(windowWidth, windowHeight, hexColor, borderRadius, rectSpecs.OriginX+mainRectSpecs.Width,
		rectSpecs.OriginY)
	// left bottom circle
	DrawCircle(windowWidth, windowHeight, hexColor, borderRadius, rectSpecs.OriginX,
		rectSpecs.OriginY+leftRectSpecs.Height)
	// right bottom circle
	DrawCircle(windowWidth, windowHeight, hexColor, borderRadius, rectSpecs.OriginX+mainRectSpecs.Width,
		rectSpecs.OriginY+leftRectSpecs.Height)
}
