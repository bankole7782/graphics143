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

	window := g143.NewWindow(800, 600, "a text display sample program", false)
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
	g143.DrawRectangle(wWidth, wHeight, "#dddddd", basics.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0})

	text1 := "OpenGL Text 1"
	textWidth1 := g143.MeasureText(text1, g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE)
	trs1 := basics.RectSpecs{Width: textWidth1, Height: 40, OriginX: 50, OriginY: 50}
	g143.DrawString(wWidth, wHeight, text1, "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE, trs1)

	text2 := "A wonderous day"
	textWidth2 := g143.MeasureText(text2, g143.DefaultFontBytes, 40)
	trs2 := basics.RectSpecs{Width: textWidth2, Height: 60, OriginX: 50, OriginY: 100}
	g143.DrawString(wWidth, wHeight, text2, "#444444", &g143.DefaultFontBytes, 40, trs2)

	window.SwapBuffers()
}
