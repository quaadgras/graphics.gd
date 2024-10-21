package EditorScenePostImportPlugin

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This plugin type exists to modify the process of importing scenes, allowing to change the content as well as add importer options at every stage of the process.
	// EditorScenePostImportPlugin methods that can be overridden by a [Class] that extends it.
	type EditorScenePostImportPlugin interface {
		//Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
		GetInternalImportOptions(category gd.Int) 
		//Return true or false whether a given option should be visible. Return null to ignore.
		GetInternalOptionVisibility(category gd.Int, for_animation bool, option gd.String) gd.Variant
		//Return true whether updating the 3D view of the import dialog needs to be updated if an option has changed.
		GetInternalOptionUpdateViewRequired(category gd.Int, option gd.String) gd.Variant
		//Process a specific node or resource for a given category.
		InternalProcess(category gd.Int, base_node object.Node, node object.Node, resource object.Resource) 
		//Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
		GetImportOptions(path gd.String) 
		//Return true or false whether a given option should be visible. Return null to ignore.
		GetOptionVisibility(path gd.String, for_animation bool, option gd.String) gd.Variant
		//Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
		PreProcess(scene object.Node) 
		//Post process the scene. This function is called after the final scene has been configured.
		PostProcess(scene object.Node) 
	}

*/
type Simple [1]classdb.EditorScenePostImportPlugin
func (self Simple) GetOptionValue(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetOptionValue(gc, gc.StringName(name)))
}
func (self Simple) AddImportOption(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddImportOption(gc.String(name), value)
}
func (self Simple) AddImportOptionAdvanced(atype gd.VariantType, name string, default_value gd.Variant, hint gd.PropertyHint, hint_string string, usage_flags int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddImportOptionAdvanced(atype, gc.String(name), default_value, hint, gc.String(hint_string), gd.Int(usage_flags))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorScenePostImportPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Override to add internal import options. These will appear in the 3D scene import dialog. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (class) _get_internal_import_options(impl func(ptr unsafe.Pointer, category gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, category)
		ctx.End()
	}
}

/*
Return true or false whether a given option should be visible. Return null to ignore.
*/
func (class) _get_internal_option_visibility(impl func(ptr unsafe.Pointer, category gd.Int, for_animation bool, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, category, for_animation, option)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Return true whether updating the 3D view of the import dialog needs to be updated if an option has changed.
*/
func (class) _get_internal_option_update_view_required(impl func(ptr unsafe.Pointer, category gd.Int, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var option = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, category, option)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Process a specific node or resource for a given category.
*/
func (class) _internal_process(impl func(ptr unsafe.Pointer, category gd.Int, base_node object.Node, node object.Node, resource object.Resource) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var category = gd.UnsafeGet[gd.Int](p_args,0)
		var base_node object.Node
		base_node[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		var node object.Node
		node[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		var resource object.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,3)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, category, base_node, node, resource)
		ctx.End()
	}
}

/*
Override to add general import options. These will appear in the main import dock on the editor. Add options via [method add_import_option] and [method add_import_option_advanced].
*/
func (class) _get_import_options(impl func(ptr unsafe.Pointer, path gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path)
		ctx.End()
	}
}

/*
Return true or false whether a given option should be visible. Return null to ignore.
*/
func (class) _get_option_visibility(impl func(ptr unsafe.Pointer, path gd.String, for_animation bool, option gd.String) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var for_animation = gd.UnsafeGet[bool](p_args,1)
		var option = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, for_animation, option)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Pre Process the scene. This function is called right after the scene format loader loaded the scene and no changes have been made.
*/
func (class) _pre_process(impl func(ptr unsafe.Pointer, scene object.Node) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var scene object.Node
		scene[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, scene)
		ctx.End()
	}
}

/*
Post process the scene. This function is called after the final scene has been configured.
*/
func (class) _post_process(impl func(ptr unsafe.Pointer, scene object.Node) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var scene object.Node
		scene[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, scene)
		ctx.End()
	}
}

/*
Query the value of an option. This function can only be called from those querying visibility, or processing.
*/
//go:nosplit
func (self class) GetOptionValue(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScenePostImportPlugin.Bind_get_option_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Add a specific import option (name and default value only). This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
//go:nosplit
func (self class) AddImportOption(name gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScenePostImportPlugin.Bind_add_import_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Add a specific import option. This function can only be called from [method _get_import_options] and [method _get_internal_import_options].
*/
//go:nosplit
func (self class) AddImportOptionAdvanced(atype gd.VariantType, name gd.String, default_value gd.Variant, hint gd.PropertyHint, hint_string gd.String, usage_flags gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(default_value))
	callframe.Arg(frame, hint)
	callframe.Arg(frame, mmm.Get(hint_string))
	callframe.Arg(frame, usage_flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorScenePostImportPlugin.Bind_add_import_option_advanced, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorScenePostImportPlugin() Expert { return self[0].AsEditorScenePostImportPlugin() }


//go:nosplit
func (self Simple) AsEditorScenePostImportPlugin() Simple { return self[0].AsEditorScenePostImportPlugin() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorScenePostImportPlugin", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
