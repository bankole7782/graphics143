package basics

func XtoFloat(x, windowWidth int) float32 {
	return float32(2.0)*(float32(x)/float32(windowWidth)) - float32(1.0)
}

func YtoFloat(y, windowHeight int) float32 {
	return float32(1.0) - (float32(2.0) * float32(y) / float32(windowHeight))
}
