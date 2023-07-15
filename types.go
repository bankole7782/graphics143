package graphics143

type RectSpecs struct {
	Width   int
	Height  int
	OriginX int
	OriginY int
}

type ShaderDef struct {
	Source     string
	ShaderType uint32
}

type BorderSide int

const (
	TOP BorderSide = iota
	LEFT
	BOTTOM
	RIGHT
)
