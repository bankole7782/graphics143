package main

import (
	"image"
	"image/draw"
	"log"
	"runtime"
	"unsafe"

	g143 "github.com/bankole7782/graphics143"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	width   = 800
	height  = 600
	DPI     = 72
	SIZE    = 40
	SPACING = 1.5
)

func main() {
	runtime.LockOSThread()

	window := g143.NewWindow(width, height, "a single text: text1", false)
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

	// texture position
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, stride, gl.PtrOffset(offset))
	gl.EnableVertexAttribArray(1)
	offset += 2 * 4

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}

func programLoop(window *glfw.Window) error {

	// OpenGL state
	// ------------
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// truetype things
	parsedFont, err := truetype.Parse(RobotoBytes)
	if err != nil {
		panic(err)
	}

	// Initialize the context.
	textWidth, textHeight := 640, 480
	// fg, bg := image.Black, image.Transparent
	fg, bg := image.Black, image.White

	// ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	rgba := image.NewRGBA(image.Rect(0, 0, textWidth, textHeight))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(DPI)
	c.SetFont(parsedFont)
	c.SetFontSize(SIZE)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingFull)

	text := "OpenGL Text 1"
	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(SIZE)>>6))
	_, err = c.DrawString(text, pt)
	if err != nil {
		panic(err)
	}
	pt.Y += c.PointToFixed(SIZE * SPACING)

	texture0, err := g143.NewTexture(rgba, gl.CLAMP_TO_EDGE, gl.CLAMP_TO_EDGE)
	if err != nil {
		panic(err.Error())
	}

	mainRectShaders := []g143.ShaderDef{
		{Source: g143.TextureVertexShaderSrc, ShaderType: gl.VERTEX_SHADER},
		{Source: g143.TextureFragmentShaderSrc, ShaderType: gl.FRAGMENT_SHADER},
	}
	shaderProgram := g143.MakeProgram(mainRectShaders)

	wWidth, wHeight := window.GetSize()
	rectSpec1 := g143.RectSpecs{Width: textWidth, Height: textHeight, OriginX: 50, OriginY: 50}
	vertices, indices := g143.ImageCoordinates(wWidth, wHeight, rectSpec1)

	VAO := createVAO(vertices, indices)

	for !window.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		// background color
		gl.ClearColor(g143.ConvertColorToShaderFloats("#D1B2B2"))

		gl.Clear(gl.COLOR_BUFFER_BIT)

		// draw vertices
		gl.UseProgram(shaderProgram)
		// set texture0 to uniform0 in the fragment shader
		texture0.Bind(gl.TEXTURE0)
		uniform1 := g143.GetUniformLocation(shaderProgram, "ourTexture0")
		texture0.SetUniform(uniform1)

		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, unsafe.Pointer(nil))
		gl.BindVertexArray(0)

		texture0.UnBind()

		// end of draw loop

		// swap in the rendered buffer
		window.SwapBuffers()
	}

	return nil
}
