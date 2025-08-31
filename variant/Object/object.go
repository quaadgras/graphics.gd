// Package Object provides methods for working with Object instances.
package Object

import (
	"reflect"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
	"graphics.gd/variant/Error"
	"graphics.gd/variant/Signal"
)

// ID uniquely and opaquely identifies an Object instance.
type ID uint64

// Instance returns the Object instance identified by this ID.
func (id ID) Instance() Instance {
	if id == 0 {
		return Nil
	}
	return Instance([1]gd.Object{gd.PointerMustAssertInstanceID[gd.Object](gdextension.Host.Objects.Lookup(gdextension.ObjectID(id)))})
}

type Notification int

const (
	NotificationPostInitialize    Notification = 0 // Notification received when the object is initialized, before its script is attached. Used internally.
	NotificationPreDelete         Notification = 1 // Notification received when the object is about to be deleted. Can be used like destructors in object-oriented programming languages.
	NotificationExtensionReloaded Notification = 2 // Notification received when the object finishes hot reloading. This notification is only sent for extensions classes and derived.
)

/*
Instance is an advanced Variant type. All classes in the engine inherit from Object. Each class may define new properties, methods or
signals, which are available to all inheriting classes. For example, a Sprite2D instance is able to call Node.add_child because it
inherits from Node.

You can create new instances, using [New].

Objects can have a Script attached to them. Once the Script is instantiated, it effectively acts as an extension to the base class,
allowing it to define and inherit new properties, methods and signals.

Each [Interface] method can be overidden independently:

	type Interface interface {
		Get(property string) any
		GetPropertyList() []ClassDB.PropertyInfo
		Notification(what int, reversed bool)
		PropertyCanRevert(property string) bool
		PropertyGetRevert(property string) any
		Set(property string, value any) bool
		ToString() string
		ValidateProperty(ClassDB.PropertyInfo)
	}
*/
type Instance [1]gdclass.Object

// Extension can be embedded in a struct to create a new class. T should be the type of the struct
// that embeds this Extension.
type Extension[T gdclass.Interface] struct {
	gdclass.Extension[T, Instance]
}

var otype gdextension.ObjectType

func init() {
	gd.Links = append(gd.Links, func() {
		sname := gdextension.Host.Strings.Intern.UTF8("Object")
		otype = gdextension.Host.Objects.Type(sname)
		gdextension.Free(gdextension.TypeStringName, &sname)
	})
}

func (self *Instance) SetObject(obj [1]gdclass.Object) bool {
	if gdextension.Host.Objects.Cast(gdextension.Object(pointers.Get(obj[0])[0]), otype) != 0 {
		self[0] = *(*gdclass.Object)(unsafe.Pointer(&obj))
		return true
	}
	return false
}

func (class *Extension[T]) AsObject() [1]gdclass.Object { return class.Super() }

// Nil is a nil Object instance. Useful for comparisons.
var Nil Instance

type Interface interface {
	Get(property string) any
	GetPropertyList() []PropertyInfo
	Notification(what int, reversed bool)
	PropertyCanRevert(property string) bool
	PropertyGetRevert(property string) any
	Set(property string, value any) bool
	ToString() string
	ValidateProperty(PropertyInfo)
}

type PropertyInfo struct {
	ClassName  string       `gd:"class_name"`
	Name       string       `gd:"name"`
	Hint       int          `gd:"hint"`
	HintString string       `gd:"hint_string"`
	Type       reflect.Type `gd:"type"`
	Usage      int          `gd:"usage"`
}

// New creates a new Object instance.
func New() Instance {
	if !gd.Linked {
		var placeholder Instance
		*(*gd.Object)(unsafe.Pointer(&placeholder)) = pointers.Add[gd.Object]([3]uint64{})
		gd.StartupFunctions = append(gd.StartupFunctions, func() {
			if gd.Linked {
				raw, _ := pointers.End(New().AsObject()[0])
				pointers.Set(*(*gd.Object)(unsafe.Pointer(&placeholder)), raw)
				gd.RegisterCleanup(func() {
					gdextension.Host.Objects.Unsafe.Free(gdextension.Object(raw[0]))
				})
			}
		})
		return placeholder
	}
	return Instance([1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gdextension.Host.Objects.Make(pointers.Get(gd.NewStringName("Object"))))})})
}

func (obj Instance) AsObject() [1]gd.Object          { return obj }
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

// Virtual method lookup.
func (obj Instance) Virtual(name string) reflect.Value { return obj[0].Virtual(name) }

// ClassName returns the object's built-in class name, as a string.
func (obj Instance) ClassName() string {
	return obj[0].GetClass().String()
}

// CanTranslateMessages returns true if the object is allowed to translate messages with tr and tr_n.
// See also [Instance.SetMessageTranslation].
func (obj Instance) CanTranslateMessages() bool {
	return bool(obj[0].CanTranslateMessages())
}

// ID returns the object's unique instance ID. This ID can be saved in EncodedObjectAsID, and can be used
// to retrieve this object instance with [ID.Instance].
func (obj Instance) ID() ID {
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(gdextension.Object(pointers.Get(obj[0])[0]), gdextension.CallReturns[gdextension.ObjectID](&id))
	return ID(id)
}

// ScriptInstance returns the object's Script instance, or false if no script is attached.
func (obj Instance) Script() ([1]gdclass.Script, bool) {
	script, ok := obj[0].GetScript().Interface().([1]gdclass.Script)
	return script, ok
}

// SignalsBlocked returns true if the object is blocking its signals from being emitted.
// See [Instance.SetSignalsBlocked].
func (obj Instance) SignalsBlocked() bool {
	return bool(obj[0].IsBlockingSignals())
}

// NotifyPropertyListChanged emits the property_list_changed signal. This is mainly used to
// refresh the editor, so that the Inspector and editor plugins are properly updated.
func (obj Instance) NotifyPropertyListChanged() {
	obj[0].NotifyPropertyListChanged()
}

// SetBlockSignals if set to true, the object becomes unable to emit signals. Signal connections will
// not work, until it is set to false.
func (obj Instance) SetSignalsBlocked(enable bool) {
	obj[0].SetBlockSignals(enable)
}

// SetMessageTranslation if set to true, allows the object to translate messages with tr and tr_n.
// Enabled by default. See also [Instance.CanTranslateMessages].
func (obj Instance) SetMessageTranslation(enable bool) {
	obj[0].SetMessageTranslation(enable)
}

// SetScript attaches script to the object, and instantiates it. As a result, the script's _init is called.
// A Script is used to extend the object's functionality.
//
// If a script already exists, its instance is detached, and its property values and state are lost. Built-in
// property values are still kept.
func (obj Instance) SetScript(script [1]gdclass.Script) {
	gd.PointerWithOwnershipTransferredToGodot(script[0])
	obj[0].SetScript(gd.NewVariant(script))
}

// String returns a String representing the object. Defaults to "<ClassName#RID>". Override _to_string to
// customize the string representation of the object.
func (obj Instance) String() string {
	return obj[0].ToString().String()
}

// Translate translates a message, using the translation catalogs configured in the Project Settings.
// Note that most Control nodes automatically translate their strings, so this method is mostly useful
// for formatted strings or custom drawn text.
//
// If [Instance.CanTranslateMessages] is false, or no translation is available, this method returns the
// message without changes. See [Instance.SetMessageTranslation].
func (obj Instance) Translate(message string) string {
	return obj[0].Tr(gd.NewStringName(message), gd.NewStringName("")).String()
}

// Translation translates a message or plural_message, using the translation catalogs configured in the Project Settings.
// Further context can be specified to help with the translation.
//
// If [Instance.CanTranslateMessages] is false, or no translation is available, this method returns message or plural_message,
// without changes. See [Instance.SetMessageTranslation].
//
// The n is the number, or amount, of the message's subject. It is used by the translation system to fetch the correct plural
// form for the current language.
//
// For detailed examples, see Localization using gettext.
//
// Note: Negative and float numbers may not properly apply to some countable subjects. It's recommended to handle these cases
// with [Instance.Translate].
//
// Note: This method can't be used without an Object instance, as it requires the [Instance.CanTranslateMessages] method.
// To translate strings in a static context, use [TranslationServer.TranslatePlural].
func (obj Instance) Translation(message string, plural_message string, n int, context string) string {
	return obj[0].TrN(gd.NewStringName(message), gd.NewStringName(plural_message), gd.Int(n), gd.NewStringName(context)).String()
}

// Connect connects a signal by name to a callable. Optional flags can be also added to configure the connection's behavior
// (see Signal.Flags constants).
//
// A signal can only be connected once to the same Callable. If the signal is already connected, this method returns
// Error.InvalidParameter and pushes an error message, unless the signal is connected with [Signal.Weak]. To prevent this,
// use [Instance.IsConnected] first to check for existing connections.
//
// If the callable's object is freed, the connection will be lost.
func (obj Instance) Connect(signal string, callable any, flags ...Signal.Flags) error {
	var all_flags Signal.Flags
	for _, f := range flags {
		all_flags |= f
	}
	err := Error.Code(obj[0].Connect(gd.NewStringName(signal), gd.NewCallable(callable), gd.Int(all_flags)))
	if err != 0 {
		return err
	}
	return nil
}

// IsConnected returns true if a connection exists between the given signal name and callable.
func (obj Instance) IsConnected(signal string, callable any) bool {
	return obj[0].IsConnected(gd.NewStringName(signal), gd.NewCallable(callable))
}

// Use keeps an object alive, preventing it from being garbage collected until the next frame.
func Use(obj interface{ AsObject() [1]gdclass.Object }) { variant.Use(obj.AsObject()[0]) }

// Signal returns the signal with the given name, or a nil signal if it does not exist.
func (obj Instance) Signal(name string) Signal.Any {
	signal := gd.NewSignalOf(obj, gd.NewStringName(name))
	if signal == (gd.Signal{}) {
		return Signal.Nil
	}
	return Signal.Via(gd.SignalProxy{}, pointers.Pack(signal))
}

// HasMethod returns true if the object has a method with the given name.
func (obj Instance) HasMethod(name string) bool {
	return obj[0].HasMethod(gd.NewStringName(name))
}
