package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"graphics.gd/internal/gdjson"
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
	fmt.Fprintln(out, `import "graphics.gd/internal/gdextension"`)
	fmt.Fprintln(out)

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

	fmt.Fprintf(out, "type typeset struct{\n")
	fmt.Fprintf(out, "\tcreation struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v [%d]gdextension.FunctionID\n", class.Name, len(class.Constructors))
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
			fmt.Fprintf(out, "\t\tBind_%v gdextension.MethodForClass `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	for _, class := range spec.Classes {
		if class.IsEphemeral {
			continue
		}
		var w = all
		pkg := class.Package
		if class.Name != "Object" && class.Name != "RefCounted" {
			fmt.Fprintf(w, "type %[1]v pointers.Trio[%[1]v]\n", class.Name, classDB.nameOf(pkg, class.Inherits))
			fmt.Fprintf(w, "func (self %[1]v) Free() { (*(*Object)(unsafe.Pointer(&self))).Free() }\n", class.Name)
			fmt.Fprintf(w, "func (self %[1]v) IsAlive(raw [3]uint64) bool { return (*(*Object)(unsafe.Pointer(&self))).IsAlive(raw) }\n", class.Name)

			if class.Inherits != "" {
				fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) AsObject() [1]Object { return (*(*[1]Object)(unsafe.Pointer(&self))) }\n", class.Name)
			}
			fmt.Fprintf(w, "\nfunc (self %[1]v) Virtual(name string) reflect.Value { return reflect.Value{} }\n", class.Name)
		}
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

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
