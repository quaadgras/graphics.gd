package AudioEffectInstance

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
An audio effect instance manipulates the audio it receives for a given effect. This instance is automatically created by an [AudioEffect] when it is added to a bus, and should usually not be created directly. If necessary, it can be fetched at run-time with [method AudioServer.get_bus_effect_instance].
	// AudioEffectInstance methods that can be overridden by a [Class] that extends it.
	type AudioEffectInstance interface {
		//Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
		//[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
		Process(src_buffer unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count gd.Int) 
		//Override this method to customize the processing behavior of this effect instance.
		//Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
		ProcessSilence() bool
	}

*/
type Simple [1]classdb.AudioEffectInstance
func (Simple) _process(impl func(ptr unsafe.Pointer, src_buffer unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var src_buffer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,1)
		var frame_count = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, src_buffer, dst_buffer, int(frame_count))
		gc.End()
	}
}
func (Simple) _process_silence(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioEffectInstance
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (class) _process(impl func(ptr unsafe.Pointer, src_buffer unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var src_buffer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,1)
		var frame_count = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, src_buffer, dst_buffer, frame_count)
		ctx.End()
	}
}

/*
Override this method to customize the processing behavior of this effect instance.
Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
*/
func (class) _process_silence(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (self class) AsAudioEffectInstance() Expert { return self[0].AsAudioEffectInstance() }


//go:nosplit
func (self Simple) AsAudioEffectInstance() Simple { return self[0].AsAudioEffectInstance() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process": return reflect.ValueOf(self._process);
	case "_process_silence": return reflect.ValueOf(self._process_silence);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_process": return reflect.ValueOf(self._process);
	case "_process_silence": return reflect.ValueOf(self._process_silence);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioEffectInstance", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
