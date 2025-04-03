// Package TranslationDomain provides methods for working with TranslationDomain object instances.
package TranslationDomain

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[TranslationDomain] is a self-contained collection of [Translation] resources. Translations can be added to or removed from it.
If you're working with the main translation domain, it is more convenient to use the wrap methods on [TranslationServer].
*/
type Instance [1]gdclass.TranslationDomain
type Expanded [1]gdclass.TranslationDomain

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTranslationDomain() Instance
}

/*
Returns the [Translation] instance that best matches [param locale]. Returns [code]null[/code] if there are no matches.
*/
func (self Instance) GetTranslationObject(locale string) [1]gdclass.Translation { //gd:TranslationDomain.get_translation_object
	return [1]gdclass.Translation(Advanced(self).GetTranslationObject(String.New(locale)))
}

/*
Adds a translation.
*/
func (self Instance) AddTranslation(translation [1]gdclass.Translation) { //gd:TranslationDomain.add_translation
	Advanced(self).AddTranslation(translation)
}

/*
Removes the given translation.
*/
func (self Instance) RemoveTranslation(translation [1]gdclass.Translation) { //gd:TranslationDomain.remove_translation
	Advanced(self).RemoveTranslation(translation)
}

/*
Removes all translations.
*/
func (self Instance) Clear() { //gd:TranslationDomain.clear
	Advanced(self).Clear()
}

/*
Returns the current locale's translation for the given message and context.
*/
func (self Instance) Translate(message string) string { //gd:TranslationDomain.translate
	return string(Advanced(self).Translate(String.Name(String.New(message)), String.Name(String.New(""))).String())
}

/*
Returns the current locale's translation for the given message and context.
*/
func (self Expanded) Translate(message string, context string) string { //gd:TranslationDomain.translate
	return string(Advanced(self).Translate(String.Name(String.New(message)), String.Name(String.New(context))).String())
}

/*
Returns the current locale's translation for the given message, plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func (self Instance) TranslatePlural(message string, message_plural string, n int) string { //gd:TranslationDomain.translate_plural
	return string(Advanced(self).TranslatePlural(String.Name(String.New(message)), String.Name(String.New(message_plural)), int64(n), String.Name(String.New(""))).String())
}

/*
Returns the current locale's translation for the given message, plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
func (self Expanded) TranslatePlural(message string, message_plural string, n int, context string) string { //gd:TranslationDomain.translate_plural
	return string(Advanced(self).TranslatePlural(String.Name(String.New(message)), String.Name(String.New(message_plural)), int64(n), String.Name(String.New(context))).String())
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
*/
func (self Instance) Pseudolocalize(message string) string { //gd:TranslationDomain.pseudolocalize
	return string(Advanced(self).Pseudolocalize(String.Name(String.New(message))).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TranslationDomain

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TranslationDomain"))
	casted := Instance{*(*gdclass.TranslationDomain)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) PseudolocalizationEnabled() bool {
	return bool(class(self).IsPseudolocalizationEnabled())
}

func (self Instance) SetPseudolocalizationEnabled(value bool) {
	class(self).SetPseudolocalizationEnabled(value)
}

func (self Instance) PseudolocalizationAccentsEnabled() bool {
	return bool(class(self).IsPseudolocalizationAccentsEnabled())
}

func (self Instance) SetPseudolocalizationAccentsEnabled(value bool) {
	class(self).SetPseudolocalizationAccentsEnabled(value)
}

func (self Instance) PseudolocalizationDoubleVowelsEnabled() bool {
	return bool(class(self).IsPseudolocalizationDoubleVowelsEnabled())
}

func (self Instance) SetPseudolocalizationDoubleVowelsEnabled(value bool) {
	class(self).SetPseudolocalizationDoubleVowelsEnabled(value)
}

func (self Instance) PseudolocalizationFakeBidiEnabled() bool {
	return bool(class(self).IsPseudolocalizationFakeBidiEnabled())
}

func (self Instance) SetPseudolocalizationFakeBidiEnabled(value bool) {
	class(self).SetPseudolocalizationFakeBidiEnabled(value)
}

func (self Instance) PseudolocalizationOverrideEnabled() bool {
	return bool(class(self).IsPseudolocalizationOverrideEnabled())
}

func (self Instance) SetPseudolocalizationOverrideEnabled(value bool) {
	class(self).SetPseudolocalizationOverrideEnabled(value)
}

func (self Instance) PseudolocalizationSkipPlaceholdersEnabled() bool {
	return bool(class(self).IsPseudolocalizationSkipPlaceholdersEnabled())
}

func (self Instance) SetPseudolocalizationSkipPlaceholdersEnabled(value bool) {
	class(self).SetPseudolocalizationSkipPlaceholdersEnabled(value)
}

func (self Instance) PseudolocalizationExpansionRatio() Float.X {
	return Float.X(Float.X(class(self).GetPseudolocalizationExpansionRatio()))
}

func (self Instance) SetPseudolocalizationExpansionRatio(value Float.X) {
	class(self).SetPseudolocalizationExpansionRatio(float64(value))
}

func (self Instance) PseudolocalizationPrefix() string {
	return string(class(self).GetPseudolocalizationPrefix().String())
}

func (self Instance) SetPseudolocalizationPrefix(value string) {
	class(self).SetPseudolocalizationPrefix(String.New(value))
}

func (self Instance) PseudolocalizationSuffix() string {
	return string(class(self).GetPseudolocalizationSuffix().String())
}

func (self Instance) SetPseudolocalizationSuffix(value string) {
	class(self).SetPseudolocalizationSuffix(String.New(value))
}

/*
Returns the [Translation] instance that best matches [param locale]. Returns [code]null[/code] if there are no matches.
*/
//go:nosplit
func (self class) GetTranslationObject(locale String.Readable) [1]gdclass.Translation { //gd:TranslationDomain.get_translation_object
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_get_translation_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Translation{gd.PointerWithOwnershipTransferredToGo[gdclass.Translation](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds a translation.
*/
//go:nosplit
func (self class) AddTranslation(translation [1]gdclass.Translation) { //gd:TranslationDomain.add_translation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(translation[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_add_translation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the given translation.
*/
//go:nosplit
func (self class) RemoveTranslation(translation [1]gdclass.Translation) { //gd:TranslationDomain.remove_translation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(translation[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_remove_translation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all translations.
*/
//go:nosplit
func (self class) Clear() { //gd:TranslationDomain.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current locale's translation for the given message and context.
*/
//go:nosplit
func (self class) Translate(message String.Name, context String.Name) String.Name { //gd:TranslationDomain.translate
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(context)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the current locale's translation for the given message, plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
*/
//go:nosplit
func (self class) TranslatePlural(message String.Name, message_plural String.Name, n int64, context String.Name) String.Name { //gd:TranslationDomain.translate_plural
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message_plural)))
	callframe.Arg(frame, n)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(context)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_translate_plural, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsPseudolocalizationEnabled() bool { //gd:TranslationDomain.is_pseudolocalization_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_is_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationEnabled(enabled bool) { //gd:TranslationDomain.set_pseudolocalization_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPseudolocalizationAccentsEnabled() bool { //gd:TranslationDomain.is_pseudolocalization_accents_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_is_pseudolocalization_accents_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationAccentsEnabled(enabled bool) { //gd:TranslationDomain.set_pseudolocalization_accents_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_accents_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPseudolocalizationDoubleVowelsEnabled() bool { //gd:TranslationDomain.is_pseudolocalization_double_vowels_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_is_pseudolocalization_double_vowels_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationDoubleVowelsEnabled(enabled bool) { //gd:TranslationDomain.set_pseudolocalization_double_vowels_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_double_vowels_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPseudolocalizationFakeBidiEnabled() bool { //gd:TranslationDomain.is_pseudolocalization_fake_bidi_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_is_pseudolocalization_fake_bidi_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationFakeBidiEnabled(enabled bool) { //gd:TranslationDomain.set_pseudolocalization_fake_bidi_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_fake_bidi_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPseudolocalizationOverrideEnabled() bool { //gd:TranslationDomain.is_pseudolocalization_override_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_is_pseudolocalization_override_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationOverrideEnabled(enabled bool) { //gd:TranslationDomain.set_pseudolocalization_override_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_override_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPseudolocalizationSkipPlaceholdersEnabled() bool { //gd:TranslationDomain.is_pseudolocalization_skip_placeholders_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_is_pseudolocalization_skip_placeholders_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationSkipPlaceholdersEnabled(enabled bool) { //gd:TranslationDomain.set_pseudolocalization_skip_placeholders_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_skip_placeholders_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPseudolocalizationExpansionRatio() float64 { //gd:TranslationDomain.get_pseudolocalization_expansion_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_get_pseudolocalization_expansion_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationExpansionRatio(ratio float64) { //gd:TranslationDomain.set_pseudolocalization_expansion_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_expansion_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPseudolocalizationPrefix() String.Readable { //gd:TranslationDomain.get_pseudolocalization_prefix
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_get_pseudolocalization_prefix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationPrefix(prefix String.Readable) { //gd:TranslationDomain.set_pseudolocalization_prefix
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(prefix)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_prefix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPseudolocalizationSuffix() String.Readable { //gd:TranslationDomain.get_pseudolocalization_suffix
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_get_pseudolocalization_suffix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationSuffix(suffix String.Readable) { //gd:TranslationDomain.set_pseudolocalization_suffix
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(suffix)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_set_pseudolocalization_suffix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
*/
//go:nosplit
func (self class) Pseudolocalize(message String.Name) String.Name { //gd:TranslationDomain.pseudolocalize
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationDomain.Bind_pseudolocalize, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) AsTranslationDomain() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTranslationDomain() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("TranslationDomain", func(ptr gd.Object) any {
		return [1]gdclass.TranslationDomain{*(*gdclass.TranslationDomain)(unsafe.Pointer(&ptr))}
	})
}
