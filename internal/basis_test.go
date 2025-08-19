package gd_test

import (
	"testing"

	"graphics.gd/classdb/GDScript"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Object"

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
