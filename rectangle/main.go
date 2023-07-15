package main

import (
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.6-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 800
	height = 600

	fps = 10

	vertexShaderSource = `
		#version 460
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"
)

func main() {

	rect1 := g143.RectangleToCoords(500, 500, 100, 200, 20, 20)
	rect2 := g143.RectangleToCoords(500, 500, 100, 200, 140, 20)

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "two rectangles")

	if err := gl.Init(); err != nil {
		panic(err)
	}

	program := makeProgram()

	vao := g143.MakeVao(rect1)
	vao2 := g143.MakeVao(rect2)
	for !window.ShouldClose() {
		t := time.Now()
		draw([]uint32{vao, vao2}, window, program, [][]float32{rect1, rect2})
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
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

func makeProgram() uint32 {
	vertexShader, err := g143.CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShaderSource, _ := g143.GetColorShader("#805F5F")
	fragmentShader, err := g143.CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}
