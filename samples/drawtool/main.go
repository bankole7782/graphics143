package main

import (
	"image"
	"image/draw"
	"runtime"
	"time"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
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

type CanvasWidget struct{}

type CircleSpec struct {
	X int
	Y int
	R int
}

var drawnIndicators []CircleSpec
var activeTool string
var lastX, lastY float64 // used in drawing

func main() {
	runtime.LockOSThread()

	objCoords = make(map[g143.RectSpecs]any)
	drawnIndicators = make([]CircleSpec, 0)

	window := g143.NewWindow(1100, 600, "a draw tool (sample)", false)
	allDraws(window)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)
	// respond to mouse movement
	window.SetCursorPosCallback(cursorPosCallback)

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

	canvasRS := g143.RectSpecs{Width: 800, Height: 500, OriginX: 200, OriginY: 60}
	objCoords[canvasRS] = CanvasWidget{}

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var objRS g143.RectSpecs
	var obj any

	for rs, anyObj := range objCoords {
		if g143.InRectSpecs(rs, xPosInt, yPosInt) {
			objRS = rs
			obj = anyObj
			break
		}
	}

	if obj != nil {
		switch obj.(type) {

		case PencilWidget:

			ggCtx := gg.NewContextForImage(currentWindowFrame)

			activeTool = "P"

			// clear indicators
			for _, cs := range drawnIndicators {
				ggCtx.SetHexColor("#dddddd")
				ggCtx.DrawCircle(float64(cs.X), float64(cs.Y), float64(cs.R))
				ggCtx.Fill()
			}
			// draw an indicator on the active tool
			ggCtx.SetHexColor("#DAC166")
			ggCtx.DrawCircle(float64(objRS.OriginX+objRS.Width-20), float64(objRS.OriginY+20), 10)
			ggCtx.Fill()
			drawnIndicators = append(drawnIndicators, CircleSpec{X: objRS.OriginX + objRS.Width - 20, Y: objRS.OriginY + 20, R: 10})

			// send the frame to glfw window
			windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
			g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
			window.SwapBuffers()

			// save the frame
			currentWindowFrame = ggCtx.Image()

		case EraserWidget:
			ggCtx := gg.NewContextForImage(currentWindowFrame)

			activeTool = "E"

			// clear indicators
			for _, cs := range drawnIndicators {
				ggCtx.SetHexColor("#dddddd")
				ggCtx.DrawCircle(float64(cs.X), float64(cs.Y), float64(cs.R))
				ggCtx.Fill()
			}
			// draw an indicator on the active tool
			ggCtx.SetHexColor("#DAC166")
			ggCtx.DrawCircle(float64(objRS.OriginX+objRS.Width-20), float64(objRS.OriginY+20), 10)
			ggCtx.Fill()
			drawnIndicators = append(drawnIndicators, CircleSpec{X: objRS.OriginX + objRS.Width - 20, Y: objRS.OriginY + 20, R: 10})

			// send the frame to glfw window
			windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
			g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
			window.SwapBuffers()

			// save the frame
			currentWindowFrame = ggCtx.Image()

		case SaveWidget:
			ggCtx := gg.NewContextForImage(currentWindowFrame)
			activeTool = ""

			// clear indicators
			for _, cs := range drawnIndicators {
				ggCtx.SetHexColor("#dddddd")
				ggCtx.DrawCircle(float64(cs.X), float64(cs.Y), float64(cs.R))
				ggCtx.Fill()
			}

			var canvasRS g143.RectSpecs
			for rs, obj := range objCoords {
				_, ok := obj.(CanvasWidget)
				if ok {
					canvasRS = rs
					break
				}
			}

			newImageRect := image.Rect(0, 0, canvasRS.Width, canvasRS.Height)
			outImg := image.NewRGBA(newImageRect)
			draw.Draw(outImg, newImageRect, currentWindowFrame, image.Pt(canvasRS.OriginX, canvasRS.OriginY), draw.Src)

			imaging.Save(outImg, time.Now().Format("20060102T150405MST")+".png")
		default:

		}
	}

}

func cursorPosCallback(window *glfw.Window, xpos float64, ypos float64) {
	wWidth, wHeight := window.GetSize()

	ggCtx := gg.NewContextForImage(currentWindowFrame)

	var canvasRS g143.RectSpecs
	for rs, obj := range objCoords {
		_, ok := obj.(CanvasWidget)
		if ok {
			canvasRS = rs
			break
		}
	}

	currentMouseAction := window.GetMouseButton(glfw.MouseButtonLeft)

	if currentMouseAction == glfw.Release {
		lastX, lastY = 0.0, 0.0
	}

	if g143.InRectSpecs(canvasRS, int(xpos), int(ypos)) && currentMouseAction == glfw.Press {
		if activeTool == "P" {
			// draw circles
			ggCtx.SetHexColor("#222222")

			if int(lastX) != 0 {
				ggCtx.SetLineWidth(3)
				ggCtx.MoveTo(lastX, lastY)
				ggCtx.LineTo(xpos, ypos)
				ggCtx.Stroke()
			} else {
				ggCtx.DrawCircle(xpos, ypos, 3)
				ggCtx.Fill()
			}

			lastX, lastY = xpos, ypos
		} else {
			ggCtx.SetHexColor("#ffffff")
			ggCtx.DrawCircle(xpos, ypos, 10)
			ggCtx.Fill()
		}
	}

	// send the frame to glfw window
	windowRS := g143.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}
