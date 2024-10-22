package OpenXRExtensionWrapperExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[OpenXRExtensionWrapperExtension] allows clients to implement OpenXR extensions with GDExtension. The extension should be registered with [method register_extension_wrapper].
	// OpenXRExtensionWrapperExtension methods that can be overridden by a [Class] that extends it.
	type OpenXRExtensionWrapperExtension interface {
		//Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
		//- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
		//- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
		GetRequestedExtensions() gd.Dictionary
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
		SetViewportCompositionLayerAndGetNextPointer(layer unsafe.Pointer, property_values gd.Dictionary, next_pointer unsafe.Pointer) int
		//Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
		GetViewportCompositionLayerExtensionProperties() gd.ArrayOf[gd.Dictionary]
		//Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
		GetViewportCompositionLayerExtensionPropertyDefaults() gd.Dictionary
		//Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
		//[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
		OnViewportCompositionLayerDestroyed(layer unsafe.Pointer) 
	}

*/
type Go [1]classdb.OpenXRExtensionWrapperExtension

/*
Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
*/
func (Go) _get_requested_extensions(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}

/*
Adds additional data structures when interogating OpenXR system abilities.
*/
func (Go) _set_system_properties_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Adds additional data structures when the OpenXR instance is created.
*/
func (Go) _set_instance_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Adds additional data structures when the OpenXR session is created.
*/
func (Go) _set_session_create_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Adds additional data structures when creating OpenXR swapchains.
*/
func (Go) _set_swapchain_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Adds additional data structures when each hand tracker is created.
*/
func (Go) _set_hand_joint_locations_and_get_next_pointer(impl func(ptr unsafe.Pointer, hand_index int, next_pointer unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var hand_index = gd.UnsafeGet[gd.Int](p_args,0)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(hand_index), next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Go) _get_composition_layer_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Go) _get_composition_layer(impl func(ptr unsafe.Pointer, index int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Go) _get_composition_layer_order(impl func(ptr unsafe.Pointer, index int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
*/
func (Go) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
*/
func (Go) _on_register_metadata(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called before the OpenXR instance is created.
*/
func (Go) _on_before_instance_created(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called right after the OpenXR instance is created.
*/
func (Go) _on_instance_created(impl func(ptr unsafe.Pointer, instance int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var instance = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(instance))
		gc.End()
	}
}

/*
Called right before the OpenXR instance is destroyed.
*/
func (Go) _on_instance_destroyed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called right after the OpenXR session is created.
*/
func (Go) _on_session_created(impl func(ptr unsafe.Pointer, session int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var session = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(session))
		gc.End()
	}
}

/*
Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
*/
func (Go) _on_process(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called right before the XR viewports begin their rendering step.
*/
func (Go) _on_pre_render(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called right after the main swapchains are (re)created.
*/
func (Go) _on_main_swapchains_created(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called right before the OpenXR session is destroyed.
*/
func (Go) _on_session_destroyed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to idle.
*/
func (Go) _on_state_idle(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
*/
func (Go) _on_state_ready(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
*/
func (Go) _on_state_synchronized(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
*/
func (Go) _on_state_visible(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
*/
func (Go) _on_state_focused(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to stopping.
*/
func (Go) _on_state_stopping(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to loss pending.
*/
func (Go) _on_state_loss_pending(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the OpenXR session state is changed to exiting.
*/
func (Go) _on_state_exiting(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
*/
func (Go) _on_event_polled(impl func(ptr unsafe.Pointer, event unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var event = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (Go) _set_viewport_composition_layer_and_get_next_pointer(impl func(ptr unsafe.Pointer, layer unsafe.Pointer, property_values gd.Dictionary, next_pointer unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var property_values = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, property_values, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
*/
func (Go) _get_viewport_composition_layer_extension_properties(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}

/*
Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
*/
func (Go) _get_viewport_composition_layer_extension_property_defaults(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}

/*
Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (Go) _on_viewport_composition_layer_destroyed(impl func(ptr unsafe.Pointer, layer unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, layer)
		gc.End()
	}
}

/*
Returns the created [OpenXRAPIExtension], which can be used to access the OpenXR API.
*/
func (self Go) GetOpenxrApi() gdclass.OpenXRAPIExtension {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.OpenXRAPIExtension(class(self).GetOpenxrApi(gc))
}

/*
Registers the extension. This should happen at core module initialization level.
*/
func (self Go) RegisterExtensionWrapper() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RegisterExtensionWrapper()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRExtensionWrapperExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRExtensionWrapperExtension"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
*/
func (class) _get_requested_extensions(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Adds additional data structures when interogating OpenXR system abilities.
*/
func (class) _set_system_properties_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Adds additional data structures when the OpenXR instance is created.
*/
func (class) _set_instance_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Adds additional data structures when the OpenXR session is created.
*/
func (class) _set_session_create_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Adds additional data structures when creating OpenXR swapchains.
*/
func (class) _set_swapchain_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Adds additional data structures when each hand tracker is created.
*/
func (class) _set_hand_joint_locations_and_get_next_pointer(impl func(ptr unsafe.Pointer, hand_index gd.Int, next_pointer unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var hand_index = gd.UnsafeGet[gd.Int](p_args,0)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, hand_index, next_pointer)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer_order(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
*/
func (class) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
*/
func (class) _on_register_metadata(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called before the OpenXR instance is created.
*/
func (class) _on_before_instance_created(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called right after the OpenXR instance is created.
*/
func (class) _on_instance_created(impl func(ptr unsafe.Pointer, instance gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var instance = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, instance)
		ctx.End()
	}
}

/*
Called right before the OpenXR instance is destroyed.
*/
func (class) _on_instance_destroyed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called right after the OpenXR session is created.
*/
func (class) _on_session_created(impl func(ptr unsafe.Pointer, session gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var session = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, session)
		ctx.End()
	}
}

/*
Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
*/
func (class) _on_process(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called right before the XR viewports begin their rendering step.
*/
func (class) _on_pre_render(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called right after the main swapchains are (re)created.
*/
func (class) _on_main_swapchains_created(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called right before the OpenXR session is destroyed.
*/
func (class) _on_session_destroyed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to idle.
*/
func (class) _on_state_idle(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
*/
func (class) _on_state_ready(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
*/
func (class) _on_state_synchronized(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
*/
func (class) _on_state_visible(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
*/
func (class) _on_state_focused(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to stopping.
*/
func (class) _on_state_stopping(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to loss pending.
*/
func (class) _on_state_loss_pending(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the OpenXR session state is changed to exiting.
*/
func (class) _on_state_exiting(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
*/
func (class) _on_event_polled(impl func(ptr unsafe.Pointer, event unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var event = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (class) _set_viewport_composition_layer_and_get_next_pointer(impl func(ptr unsafe.Pointer, layer unsafe.Pointer, property_values gd.Dictionary, next_pointer unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		var property_values = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, property_values, next_pointer)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
*/
func (class) _get_viewport_composition_layer_extension_properties(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
*/
func (class) _get_viewport_composition_layer_extension_property_defaults(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (class) _on_viewport_composition_layer_destroyed(impl func(ptr unsafe.Pointer, layer unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, layer)
		ctx.End()
	}
}

/*
Returns the created [OpenXRAPIExtension], which can be used to access the OpenXR API.
*/
//go:nosplit
func (self class) GetOpenxrApi(ctx gd.Lifetime) gdclass.OpenXRAPIExtension {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRExtensionWrapperExtension.Bind_get_openxr_api, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRAPIExtension
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Registers the extension. This should happen at core module initialization level.
*/
//go:nosplit
func (self class) RegisterExtensionWrapper()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRExtensionWrapperExtension.Bind_register_extension_wrapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRExtensionWrapperExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRExtensionWrapperExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_requested_extensions": return reflect.ValueOf(self._get_requested_extensions);
	case "_set_system_properties_and_get_next_pointer": return reflect.ValueOf(self._set_system_properties_and_get_next_pointer);
	case "_set_instance_create_info_and_get_next_pointer": return reflect.ValueOf(self._set_instance_create_info_and_get_next_pointer);
	case "_set_session_create_and_get_next_pointer": return reflect.ValueOf(self._set_session_create_and_get_next_pointer);
	case "_set_swapchain_create_info_and_get_next_pointer": return reflect.ValueOf(self._set_swapchain_create_info_and_get_next_pointer);
	case "_set_hand_joint_locations_and_get_next_pointer": return reflect.ValueOf(self._set_hand_joint_locations_and_get_next_pointer);
	case "_get_composition_layer_count": return reflect.ValueOf(self._get_composition_layer_count);
	case "_get_composition_layer": return reflect.ValueOf(self._get_composition_layer);
	case "_get_composition_layer_order": return reflect.ValueOf(self._get_composition_layer_order);
	case "_get_suggested_tracker_names": return reflect.ValueOf(self._get_suggested_tracker_names);
	case "_on_register_metadata": return reflect.ValueOf(self._on_register_metadata);
	case "_on_before_instance_created": return reflect.ValueOf(self._on_before_instance_created);
	case "_on_instance_created": return reflect.ValueOf(self._on_instance_created);
	case "_on_instance_destroyed": return reflect.ValueOf(self._on_instance_destroyed);
	case "_on_session_created": return reflect.ValueOf(self._on_session_created);
	case "_on_process": return reflect.ValueOf(self._on_process);
	case "_on_pre_render": return reflect.ValueOf(self._on_pre_render);
	case "_on_main_swapchains_created": return reflect.ValueOf(self._on_main_swapchains_created);
	case "_on_session_destroyed": return reflect.ValueOf(self._on_session_destroyed);
	case "_on_state_idle": return reflect.ValueOf(self._on_state_idle);
	case "_on_state_ready": return reflect.ValueOf(self._on_state_ready);
	case "_on_state_synchronized": return reflect.ValueOf(self._on_state_synchronized);
	case "_on_state_visible": return reflect.ValueOf(self._on_state_visible);
	case "_on_state_focused": return reflect.ValueOf(self._on_state_focused);
	case "_on_state_stopping": return reflect.ValueOf(self._on_state_stopping);
	case "_on_state_loss_pending": return reflect.ValueOf(self._on_state_loss_pending);
	case "_on_state_exiting": return reflect.ValueOf(self._on_state_exiting);
	case "_on_event_polled": return reflect.ValueOf(self._on_event_polled);
	case "_set_viewport_composition_layer_and_get_next_pointer": return reflect.ValueOf(self._set_viewport_composition_layer_and_get_next_pointer);
	case "_get_viewport_composition_layer_extension_properties": return reflect.ValueOf(self._get_viewport_composition_layer_extension_properties);
	case "_get_viewport_composition_layer_extension_property_defaults": return reflect.ValueOf(self._get_viewport_composition_layer_extension_property_defaults);
	case "_on_viewport_composition_layer_destroyed": return reflect.ValueOf(self._on_viewport_composition_layer_destroyed);
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_requested_extensions": return reflect.ValueOf(self._get_requested_extensions);
	case "_set_system_properties_and_get_next_pointer": return reflect.ValueOf(self._set_system_properties_and_get_next_pointer);
	case "_set_instance_create_info_and_get_next_pointer": return reflect.ValueOf(self._set_instance_create_info_and_get_next_pointer);
	case "_set_session_create_and_get_next_pointer": return reflect.ValueOf(self._set_session_create_and_get_next_pointer);
	case "_set_swapchain_create_info_and_get_next_pointer": return reflect.ValueOf(self._set_swapchain_create_info_and_get_next_pointer);
	case "_set_hand_joint_locations_and_get_next_pointer": return reflect.ValueOf(self._set_hand_joint_locations_and_get_next_pointer);
	case "_get_composition_layer_count": return reflect.ValueOf(self._get_composition_layer_count);
	case "_get_composition_layer": return reflect.ValueOf(self._get_composition_layer);
	case "_get_composition_layer_order": return reflect.ValueOf(self._get_composition_layer_order);
	case "_get_suggested_tracker_names": return reflect.ValueOf(self._get_suggested_tracker_names);
	case "_on_register_metadata": return reflect.ValueOf(self._on_register_metadata);
	case "_on_before_instance_created": return reflect.ValueOf(self._on_before_instance_created);
	case "_on_instance_created": return reflect.ValueOf(self._on_instance_created);
	case "_on_instance_destroyed": return reflect.ValueOf(self._on_instance_destroyed);
	case "_on_session_created": return reflect.ValueOf(self._on_session_created);
	case "_on_process": return reflect.ValueOf(self._on_process);
	case "_on_pre_render": return reflect.ValueOf(self._on_pre_render);
	case "_on_main_swapchains_created": return reflect.ValueOf(self._on_main_swapchains_created);
	case "_on_session_destroyed": return reflect.ValueOf(self._on_session_destroyed);
	case "_on_state_idle": return reflect.ValueOf(self._on_state_idle);
	case "_on_state_ready": return reflect.ValueOf(self._on_state_ready);
	case "_on_state_synchronized": return reflect.ValueOf(self._on_state_synchronized);
	case "_on_state_visible": return reflect.ValueOf(self._on_state_visible);
	case "_on_state_focused": return reflect.ValueOf(self._on_state_focused);
	case "_on_state_stopping": return reflect.ValueOf(self._on_state_stopping);
	case "_on_state_loss_pending": return reflect.ValueOf(self._on_state_loss_pending);
	case "_on_state_exiting": return reflect.ValueOf(self._on_state_exiting);
	case "_on_event_polled": return reflect.ValueOf(self._on_event_polled);
	case "_set_viewport_composition_layer_and_get_next_pointer": return reflect.ValueOf(self._set_viewport_composition_layer_and_get_next_pointer);
	case "_get_viewport_composition_layer_extension_properties": return reflect.ValueOf(self._get_viewport_composition_layer_extension_properties);
	case "_get_viewport_composition_layer_extension_property_defaults": return reflect.ValueOf(self._get_viewport_composition_layer_extension_property_defaults);
	case "_on_viewport_composition_layer_destroyed": return reflect.ValueOf(self._on_viewport_composition_layer_destroyed);
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("OpenXRExtensionWrapperExtension", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
