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
	"sort"
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
	case "int", "Int":
		return "Int"
	case "float", "Float":
		return "Float"
	case "bool", "Bool":
		return "bool"
	case "String":
		return "String"
	case "StringName":
		return "StringName"
	case "enum::GDExtension.InitializationLevel":
		return "ExtensionInitializationLevel"
	case "enum::GDExtensionManager.LoadStatus":
		return "ExtensionManagerLoadStatus"
	case "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedColorArray", "PackedByteArray":
		return gdType
	case "Variant":
		return "Variant"
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
	case "interface":
		return "intf"
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

func classWhitelisted(s string) bool {
	switch s {
	case "Object", "Node", "CanvasItem", "Node2D", "GDExtension", "Texture2D",
		"Font", "XMLParser", "FileAccess", "Script", "ScriptLanguage", "StyleBox",
		"RefCounted", "Resource", "TextServer", "Mesh", "MultiMesh", "World2D",
		"Material", "InputEvent", "Window", "SceneTree", "Texture", "ConcavePolygonShape3D",
		"ConvexPolygonShape3D", "TriangleMesh", "Tween", "Viewport", "MultiplayerAPI",
		"MainLoop", "SceneTreeTimer", "Shape3D", "MultiplayerPeer", "Image", "PackedScene",
		"PacketPeer", "SceneState", "ArrayMesh", "PropertyTweener", "IntervalTweener",
		"CallbackTweener", "MethodTweener", "ViewportTexture", "Camera2D", "Tweener",
		"Control", "Theme", "World3D", "Camera3D", "PhysicsDirectSpaceState2D", "Node3D",
		"Environment", "CameraAttributes", "PhysicsPointQueryParameters2D",
		"PhysicsRayQueryParameters2D", "PhysicsShapeQueryParameters2D", "Sky",
		"Node3DGizmo", "PhysicsDirectSpaceState3D", "PhysicsPointQueryParameters3D",
		"PhysicsRayQueryParameters3D", "PhysicsShapeQueryParameters3D", "Engine",
		"RenderingServer", "RenderingDevice", "RDTextureFormat", "RDTextureView",
		"RDAttachmentFormat", "RDFramebufferPass", "Shader", "RDSamplerState",
		"RDVertexAttribute", "RDShaderSource", "RDShaderSPIRV", "RDUniform",
		"RDPipelineRasterizationState", "RDPipelineMultisampleState",
		"RDPipelineDepthStencilState", "RDPipelineColorBlendState", "RDPipelineColorBlendStateAttachment",
		"RDPipelineSpecializationConstant":
		return true
	default:
		return false
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

	out, err := os.Create("./out.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, `//go:build !generate`)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(out, `package gd`)

	all, err := os.Create("./all.go")
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
	fmt.Fprintln(all, `import "runtime/cgo"`)
	fmt.Fprintln(all, `import "runtime.link/mmm"`)
	fmt.Fprintln(all)

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
		fmt.Fprintf(out, "\t%v func(CallFrameBack,CallFrameArgs) `hash:\"%v\"`\n", utility.Name, utility.Hash)
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type builtin struct{\n")
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			genEnum("gd", val, class.Name, enum)
		}
		fmt.Fprintf(out, "\t%v struct{\n", class.Name)
		for _, method := range class.Methods {
			if method.Name == "map" {
				method.Name = "map_"
			}
			fmt.Fprintf(out, "\t\t%v func(base, args CallFrameArgs, ret CallFrameBack, c int32) `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type variant struct{\n")
	fmt.Fprintf(out, "\tcreation struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v [%d]func(CallFrameBack,CallFrameArgs)\n", class.Name, len(class.Constructors))
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
				fmt.Fprintf(out, "\t\t\t%v func(uintptr,uintptr,uintptr)\n", mapOperator(op.Name))
				last = ""
			} else {
				if op.Name != last {
					if last != "" {
						fmt.Fprintf(out, "\t\t\t}\n")
					}
					fmt.Fprintf(out, "\t\t\t%v struct {\n", mapOperator(op.Name))
				}
				fmt.Fprintf(out, "\t\t\t\t%v func(uintptr,uintptr,uintptr)\n", op.RightType)
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
		fmt.Fprintf(out, "\t\t%v func(CallFrameArgs)\n", class.Name)
		for _, method := range class.Methods {
			methodCall(all, class, method, "")
		}
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type methods struct{\n")
	for _, class := range spec.Classes {
		if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
			continue
		}
		if !classWhitelisted(class.Name) {
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
			fmt.Fprintf(out, "\t\t%v MethodBind `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	var classDB = make(map[string]Class)
	for _, class := range spec.Classes {
		classDB[class.Name] = class
	}
	for _, class := range spec.Classes {
		if !classWhitelisted(class.Name) {
			continue
		}

		for _, enum := range class.Enums {
			genEnum("gd", val, class.Name, enum)
		}
		fmt.Fprintf(pub, "type %v = class%v \n", class.Name, class.Name)
		if class.Inherits != "" {
			fmt.Fprintf(all, "type class%[1]v struct {Class[class%[1]v, %v]}\n", class.Name, class.Inherits)
		} else {
			fmt.Fprintf(all, "type class%[1]v struct {Class[class%[1]v, Pointer]}\n", class.Name)
		}
		if class.Inherits != "" {
			var i = 1
			super := classDB[class.Inherits]
			for super.Name != "" {
				fmt.Fprintf(all, "\nfunc (self class%[1]v) %[2]v() %[2]v { return *self%s }\n", class.Name, super.Name, strings.Repeat(".Super()", i))
				i++
				super = classDB[super.Inherits]
			}

		}
		if singletons[class.Name] {
			fmt.Fprintf(all, "\nfunc (self class%[1]v) isSingleton() {}\n", class.Name)
		}
		for _, method := range class.Methods {
			methodCall(all, class, method, "class")
		}
		fmt.Fprintf(all, "\nfunc (self class%[1]v) virtual(name string) reflect.Value {\n", class.Name)
		fmt.Fprintf(all, "\tswitch name {\n")
		for _, method := range class.Methods {
			if method.IsVirtual {
				fmt.Fprintf(all, "\tcase \"%v\": return reflect.ValueOf(self.%v);\n", method.Name, method.Name)
			}
		}
		fmt.Fprintf(all, "\tdefault: return self.Super().virtual(name)\n")
		fmt.Fprintf(all, "\t}\n")
		fmt.Fprintf(all, "}\n")
	}

	return nil
}

func isPointer(t string) (string, bool) {
	switch t {
	case "String", "StringName", "NodePath",
		"Dictionary", "Array":
		return "uintptr", true
	case "Callable", "Signal",
		"PackedByteArray", "PackedInt32Array",
		"PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedStringArray",
		"PackedVector2Array", "PackedVector3Array",
		"PackedColorArray":
		return "[2]uintptr", true
	case "Variant":
		return "[3]uintptr", true
	default:
		return "", false
	}
}

func methodCall(w io.Writer, class Class, method Method, prefix string) {
	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	result := convertType("gd", method.ReturnValue.Meta, method.ReturnValue.Type)
	if method.ReturnType != "" {
		result = convertType("gd", "", method.ReturnType)
	}
	ptrKind, isPtr := isPointer(result)

	if method.IsVirtual {
		fmt.Fprintf(w, "\nfunc (%s%v) %v(impl func(ptr unsafe.Pointer, ctx Context", prefix, class.Name, method.Name)
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), convertType("gd", arg.Meta, arg.Type))
		}
		fmt.Fprintf(w, ") %v, api *API) (cb ExtensionClassCallVirtualFunc) {\n", result)
		fmt.Fprintf(w, "\tcb.Set(func(class cgo.Handle, p_args godotArgs, p_back godotBack) {\n")
		fmt.Fprintf(w, "\tctx := newContext(api)\n")
		for i, arg := range method.Arguments {
			fmt.Fprintf(w, "\t\tvar %v = godotGet[%v](p_args,%d)\n", fixReserved(arg.Name), convertType("gd", arg.Meta, arg.Type), i)
		}
		fmt.Fprintf(w, "\tself := reflect.ValueOf(class.Value()).UnsafePointer()\n")
		if result != "" {
			fmt.Fprintf(w, "\t\tret := ")
		}
		fmt.Fprintf(w, "impl(self,ctx")
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v", fixReserved(arg.Name))
		}
		fmt.Fprintf(w, ")\n")
		if isPtr {
			fmt.Fprintf(w, "\t\tmmm.MarkFree(ret)\n")
		}
		fmt.Fprintf(w, "\tctx.Free()\n")
		if result != "" {
			fmt.Fprintf(w, "\t\tgodotSet[%v](p_back, ret)\n", result)
		}
		fmt.Fprintf(w, "\t})\n")
		fmt.Fprintf(w, "\treturn\n")
		fmt.Fprintf(w, "}\n")
		return
	}
	if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
		return
	}
	var context = "ctx Context"
	if !isPtr {
		context = ""
	}
	fmt.Fprintf(w, "\nfunc (self %s%v) %v(%s", prefix, class.Name, convertName(method.Name), context)

	if method.Name == "select" {
		method.Name = "select_"
	}
	if method.Name == "map" {
		method.Name = "map_"
	}

	for i, arg := range method.Arguments {
		if i > 0 || isPtr {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), convertType("gd", arg.Meta, arg.Type))
	}
	fmt.Fprintf(w, ") %v {\n", result)
	if prefix == "" {
		fmt.Fprintf(w, "\tvar selfPtr = self\n")
	} else {
		fmt.Fprintf(w, "\tvar selfPtr = self.getPointer()\n")
	}
	fmt.Fprintf(w, "\tvar abi = selfPtr.API.newFrame()\n")
	if prefix == "" {
		for i, arg := range method.Arguments {
			fmt.Fprintf(w, "\tframeSet[%[2]v](%[1]v, abi, %[3]v)\n", i+1, convertType("gd", arg.Meta, arg.Type), fixReserved(arg.Name))
		}
		fmt.Fprintf(w, "\tframeSet(0, abi, selfPtr.Pointer())\n")
		fmt.Fprintf(w, "\tselfPtr.API.builtin.%v.%v(abi.Get(0), abi.Get(1), abi.Back(), %d)\n", class.Name, method.Name, len(method.Arguments))
	} else {
		for i, arg := range method.Arguments {
			fmt.Fprintf(w, "\tframeSet[%[2]v](%[1]v, abi, %[3]v)\n", i, convertType("gd", arg.Meta, arg.Type), fixReserved(arg.Name))
		}
		fmt.Fprintf(w, "\tselfPtr.API.Object.UnsafeCall(selfPtr.API.methods.%v.%v, selfPtr.Pointer(), abi.Args(), abi.Back())\n", class.Name, method.Name)
	}

	if isPtr {
		fmt.Fprintf(w, "\tvar ret = frameGet[%v](abi)\n", ptrKind)
	} else {
		if result != "" {
			fmt.Fprintf(w, "\tvar ret = frameGet[%v](abi)\n", result)
		}
	}
	fmt.Fprintf(w, "\tabi.free()\n")
	if result != "" {
		if isPtr {
			fmt.Fprintf(w, "\treturn mmm.Make[API,%v,%v](ctx,selfPtr.API,ret)", result, ptrKind)
		} else {
			fmt.Fprintf(w, "\treturn ret\n")
		}
	}
	fmt.Fprintf(w, "}\n")
}

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
