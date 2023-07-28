package basics

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	BasicVertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	TextureVertexShaderSrc = `
	#version 410

	layout (location = 0) in vec3 position;
	layout (location = 1) in vec2 texCoord;
	
	out vec2 TexCoord;
	
	void main()
	{
			gl_Position = vec4(position, 1.0);
			TexCoord = texCoord;    // pass the texture coords on to the fragment shader
	}
		` + "\x00"

	TextureFragmentShaderSrc = `
		#version 410 core
		in vec2 TexCoord;

		out vec4 color;

		uniform sampler2D ourTexture0;

		void main()
		{
			// mix the two textures together (texture1 is colored with "ourColor")
			color = texture(ourTexture0, TexCoord);
		}
	` + "\x00"
)

func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func MakeProgram(vertexShaderSource, fragmentShaderSource string) uint32 {
	prog := gl.CreateProgram()

	vertShader, err := CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	gl.AttachShader(prog, vertShader)

	fragShader, err := CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	gl.AttachShader(prog, fragShader)

	gl.LinkProgram(prog)
	return prog
}

func ConvertColorToShaderFloats(hexColor string) (float32, float32, float32, float32) {
	c, err := colorful.Hex(hexColor)
	if err != nil {
		panic(err)
	}

	r, g, b, a := c.RGBA()
	rNormalized := float32(r) / 65535.0
	gNormalized := float32(g) / 65535.0
	bNormalized := float32(b) / 65535.0
	aNormalized := float32(a) / 65535.0

	return rNormalized, gNormalized, bNormalized, aNormalized
}

func GetRectColorShader(hexColor string) (string, error) {
	rf, gf, bf, af := ConvertColorToShaderFloats(hexColor)
	fragmentShaderSource := fmt.Sprintf(`
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(%f, %f, %f, %f);
		}
	`, rf, gf, bf, af)

	return fragmentShaderSource + "\x00", nil
}

func GetPointShader(hexColor string) (string, error) {
	rf, gf, bf, af := ConvertColorToShaderFloats(hexColor)
	circlePointFragmentSource := fmt.Sprintf(`
	#version 410
	out vec4 frag_colour;
	void main() {
		frag_colour = vec4(%f, %f, %f, %f);

		vec2 coord = gl_PointCoord - vec2(0.5); 
		if(length(coord) > 0.5)
			discard;
	}
	`, rf, gf, bf, af)

	return circlePointFragmentSource + "\x00", nil
}

func GetUniformLocation(openglProgram uint32, name string) int32 {
	return gl.GetUniformLocation(openglProgram, gl.Str(name+"\x00"))
}
