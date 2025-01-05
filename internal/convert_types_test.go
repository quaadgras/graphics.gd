package gd_test

import (
	"errors"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/variant"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Dictionary"
	"graphics.gd/variant/NodePath"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"
	"graphics.gd/variant/StringName"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"

	_ "embed"
)

//go:embed convert_types_test.gd
var convert_types_test string

var doneConversionsTest = make(chan error, 2)

type Converter struct {
	classdb.Extension[Converter, Node.Instance]
}

func (c Converter) Ready()                { doneConversionsTest <- errors.New("script didn't call Done") }
func (c Converter) AsNode() Node.Instance { return c.Super() }

func (c Converter) Done() { doneConversionsTest <- nil }
func (c Converter) Fail() { doneConversionsTest <- errors.New("test failed") }

func (c Converter) HelloWorld() string     { return "Hello, World!" }
func (c Converter) Echo(s string) string   { return s }
func (c Converter) Int() int               { return 22 }
func (c Converter) Bool() bool             { return true }
func (c Converter) Int8() int8             { return 8 }
func (c Converter) Int16() int16           { return 16 }
func (c Converter) Int32() int32           { return 32 }
func (c Converter) Int64() int64           { return 64 }
func (c Converter) Uint() uint             { return 22 }
func (c Converter) Uint8() uint8           { return 8 }
func (c Converter) Uint16() uint16         { return 16 }
func (c Converter) Uint32() uint32         { return 32 }
func (c Converter) Uint64() uint64         { return 64 }
func (c Converter) Float32() float32       { return 2.0 }
func (c Converter) Float64() float64       { return 2.2 }
func (c Converter) Complex64() complex64   { return 2 + 2i }
func (c Converter) Complex128() complex128 { return 2.2 + 2.2i }

func (c Converter) Float() gd.Float                      { return 2.2 }
func (c Converter) String() String.Advanced              { return String.New("testing") }
func (c Converter) Vector2() Vector2.XY                  { return Vector2.New(1, 2) }
func (c Converter) Vector2i() Vector2i.XY                { return Vector2i.New(1, 2) }
func (c Converter) Rect2() Rect2.PositionSize            { return Rect2.New(1, 2, 3, 4) }
func (c Converter) Rect2i() Rect2i.PositionSize          { return Rect2i.New(1, 2, 3, 4) }
func (c Converter) Vector3() Vector3.XYZ                 { return Vector3.New(1, 2, 3) }
func (c Converter) Vector3i() Vector3i.XYZ               { return Vector3i.New(1, 2, 3) }
func (c Converter) GetTransform2D() Transform2D.OriginXY { return Transform2D.Identity }
func (c Converter) Vector4() Vector4.XYZW                { return Vector4.New(1, 2, 3, 4) }
func (c Converter) Vector4i() Vector4i.XYZW              { return Vector4i.New(1, 2, 3, 4) }
func (c Converter) Plane() Plane.NormalD                 { return Plane.New(1, 2, 3, 4) }
func (c Converter) GetQuaternion() Quaternion.IJKX       { return Quaternion.New() }
func (c Converter) AABB() AABB.PositionSize {
	return AABB.PositionSize{Position: Vector3.New(1, 2, 3), Size: Vector3.New(4, 5, 6)}
}
func (c Converter) GetBasis() Basis.XYZ                     { return Basis.New() }
func (c Converter) GetTransform3D() Transform3D.BasisOrigin { return Transform3D.Identity }
func (c Converter) GetProjection() Projection.XYZW          { return Projection.New() }
func (c Converter) Color() Color.RGBA                       { return Color.Bytes(255, 0, 0, 255) }
func (c Converter) StringName() StringName.Advanced         { return StringName.New("testing") }
func (c Converter) NodePath() NodePath.String               { return "/" }
func (c Converter) RID() Resource.ID                        { return 22 }
func (c Converter) Object() Object.Instance                 { return Object.New() }
func (c Converter) Callable() Callable.Any {
	return Callable.New(func() string {
		return "Hello, World!"
	})
}
func (c Converter) Signal() Signal.Any {
	return gd.NewSignalOf(c.AsObject(), StringName.New("property_list_changed"))
}
func (c Converter) Dictionary() Dictionary.Any {
	return variant.New(map[string]string{
		"hello": "world",
	}).Interface().(Dictionary.Any)
}
func (c Converter) Array() []int                  { return []int{1, 2, 3} }
func (c Converter) PackedByteArray() []byte       { return []byte{1, 2, 3} }
func (c Converter) PackedInt32Array() []int32     { return []int32{1, 2, 3} }
func (c Converter) PackedInt64Array() []int64     { return []int64{1, 2, 3} }
func (c Converter) PackedFloat32Array() []float32 { return []float32{1, 2, 3} }
func (c Converter) PackedFloat64Array() []float64 { return []float64{1, 2, 3} }
func (c Converter) PackedStringArray() []string   { return []string{"hello", "world"} }
func (c Converter) PackedVector2Array() []Vector2.XY {
	return []Vector2.XY{{1, 2}, {3, 4}}
}
func (c Converter) PackedVector3Array() []Vector3.XYZ {
	return []Vector3.XYZ{{1, 2, 3}, {4, 5, 6}}
}
func (c Converter) PackedColorArray() []Color.RGBA {
	return []Color.RGBA{{1, 2, 3, 4}, {5, 6, 7, 8}}
}
func (c Converter) PackedVector4Array() []Vector4.XYZW {
	return []Vector4.XYZW{{1, 2, 3, 4}, {5, 6, 7, 8}}
}

func TestConversions(t *testing.T) {
	//fmt.Println(variant.New(Resource.ID(22)).Interface())

	converter := &Converter{}
	var script = GDScript.New().AsScript()
	script.SetSourceCode(convert_types_test)
	script.Reload()
	Object.Instance{converter.AsObject()}.SetScript(script)
	SceneTree.Add(converter)
	if err := <-doneConversionsTest; err != nil {
		t.Fatal(err)
	}
}
