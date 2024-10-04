package main

import (
	"runtime"
	"time"

	_ "image/jpeg"
	_ "image/png"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10

	BEGIN_WIDTH  = 1000
	BEGIN_HEIGHT = 600
)

var objCoords map[int]g143.Rect

func main() {
	runtime.LockOSThread()

	objCoords = make(map[int]g143.Rect)

	window := g143.NewWindow(BEGIN_WIDTH, BEGIN_HEIGHT, "a resizable program sample", true)
	allDraws(window)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)
	// necessary for resizing
	window.SetFramebufferSizeCallback(frameBufferSizeCallback)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}

}

func allDraws(window *glfw.Window) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	fRW := 800
	sRW := 300

	if wWidth >= fRW+sRW+20 {
		fRX := (wWidth - (fRW + sRW + 20)) / 2
		//first rectangle
		ggCtx.SetHexColor("#65805E")
		ggCtx.DrawRectangle(float64(fRX), 10, float64(fRW), 400)
		ggCtx.Fill()

		// second rectangle
		ggCtx.SetHexColor("#C99481")
		ggCtx.DrawRectangle(float64(fRW+fRX)+20, 10, float64(sRW), 200)
		ggCtx.Fill()

	} else {
		// first rectangle
		ggCtx.SetHexColor("#65805E")
		ggCtx.DrawRectangle(10, 10, float64(fRW), 300)
		ggCtx.Fill()

		// second rectangle
		ggCtx.SetHexColor("#C99481")
		ggCtx.DrawRectangle(10, 310+10, float64(sRW), 200)
		ggCtx.Fill()
	}
	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()
}

func frameBufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	allDraws(w)
}

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	if action != glfw.Release {
		return
	}
	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	// wWidth, wHeight := window.GetSize()

	// var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range objCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			// widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

}
