package TranslationServer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The server that manages all language translations. Translations can be added to or removed from it.

*/
type Simple [1]classdb.TranslationServer
func (self Simple) SetLocale(locale string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLocale(gc.String(locale))
}
func (self Simple) GetLocale() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLocale(gc).String())
}
func (self Simple) GetToolLocale() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetToolLocale(gc).String())
}
func (self Simple) CompareLocales(locale_a string, locale_b string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).CompareLocales(gc.String(locale_a), gc.String(locale_b))))
}
func (self Simple) StandardizeLocale(locale string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).StandardizeLocale(gc, gc.String(locale)).String())
}
func (self Simple) GetAllLanguages() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetAllLanguages(gc))
}
func (self Simple) GetLanguageName(language string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLanguageName(gc, gc.String(language)).String())
}
func (self Simple) GetAllScripts() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetAllScripts(gc))
}
func (self Simple) GetScriptName(script string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetScriptName(gc, gc.String(script)).String())
}
func (self Simple) GetAllCountries() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetAllCountries(gc))
}
func (self Simple) GetCountryName(country string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCountryName(gc, gc.String(country)).String())
}
func (self Simple) GetLocaleName(locale string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLocaleName(gc, gc.String(locale)).String())
}
func (self Simple) Translate(message string, context string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).Translate(gc, gc.StringName(message), gc.StringName(context)).String())
}
func (self Simple) TranslatePlural(message string, plural_message string, n int, context string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).TranslatePlural(gc, gc.StringName(message), gc.StringName(plural_message), gd.Int(n), gc.StringName(context)).String())
}
func (self Simple) AddTranslation(translation [1]classdb.Translation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddTranslation(translation)
}
func (self Simple) RemoveTranslation(translation [1]classdb.Translation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTranslation(translation)
}
func (self Simple) GetTranslationObject(locale string) [1]classdb.Translation {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Translation(Expert(self).GetTranslationObject(gc, gc.String(locale)))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) GetLoadedLocales() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetLoadedLocales(gc))
}
func (self Simple) IsPseudolocalizationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPseudolocalizationEnabled())
}
func (self Simple) SetPseudolocalizationEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPseudolocalizationEnabled(enabled)
}
func (self Simple) ReloadPseudolocalization() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReloadPseudolocalization()
}
func (self Simple) Pseudolocalize(message string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).Pseudolocalize(gc, gc.StringName(message)).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TranslationServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the locale of the project. The [param locale] string will be standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
If translations have been loaded beforehand for the new locale, they will be applied.
*/
//go:nosplit
func (self class) SetLocale(locale gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_set_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current locale of the project.
See also [method OS.get_locale] and [method OS.get_locale_language] to query the locale of the user system.
*/
//go:nosplit
func (self class) GetLocale(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current locale of the editor.
[b]Note:[/b] When called from an exported project returns the same value as [method get_locale].
*/
//go:nosplit
func (self class) GetToolLocale(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_tool_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Compares two locales and returns a similarity score between [code]0[/code] (no match) and [code]10[/code] (full match).
*/
//go:nosplit
func (self class) CompareLocales(locale_a gd.String, locale_b gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale_a))
	callframe.Arg(frame, mmm.Get(locale_b))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_compare_locales, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
*/
//go:nosplit
func (self class) StandardizeLocale(ctx gd.Lifetime, locale gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_standardize_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns array of known language codes.
*/
//go:nosplit
func (self class) GetAllLanguages(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_all_languages, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a readable language name for the [param language] code.
*/
//go:nosplit
func (self class) GetLanguageName(ctx gd.Lifetime, language gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_language_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of known script codes.
*/
//go:nosplit
func (self class) GetAllScripts(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_all_scripts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a readable script name for the [param script] code.
*/
//go:nosplit
func (self class) GetScriptName(ctx gd.Lifetime, script gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(script))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_script_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array of known country codes.
*/
//go:nosplit
func (self class) GetAllCountries(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_all_countries, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a readable country name for the [param country] code.
*/
//go:nosplit
func (self class) GetCountryName(ctx gd.Lifetime, country gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(country))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_country_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
*/
//go:nosplit
func (self class) GetLocaleName(ctx gd.Lifetime, locale gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_locale_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current locale's translation for the given message (key) and context.
*/
//go:nosplit
func (self class) Translate(ctx gd.Lifetime, message gd.StringName, context gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	callframe.Arg(frame, mmm.Get(context))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current locale's translation for the given message (key), plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
//go:nosplit
func (self class) TranslatePlural(ctx gd.Lifetime, message gd.StringName, plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	callframe.Arg(frame, mmm.Get(plural_message))
	callframe.Arg(frame, n)
	callframe.Arg(frame, mmm.Get(context))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_translate_plural, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds a [Translation] resource.
*/
//go:nosplit
func (self class) AddTranslation(translation object.Translation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(translation[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_add_translation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given translation from the server.
*/
//go:nosplit
func (self class) RemoveTranslation(translation object.Translation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(translation[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_remove_translation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Translation] instance based on the [param locale] passed in.
It will return [code]null[/code] if there is no [Translation] instance that matches the [param locale].
*/
//go:nosplit
func (self class) GetTranslationObject(ctx gd.Lifetime, locale gd.String) object.Translation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_translation_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Translation
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Clears the server from all translations.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an array of all loaded locales of the project.
*/
//go:nosplit
func (self class) GetLoadedLocales(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_loaded_locales, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsPseudolocalizationEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_is_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPseudolocalizationEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_set_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Reparses the pseudolocalization options and reloads the translation.
*/
//go:nosplit
func (self class) ReloadPseudolocalization()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_reload_pseudolocalization, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the pseudolocalized string based on the [param message] passed in.
*/
//go:nosplit
func (self class) Pseudolocalize(ctx gd.Lifetime, message gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_pseudolocalize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTranslationServer() Expert { return self[0].AsTranslationServer() }


//go:nosplit
func (self Simple) AsTranslationServer() Simple { return self[0].AsTranslationServer() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TranslationServer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
