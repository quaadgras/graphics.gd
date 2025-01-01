package RenderSceneBuffersConfiguration

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This configuration object is created and populated by the render engine on a viewport change and used to (re)configure a [RenderSceneBuffers] object.
*/
type Instance [1]classdb.RenderSceneBuffersConfiguration
type Any interface {
	gd.IsClass
	AsRenderSceneBuffersConfiguration() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderSceneBuffersConfiguration

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneBuffersConfiguration"))
	return Instance{classdb.RenderSceneBuffersConfiguration(object)}
}

func (self Instance) RenderTarget() Resource.ID {
	return Resource.ID(class(self).GetRenderTarget())
}

func (self Instance) SetRenderTarget(value Resource.ID) {
	class(self).SetRenderTarget(value)
}

func (self Instance) InternalSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetInternalSize())
}

func (self Instance) SetInternalSize(value Vector2i.XY) {
	class(self).SetInternalSize(gd.Vector2i(value))
}

func (self Instance) TargetSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetTargetSize())
}

func (self Instance) SetTargetSize(value Vector2i.XY) {
	class(self).SetTargetSize(gd.Vector2i(value))
}

func (self Instance) ViewCount() int {
	return int(int(class(self).GetViewCount()))
}

func (self Instance) SetViewCount(value int) {
	class(self).SetViewCount(gd.Int(value))
}

func (self Instance) Scaling3dMode() classdb.RenderingServerViewportScaling3DMode {
	return classdb.RenderingServerViewportScaling3DMode(class(self).GetScaling3dMode())
}

func (self Instance) SetScaling3dMode(value classdb.RenderingServerViewportScaling3DMode) {
	class(self).SetScaling3dMode(value)
}

func (self Instance) Msaa3d() classdb.RenderingServerViewportMSAA {
	return classdb.RenderingServerViewportMSAA(class(self).GetMsaa3d())
}

func (self Instance) SetMsaa3d(value classdb.RenderingServerViewportMSAA) {
	class(self).SetMsaa3d(value)
}

func (self Instance) ScreenSpaceAa() classdb.RenderingServerViewportScreenSpaceAA {
	return classdb.RenderingServerViewportScreenSpaceAA(class(self).GetScreenSpaceAa())
}

func (self Instance) SetScreenSpaceAa(value classdb.RenderingServerViewportScreenSpaceAA) {
	class(self).SetScreenSpaceAa(value)
}

func (self Instance) FsrSharpness() Float.X {
	return Float.X(Float.X(class(self).GetFsrSharpness()))
}

func (self Instance) SetFsrSharpness(value Float.X) {
	class(self).SetFsrSharpness(gd.Float(value))
}

func (self Instance) TextureMipmapBias() Float.X {
	return Float.X(Float.X(class(self).GetTextureMipmapBias()))
}

func (self Instance) SetTextureMipmapBias(value Float.X) {
	class(self).SetTextureMipmapBias(gd.Float(value))
}

//go:nosplit
func (self class) GetRenderTarget() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderTarget(render_target gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_render_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInternalSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_internal_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInternalSize(internal_size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, internal_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_internal_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetSize(target_size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, target_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_target_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetViewCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetViewCount(view_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, view_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_view_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaling3dMode() classdb.RenderingServerViewportScaling3DMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportScaling3DMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaling3dMode(scaling_3d_mode classdb.RenderingServerViewportScaling3DMode) {
	var frame = callframe.New()
	callframe.Arg(frame, scaling_3d_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsaa3d() classdb.RenderingServerViewportMSAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportMSAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsaa3d(msaa_3d classdb.RenderingServerViewportMSAA) {
	var frame = callframe.New()
	callframe.Arg(frame, msaa_3d)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenSpaceAa() classdb.RenderingServerViewportScreenSpaceAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingServerViewportScreenSpaceAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenSpaceAa(screen_space_aa classdb.RenderingServerViewportScreenSpaceAA) {
	var frame = callframe.New()
	callframe.Arg(frame, screen_space_aa)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFsrSharpness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFsrSharpness(fsr_sharpness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fsr_sharpness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMipmapBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_get_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMipmapBias(texture_mipmap_bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_mipmap_bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderSceneBuffersConfiguration.Bind_set_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsRenderSceneBuffersConfiguration() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneBuffersConfiguration() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("RenderSceneBuffersConfiguration", func(ptr gd.Object) any {
		return [1]classdb.RenderSceneBuffersConfiguration{classdb.RenderSceneBuffersConfiguration(ptr)}
	})
}
