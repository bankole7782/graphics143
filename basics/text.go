package basics

import (
	"image"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func MeasureText(string1 string, fontBytes []byte, size, dpi float64) int {
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
			DPI:     dpi,
			Hinting: font.HintingNone,
		}),
	}

	return fontDrawer.MeasureString(string1).Round()
}
