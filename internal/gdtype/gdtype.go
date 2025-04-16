package gdtype

import (
	"fmt"
	"strings"

	"graphics.gd/internal/gdjson"
)

var ClassDB map[string]gdjson.Class

func (s Name) InCore() bool {
	switch s {
	case "Object", "RefCounted":
		return true
	default:
		return false
	}
}

var Builtins = []Name{
	"String", "StringName", "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
	"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedVector4Array", "PackedColorArray", "PackedByteArray",
	"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
	"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID",
	"Callable", "Signal", "Dictionary", "Array",
}

var builtinsMap = make(map[string]struct{})

func init() {
	for _, b := range Builtins {
		builtinsMap[string(b)] = struct{}{}
	}
}

func (s Name) IsBuiltin() bool {
	_, ok := builtinsMap[string(s)]
	return ok
}

type Name string

func (name Name) Let(ctx string, val string) string {
	prefix := ""
	if strings.HasPrefix(string(name), "gd.") {
		prefix = "gd."
	}
	base := strings.TrimPrefix(string(name), "gd.")
	if strings.HasPrefix(base, "ArrayOf") {
		typed := prefix + "TypedArray" + strings.TrimPrefix(base, "ArrayOf")
		return fmt.Sprintf("%s(mmm.Let["+prefix+"Array](%v.Lifetime, %v.API, %v))", typed, ctx, ctx, val)
	}
	return fmt.Sprintf("mmm.Let[%v](%v.Lifetime, %v.API, %v)", name, ctx, ctx, val)
}

func (name Name) Underlying() string {
	prefix := ""
	if strings.HasPrefix(string(name), "gd.") {
		prefix = "gd."
	}
	base := strings.TrimPrefix(string(name), "gd.")
	if strings.HasPrefix(base, "ArrayOf") {
		return prefix + "Array"
	}
	return string(name)
}

func (name Name) ToUnderlying(val string) string {
	if name == "Basis.XYZ" {
		return fmt.Sprintf("Basis.Transposed(%v)", val)
	}
	if name == "Transform3D.BasisOrigin" {
		return fmt.Sprintf("gd.Transposed(%v)", val)
	}
	if name.Underlying() != string(name) {
		return fmt.Sprintf("%v.%v()", val, strings.TrimPrefix(name.Underlying(), "gd."))
	}
	return val
}

func (name Name) ConvertToSimple(val string) string {
	if strings.HasPrefix(val, "(") {
		val = strings.TrimPrefix(val, "(")
		val = strings.TrimSuffix(val, ")")
	}
	if strings.HasPrefix(string(name), "Array.Contains[") {
		return fmt.Sprintf("gd.ArrayFromSlice[%v](%v)", name, val)
	}
	switch name {
	case "Array.Any":
		if val == "Array.Nil" {
			return "Array.Nil"
		}
		return fmt.Sprintf("gd.EngineArrayFromSlice(%v)", val)
	case "Dictionary.Any":
		if val == "Dictionary.Nil" {
			return "Dictionary.Nil"
		}
		return fmt.Sprintf("gd.DictionaryFromMap(%v)", val)
	case "String.Readable":
		return fmt.Sprintf("String.New(%v)", val)
	case "Path.ToNode":
		return fmt.Sprintf("Path.ToNode(String.New(%v))", val)
	case "String.Name":
		return fmt.Sprintf("String.Name(String.New(%v))", val)
	case "int64", "float64", "Vector2.XY", "Vector2i.XY", "Rect2.PositionSize", "Rect2i.PositionSize",
		"Vector3.XYZ", "Vector3i.XYZ", "Transform2D.OriginXY", "Quaternion.IJKL",
		"AABB.PositionSize", "Color.RGBA", "Plane.NormalD", "Basis.XYZ", "Transform3D.BasisOrigin",
		"Vector4.XYZW", "Vector4i.XYZW", "RID.Any":
		return fmt.Sprintf("%s(%v)", name, val)
	case "Packed.Bytes":
		return fmt.Sprintf("Packed.Bytes(Packed.New(%v...))", val)
	case "Packed.Strings":
		return fmt.Sprintf("Packed.MakeStrings(%v...)", val)
	case "Packed.Array[int32]", "Packed.Array[int64]", "Packed.Array[float32]", "Packed.Array[float64]":
		return fmt.Sprintf("Packed.New(%v...)", val)
	case "Packed.Array[Vector2.XY]", "Packed.Array[Vector3.XYZ]", "Packed.Array[Vector4.XYZW]", "Packed.Array[Color.RGBA]":
		elem := strings.TrimPrefix(string(name), "Packed.Array[")
		elem = strings.TrimSuffix(elem, "]")
		if val == "[1][]Vector2.XY{}[0]" {
			return fmt.Sprintf("Packed.New[%s]()", elem)
		}
		return fmt.Sprintf("Packed.New(%v...)", val)
	case "Callable.Function":
		return fmt.Sprintf("Callable.New(%v)", val)
	case "variant.Any":
		return fmt.Sprintf("variant.New(%v)", val)
	case "Error.Code":
		return fmt.Sprintf("Error.New(%v)", val)
	default:
		return val
	}
}

func (name Name) ConvertToGo(val string, simple string) string {
	if strings.HasPrefix(string(name), "Array.Contains[") {
		return fmt.Sprintf("gd.ArrayAs[%s](gd.InternalArray(%s))", simple, val)
	}
	switch name {
	case "String.Readable", "Path.ToNode", "String.Name":
		return fmt.Sprintf("%v.String()", val)
	case "Error.Code":
		return fmt.Sprintf("gd.ToError(%v)", val)
	case "int64":
		return fmt.Sprintf("int(%v)", val)
	case "float64":
		return fmt.Sprintf("Float.X(%v)", val)
	case "Packed.Bytes":
		return fmt.Sprintf("%v.Bytes()", val)
	case "Packed.Strings":
		return fmt.Sprintf("%v.Strings()", val)
	case "Packed.Array[int32]", "Packed.Array[int64]", "Packed.Array[float32]", "Packed.Array[float64]",
		"Packed.Array[Vector2.XY]", "Packed.Array[Vector3.XYZ]", "Packed.Array[Vector4.XYZW]", "Packed.Array[Color.RGBA]":
		return fmt.Sprintf("slices.Collect(%v.Values())", val)
	case "variant.Any":
		return fmt.Sprintf("%v.Interface()", val)
	case "Array.Any":
		return fmt.Sprintf("gd.ArrayAs[%s](gd.InternalArray(%s))", simple, val)
	case "Dictionary.Any":
		return fmt.Sprintf("gd.DictionaryAs[%s](%s)", simple, val)
	default:
		return val
	}
}

func (name Name) LoadFromRawPointerValue(val string) string {
	if strings.HasPrefix(string(name), "Array.Contains[") {
		_, elem, _ := strings.Cut(string(name), "Array.Contains[")
		elem = strings.TrimSuffix(elem, "]")
		return fmt.Sprintf("Array.Through(gd.ArrayProxy[%s]{}, pointers.Pack(pointers.New[gd.Array](%s)))", elem, val)
	}
	switch name {
	case "Signal.Any":
		return fmt.Sprintf("Signal.Via(gd.SignalProxy{}, pointers.Pack(pointers.New[gd.Signal](%s)))", val)
	case "Array.Any":
		return fmt.Sprintf("Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](%s)))", val)
	case "String.Readable":
		return fmt.Sprintf("String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](%s)))", val)
	case "Dictionary.Any":
		return fmt.Sprintf("Dictionary.Through(gd.DictionaryProxy[variant.Any,variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](%s)))", val)
	case "[1]gd.Object":
		return fmt.Sprintf("[1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(%s)})}", val)
	case "Callable.Function":
		return fmt.Sprintf("Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](%s)))", val)
	case "Path.ToNode":
		return fmt.Sprintf("Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](%s))))", val)
	case "String.Name":
		return fmt.Sprintf("String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](%s))))", val)
	case "Packed.Bytes":
		return fmt.Sprintf("Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](%s))))", val)
	case "Packed.Strings":
		return fmt.Sprintf("Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](%s))))", val)
	case "Packed.Array[int32]", "Packed.Array[int64]", "Packed.Array[float32]", "Packed.Array[float64]",
		"Packed.Array[Vector2.XY]", "Packed.Array[Vector3.XYZ]", "Packed.Array[Vector4.XYZW]", "Packed.Array[Color.RGBA]":
		elem := strings.TrimPrefix(string(name), "Packed.Array[")
		elem = strings.TrimSuffix(elem, "]")
		title, _, _ := strings.Cut(elem, ".")
		title = strings.Title(title)
		return fmt.Sprintf("Packed.Array[%s](Array.Through(gd.PackedProxy[gd.Packed%sArray, %s]{}, pointers.Pack(pointers.New[gd.PackedStringArray](%s))))",
			elem, title, elem, val)
	case "variant.Any":
		return fmt.Sprintf("variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](%s)))", val)
	case "Error.Code":
		return fmt.Sprintf("Error.Code(%s)", val)
	case "Basis.XYZ":
		return fmt.Sprintf("Basis.Transposed(%v)", val)
	case "Transform3D.BasisOrigin":
		return fmt.Sprintf("gd.Transposed(%v)", val)
	default:
		_, argIsPtr := name.IsPointer()
		if name == "Object" {
			return fmt.Sprintf("[1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(%v)})}\n", val)
		}
		class_name := strings.TrimPrefix(string(name), "[1]gdclass.")
		_, ok := ClassDB[class_name]
		if ok && argIsPtr {
			return fmt.Sprintf("%s{pointers.New[gdclass.%v]([3]uint64{uint64(%v)})}\n", name, class_name, val)
		}
		if argIsPtr {
			return fmt.Sprintf("pointers.New[%v](%v)", name, val)
		} else {
			return fmt.Sprintf("%v", val)
		}
	}
}

func (name Name) EndPointer(val string) string {
	if strings.HasPrefix(string(name), "Array.Contains[") {
		return fmt.Sprintf("pointers.End(gd.InternalArray(%v))", val)
	}
	switch name {
	case "Signal.Any":
		return fmt.Sprintf("pointers.End(gd.InternalSignal(%v))", val)
	case "Array.Any":
		return fmt.Sprintf("pointers.End(gd.InternalArray(%v))", val)
	case "String.Readable":
		return fmt.Sprintf("pointers.End(gd.InternalString(%v))", val)
	case "Dictionary.Any":
		return fmt.Sprintf("pointers.End(gd.InternalDictionary(%v))", val)
	case "Callable.Function":
		return fmt.Sprintf("pointers.End(gd.InternalCallable(%v))", val)
	case "Path.ToNode":
		return fmt.Sprintf("pointers.End(gd.InternalNodePath(%v))", val)
	case "String.Name":
		return fmt.Sprintf("pointers.End(gd.InternalStringName(%v))", val)
	case "Packed.Bytes":
		return fmt.Sprintf("pointers.End(gd.InternalPacked[gd.PackedByteArray,byte](Packed.Array[byte](%v)))", val)
	case "Packed.Strings":
		return fmt.Sprintf("pointers.End(gd.InternalPackedStrings(%v))", val)
	case "variant.Any":
		return fmt.Sprintf("pointers.End(gd.InternalVariant(%v))", val)
	case "Error.Code":
		return fmt.Sprintf("func(e Error.Code)(int64,bool){return int64(e),true}(%s)", val)
	case "Packed.Array[int32]", "Packed.Array[int64]", "Packed.Array[float32]", "Packed.Array[float64]",
		"Packed.Array[Vector2.XY]", "Packed.Array[Vector3.XYZ]", "Packed.Array[Vector4.XYZW]", "Packed.Array[Color.RGBA]":
		elem := strings.TrimPrefix(string(name), "Packed.Array[")
		elem = strings.TrimSuffix(elem, "]")
		title, _, _ := strings.Cut(elem, ".")
		title = strings.Title(title)
		return fmt.Sprintf("pointers.End(gd.InternalPacked[gd.Packed%sArray,%s](%v))", title, elem, val)
	default:
		name := strings.TrimPrefix(string(name), "classdb.")
		name = strings.TrimPrefix(name, "[1]gdclass.")
		_, ok := ClassDB[strings.TrimPrefix(name, "[1]gdclass.")]
		if ok || name == "[1]gd.Object" {
			return fmt.Sprintf("pointers.End(%v[0])", val)
		}
		return fmt.Sprintf("pointers.End(%v)", val)
	}
}

func (name Name) LoadOntoCallFrame(val string) string {
	if strings.HasPrefix(string(name), "Array.Contains[") {
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalArray(%v)))\n", val)
	}
	switch name {
	case "variant.Any":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalVariant(%v)))\n", val)
	case "Error.Code":
		return fmt.Sprintf("\tcallframe.Arg(frame, int64(%v))\n", val)
	case "Basis.XYZ":
		return fmt.Sprintf("\tcallframe.Arg(frame, Basis.Transposed(%v))\n", val)
	case "Transform3D.BasisOrigin":
		return fmt.Sprintf("\tcallframe.Arg(frame, gd.Transposed(%v))\n", val)
	case "Signal.Any":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalSignal(%v)))\n", val)
	case "Array.Any":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalArray(%v)))\n", val)
	case "String.Readable":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalString(%v)))\n", val)
	case "Dictionary.Any":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalDictionary(%v)))\n", val)
	case "Callable.Function":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalCallable(%v)))\n", val)
	case "Path.ToNode":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalNodePath(%v)))\n", val)
	case "String.Name":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalStringName(%v)))\n", val)
	case "Packed.Bytes":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray,byte](Packed.Array[byte](%v))))\n", val)
	case "Packed.Strings":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(%v)))\n", val)
	case "Packed.Array[int32]", "Packed.Array[int64]", "Packed.Array[float32]", "Packed.Array[float64]",
		"Packed.Array[Vector2.XY]", "Packed.Array[Vector3.XYZ]", "Packed.Array[Vector4.XYZW]", "Packed.Array[Color.RGBA]":
		elem := strings.TrimPrefix(string(name), "Packed.Array[")
		elem = strings.TrimSuffix(elem, "]")
		title, _, _ := strings.Cut(elem, ".")
		title = strings.Title(title)
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.Packed%sArray, %s](%v)))\n", title, elem, val)
	}
	_, argIsPtr := name.IsPointer()
	if argIsPtr {
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(%v))\n", val)
	} else {
		return fmt.Sprintf("\tcallframe.Arg(frame, %v)\n", val)
	}
}

func (name Name) IsPointer() (string, bool) {
	t := string(name)
	t = strings.TrimPrefix(t, "[1]")
	t = strings.TrimPrefix(t, "gd.")
	t = strings.TrimPrefix(t, "gdclass.")
	t = strings.TrimPrefix(t, "[1]gdclass.")
	if strings.HasPrefix(t, "Array.Contains[") {
		return "[1]gd.EnginePointer", true
	}
	switch t {
	case "String", "StringName", "NodePath",
		"Dictionary.Any", "Array.Any", "String.Readable", "Path.ToNode", "String.Name":
		return "[1]gd.EnginePointer", true
	case "Signal", "Signal.Any":
		return "[2]uint64", true
	case "Callable":
		return "[2]uint64", true
	case "Callable.Function":
		return "[2]uint64", true
	case "Packed.Bytes", "Packed.Array[int32]",
		"Packed.Array[int64]", "Packed.Array[float32]",
		"Packed.Array[float64]", "Packed.Strings",
		"Packed.Array[Vector2.XY]", "Packed.Array[Vector3.XYZ]",
		"Packed.Array[Vector4.XYZW]",
		"Packed.Array[Color.RGBA]":
		return "gd.PackedPointers", true
	case "Variant", "variant.Any":
		return "[3]uint64", true
	case "Object":
		return "gd.EnginePointer", true
	case "Error.Code":
		return "int64", true
	default:
		if entry, ok := ClassDB[t]; ok && !entry.IsEnum {
			return "gd.EnginePointer", true
		}
		return "", false
	}
}
