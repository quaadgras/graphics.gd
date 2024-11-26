package EditorInspectorPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[EditorInspectorPlugin] allows adding custom property editors to [EditorInspector].
When an object is edited, the [method _can_handle] function is called and must return [code]true[/code] if the object type is supported.
If supported, the function [method _parse_begin] will be called, allowing to place custom controls at the beginning of the class.
Subsequently, the [method _parse_category] and [method _parse_property] are called for every category and property. They offer the ability to add custom controls to the inspector too.
Finally, [method _parse_end] will be called.
On each of these calls, the "add" functions can be called.
To use [EditorInspectorPlugin], register it using the [method EditorPlugin.add_inspector_plugin] method first.

	// EditorInspectorPlugin methods that can be overridden by a [Class] that extends it.
	type EditorInspectorPlugin interface {
		//Returns [code]true[/code] if this object can be handled by this plugin.
		CanHandle(obj gd.Object) bool
		//Called to allow adding controls at the beginning of the property list for [param object].
		ParseBegin(obj gd.Object)
		//Called to allow adding controls at the beginning of a category in the property list for [param object].
		ParseCategory(obj gd.Object, category string)
		//Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
		ParseGroup(obj gd.Object, group string)
		//Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
		ParseProperty(obj gd.Object, atype gd.VariantType, name string, hint_type PropertyHint, hint_string string, usage_flags PropertyUsageFlags, wide bool) bool
		//Called to allow adding controls at the end of the property list for [param object].
		ParseEnd(obj gd.Object)
	}
*/
type Instance [1]classdb.EditorInspectorPlugin

/*
Returns [code]true[/code] if this object can be handled by this plugin.
*/
func (Instance) _can_handle(impl func(ptr unsafe.Pointer, obj gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the beginning of the property list for [param object].
*/
func (Instance) _parse_begin(impl func(ptr unsafe.Pointer, obj gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Called to allow adding controls at the beginning of a category in the property list for [param object].
*/
func (Instance) _parse_category(impl func(ptr unsafe.Pointer, obj gd.Object, category string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var category = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(category)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj, category.String())
	}
}

/*
Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
*/
func (Instance) _parse_group(impl func(ptr unsafe.Pointer, obj gd.Object, group string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var group = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(group)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj, group.String())
	}
}

/*
Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
*/
func (Instance) _parse_property(impl func(ptr unsafe.Pointer, obj gd.Object, atype gd.VariantType, name string, hint_type PropertyHint, hint_string string, usage_flags PropertyUsageFlags, wide bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var atype = gd.UnsafeGet[gd.VariantType](p_args, 1)
		var name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(name)
		var hint_type = gd.UnsafeGet[PropertyHint](p_args, 3)
		var hint_string = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 4))
		defer pointers.End(hint_string)
		var usage_flags = gd.UnsafeGet[PropertyUsageFlags](p_args, 5)
		var wide = gd.UnsafeGet[bool](p_args, 6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, atype, name.String(), hint_type, hint_string.String(), usage_flags, wide)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the end of the property list for [param object].
*/
func (Instance) _parse_end(impl func(ptr unsafe.Pointer, obj gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Adds a custom control, which is not necessarily a property editor.
*/
func (self Instance) AddCustomControl(control objects.Control) {
	class(self).AddCustomControl(control)
}

/*
Adds a property editor for an individual property. The [param editor] control must extend [EditorProperty].
There can be multiple property editors for a property. If [param add_to_end] is [code]true[/code], this newly added editor will be displayed after all the other editors of the property whose [param add_to_end] is [code]false[/code]. For example, the editor uses this parameter to add an "Edit Region" button for [member Sprite2D.region_rect] below the regular [Rect2] editor.
[param label] can be used to choose a custom label for the property editor in the inspector. If left empty, the label is computed from the name of the property instead.
*/
func (self Instance) AddPropertyEditor(property string, editor objects.Control) {
	class(self).AddPropertyEditor(gd.NewString(property), editor, false, gd.NewString(""))
}

/*
Adds an editor that allows modifying multiple properties. The [param editor] control must extend [EditorProperty].
*/
func (self Instance) AddPropertyEditorForMultipleProperties(label string, properties []string, editor objects.Control) {
	class(self).AddPropertyEditorForMultipleProperties(gd.NewString(label), gd.NewPackedStringSlice(properties), editor)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorInspectorPlugin

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorInspectorPlugin"))
	return Instance{classdb.EditorInspectorPlugin(object)}
}

/*
Returns [code]true[/code] if this object can be handled by this plugin.
*/
func (class) _can_handle(impl func(ptr unsafe.Pointer, obj gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the beginning of the property list for [param object].
*/
func (class) _parse_begin(impl func(ptr unsafe.Pointer, obj gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Called to allow adding controls at the beginning of a category in the property list for [param object].
*/
func (class) _parse_category(impl func(ptr unsafe.Pointer, obj gd.Object, category gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var category = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj, category)
	}
}

/*
Called to allow adding controls at the beginning of a group or a sub-group in the property list for [param object].
*/
func (class) _parse_group(impl func(ptr unsafe.Pointer, obj gd.Object, group gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var group = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj, group)
	}
}

/*
Called to allow adding property-specific editors to the property list for [param object]. The added editor control must extend [EditorProperty]. Returning [code]true[/code] removes the built-in editor for this property, otherwise allows to insert a custom editor before the built-in one.
*/
func (class) _parse_property(impl func(ptr unsafe.Pointer, obj gd.Object, atype gd.VariantType, name gd.String, hint_type PropertyHint, hint_string gd.String, usage_flags PropertyUsageFlags, wide bool) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		var atype = gd.UnsafeGet[gd.VariantType](p_args, 1)
		var name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var hint_type = gd.UnsafeGet[PropertyHint](p_args, 3)
		var hint_string = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 4))
		var usage_flags = gd.UnsafeGet[PropertyUsageFlags](p_args, 5)
		var wide = gd.UnsafeGet[bool](p_args, 6)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, atype, name, hint_type, hint_string, usage_flags, wide)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to allow adding controls at the end of the property list for [param object].
*/
func (class) _parse_end(impl func(ptr unsafe.Pointer, obj gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Adds a custom control, which is not necessarily a property editor.
*/
//go:nosplit
func (self class) AddCustomControl(control objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspectorPlugin.Bind_add_custom_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a property editor for an individual property. The [param editor] control must extend [EditorProperty].
There can be multiple property editors for a property. If [param add_to_end] is [code]true[/code], this newly added editor will be displayed after all the other editors of the property whose [param add_to_end] is [code]false[/code]. For example, the editor uses this parameter to add an "Edit Region" button for [member Sprite2D.region_rect] below the regular [Rect2] editor.
[param label] can be used to choose a custom label for the property editor in the inspector. If left empty, the label is computed from the name of the property instead.
*/
//go:nosplit
func (self class) AddPropertyEditor(property gd.String, editor objects.Control, add_to_end bool, label gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(property))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(editor[0])))
	callframe.Arg(frame, add_to_end)
	callframe.Arg(frame, pointers.Get(label))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspectorPlugin.Bind_add_property_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an editor that allows modifying multiple properties. The [param editor] control must extend [EditorProperty].
*/
//go:nosplit
func (self class) AddPropertyEditorForMultipleProperties(label gd.String, properties gd.PackedStringArray, editor objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(properties))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(editor[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorInspectorPlugin.Bind_add_property_editor_for_multiple_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsEditorInspectorPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorInspectorPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted          { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted       { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_can_handle":
		return reflect.ValueOf(self._can_handle)
	case "_parse_begin":
		return reflect.ValueOf(self._parse_begin)
	case "_parse_category":
		return reflect.ValueOf(self._parse_category)
	case "_parse_group":
		return reflect.ValueOf(self._parse_group)
	case "_parse_property":
		return reflect.ValueOf(self._parse_property)
	case "_parse_end":
		return reflect.ValueOf(self._parse_end)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_can_handle":
		return reflect.ValueOf(self._can_handle)
	case "_parse_begin":
		return reflect.ValueOf(self._parse_begin)
	case "_parse_category":
		return reflect.ValueOf(self._parse_category)
	case "_parse_group":
		return reflect.ValueOf(self._parse_group)
	case "_parse_property":
		return reflect.ValueOf(self._parse_property)
	case "_parse_end":
		return reflect.ValueOf(self._parse_end)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorInspectorPlugin", func(ptr gd.Object) any { return classdb.EditorInspectorPlugin(ptr) })
}

type PropertyHint int

const (
	/*The property has no hint for the editor.*/
	PropertyHintNone PropertyHint = 0
	/*Hints that an [int] or [float] property should be within a range specified via the hint string [code]"min,max"[/code] or [code]"min,max,step"[/code]. The hint string can optionally include [code]"or_greater"[/code] and/or [code]"or_less"[/code] to allow manual input going respectively above the max or below the min values.
	  [b]Example:[/b] [code]"-360,360,1,or_greater,or_less"[/code].
	  Additionally, other keywords can be included: [code]"exp"[/code] for exponential range editing, [code]"radians_as_degrees"[/code] for editing radian angles in degrees (the range values are also in degrees), [code]"degrees"[/code] to hint at an angle and [code]"hide_slider"[/code] to hide the slider.*/
	PropertyHintRange PropertyHint = 1
	/*Hints that an [int] or [String] property is an enumerated value to pick in a list specified via a hint string.
	  The hint string is a comma separated list of names such as [code]"Hello,Something,Else"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. For integer properties, the first name in the list has value 0, the next 1, and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"Zero,One,Three:3,Four,Six:6"[/code].*/
	PropertyHintEnum PropertyHint = 2
	/*Hints that a [String] property can be an enumerated value to pick in a list specified via a hint string such as [code]"Hello,Something,Else"[/code].
	  Unlike [constant PROPERTY_HINT_ENUM], a property with this hint still accepts arbitrary values and can be empty. The list of values serves to suggest possible values.*/
	PropertyHintEnumSuggestion PropertyHint = 3
	/*Hints that a [float] property should be edited via an exponential easing function. The hint string can include [code]"attenuation"[/code] to flip the curve horizontally and/or [code]"positive_only"[/code] to exclude in/out easing and limit values to be greater than or equal to zero.*/
	PropertyHintExpEasing PropertyHint = 4
	/*Hints that a vector property should allow its components to be linked. For example, this allows [member Vector2.x] and [member Vector2.y] to be edited together.*/
	PropertyHintLink PropertyHint = 5
	/*Hints that an [int] property is a bitmask with named bit flags.
	  The hint string is a comma separated list of names such as [code]"Bit0,Bit1,Bit2,Bit3"[/code]. Whitespaces are [b]not[/b] removed from either end of a name. The first name in the list has value 1, the next 2, then 4, 8, 16 and so on. Explicit values can also be specified by appending [code]:integer[/code] to the name, e.g. [code]"A:4,B:8,C:16"[/code]. You can also combine several flags ([code]"A:4,B:8,AB:12,C:16"[/code]).
	  [b]Note:[/b] A flag value must be at least [code]1[/code] and at most [code]2 ** 32 - 1[/code].
	  [b]Note:[/b] Unlike [constant PROPERTY_HINT_ENUM], the previous explicit value is not taken into account. For the hint [code]"A:16,B,C"[/code], A is 16, B is 2, C is 4.*/
	PropertyHintFlags PropertyHint = 6
	/*Hints that an [int] property is a bitmask using the optionally named 2D render layers.*/
	PropertyHintLayers2dRender PropertyHint = 7
	/*Hints that an [int] property is a bitmask using the optionally named 2D physics layers.*/
	PropertyHintLayers2dPhysics PropertyHint = 8
	/*Hints that an [int] property is a bitmask using the optionally named 2D navigation layers.*/
	PropertyHintLayers2dNavigation PropertyHint = 9
	/*Hints that an [int] property is a bitmask using the optionally named 3D render layers.*/
	PropertyHintLayers3dRender PropertyHint = 10
	/*Hints that an [int] property is a bitmask using the optionally named 3D physics layers.*/
	PropertyHintLayers3dPhysics PropertyHint = 11
	/*Hints that an [int] property is a bitmask using the optionally named 3D navigation layers.*/
	PropertyHintLayers3dNavigation PropertyHint = 12
	/*Hints that an integer property is a bitmask using the optionally named avoidance layers.*/
	PropertyHintLayersAvoidance PropertyHint = 37
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code].*/
	PropertyHintFile PropertyHint = 13
	/*Hints that a [String] property is a path to a directory. Editing it will show a file dialog for picking the path.*/
	PropertyHintDir PropertyHint = 14
	/*Hints that a [String] property is an absolute path to a file outside the project folder. Editing it will show a file dialog for picking the path. The hint string can be a set of filters with wildcards, like [code]"*.png,*.jpg"[/code].*/
	PropertyHintGlobalFile PropertyHint = 15
	/*Hints that a [String] property is an absolute path to a directory outside the project folder. Editing it will show a file dialog for picking the path.*/
	PropertyHintGlobalDir PropertyHint = 16
	/*Hints that a property is an instance of a [Resource]-derived type, optionally specified via the hint string (e.g. [code]"Texture2D"[/code]). Editing it will show a popup menu of valid resource types to instantiate.*/
	PropertyHintResourceType PropertyHint = 17
	/*Hints that a [String] property is text with line breaks. Editing it will show a text input field where line breaks can be typed.*/
	PropertyHintMultilineText PropertyHint = 18
	/*Hints that a [String] property is an [Expression].*/
	PropertyHintExpression PropertyHint = 19
	/*Hints that a [String] property should show a placeholder text on its input field, if empty. The hint string is the placeholder text to use.*/
	PropertyHintPlaceholderText PropertyHint = 20
	/*Hints that a [Color] property should be edited without affecting its transparency ([member Color.a] is not editable).*/
	PropertyHintColorNoAlpha PropertyHint = 21
	/*Hints that the property's value is an object encoded as object ID, with its type specified in the hint string. Used by the debugger.*/
	PropertyHintObjectId PropertyHint = 22
	/*If a property is [String], hints that the property represents a particular type (class). This allows to select a type from the create dialog. The property will store the selected type as a string.
	  If a property is [Array], hints the editor how to show elements. The [code]hint_string[/code] must encode nested types using [code]":"[/code] and [code]"/"[/code].
	  [codeblocks]
	  [gdscript]
	  # Array of elem_type.
	  hint_string = "%d:" % [elem_type]
	  hint_string = "%d/%d:%s" % [elem_type, elem_hint, elem_hint_string]
	  # Two-dimensional array of elem_type (array of arrays of elem_type).
	  hint_string = "%d:%d:" % [TYPE_ARRAY, elem_type]
	  hint_string = "%d:%d/%d:%s" % [TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
	  # Three-dimensional array of elem_type (array of arrays of arrays of elem_type).
	  hint_string = "%d:%d:%d:" % [TYPE_ARRAY, TYPE_ARRAY, elem_type]
	  hint_string = "%d:%d:%d/%d:%s" % [TYPE_ARRAY, TYPE_ARRAY, elem_type, elem_hint, elem_hint_string]
	  [/gdscript]
	  [csharp]
	  // Array of elemType.
	  hintString = $"{elemType:D}:";
	  hintString = $"{elemType:}/{elemHint:D}:{elemHintString}";
	  // Two-dimensional array of elemType (array of arrays of elemType).
	  hintString = $"{Variant.Type.Array:D}:{elemType:D}:";
	  hintString = $"{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
	  // Three-dimensional array of elemType (array of arrays of arrays of elemType).
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}:";
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Array:D}:{elemType:D}/{elemHint:D}:{elemHintString}";
	  [/csharp]
	  [/codeblocks]
	  Examples:
	  [codeblocks]
	  [gdscript]
	  hint_string = "%d:" % [TYPE_INT] # Array of integers.
	  hint_string = "%d/%d:1,10,1" % [TYPE_INT, PROPERTY_HINT_RANGE] # Array of integers (in range from 1 to 10).
	  hint_string = "%d/%d:Zero,One,Two" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
	  hint_string = "%d/%d:Zero,One,Three:3,Six:6" % [TYPE_INT, PROPERTY_HINT_ENUM] # Array of integers (an enum).
	  hint_string = "%d/%d:*.png" % [TYPE_STRING, PROPERTY_HINT_FILE] # Array of strings (file paths).
	  hint_string = "%d/%d:Texture2D" % [TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Array of textures.

	  hint_string = "%d:%d:" % [TYPE_ARRAY, TYPE_FLOAT] # Two-dimensional array of floats.
	  hint_string = "%d:%d/%d:" % [TYPE_ARRAY, TYPE_STRING, PROPERTY_HINT_MULTILINE_TEXT] # Two-dimensional array of multiline strings.
	  hint_string = "%d:%d/%d:-1,1,0.1" % [TYPE_ARRAY, TYPE_FLOAT, PROPERTY_HINT_RANGE] # Two-dimensional array of floats (in range from -1 to 1).
	  hint_string = "%d:%d/%d:Texture2D" % [TYPE_ARRAY, TYPE_OBJECT, PROPERTY_HINT_RESOURCE_TYPE] # Two-dimensional array of textures.
	  [/gdscript]
	  [csharp]
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Range:D}:1,10,1"; // Array of integers (in range from 1 to 10).
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Two"; // Array of integers (an enum).
	  hintString = $"{Variant.Type.Int:D}/{PropertyHint.Enum:D}:Zero,One,Three:3,Six:6"; // Array of integers (an enum).
	  hintString = $"{Variant.Type.String:D}/{PropertyHint.File:D}:*.png"; // Array of strings (file paths).
	  hintString = $"{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Array of textures.

	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}:"; // Two-dimensional array of floats.
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.String:D}/{PropertyHint.MultilineText:D}:"; // Two-dimensional array of multiline strings.
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Float:D}/{PropertyHint.Range:D}:-1,1,0.1"; // Two-dimensional array of floats (in range from -1 to 1).
	  hintString = $"{Variant.Type.Array:D}:{Variant.Type.Object:D}/{PropertyHint.ResourceType:D}:Texture2D"; // Two-dimensional array of textures.
	  [/csharp]
	  [/codeblocks]
	  [b]Note:[/b] The trailing colon is required for properly detecting built-in types.*/
	PropertyHintTypeString           PropertyHint = 23
	PropertyHintNodePathToEditedNode PropertyHint = 24
	/*Hints that an object is too big to be sent via the debugger.*/
	PropertyHintObjectTooBig PropertyHint = 25
	/*Hints that the hint string specifies valid node types for property of type [NodePath].*/
	PropertyHintNodePathValidTypes PropertyHint = 26
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the project's directory. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintSaveFile PropertyHint = 27
	/*Hints that a [String] property is a path to a file. Editing it will show a file dialog for picking the path for the file to be saved at. The dialog has access to the entire filesystem. The hint string can be a set of filters with wildcards like [code]"*.png,*.jpg"[/code]. See also [member FileDialog.filters].*/
	PropertyHintGlobalSaveFile PropertyHint = 28
	PropertyHintIntIsObjectid  PropertyHint = 29
	/*Hints that an [int] property is a pointer. Used by GDExtension.*/
	PropertyHintIntIsPointer PropertyHint = 30
	/*Hints that a property is an [Array] with the stored type specified in the hint string.*/
	PropertyHintArrayType PropertyHint = 31
	/*Hints that a string property is a locale code. Editing it will show a locale dialog for picking language and country.*/
	PropertyHintLocaleId PropertyHint = 32
	/*Hints that a dictionary property is string translation map. Dictionary keys are locale codes and, values are translated strings.*/
	PropertyHintLocalizableString PropertyHint = 33
	/*Hints that a property is an instance of a [Node]-derived type, optionally specified via the hint string (e.g. [code]"Node2D"[/code]). Editing it will show a dialog for picking a node from the scene.*/
	PropertyHintNodeType PropertyHint = 34
	/*Hints that a quaternion property should disable the temporary euler editor.*/
	PropertyHintHideQuaternionEdit PropertyHint = 35
	/*Hints that a string property is a password, and every character is replaced with the secret character.*/
	PropertyHintPassword PropertyHint = 36
	/*Represents the size of the [enum PropertyHint] enum.*/
	PropertyHintMax PropertyHint = 38
)

type PropertyUsageFlags int

const (
	/*The property is not stored, and does not display in the editor. This is the default for non-exported properties.*/
	PropertyUsageNone PropertyUsageFlags = 0
	/*The property is serialized and saved in the scene file (default for exported properties).*/
	PropertyUsageStorage PropertyUsageFlags = 2
	/*The property is shown in the [EditorInspector] (default for exported properties).*/
	PropertyUsageEditor PropertyUsageFlags = 4
	/*The property is excluded from the class reference.*/
	PropertyUsageInternal PropertyUsageFlags = 8
	/*The property can be checked in the [EditorInspector].*/
	PropertyUsageCheckable PropertyUsageFlags = 16
	/*The property is checked in the [EditorInspector].*/
	PropertyUsageChecked PropertyUsageFlags = 32
	/*Used to group properties together in the editor. See [EditorInspector].*/
	PropertyUsageGroup PropertyUsageFlags = 64
	/*Used to categorize properties together in the editor.*/
	PropertyUsageCategory PropertyUsageFlags = 128
	/*Used to group properties together in the editor in a subgroup (under a group). See [EditorInspector].*/
	PropertyUsageSubgroup PropertyUsageFlags = 256
	/*The property is a bitfield, i.e. it contains multiple flags represented as bits.*/
	PropertyUsageClassIsBitfield PropertyUsageFlags = 512
	/*The property does not save its state in [PackedScene].*/
	PropertyUsageNoInstanceState PropertyUsageFlags = 1024
	/*Editing the property prompts the user for restarting the editor.*/
	PropertyUsageRestartIfChanged PropertyUsageFlags = 2048
	/*The property is a script variable which should be serialized and saved in the scene file.*/
	PropertyUsageScriptVariable PropertyUsageFlags = 4096
	/*The property value of type [Object] will be stored even if its value is [code]null[/code].*/
	PropertyUsageStoreIfNull PropertyUsageFlags = 8192
	/*If this property is modified, all inspector fields will be refreshed.*/
	PropertyUsageUpdateAllIfModified PropertyUsageFlags = 16384
	PropertyUsageScriptDefaultValue  PropertyUsageFlags = 32768
	/*The property is an enum, i.e. it only takes named integer constants from its associated enumeration.*/
	PropertyUsageClassIsEnum PropertyUsageFlags = 65536
	/*If property has [code]nil[/code] as default value, its type will be [Variant].*/
	PropertyUsageNilIsVariant PropertyUsageFlags = 131072
	/*The property is an array.*/
	PropertyUsageArray PropertyUsageFlags = 262144
	/*When duplicating a resource with [method Resource.duplicate], and this flag is set on a property of that resource, the property should always be duplicated, regardless of the [code]subresources[/code] bool parameter.*/
	PropertyUsageAlwaysDuplicate PropertyUsageFlags = 524288
	/*When duplicating a resource with [method Resource.duplicate], and this flag is set on a property of that resource, the property should never be duplicated, regardless of the [code]subresources[/code] bool parameter.*/
	PropertyUsageNeverDuplicate PropertyUsageFlags = 1048576
	/*The property is only shown in the editor if modern renderers are supported (the Compatibility rendering method is excluded).*/
	PropertyUsageHighEndGfx PropertyUsageFlags = 2097152
	/*The [NodePath] property will always be relative to the scene's root. Mostly useful for local resources.*/
	PropertyUsageNodePathFromSceneRoot PropertyUsageFlags = 4194304
	/*Use when a resource is created on the fly, i.e. the getter will always return a different instance. [ResourceSaver] needs this information to properly save such resources.*/
	PropertyUsageResourceNotPersistent PropertyUsageFlags = 8388608
	/*Inserting an animation key frame of this property will automatically increment the value, allowing to easily keyframe multiple values in a row.*/
	PropertyUsageKeyingIncrements    PropertyUsageFlags = 16777216
	PropertyUsageDeferredSetResource PropertyUsageFlags = 33554432
	/*When this property is a [Resource] and base object is a [Node], a resource instance will be automatically created whenever the node is created in the editor.*/
	PropertyUsageEditorInstantiateObject PropertyUsageFlags = 67108864
	/*The property is considered a basic setting and will appear even when advanced mode is disabled. Used for project settings.*/
	PropertyUsageEditorBasicSetting PropertyUsageFlags = 134217728
	/*The property is read-only in the [EditorInspector].*/
	PropertyUsageReadOnly PropertyUsageFlags = 268435456
	/*An export preset property with this flag contains confidential information and is stored separately from the rest of the export preset configuration.*/
	PropertyUsageSecret PropertyUsageFlags = 536870912
	/*Default usage (storage and editor).*/
	PropertyUsageDefault PropertyUsageFlags = 6
	/*Default usage but without showing the property in the editor (storage).*/
	PropertyUsageNoEditor PropertyUsageFlags = 2
)

type VariantType int

const (
	/*Variable is [code]null[/code].*/
	TypeNil VariantType = 0
	/*Variable is of type [bool].*/
	TypeBool VariantType = 1
	/*Variable is of type [int].*/
	TypeInt VariantType = 2
	/*Variable is of type [float].*/
	TypeFloat VariantType = 3
	/*Variable is of type [String].*/
	TypeString VariantType = 4
	/*Variable is of type [Vector2].*/
	TypeVector2 VariantType = 5
	/*Variable is of type [Vector2i].*/
	TypeVector2i VariantType = 6
	/*Variable is of type [Rect2].*/
	TypeRect2 VariantType = 7
	/*Variable is of type [Rect2i].*/
	TypeRect2i VariantType = 8
	/*Variable is of type [Vector3].*/
	TypeVector3 VariantType = 9
	/*Variable is of type [Vector3i].*/
	TypeVector3i VariantType = 10
	/*Variable is of type [Transform2D].*/
	TypeTransform2d VariantType = 11
	/*Variable is of type [Vector4].*/
	TypeVector4 VariantType = 12
	/*Variable is of type [Vector4i].*/
	TypeVector4i VariantType = 13
	/*Variable is of type [Plane].*/
	TypePlane VariantType = 14
	/*Variable is of type [Quaternion].*/
	TypeQuaternion VariantType = 15
	/*Variable is of type [AABB].*/
	TypeAabb VariantType = 16
	/*Variable is of type [Basis].*/
	TypeBasis VariantType = 17
	/*Variable is of type [Transform3D].*/
	TypeTransform3d VariantType = 18
	/*Variable is of type [Projection].*/
	TypeProjection VariantType = 19
	/*Variable is of type [Color].*/
	TypeColor VariantType = 20
	/*Variable is of type [StringName].*/
	TypeStringName VariantType = 21
	/*Variable is of type [NodePath].*/
	TypeNodePath VariantType = 22
	/*Variable is of type [RID].*/
	TypeRid VariantType = 23
	/*Variable is of type [Object].*/
	TypeObject VariantType = 24
	/*Variable is of type [Callable].*/
	TypeCallable VariantType = 25
	/*Variable is of type [Signal].*/
	TypeSignal VariantType = 26
	/*Variable is of type [Dictionary].*/
	TypeDictionary VariantType = 27
	/*Variable is of type [Array].*/
	TypeArray VariantType = 28
	/*Variable is of type [PackedByteArray].*/
	TypePackedByteArray VariantType = 29
	/*Variable is of type [PackedInt32Array].*/
	TypePackedInt32Array VariantType = 30
	/*Variable is of type [PackedInt64Array].*/
	TypePackedInt64Array VariantType = 31
	/*Variable is of type [PackedFloat32Array].*/
	TypePackedFloat32Array VariantType = 32
	/*Variable is of type [PackedFloat64Array].*/
	TypePackedFloat64Array VariantType = 33
	/*Variable is of type [PackedStringArray].*/
	TypePackedStringArray VariantType = 34
	/*Variable is of type [PackedVector2Array].*/
	TypePackedVector2Array VariantType = 35
	/*Variable is of type [PackedVector3Array].*/
	TypePackedVector3Array VariantType = 36
	/*Variable is of type [PackedColorArray].*/
	TypePackedColorArray VariantType = 37
	/*Variable is of type [PackedVector4Array].*/
	TypePackedVector4Array VariantType = 38
	/*Represents the size of the [enum Variant.Type] enum.*/
	TypeMax VariantType = 39
)
