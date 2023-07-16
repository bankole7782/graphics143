package graphics143

import (
	"fmt"
	"math"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lucasb-eyer/go-colorful"
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

func CircleCoords(windowWidth, windowHeight, originX, originY, radius int) []float32 {
	twicePi := 2 * math.Pi
	triangleAmount := 128

	radiusX := float64(radius) / float64(windowWidth)
	originXf32 := XtoFloat(originX, windowWidth)
	originYf32 := YtoFloat(originY, windowHeight)

	vertices := make([]float32, 0)
	// vertices = append(vertices, originX, originY, 0)
	for i := 0; i < triangleAmount; i++ {
		x := originXf32 + float32(radiusX*math.Cos(float64(i)*twicePi/float64(triangleAmount)))
		y := originYf32 + float32(radiusX*math.Sin(float64(i)*twicePi/float64(triangleAmount)))
		vertices = append(vertices, x, y, 0)
	}

	return vertices
}

func ConvertColorToShaderFloats(hexColor string) (float32, float32, float32, float32) {
	c, err := colorful.Hex(hexColor)
	if err != nil {
		panic(err)
	}

	r, g, b, a := c.RGBA()
	rNormalized := float32(r) / 65535.0
	gNormalized := float32(g) / 65535.0
	bNormalized := float32(b) / 65535.0
	aNormalized := float32(a) / 65535.0

	return rNormalized, gNormalized, bNormalized, aNormalized
}

func GetRectColorShader(hexColor string) (string, error) {
	rf, gf, bf, af := ConvertColorToShaderFloats(hexColor)
	fragmentShaderSource := fmt.Sprintf(`
		#version 460
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(%f, %f, %f, %f);
		}
	`, rf, gf, bf, af)

	return fragmentShaderSource + "\x00", nil
}

func GetPointShader(hexColor string) (string, error) {
	rf, gf, bf, af := ConvertColorToShaderFloats(hexColor)
	circlePointFragmentSource := fmt.Sprintf(`
	#version 460
	out vec4 frag_colour;
	void main() {
		frag_colour = vec4(%f, %f, %f, %f);

		vec2 coord = gl_PointCoord - vec2(0.5); 
		if(length(coord) > 0.5)
			discard;
	}
	`, rf, gf, bf, af)

	return circlePointFragmentSource + "\x00", nil
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

// makeVao initializes and returns a vertex array from the points provided.
func MakeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
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
