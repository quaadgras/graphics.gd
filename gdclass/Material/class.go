package Material

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Material] is a base resource used for coloring and shading geometry. All materials inherit from it and almost all [VisualInstance3D] derived nodes carry a [Material]. A few flags and parameters are shared between all material types and are configured here.
Importantly, you can inherit from [Material] to create your own custom material type in script or in GDExtension.
	// Material methods that can be overridden by a [Class] that extends it.
	type Material interface {
		//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools. Used to access the RID of the [Material]'s [Shader].
		GetShaderRid() gd.RID
		//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools.
		GetShaderMode() classdb.ShaderMode
		//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member next_pass] should be shown in the editor or not.
		CanDoNextPass() bool
		//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member render_priority] should be shown in the editor or not.
		CanUseRenderPriority() bool
	}

*/
type Go [1]classdb.Material

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools. Used to access the RID of the [Material]'s [Shader].
*/
func (Go) _get_shader_rid(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools.
*/
func (Go) _get_shader_mode(impl func(ptr unsafe.Pointer) classdb.ShaderMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member next_pass] should be shown in the editor or not.
*/
func (Go) _can_do_next_pass(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member render_priority] should be shown in the editor or not.
*/
func (Go) _can_use_render_priority(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Only available when running in the editor. Opens a popup that visualizes the generated shader code, including all variants and internal shader code.
*/
func (self Go) InspectNativeShaderCode() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InspectNativeShaderCode()
}

/*
Creates a placeholder version of this resource ([PlaceholderMaterial]).
*/
func (self Go) CreatePlaceholder() gdclass.Resource {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Resource(class(self).CreatePlaceholder(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Material
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Material"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RenderPriority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetRenderPriority()))
}

func (self Go) SetRenderPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRenderPriority(gd.Int(value))
}

func (self Go) NextPass() gdclass.Material {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Material(class(self).GetNextPass(gc))
}

func (self Go) SetNextPass(value gdclass.Material) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNextPass(value)
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools. Used to access the RID of the [Material]'s [Shader].
*/
func (class) _get_shader_rid(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools.
*/
func (class) _get_shader_mode(impl func(ptr unsafe.Pointer) classdb.ShaderMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member next_pass] should be shown in the editor or not.
*/
func (class) _can_do_next_pass(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member render_priority] should be shown in the editor or not.
*/
func (class) _can_use_render_priority(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetNextPass(next_pass gdclass.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(next_pass[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Material.Bind_set_next_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNextPass(ctx gd.Lifetime) gdclass.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Material.Bind_get_next_pass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Material.Bind_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Material.Bind_get_render_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Only available when running in the editor. Opens a popup that visualizes the generated shader code, including all variants and internal shader code.
*/
//go:nosplit
func (self class) InspectNativeShaderCode()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Material.Bind_inspect_native_shader_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a placeholder version of this resource ([PlaceholderMaterial]).
*/
//go:nosplit
func (self class) CreatePlaceholder(ctx gd.Lifetime) gdclass.Resource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Material.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Resource
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_shader_rid": return reflect.ValueOf(self._get_shader_rid);
	case "_get_shader_mode": return reflect.ValueOf(self._get_shader_mode);
	case "_can_do_next_pass": return reflect.ValueOf(self._can_do_next_pass);
	case "_can_use_render_priority": return reflect.ValueOf(self._can_use_render_priority);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_shader_rid": return reflect.ValueOf(self._get_shader_rid);
	case "_get_shader_mode": return reflect.ValueOf(self._get_shader_mode);
	case "_can_do_next_pass": return reflect.ValueOf(self._can_do_next_pass);
	case "_can_use_render_priority": return reflect.ValueOf(self._can_use_render_priority);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Material", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
