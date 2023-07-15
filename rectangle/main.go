package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.6-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 500
	height = 500

	fps = 10

	vertexShaderSource = `
		#version 460
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 460
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(0, 0, 0, 1.0);
		}
	` + "\x00"
)

func main() {

	// fmt.Println("square from tutorial:")
	// PrintF32Arr(square)

	// fmt.Println()
	// fmt.Println("rect1:")
	rect1 := graphics143.RectangleToCoords(500, 500, 100, 200, 20, 20)
	rect2 := graphics143.RectangleToCoords(500, 500, 100, 200, 140, 20)
	// PrintF32Arr(rect1)

	runtime.LockOSThread()

	// window := initGlfw()
	window := graphics143.NewWindow(width, height, "two rectangles")

	if err := gl.Init(); err != nil {
		panic(err)
	}

	// defer glfw.Terminate()
	program := makeProgram()

	vao := makeVao(rect1)
	vao2 := makeVao(rect2)
	for !window.ShouldClose() {
		// draw(vao, window, program, rect1)
		t := time.Now()
		draw2([]uint32{vao, vao2}, window, program, [][]float32{rect1, rect2})
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

// func draw(vao uint32, window *glfw.Window, program uint32, vertices []float32) {
// 	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
// 	gl.UseProgram(program)

// 	gl.BindVertexArray(vao)
// 	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))

// 	glfw.PollEvents()
// 	window.SwapBuffers()
// }

func draw2(vaos []uint32, window *glfw.Window, program uint32, vertices [][]float32) {
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
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShaderSource, _ := graphics143.GetColorShader("#805F5F")
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	// fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	// if err != nil {
	// 	panic(err)
	// }

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
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

func compileShader(source string, shaderType uint32) (uint32, error) {
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
