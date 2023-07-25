package main

import (
	"runtime"
	"time"
	"unsafe"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.1-core/gl"
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

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)

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
	rect1, rect1Indices := g143.RectangleToCoords2(wWidth, wHeight, g143.RectSpecs{Width: 100, Height: 200, OriginX: 20, OriginY: 20})
	rect2, rect2Indices := g143.RectangleToCoords2(wWidth, wHeight, g143.RectSpecs{Width: 100, Height: 200, OriginX: 140, OriginY: 20})
	// fmt.Println(rect1)
	// fmt.Println(rect1Indices)
	vao := createVAO(rect1, rect1Indices)
	// vao2 := makeVao(rect2)

	fragmentShaderSource, _ := g143.GetRectColorShader("#7B4747")

	mainRectProgram := g143.MakeProgram(g143.BasicVertexShaderSource, fragmentShaderSource)

	// draw([]uint32{vao}, window, mainRectProgram, [][]float32{rect1})
	gl.UseProgram(mainRectProgram)
	gl.BindVertexArray(vao)
	// gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices[i])/3))
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

	vao2 := createVAO(rect2, rect2Indices)
	gl.BindVertexArray(vao2)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))

	glfw.PollEvents()
	window.SwapBuffers()
}

func createVAO(vertices []float32, indices []uint32) uint32 {

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)

	var EBO uint32
	gl.GenBuffers(1, &EBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// copy indices into element buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// size of one whole vertex (sum of attrib sizes)
	// var stride int32 = 3*4 + 3*4 + 2*4
	var stride int32 = 3 * 4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}
