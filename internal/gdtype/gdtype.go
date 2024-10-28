package gdtype

import (
	"fmt"
	"strings"
)

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
	switch name {
	case "gd.String":
		return fmt.Sprintf("gd.NewString(%v)", val)
	case "gd.StringName":
		return fmt.Sprintf("gd.NewStringName(%v)", val)
	case "gd.NodePath":
		return fmt.Sprintf("gd.NewString(%v).NodePath()", val)
	case "gd.Int":
		return fmt.Sprintf("gd.Int(%v)", val)
	case "gd.Float":
		return fmt.Sprintf("gd.Float(%v)", val)
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
		return fmt.Sprintf("gd.NewPackedVector2Slice(%v)", val)
	case "gd.PackedVector3Array":
		return fmt.Sprintf("gd.NewPackedVector3Slice(%v)", val)
	case "gd.PackedVector4Array":
		return fmt.Sprintf("gd.NewPackedVector4Slice(%v)", val)
	case "gd.PackedColorArray":
		return fmt.Sprintf("gd.NewPackedColorSlice(%v)", val)
	default:
		return val
	}
}

func (name Name) ConvertToGo(val string) string {
	switch name {
	case "gd.String", "gd.StringName", "gd.NodePath":
		return fmt.Sprintf("%v.String()", val)
	case "gd.Int":
		return fmt.Sprintf("int(%v)", val)
	case "gd.Float":
		return fmt.Sprintf("float64(%v)", val)
	case "gd.PackedByteArray":
		return fmt.Sprintf("%v.Bytes()", val)
	case "gd.PackedStringArray":
		return fmt.Sprintf("%v.Strings()", val)
	case "gd.PackedInt32Array", "gd.PackedInt64Array", "gd.PackedFloat32Array", "gd.PackedFloat64Array",
		"gd.PackedVector2Array", "gd.PackedVector3Array", "gd.PackedVector4Array", "gd.PackedColorArray":
		return fmt.Sprintf("%v.AsSlice()", val)
	default:
		return val
	}
}
