package main

import (
	"image"
	"runtime"
	"time"

	_ "image/jpeg"
	_ "image/png"

	g143 "github.com/bankole7782/graphics143"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sqweek/dialog"
)

const (
	fps = 10

	ImagePicker   = 101
	TextEntryName = 102
	TextEntryAge  = 103
	DoneBtn       = 104
)

var objCoords map[int]g143.Rect
var currentWindowFrame image.Image
var inputsStore map[string]string
var activeEntryIndex int
var enteredTextName string
var enteredTextAge string

func main() {
	runtime.LockOSThread()

	objCoords = make(map[int]g143.Rect)
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
	prs := g143.Rect{Width: 200, Height: 200, OriginX: 20, OriginY: 70}
	objCoords[ImagePicker] = prs

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

		ibrs := g143.Rect{Width: iBoxWidth, Height: iBoxHeight, OriginX: int(iBoxX), OriginY: int(iBoxY)}
		if i == 0 {
			objCoords[TextEntryName] = ibrs
		} else {
			objCoords[TextEntryAge] = ibrs
		}

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

	bgBtnTextRS := g143.Rect{Width: btnBGWidth, Height: int(btnBGX), OriginY: 300}
	objCoords[DoneBtn] = bgBtnTextRS

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
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

	var widgetRS g143.Rect
	var widgetCode int

	for code, RS := range objCoords {
		if g143.InRect(RS, xPosInt, yPosInt) {
			widgetRS = RS
			widgetCode = code
			break
		}
	}

	if widgetCode == 0 {
		return
	}

	switch widgetCode {

	case ImagePicker:
		filename, err := dialog.File().Filter("Passport file", "jpg").Load()
		if err != nil {
			return
		}
		inputsStore["passport"] = filename

		// show picked image
		ggCtx := gg.NewContextForImage(currentWindowFrame)

		img, _ := imaging.Open(filename)
		img = imaging.Resize(img, widgetRS.Width-20, widgetRS.Height-20, imaging.Lanczos)
		ggCtx.DrawImage(img, widgetRS.OriginX+10, widgetRS.OriginY+10)

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()
	case DoneBtn:
		window.SetShouldClose(true)
	case TextEntryName, TextEntryAge:
		// switching where the keys would be placed.
		// if widgetClass.Index != activeEntryIndex {
		// 	oldText, ok := inputsStore[strconv.Itoa(widgetClass.Index)]
		// 	inputsStore[strconv.Itoa(widgetClass.Index)] = enteredText
		// 	if ok {
		// 		enteredText = oldText
		// 	} else {
		// 		enteredText = ""
		// 	}
		// }
		if widgetCode == TextEntryName {
			activeEntryIndex = 1
		} else if widgetCode == TextEntryAge {
			activeEntryIndex = 2
		}

		ggCtx := gg.NewContextForImage(currentWindowFrame)

		// draw border input
		ggCtx.SetHexColor("#63D171")

		ggCtx.DrawRectangle(float64(widgetRS.OriginX), float64(widgetRS.OriginY), float64(widgetRS.Width), float64(widgetRS.Height))
		ggCtx.Fill()

		ggCtx.SetHexColor("#fff")
		ggCtx.DrawRectangle(float64(widgetRS.OriginX+5), float64(widgetRS.OriginY+5), 340, 30)
		ggCtx.Fill()

		// send the frame to glfw window
		windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
		g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
		window.SwapBuffers()

		// save the frame
		currentWindowFrame = ggCtx.Image()
	default:
		activeEntryIndex = 0
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

	var widgetRS g143.Rect
	if activeEntryIndex == 1 {
		widgetRS = objCoords[TextEntryName]
	} else if activeEntryIndex == 2 {
		widgetRS = objCoords[TextEntryAge]
	}

	if activeEntryIndex == 1 {
		if key == glfw.KeyBackspace && len(enteredTextName) != 0 {
			enteredTextName = enteredTextName[:len(enteredTextName)-1]
		} else {
			enteredTextName += glfw.GetKeyName(key, scancode)
		}
	} else if activeEntryIndex == 2 {
		// enforce number types
		if isKeyNumeric(key) {
			enteredTextAge += glfw.GetKeyName(key, scancode)
		} else if key == glfw.KeyBackspace && len(enteredTextAge) != 0 {
			enteredTextAge = enteredTextAge[:len(enteredTextAge)-1]
		}
	}

	ggCtx := gg.NewContextForImage(currentWindowFrame)
	err := ggCtx.LoadFontFace("Roboto-Light.ttf", 20)
	if err != nil {
		panic(err)
	}

	ggCtx.SetHexColor("#fff")
	ggCtx.DrawRectangle(float64(widgetRS.OriginX+5), float64(widgetRS.OriginY+5), 340, 30)
	ggCtx.Fill()

	ggCtx.SetHexColor("#444444")
	if activeEntryIndex == 1 {
		ggCtx.DrawString(enteredTextName, float64(widgetRS.OriginX+25), float64(widgetRS.OriginY+25))
	} else if activeEntryIndex == 2 {
		ggCtx.DrawString(enteredTextAge, float64(widgetRS.OriginX+25), float64(widgetRS.OriginY+25))
	}

	// send the frame to glfw window
	windowRS := g143.Rect{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0}
	g143.DrawImage(wWidth, wHeight, ggCtx.Image(), windowRS)
	window.SwapBuffers()

	// save the frame
	currentWindowFrame = ggCtx.Image()
}
