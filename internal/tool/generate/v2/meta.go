package main

import (
	"iter"
	"reflect"
	"strings"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdjson"
)

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
	case "class":
		return "class_"
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
	case "object":
		return "obj"
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

func isBuiltin(s string) bool {
	switch s {
	case "String", "StringName", "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedVector4Array", "PackedColorArray", "PackedByteArray",
		"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID",
		"Callable", "Signal", "Dictionary", "Array":
		return true
	default:
		return false
	}
}

func importsVariant(class gdjson.Class, identifier, s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if strings.HasPrefix(s, "typedarray::") {
			s = strings.TrimPrefix(s, "typedarray::")
			for pkg := range importsVariant(class, identifier, s) {
				if !yield(pkg) {
					return
				}
			}
			return
		}
		switch s {
		case "Float", "float":
			yield("graphics.gd/variant/Float")
		case "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
			"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath":
			yield("graphics.gd/variant/" + s)
		case "Signal":
			//return "graphics.gd/variant/" + s
			return
		case "PackedVector2Array":
			yield("graphics.gd/variant/Vector2")
		case "PackedVector3Array":
			yield("graphics.gd/variant/Vector3")
		case "PackedVector4Array":
			yield("graphics.gd/variant/Vector4")
		case "PackedColorArray":
			yield("graphics.gd/variant/Color")
		case "Callable":
			details := gdjson.Callables[identifier]
			if len(details) == 0 {
				return
			}
			for _, detail := range details {
				if detail == "void" {
					continue
				}
				detail, _, _ = strings.Cut(detail, " ")
				for pkg := range importsVariant(class, "", detail) {
					if !yield(pkg) {
						return
					}
				}
			}
		}
	}
}

func (classDB ClassDB) convertType(pkg, meta string, gdType string) string {
	maybeInternal := func(name string) string {
		return "gd." + name
	}
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		meta, rest, ok := strings.Cut(gdType, ":")
		if ok {
			gdType = rest
		}
		return "Array.Contains[" + classDB.convertType(pkg, meta, gdType) + "]"
	}
	switch gdType {
	case "int", "Int":
		return maybeInternal("Int")
	case "float", "Float":
		return maybeInternal("Float")
	case "bool", "Bool":
		return "bool"
	case "StringName":
		return maybeInternal("StringName")
	case "enum::GDExtension.InitializationLevel":
		return maybeInternal("GDExtensionInitializationLevel")
	case "enum::GDExtensionManager.LoadStatus":
		return "gdclass.GDExtensionManagerLoadStatus"
	case "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedVector4Array", "PackedColorArray", "PackedByteArray",
		"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID",
		"Signal":
		return maybeInternal(gdType)
	case "Array":
		return "Array.Any"
	case "Dictionary":
		return "Dictionary.Any"
	case "String":
		return "String.Readable"
	case "Callable":
		return "Callable.Function"
	case "Variant":
		return maybeInternal("Variant")
	case "enum::Variant.Type":
		return maybeInternal("VariantType")

	// strange C++ cases
	case "enum::Error":
		return "gd.Error"
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
	case "Object":
		return "[1]gd.Object"
	case "RefCounted":
		return "gd." + gdType
	default:
		gdType = strings.TrimPrefix(gdType, "const")

		if strings.HasSuffix(gdType, "*") {
			gdType = strings.TrimPrefix(gdType, pkg)
			return "*" + gdType[:len(gdType)-1]
		}

		if strings.HasPrefix(gdType, "enum::") || strings.HasPrefix(gdType, "bitfield::") {
			gdType = strings.TrimPrefix(gdType, "enum::")
			gdType = strings.TrimPrefix(gdType, "bitfield::")
			host, sub, hasHost := strings.Cut(gdType, ".")
			if isBuiltin(host) {
				return "gd." + host + sub
			}
			if hasHost {
				return "gdclass." + host + sub
			} else {
				return gdType
			}
		}

		gdType = strings.Replace(gdType, ".", "", -1)
		gdType = strings.Replace(gdType, ".", "", -1)

		gdType = strings.Replace(gdType, ".", "", -1)

		if gdType == "Object" {
			return "[1]gd.Object"
		}

		if class, ok := classDB[gdType]; ok {
			return "[1]gdclass." + class.Name
		}

		if inCore(gdType) {
			return maybeInternal(gdType)
		}

		if gdType != "" {
			return "[1]gdclass." + gdType
		}
		return gdType
	}
}

var StructablesInThisPackageGlobalHack = make(map[reflect.Type]bool)

func (classDB ClassDB) convertTypeSimple(class gdjson.Class, lookup, meta string, gdType string) string {
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		meta, rest, ok := strings.Cut(gdType, ":")
		if ok {
			gdType = rest
		}
		return "[]" + classDB.convertTypeSimple(class, lookup, meta, gdType)
	}
	switch gdType {
	case "int", "Int":
		return "int"
	case "float", "Float":
		return "Float.X"
	case "bool", "Bool":
		return "bool"
	case "String", "StringName":
		return "string"
	case "PackedByteArray":
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
		return "[]Vector2.XY"
	case "PackedVector3Array":
		return "[]Vector3.XYZ"
	case "PackedVector4Array":
		return "[]Vector4.XYZW"
	case "PackedColorArray":
		return "[]Color.RGBA"
	case "enum::Error":
		return "error"
	case "Vector2":
		return "Vector2.XY"
	case "Vector2i":
		return "Vector2i.XY"
	case "Rect2":
		return "Rect2.PositionSize"
	case "Rect2i":
		return "Rect2i.PositionSize"
	case "Vector3":
		return "Vector3.XYZ"
	case "Vector3i":
		return "Vector3i.XYZ"
	case "Transform2D":
		return "Transform2D.OriginXY"
	case "Vector4":
		return "Vector4.XYZW"
	case "Vector4i":
		return "Vector4i.XYZW"
	case "Plane":
		return "Plane.NormalD"
	case "Quaternion":
		return "Quaternion.IJKX"
	case "AABB":
		return "AABB.PositionSize"
	case "Basis":
		return "Basis.XYZ"
	case "Transform3D":
		return "Transform3D.BasisOrigin"
	case "Projection":
		return "Projection.XYZW"
	case "Color":
		return "Color.RGBA"
	case "NodePath":
		return "NodePath.String"
	case "RID", "gd.RID":
		if class.Name == "Resource" {
			return "ID"
		}
		if typed, ok := gdjson.Resources[lookup]; ok {
			if typed == reflect.TypeFor[gd.RID]() {
				return "RID.Any"
			}
			return typed.String()
		}
		return "RID.Any"
	case "ObjectID":
		return "Object.ID"
	case "Signal":
		return "chan []any"
	case "Dictionary":
		rtype, ok := gdjson.Structables[lookup]
		if !ok {
			return "map[any]any"
		}
		registerStructables(rtype)
		if rtype.Name() == "" {
			return strings.Replace(rtype.String(), "gdjson.", "", -1)
		}
		return strings.Replace(rtype.Name(), "gdjson.", "", -1)
	case "Array":
		return "[]any"
	case "Variant":
		return "any"
	case "Object":
		return "Object.Instance"
	case "Callable":
		details, ok := gdjson.Callables[lookup]
		if !ok || len(details) == 0 || strings.HasSuffix(lookup, ".") {
			return "Callable.Function"
		}
		var ftype string = "func("
		for i, arg := range details[1:] {
			if i > 0 {
				ftype += ", "
			}
			atype, name, ok := strings.Cut(arg, " ")
			if ok {
				ftype += name + " " + classDB.convertTypeSimple(class, "", "", atype)
			} else {
				ftype += classDB.convertTypeSimple(class, "", "", arg)
			}
		}
		ftype += ")"
		if details[0] != "void" {
			ftype += " " + classDB.convertTypeSimple(class, "", "", details[0])
		}
		return ftype
	default:
		return classDB.convertType(class.Name, meta, gdType)
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
