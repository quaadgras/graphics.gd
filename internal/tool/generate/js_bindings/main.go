package main

import (
	"fmt"
	"reflect"

	"graphics.gd/internal/gdextension"
)

func main() {
	GDExtension := reflect.TypeFor[gdextension.API]()
	for i := range GDExtension.NumField() {
		field := GDExtension.Field(i)
		if field.Type.Kind() != reflect.Func {
			continue
		}
		ftype := field.Type
		if ftype.NumOut() == 0 {
			fmt.Print("void")
		} else {
			result, _ := convertGoTypeToCpp(ftype.Out(0))
			fmt.Print(result)
		}
		fmt.Print(" js_" + field.Name)
		fmt.Print("(")
		var num int
		for j := range ftype.NumIn() {
			atype, count := convertGoTypeToCpp(ftype.In(j))
			if count == 0 {
				continue
			}
			for k := 0; k < count; k++ {
				if num > 0 {
					fmt.Print(", ")
				}
				fmt.Print(atype)
				fmt.Print(" arg", j)
				if count > 1 {
					fmt.Print("_", k)
				}
				if atype == "val" {
					fmt.Print("_js")
				}
				num++
			}
		}
		fmt.Println(") {")
		for j := range ftype.NumIn() {
			atype, count := convertGoTypeToCpp(ftype.In(j))
			if atype == "val" {
				fmt.Printf("\t%s arg%d;\n", convertGoTypeToInternal(ftype.In(j)), j)
			}
			if count <= 1 {
				if count == 0 {
					fmt.Printf("\tvoid *arg%d = %s;\n", j, atype)
				}
				continue
			}
			fmt.Printf("\t%s arg%d[%d] = {", atype, j, count)
			for k := 0; k < count; k++ {
				fmt.Printf(" arg%d_%d", j, k)
				if k < count-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println("};")
		}
		if ftype.NumOut() > 0 {
			result, _ := convertGoTypeToCpp(ftype.Out(0))
			fmt.Printf("\treturn (%s)(gdextension_%s(", result, field.Name)
		} else {
			fmt.Printf("\tgdextension_%s(", field.Name)
		}
		for j := range ftype.NumIn() {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf(convertValueTo(ftype.In(j)), fmt.Sprint("arg", j))
		}
		if ftype.NumOut() > 0 {
			fmt.Println("));")
		} else {
			fmt.Println(");")
		}
		if vtype := reflect.TypeFor[gdextension.Uninitialized[*gdextension.GodotVersion]](); ftype.NumIn() > 0 && ftype.In(0) == vtype {
			vtype = reflect.TypeFor[gdextension.GodotVersion]()
			for i := 0; i < vtype.NumField(); i++ {
				field := vtype.Field(i)
				if field.Name == "_" {
					continue
				}
				fmt.Printf("\targ%d_js.set(%q, arg%d.%s);\n", 0, field.Name, 0, field.Name)
			}
		}
		fmt.Println("}")
		if i > 20 {
			return
		}
	}
}

func convertValueTo(rtype reflect.Type) string {
	switch rtype.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "%v"
	case reflect.String:
		return "%v.c_str()"
	default:
		switch reflect.Zero(rtype).Interface().(type) {
		case gdextension.Uninitialized[*gdextension.GodotVersion]:
			return "&%v"
		}
		return fmt.Sprintf("(%s)%v", convertGoTypeToInternal(rtype), "%v")
	}
}

func convertGoTypeToInternal(rtype reflect.Type) string {
	ctype, _ := convertGoTypeToCpp(rtype)
	switch rtype.Kind() {
	case reflect.UnsafePointer:
		return "void *"
	default:
		switch reflect.Zero(rtype).Interface().(type) {
		case gdextension.Uninitialized[*gdextension.GodotVersion]:
			return "GDExtensionGodotVersion"
		case gdextension.Const[gdextension.StringNamePtr]:
			return "GDExtensionConstStringNamePtr"
		case gdextension.Const[gdextension.VariantPtr]:
			return "GDExtensionConstVariantPtr"
		case gdextension.VariantPtr:
			return "GDExtensionVariantPtr"
		case gdextension.Uninitialized[gdextension.VariantPtr]:
			return "GDExtensionUninitializedVariantPtr"
		}
		return ctype
	}
}

func convertGoTypeToCpp(rtype reflect.Type) (string, int) {
	switch rtype.Kind() {
	case reflect.Int:
		return "int", 1
	case reflect.Uint:
		return "unsigned int", 1
	case reflect.Int8:
		return "int8_t", 1
	case reflect.Uint8:
		return "uint8_t", 1
	case reflect.Int16:
		return "int16_t", 1
	case reflect.Uint16:
		return "uint16_t", 1
	case reflect.Uintptr:
		return "size_t", 1
	case reflect.Int32:
		return "int32_t", 1
	case reflect.Uint32:
		return "uint32_t", 1
	case reflect.Int64:
		return "int64_t", 1
	case reflect.Uint64:
		return "uint64_t", 1
	case reflect.Float32:
		return "float", 1
	case reflect.Float64:
		return "double", 1
	case reflect.String:
		return "std::string", 1
	case reflect.Bool:
		return "bool", 1
	case reflect.Pointer:
		switch rtype.Elem().Kind() {
		case reflect.Struct:
			return "val", 1
		}
		fallthrough
	case reflect.UnsafePointer:
		return "uintptr_t", 1
	default:
		switch reflect.Zero(rtype).Interface().(type) {
		case gdextension.Uninitialized[*gdextension.GodotVersion]:
			return "val", 1
		case gdextension.Const[gdextension.StringNamePtr]:
			return "uint32_t", 1
		case gdextension.Const[gdextension.VariantPtr]:
			return "uint32_t", 6
		case gdextension.Uninitialized[gdextension.VariantPtr]:
			return "&result_buffer[0]", 0
		}
		return "void *", 1
	}
}
