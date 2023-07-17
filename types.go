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

type Texture struct {
	handle  uint32
	target  uint32 // same target as gl.BindTexture(<this param>, ...)
	texUnit uint32 // Texture unit that is currently bound to ex: gl.TEXTURE0
}
