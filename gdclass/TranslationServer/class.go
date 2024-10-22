package TranslationServer

import "unsafe"
import "sync"
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
The server that manages all language translations. Translations can be added to or removed from it.

*/
var self gdclass.TranslationServer
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.TranslationServer)
	self = *(*gdclass.TranslationServer)(unsafe.Pointer(&obj))
}

/*
Sets the locale of the project. The [param locale] string will be standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
If translations have been loaded beforehand for the new locale, they will be applied.
*/
func SetLocale(locale string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SetLocale(gc.String(locale))
}

/*
Returns the current locale of the project.
See also [method OS.get_locale] and [method OS.get_locale_language] to query the locale of the user system.
*/
func GetLocale() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetLocale(gc).String())
}

/*
Returns the current locale of the editor.
[b]Note:[/b] When called from an exported project returns the same value as [method get_locale].
*/
func GetToolLocale() string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetToolLocale(gc).String())
}

/*
Compares two locales and returns a similarity score between [code]0[/code] (no match) and [code]10[/code] (full match).
*/
func CompareLocales(locale_a string, locale_b string) int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).CompareLocales(gc.String(locale_a), gc.String(locale_b))))
}

/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
*/
func StandardizeLocale(locale string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).StandardizeLocale(gc, gc.String(locale)).String())
}

/*
Returns array of known language codes.
*/
func GetAllLanguages() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetAllLanguages(gc).Strings(gc))
}

/*
Returns a readable language name for the [param language] code.
*/
func GetLanguageName(language string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetLanguageName(gc, gc.String(language)).String())
}

/*
Returns an array of known script codes.
*/
func GetAllScripts() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetAllScripts(gc).Strings(gc))
}

/*
Returns a readable script name for the [param script] code.
*/
func GetScriptName(script string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetScriptName(gc, gc.String(script)).String())
}

/*
Returns an array of known country codes.
*/
func GetAllCountries() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetAllCountries(gc).Strings(gc))
}

/*
Returns a readable country name for the [param country] code.
*/
func GetCountryName(country string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetCountryName(gc, gc.String(country)).String())
}

/*
Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
*/
func GetLocaleName(locale string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetLocaleName(gc, gc.String(locale)).String())
}

/*
Returns the current locale's translation for the given message (key) and context.
*/
func Translate(message string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).Translate(gc, gc.StringName(message), gc.StringName("")).String())
}

/*
Returns the current locale's translation for the given message (key), plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func TranslatePlural(message string, plural_message string, n int) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).TranslatePlural(gc, gc.StringName(message), gc.StringName(plural_message), gd.Int(n), gc.StringName("")).String())
}

/*
Adds a [Translation] resource.
*/
func AddTranslation(translation gdclass.Translation) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).AddTranslation(translation)
}

/*
Removes the given translation from the server.
*/
func RemoveTranslation(translation gdclass.Translation) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).RemoveTranslation(translation)
}

/*
Returns the [Translation] instance based on the [param locale] passed in.
It will return [code]null[/code] if there is no [Translation] instance that matches the [param locale].
*/
func GetTranslationObject(locale string) gdclass.Translation {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.Translation(class(self).GetTranslationObject(gc, gc.String(locale)))
}

/*
Clears the server from all translations.
*/
func Clear() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).Clear()
}

/*
Returns an array of all loaded locales of the project.
*/
func GetLoadedLocales() []string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return []string(class(self).GetLoadedLocales(gc).Strings(gc))
}

/*
Reparses the pseudolocalization options and reloads the translation.
*/
func ReloadPseudolocalization() {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).ReloadPseudolocalization()
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
*/
func Pseudolocalize(message string) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).Pseudolocalize(gc, gc.StringName(message)).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.TranslationServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func PseudolocalizationEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPseudolocalizationEnabled())
}

func SetPseudolocalizationEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPseudolocalizationEnabled(value)
}

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
func (self class) AddTranslation(translation gdclass.Translation)  {
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
func (self class) RemoveTranslation(translation gdclass.Translation)  {
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
func (self class) GetTranslationObject(ctx gd.Lifetime, locale gd.String) gdclass.Translation {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(locale))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TranslationServer.Bind_get_translation_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Translation
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
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("TranslationServer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
