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
	height = 600

	fps = 10
)

func main() {

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "two rectangles", true)
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
	wWidth, wHeight := window.GetSize()
	rect1 := g143.RectangleToCoords(wWidth, wHeight, g143.RectSpecs{Width: 100, Height: 200, OriginX: 20, OriginY: 20})
	rect2 := g143.RectangleToCoords(wWidth, wHeight, g143.RectSpecs{Width: 100, Height: 200, OriginX: 140, OriginY: 20})
	vao := makeVao(rect1)
	vao2 := makeVao(rect2)

	fragmentShaderSource, _ := g143.GetRectColorShader("#7B4747")
	mainRectShaders := []g143.ShaderDef{
		{Source: g143.BasicVertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: fragmentShaderSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	mainRectProgram := g143.MakeProgram(mainRectShaders)

	draw([]uint32{vao, vao2}, window, mainRectProgram, [][]float32{rect1, rect2})
}

func draw(vaos []uint32, window *glfw.Window, program uint32, vertices [][]float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.UseProgram(program)

	for i, vao := range vaos {
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices[i])/3))
	}

	glfw.PollEvents()
	window.SwapBuffers()
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
