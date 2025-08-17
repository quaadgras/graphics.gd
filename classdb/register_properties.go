package classdb

import (
	"fmt"
	"reflect"

	EngineClass "graphics.gd/classdb/Engine"
	NodeClass "graphics.gd/classdb/Node"
	ResourceClass "graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"
)

func propertyOf(class gd.StringName, field reflect.StructField) (gd.PropertyInfo, bool) {
	var name = String.ToSnakeCase(field.Name)
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var vtype gd.VariantType
	var hint PropertyHint
	var hintString = nameOf(field.Type)
	var enum = registerEnumsFor(class, field.Type)
	var className = nameOf(field.Type)
	if instance, ok := field.Type.MethodByName("Instance"); ok && instance.Type.NumOut() == 2 && field.Type.Name() == "ID" {
		vtype = gd.TypeObject
		className = nameOf(instance.Type.Out(0))
		hintString = className
		field.Type = instance.Type.Out(0)
	} else {
		switch {
		case enum != nil:
			vtype = gd.TypeInt
			hint |= PropertyHintEnum
			hintString = ""
			var first = true
			for name, value := range enum {
				if !first {
					hintString += ","
				}
				hintString += fmt.Sprintf("%s:%d", name, value)
				first = false
			}
		case field.Type.Kind() == reflect.Pointer && field.Type.Implements(reflect.TypeOf([0]interface{ Super() ResourceClass.Instance }{}).Elem()):
			vtype = gd.TypeObject
			hint |= PropertyHintResourceType
			hintString = nameOf(field.Type.Elem())
		default:
			vtype, ok = gd.VariantTypeOf(field.Type)
			if !ok {
				return gd.PropertyInfo{}, false
			}
			if vtype == gd.TypeArray && (field.Type.Kind() == reflect.Array || field.Type.Kind() == reflect.Slice) {
				elem := field.Type.Elem()
				etype, ok := gd.VariantTypeOf(elem)
				if !ok {
					return gd.PropertyInfo{}, false
				}
				if elem.Implements(reflect.TypeFor[ResourceClass.Any]()) {
					hintString = fmt.Sprintf("%d/%d:%s", gd.TypeObject, PropertyHintResourceType, nameOf(elem)) // MAKE_RESOURCE_TYPE_HINT
				} else if etype != gd.TypeNil {
					hint |= PropertyHintArrayType
					hintString = etype.String()
				}
			}
			if vtype == gd.TypeArray && field.Type.Implements(reflect.TypeFor[Array.Interface]()) {
				elem := reflect.Zero(field.Type).Interface().(Array.Interface).ElemType()
				etype, ok := gd.VariantTypeOf(elem)
				if !ok {
					return gd.PropertyInfo{}, false
				}
				if etype != gd.TypeNil {
					hint |= PropertyHintArrayType
					hintString = etype.String()
				}
			}
		}
	}
	if field.Type.Implements(reflect.TypeOf([0]interface{ AsResource() ResourceClass.Instance }{}).Elem()) {
		hint |= PropertyHintResourceType
	}
	if field.Type.Implements(reflect.TypeOf([0]interface{ AsNode() NodeClass.Instance }{}).Elem()) {
		hint |= PropertyHintNodeType
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
		ClassName:  gd.NewStringName(className),
		Hint:       int64(hint),
		HintString: gd.NewString(hintString),
		Usage:      int64(usage),
	}, true
}

// Set needs to reference++ any resources that are sucessfully set.
func (instance *instanceImplementation) Set(name gd.StringName, value gd.Variant) bool {
	sname := name.String()
	if impl, ok := instance.Value.(interface {
		Set(string, any) bool
	}); ok {
		ok := bool(impl.Set(sname, value.Interface()))
		if ok {
			if impl, ok := instance.Value.(interface {
				OnSet(string, any)
			}); ok {
				impl.OnSet(sname, value.Interface())
			}
		}
		return ok
	}
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(sname)
	if !field.IsValid() {
		for _, rfield := range reflect.VisibleFields(rvalue.Type()) {
			if !rfield.IsExported() {
				continue
			}
			tag, hasTag := rfield.Tag.Lookup("gd")
			if hasTag && !rfield.Anonymous && tag == sname {
				field = rvalue.FieldByIndex(rfield.Index)
				break
			}
			if !hasTag && String.ToSnakeCase(rfield.Name) == sname {
				field = rvalue.FieldByIndex(rfield.Index)
				break
			}
		}
		if !field.IsValid() {
			return false
		}
	}
	if !field.CanSet() {
		return false
	}
	if value.Type() == gd.TypeNil {
		field.Set(reflect.Zero(field.Type()))
		return true
	}
	if reflect.PointerTo(field.Type()).Implements(reflect.TypeFor[Enum.Pointer]()) {
		if value.Type() != gd.TypeInt {
			return false
		}
		field.Addr().Interface().(Enum.Pointer).SetInt(int(value.Int()))
		return true
	}
	var isExtensionClass bool
	var converted reflect.Value
	if value.Type() == gd.TypeObject {
		obj := gd.LetVariantAsPointerType[gd.Object](value, gd.TypeObject)
		ext, ok := gd.ExtensionInstances.Load(pointers.Get(obj)[0])
		if ok {
			converted = reflect.ValueOf(ext)
			isExtensionClass = true
		}
	}
	if !converted.IsValid() {
		var err error
		converted, err = gd.ConvertToDesiredGoType(value, field.Type())
		if err != nil {
			EngineClass.Raise(err)
			return false
		}
	}
	if converted.Kind() == reflect.Array || isExtensionClass {
		if !field.IsZero() {
			freeable, ok := field.Interface().(interface{ Free() })
			if ok {
				freeable.Free()
			}
		}
		obj, ok := converted.Interface().(interface {
			AsObject() [1]gd.Object
		})
		if !ok {
			return false
		}
		ref, ok := gd.As[gd.RefCounted](obj.AsObject()[0])
		if ok {
			ref.Reference()
		}
		pointers.Pin(obj.AsObject()[0])
	}
	field.Set(converted)
	if impl, ok := instance.Value.(interface {
		OnSet(gd.StringName, gd.Variant)
	}); ok {
		impl.OnSet(name, value)
	}
	return true
}

func (instance *instanceImplementation) Get(name gd.StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		Get(string) any
	}); ok {
		return gd.NewVariant(impl.Get(name.String())), true
	}
	sname := name.String()
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(String.ToPascalCase(sname))
	if !field.IsValid() {
		for _, rfield := range reflect.VisibleFields(rvalue.Type()) {
			if !rfield.IsExported() {
				continue
			}
			tag, hasTag := rfield.Tag.Lookup("gd")
			if hasTag && !rfield.Anonymous && tag == sname {
				field = rvalue.FieldByIndex(rfield.Index)
				break
			}
			if !hasTag && String.ToSnakeCase(rfield.Name) == sname {
				field = rvalue.FieldByIndex(rfield.Index)
				break
			}
		}
		if !field.IsValid() {
			return gd.Variant{}, false
		}
		if field.Type().Kind() == reflect.Chan || reflect.PointerTo(field.Type()).Implements(reflect.TypeFor[Signal.Pointer]()) {
			return gd.Variant{}, false
		}
	}
	if field.Type().Implements(reflect.TypeFor[interface{ superType() reflect.Type }]()) {
		if field.IsZero() {
			return gd.Variant{}, false
		}
		obj := field.Addr().Interface().(interface{ AsObject() [1]gd.Object }).AsObject()[0]
		vary := gd.NewVariant(obj)
		return vary, true
	}
	return gd.NewVariant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList() []gd.PropertyInfo {
	if impl, ok := instance.Value.(interface {
		GetPropertyList() []Object.PropertyInfo
	}); ok {
		var list = impl.GetPropertyList()
		var results = make([]gd.PropertyInfo, len(list))
		for i, info := range list {
			vtype, _ := gd.VariantTypeOf(info.Type)
			results[i] = gd.PropertyInfo{
				Type:       vtype,
				Name:       gd.NewStringName(info.Name),
				ClassName:  gd.NewStringName(info.ClassName),
				HintString: gd.NewString(info.HintString),
				Usage:      int64(info.Usage),
				Hint:       int64(info.Hint),
			}
		}
		return results
	}
	return nil
}

func (instance *instanceImplementation) PropertyCanRevert(name gd.StringName) bool {
	if impl, ok := instance.Value.(interface {
		PropertyCanRevert(string) bool
	}); ok {
		return bool(impl.PropertyCanRevert(name.String()))
	}
	sname := name.String()
	rtype := reflect.TypeOf(instance.Value).Elem()
	field, ok := rtype.FieldByName(sname)
	if !ok {
		for _, rfield := range reflect.VisibleFields(rtype) {
			if rfield.Tag.Get("gd") == sname {
				field = rtype.FieldByIndex(rfield.Index)
				ok = true
				break
			}
		}
	}
	if !ok {
		return false
	}
	_, ok = field.Tag.Lookup("default")
	return ok
}
func (instance *instanceImplementation) PropertyGetRevert(name gd.StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		PropertyGetRevert(string) (any, bool)
	}); ok {
		val, ok := impl.PropertyGetRevert(name.String())
		return gd.NewVariant(val), ok
	}
	sname := name.String()
	rtype := reflect.TypeOf(instance.Value).Elem()
	field, ok := rtype.FieldByName(sname)
	if !ok {
		for _, rfield := range reflect.VisibleFields(rtype) {
			if rfield.Tag.Get("gd") == sname {
				field = rtype.FieldByIndex(rfield.Index)
				ok = true
				break
			}
		}
	}
	if !ok {
		return gd.Variant{}, false
	}
	var value = reflect.New(field.Type)
	if tag := field.Tag.Get("default"); tag != "" {
		_, err := fmt.Sscanf(tag, "%v", value.Interface())
		if err != nil {
			return gd.Variant{}, false
		}
	}
	return gd.NewVariant(value.Elem().Interface()), true
}

func (instance *instanceImplementation) ValidateProperty(info *gd.PropertyInfo) bool {
	switch validate := instance.Value.(type) {
	case interface {
		ValidateProperty(Object.PropertyInfo) bool
	}:
		return bool(validate.ValidateProperty(Object.PropertyInfo{
			ClassName:  info.ClassName.String(),
			Usage:      int(info.Usage),
			Type:       gd.ConvieniantGoTypeOf(info.Type),
			HintString: info.HintString.String(),
			Hint:       int(info.Hint),
			Name:       info.Name.String(),
		}))
	}
	return true
}
