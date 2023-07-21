package main

import (
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 800
	height = 400

	fps = 10
)

func main() {

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "a single pill", true)
	window.SetFramebufferSizeCallback(frameBufferSizeCallback)
	if err := gl.Init(); err != nil {
		panic(err)
	}

	for !window.ShouldClose() {
		t := time.Now()

		allDraws(window)
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func frameBufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	allDraws(w)
}

func allDraws(window *glfw.Window) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	wWidth, wHeight := window.GetSize()

	// draw edge circles
	pointVertices := []float32{g143.XtoFloat(100, wWidth), g143.YtoFloat(100, wHeight), 0}
	pvVao := makeVao(pointVertices)
	pointFragmentSource, _ := g143.GetPointShader("#626193")
	pt1Shaders := []g143.ShaderDef{
		{Source: g143.BasicVertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: pointFragmentSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	pt1Program := g143.MakeProgram(pt1Shaders)

	gl.PointSize(100)
	draw(pvVao, pt1Program, pointVertices)

	point2Vertices := []float32{g143.XtoFloat(300, wWidth), g143.YtoFloat(100, wHeight), 0}
	pv2Vao := makeVao(point2Vertices)

	gl.PointSize(100)
	draw(pv2Vao, pt1Program, point2Vertices)

	// draw connecting rectangle
	rect1Specs := g143.RectSpecs{Width: 200, Height: 100, OriginX: 100, OriginY: 50}
	rect1Vertices := g143.RectangleToCoords(wWidth, wHeight, rect1Specs)
	rectVao := makeVao(rect1Vertices)
	rectFragmentShaderSource, _ := g143.GetRectColorShader("#626193")
	rectShaders := []g143.ShaderDef{
		{Source: g143.BasicVertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: rectFragmentShaderSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	rectProgram := g143.MakeProgram(rectShaders)
	gl.UseProgram(rectProgram)
	gl.BindVertexArray(rectVao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(rect1Vertices)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

func draw(vao uint32, program uint32, vertices []float32) {
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.POINTS, 0, int32(len(vertices)/3))
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
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
