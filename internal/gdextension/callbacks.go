package gdextension

var On Callbacks

type Callbacks struct {
	Engine CallbacksForEngine
	Editor CallbacksForEditor

	MainLoop  CallbacksForMainLoop
	Callables CallbacksForCallables
	Extension CallbacksForExtension
	Threading CallbacksForThreading
}

type CallbacksForEngine struct {
	Init func(level InitializationLevel) `gd:"on_engine_init"`
	Exit func(level InitializationLevel) `gd:"on_engine_exit"`
}

type CallbacksForMainLoop struct {
	FirstFrame func() `gd:"on_first_frame"`
	EveryFrame func() `gd:"on_every_frame"`
	FinalFrame func() `gd:"on_final_frame"`
}

type CallbacksForExtension struct {
	Binding  CallbacksForExtensionBinding
	Instance CallbacksForExtensionInstance
	Class    CallbacksForExtensionClass
	Script   CallbacksForExtensionScript
}

type CallbacksForExtensionBinding struct {
	Created   func(instance ExtensionInstanceID) ExtensionBindingID          `gd:"on_extension_binding_created"`
	Removed   func(instance ExtensionInstanceID, binding ExtensionBindingID) `gd:"on_extension_binding_removed"`
	Reference func(instance ExtensionInstanceID, increment bool) bool        `gd:"on_extension_binding_reference"`
}

type CallbacksForExtensionInstance struct {
	Set                func(instance ExtensionInstanceID, field StringName, value Variant) bool                                                                 `gd:"on_extension_instance_set"`
	Get                func(instance ExtensionInstanceID, field StringName, result Returns[Variant]) bool                                                       `gd:"on_extension_instance_get"`
	PropertyList       func(instance ExtensionInstanceID) PropertyList                                                                                          `gd:"on_extension_instance_property_list"`
	PropertyHasDefault func(instance ExtensionInstanceID, field StringName) bool                                                                                `gd:"on_extension_instance_property_has_default"`
	PropertyGetDefault func(instance ExtensionInstanceID, field StringName, result Returns[Variant]) bool                                                       `gd:"on_extension_instance_property_get_default"`
	PropertyValidation func(instance ExtensionInstanceID, field PropertyList) bool                                                                              `gd:"on_extension_instance_property_validation"`
	Notification       func(instance ExtensionInstanceID, what int32, reverse bool)                                                                             `gd:"on_extension_instance_notification"`
	Stringify          func(instance ExtensionInstanceID) String                                                                                                `gd:"on_extension_instance_stringify"`
	Reference          func(instance ExtensionInstanceID, increment bool) bool                                                                                  `gd:"on_extension_instance_reference"`
	RID                func(instance ExtensionInstanceID, rid Returns[uint64])                                                                                  `gd:"on_extension_instance_rid"`
	CheckedCall        func(instance ExtensionInstanceID, fn FunctionID, result Returns[any], args Accepts[any])                                                `gd:"on_extension_instance_checked_call"`
	VariantCall        func(instance ExtensionInstanceID, fn FunctionID, result Returns[Variant], args Accepts[Variant])                                        `gd:"on_extension_instance_variant_call"`
	DynamicCall        func(instance ExtensionInstanceID, fn FunctionID, result Returns[Variant], arg_count int, args Accepts[Variant], err Returns[CallError]) `gd:"on_extension_instance_dynamic_call"`
	Free               func(instance ExtensionInstanceID)                                                                                                       `gd:"on_extension_instance_free"`
}

type CallbacksForExtensionClass struct {
	Create func(class ExtensionClassID, notify_postinitialize bool) Object         `gd:"on_extension_class_create"`
	Method func(class ExtensionClassID, method StringName, hash uint32) FunctionID `gd:"on_extension_class_method"`
}

type CallbacksForExtensionScript struct {
	Categorization      func(instance ExtensionInstanceID, into PropertyList) bool      `gd:"on_extension_script_categorization"`
	PropertyType        func(field StringName, err Returns[CallError]) VariantType      `gd:"on_extension_script_get_property_type"`
	Owner               func(instance ExtensionInstanceID) Object                       `gd:"on_extension_script_get_owner"`
	PropertyState       func(instance ExtensionInstanceID, add FunctionID, arg Pointer) `gd:"on_extension_script_get_property_state"`
	Methods             func(instance ExtensionInstanceID) MethodList                   `gd:"on_extension_script_get_methods"`
	HasMethod           func(instance ExtensionInstanceID, method StringName) bool      `gd:"on_extension_script_has_method"`
	MethodArgumentCount func(instance ExtensionInstanceID, method StringName) int       `gd:"on_extension_script_get_method_argument_count"`
	Get                 func(instance ExtensionInstanceID) Object                       `gd:"on_extension_script_get"`
	IsPlaceholder       func(instance ExtensionInstanceID) bool                         `gd:"on_extension_script_is_placeholder"`
	Language            func(instance ExtensionInstanceID) Object                       `gd:"on_extension_script_get_language"`
}

type CallbacksForCallables struct {
	Call          func(fn FunctionID, result Returns[Variant], arg_count int, args Accepts[Variant], err Returns[CallError]) `gd:"on_callable_call"`
	Validation    func(fn FunctionID) bool                                                                                   `gd:"on_callable_validation"`
	Free          func(fn FunctionID)                                                                                        `gd:"on_callable_free"`
	Hash          func(fn FunctionID) uint32                                                                                 `gd:"on_callable_hash"`
	Compare       func(fn FunctionID, other FunctionID) bool                                                                 `gd:"on_callable_compare"`
	LessThan      func(fn FunctionID, other FunctionID) bool                                                                 `gd:"on_callable_less_than"`
	Stringify     func(fn FunctionID, err Returns[CallError]) String                                                         `gd:"on_callable_stringify"`
	ArgumentCount func(fn FunctionID, err Returns[CallError]) int                                                            `gd:"on_callable_get_argument_count"`
}

type CallbacksForEditor struct {
	ClassInUseDetection func(classes PackedArray[String], result Returns[PackedArray[String]]) `gd:"on_editor_class_in_use_detection"`
}

type CallbacksForThreading struct {
	Run        func(task TaskID)           `gd:"on_worker_thread_pool_task"`
	RunInGroup func(task TaskID, n uint32) `gd:"on_worker_thread_pool_group_task"`
}
