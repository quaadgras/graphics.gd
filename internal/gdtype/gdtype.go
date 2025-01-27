package gdtype

import (
	"fmt"
	"strings"

	"graphics.gd/internal/gdjson"
)

var ClassDB map[string]gdjson.Class

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
	if name.Underlying() != string(name) {
		return fmt.Sprintf("%v.%v()", val, strings.TrimPrefix(name.Underlying(), "gd."))
	}
	return val
}

func (name Name) End(val string) string {
	if name.Underlying() != string(name) {
		return fmt.Sprintf("if %s != nil { mmm.End(%s.%v()) }", val, val, strings.TrimPrefix(name.Underlying(), "gd."))
	}
	return "mmm.End(" + val + ")"
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
	case "gd.RID":
		return fmt.Sprintf("gd.RID(%v)", val)
	case "gd.String":
		return fmt.Sprintf("gd.NewString(%v)", val)
	case "gd.StringName":
		return fmt.Sprintf("gd.NewStringName(%v)", val)
	case "gd.NodePath":
		return fmt.Sprintf("gd.NewString(string(%v)).NodePath()", val)
	case "gd.Int", "gd.Float", "gd.Vector2", "gd.Vector2i", "gd.Rect2", "gd.Rect2i",
		"gd.Vector3", "gd.Vector3i", "gd.Transform2D", "gd.Quaternion",
		"gd.AABB", "gd.Color", "gd.Plane", "gd.Basis", "gd.Transform3D",
		"gd.Vector4", "gd.Vector4i":
		return fmt.Sprintf("%s(%v)", name, val)
	case "gd.Dictionary":
		return fmt.Sprintf("gd.NewVariant(%v).Interface().(gd.Dictionary)", val)
	case "gd.PackedByteArray":
		return fmt.Sprintf("gd.NewPackedByteSlice(%v)", val)
	case "gd.PackedStringArray":
		return fmt.Sprintf("gd.NewPackedStringSlice(%v)", val)
	case "gd.PackedInt32Array":
		return fmt.Sprintf("gd.NewPackedInt32Slice(%v)", val)
	case "gd.PackedInt64Array":
		return fmt.Sprintf("gd.NewPackedInt64Slice(%v)", val)
	case "gd.PackedFloat32Array":
		return fmt.Sprintf("gd.NewPackedFloat32Slice(%v)", val)
	case "gd.PackedFloat64Array":
		return fmt.Sprintf("gd.NewPackedFloat64Slice(%v)", val)
	case "gd.PackedVector2Array":
		if val == "[1][]Vector2.XY{}[0]" {
			return "gd.NewPackedVector2Slice(nil)"
		}
		return fmt.Sprintf("gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&%v)))", val)
	case "gd.PackedVector3Array":
		if val == "[1][]Vector3.XYZ{}[0]" {
			return "gd.NewPackedVector3Slice(nil)"
		}
		return fmt.Sprintf("gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&%v)))", val)
	case "gd.PackedVector4Array":
		if val == "[1][]Vector4.XYZW{}[0]" {
			return "gd.NewPackedVector4Slice(nil)"
		}
		return fmt.Sprintf("gd.NewPackedVector4Slice(*(*[]gd.Vector4)(unsafe.Pointer(&%v)))", val)
	case "gd.PackedColorArray":
		if val == "[1][]Color.RGBA{}[0]" {
			return "gd.NewPackedColorSlice(nil)"
		}
		return fmt.Sprintf("gd.NewPackedColorSlice(*(*[]gd.Color)(unsafe.Pointer(&%v)))", val)
	case "gd.Variant":
		return fmt.Sprintf("gd.NewVariant(%v)", val)
	case "Callable.Function":
		return fmt.Sprintf("Callable.New(%v)", val)
	default:
		return val
	}
}

func (name Name) ConvertToGo(val string, simple string) string {
	if strings.HasPrefix(string(name), "Array.Contains[") {
		return fmt.Sprintf("gd.ArrayAs[%s](gd.InternalArray(%s))", simple, val)
	}
	switch name {
	case "gd.String", "gd.StringName", "gd.NodePath", "String.Readable":
		return fmt.Sprintf("%v.String()", val)
	case "gd.Error":
		return fmt.Sprintf("gd.ToError(%v)", val)
	case "gd.Int":
		return fmt.Sprintf("int(%v)", val)
	case "gd.Float":
		return fmt.Sprintf("Float.X(%v)", val)
	case "gd.PackedByteArray":
		return fmt.Sprintf("%v.Bytes()", val)
	case "gd.PackedStringArray":
		return fmt.Sprintf("%v.Strings()", val)
	case "gd.PackedInt32Array", "gd.PackedInt64Array", "gd.PackedFloat32Array", "gd.PackedFloat64Array",
		"gd.PackedVector2Array", "gd.PackedVector3Array", "gd.PackedVector4Array", "gd.PackedColorArray":
		return fmt.Sprintf("%v.AsSlice()", val)
	case "gd.Variant":
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
			return fmt.Sprintf("%v\n", val)
		}
	}
}

func (name Name) EndPointer(val string) string {
	if strings.HasPrefix(string(name), "Array.Contains[") {
		return fmt.Sprintf("pointers.End(gd.InternalArray(%v))", val)
	}
	switch name {
	case "Array.Any":
		return fmt.Sprintf("pointers.End(gd.InternalArray(%v))", val)
	case "String.Readable":
		return fmt.Sprintf("pointers.End(gd.InternalString(%v))", val)
	case "Dictionary.Any":
		return fmt.Sprintf("pointers.End(gd.InternalDictionary(%v))", val)
	case "Callable.Function":
		return fmt.Sprintf("pointers.End(gd.InternalCallable(%v))", val)
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
	case "Array.Any":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalArray(%v)))\n", val)
	case "String.Readable":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalString(%v)))\n", val)
	case "Dictionary.Any":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalDictionary(%v)))\n", val)
	case "Callable.Function":
		return fmt.Sprintf("\tcallframe.Arg(frame, pointers.Get(gd.InternalCallable(%v)))\n", val)
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
		"Dictionary.Any", "Array.Any", "String.Readable":
		return "[1]gd.EnginePointer", true
	case "Signal":
		return "[2]uint64", true
	case "Callable":
		return "[2]uint64", true
	case "Callable.Function":
		return "[2]uint64", true
	case "PackedByteArray", "PackedInt32Array",
		"PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedStringArray",
		"PackedVector2Array", "PackedVector3Array",
		"PackedVector4Array",
		"PackedColorArray":
		return "gd.PackedPointers", true
	case "Variant":
		return "[3]uint64", true
	case "Object":
		return "gd.EnginePointer", true
	default:
		if entry, ok := ClassDB[t]; ok && !entry.IsEnum {
			return "gd.EnginePointer", true
		}
		return "", false
	}
}
