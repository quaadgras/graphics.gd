package Translation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Translation]s are resources that can be loaded and unloaded on demand. They map a collection of strings to their individual translations, and they also provide convenience methods for pluralization.
	// Translation methods that can be overridden by a [Class] that extends it.
	type Translation interface {
		//Virtual method to override [method get_plural_message].
		GetPluralMessage(src_message gd.StringName, src_plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName
		//Virtual method to override [method get_message].
		GetMessage(src_message gd.StringName, context gd.StringName) gd.StringName
	}

*/
type Simple [1]classdb.Translation
func (Simple) _get_plural_message(impl func(ptr unsafe.Pointer, src_message string, src_plural_message string, n int, context string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Simple) _get_message(impl func(ptr unsafe.Pointer, src_message string, context string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (self Simple) SetLocale(locale string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLocale(gc.String(locale))
}
func (self Simple) GetLocale() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLocale(gc).String())
}
func (self Simple) AddMessage(src_message string, xlated_message string, context string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMessage(gc.StringName(src_message), gc.StringName(xlated_message), gc.StringName(context))
}
func (self Simple) AddPluralMessage(src_message string, xlated_messages gd.PackedStringArray, context string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPluralMessage(gc.StringName(src_message), xlated_messages, gc.StringName(context))
}
func (self Simple) GetMessage(src_message string, context string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetMessage(gc, gc.StringName(src_message), gc.StringName(context)).String())
}
func (self Simple) GetPluralMessage(src_message string, src_plural_message string, n int, context string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPluralMessage(gc, gc.StringName(src_message), gc.StringName(src_plural_message), gd.Int(n), gc.StringName(context)).String())
}
func (self Simple) EraseMessage(src_message string, context string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseMessage(gc.StringName(src_message), gc.StringName(context))
}
func (self Simple) GetMessageList() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetMessageList(gc))
}
func (self Simple) GetTranslatedMessageList() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetTranslatedMessageList(gc))
}
func (self Simple) GetMessageCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMessageCount()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Translation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsTranslation() Expert { return self[0].AsTranslation() }


//go:nosplit
func (self Simple) AsTranslation() Simple { return self[0].AsTranslation() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_plural_message": return reflect.ValueOf(self._get_plural_message);
	case "_get_message": return reflect.ValueOf(self._get_message);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_get_plural_message": return reflect.ValueOf(self._get_plural_message);
	case "_get_message": return reflect.ValueOf(self._get_message);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Translation", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
