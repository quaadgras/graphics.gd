package EditorScenePostImportPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This plugin type exists to modify the process of importing scenes, allowing to change the content as well as add importer options at every stage of the process.
	// EditorScenePostImportPlugin methods that can be overridden by a [Class] that extends it.
	type EditorScenePostImportPlugin interface {
		//Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
		GetInternalImportOptions(category int) 
		//Return true or false whether a given option should be visible. Return null to ignore.
		GetInternalOptionVisibility(category int, for_animation bool, option string) gd.Variant
		//Return true whether updating the 3D view of the import dialog needs to be updated if an option has changed.
		GetInternalOptionUpdateViewRequired(category int, option string) gd.Variant
		//Process a specific node or resource for a given category.
		InternalProcess(category int, base_node gdclass.Node, node gdclass.Node, resource gdclass.Resource) 
		//Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
		GetImportOptions(path string) 
		//Return true or false whether a given option should be visible. Return null to ignore.
		GetOptionVisibility(path string, for_animation bool, option string) gd.Variant
		//Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
		PreProcess(scene gdclass.Node) 
		//Post process the scene. This function is called after the final scene has been configured.
		PostProcess(scene gdclass.Node) 
	}

*/
type Go [1]classdb.EditorScenePostImportPlugin

/*
Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (Go) _get_internal_import_options(impl func(ptr unsafe.Pointer, category int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(category))
	}
}

/*
Return true or false whether a given option should be visible. Return null to ignore.
*/
func (Go) _get_internal_option_visibility(impl func(ptr unsafe.Pointer, category int, for_animation bool, option string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(category), for_animation, option.String())
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return true whether updating the 3D view of the import dialog needs to be updated if an option has changed.
*/
func (Go) _get_internal_option_update_view_required(impl func(ptr unsafe.Pointer, category int, option string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(category), option.String())
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Process a specific node or resource for a given category.
*/
func (Go) _internal_process(impl func(ptr unsafe.Pointer, category int, base_node gdclass.Node, node gdclass.Node, resource gdclass.Resource) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var base_node = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,1)})}
		defer discreet.End(base_node[0])
		var node = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,2)})}
		defer discreet.End(node[0])
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,3)})}
		defer discreet.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(category), base_node, node, resource)
	}
}

/*
Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (Go) _get_import_options(impl func(ptr unsafe.Pointer, path string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path.String())
	}
}

/*
Return true or false whether a given option should be visible. Return null to ignore.
*/
func (Go) _get_option_visibility(impl func(ptr unsafe.Pointer, path string, for_animation bool, option string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(path)
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		defer discreet.End(option)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), for_animation, option.String())
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
*/
func (Go) _pre_process(impl func(ptr unsafe.Pointer, scene gdclass.Node) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var scene = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, scene)
	}
}

/*
Post process the scene. This function is called after the final scene has been configured.
*/
func (Go) _post_process(impl func(ptr unsafe.Pointer, scene gdclass.Node) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var scene = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, scene)
	}
}

/*
Query the value of an option. This function can only be called from those querying visibility, or processing.
*/
func (self Go) GetOptionValue(name string) gd.Variant {
	return gd.Variant(class(self).GetOptionValue(gd.NewStringName(name)))
}

/*
Add a specific import option (name and default value only). This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
func (self Go) AddImportOption(name string, value gd.Variant) {
	class(self).AddImportOption(gd.NewString(name), value)
}

/*
Add a specific import option. This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
func (self Go) AddImportOptionAdvanced(atype gd.VariantType, name string, default_value gd.Variant) {
	class(self).AddImportOptionAdvanced(atype, gd.NewString(name), default_value, 0, gd.NewString(""), gd.Int(6))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorScenePostImportPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorScenePostImportPlugin"))
	return Go{classdb.EditorScenePostImportPlugin(object)}
}

/*
Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (class) _get_internal_import_options(impl func(ptr unsafe.Pointer, category gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, category)
	}
}

/*
Return true or false whether a given option should be visible. Return null to ignore.
*/
func (class) _get_internal_option_visibility(impl func(ptr unsafe.Pointer, category gd.Int, for_animation bool, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, category, for_animation, option)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Return true whether updating the 3D view of the import dialog needs to be updated if an option has changed.
*/
func (class) _get_internal_option_update_view_required(impl func(ptr unsafe.Pointer, category gd.Int, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, category, option)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Process a specific node or resource for a given category.
*/
func (class) _internal_process(impl func(ptr unsafe.Pointer, category gd.Int, base_node gdclass.Node, node gdclass.Node, resource gdclass.Resource) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var base_node = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,1)})}
		defer discreet.End(base_node[0])
		var node = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,2)})}
		defer discreet.End(node[0])
		var resource = gdclass.Resource{discreet.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args,3)})}
		defer discreet.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, category, base_node, node, resource)
	}
}

/*
Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (class) _get_import_options(impl func(ptr unsafe.Pointer, path gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path)
	}
}

/*
Return true or false whether a given option should be visible. Return null to ignore.
*/
func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path gd.String, for_animation bool, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,0))
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = discreet.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, for_animation, option)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
*/
func (class) _pre_process(impl func(ptr unsafe.Pointer, scene gdclass.Node) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var scene = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, scene)
	}
}

/*
Post process the scene. This function is called after the final scene has been configured.
*/
func (class) _post_process(impl func(ptr unsafe.Pointer, scene gdclass.Node) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var scene = gdclass.Node{discreet.New[classdb.Node]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(scene[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, scene)
	}
}

/*
Query the value of an option. This function can only be called from those querying visibility, or processing.
*/
//go:nosplit
func (self class) GetOptionValue(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImportPlugin.Bind_get_option_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Add a specific import option (name and default value only). This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
//go:nosplit
func (self class) AddImportOption(name gd.String, value gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImportPlugin.Bind_add_import_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Add a specific import option. This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
//go:nosplit
func (self class) AddImportOptionAdvanced(atype gd.VariantType, name gd.String, default_value gd.Variant, hint gd.PropertyHint, hint_string gd.String, usage_flags gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(default_value))
	callframe.Arg(frame, hint)
	callframe.Arg(frame, discreet.Get(hint_string))
	callframe.Arg(frame, usage_flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorScenePostImportPlugin.Bind_add_import_option_advanced, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsEditorScenePostImportPlugin() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorScenePostImportPlugin() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_internal_import_options": return reflect.ValueOf(self._get_internal_import_options);
	case "_get_internal_option_visibility": return reflect.ValueOf(self._get_internal_option_visibility);
	case "_get_internal_option_update_view_required": return reflect.ValueOf(self._get_internal_option_update_view_required);
	case "_internal_process": return reflect.ValueOf(self._internal_process);
	case "_get_import_options": return reflect.ValueOf(self._get_import_options);
	case "_get_option_visibility": return reflect.ValueOf(self._get_option_visibility);
	case "_pre_process": return reflect.ValueOf(self._pre_process);
	case "_post_process": return reflect.ValueOf(self._post_process);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_internal_import_options": return reflect.ValueOf(self._get_internal_import_options);
	case "_get_internal_option_visibility": return reflect.ValueOf(self._get_internal_option_visibility);
	case "_get_internal_option_update_view_required": return reflect.ValueOf(self._get_internal_option_update_view_required);
	case "_internal_process": return reflect.ValueOf(self._internal_process);
	case "_get_import_options": return reflect.ValueOf(self._get_import_options);
	case "_get_option_visibility": return reflect.ValueOf(self._get_option_visibility);
	case "_pre_process": return reflect.ValueOf(self._pre_process);
	case "_post_process": return reflect.ValueOf(self._post_process);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorScenePostImportPlugin", func(ptr gd.Object) any { return classdb.EditorScenePostImportPlugin(ptr) })}
type InternalImportCategory = classdb.EditorScenePostImportPluginInternalImportCategory

const (
	InternalImportCategoryNode InternalImportCategory = 0
	InternalImportCategoryMesh3dNode InternalImportCategory = 1
	InternalImportCategoryMesh InternalImportCategory = 2
	InternalImportCategoryMaterial InternalImportCategory = 3
	InternalImportCategoryAnimation InternalImportCategory = 4
	InternalImportCategoryAnimationNode InternalImportCategory = 5
	InternalImportCategorySkeleton3dNode InternalImportCategory = 6
	InternalImportCategoryMax InternalImportCategory = 7
)
