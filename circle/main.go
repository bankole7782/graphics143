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

	vertexShaderSource = `
		#version 460
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"
)

func main() {

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "a circle")

	if err := gl.Init(); err != nil {
		panic(err)
	}

	// originX := float32(0.3)
	// originY := float32(0.3)
	// radius := float32(.3)
	// triangleAmount := 120

	// twicePi := 2 * math.Pi

	// vertices := make([]float32, 0)
	// // vertices = append(vertices, originX, originY, 0)
	// for i := 0; i < triangleAmount; i++ {
	// 	x := originX + (radius * float32(math.Cos(float64(i)*twicePi/float64(triangleAmount))))
	// 	y := originY + (radius * float32(math.Sin(float64(i)*twicePi/float64(triangleAmount))))
	// 	vertices = append(vertices, x, y, 0)
	// }

	vertices := g143.CircleCoords(width, height, width/2, height/2, 100)
	vertices2 := g143.CircleCoords(width, height, width/2, height/2, 110)
	// vertices := CircleCoords(width, height, 0, 0, 100)
	// vertices2 := CircleCoords(width, height, 0, 0, 110)

	fragmentShaderSource, _ := g143.GetColorShader("#BB97B7")
	circleShaders1 := []g143.ShaderDef{
		{Source: vertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: fragmentShaderSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	circleProgram1 := g143.MakeProgram(circleShaders1)
	fragmentShaderSource2, _ := g143.GetColorShader("#855980")

	circleShaders2 := []g143.ShaderDef{
		{Source: vertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: fragmentShaderSource2, ShaderType: gl.FRAGMENT_SHADER},
	}
	circleProgram2 := g143.MakeProgram(circleShaders2)

	vao := g143.MakeVao(vertices)
	vao2 := g143.MakeVao(vertices2)

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
