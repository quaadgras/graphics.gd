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
)

func callBuiltinMethod[T any, S any](self *S, method func(self *S, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer), shape gdextension.Shape, args unsafe.Pointer) T {
	ring.Main.Flush()
	var result T
	method((*S)(noescape.Pointer(unsafe.Pointer(self))), noescape.Pointer(unsafe.Pointer(&result)), shape, args)
	return result
}

// builtin methods that are strictly required for graphics.gd to function.
var builtin struct {
	typeset

	Array struct {
		size           func(self *gdunsafe.Array, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		resize         func(self *gdunsafe.Array, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		is_read_only   func(self *gdunsafe.Array, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3918633141"`
		make_read_only func(self *gdunsafe.Array, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3218959716"`
	}
	Callable struct {
		get_method          func(self *gdunsafe.Callable, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"1825232092"`
		get_bound_arguments func(self *gdunsafe.Callable, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"4144163970"`
		get_argument_count  func(self *gdunsafe.Callable, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		callv               func(self *gdunsafe.Callable, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"413578926"`
		bindv               func(self *gdunsafe.Callable, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3564560322"`
		call_deferred       func(self *gdunsafe.Callable, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3286317445"`
	}
	Dictionary struct {
		keys           func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"4144163970"`
		has            func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3680194679"`
		clear          func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3218959716"`
		sort           func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3218959716"`
		erase          func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"1776646889"`
		hash           func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		size           func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		is_read_only   func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3918633141"`
		make_read_only func(self *gdunsafe.Dictionary, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3218959716"`
	}
	PackedByteArray struct {
		resize    func(self *gdunsafe.PackedArray[byte], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size      func(self *gdunsafe.PackedArray[byte], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		duplicate func(self *gdunsafe.PackedArray[byte], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"851781288"`
	}
	PackedColorArray struct {
		resize func(self *gdunsafe.PackedArray[Color], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[Color], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedFloat32Array struct {
		resize func(self *gdunsafe.PackedArray[float32], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[float32], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedFloat64Array struct {
		resize func(self *gdunsafe.PackedArray[float64], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[float64], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedInt32Array struct {
		resize func(self *gdunsafe.PackedArray[int32], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[int32], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedStringArray struct {
		resize func(self *gdunsafe.PackedArray[gdunsafe.String], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[gdunsafe.String], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedVector2Array struct {
		resize func(self *gdunsafe.PackedArray[Vector2], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[Vector2], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedVector3Array struct {
		resize func(self *gdunsafe.PackedArray[Vector3], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[Vector3], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedVector4Array struct {
		resize func(self *gdunsafe.PackedArray[Vector4], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[Vector4], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	PackedInt64Array struct {
		resize func(self *gdunsafe.PackedArray[int64], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"848867239"`
		size   func(self *gdunsafe.PackedArray[int64], ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
	}
	Signal struct {
		emit            func(self *gdunsafe.Signal, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3286317445"`
		connect         func(self *gdunsafe.Signal, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"979702392"`
		disconnect      func(self *gdunsafe.Signal, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3470848906"`
		get_name        func(self *gdunsafe.Signal, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"1825232092"`
		get_connections func(self *gdunsafe.Signal, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"4144163970"`
		get_object      func(self *gdunsafe.Signal, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"4008621732"`
	}
	String struct {
		length     func(self *gdunsafe.String, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		substr     func(self *gdunsafe.String, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"787537301"`
		casecmp_to func(self *gdunsafe.String, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"2920860731"`
	}
	StringName struct {
		length     func(self *gdunsafe.StringName, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"3173160232"`
		substr     func(self *gdunsafe.StringName, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"787537301"`
		casecmp_to func(self *gdunsafe.StringName, ret unsafe.Pointer, shape gdunsafe.Shape, args unsafe.Pointer) `hash:"2920860731"`
	}
}

func (a Array) Size() int64 {
	self := gdunsafe.Array(pointers.Get(a)[0])
	return callBuiltinMethod[int64](&self, builtin.Array.size, gdextension.SizeInt|gdextension.SizeArray<<4, nil)
}

func (a Array) Resize(size Int) Int {
	self := gdunsafe.Array(pointers.Get(a)[0])
	return Int(callBuiltinMethod[int64](&self, builtin.Array.resize, 0|gdextension.SizeArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a Array) IsReadOnly() bool {
	self := gdunsafe.Array(pointers.Get(a)[0])
	return callBuiltinMethod[bool](&self, builtin.Array.is_read_only, gdextension.SizeBool|gdextension.SizeArray<<4, nil)
}
func (a Array) MakeReadOnly() {
	self := gdunsafe.Array(pointers.Get(a)[0])
	callBuiltinMethod[bool](&self, builtin.Array.make_read_only, 0|gdextension.SizeArray<<4, nil)
}

func (c Callable) GetMethod() StringName {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	return pointers.New[StringName](callBuiltinMethod[gdextension.StringName](&ptr, builtin.Callable.get_method, gdextension.SizeStringName|gdextension.SizeCallable<<4, nil))
}
func (c Callable) GetBoundArguments() Array {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	return pointers.New[Array](callBuiltinMethod[gdextension.Array](&ptr, builtin.Callable.get_bound_arguments, gdextension.SizeArray|gdextension.SizeCallable<<4, nil))
}
func (c Callable) GetArgumentCount() int64 {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	return callBuiltinMethod[int64](&ptr, builtin.Callable.get_argument_count, gdextension.SizeInt|gdextension.SizeCallable<<4, nil)
}
func (c Callable) Call(args ...Variant) Variant {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	var array = NewArray()
	array.Resize(int64(len(args)))
	for i, arg := range args {
		array.SetIndex(int64(i), arg)
	}
	defer array.Free()
	return pointers.New[Variant](callBuiltinMethod[gdextension.Variant](&ptr, builtin.Callable.callv, gdextension.SizeVariant|gdextension.SizeCallable<<4|gdextension.SizeArray<<8, unsafe.Pointer(&struct {
		gdextension.Array
	}{
		pointers.Get(array),
	})))
}
func (c Callable) CallDeferred() Variant {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	return pointers.New[Variant](callBuiltinMethod[gdextension.Variant](&ptr, builtin.Callable.call_deferred, gdextension.SizeVariant|gdextension.SizeCallable<<4, nil))
}
func (c Callable) Bind(args ...Variant) Callable {
	var ptr = gdunsafe.Callable(pointers.Get(c))
	var array = NewArray()
	array.Resize(int64(len(args)))
	for i, arg := range args {
		array.SetIndex(int64(i), arg)
	}
	defer array.Free()
	return pointers.New[Callable](callBuiltinMethod[gdextension.Callable](&ptr, builtin.Callable.bindv, gdextension.SizeCallable|gdextension.SizeCallable<<4|gdextension.SizeArray<<8, unsafe.Pointer(&struct {
		gdextension.Array
	}{
		pointers.Get(array),
	})))
}

func (d Dictionary) Keys() Array {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	return pointers.New[Array](callBuiltinMethod[gdextension.Array](&self, builtin.Dictionary.keys, gdextension.SizeArray|gdextension.SizeDictionary<<4, nil))
}
func (d Dictionary) Has(key Variant) bool {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	return callBuiltinMethod[bool](&self, builtin.Dictionary.has, gdextension.SizeBool|gdextension.SizeDictionary<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		gdextension.Variant
	}{
		gdextension.Variant(pointers.Get(key)),
	}))
}
func (d Dictionary) Clear() {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	callBuiltinMethod[bool](&self, builtin.Dictionary.clear, 0|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) Sort() {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	callBuiltinMethod[bool](&self, builtin.Dictionary.sort, 0|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) Erase(key Variant) bool {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	return callBuiltinMethod[bool](&self, builtin.Dictionary.erase, gdextension.SizeBool|gdextension.SizeDictionary<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		gdextension.Variant
	}{
		gdextension.Variant(pointers.Get(key)),
	}))
}
func (d Dictionary) Hash() int64 {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	return callBuiltinMethod[int64](&self, builtin.Dictionary.hash, gdextension.SizeInt|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) Size() int64 {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	return callBuiltinMethod[int64](&self, builtin.Dictionary.size, gdextension.SizeInt|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) IsReadOnly() bool {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	return callBuiltinMethod[bool](&self, builtin.Dictionary.is_read_only, gdextension.SizeBool|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) MakeReadOnly() {
	self := gdunsafe.Dictionary(pointers.Get(d)[0])
	callBuiltinMethod[bool](&self, builtin.Dictionary.make_read_only, 0|gdextension.SizeDictionary<<4, nil)
}

func (a PackedByteArray) Resize(size Int) Int {
	self := gdunsafe.PackedArray[byte](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedByteArray.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[byte](self))
	return result
}
func (a PackedByteArray) Size() int64 {
	self := gdunsafe.PackedArray[byte](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedByteArray.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedByteArray) Duplicate() PackedByteArray {
	self := gdunsafe.PackedArray[byte](pointers.Get(a))
	return pointers.New[PackedByteArray](callBuiltinMethod[gdextension.PackedArray[byte]](&self, builtin.PackedByteArray.duplicate, gdextension.SizePackedArray|gdextension.SizePackedArray<<4, nil))
}
func (a PackedColorArray) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Color](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedColorArray.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[Color](self))
	return result
}
func (a PackedColorArray) Size() int64 {
	self := gdunsafe.PackedArray[Color](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedColorArray.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedFloat32Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[float32](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedFloat32Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[float32](self))
	return result
}
func (a PackedFloat32Array) Size() int64 {
	self := gdunsafe.PackedArray[float32](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedFloat32Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedFloat64Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[float64](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedFloat64Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[float64](self))
	return result
}
func (a PackedFloat64Array) Size() int64 {
	self := gdunsafe.PackedArray[float64](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedFloat64Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedInt32Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[int32](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedInt32Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[int32](self))
	return result
}
func (a PackedInt32Array) Size() int64 {
	self := gdunsafe.PackedArray[int32](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedInt32Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedStringArray) Resize(size Int) Int {
	self := gdunsafe.PackedArray[gdunsafe.String](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedStringArray.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[gdextension.String](self))
	return result
}
func (a PackedStringArray) Size() int64 {
	self := gdunsafe.PackedArray[gdunsafe.String](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedStringArray.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedVector2Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Vector2](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedVector2Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[Vector2](self))
	return result
}
func (a PackedVector2Array) Size() int64 {
	self := gdunsafe.PackedArray[Vector2](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedVector2Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedVector3Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Vector3](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedVector3Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[Vector3](self))
	return result
}
func (a PackedVector3Array) Size() int64 {
	self := gdunsafe.PackedArray[Vector3](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedVector3Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedVector4Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[Vector4](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedVector4Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[Vector4](self))
	return result
}
func (a PackedVector4Array) Size() int64 {
	self := gdunsafe.PackedArray[Vector4](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedVector4Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedInt64Array) Resize(size Int) Int {
	self := gdunsafe.PackedArray[int64](pointers.Get(a))
	result := Int(callBuiltinMethod[int64](&self, builtin.PackedInt64Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
	pointers.Set(a, gdextension.PackedArray[int64](self))
	return result
}
func (a PackedInt64Array) Size() int64 {
	self := gdunsafe.PackedArray[int64](pointers.Get(a))
	return callBuiltinMethod[int64](&self, builtin.PackedInt64Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}

func (s Signal) Emit(args ...Variant) {
	var converted = make([]gdextension.Variant, len(args))
	for i, arg := range args {
		converted[i] = gdextension.Variant(pointers.Get(arg))
	}
	self := gdunsafe.Signal(pointers.Get(s))
	callBuiltinMethod[struct{}](&self, builtin.Signal.emit, gdextension.SizeSignal<<4|gdextension.ShapeVariants(len(args))<<4, unsafe.Pointer(unsafe.SliceData(converted)))
}

func (s Signal) Connect(callable Callable, flags int64) int64 {
	self := gdunsafe.Signal(pointers.Get(s))
	return callBuiltinMethod[int64](&self, builtin.Signal.connect, gdextension.SizeInt|gdextension.SizeSignal<<4|gdextension.SizeCallable<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		gdextension.Callable
		Flags int64
	}{
		gdextension.Callable(pointers.Get(callable)), flags,
	}))
}
func (s Signal) Disconnect(callable Callable) {
	self := gdunsafe.Signal(pointers.Get(s))
	callBuiltinMethod[bool](&self, builtin.Signal.disconnect, 0|gdextension.SizeSignal<<4|gdextension.SizeCallable<<8, unsafe.Pointer(&struct {
		gdextension.Callable
	}{
		gdextension.Callable(pointers.Get(callable)),
	}))
}
func (s Signal) GetName() StringName {
	self := gdunsafe.Signal(pointers.Get(s))
	return pointers.New[StringName](callBuiltinMethod[gdextension.StringName](&self, builtin.Signal.get_name, gdextension.SizeStringName|gdextension.SizeSignal<<4, nil))
}
func (s Signal) GetConnections() Array {
	self := gdunsafe.Signal(pointers.Get(s))
	return pointers.New[Array](callBuiltinMethod[gdextension.Array](&self, builtin.Signal.get_connections, gdextension.SizeArray|gdextension.SizeSignal<<4, nil))
}
func (s Signal) GetObject() gdreference.Object {
	self := gdunsafe.Signal(pointers.Get(s))
	return gdreference.OwnObject(callBuiltinMethod[gdextension.Object](&self, builtin.Signal.get_object, gdextension.SizeObject|gdextension.SizeSignal<<4, nil), Free)
}
func (s String) Length() int64 {
	self := gdunsafe.String(pointers.Get(s)[0])
	return callBuiltinMethod[int64](&self, builtin.String.length, gdextension.SizeInt|gdextension.SizeString<<4, nil)
}
func (s String) Substr(begin, end int64) String {
	self := gdunsafe.String(pointers.Get(s)[0])
	return pointers.New[String](callBuiltinMethod[gdextension.String](&self, builtin.String.substr, gdextension.SizeString|gdextension.SizeString<<4|gdextension.SizeInt<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		Begin int64
		End   int64
	}{
		begin, end,
	})))
}
func (s String) CasecmpTo(other String) int64 {
	self := gdunsafe.String(pointers.Get(s)[0])
	return callBuiltinMethod[int64](&self, builtin.String.casecmp_to, gdextension.SizeInt|gdextension.SizeString<<4|gdextension.SizeString<<8, unsafe.Pointer(&struct {
		Other gdextension.String
	}{
		pointers.Get(other),
	}))
}
func (s StringName) Length() int64 {
	self := gdunsafe.StringName(pointers.Get(s)[0])
	return callBuiltinMethod[int64](&self, builtin.StringName.length, gdextension.SizeInt|gdextension.SizeStringName<<4, nil)
}
func (s StringName) Substr(begin, end int64) String {
	self := gdunsafe.StringName(pointers.Get(s)[0])
	return pointers.New[String](callBuiltinMethod[gdextension.String](&self, builtin.StringName.substr, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeInt<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		Begin int64
		End   int64
	}{
		begin, end,
	})))
}
func (s StringName) CasecmpTo(other String) int64 {
	self := gdunsafe.StringName(pointers.Get(s)[0])
	return callBuiltinMethod[int64](&self, builtin.StringName.casecmp_to, gdextension.SizeInt|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8, unsafe.Pointer(&struct {
		Other gdextension.String
	}{
		pointers.Get(other),
	}))
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

func ObjectGet(o gdreference.Object, name StringName) Variant {
	return pointers.New[Variant](noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get, gdextension.SizeVariant|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdextension.StringName
	}{
		pointers.Get(name),
	})))
}
func ObjectSet(o gdreference.Object, name StringName, value Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set, 0|gdextension.SizeStringName<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		Name  gdextension.StringName
		Value gdextension.Variant
	}{
		pointers.Get(name), gdextension.Variant(pointers.Get(value)),
	}))
}

func ObjectGetMeta(o gdreference.Object, name StringName) Variant {
	return pointers.New[Variant](noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get_meta, gdextension.SizeVariant|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdextension.StringName
	}{
		pointers.Get(name),
	})))
}
func ObjectSetMeta(o gdreference.Object, name StringName, value Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set_meta, 0|gdextension.SizeStringName<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		Name  gdextension.StringName
		Value gdextension.Variant
	}{
		pointers.Get(name), gdextension.Variant(pointers.Get(value)),
	}))
}

func ObjectHasMethod(o gdreference.Object, name StringName) bool {
	return noescape.Call[bool](gdreference.GetObject(o), object_methods.has_method, gdextension.SizeBool|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdextension.StringName
	}{
		pointers.Get(name),
	}))
}
func ObjectCall(o gdreference.Object, method StringName, args ...Variant) (Variant, error) {
	ring.Main.Flush()
	self := gdreference.GetObject(o)
	name := pointers.Get(method)
	if gdunsafe.Script(self).HasMethod(gdunsafe.StringName(name[0])) {
		var converted []gdunsafe.Variant
		for _, arg := range args {
			converted = append(converted, gdunsafe.Variant(pointers.Get(arg)))
		}
		result, err := gdunsafe.Script(self).Call(gdunsafe.StringName(name[0]),
			converted...,
		)
		if err != (gdunsafe.Error{}) {
			return pointers.New[Variant](result), err
		}
		return pointers.New[Variant](result), nil
	}
	return NewVariant(o).Call(method, args...) // FIXME is this ok?
}

func ObjectCanTranslateMessages(o gdreference.Object) bool {
	return jumponly.Call[bool](gdreference.GetObject(o), object_methods.can_translate_messages, gdextension.SizeBool, nil)
}
func ObjectGetScript(o gdreference.Object) Variant {
	return pointers.New[Variant](noescape.Call[gdextension.Variant](gdreference.GetObject(o), object_methods.get_script, gdextension.SizeVariant, nil))
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
func ObjectSetScript(o gdreference.Object, script Variant) {
	noescape.Call[struct{}](gdreference.GetObject(o), object_methods.set_script, 0|gdextension.SizeVariant<<4, unsafe.Pointer(&struct {
		Script gdextension.Variant
	}{
		gdextension.Variant(pointers.Get(script)),
	}))
}
func ObjectToString(o gdreference.Object) String {
	return pointers.New[String](noescape.Call[gdextension.String](gdreference.GetObject(o), object_methods.to_string, gdextension.SizeString, nil))
}
func ObjectTr(o gdreference.Object, message StringName, context StringName) String {
	return pointers.New[String](noescape.Call[gdextension.String](gdreference.GetObject(o), object_methods.tr, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8, unsafe.Pointer(&struct {
		Message gdextension.StringName
		Context gdextension.StringName
	}{
		pointers.Get(message), pointers.Get(context),
	})))
}
func ObjectTrN(o gdreference.Object, message StringName, plural StringName, n int64, context StringName) String {
	return pointers.New[String](noescape.Call[gdextension.String](gdreference.GetObject(o), object_methods.tr_n, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8|gdextension.SizeInt<<12|gdextension.SizeStringName<<16, unsafe.Pointer(&struct {
		Message gdextension.StringName
		Plural  gdextension.StringName
		N       int64
		Context gdextension.StringName
	}{
		pointers.Get(message), pointers.Get(plural), n, pointers.Get(context),
	})))
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
func ObjectGetClass(o gdreference.Object) String {
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
