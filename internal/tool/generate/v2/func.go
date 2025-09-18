package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
)

func (classDB ClassDB) signalCall(w io.Writer, class gdjson.Class, signal gdjson.Signal, singleton bool) {
	if signal.Description != "" {
		fmt.Fprintln(w, "\n/*")
		fmt.Fprint(w, gdjson.DocsToGoDoc(signal.Description, classDB, class.Name, class.Name+"_"+convertName(signal.Name)))
		fmt.Fprint(w, "\n*/")
	}
	if singleton {
		fmt.Fprintf(w, "\nfunc On%v(cb func(", convertName(signal.Name))
	} else {
		fmt.Fprintf(w, "\nfunc (self Instance) On%v(cb func(", convertName(signal.Name))
	}
	for i, arg := range signal.Arguments {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, class.Name+"."+signal.Name+"."+arg.Name, arg.Meta, arg.Type))
	}
	fmt.Fprint(w, "), flags ...Signal.Flags) {\n\t")
	fmt.Fprintln(w, "var flags_together Signal.Flags")
	fmt.Fprint(w, "\tfor _, flag := range flags {\n")
	fmt.Fprint(w, "\t\tflags_together |= flag\n")
	fmt.Fprint(w, "\t}\n\t")
	fmt.Fprintf(w, `self[0].AsObject()[0].Connect(gd.NewStringName("%s"), gd.NewCallable(cb), int64(flags_together))`, signal.Name)
	fmt.Fprint(w, "\n}\n\n")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "\nfunc (self class) %s() Signal.Any { return Signal.Via(gd.SignalProxy{}, pointers.Pack(gd.NewSignalOf(self.AsObject(), gd.NewStringName(`%[1]s`)))) }\n", convertName(signal.Name))
}

func (classDB ClassDB) simpleCall(w io.Writer, class gdjson.Class, method gdjson.Method, singleton, defaults bool) {
	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	if method.Description != "" {
		fmt.Fprintln(w, "\n/*")
		fmt.Fprint(w, gdjson.DocsToGoDoc(method.Description, classDB, class.Name, class.Name+"_"+convertName(method.Name)))
		fmt.Fprintln(w, "\n*/")
	}
	if method.IsVirtual {
		classDB.simpleVirtualCall(w, class, method)
		return
	}

	resultSimple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+".", method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := gdtype.EngineTypeAsGoType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
	resultsSimple := []string{resultSimple}

	var skips map[string]reflect.Type
	if returns, ok := gdjson.Returnables[class.Name+"."+method.Name]; ok {
		skips = make(map[string]reflect.Type, len(returns))
		for name := range returns {
			skips[name] = returns[name]
			resultsSimple = append([]string{returns[name].String()}, resultsSimple...)
		}
	}

	if singleton || method.IsStatic {
		if defaults {
			fmt.Fprintf(w, "func %v(", convertName(method.Name))
		} else {
			fmt.Fprintf(w, "func %vOptions(", convertName(method.Name))
		}
	} else {
		if defaults {
			fmt.Fprintf(w, "func (self Instance) %v(", convertName(method.Name))
		} else {
			fmt.Fprintf(w, "func (self MoreArgs) %v(", convertName(method.Name))
		}
	}
	var first = true
	for _, arg := range method.Arguments {
		if skips[arg.Name] != nil {
			continue
		}
		if !defaults || arg.DefaultValue == nil || ((singleton || method.IsStatic) && arg.DefaultValue != nil && gdjson.IsTheDefaultValueZero(*arg.DefaultValue)) {
			if !first {
				fmt.Fprint(w, ", ")
			}
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type))
			first = false
		}
	}
	if method.IsVararg {
		if !first {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprint(w, "args ...any")
	}
	fmt.Fprint(w, ") ")

	if multiple, ok := gdjson.Unpackables[class.Name+"."+method.Name]; ok {
		fmt.Fprintf(w, "(")
		for i, ret := range multiple {
			if i > 0 {
				fmt.Fprint(w, ", ")
			}
			resultsSimple = append(resultsSimple, ret.String())
			fmt.Fprint(w, ret)
		}
		fmt.Fprintf(w, ") ")
	} else {
		if len(resultsSimple) > 1 {
			fmt.Fprintf(w, "(")
			for i, ret := range resultsSimple {
				if i > 0 {
					fmt.Fprint(w, ", ")
				}
				fmt.Fprint(w, ret)
			}
			fmt.Fprintf(w, ") ")
		} else {
			if method.ReturnValue.Type != "" {
				fmt.Fprintf(w, "%v ", resultSimple)
			}
		}
	}
	fmt.Fprintf(w, "{ //gd:%s.%s\n\t", class.Name, method.Name)
	if singleton {
		fmt.Fprintf(w, "once.Do(singleton)\n\t")
	}
	for name := range skips {
		fmt.Fprintf(w, "var returns_%s = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(gd.NewArray()))\n\t", name)
	}
	if method.IsStatic && !singleton {
		fmt.Fprintf(w, "self := Instance{}\n")
	}
	if method.IsVararg {
		fmt.Fprint(w, "var converted_variants = make([]gd.Variant, len(args))\n")
		fmt.Fprint(w, "for i, arg := range args {\n")
		fmt.Fprint(w, "\tconverted_variants[i] = gd.NewVariant(arg)\n")
		fmt.Fprint(w, "}\n")
	}
	if len(resultsSimple) == 1 {
		if method.ReturnValue.Type != "" {
			fmt.Fprintf(w, "return %s(", resultSimple)
		}
	}
	if len(resultsSimple) > 1 {
		fmt.Fprintf(w, "results := ")
	}
	var call strings.Builder
	if singleton {
		fmt.Fprintf(&call, "Advanced().%v(", convertName(method.Name))
	} else {
		fmt.Fprintf(&call, "Advanced(self).%v(", convertName(method.Name))
	}
	for i, arg := range method.Arguments {
		if i > 0 {
			fmt.Fprint(&call, ", ")
		}
		if skips[arg.Name] != nil {
			fmt.Fprint(&call, "returns_"+arg.Name)
			continue
		}
		val := fixReserved(arg.Name)
		if arg.DefaultValue != nil && defaults && !((singleton || method.IsStatic) && gdjson.IsTheDefaultValueZero(*arg.DefaultValue)) {
			switch arg.Type {
			case "Array":
				val = "Array.Nil"
			case "Callable":
				val = "Callable.Nil"
			case "Dictionary":
				val = "Dictionary.Nil"
			default:
				val = *arg.DefaultValue
				val = strings.TrimPrefix(val, "&")
				if val == "null" || val == "[]" || val == "{}" || strings.HasSuffix(val, "()") || strings.HasSuffix(val, "[])") {
					if arg.Type == "Callable" {
						val = "nil"
					} else {
						val = "([1]" + classDB.convertTypeSimple(class, "", arg.Meta, arg.Type) + "{}[0])"
					}
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
			}
			if gdtype.Name(gdtype.EngineTypeAsGoType(class.Name, arg.Meta, arg.Type)) == "gd.Variant" {
				val = `gd.NewVariant(` + val + `)`
			}
		}
		simple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type)
		fmt.Fprint(&call, gdtype.Name(gdtype.EngineTypeAsGoType(class.Name, arg.Meta, arg.Type)).ConvertToSimple(val, simple))
	}
	if method.IsVararg {
		if len(method.Arguments) > 0 {
			fmt.Fprint(&call, ", ")
		}
		fmt.Fprint(&call, "converted_variants...")
	}
	fmt.Fprint(&call, ")")
	if len(resultsSimple) < 2 {
		fmt.Fprint(w, gdtype.Name(resultExpert).ConvertToGo(call.String(), resultSimple))
	} else if len(skips) == 0 && strings.HasPrefix(method.ReturnValue.Type, "Array") {
		fmt.Fprint(w, "gd.InternalArray(", call.String(), ")")
	} else {
		fmt.Fprint(w, call.String())
	}
	if len(resultsSimple) == 1 {
		if method.ReturnValue.Type != "" {
			fmt.Fprint(w, ")")
		}
	}
	if len(resultsSimple) > 1 {
		if len(skips) > 0 {
			fmt.Fprint(w, "\n\treturn ")
			var first = true
			for skip, rtype := range skips {
				if !first {
					fmt.Fprint(w, ",")
				}
				first = false
				if rtype.Kind() == reflect.Slice {
					fmt.Fprintf(w, "gd.ArrayAs[%s](gd.InternalArray(returns_%s)), ", rtype.String(), skip)
				} else {
					fmt.Fprintf(w, "gd.VariantAs[%s](gd.InternalArray(returns_%s).Index(0)), ", rtype.String(), skip)
				}
			}
			fmt.Fprint(w, gdtype.Name(resultExpert).ConvertToGo("results", resultSimple))
		} else {
			fmt.Fprint(w, "\n\treturn ")
			if resultSimple == "Vector2i.XY" || resultSimple == "Vector2.XY" {
				fmt.Fprintf(w, "%s(results.X), %s(results.Y)", resultsSimple[1], resultsSimple[1])
			} else {
				for i, ret := range resultsSimple[1:] {
					if i > 0 {
						fmt.Fprint(w, ",")
					}
					fmt.Fprintf(w, "gd.VariantAs[%s](results.Index(%d))", ret, i)
				}
			}
		}
	}
	fmt.Fprintf(w, "\n}\n")
}

func (classDB ClassDB) simpleRelocatedCall(w io.Writer, class gdjson.Class, method gdjson.Method, relocated_from gdjson.Class, original_name string, defaults bool) {
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
	resultSimple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+".", method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := gdtype.EngineTypeAsGoType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
	if method.IsStatic {
		if defaults {
			fmt.Fprintf(w, "func %v(", convertName(method.Name))
		} else {
			fmt.Fprintf(w, "func %vOptions(", convertName(method.Name))
		}
	} else {
		if defaults {
			fmt.Fprintf(w, "func (self Instance) %v(", convertName(method.Name))
		} else {
			fmt.Fprintf(w, "func (self MoreArgs) %v(", convertName(method.Name))
		}
	}
	var first = true
	for _, arg := range method.Arguments {
		if !defaults || arg.DefaultValue == nil || (method.IsStatic && arg.DefaultValue != nil && gdjson.IsTheDefaultValueZero(*arg.DefaultValue)) {
			if !first {
				fmt.Fprint(w, ", ")
			}
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type))
			first = false
		}
	}
	if method.IsVararg {
		if !first {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprint(w, "args ...any")
	}
	fmt.Fprint(w, ") ")
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "%v ", resultSimple)
	}
	fmt.Fprintf(w, "{ //gd:%s.%s\n\t", relocated_from.Name, original_name)
	if method.IsVararg {
		fmt.Fprint(w, "var converted_variants = make([]gd.Variant, len(args))\n")
		fmt.Fprint(w, "for i, arg := range args {\n")
		fmt.Fprint(w, "\tconverted_variants[i] = gd.NewVariant(arg)\n")
		fmt.Fprint(w, "}\n")
	}
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "return %s(", resultSimple)
	}
	var call strings.Builder
	fmt.Fprintf(&call, "%s.Advanced(peer).%v(", relocated_from.Name, convertName(original_name))
	for i, arg := range method.ArgumentsRemapped {
		if i > 0 {
			fmt.Fprint(&call, ", ")
		}
		val := fixReserved(arg.Name)
		if arg.DefaultValue != nil && defaults && !((method.IsStatic) && gdjson.IsTheDefaultValueZero(*arg.DefaultValue)) {
			switch arg.Type {
			case "Array":
				val = "Array.Nil"
			case "Callable":
				val = "Callable.Nil"
			case "Dictionary":
				val = "Dictionary.Nil"
			default:
				val = *arg.DefaultValue
				val = strings.TrimPrefix(val, "&")
				if val == "null" || val == "[]" || val == "{}" || strings.HasSuffix(val, "()") || strings.HasSuffix(val, "[])") {
					if arg.Type == "Callable" {
						val = "nil"
					} else {
						val = "([1]" + classDB.convertTypeSimple(class, "", arg.Meta, arg.Type) + "{}[0])"
					}
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
			}
			if gdtype.Name(gdtype.EngineTypeAsGoType(class.Name, arg.Meta, arg.Type)) == "gd.Variant" {
				val = `gd.NewVariant(` + val + `)`
			}
		}
		simple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type)
		fmt.Fprint(&call, gdtype.Name(gdtype.EngineTypeAsGoType(class.Name, arg.Meta, arg.Type)).ConvertToSimple(val, simple))
	}
	if method.IsVararg {
		if len(method.Arguments) > 0 {
			fmt.Fprint(&call, ", ")
		}
		fmt.Fprint(&call, "converted_variants...")
	}
	fmt.Fprint(&call, ")")
	fmt.Fprint(w, gdtype.Name(resultExpert).ConvertToGo(call.String(), resultSimple))
	if method.ReturnValue.Type != "" {
		fmt.Fprint(w, ")")
	}
	fmt.Fprintf(w, "\n}\n")
}

func (classDB ClassDB) simpleVirtualCall(w io.Writer, class gdjson.Class, method gdjson.Method) {
	resultSimple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+".", method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := gdtype.EngineTypeAsGoType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
	_, needsLifetime := gdtype.Name(resultExpert).IsPointer()
	if method.IsStatic {
		needsLifetime = true
	}
	fmt.Fprintf(w, "func (Instance) %s(impl func(ptr gdclass.Receiver", method.Name)
	for _, arg := range method.Arguments {
		fmt.Fprint(w, ", ")
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type))
	}
	fmt.Fprintf(w, ") %v) (cb gd.ExtensionClassCallVirtualFunc) {\n", resultSimple)
	fmt.Fprintf(w, "\treturn func(class any, p_args, p_back gdextension.Pointer) {\n")
	for i, arg := range method.Arguments {
		var expert = gdtype.EngineTypeAsGoType(class.Name, arg.Meta, arg.Type)
		pointerKind, argIsPtr := gdtype.Name(expert).IsPointer()
		if !argIsPtr {
			pointerKind = expert
		}
		fmt.Fprintf(w, "\t\tvar %v = %v\n", fixReserved(arg.Name), gdtype.Name(expert).LoadFromRawPointerValue(
			fmt.Sprintf("gd.UnsafeGet[%v](p_args,%d)", pointerKind, i),
		))
		if argIsPtr {
			fmt.Fprintf(w, "\t\tdefer %s\n", gdtype.Name(expert).EndPointer(fixReserved(arg.Name)))
		}
	}
	fmt.Fprintf(w, "\t\tself := gdclass.Receiver(reflect.ValueOf(class).UnsafePointer())\n")
	if resultSimple != "" {
		fmt.Fprintf(w, "\t\tret := ")
	}
	fmt.Fprintf(w, "impl(self")
	for _, arg := range method.Arguments {
		simple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type)
		fmt.Fprint(w, ", ")
		fmt.Fprintf(w, "%v", gdtype.Name(gdtype.EngineTypeAsGoType(class.Name, arg.Meta, arg.Type)).ConvertToGo(fixReserved(arg.Name), simple))
	}
	fmt.Fprintf(w, ")\n")
	if resultSimple != "" {
		simple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+".", method.ReturnValue.Meta, method.ReturnValue.Type)
		ret := gdtype.Name(resultExpert).ToUnderlying(gdtype.Name(resultExpert).ConvertToSimple("ret", simple))
		if needsLifetime {
			fmt.Fprintf(w, "ptr, ok := %s\n", gdtype.Name(resultExpert).EndPointer(ret))
			fmt.Fprintf(w, "\n\t\tif !ok {\n")
			fmt.Fprintf(w, "\t\t\treturn\n")
			fmt.Fprintf(w, "\t\t}\n")
			ret = "ptr"
		}
		fmt.Fprintf(w, "\t\tgd.UnsafeSet(p_back, %s)\n", ret)
	}
	fmt.Fprintf(w, "\t}\n")
	fmt.Fprintf(w, "}\n")
}
