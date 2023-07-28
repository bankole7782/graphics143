package main

import (
	"image"
	"os"
	"runtime"
	"time"

	_ "image/jpeg"
	_ "image/png"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/graphics143/basics"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10
)

func main() {
	if len(os.Args) == 1 {
		panic("Expecting a picture path as the only argument")
	}

	runtime.LockOSThread()

	window := g143.NewWindow(800, 600, "an image viewer", false)
	window.SetFramebufferSizeCallback(frameBufferSizeCallback)

	for !window.ShouldClose() {
		t := time.Now()

		allDraws(window)
		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}

}

func allDraws(window *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	wWidth, wHeight := window.GetSize()

	imgFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	// Decode detexts the type of image as long as its image/<type> is imported
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	irs := basics.RectSpecs{Width: 400, Height: 400, OriginX: 50, OriginY: 50}

	g143.DrawImage(wWidth, wHeight, img, irs)

	glfw.PollEvents()
	window.SwapBuffers()
}

func frameBufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	allDraws(w)
}
