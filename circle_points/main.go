package main

import (
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 800
	height = 400

	fps = 10
)

func main() {

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "bordered circle", true)
	window.SetFramebufferSizeCallback(frameBufferSizeCallback)
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.MULTISAMPLE)

	// gl.Enable(gl.LINE_SMOOTH)
	// gl.Enable(gl.POLYGON_SMOOTH)
	// gl.Enable(gl.POINT)

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

	pointVertices := []float32{g143.XtoFloat(wWidth/2, wWidth), g143.YtoFloat(wHeight/2, wHeight), 0}
	pvVao := makeVao(pointVertices)

	pointFragmentSource, _ := g143.GetPointShader("#aaaaaa")
	pointFragmentSource2, _ := g143.GetPointShader("#666666")

	pt1Program := g143.MakeProgram(g143.BasicVertexShaderSource, pointFragmentSource)
	pt2Program := g143.MakeProgram(g143.BasicVertexShaderSource, pointFragmentSource2)

	gl.PointSize(250)
	draw(pvVao, pt2Program, pointVertices)

	gl.PointSize(200)
	draw(pvVao, pt1Program, pointVertices)

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
