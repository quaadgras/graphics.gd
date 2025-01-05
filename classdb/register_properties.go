package classdb

import (
	"fmt"
	"reflect"

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
	vtype, ok := gd.VariantTypeOf(field.Type)
	if !ok {
		return gd.PropertyInfo{}, false
	}
	if vtype == gd.TypeArray {
		elem := field.Type.Elem()
		etype, ok := gd.VariantTypeOf(elem)
		if !ok {
			return gd.PropertyInfo{}, false
		}
		if elem.Implements(reflect.TypeFor[ResourceClass.Any]()) {
			hintString = fmt.Sprintf("%d/%d:%s", gd.TypeObject, PropertyHintResourceType, nameOf(elem)) // MAKE_RESOURCE_TYPE_HINT
		} else {
			hint |= PropertyHintArrayType
			hintString = etype.String()
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
