package main

import (
	"fmt"
	"math"
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
