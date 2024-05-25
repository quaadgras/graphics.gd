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
