// These tests only generate code, so they run without the engine:
//
//	CGO_ENABLED=0 go test ./shaders
package shaders_test

import (
	"strings"
	"testing"

	"graphics.gd/shaders"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/int"
	"graphics.gd/shaders/over"
	"graphics.gd/shaders/pipeline/Spatial"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/vec3"
)

type Branchy struct {
	Spatial.Shader[Branchy]

	Threshold float.X
	Red       vec3.RGB
	Blue      vec3.RGB
}

func (s *Branchy) Material(frag Spatial.Fragment) Spatial.Material {
	light := float.Mul(frag.UV.X, 3.0)
	color := s.Red
	if float.Gt(frag.UV.Y, s.Threshold).Evaluate() {
		color = rgb.Mul(s.Blue, float.Sin(light))
	}
	return Spatial.Material{
		Albedo: rgb.Mul(rgb.Mul(color, light), 2.0),
	}
}

// expect fails the test unless the code generated for prog contains the
// function body want.
func expect(t *testing.T, prog shaders.Any, want string) {
	t.Helper()
	code := shaders.Generate(prog)
	if !strings.Contains(code, strings.TrimSpace(want)) {
		t.Fatalf("generated:\n%s\nwant:\n%s", code, want)
	}
}

// Only the value that differs between the branches is computed inside of the
// if, the work done on it afterwards is not repeated in each branch.
func TestBranch(t *testing.T) {
	expect(t, new(Branchy), `
void fragment() {
	float _v1 = (UV.x * 3.0);
	vec3 _r0;
	if ((UV.y > threshold)) {
		_r0 = (blue * sin(_v1));
	} else {
		_r0 = red;
	}
	ALBEDO = ((_r0 * _v1) * 2.0);
}`)
}

type Sequential struct {
	Spatial.Shader[Sequential]
}

func (s *Sequential) Material(frag Spatial.Fragment) Spatial.Material {
	a, b := float.New(0.0), float.New(0.0)
	if float.Gt(frag.UV.X, 0.5).Evaluate() {
		a = float.Sin(frag.UV.X)
	}
	if float.Gt(frag.UV.Y, 0.5).Evaluate() {
		b = float.Cos(frag.UV.Y)
	} else {
		b = float.Mul(a, 2.0)
	}
	return Spatial.Material{Roughness: a, Metallic: b}
}

// Two conditions evaluated one after the other become two ifs, rather than
// the second if being repeated inside of each branch of the first.
func TestSequential(t *testing.T) {
	expect(t, new(Sequential), `
void fragment() {
	float _r1;
	if ((UV.x > 0.5)) {
		_r1 = sin(UV.x);
	} else {
		_r1 = 0.0;
	}
	float _r0;
	if ((UV.y > 0.5)) {
		_r0 = cos(UV.y);
	} else {
		_r0 = (_r1 * 2.0);
	}
	METALLIC = _r0;
	ROUGHNESS = _r1;
}`)
}

type Looping struct {
	Spatial.Shader[Looping]

	Steps float.X
	Sun   vec3.RGB
}

func (s *Looping) Material(frag Spatial.Fragment) Spatial.Material {
	step := vec3.Mul(frag.Normal, 0.1)
	sun := rgb.Mul(s.Sun, float.Exp(s.Steps)) // computed once, outside of the loop.
	position, alpha, light := frag.VertexPosition, float.New(0.0), rgb.New(0.0, 0.0, 0.0)
	for i := range over.Int(8) {
		position = vec3.Add(position, step)
		t := float.Sin(float.Add(position.X, float.FromInt(i)))
		if float.Gt(t, 0.0).Evaluate() {
			alpha = float.Add(alpha, float.Mul(t, float.Sub(1.0, alpha)))
			light = rgb.Add(light, rgb.Mul(sun, t))
		}
	}
	return Spatial.Material{
		Albedo: light,
		Alpha:  alpha,
	}
}

// Values computed outside of the loop body are computed once, before the loop,
// the variables that the body changes become the variables of the loop, and a
// condition evaluated inside of the body becomes an if inside of the loop.
func TestLoop(t *testing.T) {
	expect(t, new(Looping), `
void fragment() {
	vec3 _v6 = (NORMAL * 0.1);
	vec3 _v9 = (sun * exp(steps));
	vec3 _s0 = vec3(0.0, 0.0, 0.0);
	float _s1 = 0.0;
	vec3 _s2 = VERTEX;
	for (int _i3 = 0; _i3 < 8; _i3++) {
		vec3 _v7 = (_s2 + _v6);
		float _v8 = sin((_v7.x + float(_i3)));
		vec3 _r4;
		float _r5;
		if ((_v8 > 0.0)) {
			_r4 = (_s0 + (_v9 * _v8));
			_r5 = (_s1 + (_v8 * (1.0 - _s1)));
		} else {
			_r4 = _s0;
			_r5 = _s1;
		}
		_s0 = _r4;
		_s1 = _r5;
		_s2 = _v7;
	}
	ALBEDO = _s0;
	ALPHA = _s1;
}`)
}

type Breaking struct {
	Spatial.Shader[Breaking]
}

func (s *Breaking) Material(frag Spatial.Fragment) Spatial.Material {
	position := frag.VertexPosition
	for range over.Int(int.FromFloat(frag.UV.X)) {
		position = vec3.Add(position, frag.Normal)
		if float.Gt(vec3.Length(position), 10.0).Evaluate() {
			break
		}
	}
	return Spatial.Material{Albedo: rgb.New(position.X, position.Y, position.Z)}
}

// A break inside of a branch stops the loop after the variables are assigned,
// and the components of a variable read after the loop are read from it.
func TestBreak(t *testing.T) {
	expect(t, new(Breaking), `
void fragment() {
	vec3 _s0 = VERTEX;
	for (int _i1 = 0; _i1 < int(UV.x); _i1++) {
		vec3 _v2 = (_s0 + NORMAL);
		_s0 = _v2;
		if ((length(_v2) > 10.0)) {
			break;
		}
	}
	ALBEDO = vec3(_s0.x, _s0.y, _s0.z);
}`)
}

type Repeated struct {
	Spatial.Shader[Repeated]
}

func (s *Repeated) Material(frag Spatial.Fragment) Spatial.Material {
	a, b := float.New(0.0), float.New(0.0)
	if float.Gt(frag.UV.X, 0.5).Evaluate() {
		a = float.Sin(frag.UV.X)
	}
	if float.Gt(frag.UV.X, 0.5).Evaluate() { // the same condition goes the same way.
		b = float.Cos(frag.UV.X)
	}
	return Spatial.Material{Roughness: a, Metallic: b}
}

func TestRepeatedCondition(t *testing.T) {
	expect(t, new(Repeated), `
void fragment() {
	float _r0;
	float _r1;
	if ((UV.x > 0.5)) {
		_r0 = cos(UV.x);
		_r1 = sin(UV.x);
	} else {
		_r0 = 0.0;
		_r1 = 0.0;
	}
	METALLIC = _r0;
	ROUGHNESS = _r1;
}`)
}

type Copying struct {
	Spatial.Shader[Copying]
}

func (s *Copying) Material(frag Spatial.Fragment) Spatial.Material {
	prev, cur := frag.UV.X, float.New(0.0)
	a, b := float.New(1.0), float.New(2.0)
	total := float.New(0.0)
	for range over.Int(4) {
		prev = cur
		cur = float.Add(cur, 1.0)
		total = float.Add(total, b)
		b = a // only read inside of the loop
		a = float.Mul(a, 2.0)
	}
	return Spatial.Material{Roughness: prev, Metallic: total}
}

type Found struct {
	Spatial.Shader[Found]
}

func (s *Found) Material(frag Spatial.Fragment) Spatial.Material {
	x := frag.UV.X
	hit := float.New(0.0)
	for i := range over.Int(16) {
		x = float.Mul(x, 1.5)
		if float.Gt(x, 1.0).Evaluate() {
			hit = float.FromInt(i)
			break
		}
	}
	return Spatial.Material{Roughness: hit}
}

// A variable set to the value of another, from before that other one changes,
// is set to the other's variable of the loop, whether or not it is used after
// the loop.
func TestLoopCopy(t *testing.T) {
	expect(t, new(Copying), `
void fragment() {
	float _s0 = 0.0;
	float _s1 = UV.x;
	float _s2 = 0.0;
	float _s3 = 2.0;
	float _s4 = 1.0;
	for (int _i5 = 0; _i5 < 4; _i5++) {
		_s0 = (_s0 + _s3);
		_s1 = _s2;
		_s2 = (_s2 + 1.0);
		_s3 = _s4;
		_s4 = (_s4 * 2.0);
	}
	METALLIC = _s0;
	ROUGHNESS = _s1;
}`)
}

// A variable that is only set right before a break is a variable of the loop.
func TestLoopFound(t *testing.T) {
	expect(t, new(Found), `
void fragment() {
	float _s0 = 0.0;
	float _s1 = UV.x;
	for (int _i2 = 0; _i2 < 16; _i2++) {
		float _v5 = (_s1 * 1.5);
		bool _v6 = (_v5 > 1.0);
		float _r3;
		float _r4;
		if (_v6) {
			_r3 = float(_i2);
			_r4 = _s1;
		} else {
			_r3 = _s0;
			_r4 = _v5;
		}
		_s0 = _r3;
		_s1 = _r4;
		if (_v6) {
			break;
		}
	}
	ROUGHNESS = _s0;
}`)
}

type Skipping struct {
	Spatial.Shader[Skipping]
}

func (s *Skipping) Material(frag Spatial.Fragment) Spatial.Material {
	sum, count := float.New(0.0), float.New(0.0)
	for i := range over.Int(8) {
		x := float.Sin(float.Add(frag.UV.X, float.FromInt(i)))
		if float.Lt(x, 0.0).Evaluate() {
			continue
		}
		sum = float.Add(sum, x)
		count = float.Add(count, 1.0)
	}
	return Spatial.Material{Roughness: float.Div(sum, count)}
}

// A continue inside of a branch skips the rest of the body for that iteration.
func TestContinue(t *testing.T) {
	expect(t, new(Skipping), `
void fragment() {
	float _s0 = 0.0;
	float _s1 = 0.0;
	for (int _i2 = 0; _i2 < 8; _i2++) {
		float _v5 = sin((UV.x + float(_i2)));
		float _r3;
		float _r4;
		if ((_v5 < 0.0)) {
			_r3 = _s0;
			_r4 = _s1;
		} else {
			_r3 = (_s0 + _v5);
			_r4 = (_s1 + 1.0);
		}
		_s0 = _r3;
		_s1 = _r4;
	}
	ROUGHNESS = (_s0 / _s1);
}`)
}
