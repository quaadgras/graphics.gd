package TranslationServer

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The server that manages all language translations. Translations can be added to or removed from it.
*/
var self objects.TranslationServer
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.TranslationServer)
	self = *(*objects.TranslationServer)(unsafe.Pointer(&obj))
}

/*
Sets the locale of the project. The [param locale] string will be standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
If translations have been loaded beforehand for the new locale, they will be applied.
*/
func SetLocale(locale string) {
	once.Do(singleton)
	class(self).SetLocale(gd.NewString(locale))
}

/*
Returns the current locale of the project.
See also [method OS.get_locale] and [method OS.get_locale_language] to query the locale of the user system.
*/
func GetLocale() string {
	once.Do(singleton)
	return string(class(self).GetLocale().String())
}

/*
Returns the current locale of the editor.
[b]Note:[/b] When called from an exported project returns the same value as [method get_locale].
*/
func GetToolLocale() string {
	once.Do(singleton)
	return string(class(self).GetToolLocale().String())
}

/*
Compares two locales and returns a similarity score between [code]0[/code] (no match) and [code]10[/code] (full match).
*/
func CompareLocales(locale_a string, locale_b string) int {
	once.Do(singleton)
	return int(int(class(self).CompareLocales(gd.NewString(locale_a), gd.NewString(locale_b))))
}

/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
*/
func StandardizeLocale(locale string) string {
	once.Do(singleton)
	return string(class(self).StandardizeLocale(gd.NewString(locale)).String())
}

/*
Returns array of known language codes.
*/
func GetAllLanguages() []string {
	once.Do(singleton)
	return []string(class(self).GetAllLanguages().Strings())
}

/*
Returns a readable language name for the [param language] code.
*/
func GetLanguageName(language string) string {
	once.Do(singleton)
	return string(class(self).GetLanguageName(gd.NewString(language)).String())
}

/*
Returns an array of known script codes.
*/
func GetAllScripts() []string {
	once.Do(singleton)
	return []string(class(self).GetAllScripts().Strings())
}

/*
Returns a readable script name for the [param script] code.
*/
func GetScriptName(script string) string {
	once.Do(singleton)
	return string(class(self).GetScriptName(gd.NewString(script)).String())
}

/*
Returns an array of known country codes.
*/
func GetAllCountries() []string {
	once.Do(singleton)
	return []string(class(self).GetAllCountries().Strings())
}

/*
Returns a readable country name for the [param country] code.
*/
func GetCountryName(country string) string {
	once.Do(singleton)
	return string(class(self).GetCountryName(gd.NewString(country)).String())
}

/*
Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
*/
func GetLocaleName(locale string) string {
	once.Do(singleton)
	return string(class(self).GetLocaleName(gd.NewString(locale)).String())
}

/*
Returns the current locale's translation for the given message (key) and context.
*/
func Translate(message string) string {
	once.Do(singleton)
	return string(class(self).Translate(gd.NewStringName(message), gd.NewStringName("")).String())
}

/*
Returns the current locale's translation for the given message (key), plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func TranslatePlural(message string, plural_message string, n int) string {
	once.Do(singleton)
	return string(class(self).TranslatePlural(gd.NewStringName(message), gd.NewStringName(plural_message), gd.Int(n), gd.NewStringName("")).String())
}

/*
Adds a [Translation] resource.
*/
func AddTranslation(translation objects.Translation) {
	once.Do(singleton)
	class(self).AddTranslation(translation)
}

/*
Removes the given translation from the server.
*/
func RemoveTranslation(translation objects.Translation) {
	once.Do(singleton)
	class(self).RemoveTranslation(translation)
}

/*
Returns the [Translation] instance based on the [param locale] passed in.
It will return [code]null[/code] if there is no [Translation] instance that matches the [param locale].
*/
func GetTranslationObject(locale string) objects.Translation {
	once.Do(singleton)
	return objects.Translation(class(self).GetTranslationObject(gd.NewString(locale)))
}

/*
Clears the server from all translations.
*/
func Clear() {
	once.Do(singleton)
	class(self).Clear()
}

/*
Returns an array of all loaded locales of the project.
*/
func GetLoadedLocales() []string {
	once.Do(singleton)
	return []string(class(self).GetLoadedLocales().Strings())
}

/*
Reparses the pseudolocalization options and reloads the translation.
*/
func ReloadPseudolocalization() {
	once.Do(singleton)
	class(self).ReloadPseudolocalization()
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
*/
func Pseudolocalize(message string) string {
	once.Do(singleton)
	return string(class(self).Pseudolocalize(gd.NewStringName(message)).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.TranslationServer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func PseudolocalizationEnabled() bool {
	return bool(class(self).IsPseudolocalizationEnabled())
}

func SetPseudolocalizationEnabled(value bool) {
	class(self).SetPseudolocalizationEnabled(value)
}

/*
Sets the locale of the project. The [param locale] string will be standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
If translations have been loaded beforehand for the new locale, they will be applied.
*/
//go:nosplit
func (self class) SetLocale(locale gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(locale))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_set_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current locale of the project.
See also [method OS.get_locale] and [method OS.get_locale_language] to query the locale of the user system.
*/
//go:nosplit
func (self class) GetLocale() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current locale of the editor.
[b]Note:[/b] When called from an exported project returns the same value as [method get_locale].
*/
//go:nosplit
func (self class) GetToolLocale() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_tool_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Compares two locales and returns a similarity score between [code]0[/code] (no match) and [code]10[/code] (full match).
*/
//go:nosplit
func (self class) CompareLocales(locale_a gd.String, locale_b gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(locale_a))
	callframe.Arg(frame, pointers.Get(locale_b))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_compare_locales, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
*/
//go:nosplit
func (self class) StandardizeLocale(locale gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(locale))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_standardize_locale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns array of known language codes.
*/
//go:nosplit
func (self class) GetAllLanguages() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_all_languages, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a readable language name for the [param language] code.
*/
//go:nosplit
func (self class) GetLanguageName(language gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_language_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of known script codes.
*/
//go:nosplit
func (self class) GetAllScripts() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_all_scripts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a readable script name for the [param script] code.
*/
//go:nosplit
func (self class) GetScriptName(script gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_script_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an array of known country codes.
*/
//go:nosplit
func (self class) GetAllCountries() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_all_countries, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a readable country name for the [param country] code.
*/
//go:nosplit
func (self class) GetCountryName(country gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(country))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_country_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
*/
//go:nosplit
func (self class) GetLocaleName(locale gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(locale))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_locale_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current locale's translation for the given message (key) and context.
*/
//go:nosplit
func (self class) Translate(message gd.StringName, context gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current locale's translation for the given message (key), plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
//go:nosplit
func (self class) TranslatePlural(message gd.StringName, plural_message gd.StringName, n gd.Int, context gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	callframe.Arg(frame, pointers.Get(plural_message))
	callframe.Arg(frame, n)
	callframe.Arg(frame, pointers.Get(context))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_translate_plural, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a [Translation] resource.
*/
//go:nosplit
func (self class) AddTranslation(translation objects.Translation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(translation[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_add_translation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the given translation from the server.
*/
//go:nosplit
func (self class) RemoveTranslation(translation objects.Translation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(translation[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_remove_translation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Translation] instance based on the [param locale] passed in.
It will return [code]null[/code] if there is no [Translation] instance that matches the [param locale].
*/
//go:nosplit
func (self class) GetTranslationObject(locale gd.String) objects.Translation {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(locale))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_translation_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Translation{classdb.Translation(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Clears the server from all translations.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all loaded locales of the project.
*/
//go:nosplit
func (self class) GetLoadedLocales() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_loaded_locales, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsPseudolocalizationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_is_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_set_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Reparses the pseudolocalization options and reloads the translation.
*/
//go:nosplit
func (self class) ReloadPseudolocalization() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_reload_pseudolocalization, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
*/
//go:nosplit
func (self class) Pseudolocalize(message gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(message))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_pseudolocalize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("TranslationServer", func(ptr gd.Object) any { return [1]classdb.TranslationServer{classdb.TranslationServer(ptr)} })
}
