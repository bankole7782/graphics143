package graphics143

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/pkg/errors"
)

func NewTextureFromFile(file string, wrapR, wrapS int32) (*Texture, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	// Decode detexts the type of image as long as its image/<type> is imported
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return NewTexture(img, wrapR, wrapS)
}

func NewTexture(img image.Image, wrapR, wrapS int32) (*Texture, error) {
	img = imaging.FlipH(img)

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, errors.New("unsupported stride")
	}

	var handle uint32
	gl.GenTextures(1, &handle)

	target := uint32(gl.TEXTURE_2D)
	internalFmt := int32(gl.RGBA)
	format := uint32(gl.RGBA)
	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	pixType := uint32(gl.UNSIGNED_BYTE)
	dataPtr := gl.Ptr(rgba.Pix)

	texture := Texture{
		handle: handle,
		target: target,
	}

	texture.Bind(gl.TEXTURE0)
	defer texture.UnBind()

	// set the texture wrapping/filtering options (applies to current bound texture obj)
	// TODO-cs
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_R, wrapR)
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_S, wrapS)
	gl.TexParameteri(texture.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR) // minification filter
	gl.TexParameteri(texture.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR) // magnification filter

	gl.TexImage2D(target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(texture.handle)

	return &texture, nil
}

func (tex *Texture) Bind(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(tex.target, tex.handle)
	tex.texUnit = texUnit
}

func (tex *Texture) UnBind() {
	tex.texUnit = 0
	gl.BindTexture(tex.target, 0)
}

func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errors.New("texture not bound")
	}
	gl.Uniform1i(uniformLoc, int32(tex.texUnit-gl.TEXTURE0))
	return nil
}

// the outputs of this is good for gl.DrawElements
func ImageCoordinates(windowWidth, windowHeight int, rectSpec RectSpecs) ([]float32, []uint32) {
	tmpVertices, indices := RectangleToCoords2(windowWidth, windowHeight, rectSpec)
	v1 := tmpVertices
	// inject texture coordinates
	vertices := []float32{
		v1[0], v1[1], v1[2], // vertices position
		1.0, 0.0, // texture coordinates

		v1[3], v1[4], v1[5],
		1.0, 1.0,

		v1[6], v1[7], v1[8],
		0.0, 1.0,

		v1[9], v1[10], v1[11],
		0.0, 0.0,
	}

	return vertices, indices
}
