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
	`

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
		`

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
	`
)

func CompileShader(source string, shaderType uint32) (uint32, error) {
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

func MakeProgram(vertexShaderSource, fragmentShaderSource string) (uint32, uint32, uint32) {
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
	return prog, vertShader, fragShader
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

	return fragmentShaderSource, nil
}

func GetRectGradientShader(hexColor1, hexColor2 string, directionX bool, windowWidth, windowHeight int, rectSpecs RectSpecs) (string, error) {
	rf1, gf1, bf1, af1 := ConvertColorToShaderFloats(hexColor1)
	rf2, gf2, bf2, af2 := ConvertColorToShaderFloats(hexColor2)

	var fragmentShaderSource string
	if directionX {
		fragmentShaderSource = fmt.Sprintf(`
			#version 410 core
			out vec4 frag_colour;

			void main() {
				vec2 u_resolution = vec2(%d, %d);
				vec2 st = gl_FragCoord.xy/u_resolution.xy;
				float mixValue = distance(st, vec2(0, 1));
				vec4 color1 = vec4(%f, %f, %f, %f);
				vec4 color2 = vec4(%f, %f, %f, %f);
				frag_colour = mix(color1, color2, mixValue);
			}
			`, windowWidth, windowHeight,
			rf1, gf1, bf1, af1,
			rf2, gf2, bf2, af2,
		)
	} else {
		fragmentShaderSource = fmt.Sprintf(`
			#version 410 core
			out vec4 frag_colour;

			void main() {
				float lerpValue = gl_FragCoord.y / %d.0f;
				vec4 color1 = vec4(%f, %f, %f, %f);
				vec4 color2 = vec4(%f, %f, %f, %f);
				frag_colour = mix(color1, color2, lerpValue);
			}
			`, (rectSpecs.Height - rectSpecs.OriginY),
			rf1, gf1, bf1, af1,
			rf2, gf2, bf2, af2,
		)

	}

	fmt.Println(fragmentShaderSource)

	return fragmentShaderSource, nil
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

	return circlePointFragmentSource, nil
}

func GetUniformLocation(openglProgram uint32, name string) int32 {
	return gl.GetUniformLocation(openglProgram, gl.Str(name+"\x00"))
}
