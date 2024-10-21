package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"grow.graphics/gd/internal/gdjson"
	"grow.graphics/gd/internal/gdtype"
)

// LoadSpecification, either from a local file or by downloading
// it from the Godot Github repository.
func LoadSpecification() (*gdjson.Specification, error) {
	file, err := os.Open("./extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		file, err = os.Create("./extension_api.json")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return nil, err
		}
		file.Seek(0, 0)
		resp.Body.Close()
	}
	var spec gdjson.Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return nil, err
	}
	return &spec, nil
}

func hasNative(gdType string) bool {
	switch gdType {
	case "int", "float", "String", "bool",
		"PackedStringArray", "PackedInt32Array", "PackedInt64Array",
		"PackedFloat32Array", "PackedFloat64Array", "PackedVector2Array",
		"PackedVector3Array", "PackedVector4Array", "PackedColorArray", "Variant":
		return true
	}
	return false
}

func (classDB ClassDB) convertType(pkg, meta string, gdType string) string {
	maybeInternal := func(name string) string {
		if pkg != "internal" {
			return "gd." + name
		}
		return name
	}
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		return maybeInternal("ArrayOf[" + classDB.convertType(pkg, meta, gdType) + "]")
	}
	switch gdType {
	case "int", "Int":
		return maybeInternal("Int")
	case "float", "Float":
		return maybeInternal("Float")
	case "bool", "Bool":
		return "bool"
	case "String":
		return maybeInternal("String")
	case "StringName":
		return maybeInternal("StringName")
	case "enum::GDExtension.InitializationLevel":
		return maybeInternal("GDExtensionInitializationLevel")
	case "enum::GDExtensionManager.LoadStatus":
		return "GDExtensionManagerLoadStatus"
	case "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedVector4Array", "PackedColorArray", "PackedByteArray",
		"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID", "Object",
		"Callable", "Signal", "Dictionary", "Array":
		return maybeInternal(gdType)
	case "Variant":
		return maybeInternal("Variant")
	case "enum::Variant.Type":
		return maybeInternal("VariantType")

	// strange C++ cases
	case "enum::Error":
		return "int64"
	case "const uint8_t **":
		return "unsafe.Pointer"
	case "const void*", "const uint8_t*", "const uint8_t *":
		return "unsafe.Pointer"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "unsafe.Pointer"

	default:
		gdType = strings.TrimPrefix(gdType, "const")

		if strings.HasSuffix(gdType, "*") {
			return "*" + gdType[:len(gdType)-1]
		}

		gdType = strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "enum::")
		gdType = strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "bitfield::")

		gdType = strings.Replace(gdType, ".", "", -1)

		if class, ok := classDB[gdType]; ok {
			if pkg != class.Package {
				if class.Package == "internal" {
					class.Package = "gd"
				}
				return class.Package + "." + class.Name
			}
			return class.Name
		}

		if inCore(gdType) {
			return maybeInternal(gdType)
		}

		return gdType
	}
}

func mapOperator(name string) string {
	switch name {
	case "==":
		return "Equals"
	case "!=":
		return "NotEqual"
	case "<":
		return "Less"
	case "<=":
		return "LessEqual"
	case ">":
		return "Greater"
	case ">=":
		return "GreaterEqual"
	case "+":
		return "Add"
	case "-":
		return "Subtract"
	case "*":
		return "Multiply"
	case "/":
		return "Divide"
	case "unary-":
		return "Negate"
	case "%":
		return "Module"
	case "**":
		return "Power"
	case "<<":
		return "ShiftLeft"
	case ">>":
		return "ShiftRight"
	case "&":
		return "BitAnd"
	case "|":
		return "BitOr"
	case "^":
		return "BitXor"
	case "~":
		return "BitNegate"
	case "and":
		return "And"
	case "or":
		return "Or"
	case "xor":
		return "Xor"
	case "not":
		return "Not"
	case "in":
		return "In"
	default:
		panic("unknown operator: " + name)
	}
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}

	fnName = strings.ToLower(fnName)

	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, strings.Title(word))
	}
	/*if joins[0] == "Get" {
		backup := joins
		joins = joins[1:]

		if len(joins) == 0 {
			joins = backup
		} else {
			if _, err := strconv.Atoi(joins[0]); err == nil {
				joins = backup
			}
		}
	}*/
	return strings.Join(joins, "")
}

func genEnum(pkg string, decl, code io.Writer, prefix string, enum gdjson.Enum) {
	name := prefix + strings.Replace(enum.Name, ".", "", -1)
	if name == "Side" || name == "EulerOrder" {
		return
	}

	if decl != nil {
		fmt.Fprintln(decl)
		fmt.Fprintf(decl, "type %v int64\n", name)

		if code != nil {
			fmt.Fprintln(code)
		}

		var topLevelPrefix = pkg
		if topLevelPrefix == "internal" {
			topLevelPrefix = "gd"
		}
		if code != nil {
			fmt.Fprintf(code, "type %v = %v.%[1]v\n", name, topLevelPrefix)
		}

		if name == "GDExtensionInitializationLevel" || name == "VariantType" || name == "Error" {
			genEnum(pkg, nil, decl, prefix, enum)
		}
	}
	if code != nil {
		if len(enum.Values) > 0 {
			fmt.Fprintln(code)
			fmt.Fprintf(code, "const (\n")
			for _, value := range enum.Values {
				n := prefix + convertName(value.Name)
				if n == name {
					n += "Default"
				}
				if value.Description != "" {
					fmt.Fprint(code, "/*")
					fmt.Fprint(code, value.Description)
					fmt.Fprintln(code, "*/")
				}
				fmt.Fprintf(code, "\t%v %v = %v\n", n, name, value.Value)
			}
			fmt.Fprintf(code, ")\n")
		}
	}
}

type ClassDB map[string]gdjson.Class

func (db ClassDB) nameOf(pkg, original string) string {
	class := db[original]
	if pkg == class.Package {
		return class.Name
	}
	if class.Package == "internal" {
		return "gd." + class.Name
	}
	return class.Package + "." + class.Name
}

func (db ClassDB) genArgs(pkg string, method gdjson.Method) string {
	var args []string
	for _, arg := range method.Arguments {
		if arg.Name == "type" {
			arg.Name = "typ"
		}
		if arg.Name == "func" {
			arg.Name = "fn"
		}
		if arg.Name == "default" {
			arg.Name = "def"
		}
		if arg.Name == "interface" {
			arg.Name = "iface"
		}
		if arg.Name == "range" {
			arg.Name = "rng"
		}
		if arg.Name == "var" {
			arg.Name = "variable"
		}
		if arg.Name == "map" {
			arg.Name = "mapping"
		}
		if arg.Name == "internal" {
			arg.Name = "internal_"
		}
		args = append(args, fmt.Sprintf("%v %v", arg.Name, db.convertType(pkg, arg.Meta, arg.Type)))
	}
	return strings.Join(args, ", ")
}

func pascal(name string) string {
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func fixReserved(name string) string {
	switch name {
	case "bool":
		return "b"
	case "type":
		return "atype"
	case "range":
		return "arange"
	case "default":
		return "def"
	case "func":
		return "fn"
	case "frame":
		return "frame_"
	case "interface":
		return "intf"
	case "internal":
		return "internal_"
	case "map":
		return "mapping"
	case "var":
		return "v"
	case "string":
		return "s"
	case "RID":
		return "rid"
	default:
		return name
	}
}

func inCore(s string) bool {
	switch s {
	case "Object", "RefCounted":
		return true
	default:
		return false
	}
}

func generate() error {
	spec, err := LoadSpecification()
	if err != nil {
		return err
	}

	var classDB = make(ClassDB)
	for _, enum := range spec.GlobalEnums {
		classDB[strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
			IsEnum:  true,
			Name:    strings.Replace(enum.Name, ".", "", -1),
			Package: "internal",
		}
	}
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			classDB[class.Name+strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
				IsEnum:  true,
				Name:    class.Name + strings.Replace(enum.Name, ".", "", -1),
				Package: "internal",
			}
		}
	}
	for i, class := range spec.Classes {
		var pkg = "internal"
		if !inCore(class.Name) {
			pkg = "classdb"
		}
		class.Package = pkg
		for _, enum := range class.Enums {
			classDB[class.Name+strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
				IsEnum:  true,
				Name:    class.Name + strings.Replace(enum.Name, ".", "", -1),
				Package: pkg,
			}
		}
		spec.Classes[i] = class
		classDB[class.Name] = class
	}

	file, err := os.Open("./extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		file, err = os.Create("../extension_api.json")
		if err != nil {
			return err
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		file.Seek(0, 0)
		resp.Body.Close()
	}

	out, err := os.Create("./internal/out.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, `//go:build !generate`)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(out, `package gd`)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `import "grow.graphics/gd/internal/callframe"`)
	fmt.Fprintln(out)

	core, err := os.Create("./internal/all.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(core, `//go:build !generate`)
	fmt.Fprintln(core)
	fmt.Fprintln(core, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(core, `package gd`)
	fmt.Fprintln(core)
	fmt.Fprintln(core, `import "unsafe"`)
	fmt.Fprintln(core, `import "reflect"`)
	fmt.Fprintln(core, `import "grow.graphics/gd/internal/mmm"`)
	fmt.Fprintln(core, `import "grow.graphics/gd/internal/callframe"`)
	fmt.Fprintln(core)

	all, err := os.Create("./internal/classdb/all.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(all, `//go:build !generate`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(all, `package gd`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `import "unsafe"`)
	fmt.Fprintln(all, `import "reflect"`)
	fmt.Fprintln(all, `import "grow.graphics/gd/internal/mmm"`)
	//fmt.Fprintln(all, `import "grow.graphics/gd/internal/callframe"`)
	fmt.Fprintln(all, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(all)

	enums, err := os.Create("./enums.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(enums, `//go:build !generate`)
	fmt.Fprintln(enums)
	fmt.Fprintln(enums, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(enums, `package gd`)
	fmt.Fprintln(enums)
	fmt.Fprintln(enums, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(enums)

	classdb, err := os.Create("./classdb.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(classdb, `//go:build !generate`)
	fmt.Fprintln(classdb)
	fmt.Fprintln(classdb, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(classdb, `package gd`)
	fmt.Fprintln(classdb)
	fmt.Fprintln(classdb, `import "unsafe"`)
	fmt.Fprintln(classdb, `import classdb "grow.graphics/gd/internal/classdb"`)
	fmt.Fprintln(classdb)

	for _, enum := range spec.GlobalEnums {
		genEnum("internal", core, enums, "", enum)
	}

	var singletons = make(map[string]bool)
	for _, class := range spec.Singletons {
		singletons[class.Name] = true
	}

	var builtin_sizes = make(map[string]int)
	for _, config := range spec.BuiltinClassSizes {
		if config.BuildConfiguration != "float_64" {
			continue
		}
		for _, size := range config.Sizes {
			builtin_sizes[size.Name] = size.Size
		}
	}

	fmt.Fprintf(out, "type utility struct{\n")
	for _, utility := range spec.UtilityFunctions {
		fmt.Fprintf(out, "\t%v func(ret uintptr, args callframe.Args, c int32) `hash:\"%v\"`\n", utility.Name, utility.Hash)

		fmt.Fprintf(core, "\nfunc (ctx Lifetime) %v(", convertName(utility.Name))
		for i, arg := range utility.Arguments {
			if i > 0 {
				fmt.Fprintf(core, ", ")
			}
			fmt.Fprintf(core, "%v %v", fixReserved(arg.Name), classDB.convertType("internal", arg.Meta, arg.Type))
		}
		if utility.IsVararg {
			if len(utility.Arguments) > 0 {
				fmt.Fprintf(core, ", ")
			}
			fmt.Fprintf(core, "args ...Variant")
		}
		fmt.Fprintf(core, ")")
		result := classDB.convertType("internal", "", utility.ReturnType)
		ptrKind, isPtr := classDB.isPointer(result)
		if result != "" {
			fmt.Fprintf(core, " %v {\n", result)
		} else {
			fmt.Fprintf(core, " {\n")
		}
		fmt.Fprintf(core, "\tvar frame = callframe.New()\n")
		for _, arg := range utility.Arguments {
			_, ok := classDB[arg.Type]
			if ok {
				fmt.Fprintf(core, "\tcallframe.Arg(frame, mmm.Get(%v))\n", fixReserved(arg.Name))
				continue
			}
			argType := classDB.convertType("internal", arg.Meta, arg.Type)
			_, argIsPtr := classDB.isPointer(argType)
			if argIsPtr {
				fmt.Fprintf(core, "\tcallframe.Arg(frame, mmm.Get(%v))\n", fixReserved(arg.Name))
			} else {
				fmt.Fprintf(core, "\tcallframe.Arg(frame, %v)\n", fixReserved(arg.Name))
			}
		}
		if utility.IsVararg {
			fmt.Fprintf(core, "\tfor _, arg := range args {\n")
			fmt.Fprintf(core, "\t\tcallframe.Arg(frame, mmm.Get(arg))\n")
			fmt.Fprintf(core, "\t}\n")
		}
		if isPtr {
			fmt.Fprintf(core, "\tvar r_ret = callframe.Ret[%v](frame)\n", ptrKind)
		} else {
			if result != "" {
				fmt.Fprintf(core, "\tvar r_ret = callframe.Ret[%v](frame)\n", result)
			} else {
				fmt.Fprintf(core, "\tvar r_ret callframe.Nil\n")
			}
		}
		fmt.Fprintf(core, "\tctx.API.utility.%v(r_ret.Uintptr(), frame.Array(0), %d)\n", utility.Name, len(utility.Arguments))
		if isPtr {
			_, ok := classDB[result]
			if ok || result == "Object" {
				fmt.Fprintf(core, "\tvar ret %v\n", result)
				fmt.Fprintf(core, "\tret.SetPointer(PointerMustAssertInstanceID(ctx, r_ret.Get()))\n")
			} else {
				fmt.Fprintf(core, "\tvar ret = mmm.New[%v](ctx.Lifetime, ctx.API, r_ret.Get())\n", result)
			}
		} else {
			if result != "" {
				fmt.Fprintf(core, "\tvar ret = r_ret.Get()\n")
			}
		}
		fmt.Fprintf(core, "\tframe.Free()\n")
		if result != "" {
			fmt.Fprintf(core, "\treturn ret\n")
		}
		fmt.Fprintf(core, "}\n")
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type builtin struct{\n")
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			genEnum("internal", core, enums, class.Name, enum)
		}
		fmt.Fprintf(out, "\t%v struct{\n", class.Name)
		for _, method := range class.Methods {
			if method.Name == "map" {
				method.Name = "map_"
			}
			fmt.Fprintf(out, "\t\t%v func(base uintptr, args callframe.Args, ret uintptr, c int32) `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type typeset struct{\n")
	fmt.Fprintf(out, "\tcreation struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v [%d]func(uintptr, callframe.Args)\n", class.Name, len(class.Constructors))
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "\toperator struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v struct {\n", class.Name)
		sort.Slice(class.Operators, func(i, j int) bool {
			return class.Operators[i].Name < class.Operators[j].Name
		})
		var last string
		for _, op := range class.Operators {
			if op.Name == "unary+" || op.RightType == "Variant" {
				continue
			}
			if strings.HasPrefix(op.Name, "unary") || op.RightType == "" {
				if last != "" {
					fmt.Fprintf(out, "\t\t\t}\n")
				}
				fmt.Fprintf(out, "\t\t\t%v func(a, b, ret uintptr)\n", mapOperator(op.Name))
				last = ""
			} else {
				if op.Name != last {
					if last != "" {
						fmt.Fprintf(out, "\t\t\t}\n")
					}
					fmt.Fprintf(out, "\t\t\t%v struct {\n", mapOperator(op.Name))
				}
				fmt.Fprintf(out, "\t\t\t\t%v func(a, b, ret uintptr)\n", op.RightType)
				last = op.Name
			}
		}
		if last != "" {
			fmt.Fprintf(out, "\t\t\t}\n")
		}
		fmt.Fprintf(out, "\t\t}\n")
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "\tdestruct struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v func(uintptr)\n", class.Name)
		for _, method := range class.Methods {
			classDB.methodCall(core, "internal", class, method, callBuiltin)
		}
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type singletons struct{\n")
	for _, class := range spec.Classes {
		if singletons[class.Name] {
			fmt.Fprintf(out, "\t%v StringName\n", class.Name)
		}
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type methods struct{\n")
	for _, class := range spec.Classes {
		if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
			continue
		}
		fmt.Fprintf(out, "\t%v struct{\n", class.Name)
		for _, method := range class.Methods {
			if method.Name == "select" {
				method.Name = "select_"
			}
			if method.IsVirtual {
				continue
			}
			fmt.Fprintf(out, "\t\tBind_%v MethodBind `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	for _, class := range spec.Classes {
		var w = core
		if class.Package != "internal" {
			w = all
		}
		pkg := class.Package

		prefix := ""
		if pkg != "internal" {
			prefix = "gd."
		}

		for _, enum := range class.Enums {
			genEnum(pkg, w, nil, class.Name, enum)
		}
		if class.Name != "Object" {
			if singletons[class.Name] {
				fmt.Fprintf(classdb, "func %v(godot Lifetime) [1]%[2]v { obj := godot.API.Object.GetSingleton(godot, godot.API.Singletons.%[1]v); return *(*[1]%[2]v)(unsafe.Pointer(&obj)) }\n",
					class.Name, classDB.nameOf("gd", class.Name))
			}
			fmt.Fprintf(w, "type %[1]v struct {_ [0]*%[1]v; ptr "+prefix+"Pointer}\n", class.Name, classDB.nameOf(pkg, class.Inherits))
			fmt.Fprintf(w, "\n//go:nosplit\nfunc (self %[1]v) AsPointer() "+prefix+"Pointer { return self.ptr }\n", class.Name)
			fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self *%[1]v) SetPointer(ptr "+prefix+"Pointer) { self.ptr = ptr }\n", class.Name)
		}
		if class.Inherits != "" {
			var i = 1
			super := classDB[class.Inherits]
			fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) Super() [1]%[3]v { return *(*[1]%[3]v)(unsafe.Pointer(&self)) }\n", class.Name, super.Name, classDB.nameOf(pkg, super.Name))
			fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) AsObject() "+prefix+"Object { return *(*"+prefix+"Object)(unsafe.Pointer(&self)) }\n", class.Name)
			fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) As%[1]v() [1]%[1]v { return [1]%[1]v{self} }\n", class.Name)
			for super.Name != "" && super.Name != "Object" {
				if super.Name == "RefCounted" {
					fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) AsRefCounted() "+prefix+"RefCounted { return *(*"+prefix+"RefCounted)(unsafe.Pointer(&self)) }\n", class.Name)
				} else {
					fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) As%[2]v() [1]%[3]v { return *(*[1]%[3]v)(unsafe.Pointer(&self)) }\n", class.Name, super.Name, classDB.nameOf(pkg, super.Name))
				}
				i++
				super = classDB[super.Inherits]
			}
		}
		if singletons[class.Name] {
			fmt.Fprintf(w, "\n//go:nosplit\n\nfunc (self %[1]v) IsSingleton() {}\n", class.Name)
		}
		if class.Name == "Object" || class.Name == "RefCounted" {
			for _, method := range class.Methods {
				classDB.methodCall(w, pkg, class, method, callDefault)
			}
		} else {
			for _, method := range class.Methods {
				if method.IsVirtual {
					classDB.methodCall(w, pkg, class, method, callDefault)
				}
			}
		}
		fmt.Fprintf(w, "\nfunc (self %[1]v) Virtual(name string) reflect.Value {\n", class.Name)
		fmt.Fprintf(w, "\tswitch name {\n")
		for _, method := range class.Methods {
			if method.IsVirtual {
				fmt.Fprintf(w, "\tcase \"%v\": return reflect.ValueOf(self.%v);\n", method.Name, method.Name)
			}
		}
		if class.Inherits != "" {
			fmt.Fprintf(w, "\tdefault: return "+prefix+"VirtualByName(self.Super()[0], name)\n")
		} else {
			fmt.Fprintf(w, "\tdefault: return reflect.Value{}\n")
		}
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "}\n")
	}

	return nil
}

func (db ClassDB) isPointer(t string) (string, bool) {
	t = strings.TrimPrefix(t, "gd.")
	if strings.HasPrefix(t, "ArrayOf") {
		return "uintptr", true
	}
	switch t {
	case "String", "StringName", "NodePath",
		"Dictionary", "Array":
		return "uintptr", true
	case "Callable", "Signal",
		"PackedByteArray", "PackedInt32Array",
		"PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedStringArray",
		"PackedVector2Array", "PackedVector3Array",
		"PackedVector4Array",
		"PackedColorArray":
		return "[2]uintptr", true
	case "Variant":
		return "[3]uintptr", true
	default:
		if entry, ok := db[t]; ok && !entry.IsEnum {
			return "uintptr", true
		}
		return "", false
	}
}

type callType int

const (
	callDefault callType = iota
	callBuiltin
	callUtility
)

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
		fmt.Fprintf(w, "func (%s) %s(impl func(ptr unsafe.Pointer", class.Name, method.Name)
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
				fmt.Fprintf(w, "\t\t%v.SetPointer(mmm.Let["+prefix+"Pointer](ctx.Lifetime, ctx.API, [2]uintptr{"+prefix+"UnsafeGet[uintptr](p_args,%d)}))\n", fixReserved(arg.Name), i)
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
				_, ok := classDB[result]
				if ok || result == "gd.Object" {
					ret = fmt.Sprintf("mmm.End(%s.AsPointer())", ret)
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
		fmt.Fprintf(w, "//go:nosplit\nfunc (self *%v) %v(%s", class.Name, convertName(method.Name), context)
	} else {
		fmt.Fprintf(w, "//go:nosplit\nfunc (self %v) %v(%s", class.Name, convertName(method.Name), context)
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
			fmt.Fprintf(w, "\tvar selfPtr = self.AsPointer()\n")
		}
	}
	fmt.Fprintf(w, "\tvar frame = callframe.New()\n")
	for _, arg := range method.Arguments {
		_, ok := classDB[arg.Type]
		if ok {
			switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name][arg.Name]; semantics {
			case gdjson.OwnershipTransferred, gdjson.LifetimeBoundToClass:
				fmt.Fprintf(w, "\tcallframe.Arg(frame, mmm.End(%v.AsPointer())[0])\n", fixReserved(arg.Name))
			case gdjson.RefCountedManagement, gdjson.IsTemporaryReference, gdjson.MustAssertInstanceID, gdjson.ReversesTheOwnership:
				fmt.Fprintf(w, "\tcallframe.Arg(frame, mmm.Get(%v.AsPointer())[0])\n", fixReserved(arg.Name))
			default:
				panic("unknown ownership: " + fmt.Sprint(semantics))
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
		_, ok := classDB[result]
		if ok || result == "gd.Object" {
			fmt.Fprintf(w, "\tvar ret %v\n", result)

			switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name]["return value"]; semantics {
			case gdjson.RefCountedManagement, gdjson.OwnershipTransferred:
				fmt.Fprintf(w, "\tret.SetPointer("+prefix+"PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))\n")
			case gdjson.LifetimeBoundToClass:
				fmt.Fprintf(w, "\tret.SetPointer("+prefix+"PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))\n")
			case gdjson.MustAssertInstanceID:
				fmt.Fprintf(w, "\tret.SetPointer("+prefix+"PointerMustAssertInstanceID(ctx, r_ret.Get()))\n")
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

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
