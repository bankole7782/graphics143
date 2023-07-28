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

	window := g143.NewWindow(800, 600, "a text display sample program", false)

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

	text1 := "OpenGL Text 1"
	textWidth1 := g143.MeasureText(text1, g143.DefaultFontBytes, 12)
	trs1 := basics.RectSpecs{Width: textWidth1, Height: 40, OriginX: 50, OriginY: 50}
	g143.DrawString(wWidth, wHeight, text1, "#8C5555", g143.DefaultFontBytes, 12, trs1)

	text2 := "A wonderous day"
	textWidth2 := g143.MeasureText(text2, g143.DefaultFontBytes, 24)
	trs2 := basics.RectSpecs{Width: textWidth2, Height: 50, OriginX: 50, OriginY: 100}
	g143.DrawString(wWidth, wHeight, text2, "#BE9898", g143.DefaultFontBytes, 24, trs2)

	glfw.PollEvents()
	window.SwapBuffers()
}
