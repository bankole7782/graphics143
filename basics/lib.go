package basics

func XtoFloat(x, windowWidth int) float32 {
	return float32(2.0)*(float32(x)/float32(windowWidth)) - float32(1.0)
}

func YtoFloat(y, windowHeight int) float32 {
	return float32(1.0) - (float32(2.0) * float32(y) / float32(windowHeight))
}

func GLToX(x float32, screenWidth int) int {
	return int((x + 1.0) * float32(screenWidth-1) / 2.0)
}

func GLToY(y float32, screenHeight int) int {
	return int((1.0 - y) * float32(screenHeight-1) / 2.0)
}
