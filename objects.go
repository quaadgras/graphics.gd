package gd

import internal "grow.graphics/gd/internal"

// NotificationType sent to an Object's notification method.
type NotificationType int32

const (
	// Notification received when the object is initialized, before its script is attached. Used internally.
	NotificationPostInitialize NotificationType = iota
	// Notification received when the object is about to be deleted. Can act as the deconstructor of some programming languages.
	NotificationPreDelete
	// Notification received when the object finishes hot reloading. This notification is only sent for extensions classes and derived.
	NotificationExtensionReloaded
)

/*
Object is an advanced [Variant] type. All classes in the engine inherit from [Object]. Each class may define
new properties, methods or signals, which are available to all inheriting classes. For example, a [Sprite2D]
instance is able to call [Node.AddChild] because it inherits from [Node].

You can create new instances, using [Create].

To delete an [Object] instance, call [Object.Free]. This is necessary for most classes inheriting [Object],
because they do not manage memory on their own, and will otherwise cause memory leaks when no longer in use.
There are a few classes that perform memory management. For example, [RefCounted] (and by extension [Resource])
deletes itself when no longer referenced, and [Node] deletes its children when freed.

Objects can have a [Script] attached to them. Once the [Script] is instantiated, it effectively acts as an
extension to the base class, allowing it to define and inherit new properties, methods and signals.

Inside a [Script], GetPropertyList may be overridden to customize properties in several ways. This allows them
to be available to the editor, display as lists of options, sub-divide into groups, save on disk, etc. Scripting
languages offer easier ways to customize properties.

Godot is very dynamic. An object's script, and therefore its properties, methods and signals, can be changed at
run-time. Because of this, there can be occasions where, for example, a property required by a method may not
exist. To prevent run-time errors, see methods such as [Object.Set], [Object.Get], [Object.Call],
[Object.HasMethod], [Object.HasSignal], etc. Note that these methods are much slower than direct references.

Notifications are int constants commonly sent and received by objects. For example, on every rendered frame,
the [SceneTree] notifies nodes inside the tree with a [NotificationProcess]. The nodes receive it and may call
Process to update. To make use of notifications, see [Object.Notification].

Lastly, every object can also contain metadata (data about data). [Object.SetMeta] can be useful to store
information that the object itself does not depend on. To keep your code clean, making excessive use of metadata
is discouraged.

Note: Unlike references to a [RefCounted], references to an object stored in a variable can become invalid
without being set to null. To check if an object has been deleted, do not compare it against null. Instead,
se [Context.IsInstanceValid]. It's also recommended to inherit from [RefCounted] for classes storing data
instead of [Object].

Note: The script is not exposed like most properties. To set or get an object's [Script] in code, use
[Object.SetScript] and [Object.GetScript], respectively.

	// Object methods that can be overridden by a [Class] that extends it.
	type Object interface {
		// Override this method to customize the behavior of get. Should return the given property's value,
		// or null if the property should be handled normally.
		//
		// Combined with Set and GetPropertyList, this method allows defining custom properties, which is
		// particularly useful for editor plugins. Note that a property must be present in GetPropertyList,
		// otherwise this method will not be called.
		Get(Context, StringName) Variant
		// Override this method to customize how script properties should be handled by the engine.
		GetPropertyList(Context) []gd.PropertyInfo
		// Called when the object's script is instantiated, oftentimes after the object is initialized in
		// memory (through [Create]). It can be also defined to take in parameters. This method is similar
		// to a constructor in most programming languages.
		Init(Context)
		// Called when the object receives a notification, which can be identified in what by comparing it
		// with a constant. See also [Object.Notification].
		Notification(Context, NotificationType)
		// Override this method to customize the given property's revert behavior. Should return true if the
		// property can be reverted in the Inspector dock. Use PropertyGetRevert to specify the property's
		// default value.
		PropertyCanRevert(Context, StringName) Bool
		// Override this method to customize the given property's revert behavior. Should return the default
		// value for the property. If the default value differs from the property's current value, a revert
		// icon is displayed in the Inspector dock.
		PropertyGetRevert(Context, StringName) Variant
		// Override this method to customize the behavior of set. Should set the property to value and return
		// true, or false if the property should be handled normally. The exact way to set the property is up
		// to this method's implementation.
		//
		// Combined with Get and GetPropertyList, this method allows defining custom properties, which is
		// particularly useful for editor plugins. Note that a property must be present in GetPropertyList,
		// otherwise this method will not be called.
		Set(Context, StringName, Variant) Bool
		// Override this method to customize the return value of [Object.ToString], and therefore the object's
		// representation as a String.
		ToString(Context) String
		// Override this method to customize existing properties. Every property info goes through this method.
		ValidateProperty(Context, StringName, gd.PropertyInfo)
	}
*/
type Object = internal.Object
