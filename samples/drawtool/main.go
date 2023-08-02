package main

import (
	"image"
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	fps = 10
)

var objCoords map[g143.RectSpecs]any
var currentWindowFrame image.Image

type PencilWidget struct{}

type EraserWidget struct{}

type SaveWidget struct{}

func main() {
	runtime.LockOSThread()

	objCoords = make(map[g143.RectSpecs]any)

	window := g143.NewWindow(1100, 600, "a draw tool (sample)", false)
	allDraws(window)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)

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
	ggCtx.SetHexColor("#ddd")
	ggCtx.Fill()

	// intro text
	err := ggCtx.LoadFontFace("Roboto-Light.ttf", 30)
	if err != nil {
		panic(err)
	}
	ggCtx.SetHexColor("#444444")
	introText := "A Draw Tool Program (Sample)"
	ggCtx.DrawString(introText, 20, 40)

	// draw the tools
	err = ggCtx.LoadFontFace("Roboto-Light.ttf", 20)
	if err != nil {
		panic(err)
	}

	// pencil tool
	ggCtx.SetHexColor("#DAC166")
	ggCtx.DrawRoundedRectangle(20, 60, 120, 200, 10)
	ggCtx.Fill()

	ggCtx.SetHexColor("#dddddd")
	ggCtx.DrawRectangle(30, 70, 100, 40)
	ggCtx.Fill()

	pencilRS := g143.RectSpecs{Width: 100, Height: 40, OriginX: 30, OriginY: 70}
	objCoords[pencilRS] = PencilWidget{}

	ggCtx.SetHexColor("#444444")
	ggCtx.DrawString("Pencil", 40, 100)

	// eraser tool
	ggCtx.SetHexColor("#dddddd")
	ggCtx.DrawRectangle(30, 130, 100, 40)
	ggCtx.Fill()

	eraserRS := g143.RectSpecs{Width: 100, Height: 40, OriginX: 30, OriginY: 130}
	objCoords[eraserRS] = EraserWidget{}

	ggCtx.SetHexColor("#444444")
	ggCtx.DrawString("Eraser", 40, 160)

	// save tool
	ggCtx.SetHexColor("#dddddd")
	ggCtx.DrawRectangle(30, 200, 100, 40)
	ggCtx.Fill()

	saveRS := g143.RectSpecs{Width: 100, Height: 40, OriginX: 30, OriginY: 200}
	objCoords[saveRS] = SaveWidget{}

	ggCtx.SetHexColor("#444444")
	ggCtx.DrawString("Save", 40, 230)

	// Canvas
	ggCtx.SetHexColor("#ffffff")
	ggCtx.DrawRectangle(200, 60, 800, 500)
	ggCtx.Fill()

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

}
