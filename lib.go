package graphics143

import (
	"fmt"
	"math"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/pkg/errors"
)

func RectangleToCoords(spaceWidth, spaceHeight, rectWidth, rectHeight, originX, originY int) []float32 {

	point1X := XtoFloat(originX, spaceWidth)
	point1Y := YtoFloat(originY, spaceHeight)

	point2X := XtoFloat(originX+rectWidth, spaceWidth)
	point2Y := YtoFloat(originY+rectHeight, spaceHeight)

	retFloat32 := []float32{
		// first triangle
		point1X, point1Y, 0,
		point1X, point2Y, 0,
		point2X, point2Y, 0,

		// second triangle
		point1X, point1Y, 0,
		point2X, point1Y, 0,
		point2X, point2Y, 0,
	}

	return retFloat32
}

func XtoFloat(x, width int) float32 {
	return float32(2.0)*(float32(x)/float32(width)) - float32(1.0)
}

func YtoFloat(y, height int) float32 {
	return float32(1.0) - float32(2.0)*float32(y)/float32(height)
}

func PrintF32Arr(arr []float32) {
	rem := math.Mod(float64(len(arr)), 3.0)
	if int(rem) != 0 {
		panic("supplied array is not a multiple of 3")
	}

	for i := 0; i < len(arr)/3; i++ {
		fmt.Println(arr[i], arr[i+1], arr[i+2])
		rem := math.Mod(float64(i)+1, float64(3))
		if int(rem) == 0 {
			fmt.Println()
		}
	}
}

func GetColorShader(hexColor string) (string, error) {

	c, err := colorful.Hex(hexColor)
	if err != nil {
		return "", errors.Wrap(err, "colorful error")
	}

	r, g, b, a := c.RGBA()
	rNormalized := float32(r) / 65535.0
	gNormalized := float32(g) / 65535.0
	bNormalized := float32(b) / 65535.0
	aNormalized := float32(a) / 65535.0

	fragmentShaderSource := fmt.Sprintf(`
		#version 460
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(%f, %f, %f, %f);
		}
	`, rNormalized, gNormalized, bNormalized, aNormalized)

	return fragmentShaderSource + "\x00", nil
}
