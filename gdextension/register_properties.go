package gdextension

import (
	"reflect"
	"runtime"
	"runtime/cgo"
	"strings"

	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

func attachProperties(ctx gd.Context, className gd.StringName, info *gd.ClassCreationInfo, rtype reflect.Type) {
	var (
		properties     []gd.PropertyInfo
		propertyPinner runtime.Pinner
	)
	defer propertyPinner.Unpin()
	/*ctx.Defer(func() {
		propertyPinner.Unpin()
	})*/
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

		info.GetPropertyList.Set(func(instance cgo.Handle, length *uint32) *gd.PropertyInfo {
			*length = uint32(len(properties))
			return &properties[0]
		})
		info.FreePropertyList.Set(func(instance cgo.Handle, properties *gd.PropertyInfo) {})
		info.Set.Set(func(instance cgo.Handle, name gd.StringNamePtr, value gd.VariantPtr) bool {
			ctx := gd.NewContext(ctx.API)
			sname := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *name)
			field := reflect.ValueOf(instance.Value()).Elem().FieldByName(sname.String())
			if !field.IsValid() {
				return false
			}
			var fieldPin = instance.Value().(gd.IsClass).Pin()
			variant := mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, *value)
			ptr, ok := field.Interface().(interface {
				Pin() gd.Context
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
		info.Get.Set(func(instance cgo.Handle, name gd.StringNamePtr, value gd.VariantPtr) bool {
			ctx := gd.NewContext(ctx.API)
			sname := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *name)
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
		gd.FrameSet[uintptr](0, frame, mmm.Get(className))
		gd.FrameSet[uintptr](1, frame, mmm.Get(ctx.String(group[0])))
		gd.FrameSet[uintptr](2, frame, mmm.Get(ctx.String(group[1])))
		ctx.API.ClassDB.RegisterClassPropertyGroup(ctx.API.ExtensionToken, frame.Get(0), frame.Get(1), frame.Get(2))
	}
	for _, group := range subgroups {
		gd.FrameSet[uintptr](0, frame, mmm.Get(className))
		gd.FrameSet[uintptr](1, frame, mmm.Get(ctx.String(group[0])))
		gd.FrameSet[uintptr](2, frame, mmm.Get(ctx.String(group[1])))
		ctx.API.ClassDB.RegisterClassPropertySubGroup(ctx.API.ExtensionToken, frame.Get(0), frame.Get(1), frame.Get(2))
	}
	frame.Free()
}

func derivePropertyInfo(godot gd.Context, propertyPinner *runtime.Pinner, rtype reflect.Type, field reflect.StructField) gd.PropertyInfo {
	if field.Name == "" {
		field.Name = strings.ToLower(field.Type.Name())
	}
	var vtype gd.VariantType
	var name = mmm.Get(godot.StringName(field.Name))
	var namePtr = &name
	propertyPinner.Pin(namePtr)
	var class gd.StringNamePtr
	if field.Type.Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()) {
		className := mmm.Get(godot.StringName(field.Type.Name()))
		vtype = gd.TypeObject
		class = &className
		propertyPinner.Pin(class)
	} else {
		className := mmm.Get(godot.StringName(""))
		class = &className
		propertyPinner.Pin(class)
	}
	if field.Type.Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()) {
		vtype = gd.TypeSignal
	}

	var HintString = mmm.Get(godot.String(""))
	var HintStringPtr = &HintString
	propertyPinner.Pin(HintStringPtr)

	if vtype == gd.TypeNil {
		switch field.Type {
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
		default:
			panic("gdextension.RegisterClass: unsupported property type " + field.Type.String())
		}
	}
	return gd.PropertyInfo{
		Type:       uint32(vtype),
		Name:       namePtr,
		ClassName:  class,
		HintString: HintStringPtr,
	}
}
