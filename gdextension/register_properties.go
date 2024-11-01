package gdextension

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"grow.graphics/gd/gdconst"
	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/gd/objects/Resource"
)

func propertyOf(field reflect.StructField) gd.PropertyInfo {
	var name = field.Name
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var hint gdconst.PropertyHint
	var hintString = classNameOf(field.Type)
	vtype := variantTypeOf(field.Type)
	if vtype == gd.TypeArray {
		_, generic, ok := strings.Cut(field.Type.String(), "[")
		if ok {
			hint |= gdconst.PropertyHintArrayType
			split := strings.Split(generic, ".")
			elem := split[len(split)-1]
			elem = elem[:len(elem)-1]
			hintString = fmt.Sprintf("%d/%d:%s", gd.TypeObject, gdconst.PropertyHintResourceType, elem) // MAKE_RESOURCE_TYPE_HINT
		}
	}
	if field.Type.Implements(reflect.TypeOf([0]interface{ AsResource() Resource.Instance }{}).Elem()) {
		hint |= gdconst.PropertyHintResourceType
	}
	var usage = gdconst.PropertyUsageStorage | gdconst.PropertyUsageEditor
	if vtype == gd.TypeNil {
		usage |= gdconst.PropertyUsageNilIsVariant
	}
	if rangeHint, ok := field.Tag.Lookup("range"); ok {
		hint |= gdconst.PropertyHintRange
		hintString = rangeHint
	}
	return gd.PropertyInfo{
		Type:       vtype,
		Name:       gd.NewStringName(name),
		ClassName:  gd.NewStringName(classNameOf(field.Type)),
		Hint:       hint,
		HintString: gd.NewString(hintString),
		Usage:      usage,
	}
}

func variantTypeOf(rtype reflect.Type) (vtype gd.VariantType) {
	switch rtype.Kind() {
	case reflect.Int32, reflect.Int64, reflect.Int:
		return gd.TypeInt
	}
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
		vtype = gd.TypeTransform2d
	case reflect.TypeOf([0]gd.Vector4{}).Elem():
		vtype = gd.TypeVector4
	case reflect.TypeOf([0]gd.Vector4i{}).Elem():
		vtype = gd.TypeVector4i
	case reflect.TypeOf([0]gd.Plane{}).Elem():
		vtype = gd.TypePlane
	case reflect.TypeOf([0]gd.Quaternion{}).Elem():
		vtype = gd.TypeQuaternion
	case reflect.TypeOf([0]gd.AABB{}).Elem():
		vtype = gd.TypeAabb
	case reflect.TypeOf([0]gd.Basis{}).Elem():
		vtype = gd.TypeBasis
	case reflect.TypeOf([0]gd.Transform3D{}).Elem():
		vtype = gd.TypeTransform3d
	case reflect.TypeOf([0]gd.Projection{}).Elem():
		vtype = gd.TypeProjection
	case reflect.TypeOf([0]gd.Color{}).Elem():
		vtype = gd.TypeColor
	case reflect.TypeOf([0]gd.StringName{}).Elem():
		vtype = gd.TypeStringName
	case reflect.TypeOf([0]gd.NodePath{}).Elem():
		vtype = gd.TypeNodePath
	case reflect.TypeOf([0]gd.RID{}).Elem():
		vtype = gd.TypeRid
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
	case reflect.TypeOf([0]*classdb.ScriptLanguageExtensionProfilingInfo{}).Elem():
		vtype = gd.TypeNil
	default:
		switch {
		case rtype.Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()):
			vtype = gd.TypeObject
		case rtype.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()):
			vtype = gd.TypeSignal
		case rtype.ConvertibleTo(reflect.TypeOf(gd.Dictionary{})):
			vtype = gd.TypeDictionary
		default:
			switch rtype {
			case reflect.TypeOf([0]string{}).Elem():
				vtype = gd.TypeString
			case reflect.TypeOf([0]int{}).Elem():
				vtype = gd.TypeInt
			default:
				panic("gdextension.RegisterClass: unsupported property type " + rtype.String())
			}
		}
	}
	return
}
