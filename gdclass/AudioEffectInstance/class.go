package AudioEffectInstance

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
An audio effect instance manipulates the audio it receives for a given effect. This instance is automatically created by an [AudioEffect] when it is added to a bus, and should usually not be created directly. If necessary, it can be fetched at run-time with [method AudioServer.get_bus_effect_instance].
	// AudioEffectInstance methods that can be overridden by a [Class] that extends it.
	type AudioEffectInstance interface {
		//Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
		//[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
		Process(src_buffer unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count int) 
		//Override this method to customize the processing behavior of this effect instance.
		//Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
		ProcessSilence() bool
	}

*/
type Go [1]classdb.AudioEffectInstance

/*
Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (Go) _process(impl func(ptr unsafe.Pointer, src_buffer unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_buffer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,1)
		var frame_count = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, src_buffer, dst_buffer, int(frame_count))
	}
}

/*
Override this method to customize the processing behavior of this effect instance.
Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
*/
func (Go) _process_silence(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectInstance
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectInstance"))
	return Go{classdb.AudioEffectInstance(object)}
}

/*
Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (class) _process(impl func(ptr unsafe.Pointer, src_buffer unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_buffer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,1)
		var frame_count = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, src_buffer, dst_buffer, frame_count)
	}
}

/*
Override this method to customize the processing behavior of this effect instance.
Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
*/
func (class) _process_silence(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsAudioEffectInstance() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectInstance() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process": return reflect.ValueOf(self._process);
	case "_process_silence": return reflect.ValueOf(self._process_silence);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_process": return reflect.ValueOf(self._process);
	case "_process_silence": return reflect.ValueOf(self._process_silence);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("AudioEffectInstance", func(ptr gd.Object) any { return classdb.AudioEffectInstance(ptr) })}
