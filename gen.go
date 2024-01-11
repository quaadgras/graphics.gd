//go:build generate

//go:generate go run -tags generate .
//go:generate go fmt ./

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Enum struct {
	Name   string `json:"name"`
	Values []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"values"`
}

type Method struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsConst     bool   `json:"is_const"`
	IsVararg    bool   `json:"is_vararg"`
	IsStatic    bool   `json:"is_static"`
	IsVirtual   bool   `json:"is_virtual"`
	Hash        int64  `json:"hash"`
	ReturnType  string `json:"return_type"` // builtin class
	ReturnValue struct {
		Type string `json:"type"`
		Meta string `json:"meta"`
	} `json:"return_value,omitempty"`
	Arguments []struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Meta         string `json:"meta"`
		DefaultValue string `json:"default_value,omitempty"`
	} `json:"arguments,omitempty"`
}

type Class struct {
	Name    string `json:"name"`
	Rename  string `json:"rename,omitempty"`
	Package string `json:"package,omitempty"`
	IsEnum  bool   `json:"is_enum"`

	Description    string `json:"description"`
	IsRefcounted   bool   `json:"is_refcounted"`
	IsInstantiable bool   `json:"is_instantiable"`

	IsKeyed   bool `json:"is_keyed"` // builtin class
	Operators []struct {
		Name       string `json:"name"`
		RightType  string `json:"right_type,omitempty"`
		ReturnType string `json:"return_type"`
	} `json:"operators"` // builtin class
	Constructors []struct {
		Index     int `json:"index"`
		Arguments []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"constructors"` // builtin class
	HasDestructor      bool   `json:"has_destructor"`                 // builtin class
	IndexingReturnType string `json:"indexing_return_type,omitempty"` // builtin class

	Inherits string   `json:"inherits,omitempty"`
	APIType  string   `json:"api_type"`
	Enums    []Enum   `json:"enums,omitempty"`
	Methods  []Method `json:"methods,omitempty"`

	Members []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"members,omitempty"` // builtin class

	Signals []struct {
		Name      string `json:"name"`
		Arguments []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"signals,omitempty"`
	Properties []struct {
		Type   string `json:"type"`
		Name   string `json:"name"`
		Setter string `json:"setter"`
		Getter string `json:"getter"`
		Index  int    `json:"index"`
	} `json:"properties,omitempty"`
	Constants []struct {
		Name  string `json:"name"`
		Type  string `json:"type"` // builtin class
		Value any    `json:"value"`
	} `json:"constants,omitempty"`

	singleton, builtin bool // internal
}

/*
Specification of the Godot Extension API.
Created with https://mholt.github.io/json-to-go/
*/
type Specification struct {
	Header struct {
		VersionMajor    int    `json:"version_major"`
		VersionMinor    int    `json:"version_minor"`
		VersionPatch    int    `json:"version_patch"`
		VersionStatus   string `json:"version_status"`
		VersionBuild    string `json:"version_build"`
		VersionFullName string `json:"version_full_name"`
	} `json:"header"`
	BuiltinClassSizes []struct {
		BuildConfiguration string `json:"build_configuration"`
		Sizes              []struct {
			Name string `json:"name"`
			Size int    `json:"size"`
		} `json:"sizes"`
	} `json:"builtin_class_sizes"`
	BuiltinClassMemberOffsets []struct {
		BuildConfiguration string `json:"build_configuration"`
		Classes            []struct {
			Name    string `json:"name"`
			Members []struct {
				Member string `json:"member"`
				Offset int    `json:"offset"`
			} `json:"members"`
		} `json:"classes"`
	} `json:"builtin_class_member_offsets"`
	GlobalConstants  []interface{} `json:"global_constants"`
	GlobalEnums      []Enum        `json:"global_enums"`
	UtilityFunctions []struct {
		Name       string `json:"name"`
		ReturnType string `json:"return_type,omitempty"`
		Category   string `json:"category"`
		IsVararg   bool   `json:"is_vararg"`
		Hash       int    `json:"hash"`
		Arguments  []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"utility_functions"`
	BuiltinClasses []Class `json:"builtin_classes"`
	Classes        []Class `json:"classes"`
	Singletons     []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"singletons"`
	NativeStructures []struct {
		Name   string `json:"name"`
		Format string `json:"format"`
	} `json:"native_structures"`
}

// Load the specification, either from a local file or by downloading
// it from the Godot Github repository.
func (spec *Specification) Load() error {
	file, err := os.Open("./api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		file, err = os.Create("./api.json")
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

	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return err
	}

	return nil
}

func hasNative(gdType string) bool {
	switch gdType {
	case "int", "float", "String", "bool",
		"PackedStringArray", "PackedInt32Array", "PackedInt64Array",
		"PackedFloat32Array", "PackedFloat64Array", "PackedVector2Array",
		"PackedVector3Array", "PackedColorArray", "Variant":
		return true
	}
	return false
}

func convertType(pkg, meta string, gdType string) string {
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		return "ArrayOf[" + convertType(pkg, meta, gdType) + "]"
	}
	switch gdType {
	case "int":
		return "int64"
	case "float":
		return "Float"
	case "bool":
		return "bool"
	case "String", "StringName":
		return "string"

	case "enum::GDExtension.InitializationLevel":
		return "ExtensionInitializationLevel"
	case "enum::GDExtensionManager.LoadStatus":
		return "ExtensionManagerLoadStatus"

	case "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedColorArray", "PackedByteArray":
		return gdType
	case "Variant":
		return "any"

	case "enum::Variant.Type":
		return "VariantType"

	// strange C++ cases
	case "enum::Error":
		return "int64"
	case "const uint8_t **":
		return "*[]byte"
	case "const void*", "const uint8_t*", "const uint8_t *":
		return "[]byte"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "[]byte"

	default:
		gdType = strings.TrimPrefix(gdType, "const")

		if strings.HasSuffix(gdType, "*") {
			return "*" + gdType[:len(gdType)-1]
		}

		gdType = strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "enum::")
		gdType = strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "bitfield::")

		return strings.Replace(gdType, ".", "", -1)
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

func genEnum(pkg string, code io.Writer, prefix string, enum Enum) {
	name := prefix + strings.Replace(enum.Name, ".", "", -1)
	name = strings.TrimPrefix(name, strings.Title(path.Base(pkg)))
	name = strings.TrimPrefix(name, strings.ToUpper(path.Base(pkg)))

	fmt.Fprintln(code)
	fmt.Fprintf(code, "type %v int64\n", name)
	if len(enum.Values) > 0 {
		fmt.Fprintln(code)
		fmt.Fprintf(code, "const (\n")
		for _, value := range enum.Values {
			n := prefix + convertName(value.Name)
			if n == name {
				n += "Default"
			}

			n = strings.TrimPrefix(n, strings.Title(path.Base(pkg)))
			n = strings.TrimPrefix(n, strings.ToUpper(path.Base(pkg)))
			n = strings.Replace(n, "XR", "", -1)
			n = strings.Replace(n, "xr", "", -1)
			n = strings.Replace(n, "Xr", "", -1)

			fmt.Fprintf(code, "\t%v %v = %v\n", n, name, value.Value)
		}
		fmt.Fprintf(code, ")\n")
	}
}

type ClassDB map[string]Class

func (db ClassDB) nameOf(pkg, original string) string {
	class := db[original]
	if class.IsEnum {
		return "enums." + class.Name
	}
	if pkg == "class" || pkg == "gd" {
		return class.Name
	}
	return "class." + class.Name
}

func (db ClassDB) genArgs(pkg string, method Method) string {
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
		args = append(args, fmt.Sprintf("%v %v", arg.Name, convertType(pkg, arg.Meta, arg.Type)))
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
	case "type":
		return "atype"
	case "range":
		return "arange"
	case "default":
		return "def"
	case "func":
		return "fn"
	case "interface":
		return "intf"
	case "map":
		return "mapping"
	case "var":
		return "v"
	case "string":
		return "s"
	default:
		return name
	}
}

func generate() error {
	var spec Specification
	if err := spec.Load(); err != nil {
		return err
	}

	file, err := os.Open("./api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		file, err = os.Create("./api.json")
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

	api, err := os.Create("./api.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(api, `//go:build !generate`)
	fmt.Fprintln(api)
	fmt.Fprintln(api, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(api, `package gd`)
	fmt.Fprintln(api)

	all, err := os.Create("./all.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(all, `//go:build !generate`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(all, `package gd`)

	val, err := os.Create("./val.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(val, `//go:build !generate`)
	fmt.Fprintln(val)
	fmt.Fprintln(val, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(val, `package gd`)
	fmt.Fprintln(val)

	pub, err := os.Create("./pub.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(pub, `//go:build !generate`)
	fmt.Fprintln(pub)
	fmt.Fprintln(pub, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(pub, `package gd`)
	fmt.Fprintln(pub)

	for _, enum := range spec.GlobalEnums {
		genEnum("gd", val, "", enum)
	}

	fmt.Fprintf(api, "type API struct {\n")
	for _, class := range spec.Singletons {
		fmt.Fprintf(api, "\t%v %[1]v\n", class.Name)
	}
	for _, utility := range spec.UtilityFunctions {
		fmt.Fprintf(api, "\tUtility_%v func(", utility.Name)
		for i, arg := range utility.Arguments {
			fmt.Fprintf(api, "%v %v", fixReserved(arg.Name), convertType("", arg.Meta, arg.Type))
			if i < len(utility.Arguments)-1 {
				fmt.Fprint(api, ", ")
			}
		}
		result := convertType("gd", "", utility.ReturnType)
		fmt.Fprintf(api, ") %v\t`hash:\"%v\"`\n", result, utility.Hash)
	}
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			genEnum("gd", val, class.Name, enum)
		}
		class.Name = strings.Title(class.Name)

		for _, method := range class.Methods {
			result := convertType("gd", "", method.ReturnType)

			fmt.Fprintf(api, "\t%v_%v func(self %[1]v", class.Name, method.Name)
			for _, arg := range method.Arguments {
				fmt.Fprint(api, ", ")
				fmt.Fprintf(api, "%v %v", fixReserved(arg.Name), convertType("gd", arg.Meta, arg.Type))
			}
			fmt.Fprintf(api, ") %v\t`hash:\"%v\"`\n", result, method.Hash)
		}
	}
	var classDB = make(map[string]Class)
	for _, class := range spec.Classes {
		classDB[class.Name] = class
	}
	for _, class := range spec.Classes {
		for _, enum := range class.Enums {
			genEnum("gd", val, class.Name, enum)
		}
		fmt.Fprintf(pub, "type %v struct { methods%v }\n", class.Name, class.Name)
		fmt.Fprintf(all, "type methods%[1]v struct{_ [0]*methods%[1]v; class }\n", class.Name)
		if class.Inherits != "" {
			fmt.Fprintf(all, "\nfunc (self %[1]v) %[2]v() %[2]v { var parent %[2]v; parent.class = self.class; return parent }\n", class.Name, class.Inherits)
		}
		var methodCount = 0
		var virtualCount = 0
		for _, method := range class.Methods {
			if !method.IsVirtual {
				methodCount++
			} else {
				virtualCount++
			}
		}
		var i int
		for _, method := range class.Methods {
			if method.IsVirtual {
				continue
			}

			result := convertType("gd", method.ReturnValue.Meta, method.ReturnValue.Type)

			fmt.Fprintf(api, "\t%v_%v func(self %[1]v", class.Name, method.Name)
			for _, arg := range method.Arguments {
				fmt.Fprint(api, ", ")
				fmt.Fprintf(api, "%v %v", fixReserved(arg.Name), convertType("gd", arg.Meta, arg.Type))
			}
			fmt.Fprintf(api, ") %v\t`hash:\"%v\"`\n", result, method.Hash)

			fmt.Fprintf(all, "\tfunc (self methods%v) %v(", class.Name, convertName(method.Name))
			for i, arg := range method.Arguments {
				fmt.Fprintf(all, "%v %v", fixReserved(arg.Name), convertType("gd", arg.Meta, arg.Type))
				if i < len(method.Arguments)-1 {
					fmt.Fprint(all, ", ")
				}
			}
			fmt.Fprintf(all, ") %v {\n", result)
			if result != "" {
				fmt.Fprintf(all, "\t\treturn ")
			} else {
				fmt.Fprintf(all, "\t\t")
			}
			fmt.Fprintf(all, "self.Runtime.%v_%v(%[1]v{self}", class.Name, method.Name)
			for _, arg := range method.Arguments {
				fmt.Fprint(all, ", ")
				fmt.Fprintf(all, "%v", fixReserved(arg.Name))
			}
			fmt.Fprintf(all, ")\n\t}\n")

			i++
		}
	}
	fmt.Fprintf(api, "}\n")

	return nil
}

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
