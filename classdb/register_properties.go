package classdb

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	ResourceClass "graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
)

func propertyOf(field reflect.StructField) (gd.PropertyInfo, bool) {
	var name = field.Name
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var hint PropertyHint
	var hintString = nameOf(field.Type)
	vtype, ok := variantTypeOf(field.Type)
	if !ok {
		return gd.PropertyInfo{}, false
	}
	if vtype == gd.TypeArray {
		_, generic, ok := strings.Cut(field.Type.String(), "[")
		if ok {
			hint |= PropertyHintArrayType
			split := strings.Split(generic, ".")
			elem := split[len(split)-1]
			elem = elem[:len(elem)-1]
			hintString = fmt.Sprintf("%d/%d:%s", gd.TypeObject, PropertyHintResourceType, elem) // MAKE_RESOURCE_TYPE_HINT
		}
	}
	if field.Type.Implements(reflect.TypeOf([0]interface{ AsResource() ResourceClass.Instance }{}).Elem()) {
		hint |= PropertyHintResourceType
	}
	var usage = PropertyUsageStorage | PropertyUsageEditor
	if vtype == gd.TypeNil {
		usage |= PropertyUsageNilIsVariant
	}
	if rangeHint, ok := field.Tag.Lookup("range"); ok {
		hint |= PropertyHintRange
		hintString = rangeHint
	}
	return gd.PropertyInfo{
		Type:       vtype,
		Name:       gd.NewStringName(name),
		ClassName:  gd.NewStringName(nameOf(field.Type)),
		Hint:       int64(hint),
		HintString: gd.NewString(hintString),
		Usage:      int64(usage),
	}, true
}

func variantTypeOf(rtype reflect.Type) (vtype gd.VariantType, ok bool) {
	switch rtype.Kind() {
	case reflect.Bool:
		return gd.TypeBool, true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint:
		return gd.TypeInt, true
	case reflect.Uint64:
		return gd.TypeRID, true
	case reflect.Float32, reflect.Float64:
		return gd.TypeFloat, true
	case reflect.Pointer:
		return variantTypeOf(rtype.Elem())
	case reflect.Func:
		return gd.TypeCallable, true
	case reflect.Array:
		return gd.TypeArray, true
	case reflect.String:
		return gd.TypeString, true
	case reflect.Slice:
		switch rtype.Elem().Kind() {
		case reflect.Uint8:
			return gd.TypePackedByteArray, true
		case reflect.Int32:
			return gd.TypePackedInt32Array, true
		case reflect.Int64:
			return gd.TypePackedInt64Array, true
		case reflect.Float32:
			return gd.TypePackedFloat32Array, true
		case reflect.Float64:
			return gd.TypePackedFloat64Array, true
		case reflect.String:
			return gd.TypePackedStringArray, true
		default:
			switch rtype.Elem() {
			case reflect.TypeFor[gd.Vector2]():
				return gd.TypePackedVector2Array, true
			case reflect.TypeFor[gd.Vector3]():
				return gd.TypePackedVector3Array, true
			case reflect.TypeFor[gd.Color]():
				return gd.TypePackedColorArray, true
			case reflect.TypeFor[gd.Vector4]():
				return gd.TypePackedVector4Array, true
			default:
				return gd.TypeArray, true
			}
		}
	case reflect.Map:
		return gd.TypeDictionary, true
	case reflect.Struct:
		switch rtype {
		case reflect.TypeOf([0]gd.Variant{}).Elem():
			vtype = gd.TypeNil
		case reflect.TypeOf([0]gd.Bool{}).Elem():
			vtype = gd.TypeBool
		case reflect.TypeOf([0]gd.Int{}).Elem():
			vtype = gd.TypeInt
		case reflect.TypeOf([0]gd.Float{}).Elem():
			vtype = gd.TypeFloat
		case reflect.TypeOf([0]gd.String{}).Elem():
			vtype = gd.TypeString
		case reflect.TypeOf([0]gd.Vector2{}).Elem():
			vtype = gd.TypeVector2
		case reflect.TypeOf([0]gd.Vector2i{}).Elem():
			vtype = gd.TypeVector2i
		case reflect.TypeOf([0]gd.Rect2{}).Elem():
			vtype = gd.TypeRect2
		case reflect.TypeOf([0]gd.Rect2i{}).Elem():
			vtype = gd.TypeRect2i
		case reflect.TypeOf([0]gd.Vector3{}).Elem():
			vtype = gd.TypeVector3
		case reflect.TypeOf([0]gd.Vector3i{}).Elem():
			vtype = gd.TypeVector3i
		case reflect.TypeOf([0]gd.Transform2D{}).Elem():
			vtype = gd.TypeTransform2D
		case reflect.TypeOf([0]gd.Vector4{}).Elem():
			vtype = gd.TypeVector4
		case reflect.TypeOf([0]gd.Vector4i{}).Elem():
			vtype = gd.TypeVector4i
		case reflect.TypeOf([0]gd.Plane{}).Elem():
			vtype = gd.TypePlane
		case reflect.TypeOf([0]gd.Quaternion{}).Elem():
			vtype = gd.TypeQuaternion
		case reflect.TypeOf([0]gd.AABB{}).Elem():
			vtype = gd.TypeAABB
		case reflect.TypeOf([0]gd.Basis{}).Elem():
			vtype = gd.TypeBasis
		case reflect.TypeOf([0]gd.Transform3D{}).Elem():
			vtype = gd.TypeTransform3D
		case reflect.TypeOf([0]gd.Projection{}).Elem():
			vtype = gd.TypeProjection
		case reflect.TypeOf([0]gd.Color{}).Elem():
			vtype = gd.TypeColor
		case reflect.TypeOf([0]gd.StringName{}).Elem():
			vtype = gd.TypeStringName
		case reflect.TypeOf([0]gd.NodePath{}).Elem():
			vtype = gd.TypeNodePath
		case reflect.TypeOf([0]gd.RID{}).Elem():
			vtype = gd.TypeRID
		case reflect.TypeOf([0]gd.Object{}).Elem():
			vtype = gd.TypeObject
		case reflect.TypeOf([0]gd.Callable{}).Elem():
			vtype = gd.TypeCallable
		case reflect.TypeOf([0]gd.Dictionary{}).Elem():
			vtype = gd.TypeDictionary
		case reflect.TypeOf([0]gd.Array{}).Elem():
			vtype = gd.TypeArray
		case reflect.TypeOf([0]gd.PackedByteArray{}).Elem():
			vtype = gd.TypePackedByteArray
		case reflect.TypeOf([0]gd.PackedInt32Array{}).Elem():
			vtype = gd.TypePackedInt32Array
		case reflect.TypeOf([0]gd.PackedInt64Array{}).Elem():
			vtype = gd.TypePackedInt64Array
		case reflect.TypeOf([0]gd.PackedFloat32Array{}).Elem():
			vtype = gd.TypePackedFloat32Array
		case reflect.TypeOf([0]gd.PackedFloat64Array{}).Elem():
			vtype = gd.TypePackedFloat64Array
		case reflect.TypeOf([0]gd.PackedStringArray{}).Elem():
			vtype = gd.TypePackedStringArray
		case reflect.TypeOf([0]gd.PackedVector2Array{}).Elem():
			vtype = gd.TypePackedVector2Array
		case reflect.TypeOf([0]gd.PackedVector3Array{}).Elem():
			vtype = gd.TypePackedVector3Array
		case reflect.TypeOf([0]gd.PackedColorArray{}).Elem():
			vtype = gd.TypePackedColorArray
		case reflect.TypeOf([0]unsafe.Pointer{}).Elem():
			vtype = gd.TypeNil
		case reflect.TypeOf([0]*gd.ScriptLanguageExtensionProfilingInfo{}).Elem():
			vtype = gd.TypeNil
		default:
			switch {
			case rtype.Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()):
				vtype = gd.TypeObject
			case rtype.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()):
				vtype = gd.TypeSignal
			default:
				vtype = gd.TypeDictionary
			}
		}
		return vtype, true
	default:
		return gd.TypeNil, false
	}
}
