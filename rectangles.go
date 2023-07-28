// You must have initialized glfw and gl before using any function here function
package graphics143

import (
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
	rectProgram := basics.MakeProgram(basics.BasicVertexShaderSource, fragmentShaderSource)
	rectVertices := basics.RectangleToCoords(windowWidth, windowHeight, rectSpecs)
	rectVAO := basics.MakeBasicVao(rectVertices)

	gl.UseProgram(rectProgram)
	gl.BindVertexArray(rectVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(rectVertices)/3))

	gl.BindVertexArray(0)
}

func GetBorderRectangles(rectSpec basics.RectSpecs, borderDepth int) []basics.RectSpecs {
	border1 := GetBorderSideRectangle(rectSpec, LEFT, borderDepth)
	border2 := GetBorderSideRectangle(rectSpec, TOP, borderDepth)
	border3 := GetBorderSideRectangle(rectSpec, RIGHT, borderDepth)
	border4 := GetBorderSideRectangle(rectSpec, BOTTOM, borderDepth)
	return []basics.RectSpecs{border1, border2, border3, border4}
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
