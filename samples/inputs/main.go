package main

import (
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

func main() {
	runtime.LockOSThread()

	window := g143.NewWindow(800, 600, "an inputs program (sample)", false)
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

	// other inputs

	fields := []string{"Name:", "Age:"}

	longestFieldSize := g143.MeasureText(fields[0], g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE)
	for i, f := range fields {
		// inputs label
		ftextSize := g143.MeasureText(f, g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE)
		trs2 := basics.RectSpecs{ftextSize, 30, 240, i*60 + 100}
		g143.DrawString(wWidth, wHeight, f, "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE, trs2)

		// inputs box
		ibrs := basics.RectSpecs{Width: 300, Height: 40, OriginX: trs2.OriginX + longestFieldSize + 20,
			OriginY: trs2.OriginY}

		g143.DrawRectangle(wWidth, wHeight, "#dddddd", ibrs)
		insetIbrs := g143.GetInsetRectangle(ibrs, 2)
		g143.DrawRectangle(wWidth, wHeight, "#ffffff", insetIbrs)
	}

	// final button
	btnText := "Submit"
	btnTextSize := g143.MeasureText(btnText, g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE*2)
	bgBtnTextRS := basics.RectSpecs{Width: btnTextSize + 40, Height: 80, OriginY: 320}
	bgBtnTextRS.OriginX = (wWidth - bgBtnTextRS.Width) / 2

	btnTextRS := g143.GetInsetRectangle(bgBtnTextRS, 20)

	g143.DrawRectangle(wWidth, wHeight, "#D5A2A2", bgBtnTextRS)
	g143.DrawString(wWidth, wHeight, btnText, "#444444", &g143.DefaultFontBytes, g143.DEFAULT_FONT_SIZE*2, btnTextRS)

	window.SwapBuffers()
}
