package main

import (
	"fmt"
	"image"
	"runtime"
	"strconv"
	"time"

	_ "image/jpeg"
	_ "image/png"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/graphics143/basics"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sqweek/dialog"
)

const (
	fps = 10
)

type ImagePicker struct {
}

type TextEntry struct {
	Index int
}

type DoneBtn struct {
}

var objCoords map[basics.RectSpecs]any
var currentWindowFrame image.Image
var inputsStore map[string]string
var activeEntryIndex int
var enteredText string

func main() {
	runtime.LockOSThread()

	objCoords = make(map[basics.RectSpecs]any)
	inputsStore = make(map[string]string)

	window := g143.NewWindow(800, 600, "an inputs program (sample)", false)
	allDraws(window)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)
	window.SetKeyCallback(keyCallback)

	for !window.ShouldClose() {
		t := time.Now()
		glfw.PollEvents()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}

}

func mouseBtnCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

	if action != glfw.Release {
		return
	}
	xPos, yPos := window.GetCursorPos()
	xPosInt := int(xPos)
	yPosInt := int(yPos)

	wWidth, wHeight := window.GetSize()

	var objRS basics.RectSpecs
	var obj any

	for rs, anyObj := range objCoords {
		if g143.InRectSpecs(rs, xPosInt, yPosInt) {
			objRS = rs
			obj = anyObj
			break
		}
	}

	if obj != nil {
		switch widgetClass := obj.(type) {
		case ImagePicker:
			filename, err := dialog.File().Filter("Passport file", "jpg").Load()
			if err != nil {
				return
			}
			inputsStore["passport"] = filename

			// show picked image
			ggCtx := gg.NewContextForImage(currentWindowFrame)

			img, _ := imaging.Open(filename)
			img = imaging.Resize(img, objRS.Width-20, objRS.Height-20, imaging.Lanczos)
			ggCtx.DrawImage(img, objRS.OriginX+10, objRS.OriginY+10)

			// send the frame to glfw window
			windowRS := basics.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
			g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
			window.SwapBuffers()

			// save the frame
			currentWindowFrame = ggCtx.Image()
		case DoneBtn:
			fmt.Println("done btn")
		case TextEntry:
			// switching where the keys would be placed.
			if widgetClass.Index != activeEntryIndex {
				oldText, ok := inputsStore[strconv.Itoa(widgetClass.Index)]
				inputsStore[strconv.Itoa(widgetClass.Index)] = enteredText
				if ok {
					enteredText = oldText
				} else {
					enteredText = ""
				}
			}
			activeEntryIndex = widgetClass.Index

			ggCtx := gg.NewContextForImage(currentWindowFrame)

			// draw border input
			ggCtx.SetHexColor("#63D171")

			ggCtx.DrawRectangle(float64(objRS.OriginX), float64(objRS.OriginY), float64(objRS.Width), float64(objRS.Height))
			ggCtx.Fill()

			ggCtx.SetHexColor("#fff")
			ggCtx.DrawRectangle(float64(objRS.OriginX+5), float64(objRS.OriginY+5), 340, 30)
			ggCtx.Fill()

			// send the frame to glfw window
			windowRS := basics.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
			g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
			window.SwapBuffers()

			// save the frame
			currentWindowFrame = ggCtx.Image()
		default:
			activeEntryIndex = 0
		}
	}
}

func isKeyNumeric(key glfw.Key) bool {
	numKeys := []glfw.Key{glfw.Key0, glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4,
		glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9}

	for _, numKey := range numKeys {
		if key == numKey {
			return true
		}
	}

	return false
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Release {
		return
	}

	wWidth, wHeight := window.GetSize()

	var objRS basics.RectSpecs
	for k, v := range objCoords {
		textEntry, ok := v.(TextEntry)
		if ok && textEntry.Index == activeEntryIndex {
			objRS = k
			break
		}
	}

	if activeEntryIndex == 1 {
		if key == glfw.KeyBackspace {
			enteredText = enteredText[:len(enteredText)-1]
		} else {
			enteredText += glfw.GetKeyName(key, scancode)
		}
	} else if activeEntryIndex == 2 {
		// enforce number types
		if isKeyNumeric(key) {
			enteredText += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeyBackspace {
			enteredText = enteredText[:len(enteredText)-1]
		}
	}

	ggCtx := gg.NewContextForImage(currentWindowFrame)
	err := ggCtx.LoadFontFace("Roboto-Light.ttf", 20)
	if err != nil {
		panic(err)
	}

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(objRS.OriginX+5), float64(objRS.OriginY+5), 340, 30)
	ggCtx.Fill()

	ggCtx.SetHexColor("#444444")
	ggCtx.DrawString(enteredText, float64(objRS.OriginX+25), float64(objRS.OriginY+25))

	// send the frame to glfw window
	windowRS := basics.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}

func allDraws(window *glfw.Window) {
	wWidth, wHeight := window.GetSize()

	// frame buffer
	ggCtx := gg.NewContext(wWidth, wHeight)

	// background rectangle
	ggCtx.DrawRectangle(0, 0, float64(wWidth), float64(wHeight))
	ggCtx.SetHexColor("#ffffff")
	ggCtx.Fill()

	// intro text
	err := ggCtx.LoadFontFace("Roboto-Light.ttf", 30)
	if err != nil {
		panic(err)
	}
	ggCtx.SetHexColor("#444444")
	introText := "An Inputs Program (Sample)"
	ggCtx.DrawString(introText, 20, 40)

	// draw passport Rectangle
	ggCtx.DrawRectangle(20, 70, 200, 200)
	ggCtx.SetHexColor("#dddddd")
	ggCtx.Fill()

	// save valuable coordinates
	prs := basics.RectSpecs{Width: 200, Height: 200, OriginX: 20, OriginY: 70}
	objCoords[prs] = ImagePicker{}

	err = ggCtx.LoadFontFace("Roboto-Light.ttf", 20)
	if err != nil {
		panic(err)
	}
	passportMsgText := []string{"Click to ", "pick passport"}
	ggCtx.SetHexColor("#444444")
	ggCtx.DrawString(passportMsgText[0], 40, 110)
	ggCtx.DrawString(passportMsgText[1], 40, 130)

	// other inputs
	labels := []string{"Name:", "Age:"}
	longestFieldX, _ := ggCtx.MeasureString(labels[0])
	for i, label := range labels {
		ggCtx.SetHexColor("#444444")
		ggCtx.DrawString(label, 260, 120+float64(i*50))

		// draw border input
		ggCtx.SetHexColor("#ddd")
		iBoxX, iBoxY := 260+longestFieldX+20, 100+float64(i*50)
		iBoxWidth, iBoxHeight := 350, 40
		ggCtx.DrawRectangle(iBoxX, iBoxY, float64(iBoxWidth), float64(iBoxHeight))
		ggCtx.Fill()

		ibrs := basics.RectSpecs{Width: iBoxWidth, Height: iBoxHeight, OriginX: int(iBoxX), OriginY: int(iBoxY)}
		objCoords[ibrs] = TextEntry{i + 1}

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawRectangle(260+longestFieldX+20+5, 100+5+float64(i*50), 340, 30)
		ggCtx.Fill()

	}

	// submit button
	err = ggCtx.LoadFontFace("Roboto-Light.ttf", 30)
	if err != nil {
		panic(err)
	}
	btnText := "Submit"
	btnTextWidth, btnTextHeight := ggCtx.MeasureString(btnText)
	btnBGWidth := wWidth - 100
	btnBGHeight := btnTextHeight + 40
	btnBGX := (float64(wWidth-btnBGWidth) / 2.0)

	ggCtx.SetHexColor("#D5A2A2")
	ggCtx.DrawRoundedRectangle(btnBGX, 300, float64(btnBGWidth), btnBGHeight, 30)
	ggCtx.Fill()
	ggCtx.SetHexColor("#fff")
	btnTextX := btnBGX + float64(float64(btnBGWidth)-btnTextWidth)/2.0
	ggCtx.DrawString(btnText, btnTextX, 300+40)

	bgBtnTextRS := basics.RectSpecs{Width: btnBGWidth, Height: int(btnBGX), OriginY: 300}
	objCoords[bgBtnTextRS] = DoneBtn{}

	// send the frame to glfw window
	windowRS := basics.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}
