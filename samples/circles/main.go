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

	window := g143.NewWindow(800, 600, "a bordered circle", false)

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

	g143.DrawCircle(wWidth, wHeight, "#aaaaaa", 50, 50, 50)
	rs1 := basics.RectSpecs{Width: 310, Height: 72, OriginX: 115, OriginY: 210}
	g143.DrawRoundedRectangle(wWidth, wHeight, "#666666", rs1, 72)

	rs2 := basics.RectSpecs{Width: 310, Height: 72, OriginX: 115, OriginY: 320}
	g143.DrawRoundedRectangle(wWidth, wHeight, "#666666", rs2, 60)

	glfw.PollEvents()
	window.SwapBuffers()
}
