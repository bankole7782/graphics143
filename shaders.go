package graphics143

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	BasicVertexShaderSource = `
	#version 460
	in vec3 vp;
	void main() {
		gl_Position = vec4(vp, 1.0);
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
		#version 460
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
	#version 460
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
