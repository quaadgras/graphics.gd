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

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
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
		return maybeInternal("Array")
	}
	switch gdType {
	case "int", "Int":
		return "int64"
	case "float", "Float":
		return "float64"
	case "bool", "Bool":
		return "bool"
	case "String":
		return maybeInternal("String")
	case "StringName":
		return maybeInternal("StringName")
	case "enum::GDExtension.InitializationLevel":
		return maybeInternal("GDExtensionInitializationLevel")
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
		if class.IsEphemeral {
			continue
		}
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
	fmt.Fprintln(out, `import "graphics.gd/internal/callframe"`)
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
	fmt.Fprintln(core, `import "reflect"`)
	fmt.Fprintln(core, `import "unsafe"`)
	fmt.Fprintln(core, `import "graphics.gd/internal/pointers"`)
	fmt.Fprintln(core, `import "graphics.gd/internal/callframe"`)
	fmt.Fprintln(core)

	all, err := os.Create("./internal/gdclass/all.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(all, `//go:build !generate`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(all, `package gdclass`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `import "reflect"`)
	fmt.Fprintln(all, `import "unsafe"`)
	fmt.Fprintln(all, `import "graphics.gd/internal/pointers"`)
	fmt.Fprintln(all)

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
		fmt.Fprintf(out, "\t%v func(ret callframe.Addr, args callframe.Args, c int32) `hash:\"%v\"`\n", utility.Name, utility.Hash)

		fmt.Fprintf(core, "\nfunc %v(", convertName(utility.Name))
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
				fmt.Fprintf(core, "\tcallframe.Arg(frame, pointers.Get(%v))\n", fixReserved(arg.Name))
				continue
			}
			argType := classDB.convertType("internal", arg.Meta, arg.Type)
			_, argIsPtr := classDB.isPointer(argType)
			if argIsPtr {
				fmt.Fprintf(core, "\tcallframe.Arg(frame, pointers.Get(%v))\n", fixReserved(arg.Name))
			} else {
				fmt.Fprintf(core, "\tcallframe.Arg(frame, %v)\n", fixReserved(arg.Name))
			}
		}
		if utility.IsVararg {
			fmt.Fprintf(core, "\tfor _, arg := range args {\n")
			fmt.Fprintf(core, "\t\tcallframe.Arg(frame, pointers.Get(arg))\n")
			fmt.Fprintf(core, "\t}\n")
		}
		if isPtr {
			fmt.Fprintf(core, "\tvar r_ret = callframe.Ret[%v](frame)\n", ptrKind)
		} else {
			if result != "" {
				fmt.Fprintf(core, "\tvar r_ret = callframe.Ret[%v](frame)\n", result)
			} else {
				fmt.Fprintf(core, "\tvar r_ret = callframe.Nil\n")
			}
		}
		fmt.Fprintf(core, "\tGlobal.utility.%v(r_ret.Addr(), frame.Array(0), %d)\n", utility.Name, len(utility.Arguments))
		if isPtr {
			_, ok := classDB[result]
			if ok || result == "Object" {
				fmt.Fprintf(core, "\tvar ret %v = PointerMustAssertInstanceID[Object](r_ret.Get())\n", result)
			} else {
				fmt.Fprintf(core, "\tvar ret = pointers.New[%v](r_ret.Get())\n", result)
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
		fmt.Fprintf(out, "\t%v struct{\n", class.Name)
		for _, method := range class.Methods {
			if method.Name == "map" {
				method.Name = "map_"
			}
			fmt.Fprintf(out, "\t\t%v func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32) `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type typeset struct{\n")
	fmt.Fprintf(out, "\tcreation struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v [%d]func(callframe.Addr, callframe.Args)\n", class.Name, len(class.Constructors))
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
				fmt.Fprintf(out, "\t\t\t%v func(a, b, ret callframe.Addr)\n", mapOperator(op.Name))
				last = ""
			} else {
				if op.Name != last {
					if last != "" {
						fmt.Fprintf(out, "\t\t\t}\n")
					}
					fmt.Fprintf(out, "\t\t\t%v struct {\n", mapOperator(op.Name))
				}
				fmt.Fprintf(out, "\t\t\t\t%v func(a, b, ret callframe.Addr)\n", op.RightType)
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
		fmt.Fprintf(out, "\t\t%v func(callframe.Addr)\n", class.Name)
		for _, method := range class.Methods {
			classDB.methodCall(core, "internal", class, method, callBuiltin)
		}
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type singletons struct{\n")
	for _, class := range spec.Classes {
		if class.IsEphemeral {
			continue
		}
		if singletons[class.Name] {
			fmt.Fprintf(out, "\t%v StringName\n", class.Name)
		}
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type methods struct{\n")
	for _, class := range spec.Classes {
		if class.IsEphemeral {
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
		if class.IsEphemeral {
			continue
		}
		var w = core
		if class.Package != "internal" {
			w = all
		}
		pkg := class.Package

		if class.Name != "Object" && class.Name != "RefCounted" {
			fmt.Fprintf(w, "type %[1]v pointers.Trio[%[1]v]\n", class.Name, classDB.nameOf(pkg, class.Inherits))
			fmt.Fprintf(w, "func (self %[1]v) Free() { (*(*Object)(unsafe.Pointer(&self))).Free() }\n", class.Name)
			fmt.Fprintf(w, "func (self %[1]v) IsAlive(raw [3]uint64) bool { return (*(*Object)(unsafe.Pointer(&self))).IsAlive(raw) }\n", class.Name)
		}
		if class.Inherits != "" {
			fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) AsObject() [1]Object { return (*(*[1]Object)(unsafe.Pointer(&self))) }\n", class.Name)
		}
		if class.Name == "Object" || class.Name == "RefCounted" {
			for _, method := range class.Methods {
				classDB.methodCall(w, pkg, class, method, callDefault)
			}
		}
		fmt.Fprintf(w, "\nfunc (self %[1]v) Virtual(name string) reflect.Value { return reflect.Value{} }\n", class.Name)
	}
	return nil
}

func (db ClassDB) isPointer(t string) (string, bool) {
	t = strings.TrimPrefix(t, "gd.")
	if strings.HasPrefix(t, "ArrayOf") {
		return "enginePointer", true
	}
	switch t {
	case "String", "StringName", "NodePath",
		"Dictionary", "Array":
		return "[1]enginePointer", true
	case "Signal":
		return "[2]uint64", true
	case "Callable":
		return "[2]uint64", true
	case "PackedByteArray", "PackedInt32Array",
		"PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedStringArray",
		"PackedVector2Array", "PackedVector3Array",
		"PackedVector4Array",
		"PackedColorArray":
		return "packedPointers", true
	case "Variant":
		return "[3]uint64", true
	case "Object":
		return "enginePointer", true
	default:
		if entry, ok := db[t]; ok && !entry.IsEnum {
			return "enginePointer", true
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
		fmt.Fprint(w, "\treturn func(class "+prefix+"ExtensionClass, p_args "+prefix+"Address, p_back "+prefix+"Address) {\n")
		for i, arg := range method.Arguments {
			var argType = classDB.convertType(pkg, arg.Meta, arg.Type)

			_, ok := classDB[arg.Type]
			if ok {
				fmt.Fprintf(w, "\t\tvar %v %v\n", fixReserved(arg.Name), argType)
				fmt.Fprintf(w, "\t\t%v.SetPointer(mmm.Let["+prefix+"Pointer]([2]uintptr{"+prefix+"UnsafeGet[uintptr](p_args,%d)}))\n", fixReserved(arg.Name), i)
				continue
			}

			argPtrKind, argIsPtr := classDB.isPointer(argType)
			if argIsPtr {
				fmt.Fprintf(w, "\t\tvar %v = %v\n", fixReserved(arg.Name),
					fmt.Sprintf(prefix+"UnsafeGet[%v](p_args,%d)", argPtrKind, i))
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
					ret = fmt.Sprintf("pointers.End(%s.AsPointer())", ret)
				} else {
					ret = fmt.Sprintf("pointers.End(%s)", ret)
				}
			}
			fmt.Fprintf(w, "\t\t"+prefix+"UnsafeSet(p_back, %s)\n", ret)
		}
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "}\n")
		return
	}
	if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
		return
	}
	fmt.Fprintf(w, "//go:nosplit\nfunc (self %v) %v(", class.Name, convertName(method.Name))
	if method.Name == "select" {
		method.Name = "select_"
	}
	if method.Name == "map" {
		method.Name = "map_"
	}

	for i, arg := range method.Arguments {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(pkg, arg.Meta, arg.Type))
	}
	if method.IsVararg {
		if len(method.Arguments) > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprint(w, "args ..."+prefix+"Variant")
	}
	fmt.Fprintf(w, ") %v {\n", result)
	fmt.Fprintf(w, "\tvar frame = callframe.New()\n")
	for _, arg := range method.Arguments {
		_, ok := classDB[arg.Type]
		if ok {
			switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name][arg.Name]; semantics {
			case gdjson.OwnershipTransferred, gdjson.LifetimeBoundToClass:
				fmt.Fprintf(w, "\tcallframe.Arg(frame, pointers.End(%v.AsPointer())[0])\n", fixReserved(arg.Name))
			case gdjson.RefCountedManagement, gdjson.IsTemporaryReference, gdjson.MustAssertInstanceID, gdjson.ReversesTheOwnership:
				fmt.Fprintf(w, "\tcallframe.Arg(frame, pointers.Get(%v.AsPointer())[0])\n", fixReserved(arg.Name))
			default:
				panic("unknown ownership: " + fmt.Sprint(semantics))
			}
			continue
		}

		argType := classDB.convertType(pkg, arg.Meta, arg.Type)
		_, argIsPtr := classDB.isPointer(argType)
		if argIsPtr {
			fmt.Fprintf(w, "\tcallframe.Arg(frame, pointers.Get(%v))\n", fixReserved(arg.Name))
		} else {
			fmt.Fprintf(w, "\tcallframe.Arg(frame, %v)\n", fixReserved(arg.Name))
		}
	}
	if method.IsVararg {
		fmt.Fprintf(w, "\tfor _, arg := range args {\n")
		fmt.Fprintf(w, "\t\tcallframe.Arg(frame, pointers.Get(arg))\n")
		fmt.Fprintf(w, "\t}\n")
	}
	if isPtr {
		fmt.Fprintf(w, "\tvar r_ret = callframe.Ret[%v](frame)\n", ptrKind)
	} else {
		if result != "" {
			fmt.Fprintf(w, "\tvar r_ret = callframe.Ret[%v](frame)\n", result)
		} else {
			fmt.Fprintf(w, "\tvar r_ret = callframe.Nil\n")
		}
	}
	if ctype == callBuiltin {
		var self = "callframe.Nil"
		if !method.IsStatic {
			fmt.Fprintf(w, "\tvar p_self = callframe.Arg(frame, pointers.Get(self))\n")
			self = "p_self.Addr()"
		}
		if method.IsVararg {
			fmt.Fprintf(w, "\tGlobal.builtin.%v.%v(%s, frame.Array(0), r_ret.Addr(), int32(len(args))+%d)\n", class.Name, method.Name, self, len(method.Arguments))
		} else {
			fmt.Fprintf(w, "\tGlobal.builtin.%v.%v(%s, frame.Array(0), r_ret.Addr(), %d)\n", class.Name, method.Name, self, len(method.Arguments))
		}
		if strings.HasPrefix(class.Name, "Packed") {
			fmt.Fprintf(w, "\tpointers.Set(self, p_self.Get())\n")
		}
	} else {
		if method.IsVararg {
			fmt.Fprintf(w, "\tif len(args) > 0 { panic(`varargs not supported for class methods yet`); }\n")
		}
		fmt.Fprintf(w, "\tGlobal.Object.MethodBindPointerCall(Global.Methods.%v.Bind_%v, self.AsObject(), frame.Array(0), r_ret.Addr())\n", class.Name, method.Name)
	}

	if isPtr {
		_, ok := classDB[result]
		if ok || result == "gd.Object" {
			switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name]["return value"]; semantics {
			case gdjson.RefCountedManagement, gdjson.OwnershipTransferred:
				fmt.Fprintf(w, "\tvar ret %v = "+prefix+"PointerWithOwnershipTransferredToGo[Object](r_ret.Get())\n", result)
			case gdjson.LifetimeBoundToClass:
				fmt.Fprintf(w, "\tvar ret %v ="+prefix+"PointerLifetimeBoundTo[Object](self.AsObject(), r_ret.Get())\n", result)
			case gdjson.MustAssertInstanceID:
				fmt.Fprintf(w, "\tvar ret %v ="+prefix+"PointerMustAssertInstanceID[Object](r_ret.Get())\n", result)
			default:
				panic("unknown ownership: " + fmt.Sprint(semantics))
			}

		} else {
			if strings.HasPrefix(result, "ArrayOf") {
				fmt.Fprint(w, "\tvar ret = pointers.New[Array](r_ret.Get())\n")
			} else if strings.HasPrefix(result, "gd.ArrayOf") {
				fmt.Fprint(w, "\tvar ret = pointers.New[gd.Array](r_ret.Get())\n")
			} else {
				fmt.Fprintf(w, "\tvar ret = pointers.New[%v](r_ret.Get())\n", result)
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
