package AudioEffect

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
The base [Resource] for every audio effect. In the editor, an audio effect can be added to the current bus layout through the Audio panel. At run-time, it is also possible to manipulate audio effects through [method AudioServer.add_bus_effect], [method AudioServer.remove_bus_effect], and [method AudioServer.get_bus_effect].
When applied on a bus, an audio effect creates a corresponding [AudioEffectInstance]. The instance is directly responsible for manipulating the sound, based on the original audio effect's properties.
	// AudioEffect methods that can be overridden by a [Class] that extends it.
	type AudioEffect interface {
		//Override this method to customize the [AudioEffectInstance] created when this effect is applied on a bus in the editor's Audio panel, or through [method AudioServer.add_bus_effect].
		//[codeblock]
		//extends AudioEffect
		//
		//@export var strength = 4.0
		//
		//func _instantiate():
		//    var effect = CustomAudioEffectInstance.new()
		//    effect.base = self
		//
		//    return effect
		//[/codeblock]
		//[b]Note:[/b] It is recommended to keep a reference to the original [AudioEffect] in the new instance. Depending on the implementation this allows the effect instance to listen for changes at run-time and be modified accordingly.
		Instantiate() gdclass.AudioEffectInstance
	}

*/
type Go [1]classdb.AudioEffect

/*
Override this method to customize the [AudioEffectInstance] created when this effect is applied on a bus in the editor's Audio panel, or through [method AudioServer.add_bus_effect].
[codeblock]
extends AudioEffect

@export var strength = 4.0

func _instantiate():
    var effect = CustomAudioEffectInstance.new()
    effect.base = self

    return effect
[/codeblock]
[b]Note:[/b] It is recommended to keep a reference to the original [AudioEffect] in the new instance. Depending on the implementation this allows the effect instance to listen for changes at run-time and be modified accordingly.
*/
func (Go) _instantiate(impl func(ptr unsafe.Pointer) gdclass.AudioEffectInstance, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffect
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffect"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Override this method to customize the [AudioEffectInstance] created when this effect is applied on a bus in the editor's Audio panel, or through [method AudioServer.add_bus_effect].
[codeblock]
extends AudioEffect

@export var strength = 4.0

func _instantiate():
    var effect = CustomAudioEffectInstance.new()
    effect.base = self

    return effect
[/codeblock]
[b]Note:[/b] It is recommended to keep a reference to the original [AudioEffect] in the new instance. Depending on the implementation this allows the effect instance to listen for changes at run-time and be modified accordingly.
*/
func (class) _instantiate(impl func(ptr unsafe.Pointer) gdclass.AudioEffectInstance, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

func (self class) AsAudioEffect() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate": return reflect.ValueOf(self._instantiate);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate": return reflect.ValueOf(self._instantiate);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("AudioEffect", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
