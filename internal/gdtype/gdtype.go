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
		return fmt.Sprintf("gc.String(%v)", val)
	case "gd.StringName":
		return fmt.Sprintf("gc.StringName(%v)", val)
	case "gd.Int":
		return fmt.Sprintf("gd.Int(%v)", val)
	case "gd.Float":
		return fmt.Sprintf("gd.Float(%v)", val)
	case "gd.PackedByteArray":
		return fmt.Sprintf("gc.PackedByteSlice(%v)", val)
	default:
		return val
	}
}

func (name Name) ConvertToGo(val string) string {
	switch name {
	case "gd.String", "gd.StringName":
		return fmt.Sprintf("%v.String()", val)
	case "gd.Int":
		return fmt.Sprintf("int(%v)", val)
	case "gd.Float":
		return fmt.Sprintf("float64(%v)", val)
	case "gd.PackedByteArray":
		return fmt.Sprintf("%v.Bytes()", val)
	default:
		return val
	}
}
