package graphics143

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

const (
	textureVertexShaderSrc = `
	#version 410

	layout (location = 0) in vec3 position;
	layout (location = 1) in vec2 texCoord;
	
	out vec2 TexCoord;
	
	void main()
	{
			gl_Position = vec4(position, 1.0);
			TexCoord = texCoord;    // pass the texture coords on to the fragment shader
	}
		`

	textureFragmentShaderSrc = `
		#version 410 core
		in vec2 TexCoord;

		out vec4 color;

		uniform sampler2D ourTexture0;

		void main()
		{
			// mix the two textures together (texture1 is colored with "ourColor")
			color = texture(ourTexture0, TexCoord);
		}
	`
)

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
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

func makeProgram(vertexShaderSource, fragmentShaderSource string) (uint32, uint32, uint32) {
	prog := gl.CreateProgram()

	vertShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	gl.AttachShader(prog, vertShader)

	fragShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	gl.AttachShader(prog, fragShader)

	gl.LinkProgram(prog)
	return prog, vertShader, fragShader
}

func getUniformLocation(openglProgram uint32, name string) int32 {
	return gl.GetUniformLocation(openglProgram, gl.Str(name+"\x00"))
}
