//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
)

func propertyOf(godot Lifetime, field reflect.StructField) gd.PropertyInfo {
	var name = field.Name
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var hint gd.PropertyHint
	var hintString = classNameOf(field.Type)
	vtype := variantTypeOf(field.Type)
	if vtype == TypeArray {
		_, generic, ok := strings.Cut(field.Type.String(), "[")
		if ok {
			hint |= PropertyHintArrayType
			split := strings.Split(generic, ".")
			elem := split[len(split)-1]
			elem = elem[:len(elem)-1]
			hintString = fmt.Sprintf("%d/%d:%s", gd.TypeObject, PropertyHintResourceType, elem) // MAKE_RESOURCE_TYPE_HINT
		}
	}
	if field.Type.Implements(reflect.TypeOf([0]isResource{}).Elem()) {
		hint |= PropertyHintResourceType
	}
	var usage = PropertyUsageStorage | PropertyUsageEditor
	if vtype == TypeNil {
		usage |= PropertyUsageNilIsVariant
	}
	if rangeHint, ok := field.Tag.Lookup("range"); ok {
		hint |= PropertyHintRange
		hintString = rangeHint
	}
	return gd.PropertyInfo{
		Type:       vtype,
		Name:       godot.StringName(name),
		ClassName:  godot.StringName(classNameOf(field.Type)),
		Hint:       hint,
		HintString: godot.String(hintString),
		Usage:      usage,
	}
}

func classNameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr {
		return classNameOf(rtype.Elem())
	}
	if rtype.Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()) {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
		}
		if rtype.Name() == "" && rtype.Field(0).Anonymous {
			return rtype.Field(0).Name
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}

func variantTypeOf(rtype reflect.Type) (vtype VariantType) {
	switch rtype.Kind() {
	case reflect.Int32, reflect.Int64, reflect.Int:
		return TypeInt
	}
	switch rtype {
	case reflect.TypeOf([0]Variant{}).Elem():
		vtype = TypeNil
	case reflect.TypeOf([0]Bool{}).Elem():
		vtype = TypeBool
	case reflect.TypeOf([0]Int{}).Elem():
		vtype = TypeInt
	case reflect.TypeOf([0]Float{}).Elem():
		vtype = TypeFloat
	case reflect.TypeOf([0]String{}).Elem():
		vtype = TypeString
	case reflect.TypeOf([0]Vector2{}).Elem():
		vtype = TypeVector2
	case reflect.TypeOf([0]Vector2i{}).Elem():
		vtype = TypeVector2i
	case reflect.TypeOf([0]Rect2{}).Elem():
		vtype = TypeRect2
	case reflect.TypeOf([0]Rect2i{}).Elem():
		vtype = TypeRect2i
	case reflect.TypeOf([0]Vector3{}).Elem():
		vtype = TypeVector3
	case reflect.TypeOf([0]Vector3i{}).Elem():
		vtype = TypeVector3i
	case reflect.TypeOf([0]Transform2D{}).Elem():
		vtype = TypeTransform2d
	case reflect.TypeOf([0]Vector4{}).Elem():
		vtype = TypeVector4
	case reflect.TypeOf([0]Vector4i{}).Elem():
		vtype = TypeVector4i
	case reflect.TypeOf([0]Plane{}).Elem():
		vtype = TypePlane
	case reflect.TypeOf([0]Quaternion{}).Elem():
		vtype = TypeQuaternion
	case reflect.TypeOf([0]AABB{}).Elem():
		vtype = TypeAabb
	case reflect.TypeOf([0]Basis{}).Elem():
		vtype = TypeBasis
	case reflect.TypeOf([0]Transform3D{}).Elem():
		vtype = TypeTransform3d
	case reflect.TypeOf([0]Projection{}).Elem():
		vtype = TypeProjection
	case reflect.TypeOf([0]Color{}).Elem():
		vtype = TypeColor
	case reflect.TypeOf([0]StringName{}).Elem():
		vtype = TypeStringName
	case reflect.TypeOf([0]NodePath{}).Elem():
		vtype = TypeNodePath
	case reflect.TypeOf([0]RID{}).Elem():
		vtype = TypeRid
	case reflect.TypeOf([0]Object{}).Elem():
		vtype = TypeObject
	case reflect.TypeOf([0]Callable{}).Elem():
		vtype = TypeCallable
	case reflect.TypeOf([0]Dictionary{}).Elem():
		vtype = TypeDictionary
	case reflect.TypeOf([0]Array{}).Elem():
		vtype = TypeArray
	case reflect.TypeOf([0]PackedByteArray{}).Elem():
		vtype = TypePackedByteArray
	case reflect.TypeOf([0]PackedInt32Array{}).Elem():
		vtype = TypePackedInt32Array
	case reflect.TypeOf([0]PackedInt64Array{}).Elem():
		vtype = TypePackedInt64Array
	case reflect.TypeOf([0]PackedFloat32Array{}).Elem():
		vtype = TypePackedFloat32Array
	case reflect.TypeOf([0]PackedFloat64Array{}).Elem():
		vtype = TypePackedFloat64Array
	case reflect.TypeOf([0]PackedStringArray{}).Elem():
		vtype = TypePackedStringArray
	case reflect.TypeOf([0]PackedVector2Array{}).Elem():
		vtype = TypePackedVector2Array
	case reflect.TypeOf([0]PackedVector3Array{}).Elem():
		vtype = TypePackedVector3Array
	case reflect.TypeOf([0]PackedColorArray{}).Elem():
		vtype = TypePackedColorArray
	case reflect.TypeOf([0]unsafe.Pointer{}).Elem():
		vtype = TypeNil
	case reflect.TypeOf([0]*classdb.ScriptLanguageExtensionProfilingInfo{}).Elem():
		vtype = TypeNil
	default:
		switch {
		case rtype.Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()):
			vtype = TypeObject
		case rtype.Implements(reflect.TypeOf([0]gd.IsArray{}).Elem()):
			vtype = TypeArray
		case rtype.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()):
			vtype = TypeSignal
		case rtype.ConvertibleTo(reflect.TypeOf(Dictionary{})):
			vtype = TypeDictionary
		default:
			panic("gdextension.RegisterClass: unsupported property type " + rtype.String())
		}
	}
	return
}
