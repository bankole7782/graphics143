package main

import (
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/graphics143/basics"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10
)

func main() {
	runtime.LockOSThread()

	window := g143.NewWindow(800, 600, "a bordered circle", false)
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

	rs3 := basics.RectSpecs{Width: 50, Height: 50, OriginX: 50, OriginY: 50}

	g143.DrawRectangle(wWidth, wHeight, "#6666aa", rs3)
	g143.DrawCircle(wWidth, wHeight, "#aaaaaa", 25, 50, 50)

	rs1 := basics.RectSpecs{Width: 210, Height: 102, OriginX: 115, OriginY: 210}
	// g143.DrawRectangle(wWidth, wHeight, "#444444", rs1)
	g143.DrawRoundedRectangle(wWidth, wHeight, "#758571", rs1, 15)

	rs2 := basics.RectSpecs{Width: 210, Height: 102, OriginX: 375, OriginY: 210}
	g143.DrawRoundedRectangle(wWidth, wHeight, "#666666", rs2, 51)

	window.SwapBuffers()
}
