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

func (classDB ClassDB) simpleCall(w io.Writer, class gdjson.Class, method gdjson.Method) {
	if method.IsVararg || method.IsVirtual {
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
	fmt.Fprintf(w, "func (self Simple) %v(", convertName(method.Name))
	for i, arg := range method.Arguments {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(arg.Meta, arg.Type))
	}
	fmt.Fprint(w, ") ")
	resultSimple := classDB.convertTypeSimple(method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := classDB.convertType("", method.ReturnValue.Meta, method.ReturnValue.Type)
	_, needsLifetime := classDB.isPointer(resultExpert)
	if method.IsStatic {
		needsLifetime = true
	}
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "%v ", resultSimple)
	}
	fmt.Fprint(w, "{\n\tgc := gd.GarbageCollector(); _ = gc\n\t")
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "return %s(", resultSimple)
	}
	var call strings.Builder
	fmt.Fprintf(&call, "Expert(self).%v(", convertName(method.Name))
	if needsLifetime {
		fmt.Fprint(&call, "gc")
	}
	for i, arg := range method.Arguments {
		if i > 0 || needsLifetime {
			fmt.Fprint(&call, ", ")
		}
		fmt.Fprint(&call, gdtype.Name(classDB.convertType("", arg.Meta, arg.Type)).ConvertToSimple(fixReserved(arg.Name)))
	}
	fmt.Fprint(&call, ")")
	fmt.Fprint(w, gdtype.Name(resultExpert).ConvertToGo(call.String()))
	if method.ReturnValue.Type != "" {
		fmt.Fprint(w, ")")
	}
	fmt.Fprintf(w, "\n}\n")
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
				_, ok := classDB[strings.TrimPrefix(result, "object.")]
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
		_, ok := classDB[strings.TrimPrefix(result, "object.")]
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
