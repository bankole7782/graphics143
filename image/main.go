package main

import (
	"log"
	"os"
	"runtime"
	"unsafe"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 800
	height = 600
)

func main() {
	if len(os.Args) == 1 {
		panic("Expecting a picture path as the only argument")
	}

	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "image view", false)
	if err := gl.Init(); err != nil {
		panic(err)
	}

	err := programLoop(window)
	if err != nil {
		log.Fatal(err)
	}

}

/*
 * Creates the Vertex Array Object for a triangle.
 */
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
	var stride int32 = 3*4 + 2*4
	var offset int = 0

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(0)
	offset += 3 * 4

	// // color
	// gl.VertexAttribPointer(1, 3, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	// gl.EnableVertexAttribArray(1)
	// offset += 3 * 4

	// texture position
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(2)
	offset += 2 * 4

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}

func programLoop(window *glfw.Window) error {

	// // the linked shader program determines how the data will be rendered
	// vertShader, err := gfx.NewShaderFromFile("shaders/basic.vert", gl.VERTEX_SHADER)
	// if err != nil {
	// 	return err
	// }

	// fragShader, err := gfx.NewShaderFromFile("shaders/basic.frag", gl.FRAGMENT_SHADER)
	// if err != nil {
	// 	return err
	// }

	// shaderProgram, err := gfx.NewProgram(vertShader, fragShader)
	// if err != nil {
	// 	return err
	// }
	// defer shaderProgram.Delete()

	mainRectShaders := []g143.ShaderDef{
		{Source: g143.TextureVertexShaderSrc, ShaderType: gl.VERTEX_SHADER},
		{Source: g143.TextureFragmentShaderSrc, ShaderType: gl.FRAGMENT_SHADER},
	}
	shaderProgram := g143.MakeProgram(mainRectShaders)

	wWidth, wHeight := window.GetSize()
	rectSpec1 := g143.RectSpecs{Width: 300, Height: 400, OriginX: 50, OriginY: 50}
	v1, i1 := g143.RectangleToCoords2(wWidth, wHeight, rectSpec1)

	// inject texture coordinates
	vertices := []float32{
		v1[0], v1[1], v1[2],
		// 1.0, 0.0, // texture coordinates
		1.0, 0.0,

		v1[3], v1[4], v1[5],
		1.0, 1.0,

		v1[6], v1[7], v1[8],
		0.0, 1.0,

		v1[9], v1[10], v1[11],
		0.0, 0.0,
	}

	// verticess := []float32{
	// 	// top left
	// 	-0.75, 0.75, 0.0, // position
	// 	// 1.0, 0.0, 0.0, // Color
	// 	1.0, 0.0, // texture coordinates

	// 	// top right
	// 	0.75, 0.75, 0.0,
	// 	// 0.0, 1.0, 0.0,
	// 	0.0, 0.0,

	// 	// bottom right
	// 	0.75, -0.75, 0.0,
	// 	// 0.0, 0.0, 1.0,
	// 	0.0, 1.0,

	// 	// bottom left
	// 	-0.75, -0.75, 0.0,
	// 	// 1.0, 1.0, 1.0,
	// 	1.0, 1.0,
	// }

	// indices := []uint32{
	// 	// rectangle
	// 	0, 1, 2, // top triangle
	// 	0, 2, 3, // bottom triangle
	// }

	VAO := createVAO(vertices, i1)
	texture0, err := g143.NewTextureFromFile(os.Args[1], gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		panic(err.Error())
	}

	// texture1, err := gfx.NewTextureFromFile("../images/trollface.png",
	// 	gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	// if err != nil {
	// 	panic(err.Error())
	// }

	for !window.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		// background color
		// gl.ClearColor(0.2, 0.5, 0.5, 1.0)
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		// draw vertices
		gl.UseProgram(shaderProgram)
		// set texture0 to uniform0 in the fragment shader
		texture0.Bind(gl.TEXTURE0)
		// texture0.SetUniform(shaderProgram.GetUniformLocation("ourTexture0"))
		uniform1 := g143.GetUniformLocation(shaderProgram, "ourTexture0")
		texture0.SetUniform(uniform1)

		// // set texture1 to uniform1 in the fragment shader
		// texture1.Bind(gl.TEXTURE1)
		// texture1.SetUniform(shaderProgram.GetUniformLocation("ourTexture1"))

		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
		gl.BindVertexArray(0)

		texture0.UnBind()
		// texture1.UnBind()

		// end of draw loop

		// swap in the rendered buffer
		window.SwapBuffers()
	}

	return nil
}
