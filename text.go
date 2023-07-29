package graphics143

import (
	_ "embed"
	"image"
	"image/draw"

	"github.com/bankole7782/graphics143/basics"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/font"
)

//go:embed "Roboto-Light.ttf"
var DefaultFontBytes []byte

const (
	DEFAULT_FONT_SIZE = 20
	dpi               = 72
)

func MeasureText(string1 string, fontBytes []byte, size float64) int {
	rgba := image.NewRGBA(image.Rect(0, 0, 1366, 768))

	fontParsed, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	fontDrawer := &font.Drawer{
		Dst: rgba,
		Src: image.Black,
		Face: truetype.NewFace(fontParsed, &truetype.Options{
			Size:    size,
			DPI:     float64(dpi),
			Hinting: font.HintingNone,
		}),
	}

	return fontDrawer.MeasureString(string1).Round()
}

func DrawString(windowWidth, windowHeight int, str, hexColor string, fontBytes *[]byte, size float64,
	strRectSpec basics.RectSpecs) {
	// truetype things
	parsedFont, err := truetype.Parse(*fontBytes)
	if err != nil {
		panic(err)
	}
	colorObj, err := colorful.Hex(hexColor)
	if err != nil {
		panic(err)
	}

	fg := image.NewUniform(colorObj)
	bg := image.Transparent

	rgba := image.NewRGBA(image.Rect(0, 0, strRectSpec.Width, strRectSpec.Height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(parsedFont)
	c.SetFontSize(float64(size))
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingFull)

	// Draw the text.
	pt := freetype.Pt(0, int(c.PointToFixed(size)>>6))
	_, err = c.DrawString(str, pt)
	if err != nil {
		panic(err)
	}

	DrawImage(windowWidth, windowHeight, rgba, strRectSpec)
}
