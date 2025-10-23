package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Image"
	"graphics.gd/variant/Object"
)

func TestStatic(t *testing.T) {
	t.Parallel()

	var image Image.Instance = Image.Create(1, 1, false, Image.FormatRgb8)
	if image.GetWidth() != 1 {
		t.Fail()
	}
}

type ClassWithStaticMethods struct {
	Object.Extension[ClassWithStaticMethods]
}

func CallStatic() {}

func TestRegisterStaticMethod(t *testing.T) {
	t.Parallel()

	classdb.Register[ClassWithStaticMethods](
		map[string]any{
			"call_static": CallStatic,
		},
	)

	var runner = Object.New()
	var script = GDScript.New().AsScript()
	script.SetSourceCode(`extends Object

func test_static(obj):
    obj.call_static()
    return true
`)
	script.Reload()
	runner.SetScript(script)

	done, ok := Object.Call(runner, "test_static", new(ClassWithStaticMethods)).(bool)
	if !done || !ok {
		t.Fail()
	}
}
