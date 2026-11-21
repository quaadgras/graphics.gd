package gd_test

import (
	"testing"

	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector3"

	_ "embed"
)

var basis_test string = `extends Object

func test_basis(angles):
    return Basis.from_euler(angles)
`

func TestBasis(t *testing.T) {
	var runner = Object.New()
	var script = GDScript.New().AsScript()
	script.SetSourceCode(basis_test)
	script.Reload()
	runner.SetScript(script)

	angles := Euler.Radians{X: -1, Y: 0.2, Z: 0}
	engine := Object.Call(runner, "test_basis", angles).(Basis.XYZ)
	if engine != Basis.FromEuler(angles, Angle.OrderYXZ) {
		t.Fatalf("Expected %v, got %v", Object.Call(runner, "test_basis", angles), Basis.FromEuler(angles, Angle.OrderYXZ))
	}
}

var basis_axes_test string = `extends Object

func axes(b: Basis):
    return [b.x, b.y, b.z]

func origin_of(t: Transform3D, p: Vector3):
    return t * p

func made():
    return Basis(Vector3(1, 2, 3), Vector3(4, 5, 6), Vector3(7, 8, 9))

func param_axes(m: ShaderMaterial, name: String):
    return axes(m.get_shader_parameter(name))

func param_point(m: ShaderMaterial, name: String, p: Vector3):
    return m.get_shader_parameter(name) * p
`

// A basis passed to the engine as a variant keeps its axes (the engine
// stores rows, Go columns), and one made by the engine comes back whole.
func TestBasisVariantAxes(t *testing.T) {
	var runner = Object.New()
	var script = GDScript.New().AsScript()
	script.SetSourceCode(basis_axes_test)
	script.Reload()
	runner.SetScript(script)

	b := Basis.XYZ{X: Vector3.XYZ{X: 1, Y: 2, Z: 3}, Y: Vector3.XYZ{X: 4, Y: 5, Z: 6}, Z: Vector3.XYZ{X: 7, Y: 8, Z: 9}}
	axes := Object.Call(runner, "axes", b).([]any)
	if axes[0] != b.X || axes[1] != b.Y || axes[2] != b.Z {
		t.Fatalf("engine saw axes %v, want %v", axes, b)
	}
	if got := Object.Call(runner, "made").(Basis.XYZ); got != b {
		t.Fatalf("engine basis came back as %v, want %v", got, b)
	}
	// Through a variant.Any (every generated method that takes one).
	material := ShaderMaterial.New()
	material.SetShaderParameter("b", b)
	if axes := Object.Call(runner, "param_axes", material, "b").([]any); axes[0] != b.X || axes[1] != b.Y || axes[2] != b.Z {
		t.Fatalf("engine saw the parameter's axes as %v, want %v", axes, b)
	}
	if got := material.GetShaderParameter("b").(Basis.XYZ); got != b {
		t.Fatalf("parameter came back as %v, want %v", got, b)
	}
	yaw := Transform3D.BasisOrigin{Basis: Basis.FromEuler(Euler.Radians{Y: 0.5}, Angle.OrderYXZ), Origin: Vector3.XYZ{X: 10}}
	p := Vector3.XYZ{X: 1, Y: 2, Z: 3}
	got := Object.Call(runner, "origin_of", yaw, p).(Vector3.XYZ)
	want := Transform3D.Transform(p, yaw)
	if Vector3.Length(Vector3.Sub(got, want)) > 1e-5 {
		t.Fatalf("engine transformed %v to %v, want %v", p, got, want)
	}
	material.SetShaderParameter("t", yaw)
	if got := Object.Call(runner, "param_point", material, "t", p).(Vector3.XYZ); Vector3.Length(Vector3.Sub(got, want)) > 1e-5 {
		t.Fatalf("engine transformed %v by the parameter to %v, want %v", p, got, want)
	}
	if got := material.GetShaderParameter("t").(Transform3D.BasisOrigin); got != yaw {
		t.Fatalf("parameter came back as %v, want %v", got, yaw)
	}
}
