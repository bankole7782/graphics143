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

	window := g143.NewWindow(width, height, "two bordered rectangles", false)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	fragmentShaderSource, _ := g143.GetRectColorShader("#D4D7BC")
	mainRectShaders := []g143.ShaderDef{
		{Source: g143.BasicVertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: fragmentShaderSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	mainRectProgram := g143.MakeProgram(mainRectShaders)
	rect1Specs := g143.RectSpecs{Width: 300, Height: 200, OriginX: 100, OriginY: 50}
	rect1 := g143.RectangleToCoords(width, height, rect1Specs)
	vao1 := makeVao(rect1)

	borderFragmentShaderSource, _ := g143.GetRectColorShader("#61636A")
	borderRectShaders := []g143.ShaderDef{
		{Source: g143.BasicVertexShaderSource, ShaderType: gl.VERTEX_SHADER},
		{Source: borderFragmentShaderSource, ShaderType: gl.FRAGMENT_SHADER},
	}
	borderRectProgram := g143.MakeProgram(borderRectShaders)
	borderRectsSpecs := g143.GetBorderRectangles(rect1Specs, 10)
	borderVaos := make([]uint32, 0)
	borderVbos := make([][]float32, 0)

	for _, rectSpec := range borderRectsSpecs {
		rectVBO := g143.RectangleToCoords(width, height, rectSpec)
		rectVAO := makeVao(rectVBO)
		borderVbos = append(borderVbos, rectVBO)
		borderVaos = append(borderVaos, rectVAO)
	}

	// bordered rectangle two
	rect2Specs := g143.RectSpecs{Width: 300, Height: 200, OriginX: 100, OriginY: 300}
	rect2Vbo := g143.RectangleToCoords(width, height, rect2Specs)
	rect2vao := makeVao(rect2Vbo)

	leftBorderSpec := g143.GetBorderSideRectangle(rect2Specs, g143.LEFT, 10)
	leftBorderVbo := g143.RectangleToCoords(width, height, leftBorderSpec)
	leftBorderVao := makeVao(leftBorderVbo)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)

		t := time.Now()
		// draw first rectangle
		draw([]uint32{vao1}, mainRectProgram, [][]float32{rect1})
		// draw first rectangles borders
		draw(borderVaos, borderRectProgram, borderVbos)

		// draw second rectangle
		draw([]uint32{rect2vao}, mainRectProgram, [][]float32{rect2Vbo})
		// draw left border
		draw([]uint32{leftBorderVao}, borderRectProgram, [][]float32{leftBorderVbo})

		glfw.PollEvents()
		window.SwapBuffers()
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func draw(vaos []uint32, program uint32, vertices [][]float32) {

	gl.UseProgram(program)

	for i, vao := range vaos {
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices[i])/3))
	}

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
