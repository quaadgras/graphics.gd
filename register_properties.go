//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strings"

	gd "grow.graphics/gd/internal"
)

/*func attachProperties(ctx Context, className StringName, info *ClassCreationInfo, rtype reflect.Type) {
	var (
		properties     []PropertyInfo
		propertyPinner runtime.Pinner
	)
	defer propertyPinner.Unpin()
	ctx.Defer(func() {
		propertyPinner.Unpin()
	})
	var groups [][2]string
	var subgroups [][2]string
	var addProperties func(rtype reflect.Type, prefix string, subgroup bool)
	addProperties = func(rtype reflect.Type, prefix string, subgroup bool) {
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			if !field.IsExported() || field.Anonymous {
				continue
			}
			if field.Type.Kind() == reflect.Struct && field.Type.Name() == "" {
				group := [2]string{field.Name, prefix + strings.ToLower(field.Name) + "_"}
				if subgroup {
					subgroups = append(subgroups, group)
				} else {
					groups = append(groups, group)
				}
				addProperties(field.Type, strings.ToLower(field.Name)+"_", true)
			} else {
				field.Name = prefix + field.Name
				properties = append(properties, derivePropertyInfo(ctx, &propertyPinner, rtype, field))
			}
		}
	}
	addProperties(rtype, "", false)
	if len(properties) > 0 {
		propertyPinner.Pin(&properties[0])

		info.GetPropertyList.Set(func(instance cgo.Handle, length *uint32) *PropertyInfo {
			*length = uint32(len(properties))
			return &properties[0]
		})
		info.FreePropertyList.Set(func(instance cgo.Handle, properties *PropertyInfo) {})
		info.Set.Set(func(instance cgo.Handle, name StringNamePtr, value VariantPtr) bool {
			ctx := NewContext(ctx.API)
			sname := mmm.Let[StringName](ctx.Lifetime, ctx.API, *name)
			field := reflect.ValueOf(instance.Value()).Elem().FieldByName(sname.String())
			if !field.IsValid() {
				return false
			}
			var fieldPin = instance.Value().(IsClass).Pin()
			variant := mmm.Let[Variant](ctx.Lifetime, ctx.API, *value)
			ptr, ok := field.Interface().(interface {
				Pin() Context
				Free()
			})
			if ok {
				fieldPin = ptr.Pin()
				if !field.IsZero() {
					ptr.Free()
				}
			}
			field.Set(reflect.ValueOf(variant.Interface(fieldPin)))
			ctx.End()
			return true
		})
		info.Get.Set(func(instance cgo.Handle, name StringNamePtr, value VariantPtr) bool {
			ctx := NewContext(ctx.API)
			sname := mmm.Let[StringName](ctx.Lifetime, ctx.API, *name)
			field := reflect.ValueOf(instance.Value()).Elem().FieldByName(sname.String())
			if !field.IsValid() {
				return false
			}
			tmp := ctx.Variant(field.Interface())
			*value = mmm.Get(tmp)
			//mmm.MarkFree(tmp)
			ctx.End()
			return true
		})
	}

	var frame = ctx.API.NewFrame()
	for _, group := range groups {
		FrameSet[uintptr](0, frame, mmm.Get(className))
		FrameSet[uintptr](1, frame, mmm.Get(ctx.String(group[0])))
		FrameSet[uintptr](2, frame, mmm.Get(ctx.String(group[1])))
		ctx.API.ClassDB.RegisterClassPropertyGroup(ctx.API.ExtensionToken, frame.Get(0), frame.Get(1), frame.Get(2))
	}
	for _, group := range subgroups {
		FrameSet[uintptr](0, frame, mmm.Get(className))
		FrameSet[uintptr](1, frame, mmm.Get(ctx.String(group[0])))
		FrameSet[uintptr](2, frame, mmm.Get(ctx.String(group[1])))
		ctx.API.ClassDB.RegisterClassPropertySubGroup(ctx.API.ExtensionToken, frame.Get(0), frame.Get(1), frame.Get(2))
	}
	frame.Free()
}*/

func propertyOf(godot Context, field reflect.StructField) gd.PropertyInfo {
	var name = field.Name
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var hint = PropertyHintResourceType
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
	var usage = PropertyUsageStorage | PropertyUsageEditor
	if vtype == TypeNil {
		usage |= PropertyUsageNilIsVariant
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
		return rtype.Name()
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
	default:
		switch {
		case rtype.Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()):
			vtype = TypeObject
		case rtype.Implements(reflect.TypeOf([0]gd.IsArray{}).Elem()):
			vtype = TypeArray
		case rtype.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()):
			vtype = TypeSignal
		default:
			panic("gdextension.RegisterClass: unsupported property type " + rtype.String())
		}
	}
	return
}
