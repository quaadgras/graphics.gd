// Package Material provides methods for working with Material object instances.
package Material

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[Material] is a base resource used for coloring and shading geometry. All materials inherit from it and almost all [VisualInstance3D] derived nodes carry a [Material]. A few flags and parameters are shared between all material types and are configured here.
Importantly, you can inherit from [Material] to create your own custom material type in script or in GDExtension.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Material)
*/
type Instance [1]gdclass.Material

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMaterial() Instance
}
type Interface interface {
	//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools. Used to access the RID of the [Material]'s [Shader].
	GetShaderRid() Resource.ID
	//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools.
	GetShaderMode() gdclass.ShaderMode
	//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member next_pass] should be shown in the editor or not.
	CanDoNextPass() bool
	//Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member render_priority] should be shown in the editor or not.
	CanUseRenderPriority() bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetShaderRid() (_ Resource.ID)         { return }
func (self implementation) GetShaderMode() (_ gdclass.ShaderMode) { return }
func (self implementation) CanDoNextPass() (_ bool)               { return }
func (self implementation) CanUseRenderPriority() (_ bool)        { return }

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools. Used to access the RID of the [Material]'s [Shader].
*/
func (Instance) _get_shader_rid(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools.
*/
func (Instance) _get_shader_mode(impl func(ptr unsafe.Pointer) gdclass.ShaderMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member next_pass] should be shown in the editor or not.
*/
func (Instance) _can_do_next_pass(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member render_priority] should be shown in the editor or not.
*/
func (Instance) _can_use_render_priority(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only available when running in the editor. Opens a popup that visualizes the generated shader code, including all variants and internal shader code.
*/
func (self Instance) InspectNativeShaderCode() {
	class(self).InspectNativeShaderCode()
}

/*
Creates a placeholder version of this resource ([PlaceholderMaterial]).
*/
func (self Instance) CreatePlaceholder() [1]gdclass.Resource {
	return [1]gdclass.Resource(class(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Material

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Material"))
	casted := Instance{*(*gdclass.Material)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) RenderPriority() int {
	return int(int(class(self).GetRenderPriority()))
}

func (self Instance) SetRenderPriority(value int) {
	class(self).SetRenderPriority(gd.Int(value))
}

func (self Instance) NextPass() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetNextPass())
}

func (self Instance) SetNextPass(value [1]gdclass.Material) {
	class(self).SetNextPass(value)
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools. Used to access the RID of the [Material]'s [Shader].
*/
func (class) _get_shader_rid(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally by various editor tools.
*/
func (class) _get_shader_mode(impl func(ptr unsafe.Pointer) gdclass.ShaderMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member next_pass] should be shown in the editor or not.
*/
func (class) _can_do_next_pass(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Only exposed for the purpose of overriding. You cannot call this function directly. Used internally to determine if [member render_priority] should be shown in the editor or not.
*/
func (class) _can_use_render_priority(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetNextPass(next_pass [1]gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(next_pass[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Material.Bind_set_next_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNextPass() [1]gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Material.Bind_get_next_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Material.Bind_set_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRenderPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Material.Bind_get_render_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Only available when running in the editor. Opens a popup that visualizes the generated shader code, including all variants and internal shader code.
*/
//go:nosplit
func (self class) InspectNativeShaderCode() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Material.Bind_inspect_native_shader_code, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a placeholder version of this resource ([PlaceholderMaterial]).
*/
//go:nosplit
func (self class) CreatePlaceholder() [1]gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Material.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_shader_rid":
		return reflect.ValueOf(self._get_shader_rid)
	case "_get_shader_mode":
		return reflect.ValueOf(self._get_shader_mode)
	case "_can_do_next_pass":
		return reflect.ValueOf(self._can_do_next_pass)
	case "_can_use_render_priority":
		return reflect.ValueOf(self._can_use_render_priority)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_shader_rid":
		return reflect.ValueOf(self._get_shader_rid)
	case "_get_shader_mode":
		return reflect.ValueOf(self._get_shader_mode)
	case "_can_do_next_pass":
		return reflect.ValueOf(self._can_do_next_pass)
	case "_can_use_render_priority":
		return reflect.ValueOf(self._can_use_render_priority)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Material", func(ptr gd.Object) any { return [1]gdclass.Material{*(*gdclass.Material)(unsafe.Pointer(&ptr))} })
}
