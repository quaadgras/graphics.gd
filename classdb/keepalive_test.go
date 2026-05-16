package classdb

import (
	"reflect"
	"testing"

	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
)

// TestKeepAlive — existing baseline: unexported Instance inside a nested
// non-extension struct should be picked up.
func TestKeepAlive(t *testing.T) {
	type Child struct {
		node Node.Instance
	}
	type TestStruct struct {
		child Child
	}
	fn := compile_keepalive(reflect.TypeFor[TestStruct]())
	if fn == nil {
		t.Fatal("not detected")
	}
	fn(reflect.ValueOf(&TestStruct{}).Elem())
}

// TestKeepAlive_ExtensionUnexportedInstance — matches CitizenEditor.scene:
// an unexported Node3D.Instance field directly on an extension class.
//
// The walker has an early `continue` for fields that implement Node.Any and
// are exported on extension classes (auto-bound scene nodes). The condition
// requires IsExported() == true, so unexported fields should NOT be skipped
// and must be walked.
func TestKeepAlive_ExtensionUnexportedInstance(t *testing.T) {
	type CitizenLike struct {
		Node3D.Extension[CitizenLike]
		scene Node3D.Instance
	}
	fn := compile_keepalive(reflect.TypeFor[CitizenLike]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil for *CitizenLike — the unexported Node3D.Instance field on an extension class was not picked up as a keepalive root")
	}
	fn(reflect.ValueOf(&CitizenLike{}).Elem())
}

// TestKeepAlive_PointerToStructWithInstance — matches CitizenEditor.body
// (*CitizenBody) where CitizenBody contains an unexported Instance field.
// The walker should: pointer-deref → struct fields → Instance match.
func TestKeepAlive_PointerToStructWithInstance(t *testing.T) {
	type Inner struct {
		mesh MeshInstance3D.Instance
	}
	type Outer struct {
		Node3D.Extension[Outer]
		body *Inner
	}
	fn := compile_keepalive(reflect.TypeFor[Outer]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil — the *Inner.mesh chain wasn't walked")
	}
	fn(reflect.ValueOf(&Outer{}).Elem())
}

// TestKeepAlive_FullCitizenShape — the exact nesting the citizen editor
// uses: extension class with both unexported Instance and unexported
// pointer-to-struct-with-Instance fields.
func TestKeepAlive_FullCitizenShape(t *testing.T) {
	type Body struct {
		mesh MeshInstance3D.Instance
	}
	type Editor struct {
		Node3D.Extension[Editor]
		scene Node3D.Instance
		body  *Body
	}
	fn := compile_keepalive(reflect.TypeFor[Editor]())
	if fn == nil {
		t.Fatal("compile_keepalive returned nil for the CitizenEditor-shaped struct")
	}
	// Smoke-call the closure on a zero value — must not panic on nil
	// pointers (body is nil here).
	val := reflect.New(reflect.TypeFor[Editor]()).Elem()
	fn(val)
}
