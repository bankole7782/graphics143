package main

import (
	"math"
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

	vertexShaderSource = `
		#version 460
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"
)

// func circleCoords()

func main() {

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "a circle")

	if err := gl.Init(); err != nil {
		panic(err)
	}

	// x := float32(0.3)
	// y := float32(0.3)
	radius := float32(.2)
	triangleAmount := 40

	twicePi := 2 * math.Pi

	vertices := make([]float32, 0)
	// vertices = append(vertices, x, y, 0)
	for i := 0; i < triangleAmount; i++ {
		x := radius * float32(math.Cos(float64(i)*twicePi/float64(triangleAmount)))
		y := radius * float32(math.Sin(float64(i)*twicePi/float64(triangleAmount)))
		vertices = append(vertices, x, y, 0)
	}

	fragmentShaderSource, _ := g143.GetColorShader("#000000")
	circleShaders := []g143.ShaderDef{
		{Source: vertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: fragmentShaderSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	circleProgram := g143.MakeProgram(circleShaders)

	vao := g143.MakeVao(vertices)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)

		t := time.Now()

		drawCircle(vao, circleProgram, vertices)
		glfw.PollEvents()
		window.SwapBuffers()
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func drawCircle(vao uint32, program uint32, vertices []float32) {
	gl.EnableVertexAttribArray(0)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(vertices)/3))

}
