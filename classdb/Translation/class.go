// Package Translation provides methods for working with Translation object instances.
package Translation

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
[Translation]s are resources that can be loaded and unloaded on demand. They map a collection of strings to their individual translations, and they also provide convenience methods for pluralization.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=Translation)
*/
type Instance [1]gdclass.Translation

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTranslation() Instance
}
type Interface interface {
	//Virtual method to override [method get_plural_message].
	GetPluralMessage(src_message string, src_plural_message string, n int, context string) string
	//Virtual method to override [method get_message].
	GetMessage(src_message string, context string) string
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetPluralMessage(src_message string, src_plural_message string, n int, context string) (_ string) {
	return
}
func (self implementation) GetMessage(src_message string, context string) (_ string) { return }

/*
Virtual method to override [method get_plural_message].
*/
func (Instance) _get_plural_message(impl func(ptr unsafe.Pointer, src_message string, src_plural_message string, n int, context string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(src_message)
		var src_plural_message = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(src_plural_message)
		var n = gd.UnsafeGet[gd.Int](p_args, 2)
		var context = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		defer pointers.End(context)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message.String(), src_plural_message.String(), int(n), context.String())
		ptr, ok := pointers.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override [method get_message].
*/
func (Instance) _get_message(impl func(ptr unsafe.Pointer, src_message string, context string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(src_message)
		var context = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(context)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message.String(), context.String())
		ptr, ok := pointers.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Adds a message if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
func (self Instance) AddMessage(src_message string, xlated_message string) {
	class(self).AddMessage(gd.NewStringName(src_message), gd.NewStringName(xlated_message), gd.NewStringName(""))
}

/*
Adds a message involving plural translation if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
func (self Instance) AddPluralMessage(src_message string, xlated_messages []string) {
	class(self).AddPluralMessage(gd.NewStringName(src_message), gd.NewPackedStringSlice(xlated_messages), gd.NewStringName(""))
}

/*
Returns a message's translation.
*/
func (self Instance) GetMessage(src_message string) string {
	return string(class(self).GetMessage(gd.NewStringName(src_message), gd.NewStringName("")).String())
}

/*
Returns a message's translation involving plurals.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func (self Instance) GetPluralMessage(src_message string, src_plural_message string, n int) string {
	return string(class(self).GetPluralMessage(gd.NewStringName(src_message), gd.NewStringName(src_plural_message), gd.Int(n), gd.NewStringName("")).String())
}

/*
Erases a message.
*/
func (self Instance) EraseMessage(src_message string) {
	class(self).EraseMessage(gd.NewStringName(src_message), gd.NewStringName(""))
}

/*
Returns all the messages (keys).
*/
func (self Instance) GetMessageList() []string {
	return []string(class(self).GetMessageList().Strings())
}

/*
Returns all the messages (translated text).
*/
func (self Instance) GetTranslatedMessageList() []string {
	return []string(class(self).GetTranslatedMessageList().Strings())
}

/*
Returns the number of existing messages.
*/
func (self Instance) GetMessageCount() int {
	return int(int(class(self).GetMessageCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Translation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Translation"))
	casted := Instance{*(*gdclass.Translation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Locale() string {
	return string(class(self).GetLocale().String())
}

func (self Instance) SetLocale(value string) {
	class(self).SetLocale(gd.NewString(value))
}

/*
Virtual method to override [method get_plural_message].
*/
func (class) _get_plural_message(impl func(ptr unsafe.Pointer, src_message gd.StringName, src_plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var src_plural_message = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		var n = gd.UnsafeGet[gd.Int](p_args, 2)
		var context = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message, src_plural_message, n, context)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override [method get_message].
*/
func (class) _get_message(impl func(ptr unsafe.Pointer, src_message gd.StringName, context gd.StringName) gd.StringName) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var context = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message, context)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

//go:nosplit
func (self class) SetLocale(locale gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(locale))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_set_locale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLocale() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a message if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
//go:nosplit
func (self class) AddMessage(src_message gd.StringName, xlated_message gd.StringName, context gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src_message))
	callframe.Arg(frame, pointers.Get(xlated_message))
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_add_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a message involving plural translation if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
//go:nosplit
func (self class) AddPluralMessage(src_message gd.StringName, xlated_messages gd.PackedStringArray, context gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src_message))
	callframe.Arg(frame, pointers.Get(xlated_messages))
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_add_plural_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a message's translation.
*/
//go:nosplit
func (self class) GetMessage(src_message gd.StringName, context gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src_message))
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a message's translation involving plurals.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
//go:nosplit
func (self class) GetPluralMessage(src_message gd.StringName, src_plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src_message))
	callframe.Arg(frame, pointers.Get(src_plural_message))
	callframe.Arg(frame, n)
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_plural_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Erases a message.
*/
//go:nosplit
func (self class) EraseMessage(src_message gd.StringName, context gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src_message))
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_erase_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns all the messages (keys).
*/
//go:nosplit
func (self class) GetMessageList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_message_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns all the messages (translated text).
*/
//go:nosplit
func (self class) GetTranslatedMessageList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_translated_message_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of existing messages.
*/
//go:nosplit
func (self class) GetMessageCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_message_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTranslation() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTranslation() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_get_plural_message":
		return reflect.ValueOf(self._get_plural_message)
	case "_get_message":
		return reflect.ValueOf(self._get_message)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_plural_message":
		return reflect.ValueOf(self._get_plural_message)
	case "_get_message":
		return reflect.ValueOf(self._get_message)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Translation", func(ptr gd.Object) any { return [1]gdclass.Translation{*(*gdclass.Translation)(unsafe.Pointer(&ptr))} })
}
