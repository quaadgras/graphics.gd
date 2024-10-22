package Translation

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
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var src_message = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var src_plural_message = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var n = gd.UnsafeGet[gd.Int](p_args,2)
		var context = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message.String(), src_plural_message.String(), int(n), context.String())
		gd.UnsafeSet(p_back, mmm.End(gc.StringName(ret)))
		gc.End()
	}
}

/*
Virtual method to override [method get_message].
*/
func (Go) _get_message(impl func(ptr unsafe.Pointer, src_message string, context string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var src_message = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var context = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message.String(), context.String())
		gd.UnsafeSet(p_back, mmm.End(gc.StringName(ret)))
		gc.End()
	}
}

/*
Adds a message if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
func (self Go) AddMessage(src_message string, xlated_message string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddMessage(gc.StringName(src_message), gc.StringName(xlated_message), gc.StringName(""))
}

/*
Adds a message involving plural translation if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
func (self Go) AddPluralMessage(src_message string, xlated_messages []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddPluralMessage(gc.StringName(src_message), gc.PackedStringSlice(xlated_messages), gc.StringName(""))
}

/*
Returns a message's translation.
*/
func (self Go) GetMessage(src_message string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetMessage(gc, gc.StringName(src_message), gc.StringName("")).String())
}

/*
Returns a message's translation involving plurals.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func (self Go) GetPluralMessage(src_message string, src_plural_message string, n int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetPluralMessage(gc, gc.StringName(src_message), gc.StringName(src_plural_message), gd.Int(n), gc.StringName("")).String())
}

/*
Erases a message.
*/
func (self Go) EraseMessage(src_message string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EraseMessage(gc.StringName(src_message), gc.StringName(""))
}

/*
Returns all the messages (keys).
*/
func (self Go) GetMessageList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetMessageList(gc).Strings(gc))
}

/*
Returns all the messages (translated text).
*/
func (self Go) GetTranslatedMessageList() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetTranslatedMessageList(gc).Strings(gc))
}

/*
Returns the number of existing messages.
*/
func (self Go) GetMessageCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMessageCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Translation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Translation"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Locale() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLocale(gc).String())
}

func (self Go) SetLocale(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLocale(gc.String(value))
}

/*
Virtual method to override [method get_plural_message].
*/
func (class) _get_plural_message(impl func(ptr unsafe.Pointer, src_message gd.StringName, src_plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var src_message = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var src_plural_message = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var n = gd.UnsafeGet[gd.Int](p_args,2)
		var context = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message, src_plural_message, n, context)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method to override [method get_message].
*/
func (class) _get_message(impl func(ptr unsafe.Pointer, src_message gd.StringName, context gd.StringName) gd.StringName, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var src_message = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var context = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, src_message, context)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

//go:nosplit
func (self class) SetLocale(locale gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_set_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLocale(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a message if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
//go:nosplit
func (self class) AddMessage(src_message gd.StringName, xlated_message gd.StringName, context gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src_message))
	callframe.Arg(frame, mmm.Get(xlated_message))
	callframe.Arg(frame, mmm.Get(context))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_add_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a message involving plural translation if nonexistent, followed by its translation.
An additional context could be used to specify the translation context or differentiate polysemic words.
*/
//go:nosplit
func (self class) AddPluralMessage(src_message gd.StringName, xlated_messages gd.PackedStringArray, context gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src_message))
	callframe.Arg(frame, mmm.Get(xlated_messages))
	callframe.Arg(frame, mmm.Get(context))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_add_plural_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a message's translation.
*/
//go:nosplit
func (self class) GetMessage(ctx gd.Lifetime, src_message gd.StringName, context gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src_message))
	callframe.Arg(frame, mmm.Get(context))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_get_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a message's translation involving plurals.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
//go:nosplit
func (self class) GetPluralMessage(ctx gd.Lifetime, src_message gd.StringName, src_plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src_message))
	callframe.Arg(frame, mmm.Get(src_plural_message))
	callframe.Arg(frame, n)
	callframe.Arg(frame, mmm.Get(context))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_get_plural_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Erases a message.
*/
//go:nosplit
func (self class) EraseMessage(src_message gd.StringName, context gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(src_message))
	callframe.Arg(frame, mmm.Get(context))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_erase_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all the messages (keys).
*/
//go:nosplit
func (self class) GetMessageList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_get_message_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns all the messages (translated text).
*/
//go:nosplit
func (self class) GetTranslatedMessageList(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_get_translated_message_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of existing messages.
*/
//go:nosplit
func (self class) GetMessageCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Translation.Bind_get_message_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("Translation", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
