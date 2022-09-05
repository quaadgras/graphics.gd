//go:build generate

package main

//go:generate go run -tags generate .

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	IsConst     bool   `json:"is_const"`
	IsVararg    bool   `json:"is_vararg"`
	IsStatic    bool   `json:"is_static"`
	IsVirtual   bool   `json:"is_virtual"`
	Hash        int64  `json:"hash"`
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
	Name           string   `json:"name"`
	IsRefcounted   bool     `json:"is_refcounted"`
	IsInstantiable bool     `json:"is_instantiable"`
	Inherits       string   `json:"inherits,omitempty"`
	APIType        string   `json:"api_type"`
	Enums          []Enum   `json:"enums,omitempty"`
	Methods        []Method `json:"methods,omitempty"`
	Signals        []struct {
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
		Value int    `json:"value"`
	} `json:"constants,omitempty"`
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
	BuiltinClasses []struct {
		Name      string `json:"name"`
		IsKeyed   bool   `json:"is_keyed"`
		Operators []struct {
			Name       string `json:"name"`
			RightType  string `json:"right_type,omitempty"`
			ReturnType string `json:"return_type"`
		} `json:"operators"`
		Constructors []struct {
			Index     int `json:"index"`
			Arguments []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Meta string `json:"meta"`
			} `json:"arguments,omitempty"`
		} `json:"constructors"`
		HasDestructor      bool   `json:"has_destructor"`
		IndexingReturnType string `json:"indexing_return_type,omitempty"`
		Methods            []struct {
			Name       string `json:"name"`
			ReturnType string `json:"return_type"`
			IsVararg   bool   `json:"is_vararg"`
			IsConst    bool   `json:"is_const"`
			IsStatic   bool   `json:"is_static"`
			Hash       int64  `json:"hash"`
			Arguments  []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Meta string `json:"meta"`
			} `json:"arguments,omitempty"`
		} `json:"methods,omitempty"`
		Members []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"members,omitempty"`
		Constants []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"constants,omitempty"`
		Enums []Enum `json:"enums,omitempty"`
	} `json:"builtin_classes"`
	Classes    []Class `json:"classes"`
	Singletons []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"singletons"`
	NativeStructures []struct {
		Name   string `json:"name"`
		Format string `json:"format"`
	} `json:"native_structures"`
}

func convertType(meta string, gdType string) string {
	switch gdType {
	case "int":
		/*switch meta {
		case "int8":
			return "int8"
		case "int16":
			return "int16"
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		case "uint8":
			return "uint8"
		case "uint16":
			return "uint16"
		case "uint32":
			return "uint32"
		case "uint64":
			return "uint64"
		}*/
		return "int64"
	case "float":
		/*switch meta {
		case "float":
			return "float32"
		case "double":
			return "float64"
		}*/
		return "float64"
	case "String":
		return "string"
	case "enum::Error":
		return "int64"
	case "const uint8_t **":
		return "*[]byte"
	case "const void*", "const uint8_t*", "const uint8_t *", "PackedByteArray":
		return "[]byte"
	case "PackedStringArray":
		return "[]string"
	case "PackedInt32Array":
		return "[]int32"
	case "PackedInt64Array":
		return "[]int64"
	case "PackedFloat32Array":
		return "[]float32"
	case "PackedFloat64Array":
		return "[]float64"
	case "PackedVector2Array":
		return "[]Vector2"
	case "PackedVector3Array":
		return "[]Vector3"
	case "PackedColorArray":
		return "[]Color"
	case "StringName":
		return "Name"
	case "Variant":
		return "any"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "[]byte"
	default:
		if strings.HasPrefix(gdType, "const") {
			gdType = strings.TrimPrefix(gdType, "const")
		}
		if strings.HasSuffix(gdType, "*") {
			return "*" + gdType[:len(gdType)-1]
		}
		if strings.HasPrefix(gdType, "enum::") {
			return strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "enum::")
		}
		if strings.HasPrefix(gdType, "bitfield::") {
			return strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "bitfield::")
		}
		return gdType
	}
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
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

func genEnum(code io.Writer, prefix string, enum Enum) {
	name := prefix + strings.Replace(enum.Name, ".", "", -1)

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
			fmt.Fprintf(code, "\t%v %v = %v\n", n, name, value.Value)
		}
		fmt.Fprintf(code, ")\n")
	}
}

func generate() error {
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

	var spec Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return err
	}

	var virtuals []Method

	code, err := os.Create("./api.go")
	if err != nil {
		return err
	}

	var exceptions = map[string]bool{
		"GDScriptNativeClass": true,
		"JavaClassWrapper":    true,
		"JavaScript":          true,
	}

	var singletons = make(map[string]bool)
	for _, class := range spec.Singletons {
		singletons[class.Type] = true
	}
	convertClass := func(name string) string {
		if singletons[name] {
			return name + "Singleton"
		}
		return name
	}

	fmt.Fprintln(code, `//go:build !generate`)
	fmt.Fprintln(code)
	fmt.Fprintln(code, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(code, `package gd`)
	fmt.Fprintln(code)
	fmt.Fprintln(code, `import "reflect"`)

	for _, enum := range spec.GlobalEnums {
		genEnum(code, "", enum)
	}

	if len(spec.UtilityFunctions) > 0 {
		fmt.Fprintf(code, "var utilities [%d]cUtilityFunction\n", len(spec.UtilityFunctions))
	}

	for i, utility := range spec.UtilityFunctions {
		fmt.Fprintf(code, "func %v(", convertName(utility.Name))
		for i, arg := range utility.Arguments {
			fmt.Fprintf(code, "%v %v", fixReserved(arg.Name), convertType(arg.Meta, arg.Type))
			if i < len(utility.Arguments)-1 {
				fmt.Fprint(code, ", ")
			}
		}

		result := convertType("", utility.ReturnType)

		fmt.Fprintf(code, ") %v {", result)

		if result == "" {
			result = "struct{}"
		} else {
			fmt.Fprintf(code, "return ")
		}

		fmt.Fprintf(code, "utilityCall[%v](utilities[%d]", result, i)
		for _, arg := range utility.Arguments {
			fmt.Fprint(code, ", ")
			fmt.Fprintf(code, "&%v", fixReserved(arg.Name))
		}
		fmt.Fprintf(code, ") }\n")
	}

	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			genEnum(code, class.Name, enum)
		}

		switch class.Name {
		case "bool", "int", "float", "Nil", "String":
			continue
		case "StringName":
			class.Name = "Name"
		}
		if strings.HasPrefix(class.Name, "Packed") {
			continue
		}

		fmt.Fprintf(code, "type %v = c%v\n", class.Name, class.Name)

		fmt.Fprintf(code, "var method%v [%d]cBuiltInMethod\n", class.Name, len(class.Methods))

		for i, method := range class.Methods {
			result := convertType("", method.ReturnType)

			fmt.Fprintf(code, "func (gdClass %v) %v(", class.Name, convertName(method.Name))
			for i, arg := range method.Arguments {
				fmt.Fprintf(code, "%v %v", fixReserved(arg.Name), convertType(arg.Meta, arg.Type))
				if i < len(method.Arguments)-1 {
					fmt.Fprint(code, ", ")
				}
			}

			fmt.Fprintf(code, ") %v { ", result)
			if result == "" {
				result = "struct{}"
			} else {
				fmt.Fprintf(code, "return ")
			}
			fmt.Fprintf(code, "builtinCall[%v](&gdClass, method%v[%d]", result, class.Name, i)
			for _, arg := range method.Arguments {
				fmt.Fprint(code, ", ")
				fmt.Fprintf(code, "&%v", fixReserved(arg.Name))
			}
			fmt.Fprintf(code, ") }\n")
		}
	}

	var classDB = make(map[string]Class)
	for _, class := range spec.Classes {
		classDB[class.Name] = class
	}

	var isRefCounted func(class Class) bool
	isRefCounted = func(class Class) bool {
		if class.Name == "RefCounted" {
			return true
		}
		if class.Name == "Object" {
			return false
		}
		return isRefCounted(classDB[class.Inherits])
	}

	for _, class := range spec.Classes {
		if class.Name == "RefCounted" {
			continue
		}

		for _, enum := range class.Enums {
			genEnum(code, class.Name, enum)
		}

		if singletons[class.Name] {
			fmt.Fprintf(code, "var %v %v\n", class.Name, convertClass(class.Name))
		}

		class.Name = convertClass(class.Name)

		fmt.Fprintln(code)
		if class.Name != "Object" {
			fmt.Fprintf(code, "type %[1]v struct{self *%[1]v; obj safeObject }\n", class.Name)
		}
		fmt.Fprintf(code, `func New%[1]v(owner InstanceOwner) *%[1]v { at := new(%[1]v); return New%[1]vAt(at, owner) }`+"\n", class.Name, isRefCounted(class))
		fmt.Fprintf(code, `func New%[1]vAt(at *%[1]v, owner InstanceOwner) *%[1]v { at.obj.new(owner, at.className(), %v); return at }`+"\n", class.Name, isRefCounted(class))
		fmt.Fprintf(code, `func (gdClass *%v) Free() { gdClass.obj.free() }`+"\n", class.Name)
		fmt.Fprintf(code, `func (gdClass %v) object() safeObject { return gdClass.obj }`+"\n", class.Name)
		fmt.Fprintf(code, `func (*%v) className() string { return "%v\000" }`+"\n", class.Name, class.Name)
		fmt.Fprintln(code)
		if class.Inherits != "" {
			fmt.Fprintf(code, `func (gdClass %v) %[2]v() %[2]v { return %[2]v{obj:gdClass.obj} }`+"\n", class.Name, convertClass(class.Inherits))
			fmt.Fprintln(code)
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

		if methodCount > 0 {
			fmt.Fprintf(code, "var method%v [%d]cMethodBind\n", class.Name, methodCount)
		}

		// check if a Go type implements that method and returns the method that implements the
		// named virtual.
		fmt.Fprintf(code, "func (gdClass *%v) virtual(rtype reflect.Type, name string) (method reflect.Method, ok bool) {\n", class.Name)
		if virtualCount > 0 {
			fmt.Fprintln(code, "\tswitch name {")
			for _, method := range class.Methods {
				if !method.IsVirtual {
					continue
				}

				result := convertType(method.ReturnValue.Meta, method.ReturnValue.Type)
				virtuals = append(virtuals, method)

				fmt.Fprintf(code, "\tcase \""+method.Name+"\":\n")
				fmt.Fprintf(code, "\t\tif rtype.Implements(reflect.TypeOf([0]interface{ %v(", convertName(method.Name))
				for i, arg := range method.Arguments {
					fmt.Fprintf(code, "%v %v", fixReserved(arg.Name), convertType(arg.Meta, arg.Type))
					if i < len(method.Arguments)-1 {
						fmt.Fprint(code, ", ")
					}
				}
				fmt.Fprintf(code, ") %v }{}).Elem()) {\n", result)
				fmt.Fprintf(code, "\t\t\treturn rtype.MethodByName(`%s`)\n", convertName(method.Name))
				fmt.Fprintln(code, "\t\t}")
				fmt.Fprintln(code, "\t\treturn")
			}
			fmt.Fprintln(code, "\t}")
		}
		if class.Inherits != "" {
			fmt.Fprintf(code, "\treturn new(%v).virtual(rtype, name)\n", convertClass(class.Inherits))
		} else {
			fmt.Fprintln(code, "\treturn")
		}
		fmt.Fprintln(code, "}")

		var i int
		for _, method := range class.Methods {
			if method.IsVirtual {
				continue
			}

			result := convertType(method.ReturnValue.Meta, method.ReturnValue.Type)

			fmt.Fprintf(code, "func (gdClass %v) %v(", class.Name, convertName(method.Name))
			for i, arg := range method.Arguments {
				fmt.Fprintf(code, "%v %v", fixReserved(arg.Name), convertType(arg.Meta, arg.Type))
				if i < len(method.Arguments)-1 {
					fmt.Fprint(code, ", ")
				}
			}
			fmt.Fprintf(code, ") %v { ", result)
			if result == "" {
				result = "struct{}"
			} else {
				fmt.Fprintf(code, "return ")
			}
			fmt.Fprintf(code, "methodCall[%v](gdClass.obj.get(), method%v[%d]", result, class.Name, i)
			for _, arg := range method.Arguments {
				fmt.Fprint(code, ", ")
				fmt.Fprintf(code, "&%v", fixReserved(arg.Name))
			}
			fmt.Fprintf(code, ") }\n")

			i++
		}
	}

	fmt.Fprintln(code)
	fmt.Fprintf(code, `func init() {`)
	fmt.Fprintf(code, "\tcOnLoad = func() {\n")
	for _, singleton := range spec.Singletons {
		fmt.Fprintf(code, "%v.obj.hold(global_get_singleton(`%v`))\n", singleton.Type, singleton.Name)
	}
	for i, utility := range spec.UtilityFunctions {
		fmt.Fprintf(code, "\t\t"+`utilities[%d] = get_utility_function("%v\000", %v)`+"\n", i, utility.Name, utility.Hash)
	}
	for _, class := range spec.BuiltinClasses {
		var name = class.Name
		switch class.Name {
		case "bool", "int", "float", "Nil", "String":
			continue
		case "StringName":
			name = "Name"
		}
		if strings.HasPrefix(name, "Packed") {
			continue
		}

		for i, method := range class.Methods {
			fmt.Fprintf(code, "\t\t"+`method%v[%d] = cVariantType%v.get_builtin_method("%v\000", %v)`+"\n", name, i, class.Name, method.Name, method.Hash)
		}
	}
	for _, class := range spec.Classes {
		var i int
		for _, method := range class.Methods {
			if _, ok := exceptions[class.Name]; ok {
				continue
			}

			if method.IsVirtual {
				continue
			}
			fmt.Fprintf(code, "\t\t"+`method%v[%d] = classDB.get_method_bind("%v\000", "%v\000", %v)`+"\n", convertClass(class.Name), i, class.Name, method.Name, method.Hash)
			i++
		}
	}
	fmt.Fprintf(code, "\t}\n}\n\n")

	return code.Close()
}

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
