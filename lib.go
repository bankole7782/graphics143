package graphics143

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func XtoFloat(x, windowWidth int) float32 {
	return float32(2.0)*(float32(x)/float32(windowWidth)) - float32(1.0)
}

func YtoFloat(y, windowHeight int) float32 {
	return float32(1.0) - (float32(2.0) * float32(y) / float32(windowHeight))
}

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

func GetBorderRectangles(rectSpec RectSpecs, borderDepth int) []RectSpecs {
	border1 := GetBorderSideRectangle(rectSpec, LEFT, borderDepth)
	border2 := GetBorderSideRectangle(rectSpec, TOP, borderDepth)
	border3 := GetBorderSideRectangle(rectSpec, RIGHT, borderDepth)
	border4 := GetBorderSideRectangle(rectSpec, BOTTOM, borderDepth)
	return []RectSpecs{border1, border2, border3, border4}
}

func GetBorderSideRectangle(rectSpec RectSpecs, borderSide BorderSide, borderDepth int) RectSpecs {
	if borderSide == LEFT {
		return RectSpecs{borderDepth, rectSpec.Height, rectSpec.OriginX, rectSpec.OriginY}
	} else if borderSide == TOP {
		return RectSpecs{rectSpec.Width, borderDepth, rectSpec.OriginX, rectSpec.OriginY}
	} else if borderSide == RIGHT {
		return RectSpecs{borderDepth, rectSpec.Height, rectSpec.OriginX + rectSpec.Width - borderDepth, rectSpec.OriginY}
	} else {
		return RectSpecs{rectSpec.Width, borderDepth, rectSpec.OriginX, rectSpec.OriginY + rectSpec.Height - borderDepth}
	}
}

// initGlfw initializes glfw and returns a Window to use.
func NewWindow(width, height int, title string, resizable bool) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	primaryMon := glfw.GetPrimaryMonitor()
	_, _, monWidth, monHeight := primaryMon.GetWorkarea()

	windowX := (monWidth - width) / 2
	windowY := (monHeight - height) / 2

	window.SetPos(windowX, windowY)

	return window
}

func MakeProgram(shaders []ShaderDef) uint32 {
	prog := gl.CreateProgram()
	for _, shaderSpec := range shaders {
		shader1, err := CompileShader(shaderSpec.Source, shaderSpec.ShaderType)
		if err != nil {
			panic(err)
		}
		gl.AttachShader(prog, shader1)
	}

	gl.LinkProgram(prog)
	return prog
}
