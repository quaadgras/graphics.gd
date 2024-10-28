package Translation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[Translation]s are resources that can be loaded and unloaded on demand. They map a collection of strings to their individual translations, and they also provide convenience methods for pluralization.
	// Translation methods that can be overridden by a [Class] that extends it.
	type Translation interface {
		//Virtual method to override [method get_plural_message].
		GetPluralMessage(src_message string, src_plural_message string, n int, context string) string
		//Virtual method to override [method get_message].
		GetMessage(src_message string, context string) string
	}

*/
type Go [1]classdb.Translation

/*
Virtual method to override [method get_plural_message].
*/
func (Go) _get_plural_message(impl func(ptr unsafe.Pointer, src_message string, src_plural_message string, n int, context string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(src_message)
		var src_plural_message = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(src_plural_message)
		var n = gd.UnsafeGet[gd.Int](p_args,2)
		var context = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,3))
		defer discreet.End(context)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message.String(), src_plural_message.String(), int(n), context.String())
ptr, ok := discreet.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override [method get_message].
*/
func (Go) _get_message(impl func(ptr unsafe.Pointer, src_message string, context string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(src_message)
		var context = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,1))
		defer discreet.End(context)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message.String(), context.String())
ptr, ok := discreet.End(gd.NewStringName(ret))
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
func (self Go) AddMessage(src_message string, xlated_message string) {
	class(self).AddMessage(gd.NewStringName(src_message), gd.NewStringName(xlated_message), gd.NewStringName(""))
}

/*
Adds a message involving plural translation if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
func (self Go) AddPluralMessage(src_message string, xlated_messages []string) {
	class(self).AddPluralMessage(gd.NewStringName(src_message), gd.NewPackedStringSlice(xlated_messages), gd.NewStringName(""))
}

/*
Returns a message's translation.
*/
func (self Go) GetMessage(src_message string) string {
	return string(class(self).GetMessage(gd.NewStringName(src_message), gd.NewStringName("")).String())
}

/*
Returns a message's translation involving plurals.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func (self Go) GetPluralMessage(src_message string, src_plural_message string, n int) string {
	return string(class(self).GetPluralMessage(gd.NewStringName(src_message), gd.NewStringName(src_plural_message), gd.Int(n), gd.NewStringName("")).String())
}

/*
Erases a message.
*/
func (self Go) EraseMessage(src_message string) {
	class(self).EraseMessage(gd.NewStringName(src_message), gd.NewStringName(""))
}

/*
Returns all the messages (keys).
*/
func (self Go) GetMessageList() []string {
	return []string(class(self).GetMessageList().Strings())
}

/*
Returns all the messages (translated text).
*/
func (self Go) GetTranslatedMessageList() []string {
	return []string(class(self).GetTranslatedMessageList().Strings())
}

/*
Returns the number of existing messages.
*/
func (self Go) GetMessageCount() int {
	return int(int(class(self).GetMessageCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Translation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Translation"))
	return Go{classdb.Translation(object)}
}

func (self Go) Locale() string {
		return string(class(self).GetLocale().String())
}

func (self Go) SetLocale(value string) {
	class(self).SetLocale(gd.NewString(value))
}

/*
Virtual method to override [method get_plural_message].
*/
func (class) _get_plural_message(impl func(ptr unsafe.Pointer, src_message gd.StringName, src_plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		var src_plural_message = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,1))
		var n = gd.UnsafeGet[gd.Int](p_args,2)
		var context = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message, src_plural_message, n, context)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override [method get_message].
*/
func (class) _get_message(impl func(ptr unsafe.Pointer, src_message gd.StringName, context gd.StringName) gd.StringName, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var src_message = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		var context = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message, context)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

//go:nosplit
func (self class) SetLocale(locale gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(locale))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_set_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLocale() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a message if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
//go:nosplit
func (self class) AddMessage(src_message gd.StringName, xlated_message gd.StringName, context gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(src_message))
	callframe.Arg(frame, discreet.Get(xlated_message))
	callframe.Arg(frame, discreet.Get(context))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_add_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a message involving plural translation if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
//go:nosplit
func (self class) AddPluralMessage(src_message gd.StringName, xlated_messages gd.PackedStringArray, context gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(src_message))
	callframe.Arg(frame, discreet.Get(xlated_messages))
	callframe.Arg(frame, discreet.Get(context))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_add_plural_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a message's translation.
*/
//go:nosplit
func (self class) GetMessage(src_message gd.StringName, context gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(src_message))
	callframe.Arg(frame, discreet.Get(context))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
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
	callframe.Arg(frame, discreet.Get(src_message))
	callframe.Arg(frame, discreet.Get(src_plural_message))
	callframe.Arg(frame, n)
	callframe.Arg(frame, discreet.Get(context))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_plural_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Erases a message.
*/
//go:nosplit
func (self class) EraseMessage(src_message gd.StringName, context gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(src_message))
	callframe.Arg(frame, discreet.Get(context))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_erase_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all the messages (keys).
*/
//go:nosplit
func (self class) GetMessageList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_message_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns all the messages (translated text).
*/
//go:nosplit
func (self class) GetTranslatedMessageList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_translated_message_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Translation.Bind_get_message_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTranslation() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTranslation() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_plural_message": return reflect.ValueOf(self._get_plural_message);
	case "_get_message": return reflect.ValueOf(self._get_message);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_plural_message": return reflect.ValueOf(self._get_plural_message);
	case "_get_message": return reflect.ValueOf(self._get_message);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Translation", func(ptr gd.Object) any { return classdb.Translation(ptr) })}
