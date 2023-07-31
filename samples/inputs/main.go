package main

import (
	"fmt"
	"runtime"
	"time"

	_ "image/jpeg"
	_ "image/png"

	g143 "github.com/bankole7782/graphics143"
	"github.com/bankole7782/graphics143/basics"
	"github.com/go-gl/glfw/v3.3/glfw"
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

func main() {
	runtime.LockOSThread()

	objCoords = make(map[basics.RectSpecs]any)

	window := g143.NewWindow(800, 600, "an inputs program (sample)", false)
	allDraws(window)

	// respond to the mouse
	window.SetMouseButtonCallback(mouseBtnCallback)
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

	var objRS basics.RectSpecs
	var obj any

	for rs, anyObj := range objCoords {
		if g143.InRectSpecs(rs, xPosInt, yPosInt) {
			objRS = rs
			obj = anyObj
			break
		}
	}

	_ = objRS

	if obj != nil {
		switch widgetClass := obj.(type) {
		case ImagePicker:
			fmt.Println("image picker")
		case DoneBtn:
			fmt.Println("done btn")
		case TextEntry:
			fmt.Println("text entry")
			fmt.Println(widgetClass.Index)
		}
	}
}

func allDraws(window *glfw.Window) {
	wWidth, wHeight := window.GetSize()
	// background rectangle
	g143.DrawRectangle(wWidth, wHeight, "#ffffff", basics.RectSpecs{Width: wWidth, Height: wHeight, OriginX: 0, OriginY: 0})

	// intro text
	introText := "An Inputs Program (Sample)"
	textWidth2 := g143.MeasureText(introText, g143.DefaultFontBytes, 60)
	trs2 := basics.RectSpecs{Width: textWidth2, Height: 80, OriginX: 10, OriginY: 10}
	g143.DrawString(wWidth, wHeight, introText, "#444444", &g143.DefaultFontBytes, 40, trs2)

	// passport input
	prs := basics.RectSpecs{Width: 200, Height: 200, OriginX: 10, OriginY: 100}
	g143.DrawRectangle(wWidth, wHeight, "#dddddd", prs)
	passportMsgText := []string{"Click to ", "pick passport"}
	tprs := g143.GetInsetRectangle(prs, 20)
	g143.DrawString(wWidth, wHeight, passportMsgText[0], "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE, tprs)
	tprs2 := tprs
	tprs2.OriginY += 30
	g143.DrawString(wWidth, wHeight, passportMsgText[1], "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE, tprs2)

	objCoords[prs] = ImagePicker{}

	// other inputs

	fields := []string{"Name:", "Age:"}

	longestFieldSize := g143.MeasureText(fields[0], g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE)
	for i, f := range fields {
		// inputs label
		ftextSize := g143.MeasureText(f, g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE)
		trs2 := basics.RectSpecs{Width: ftextSize, Height: 30, OriginX: 240, OriginY: i*60 + 100}
		g143.DrawString(wWidth, wHeight, f, "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE, trs2)

		// inputs box
		ibrs := basics.RectSpecs{Width: 300, Height: 40, OriginX: trs2.OriginX + longestFieldSize + 20,
			OriginY: trs2.OriginY}

		g143.DrawRectangle(wWidth, wHeight, "#dddddd", ibrs)
		insetIbrs := g143.GetInsetRectangle(ibrs, 2)
		g143.DrawRectangle(wWidth, wHeight, "#ffffff", insetIbrs)

		objCoords[ibrs] = TextEntry{i}
	}

	// final button
	btnText := "Submit"
	btnTextSize := g143.MeasureText(btnText, g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE*2)
	bgBtnTextRS := basics.RectSpecs{Width: btnTextSize + 40, Height: 80, OriginY: 320}
	bgBtnTextRS.OriginX = (wWidth - bgBtnTextRS.Width) / 2

	btnTextRS := g143.GetInsetRectangle(bgBtnTextRS, 20)

	g143.DrawRectangle(wWidth, wHeight, "#D5A2A2", bgBtnTextRS)
	g143.DrawString(wWidth, wHeight, btnText, "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE*2, btnTextRS)

	objCoords[bgBtnTextRS] = DoneBtn{}
	window.SwapBuffers()
}
