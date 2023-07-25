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
	height = 500

	fps = 10
)

func main() {

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "a circle", false)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	vertices := g143.CircleCoords(width, height, width/2, height/2, 200)
	vertices2 := g143.CircleCoords(width, height, width/2, height/2, 230)

	fragmentShaderSource, _ := g143.GetRectColorShader("#BB97B7")
	circleProgram1 := g143.MakeProgram(g143.BasicVertexShaderSource, fragmentShaderSource)
	fragmentShaderSource2, _ := g143.GetRectColorShader("#855980")
	circleProgram2 := g143.MakeProgram(g143.BasicVertexShaderSource, fragmentShaderSource2)

	vao := makeVao(vertices)
	vao2 := makeVao(vertices2)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)

		t := time.Now()

		drawCircle(vao2, circleProgram2, vertices2)
		drawCircle(vao, circleProgram1, vertices)

		glfw.PollEvents()
		window.SwapBuffers()
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func drawCircle(vao uint32, program uint32, vertices []float32) {
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(vertices)/3))

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
