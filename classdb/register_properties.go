package classdb

import (
	"fmt"
	"reflect"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RefCounted"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"
)

func propertyOf(class gd.StringName, field reflect.StructField, push_into gdextension.PropertyList) bool {
	var name = String.ToSnakeCase(field.Name)
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var vtype gdextension.VariantType
	var hint PropertyHint
	var hintString = nameOf(field.Type)
	var enum = registerEnumsFor(class, field.Type)
	var className = nameOf(field.Type)
	if instance, ok := field.Type.MethodByName("Instance"); ok && instance.Type.NumOut() == 2 && field.Type.Name() == "ID" {
		vtype = gdextension.TypeObject
		className = nameOf(instance.Type.Out(0))
		hintString = className
		field.Type = instance.Type.Out(0)
	} else {
		switch {
		case enum != nil:
			vtype = gdextension.TypeInt
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
		case field.Type.Kind() == reflect.Pointer && field.Type.Implements(reflect.TypeFor[[0]interface{ Super() Resource.Instance }]().Elem()):
			vtype = gdextension.TypeObject
			hint |= PropertyHintResourceType
			hintString = nameOf(field.Type.Elem())
		default:
			vtype, ok = gd.VariantTypeOf(field.Type)
			if !ok {
				return false
			}
			if vtype == gdextension.TypeObject {
				if field.Type.Implements(reflect.TypeFor[Resource.Any]()) {
					hintString = fmt.Sprintf("%d/%d:%s", gdextension.TypeObject, PropertyHintResourceType, nameOf(field.Type)) // MAKE_RESOURCE_TYPE_HINT
				} else {
					hintString = nameOf(field.Type)
				}
			}
			if vtype == gdextension.TypeArray {
				if field.Type.Implements(reflect.TypeFor[Array.Interface]()) {
					elem := reflect.Zero(field.Type).Interface().(Array.Interface).ElemType()
					etype, ok := gd.VariantTypeOf(elem)
					if !ok {
						return false
					}
					if etype != gdextension.TypeNil {
						hint |= PropertyHintArrayType
						hintString = etype.String()
					}
				} else {
					etype, ok := gd.VariantTypeOf(field.Type.Elem())
					if !ok {
						return false
					}
					hint |= PropertyHintArrayType
					hintString = etype.String()
				}
			}
		}
	}
	if field.Type.Implements(reflect.TypeFor[[0]interface{ AsResource() Resource.Instance }]().Elem()) {
		hint |= PropertyHintResourceType
	}
	if field.Type.Implements(reflect.TypeFor[[0]interface{ AsNode() Node.Instance }]().Elem()) {
		hint |= PropertyHintNodeType
	}
	var usage = PropertyUsageStorage | PropertyUsageEditor
	if vtype == gdextension.TypeNil {
		usage |= PropertyUsageNilIsVariant
	}
	if rangeHint, ok := field.Tag.Lookup("range"); ok {
		hint |= PropertyHintRange
		hintString = rangeHint
	}
	gdextension.Host.ClassDB.PropertyList.Push(push_into,
		vtype,
		pointers.Get(gd.NewStringName(name)),
		pointers.Get(gd.NewStringName(className)),
		uint32(hint),
		pointers.Get(gd.NewString(hintString)),
		uint32(usage),
		0,
	)
	return true
}

// Set needs to reference++ any resources that are sucessfully set.
func (instance *instanceImplementation) Set(name gd.StringName, value gd.Variant) bool {
	val, ok := instance.Interface()
	if !ok {
		return false
	}
	sname := name.String()
	if impl, ok := val.(interface {
		Set(string, any) bool
	}); ok {
		ok := bool(impl.Set(sname, value.Interface()))
		if ok {
			if impl, ok := val.(interface {
				OnSet(string, any)
			}); ok {
				impl.OnSet(sname, value.Interface())
			}
		}
		return ok
	}
	rvalue := reflect.ValueOf(val).Elem()
	field := rvalue.FieldByName(sname)
	if !field.IsValid() {
		for _, rfield := range reflect.VisibleFields(rvalue.Type()) {
			if !rfield.IsExported() {
				continue
			}
			tag, hasTag := rfield.Tag.Lookup("gd")
			if tag == "-" {
				return false
			}
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
	if value.Type() == gdextension.TypeNil {
		field.Set(reflect.Zero(field.Type()))
		return true
	}
	if reflect.PointerTo(field.Type()).Implements(reflect.TypeFor[Enum.Pointer]()) {
		if value.Type() != gdextension.TypeInt {
			return false
		}
		field.Addr().Interface().(Enum.Pointer).SetInt(int(value.Int()))
		return true
	}
	var isExtensionClass bool
	var converted reflect.Value
	if value.Type() == gdextension.TypeObject && field.Kind() != reflect.Uint64 { // support setting Object.ID fields with Object
		obj := gd.VariantAsObject(value)
		ext := gd.ExtensionInstanceLookup(gdreference.GetObject(obj))
		if ext != nil {
			converted = reflect.ValueOf(ext)
			isExtensionClass = true
		}
	}
	if !converted.IsValid() {
		var err error
		converted, err = gd.ConvertToDesiredGoType(value, field.Type())
		if err != nil {
			Engine.Raise(err)
			return false
		}
	}
	if converted.Kind() == reflect.Array || isExtensionClass {
		if !field.IsZero() {
			// we need to unreference any existing pinned resources (pinned here means that the engine `set` them).
			if ref, ok := field.Interface().(RefCounted.Any); ok {
				ref := ref.AsRefCounted()[0]
				if ref.Unreference() {
					gdextension.Host.Objects.Unsafe.Free(gdreference.GetObject(gdreference.Object(ref)))
				}
			}
		}
		obj, ok := converted.Interface().(gd.IsClass)
		if !ok {
			return false
		}
		if !isExtensionClass {
			ref, ok := Object.As[gd.RefCounted](obj)
			if ok {
				ref.Reference()
			}
		}
	}
	field.Set(converted)
	if impl, ok := val.(interface {
		OnSet(string, any)
	}); ok {
		impl.OnSet(name.String(), value)
	}
	return true
}

func (instance *instanceImplementation) Get(name gd.StringName) (gd.Variant, bool) {
	val, ok := instance.Interface()
	if !ok {
		return gd.Variant{}, false
	}
	if impl, ok := val.(interface {
		Get(string) any
	}); ok {
		return gd.NewVariant(impl.Get(name.String())), true
	}
	sname := name.String()
	rvalue := reflect.ValueOf(val).Elem()
	field := rvalue.FieldByName(String.ToPascalCase(sname))
	if !field.IsValid() {
		for _, rfield := range reflect.VisibleFields(rvalue.Type()) {
			if !rfield.IsExported() {
				continue
			}
			tag, hasTag := rfield.Tag.Lookup("gd")
			if tag == "-" {
				return gd.Variant{}, false
			}
			if hasTag && !rfield.Anonymous && tag == sname {
				field = rvalue.FieldByIndex(rfield.Index)
				break
			}
			if !hasTag && String.ToSnakeCase(rfield.Name) == sname {
				field = rvalue.FieldByIndex(rfield.Index)
				break
			}
		}
	}
	if !field.IsValid() {
		return gd.Variant{}, false
	}
	if field.Type().Kind() == reflect.Chan || reflect.PointerTo(field.Type()).Implements(reflect.TypeFor[Signal.Pointer]()) {
		return gd.Variant{}, false
	}
	return gd.NewVariant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList() gdextension.PropertyList {
	val, ok := instance.Interface()
	if !ok {
		return 0
	}
	if impl, ok := val.(interface {
		GetPropertyList() []Object.PropertyInfo
	}); ok {
		var list = impl.GetPropertyList()
		var results = gdextension.Host.ClassDB.PropertyList.Make(len(list))
		for _, info := range list {
			vtype, _ := gd.VariantTypeOf(info.Type)
			gdextension.Host.ClassDB.PropertyList.Push(results,
				vtype,
				pointers.Get(gd.NewStringName(info.Name)),
				pointers.Get(gd.NewStringName(info.ClassName)),
				uint32(info.Hint),
				pointers.Get(gd.NewString(info.HintString)),
				uint32(info.Usage),
				0,
			)
		}
		return results
	}
	return 0
}

func (instance *instanceImplementation) PropertyCanRevert(name gd.StringName) bool {
	val, ok := instance.Interface()
	if !ok {
		return false
	}
	if impl, ok := val.(interface {
		PropertyCanRevert(string) bool
	}); ok {
		return bool(impl.PropertyCanRevert(name.String()))
	}
	sname := name.String()
	rtype := reflect.TypeOf(val).Elem()
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
	val, ok := instance.Interface()
	if !ok {
		return gd.Variant{}, false
	}
	if impl, ok := val.(interface {
		PropertyGetRevert(string) (any, bool)
	}); ok {
		val, ok := impl.PropertyGetRevert(name.String())
		return gd.NewVariant(val), ok
	}
	sname := name.String()
	rtype := reflect.TypeOf(val).Elem()
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

func (instance *instanceImplementation) ValidateProperty(list gdextension.PropertyList) bool {
	val, ok := instance.Interface()
	if !ok {
		return false
	}
	switch validate := val.(type) {
	case interface {
		ValidateProperty(Object.PropertyInfo) bool
	}:
		return bool(validate.ValidateProperty(Object.PropertyInfo{
			ClassName:  pointers.Raw[gd.StringName](gdextension.Host.ClassDB.PropertyList.Info.ClassName(list)).String(),
			Usage:      int(gdextension.Host.ClassDB.PropertyList.Info.Usage(list)),
			Type:       gd.ConvieniantGoTypeOf(gdextension.Host.ClassDB.PropertyList.Info.Type(list)),
			HintString: pointers.Raw[gd.String](gdextension.Host.ClassDB.PropertyList.Info.HinString(list)).String(),
			Hint:       int(gdextension.Host.ClassDB.PropertyList.Info.Hint(list)),
			Name:       pointers.Raw[gd.StringName](gdextension.Host.ClassDB.PropertyList.Info.Name(list)).String(),
		}))
	}
	return true
}
