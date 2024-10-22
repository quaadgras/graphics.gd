package main

import (
	"fmt"
	"io"
	"strings"

	"grow.graphics/gd/internal/gdjson"
	"grow.graphics/gd/internal/gdtype"
)

type callType int

const (
	callDefault callType = iota
	callBuiltin
	callUtility
)

func (classDB ClassDB) signalCall(w io.Writer, class gdjson.Class, signal gdjson.Signal, singleton bool) {
	if singleton {
		fmt.Fprintf(w, "\nfunc On%v(cb func(", convertName(signal.Name))
	} else {
		fmt.Fprintf(w, "\nfunc (self Go) On%v(cb func(", convertName(signal.Name))
	}
	for i, arg := range signal.Arguments {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(arg.Meta, arg.Type))
	}
	fmt.Fprint(w, ")) {\n\tgc := gd.GarbageCollector(); _ = gc\n\t")
	fmt.Fprintf(w, `self[0].AsObject().Connect(gc.StringName("%s"), gc.Callable(cb), 0)`, signal.Name)
	fmt.Fprint(w, "\n}\n\n")
}

func (classDB ClassDB) simpleCall(w io.Writer, class gdjson.Class, method gdjson.Method, singleton bool) {
	if method.IsVararg {
		return
	}
	if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
		return
	}
	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	if method.Description != "" {
		fmt.Fprintln(w, "\n/*")
		fmt.Fprint(w, method.Description)
		fmt.Fprintln(w, "\n*/")
	}
	if method.IsVirtual {
		classDB.simpleVirtualCall(w, class, method)
		return
	}
	resultSimple := classDB.convertTypeSimple(method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := classDB.convertType("", method.ReturnValue.Meta, method.ReturnValue.Type)
	_, needsLifetime := classDB.isPointer(resultExpert)
	if method.IsStatic {
		needsLifetime = true
	}
	if singleton {
		fmt.Fprintf(w, "func %v(", convertName(method.Name))
	} else {
		fmt.Fprintf(w, "func (self Go) %v(", convertName(method.Name))
	}
	for i, arg := range method.Arguments {
		if arg.DefaultValue == nil {
			if i > 0 {
				fmt.Fprint(w, ", ")
			}
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(arg.Meta, arg.Type))
		}
	}
	fmt.Fprint(w, ") ")
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "%v ", resultSimple)
	}
	fmt.Fprint(w, "{\n\tgc := gd.GarbageCollector(); _ = gc\n\t")
	if singleton {
		fmt.Fprintf(w, "once.Do(singleton)\n\t")
	}
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "return %s(", resultSimple)
	}
	var call strings.Builder
	fmt.Fprintf(&call, "class(self).%v(", convertName(method.Name))
	if needsLifetime {
		fmt.Fprint(&call, "gc")
	}
	for i, arg := range method.Arguments {
		if i > 0 || needsLifetime {
			fmt.Fprint(&call, ", ")
		}
		val := fixReserved(arg.Name)
		if arg.DefaultValue != nil {
			val = *arg.DefaultValue
			val = strings.TrimPrefix(val, "&")
			if val == "null" || val == "[]" || val == "{}" || strings.HasSuffix(val, "()") || strings.HasSuffix(val, "[])") {
				val = "([1]" + classDB.convertTypeSimple(arg.Meta, arg.Type) + "{}[0])"
			} else {
				if strings.Contains(val, "(") {
					switch {
					case strings.HasPrefix(val, "Rect2("), strings.HasPrefix(val, "Rect2i("), strings.HasPrefix(val, "Transform2D("),
						strings.HasPrefix(val, "Transform3D("):
						val = "gd.New" + val
					case strings.HasPrefix(val, "StringName(\""):
						val = strings.TrimSuffix(strings.TrimPrefix(val, "StringName(\""), "\")")
					case strings.HasPrefix(val, "NodePath(\""):
						val = strings.TrimSuffix(strings.TrimPrefix(val, "NodePath(\""), "\")")
						if val == "" {
							val = `""`
						}
					default:
						val = "gd." + strings.ReplaceAll(strings.ReplaceAll(val, "(", "{"), ")", "}")
					}
				}
			}
			if gdtype.Name(classDB.convertType("", arg.Meta, arg.Type)) == "gd.Variant" {
				val = `gc.Variant(` + val + `)`
			}
		}
		fmt.Fprint(&call, gdtype.Name(classDB.convertType("", arg.Meta, arg.Type)).ConvertToSimple(val))
	}
	fmt.Fprint(&call, ")")
	fmt.Fprint(w, gdtype.Name(resultExpert).ConvertToGo(call.String()))
	if method.ReturnValue.Type != "" {
		fmt.Fprint(w, ")")
	}
	fmt.Fprintf(w, "\n}\n")
}

func (classDB ClassDB) simpleVirtualCall(w io.Writer, class gdjson.Class, method gdjson.Method) {
	resultSimple := classDB.convertTypeSimple(method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := classDB.convertType("", method.ReturnValue.Meta, method.ReturnValue.Type)
	_, needsLifetime := classDB.isPointer(resultExpert)
	if method.IsStatic {
		needsLifetime = true
	}
	fmt.Fprintf(w, "func (Go) %s(impl func(ptr unsafe.Pointer", method.Name)
	for _, arg := range method.Arguments {
		fmt.Fprint(w, ", ")
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(arg.Meta, arg.Type))
	}
	fmt.Fprintf(w, ") %v, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {\n", resultSimple)
	fmt.Fprintf(w, "\treturn func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {\n")
	fmt.Fprintf(w, "\t\tgc := gd.NewLifetime(api)\n")
	fmt.Fprintf(w, "\t\tclass.SetTemporary(gc)\n")
	for i, arg := range method.Arguments {
		var expert = classDB.convertType("", arg.Meta, arg.Type)

		_, ok := classDB[arg.Type]
		if ok {
			fmt.Fprintf(w, "\t\tvar %v %v\n", fixReserved(arg.Name), expert)
			if arg.Type == "Object" {
				fmt.Fprintf(w, "\t\t%v.SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,%d)}))\n", fixReserved(arg.Name), i)
			} else {
				fmt.Fprintf(w, "\t\t%v[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,%d)}))\n", fixReserved(arg.Name), i)
			}
			continue
		}
		argPtrKind, argIsPtr := classDB.isPointer(expert)
		if argIsPtr {
			fmt.Fprintf(w, "\t\tvar %v = %v\n", fixReserved(arg.Name),
				gdtype.Name(expert).Let("gc", fmt.Sprintf("gd.UnsafeGet[%v](p_args,%d)", argPtrKind, i)))
		} else {
			fmt.Fprintf(w, "\t\tvar %v = gd.UnsafeGet[%v](p_args,%d)\n", fixReserved(arg.Name), expert, i)
		}
	}
	fmt.Fprintf(w, "\t\tself := reflect.ValueOf(class).UnsafePointer()\n")
	if resultSimple != "" {
		fmt.Fprintf(w, "\t\tret := ")
	}
	fmt.Fprintf(w, "impl(self")
	for _, arg := range method.Arguments {
		fmt.Fprint(w, ", ")
		fmt.Fprintf(w, "%v", gdtype.Name(classDB.convertType("", arg.Meta, arg.Type)).ConvertToGo(fixReserved(arg.Name)))
	}
	fmt.Fprintf(w, ")\n")
	if resultSimple != "" {
		ret := gdtype.Name(resultExpert).ToUnderlying(gdtype.Name(resultExpert).ConvertToSimple("ret"))
		if needsLifetime {
			name := strings.TrimPrefix(resultExpert, "gdclass.")
			name = strings.TrimPrefix(name, "[1]classdb.")
			_, ok := classDB[name]
			if resultExpert == "gd.Object" {
				ret = fmt.Sprintf("mmm.End(%s.AsPointer())", ret)
			} else if ok {
				ret = fmt.Sprintf("mmm.End(%s[0].AsPointer())", ret)
			} else {
				ret = fmt.Sprintf("mmm.End(%s)", ret)
			}
		}
		fmt.Fprintf(w, "\t\tgd.UnsafeSet(p_back, %s)\n", ret)
	}
	fmt.Fprintf(w, "\t\tgc.End()\n")
	fmt.Fprintf(w, "\t}\n")
	fmt.Fprintf(w, "}\n")
	return
}

func (classDB ClassDB) methodCall(w io.Writer, pkg string, class gdjson.Class, method gdjson.Method, ctype callType) {
	if ctype == callDefault && method.IsVararg {
		return
	}

	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	result := classDB.convertType(pkg, method.ReturnValue.Meta, method.ReturnValue.Type)
	if method.ReturnType != "" {
		result = classDB.convertType(pkg, "", method.ReturnType)
	}
	ptrKind, isPtr := classDB.isPointer(result)

	prefix := ""
	if pkg != "internal" {
		prefix = "gd."
	}

	fmt.Fprintln(w)
	if method.Description != "" {
		fmt.Fprintln(w, "/*")
		fmt.Fprint(w, method.Description)
		fmt.Fprintln(w, "\n*/")
	}

	if method.IsVirtual {
		fmt.Fprintf(w, "func (class) %s(impl func(ptr unsafe.Pointer", method.Name)
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(pkg, arg.Meta, arg.Type))
		}
		fmt.Fprintf(w, ") %v, api *"+prefix+"API) (cb "+prefix+"ExtensionClassCallVirtualFunc) {\n", result)
		fmt.Fprintf(w, "\treturn func(class "+prefix+"ExtensionClass, p_args "+prefix+"UnsafeArgs, p_back "+prefix+"UnsafeBack) {\n")
		fmt.Fprintf(w, "\t\tctx := %vNewLifetime(api)\n", prefix)
		fmt.Fprintf(w, "\t\tclass.SetTemporary(ctx)\n")
		for i, arg := range method.Arguments {
			var argType = classDB.convertType(pkg, arg.Meta, arg.Type)

			_, ok := classDB[arg.Type]
			if ok {
				fmt.Fprintf(w, "\t\tvar %v %v\n", fixReserved(arg.Name), argType)
				if arg.Type == "Object" {
					fmt.Fprintf(w, "\t\t%v.SetPointer(mmm.Let["+prefix+"Pointer](ctx.Lifetime, ctx.API, [2]uintptr{"+prefix+"UnsafeGet[uintptr](p_args,%d)}))\n", fixReserved(arg.Name), i)
				} else {
					fmt.Fprintf(w, "\t\t%v[0].SetPointer(mmm.Let["+prefix+"Pointer](ctx.Lifetime, ctx.API, [2]uintptr{"+prefix+"UnsafeGet[uintptr](p_args,%d)}))\n", fixReserved(arg.Name), i)
				}
				continue
			}

			argPtrKind, argIsPtr := classDB.isPointer(argType)
			if argIsPtr {
				fmt.Fprintf(w, "\t\tvar %v = %v\n", fixReserved(arg.Name),
					gdtype.Name(argType).Let("ctx", fmt.Sprintf(prefix+"UnsafeGet[%v](p_args,%d)", argPtrKind, i)))
			} else {
				fmt.Fprintf(w, "\t\tvar %v = "+prefix+"UnsafeGet[%v](p_args,%d)\n", fixReserved(arg.Name), argType, i)
			}
		}
		fmt.Fprintf(w, "\t\tself := reflect.ValueOf(class).UnsafePointer()\n")
		if result != "" {
			fmt.Fprintf(w, "\t\tret := ")
		}
		fmt.Fprintf(w, "impl(self")
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v", fixReserved(arg.Name))
		}
		fmt.Fprintf(w, ")\n")
		if result != "" {
			ret := gdtype.Name(result).ToUnderlying("ret")
			if isPtr {
				_, ok := classDB[strings.TrimPrefix(result, "gdclass.")]
				if result == "gd.Object" {
					ret = fmt.Sprintf("mmm.End(%s.AsPointer())", ret)
				} else if ok {
					ret = fmt.Sprintf("mmm.End(%s[0].AsPointer())", ret)
				} else {
					ret = fmt.Sprintf("mmm.End(%s)", ret)
				}
			}
			fmt.Fprintf(w, "\t\t"+prefix+"UnsafeSet(p_back, %s)\n", ret)
		}
		fmt.Fprintf(w, "\t\tctx.End()\n")
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "}\n")
		return
	}
	if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
		return
	}
	var context = "ctx " + prefix + "Lifetime"
	if !isPtr && !method.IsStatic {
		context = ""
	}

	if ctype == callBuiltin && strings.HasPrefix(class.Name, "Packed") {
		fmt.Fprintf(w, "//go:nosplit\nfunc (self *class) %v(%s", convertName(method.Name), context)
	} else {
		fmt.Fprintf(w, "//go:nosplit\nfunc (self class) %v(%s", convertName(method.Name), context)
	}

	if method.Name == "select" {
		method.Name = "select_"
	}
	if method.Name == "map" {
		method.Name = "map_"
	}

	for i, arg := range method.Arguments {
		if i > 0 || context != "" {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(pkg, arg.Meta, arg.Type))
	}
	if method.IsVararg {
		if len(method.Arguments) > 0 || context != "" {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "args ..."+prefix+"Variant")
	}
	fmt.Fprintf(w, ") %v {\n", result)
	if !method.IsStatic {
		if ctype == callBuiltin {
			fmt.Fprintf(w, "\tvar selfPtr = self\n")
		} else {
			fmt.Fprintf(w, "\tvar selfPtr = self[0].AsPointer()\n")
		}
	}
	fmt.Fprintf(w, "\tvar frame = callframe.New()\n")
	for _, arg := range method.Arguments {
		_, ok := classDB[arg.Type]
		if ok {
			if arg.Type == "Object" {
				fmt.Fprintf(w, "\tcallframe.Arg(frame, mmm.End(%v.AsPointer())[0])\n", fixReserved(arg.Name))
			} else {
				switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name][arg.Name]; semantics {
				case gdjson.OwnershipTransferred, gdjson.LifetimeBoundToClass:
					fmt.Fprintf(w, "\tcallframe.Arg(frame, mmm.End(%v[0].AsPointer())[0])\n", fixReserved(arg.Name))
				case gdjson.RefCountedManagement, gdjson.IsTemporaryReference, gdjson.MustAssertInstanceID, gdjson.ReversesTheOwnership:
					fmt.Fprintf(w, "\tcallframe.Arg(frame, mmm.Get(%v[0].AsPointer())[0])\n", fixReserved(arg.Name))
				default:
					panic("unknown ownership: " + fmt.Sprint(semantics))
				}
			}
			continue
		}

		argType := classDB.convertType(pkg, arg.Meta, arg.Type)
		_, argIsPtr := classDB.isPointer(argType)
		if argIsPtr {
			fmt.Fprintf(w, "\tcallframe.Arg(frame, mmm.Get(%v))\n", fixReserved(arg.Name))
		} else {
			fmt.Fprintf(w, "\tcallframe.Arg(frame, %v)\n", fixReserved(arg.Name))
		}
	}
	if method.IsVararg {
		fmt.Fprintf(w, "\tfor _, arg := range args {\n")
		fmt.Fprintf(w, "\t\tcallframe.Arg(frame, mmm.Get(arg))\n")
		fmt.Fprintf(w, "\t}\n")
	}
	if isPtr {
		fmt.Fprintf(w, "\tvar r_ret = callframe.Ret[%v](frame)\n", ptrKind)
	} else {
		if result != "" {
			fmt.Fprintf(w, "\tvar r_ret = callframe.Ret[%v](frame)\n", result)
		} else {
			fmt.Fprintf(w, "\tvar r_ret callframe.Nil\n")
		}
	}

	var API = "mmm.API(selfPtr)"
	if method.IsStatic {
		API = "ctx.API"
	} else if ctype == callBuiltin {
		API = "mmm.API(self)"
		if strings.HasPrefix(class.Name, "Packed") {
			API = "mmm.API(*self)"
		}
	}

	if ctype == callBuiltin {
		if strings.HasPrefix(class.Name, "Packed") {
			var self = "0"
			if !method.IsStatic {
				fmt.Fprintf(w, "\tvar p_self = callframe.Arg(frame, mmm.Get(selfPtr))\n")
				self = "p_self.Uintptr()"
			}
			if method.IsVararg {
				fmt.Fprintf(w, "\t%s.builtin.%v.%v(%s, frame.Array(0), r_ret.Uintptr(), int32(len(args))+%d)\n", API, class.Name, method.Name, self, len(method.Arguments))
			} else {
				fmt.Fprintf(w, "\t%s.builtin.%v.%v(%s, frame.Array(0), r_ret.Uintptr(), %d)\n", API, class.Name, method.Name, self, len(method.Arguments))
			}
			fmt.Fprintf(w, "\tmmm.Set(selfPtr, p_self.Get())\n")
		} else {
			var self = "0"
			if !method.IsStatic {
				fmt.Fprintf(w, "\tvar p_self = callframe.Arg(frame, mmm.Get(selfPtr))\n")
				self = "p_self.Uintptr()"
			}
			if method.IsVararg {
				fmt.Fprintf(w, "\t%s.builtin.%v.%v(%s, frame.Array(0), r_ret.Uintptr(), int32(len(args))+%d)\n", API, class.Name, method.Name, self, len(method.Arguments))
			} else {
				fmt.Fprintf(w, "\t%s.builtin.%v.%v(%s, frame.Array(0), r_ret.Uintptr(), %d)\n", API, class.Name, method.Name, self, len(method.Arguments))
			}
		}
	} else {
		if method.IsVararg {
			fmt.Fprintf(w, "\tif len(args) > 0 { panic(`varargs not supported for class methods yet`); }\n")
		}
		fmt.Fprintf(w, "\t%s.Object.MethodBindPointerCall(%[1]s.Methods.%v.Bind_%v, self.AsObject(), frame.Array(0), r_ret.Uintptr())\n", API, class.Name, method.Name)
	}

	if isPtr {
		_, ok := classDB[strings.TrimPrefix(result, "gdclass.")]
		if ok || result == "gd.Object" {
			fmt.Fprintf(w, "\tvar ret %v\n", result)

			returns := "\tret"
			if result != "gd.Object" {
				returns = "\tret[0]"
			}

			switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name]["return value"]; semantics {
			case gdjson.RefCountedManagement, gdjson.OwnershipTransferred:
				fmt.Fprintf(w, returns+".SetPointer("+prefix+"PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))\n")
			case gdjson.LifetimeBoundToClass:
				fmt.Fprintf(w, returns+".SetPointer("+prefix+"PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))\n")
			case gdjson.MustAssertInstanceID:
				fmt.Fprintf(w, returns+".SetPointer("+prefix+"PointerMustAssertInstanceID(ctx, r_ret.Get()))\n")
			default:
				panic("unknown ownership: " + fmt.Sprint(semantics))
			}

		} else {
			if strings.HasPrefix(result, "ArrayOf") {
				fmt.Fprint(w, "\tvar ret = mmm.New[Array](ctx.Lifetime, ctx.API, r_ret.Get())\n")
			} else if strings.HasPrefix(result, "gd.ArrayOf") {
				fmt.Fprint(w, "\tvar ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())\n")
			} else {
				fmt.Fprintf(w, "\tvar ret = mmm.New[%v](ctx.Lifetime, ctx.API, r_ret.Get())\n", result)
			}
		}
	} else if result != "" {
		fmt.Fprintf(w, "\tvar ret = r_ret.Get()\n")
	}
	fmt.Fprintf(w, "\tframe.Free()\n")
	if result != "" {
		if strings.HasPrefix(result, "ArrayOf") || strings.HasPrefix(result, "gd.ArrayOf") {
			result = strings.ReplaceAll(result, "gd.ArrayOf", "gd.TypedArray")
			result = strings.ReplaceAll(result, "ArrayOf", "TypedArray")
			fmt.Fprintf(w, "\treturn %s(ret)\n", result)
		} else {
			fmt.Fprintf(w, "\treturn ret\n")
		}
	}
	/*if result != "" {
		fmt.Fprintf(w, "\tvar ret %v\n", result)
		fmt.Fprintf(w, "\treturn ret\n")
	}*/
	fmt.Fprintf(w, "}")
}
