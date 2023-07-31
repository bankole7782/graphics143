package main

import (
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/graphics143/basics"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10
)

func main() {
	runtime.LockOSThread()

	window := g143.NewWindow(800, 600, "many rectangles", true)
	window.SetFramebufferSizeCallback(frameBufferSizeCallback)
	allDraws(window)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func allDraws(window *glfw.Window) {
	wWidth, wHeight := window.GetSize()
	// background rectangle
	g143.DrawRectangle(wWidth, wHeight, "#ffffff", basics.RectSpecs{wWidth, wHeight, 0, 0})

	rs1 := basics.RectSpecs{200, 100, 50, 50}
	rs2 := basics.RectSpecs{200, 100, 300, 50}
	rs3 := basics.RectSpecs{200, 100, 50, 200}
	rs4 := g143.GetInsetRectangle(rs3, 5)

	rs5 := basics.RectSpecs{200, 100, 300, 200}
	rs6 := g143.GetBorderSideRectangle(rs5, g143.RIGHT, 5)

	g143.DrawRectangle(wWidth, wHeight, "#D4D7BC", rs1)
	g143.DrawRectangle(wWidth, wHeight, "#61636A", rs2)
	g143.DrawRectangle(wWidth, wHeight, "#61636A", rs3)
	g143.DrawRectangle(wWidth, wHeight, "#D4D7BC", rs4)
	g143.DrawRectangle(wWidth, wHeight, "#D4D7BC", rs5)
	g143.DrawRectangle(wWidth, wHeight, "#61636A", rs6)

	window.SwapBuffers()
}

func frameBufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	allDraws(w)
}
