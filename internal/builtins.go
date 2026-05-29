package gd

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/jumponly"
	"graphics.gd/internal/noescape"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

func callBuiltinMethod[T any, S any](self *S, method func(self *S, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer), shape gdextension.Shape, args unsafe.Pointer) T {
	ring.Main.Flush()
	var result T
	method((*S)(noescape.Pointer(unsafe.Pointer(self))), noescape.Pointer(unsafe.Pointer(&result)), shape, args)
	return result
}

// call wraps BuiltinMethodPointer.Call with a ring flush to maintain ordering
// with ring-buffered object method calls.
func call[T gdunsafe.Any, Args any, Result gdunsafe.Returnable](builtin gdunsafe.BuiltinMethod[T, Args, Result], self T, args Args) Result {
	ring.Main.Flush()
	return builtin.Call(self, args)
}

type Int = int64
type Void = struct{}

// builtin methods that are strictly required for graphics.gd to function.
var builtin = gdunsafe.Import[struct {
	typeset

	Array struct {
		size           gdunsafe.BuiltinMethod[gdunsafe.Array, struct{}, Int]           `hash:"3173160232"`
		resize         gdunsafe.BuiltinMethod[gdunsafe.Array, struct{ size Int }, Int] `hash:"848867239"`
		is_read_only   gdunsafe.BuiltinMethod[gdunsafe.Array, struct{}, bool]          `hash:"3918633141"`
		make_read_only gdunsafe.BuiltinMethod[gdunsafe.Array, struct{}, Void]          `hash:"3218959716"`
	}
	Callable struct {
		get_method          gdunsafe.BuiltinMethod[gdunsafe.Callable, struct{}, gdunsafe.StringName]                    `hash:"1825232092"`
		get_bound_arguments gdunsafe.BuiltinMethod[gdunsafe.Callable, struct{}, gdunsafe.Array]                         `hash:"4144163970"`
		get_argument_count  gdunsafe.BuiltinMethod[gdunsafe.Callable, struct{}, Int]                                    `hash:"3173160232"`
		callv               gdunsafe.BuiltinMethod[gdunsafe.Callable, struct{ args gdunsafe.Array }, gdunsafe.Variant]  `hash:"413578926"`
		bindv               gdunsafe.BuiltinMethod[gdunsafe.Callable, struct{ args gdunsafe.Array }, gdunsafe.Callable] `hash:"3564560322"`
		call_deferred       gdunsafe.BuiltinMethod[gdunsafe.Callable, []gdunsafe.Variant, gdunsafe.Variant]             `hash:"3286317445"`
	}
	Dictionary struct {
		keys           gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, gdunsafe.Array]             `hash:"4144163970"`
		has            gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{ key gdunsafe.Variant }, bool] `hash:"3680194679"`
		clear          gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, Void]                       `hash:"3218959716"`
		sort           gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, Void]                       `hash:"3218959716"`
		erase          gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{ key gdunsafe.Variant }, bool] `hash:"1776646889"`
		hash           gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, Int]                        `hash:"3173160232"`
		size           gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, Int]                        `hash:"3173160232"`
		is_read_only   gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, bool]                       `hash:"3918633141"`
		make_read_only gdunsafe.BuiltinMethod[gdunsafe.Dictionary, struct{}, Void]                       `hash:"3218959716"`
	}
	PackedByteArray struct {
		resize    gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[byte], struct{ size Int }, Int]       `hash:"848867239"`
		size      gdunsafe.BuiltinMethod[gdunsafe.PackedArray[byte], struct{}, Int]                        `hash:"3173160232"`
		duplicate gdunsafe.BuiltinMethod[gdunsafe.PackedArray[byte], struct{}, gdunsafe.PackedArray[byte]] `hash:"851781288"`
	}
	PackedColorArray struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[Color.RGBA], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[Color.RGBA], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedFloat32Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[float32], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[float32], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedFloat64Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[float64], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[float64], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedInt32Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[int32], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[int32], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedStringArray struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[gdunsafe.String], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[gdunsafe.String], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedVector2Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[Vector2.XY], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[Vector2.XY], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedVector3Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[Vector3.XYZ], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[Vector3.XYZ], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedVector4Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[Vector4.XYZW], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[Vector4.XYZW], struct{}, Int]                  `hash:"3173160232"`
	}
	PackedInt64Array struct {
		resize gdunsafe.BuiltinMethodMutable[gdunsafe.PackedArray[int64], struct{ size Int }, Int] `hash:"848867239"`
		size   gdunsafe.BuiltinMethod[gdunsafe.PackedArray[int64], struct{}, Int]                  `hash:"3173160232"`
	}
	Signal struct {
		emit    gdunsafe.BuiltinMethod[gdunsafe.Signal, []gdunsafe.Variant, Void] `hash:"3286317445"`
		connect gdunsafe.BuiltinMethod[gdunsafe.Signal, struct {
			callable gdunsafe.Callable
			flags    int64
		}, Int] `hash:"979702392"`
		disconnect      gdunsafe.BuiltinMethod[gdunsafe.Signal, struct{ callable gdunsafe.Callable }, Void] `hash:"3470848906"`
		get_name        gdunsafe.BuiltinMethod[gdunsafe.Signal, struct{}, gdunsafe.StringName]              `hash:"1825232092"`
		get_connections gdunsafe.BuiltinMethod[gdunsafe.Signal, struct{}, gdunsafe.Array]                   `hash:"4144163970"`
		get_object      gdunsafe.BuiltinMethod[gdunsafe.Signal, struct{}, gdunsafe.Object]                  `hash:"4008621732"`
	}
	String struct {
		length     gdunsafe.BuiltinMethod[gdunsafe.String, struct{}, Int]                               `hash:"3173160232"`
		substr     gdunsafe.BuiltinMethod[gdunsafe.String, struct{ begin, end int64 }, gdunsafe.String] `hash:"787537301"`
		casecmp_to gdunsafe.BuiltinMethod[gdunsafe.String, struct{ other gdunsafe.String }, Int]        `hash:"2920860731"`
	}
	StringName struct {
		length     gdunsafe.BuiltinMethod[gdunsafe.StringName, struct{}, Int]                               `hash:"3173160232"`
		substr     gdunsafe.BuiltinMethod[gdunsafe.StringName, struct{ begin, end int64 }, gdunsafe.String] `hash:"787537301"`
		casecmp_to gdunsafe.BuiltinMethod[gdunsafe.StringName, struct{ other gdunsafe.String }, Int]        `hash:"2920860731"`
	}
}]()

func (a Array) Size() int64 {
	return call(builtin.Array.size, gdunsafe.Array(pointers.Get(a)[0]), struct{}{})
}
func (a Array) Resize(size Int) Int {
	return call(builtin.Array.resize, gdunsafe.Array(pointers.Get(a)[0]), struct{ size Int }{size})
}
func (a Array) IsReadOnly() bool {
	return call(builtin.Array.is_read_only, gdunsafe.Array(pointers.Get(a)[0]), struct{}{})
}
func (a Array) MakeReadOnly() {
	call(builtin.Array.make_read_only, gdunsafe.Array(pointers.Get(a)[0]), struct{}{})
}

func (c Callable) GetMethod() StringName {
	result := call(builtin.Callable.get_method, gdunsafe.Callable(pointers.Get(c)), struct{}{})
	return pointers.New[StringName](gdextension.StringName{gdextension.Pointer(result)})
}
func (c Callable) GetBoundArguments() Array {
	result := call(builtin.Callable.get_bound_arguments, gdunsafe.Callable(pointers.Get(c)), struct{}{})
	return pointers.New[Array](gdextension.Array{gdextension.Pointer(result)})
}
func (c Callable) GetArgumentCount() int64 {
	return call(builtin.Callable.get_argument_count, gdunsafe.Callable(pointers.Get(c)), struct{}{})
}
func (c Callable) Call(args ...Variant) Variant {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	var array = NewArray()
	array.Resize(int64(len(args)))
	for i, arg := range args {
		array.SetIndex(int64(i), arg)
	}
	defer array.Free()
	return pointers.New[Variant](gdextension.Variant(builtin.Callable.callv.Call(ptr, struct{ args gdunsafe.Array }{gdunsafe.Array(pointers.Get(array)[0])})))
}
func (c Callable) CallDeferred() Variant {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	return pointers.New[Variant](gdextension.Variant(builtin.Callable.call_deferred.Call(ptr, nil)))
}
func (c Callable) Bind(args ...Variant) Callable {
	var array = NewArray()
	array.Resize(int64(len(args)))
	for i, arg := range args {
		array.SetIndex(int64(i), arg)
	}
	defer array.Free()
	result := call(builtin.Callable.bindv, gdunsafe.Callable(pointers.Get(c)), struct{ args gdunsafe.Array }{gdunsafe.Array(pointers.Get(array)[0])})
	return pointers.New[Callable](gdextension.Callable(result))
}

func (d Dictionary) Keys() Array {
	result := call(builtin.Dictionary.keys, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
	return pointers.New[Array](gdextension.Array{gdextension.Pointer(result)})
}
func (d Dictionary) Has(key Variant) bool {
	return call(builtin.Dictionary.has, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{ key gdunsafe.Variant }{key: pointers.Get(key)})
}
func (d Dictionary) Clear() {
	call(builtin.Dictionary.clear, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
}
func (d Dictionary) Sort() {
	call(builtin.Dictionary.sort, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
}
func (d Dictionary) Erase(key Variant) bool {
	return call(builtin.Dictionary.erase, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{ key gdunsafe.Variant }{key: pointers.Get(key)})
}
func (d Dictionary) Hash() int64 {
	return call(builtin.Dictionary.hash, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
}
func (d Dictionary) Size() int64 {
	return call(builtin.Dictionary.size, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
}
func (d Dictionary) IsReadOnly() bool {
	return call(builtin.Dictionary.is_read_only, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
}
func (d Dictionary) MakeReadOnly() {
	call(builtin.Dictionary.make_read_only, gdunsafe.Dictionary(pointers.Get(d)[0]), struct{}{})
}

func (a PackedByteArray) Resize(size Int) Int {
	self := gdunsafe.PackedArray[byte](pointers.Get(a))
	result := builtin.PackedByteArray.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[byte](self))
	return result
}
func (a PackedByteArray) Size() int64 {
	return call(builtin.PackedByteArray.size, gdunsafe.PackedArray[byte](pointers.Get(a)), struct{}{})
}
func (a PackedByteArray) Duplicate() PackedByteArray {
	result := call(builtin.PackedByteArray.duplicate, gdunsafe.PackedArray[byte](pointers.Get(a)), struct{}{})
	return pointers.New[PackedByteArray](gdextension.PackedArray[byte](result))
}
func (a PackedColorArray) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Color](pointers.Get(a))
	result := builtin.PackedColorArray.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[Color](self))
	return result
}
func (a PackedColorArray) Size() int64 {
	return call(builtin.PackedColorArray.size, gdunsafe.PackedArray[Color](pointers.Get(a)), struct{}{})
}
func (a PackedFloat32Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[float32](pointers.Get(a))
	result := builtin.PackedFloat32Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[float32](self))
	return result
}
func (a PackedFloat32Array) Size() int64 {
	return call(builtin.PackedFloat32Array.size, gdunsafe.PackedArray[float32](pointers.Get(a)), struct{}{})
}
func (a PackedFloat64Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[float64](pointers.Get(a))
	result := builtin.PackedFloat64Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[float64](self))
	return result
}
func (a PackedFloat64Array) Size() int64 {
	return call(builtin.PackedFloat64Array.size, gdunsafe.PackedArray[float64](pointers.Get(a)), struct{}{})
}
func (a PackedInt32Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[int32](pointers.Get(a))
	result := builtin.PackedInt32Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[int32](self))
	return result
}
func (a PackedInt32Array) Size() int64 {
	return call(builtin.PackedInt32Array.size, gdunsafe.PackedArray[int32](pointers.Get(a)), struct{}{})
}
func (a PackedStringArray) Resize(size Int) Int {
	self := gdunsafe.PackedArray[gdunsafe.String](pointers.Get(a))
	result := builtin.PackedStringArray.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[gdextension.String](self))
	return result
}
func (a PackedStringArray) Size() int64 {
	return call(builtin.PackedStringArray.size, gdunsafe.PackedArray[gdunsafe.String](pointers.Get(a)), struct{}{})
}
func (a PackedVector2Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Vector2](pointers.Get(a))
	result := builtin.PackedVector2Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[Vector2](self))
	return result
}
func (a PackedVector2Array) Size() int64 {
	return call(builtin.PackedVector2Array.size, gdunsafe.PackedArray[Vector2](pointers.Get(a)), struct{}{})
}
func (a PackedVector3Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Vector3](pointers.Get(a))
	result := builtin.PackedVector3Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[Vector3](self))
	return result
}
func (a PackedVector3Array) Size() int64 {
	return call(builtin.PackedVector3Array.size, gdunsafe.PackedArray[Vector3](pointers.Get(a)), struct{}{})
}
func (a PackedVector4Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Vector4](pointers.Get(a))
	result := builtin.PackedVector4Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[Vector4](self))
	return result
}
func (a PackedVector4Array) Size() int64 {
	return call(builtin.PackedVector4Array.size, gdunsafe.PackedArray[Vector4](pointers.Get(a)), struct{}{})
}
func (a PackedInt64Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[int64](pointers.Get(a))
	result := builtin.PackedInt64Array.resize.Call(&self, unsafe.Pointer(&self), struct{ size Int }{size})
	pointers.Set(a, gdextension.PackedArray[int64](self))
	return result
}
func (a PackedInt64Array) Size() int64 {
	return call(builtin.PackedInt64Array.size, gdunsafe.PackedArray[int64](pointers.Get(a)), struct{}{})
}

func (s Signal) Emit(args ...Variant) {
	var converted = make([]gdextension.Variant, len(args))
	for i, arg := range args {
		converted[i] = gdextension.Variant(pointers.Get(arg))
	}
	self := gdunsafe.Signal(pointers.Get(s))
	builtin.Signal.emit.Call(self, converted)
}

func (s Signal) Connect(callable Callable, flags int64) int64 {
	return call(builtin.Signal.connect, gdunsafe.Signal(pointers.Get(s)), struct {
		callable gdunsafe.Callable
		flags    int64
	}{gdunsafe.Callable(pointers.Get(callable)), flags})
}
func (s Signal) Disconnect(callable Callable) {
	call(builtin.Signal.disconnect, gdunsafe.Signal(pointers.Get(s)), struct{ callable gdunsafe.Callable }{gdunsafe.Callable(pointers.Get(callable))})
}
func (s Signal) GetName() StringName {
	result := call(builtin.Signal.get_name, gdunsafe.Signal(pointers.Get(s)), struct{}{})
	return pointers.New[StringName](gdextension.StringName{gdextension.Pointer(result)})
}
func (s Signal) GetConnections() Array {
	result := call(builtin.Signal.get_connections, gdunsafe.Signal(pointers.Get(s)), struct{}{})
	return pointers.New[Array](gdextension.Array{gdextension.Pointer(result)})
}
func (s Signal) GetObject() gdreference.Object {
	result := call(builtin.Signal.get_object, gdunsafe.Signal(pointers.Get(s)), struct{}{})
	return gdreference.OwnObject(gdextension.Object(result), Free)
}
func (s String) Length() int64 {
	return call(builtin.String.length, gdunsafe.String(pointers.Get(s)[0]), struct{}{})
}
func (s String) Substr(begin, end int64) String {
	result := call(builtin.String.substr, gdunsafe.String(pointers.Get(s)[0]), struct{ begin, end int64 }{begin, end})
	return pointers.New[String](gdextension.String{gdextension.Pointer(result)})
}
func (s String) CasecmpTo(other String) int64 {
	return call(builtin.String.casecmp_to, gdunsafe.String(pointers.Get(s)[0]), struct{ other gdunsafe.String }{gdunsafe.String(pointers.Get(other)[0])})
}
func (s StringName) Length() int64 {
	return call(builtin.StringName.length, gdunsafe.StringName(pointers.Get(s)[0]), struct{}{})
}
func (s StringName) Substr(begin, end int64) String {
	result := call(builtin.StringName.substr, gdunsafe.StringName(pointers.Get(s)[0]), struct{ begin, end int64 }{begin, end})
	return pointers.New[String](gdextension.String{gdextension.Pointer(result)})
}
func (s StringName) CasecmpTo(other String) int64 {
	return call(builtin.StringName.casecmp_to, gdunsafe.StringName(pointers.Get(s)[0]), struct{ other gdunsafe.String }{gdunsafe.String(pointers.Get(other)[0])})
}

var object_methods struct {
	get_class                    gdextension.MethodForClass `hash:"201670096"`
	is_class                     gdextension.MethodForClass `hash:"3927539163"`
	set                          gdextension.MethodForClass `hash:"3776071444"`
	get                          gdextension.MethodForClass `hash:"2760726917"`
	set_indexed                  gdextension.MethodForClass `hash:"3500910842"`
	get_indexed                  gdextension.MethodForClass `hash:"4006125091"`
	get_property_list            gdextension.MethodForClass `hash:"3995934104"`
	get_method_list              gdextension.MethodForClass `hash:"3995934104"`
	property_can_revert          gdextension.MethodForClass `hash:"2619796661"`
	property_get_revert          gdextension.MethodForClass `hash:"2760726917"`
	notification                 gdextension.MethodForClass `hash:"4023243586"`
	to_string                    gdextension.MethodForClass `hash:"2841200299"`
	get_instance_id              gdextension.MethodForClass `hash:"3905245786"`
	set_script                   gdextension.MethodForClass `hash:"1114965689"`
	get_script                   gdextension.MethodForClass `hash:"1214101251"`
	set_meta                     gdextension.MethodForClass `hash:"3776071444"`
	remove_meta                  gdextension.MethodForClass `hash:"3304788590"`
	get_meta                     gdextension.MethodForClass `hash:"3990617847"`
	has_meta                     gdextension.MethodForClass `hash:"2619796661"`
	get_meta_list                gdextension.MethodForClass `hash:"3995934104"`
	add_user_signal              gdextension.MethodForClass `hash:"85656714"`
	has_user_signal              gdextension.MethodForClass `hash:"2619796661"`
	remove_user_signal           gdextension.MethodForClass `hash:"3304788590"`
	emit_signal                  gdextension.MethodForClass `hash:"4047867050"`
	call                         gdextension.MethodForClass `hash:"3400424181"`
	call_deferred                gdextension.MethodForClass `hash:"3400424181"`
	set_deferred                 gdextension.MethodForClass `hash:"3776071444"`
	callv                        gdextension.MethodForClass `hash:"1260104456"`
	has_method                   gdextension.MethodForClass `hash:"2619796661"`
	get_method_argument_count    gdextension.MethodForClass `hash:"2458036349"`
	has_signal                   gdextension.MethodForClass `hash:"2619796661"`
	get_signal_list              gdextension.MethodForClass `hash:"3995934104"`
	get_signal_connection_list   gdextension.MethodForClass `hash:"3147814860"`
	get_incoming_connections     gdextension.MethodForClass `hash:"3995934104"`
	connect                      gdextension.MethodForClass `hash:"1518946055"`
	disconnect                   gdextension.MethodForClass `hash:"1874754934"`
	is_connected                 gdextension.MethodForClass `hash:"768136979"`
	has_connections              gdextension.MethodForClass `hash:"2619796661"`
	set_block_signals            gdextension.MethodForClass `hash:"2586408642"`
	is_blocking_signals          gdextension.MethodForClass `hash:"36873697"`
	notify_property_list_changed gdextension.MethodForClass `hash:"3218959716"`
	set_message_translation      gdextension.MethodForClass `hash:"2586408642"`
	can_translate_messages       gdextension.MethodForClass `hash:"36873697"`
	tr                           gdextension.MethodForClass `hash:"1195764410"`
	tr_n                         gdextension.MethodForClass `hash:"162698058"`
	get_translation_domain       gdextension.MethodForClass `hash:"2002593661"`
	set_translation_domain       gdextension.MethodForClass `hash:"3304788590"`
	is_queued_for_deletion       gdextension.MethodForClass `hash:"36873697"`
	cancel_free                  gdextension.MethodForClass `hash:"3218959716"`
}

var refcounted_methods struct {
	init_ref            gdextension.MethodForClass `hash:"2240911060"`
	reference           gdextension.MethodForClass `hash:"2240911060"`
	unreference         gdextension.MethodForClass `hash:"2240911060"`
	get_reference_count gdextension.MethodForClass `hash:"3905245786"`
}

func ObjectGet(o gdreference.Object, name gdunsafe.StringName) gdunsafe.Variant {
	return noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get, gdextension.SizeVariant|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdunsafe.StringName
	}{
		name,
	}))
}
func ObjectSet(o gdreference.Object, name gdunsafe.StringName, value gdunsafe.Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set, 0|gdextension.SizeStringName<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		Name  gdunsafe.StringName
		Value gdunsafe.Variant
	}{
		name, gdextension.Variant(pointers.Get(value)),
	}))
}

func ObjectGetMeta(o gdreference.Object, name gdunsafe.StringName) gdunsafe.Variant {
	return noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get_meta, gdextension.SizeVariant|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdunsafe.StringName
	}{
		name,
	}))
}
func ObjectSetMeta(o gdreference.Object, name gdunsafe.StringName, value gdunsafe.Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set_meta, 0|gdextension.SizeStringName<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		Name  gdunsafe.StringName
		Value gdunsafe.Variant
	}{
		name, value,
	}))
}

func ObjectHasMethod(o gdreference.Object, name gdunsafe.StringName) bool {
	return noescape.Call[bool](gdreference.GetObject(o), object_methods.has_method, gdextension.SizeBool|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdunsafe.StringName
	}{
		name,
	}))
}
func ObjectCall(o gdreference.Object, method gdunsafe.StringName, args ...gdunsafe.Variant) (gdunsafe.Variant, error) {
	ring.Main.Flush()
	self := gdreference.GetObject(o)
	if gdunsafe.Script(self).HasMethod(method) {
		var converted []gdunsafe.Variant
		for _, arg := range args {
			converted = append(converted, arg)
		}
		result, err := gdunsafe.Script(self).Call(method,
			converted...,
		)
		if err != (gdunsafe.Error{}) {
			return result, err
		}
		return result, nil
	}
	return NewVariant(o).Call(method, args...) // FIXME is this ok?
}

func ObjectCanTranslateMessages(o gdreference.Object) bool {
	return jumponly.Call[bool](gdreference.GetObject(o), object_methods.can_translate_messages, gdextension.SizeBool, nil)
}
func ObjectGetScript(o gdreference.Object) gdunsafe.Variant {
	return noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get_script, gdextension.SizeVariant, nil)
}
func ObjectNotifyPropertyListChanged(o gdreference.Object) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.notify_property_list_changed, 0, nil)
}
func ObjectSetBlockSignals(o gdreference.Object, blocking bool) {
	jumponly.Call[struct{}](gdreference.GetObject(o), object_methods.set_block_signals, 0|gdextension.SizeBool<<4, unsafe.Pointer(&struct {
		Blocking bool
	}{
		blocking,
	}))
}
func ObjectSetScript(o gdreference.Object, script gdunsafe.Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set_script, 0|gdextension.SizeVariant<<4, unsafe.Pointer(&struct {
		Script gdunsafe.Variant
	}{
		script,
	}))
}
func ObjectToString(o gdreference.Object) gdunsafe.String {
	return noescape.Call[gdunsafe.String](gdreference.GetObject(o), object_methods.to_string, gdextension.SizeString, nil)
}
func ObjectTr(o gdreference.Object, message gdunsafe.StringName, context gdunsafe.StringName) gdunsafe.String {
	return noescape.Call[gdunsafe.String](gdreference.GetObject(o), object_methods.tr, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8, unsafe.Pointer(&struct {
		Message gdunsafe.StringName
		Context gdunsafe.StringName
	}{
		message, context,
	}))
}
func ObjectTrN(o gdreference.Object, message gdunsafe.StringName, plural gdunsafe.StringName, n int64, context gdunsafe.StringName) gdunsafe.String {
	return noescape.Call[gdunsafe.String](gdreference.GetObject(o), object_methods.tr_n, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8|gdextension.SizeInt<<12|gdextension.SizeStringName<<16, unsafe.Pointer(&struct {
		Message gdunsafe.StringName
		Plural  gdunsafe.StringName
		N       int64
		Context gdunsafe.StringName
	}{
		message, plural, n, context,
	}))
}
func ObjectSetMessageTranslation(o gdreference.Object, enable bool) {
	jumponly.Call[struct{}](gdreference.GetObject(o), object_methods.set_message_translation, 0|gdextension.SizeBool<<4, unsafe.Pointer(&struct {
		Enable bool
	}{
		enable,
	}))
}
func ObjectIsBlockingSignals(o gdreference.Object) bool {
	return jumponly.Call[bool](gdreference.GetObject(o), object_methods.is_blocking_signals, gdextension.SizeBool, nil)
}
func ObjectGetClass(o gdreference.Object) gdunsafe.String {
	return pointers.New[String](noescape.Call[gdextension.String](gdreference.GetObject(o), object_methods.get_class, gdextension.SizeString, nil))
}
func ObjectConnect(o gdreference.Object, signal StringName, callable Callable, flags int64) int64 {
	return noescape.Call[int64](gdreference.GetObject(o), object_methods.connect, gdextension.SizeInt|gdextension.SizeStringName<<4|gdextension.SizeCallable<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		Signal   gdextension.StringName
		Callable gdextension.Callable
		Flags    int64
	}{
		pointers.Get(signal), gdextension.Callable(pointers.Get(callable)), flags,
	}))
}
func ObjectIsConnected(o gdreference.Object, signal StringName, callable Callable) bool {
	return noescape.Call[bool](gdreference.GetObject(o), object_methods.is_connected, gdextension.SizeBool|gdextension.SizeStringName<<4|gdextension.SizeCallable<<8, unsafe.Pointer(&struct {
		Signal   gdextension.StringName
		Callable gdextension.Callable
	}{
		pointers.Get(signal), gdextension.Callable(pointers.Get(callable)),
	}))
}
func ObjectDisconnect(o gdreference.Object, signal StringName, callable Callable) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.disconnect, 0|gdextension.SizeStringName<<4|gdextension.SizeCallable<<8, unsafe.Pointer(&struct {
		Signal   gdextension.StringName
		Callable gdextension.Callable
	}{
		pointers.Get(signal), gdextension.Callable(pointers.Get(callable)),
	}))
}
func ObjectIsQueuedForDeletion(o gdreference.Object) bool {
	return jumponly.Call[bool](gdreference.GetObject(o), object_methods.is_queued_for_deletion, gdextension.SizeBool, nil)
}
func ObjectNotification(o gdreference.Object, what Int, reversed bool) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.notification, 0|gdextension.SizeInt<<4|gdextension.SizeBool<<8, unsafe.Pointer(&struct {
		What     int64
		Reversed bool
	}{
		int64(what), reversed,
	}))
}
func ObjectGetPropertyList(o gdreference.Object) Array {
	return pointers.New[Array](noescape.Call[gdextension.Array](gdreference.GetObject(o), object_methods.get_property_list, gdextension.SizeArray, nil))
}

func ObjectSetIndex(o gdreference.Object, i int, v Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set_indexed, 0|gdextension.SizeInt<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		Index   int64
		Element gdextension.Variant
	}{
		int64(i), gdextension.Variant(pointers.Get(v)),
	}))
}

func ObjectGetIndex(o gdreference.Object, i int) Variant {
	return pointers.New[Variant](noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get_indexed, gdextension.SizeVariant|gdextension.SizeInt<<4, unsafe.Pointer(&struct {
		Index int64
	}{
		int64(i),
	})))
}

func (rc RefCounted) Reference() {
	noescape.Call[struct{}](ObjectChecked(rc.AsObject()), refcounted_methods.reference, 0, nil)
}
func (rc RefCounted) Unreference() bool {
	raw := ObjectChecked(rc.AsObject())
	if raw == 0 {
		return false
	}
	return noescape.Call[bool](raw, refcounted_methods.unreference, gdextension.SizeBool, nil)
}
func (rc RefCounted) InitRef() bool {
	return noescape.Call[bool](ObjectChecked(rc.AsObject()), refcounted_methods.init_ref, gdextension.SizeBool, nil)
}

func (rc RefCounted) GetReferenceCount() int {
	return int(noescape.Call[int64](ObjectChecked(rc.AsObject()), refcounted_methods.get_reference_count, gdextension.SizeInt, nil))
}
