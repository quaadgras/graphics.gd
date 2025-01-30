// Package OpenXRExtensionWrapperExtension provides methods for working with OpenXRExtensionWrapperExtension object instances.
package OpenXRExtensionWrapperExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[OpenXRExtensionWrapperExtension] allows clients to implement OpenXR extensions with GDExtension. The extension should be registered with [method register_extension_wrapper].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=OpenXRExtensionWrapperExtension)
*/
type Instance [1]gdclass.OpenXRExtensionWrapperExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRExtensionWrapperExtension() Instance
}
type Interface interface {
	//Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
	//- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
	//- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
	GetRequestedExtensions() map[any]any
	//Adds additional data structures when interogating OpenXR system abilities.
	SetSystemPropertiesAndGetNextPointer(next_pointer unsafe.Pointer) int
	//Adds additional data structures when the OpenXR instance is created.
	SetInstanceCreateInfoAndGetNextPointer(next_pointer unsafe.Pointer) int
	//Adds additional data structures when the OpenXR session is created.
	SetSessionCreateAndGetNextPointer(next_pointer unsafe.Pointer) int
	//Adds additional data structures when creating OpenXR swapchains.
	SetSwapchainCreateInfoAndGetNextPointer(next_pointer unsafe.Pointer) int
	//Adds additional data structures when each hand tracker is created.
	SetHandJointLocationsAndGetNextPointer(hand_index int, next_pointer unsafe.Pointer) int
	//Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
	//This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
	GetCompositionLayerCount() int
	//Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
	//This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
	GetCompositionLayer(index int) int
	//Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
	//This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
	GetCompositionLayerOrder(index int) int
	//Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
	GetSuggestedTrackerNames() []string
	//Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
	//Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
	OnRegisterMetadata()
	//Called before the OpenXR instance is created.
	OnBeforeInstanceCreated()
	//Called right after the OpenXR instance is created.
	OnInstanceCreated(instance int)
	//Called right before the OpenXR instance is destroyed.
	OnInstanceDestroyed()
	//Called right after the OpenXR session is created.
	OnSessionCreated(session int)
	//Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
	OnProcess()
	//Called right before the XR viewports begin their rendering step.
	OnPreRender()
	//Called right after the main swapchains are (re)created.
	OnMainSwapchainsCreated()
	//Called right before the OpenXR session is destroyed.
	OnSessionDestroyed()
	//Called when the OpenXR session state is changed to idle.
	OnStateIdle()
	//Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
	OnStateReady()
	//Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
	OnStateSynchronized()
	//Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
	OnStateVisible()
	//Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
	OnStateFocused()
	//Called when the OpenXR session state is changed to stopping.
	OnStateStopping()
	//Called when the OpenXR session state is changed to loss pending.
	OnStateLossPending()
	//Called when the OpenXR session state is changed to exiting.
	OnStateExiting()
	//Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
	OnEventPolled(event unsafe.Pointer) bool
	//Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
	//[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
	//[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
	SetViewportCompositionLayerAndGetNextPointer(layer unsafe.Pointer, property_values map[any]any, next_pointer unsafe.Pointer) int
	//Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
	GetViewportCompositionLayerExtensionProperties() []map[any]any
	//Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
	GetViewportCompositionLayerExtensionPropertyDefaults() map[any]any
	//Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
	//[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
	OnViewportCompositionLayerDestroyed(layer unsafe.Pointer)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetRequestedExtensions() (_ map[any]any) { return }
func (self implementation) SetSystemPropertiesAndGetNextPointer(next_pointer unsafe.Pointer) (_ int) {
	return
}
func (self implementation) SetInstanceCreateInfoAndGetNextPointer(next_pointer unsafe.Pointer) (_ int) {
	return
}
func (self implementation) SetSessionCreateAndGetNextPointer(next_pointer unsafe.Pointer) (_ int) {
	return
}
func (self implementation) SetSwapchainCreateInfoAndGetNextPointer(next_pointer unsafe.Pointer) (_ int) {
	return
}
func (self implementation) SetHandJointLocationsAndGetNextPointer(hand_index int, next_pointer unsafe.Pointer) (_ int) {
	return
}
func (self implementation) GetCompositionLayerCount() (_ int)           { return }
func (self implementation) GetCompositionLayer(index int) (_ int)       { return }
func (self implementation) GetCompositionLayerOrder(index int) (_ int)  { return }
func (self implementation) GetSuggestedTrackerNames() (_ []string)      { return }
func (self implementation) OnRegisterMetadata()                         { return }
func (self implementation) OnBeforeInstanceCreated()                    { return }
func (self implementation) OnInstanceCreated(instance int)              { return }
func (self implementation) OnInstanceDestroyed()                        { return }
func (self implementation) OnSessionCreated(session int)                { return }
func (self implementation) OnProcess()                                  { return }
func (self implementation) OnPreRender()                                { return }
func (self implementation) OnMainSwapchainsCreated()                    { return }
func (self implementation) OnSessionDestroyed()                         { return }
func (self implementation) OnStateIdle()                                { return }
func (self implementation) OnStateReady()                               { return }
func (self implementation) OnStateSynchronized()                        { return }
func (self implementation) OnStateVisible()                             { return }
func (self implementation) OnStateFocused()                             { return }
func (self implementation) OnStateStopping()                            { return }
func (self implementation) OnStateLossPending()                         { return }
func (self implementation) OnStateExiting()                             { return }
func (self implementation) OnEventPolled(event unsafe.Pointer) (_ bool) { return }
func (self implementation) SetViewportCompositionLayerAndGetNextPointer(layer unsafe.Pointer, property_values map[any]any, next_pointer unsafe.Pointer) (_ int) {
	return
}
func (self implementation) GetViewportCompositionLayerExtensionProperties() (_ []map[any]any) { return }
func (self implementation) GetViewportCompositionLayerExtensionPropertyDefaults() (_ map[any]any) {
	return
}
func (self implementation) OnViewportCompositionLayerDestroyed(layer unsafe.Pointer) { return }

/*
Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
*/
func (Instance) _get_requested_extensions(impl func(ptr unsafe.Pointer) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Adds additional data structures when interogating OpenXR system abilities.
*/
func (Instance) _set_system_properties_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Adds additional data structures when the OpenXR instance is created.
*/
func (Instance) _set_instance_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Adds additional data structures when the OpenXR session is created.
*/
func (Instance) _set_session_create_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Adds additional data structures when creating OpenXR swapchains.
*/
func (Instance) _set_swapchain_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Adds additional data structures when each hand tracker is created.
*/
func (Instance) _set_hand_joint_locations_and_get_next_pointer(impl func(ptr unsafe.Pointer, hand_index int, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var hand_index = gd.UnsafeGet[int64](p_args, 0)

		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(hand_index), next_pointer)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Instance) _get_composition_layer_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Instance) _get_composition_layer(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Instance) _get_composition_layer_order(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
*/
func (Instance) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
*/
func (Instance) _on_register_metadata(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called before the OpenXR instance is created.
*/
func (Instance) _on_before_instance_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR instance is created.
*/
func (Instance) _on_instance_created(impl func(ptr unsafe.Pointer, instance int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var instance = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(instance))
	}
}

/*
Called right before the OpenXR instance is destroyed.
*/
func (Instance) _on_instance_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR session is created.
*/
func (Instance) _on_session_created(impl func(ptr unsafe.Pointer, session int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var session = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(session))
	}
}

/*
Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
*/
func (Instance) _on_process(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the XR viewports begin their rendering step.
*/
func (Instance) _on_pre_render(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the main swapchains are (re)created.
*/
func (Instance) _on_main_swapchains_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the OpenXR session is destroyed.
*/
func (Instance) _on_session_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to idle.
*/
func (Instance) _on_state_idle(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
*/
func (Instance) _on_state_ready(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
*/
func (Instance) _on_state_synchronized(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
*/
func (Instance) _on_state_visible(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
*/
func (Instance) _on_state_focused(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to stopping.
*/
func (Instance) _on_state_stopping(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to loss pending.
*/
func (Instance) _on_state_loss_pending(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to exiting.
*/
func (Instance) _on_state_exiting(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
*/
func (Instance) _on_event_polled(impl func(ptr unsafe.Pointer, event unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (Instance) _set_viewport_composition_layer_and_get_next_pointer(impl func(ptr unsafe.Pointer, layer unsafe.Pointer, property_values map[any]any, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var property_values = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(property_values))
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, gd.DictionaryAs[map[any]any](property_values), next_pointer)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
*/
func (Instance) _get_viewport_composition_layer_extension_properties(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
*/
func (Instance) _get_viewport_composition_layer_extension_property_defaults(impl func(ptr unsafe.Pointer) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (Instance) _on_viewport_composition_layer_destroyed(impl func(ptr unsafe.Pointer, layer unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, layer)
	}
}

/*
Returns the created [OpenXRAPIExtension], which can be used to access the OpenXR API.
*/
func (self Instance) GetOpenxrApi() [1]gdclass.OpenXRAPIExtension { //gd:OpenXRExtensionWrapperExtension.get_openxr_api
	return [1]gdclass.OpenXRAPIExtension(class(self).GetOpenxrApi())
}

/*
Registers the extension. This should happen at core module initialization level.
*/
func (self Instance) RegisterExtensionWrapper() { //gd:OpenXRExtensionWrapperExtension.register_extension_wrapper
	class(self).RegisterExtensionWrapper()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRExtensionWrapperExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRExtensionWrapperExtension"))
	casted := Instance{*(*gdclass.OpenXRExtensionWrapperExtension)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
*/
func (class) _get_requested_extensions(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Adds additional data structures when interogating OpenXR system abilities.
*/
func (class) _set_system_properties_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when the OpenXR instance is created.
*/
func (class) _set_instance_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when the OpenXR session is created.
*/
func (class) _set_session_create_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when creating OpenXR swapchains.
*/
func (class) _set_swapchain_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when each hand tracker is created.
*/
func (class) _set_hand_joint_locations_and_get_next_pointer(impl func(ptr unsafe.Pointer, hand_index int64, next_pointer unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var hand_index = gd.UnsafeGet[int64](p_args, 0)

		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, hand_index, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer_count(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer(impl func(ptr unsafe.Pointer, index int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer_order(impl func(ptr unsafe.Pointer, index int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var index = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
*/
func (class) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
*/
func (class) _on_register_metadata(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called before the OpenXR instance is created.
*/
func (class) _on_before_instance_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR instance is created.
*/
func (class) _on_instance_created(impl func(ptr unsafe.Pointer, instance int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var instance = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, instance)
	}
}

/*
Called right before the OpenXR instance is destroyed.
*/
func (class) _on_instance_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR session is created.
*/
func (class) _on_session_created(impl func(ptr unsafe.Pointer, session int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var session = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, session)
	}
}

/*
Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
*/
func (class) _on_process(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the XR viewports begin their rendering step.
*/
func (class) _on_pre_render(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the main swapchains are (re)created.
*/
func (class) _on_main_swapchains_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the OpenXR session is destroyed.
*/
func (class) _on_session_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to idle.
*/
func (class) _on_state_idle(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
*/
func (class) _on_state_ready(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
*/
func (class) _on_state_synchronized(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
*/
func (class) _on_state_visible(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
*/
func (class) _on_state_focused(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to stopping.
*/
func (class) _on_state_stopping(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to loss pending.
*/
func (class) _on_state_loss_pending(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to exiting.
*/
func (class) _on_state_exiting(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
*/
func (class) _on_event_polled(impl func(ptr unsafe.Pointer, event unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var event = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (class) _set_viewport_composition_layer_and_get_next_pointer(impl func(ptr unsafe.Pointer, layer unsafe.Pointer, property_values Dictionary.Any, next_pointer unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var property_values = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(property_values))
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, property_values, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
*/
func (class) _get_viewport_composition_layer_extension_properties(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
*/
func (class) _get_viewport_composition_layer_extension_property_defaults(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (class) _on_viewport_composition_layer_destroyed(impl func(ptr unsafe.Pointer, layer unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, layer)
	}
}

/*
Returns the created [OpenXRAPIExtension], which can be used to access the OpenXR API.
*/
//go:nosplit
func (self class) GetOpenxrApi() [1]gdclass.OpenXRAPIExtension { //gd:OpenXRExtensionWrapperExtension.get_openxr_api
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRExtensionWrapperExtension.Bind_get_openxr_api, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRAPIExtension{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRAPIExtension](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Registers the extension. This should happen at core module initialization level.
*/
//go:nosplit
func (self class) RegisterExtensionWrapper() { //gd:OpenXRExtensionWrapperExtension.register_extension_wrapper
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRExtensionWrapperExtension.Bind_register_extension_wrapper, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsOpenXRExtensionWrapperExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRExtensionWrapperExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_requested_extensions":
		return reflect.ValueOf(self._get_requested_extensions)
	case "_set_system_properties_and_get_next_pointer":
		return reflect.ValueOf(self._set_system_properties_and_get_next_pointer)
	case "_set_instance_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_instance_create_info_and_get_next_pointer)
	case "_set_session_create_and_get_next_pointer":
		return reflect.ValueOf(self._set_session_create_and_get_next_pointer)
	case "_set_swapchain_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_swapchain_create_info_and_get_next_pointer)
	case "_set_hand_joint_locations_and_get_next_pointer":
		return reflect.ValueOf(self._set_hand_joint_locations_and_get_next_pointer)
	case "_get_composition_layer_count":
		return reflect.ValueOf(self._get_composition_layer_count)
	case "_get_composition_layer":
		return reflect.ValueOf(self._get_composition_layer)
	case "_get_composition_layer_order":
		return reflect.ValueOf(self._get_composition_layer_order)
	case "_get_suggested_tracker_names":
		return reflect.ValueOf(self._get_suggested_tracker_names)
	case "_on_register_metadata":
		return reflect.ValueOf(self._on_register_metadata)
	case "_on_before_instance_created":
		return reflect.ValueOf(self._on_before_instance_created)
	case "_on_instance_created":
		return reflect.ValueOf(self._on_instance_created)
	case "_on_instance_destroyed":
		return reflect.ValueOf(self._on_instance_destroyed)
	case "_on_session_created":
		return reflect.ValueOf(self._on_session_created)
	case "_on_process":
		return reflect.ValueOf(self._on_process)
	case "_on_pre_render":
		return reflect.ValueOf(self._on_pre_render)
	case "_on_main_swapchains_created":
		return reflect.ValueOf(self._on_main_swapchains_created)
	case "_on_session_destroyed":
		return reflect.ValueOf(self._on_session_destroyed)
	case "_on_state_idle":
		return reflect.ValueOf(self._on_state_idle)
	case "_on_state_ready":
		return reflect.ValueOf(self._on_state_ready)
	case "_on_state_synchronized":
		return reflect.ValueOf(self._on_state_synchronized)
	case "_on_state_visible":
		return reflect.ValueOf(self._on_state_visible)
	case "_on_state_focused":
		return reflect.ValueOf(self._on_state_focused)
	case "_on_state_stopping":
		return reflect.ValueOf(self._on_state_stopping)
	case "_on_state_loss_pending":
		return reflect.ValueOf(self._on_state_loss_pending)
	case "_on_state_exiting":
		return reflect.ValueOf(self._on_state_exiting)
	case "_on_event_polled":
		return reflect.ValueOf(self._on_event_polled)
	case "_set_viewport_composition_layer_and_get_next_pointer":
		return reflect.ValueOf(self._set_viewport_composition_layer_and_get_next_pointer)
	case "_get_viewport_composition_layer_extension_properties":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_properties)
	case "_get_viewport_composition_layer_extension_property_defaults":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_property_defaults)
	case "_on_viewport_composition_layer_destroyed":
		return reflect.ValueOf(self._on_viewport_composition_layer_destroyed)
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_requested_extensions":
		return reflect.ValueOf(self._get_requested_extensions)
	case "_set_system_properties_and_get_next_pointer":
		return reflect.ValueOf(self._set_system_properties_and_get_next_pointer)
	case "_set_instance_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_instance_create_info_and_get_next_pointer)
	case "_set_session_create_and_get_next_pointer":
		return reflect.ValueOf(self._set_session_create_and_get_next_pointer)
	case "_set_swapchain_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_swapchain_create_info_and_get_next_pointer)
	case "_set_hand_joint_locations_and_get_next_pointer":
		return reflect.ValueOf(self._set_hand_joint_locations_and_get_next_pointer)
	case "_get_composition_layer_count":
		return reflect.ValueOf(self._get_composition_layer_count)
	case "_get_composition_layer":
		return reflect.ValueOf(self._get_composition_layer)
	case "_get_composition_layer_order":
		return reflect.ValueOf(self._get_composition_layer_order)
	case "_get_suggested_tracker_names":
		return reflect.ValueOf(self._get_suggested_tracker_names)
	case "_on_register_metadata":
		return reflect.ValueOf(self._on_register_metadata)
	case "_on_before_instance_created":
		return reflect.ValueOf(self._on_before_instance_created)
	case "_on_instance_created":
		return reflect.ValueOf(self._on_instance_created)
	case "_on_instance_destroyed":
		return reflect.ValueOf(self._on_instance_destroyed)
	case "_on_session_created":
		return reflect.ValueOf(self._on_session_created)
	case "_on_process":
		return reflect.ValueOf(self._on_process)
	case "_on_pre_render":
		return reflect.ValueOf(self._on_pre_render)
	case "_on_main_swapchains_created":
		return reflect.ValueOf(self._on_main_swapchains_created)
	case "_on_session_destroyed":
		return reflect.ValueOf(self._on_session_destroyed)
	case "_on_state_idle":
		return reflect.ValueOf(self._on_state_idle)
	case "_on_state_ready":
		return reflect.ValueOf(self._on_state_ready)
	case "_on_state_synchronized":
		return reflect.ValueOf(self._on_state_synchronized)
	case "_on_state_visible":
		return reflect.ValueOf(self._on_state_visible)
	case "_on_state_focused":
		return reflect.ValueOf(self._on_state_focused)
	case "_on_state_stopping":
		return reflect.ValueOf(self._on_state_stopping)
	case "_on_state_loss_pending":
		return reflect.ValueOf(self._on_state_loss_pending)
	case "_on_state_exiting":
		return reflect.ValueOf(self._on_state_exiting)
	case "_on_event_polled":
		return reflect.ValueOf(self._on_event_polled)
	case "_set_viewport_composition_layer_and_get_next_pointer":
		return reflect.ValueOf(self._set_viewport_composition_layer_and_get_next_pointer)
	case "_get_viewport_composition_layer_extension_properties":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_properties)
	case "_get_viewport_composition_layer_extension_property_defaults":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_property_defaults)
	case "_on_viewport_composition_layer_destroyed":
		return reflect.ValueOf(self._on_viewport_composition_layer_destroyed)
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("OpenXRExtensionWrapperExtension", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRExtensionWrapperExtension{*(*gdclass.OpenXRExtensionWrapperExtension)(unsafe.Pointer(&ptr))}
	})
}
