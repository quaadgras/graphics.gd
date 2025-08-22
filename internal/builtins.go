package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

func callBuiltinMethod[T any](self unsafe.Pointer, method gdextension.MethodForBuiltinType, shape gdextension.Shape, args unsafe.Pointer) T {
	var result T
	call_builtin_noescape(self, method, unsafe.Pointer(&result), shape, args)
	return result
}

//go:noescape
func call_builtin_noescape(self unsafe.Pointer, method gdextension.MethodForBuiltinType, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call_builtin graphics.gd/internal.call_builtin_noescape
func call_builtin(self unsafe.Pointer, method gdextension.MethodForBuiltinType, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	gdextension.Host.Builtin.Types.Unsafe.Call(gdextension.CallMutates[any](self), method, gdextension.CallReturns[any](result), shape, gdextension.CallAccepts[any](args))
}

// builtin methods that are strictly required for graphics.gd to function.
type builtin struct {
	Array struct {
		size           gdextension.MethodForBuiltinType `hash:"3173160232"`
		resize         gdextension.MethodForBuiltinType `hash:"848867239"`
		is_read_only   gdextension.MethodForBuiltinType `hash:"3918633141"`
		make_read_only gdextension.MethodForBuiltinType `hash:"3218959716"`
	}
	Callable struct {
		get_method          gdextension.MethodForBuiltinType `hash:"1825232092"`
		get_bound_arguments gdextension.MethodForBuiltinType `hash:"4144163970"`
		get_argument_count  gdextension.MethodForBuiltinType `hash:"3173160232"`
		callv               gdextension.MethodForBuiltinType `hash:"413578926"`
		bindv               gdextension.MethodForBuiltinType `hash:"3564560322"`
		call_deferred       gdextension.MethodForBuiltinType `hash:"3286317445"`
	}
	Dictionary struct {
		keys           gdextension.MethodForBuiltinType `hash:"4144163970"`
		has            gdextension.MethodForBuiltinType `hash:"3680194679"`
		clear          gdextension.MethodForBuiltinType `hash:"3218959716"`
		sort           gdextension.MethodForBuiltinType `hash:"3218959716"`
		erase          gdextension.MethodForBuiltinType `hash:"1776646889"`
		hash           gdextension.MethodForBuiltinType `hash:"3173160232"`
		size           gdextension.MethodForBuiltinType `hash:"3173160232"`
		is_read_only   gdextension.MethodForBuiltinType `hash:"3918633141"`
		make_read_only gdextension.MethodForBuiltinType `hash:"3218959716"`
	}
	PackedByteArray struct {
		resize    gdextension.MethodForBuiltinType `hash:"848867239"`
		size      gdextension.MethodForBuiltinType `hash:"3173160232"`
		duplicate gdextension.MethodForBuiltinType `hash:"851781288"`
	}
	PackedColorArray struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedFloat32Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedFloat64Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedInt32Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedStringArray struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedVector2Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedVector3Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedVector4Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	PackedInt64Array struct {
		resize gdextension.MethodForBuiltinType `hash:"848867239"`
		size   gdextension.MethodForBuiltinType `hash:"3173160232"`
	}
	Signal struct {
		emit            gdextension.MethodForBuiltinType `hash:"3286317445"`
		connect         gdextension.MethodForBuiltinType `hash:"979702392"`
		disconnect      gdextension.MethodForBuiltinType `hash:"3470848906"`
		get_name        gdextension.MethodForBuiltinType `hash:"1825232092"`
		get_connections gdextension.MethodForBuiltinType `hash:"4144163970"`
		get_object      gdextension.MethodForBuiltinType `hash:"4008621732"`
	}
	String struct {
		length     gdextension.MethodForBuiltinType `hash:"3173160232"`
		substr     gdextension.MethodForBuiltinType `hash:"787537301"`
		casecmp_to gdextension.MethodForBuiltinType `hash:"2920860731"`
	}
	StringName struct {
		length     gdextension.MethodForBuiltinType `hash:"3173160232"`
		substr     gdextension.MethodForBuiltinType `hash:"787537301"`
		casecmp_to gdextension.MethodForBuiltinType `hash:"2920860731"`
	}
}

func (a Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.Array.size, gdextension.SizeInt|gdextension.SizeArray<<4, nil)
}

func (a Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.Array.resize, 0|gdextension.SizeArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a Array) IsReadOnly() bool {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Array.is_read_only, gdextension.SizeBool|gdextension.SizeArray<<4, nil)
}
func (a Array) MakeReadOnly() {
	var ptr = pointers.Get(a)
	callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Array.make_read_only, 0|gdextension.SizeArray<<4, nil)
}

func (c Callable) GetMethod() StringName {
	var ptr = pointers.Get(c)
	return pointers.New[StringName](callBuiltinMethod[gdextension.StringName](unsafe.Pointer(&ptr), Global.builtin.Callable.get_method, gdextension.SizeStringName|gdextension.SizeCallable<<4, nil))
}
func (c Callable) GetBoundArguments() Array {
	var ptr = pointers.Get(c)
	return pointers.New[Array](callBuiltinMethod[gdextension.Array](unsafe.Pointer(&ptr), Global.builtin.Callable.get_bound_arguments, gdextension.SizeArray|gdextension.SizeCallable<<4, nil))
}
func (c Callable) GetArgumentCount() int64 {
	var ptr = pointers.Get(c)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.Callable.get_argument_count, gdextension.SizeInt|gdextension.SizeCallable<<4, nil)
}
func (c Callable) Call(args ...Variant) Variant {
	var ptr = pointers.Get(c)
	var array = NewArray()
	array.Resize(int64(len(args)))
	for i, arg := range args {
		array.SetIndex(int64(i), arg)
	}
	defer array.Free()
	return pointers.New[Variant]([3]uint64(callBuiltinMethod[gdextension.Variant](unsafe.Pointer(&ptr), Global.builtin.Callable.callv, gdextension.SizeVariant|gdextension.SizeCallable<<4|gdextension.SizeArray<<8, unsafe.Pointer(&struct {
		gdextension.Array
	}{
		pointers.Get(array),
	}))))
}
func (c Callable) CallDeferred() Variant {
	var ptr = gdextension.Callable(pointers.Get(c))
	return pointers.New[Variant]([3]uint64(callBuiltinMethod[gdextension.Variant](unsafe.Pointer(&ptr), Global.builtin.Callable.call_deferred, gdextension.SizeVariant|gdextension.SizeCallable<<4, nil)))
}
func (c Callable) Bind(args ...Variant) Callable {
	var ptr = gdextension.Callable(pointers.Get(c))
	var array = NewArray()
	array.Resize(int64(len(args)))
	for i, arg := range args {
		array.SetIndex(int64(i), arg)
	}
	defer array.Free()
	return pointers.New[Callable]([2]uint64(callBuiltinMethod[gdextension.Callable](unsafe.Pointer(&ptr), Global.builtin.Callable.bindv, gdextension.SizeCallable|gdextension.SizeCallable<<4|gdextension.SizeArray<<8, unsafe.Pointer(&struct {
		gdextension.Array
	}{
		pointers.Get(array),
	}))))
}

func (d Dictionary) Keys() Array {
	var ptr = pointers.Get(d)
	return pointers.New[Array](callBuiltinMethod[gdextension.Array](unsafe.Pointer(&ptr), Global.builtin.Dictionary.keys, gdextension.SizeArray|gdextension.SizeDictionary<<4, nil))
}
func (d Dictionary) Has(key Variant) bool {
	var ptr = pointers.Get(d)
	return callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Dictionary.has, gdextension.SizeBool|gdextension.SizeDictionary<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		gdextension.Variant
	}{
		gdextension.Variant(pointers.Get(key)),
	}))
}
func (d Dictionary) Clear() {
	var ptr = pointers.Get(d)
	callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Dictionary.clear, 0|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) Sort() {
	var ptr = pointers.Get(d)
	callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Dictionary.sort, 0|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) Erase(key Variant) bool {
	var ptr = pointers.Get(d)
	return callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Dictionary.erase, gdextension.SizeBool|gdextension.SizeDictionary<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		gdextension.Variant
	}{
		gdextension.Variant(pointers.Get(key)),
	}))
}
func (d Dictionary) Hash() int64 {
	var ptr = pointers.Get(d)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.Dictionary.hash, gdextension.SizeInt|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) Size() int64 {
	var ptr = pointers.Get(d)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.Dictionary.size, gdextension.SizeInt|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) IsReadOnly() bool {
	var ptr = pointers.Get(d)
	return callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Dictionary.is_read_only, gdextension.SizeBool|gdextension.SizeDictionary<<4, nil)
}
func (d Dictionary) MakeReadOnly() {
	var ptr = pointers.Get(d)
	callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Dictionary.make_read_only, 0|gdextension.SizeDictionary<<4, nil)
}

func (a PackedByteArray) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedByteArray.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedByteArray) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedByteArray.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedByteArray) Duplicate() PackedByteArray {
	var ptr = pointers.Get(a)
	return pointers.New[PackedByteArray](callBuiltinMethod[gdextension.PackedArray[byte]](unsafe.Pointer(&ptr), Global.builtin.PackedByteArray.duplicate, gdextension.SizePackedArray|gdextension.SizePackedArray<<4, nil))
}
func (a PackedColorArray) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedColorArray.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedColorArray) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedColorArray.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedFloat32Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedFloat32Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedFloat32Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedFloat32Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedFloat64Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedFloat64Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedFloat64Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedFloat64Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedInt32Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedInt32Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedInt32Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedInt32Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedStringArray) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedStringArray.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedStringArray) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedStringArray.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedVector2Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedVector2Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedVector2Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedVector2Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedVector3Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedVector3Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedVector3Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedVector3Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedVector4Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedVector4Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedVector4Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedVector4Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}
func (a PackedInt64Array) Resize(size Int) Int {
	var ptr = pointers.Get(a)
	defer func() { pointers.Set(a, ptr) }()
	return Int(callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedInt64Array.resize, 0|gdextension.SizePackedArray<<4|gdextension.SizeInt<<8, unsafe.Pointer(&struct {
		Size int64
	}{
		int64(size),
	})))
}
func (a PackedInt64Array) Size() int64 {
	var ptr = pointers.Get(a)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.PackedInt64Array.size, gdextension.SizeInt|gdextension.SizePackedArray<<4, nil)
}

func (s Signal) Emit(args ...Variant) {
	var ptr = gdextension.Signal(pointers.Get(s))
	var converted = make([]gdextension.Variant, len(args))
	for _, arg := range args {
		converted = append(converted, gdextension.Variant(pointers.Get(arg)))
	}
	callBuiltinMethod[struct{}](unsafe.Pointer(&ptr), Global.builtin.Signal.emit, gdextension.ShapeVariants(len(args)), unsafe.Pointer(unsafe.SliceData(converted)))
}

func (s Signal) Connect(callable Callable, flags int64) int64 {
	var ptr = gdextension.Signal(pointers.Get(s))
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.Signal.connect, gdextension.SizeInt|gdextension.SizeSignal<<4|gdextension.SizeCallable<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		gdextension.Callable
		Flags int64
	}{
		gdextension.Callable(pointers.Get(callable)), flags,
	}))
}
func (s Signal) Disconnect(callable Callable) {
	var ptr = gdextension.Signal(pointers.Get(s))
	callBuiltinMethod[bool](unsafe.Pointer(&ptr), Global.builtin.Signal.disconnect, 0|gdextension.SizeSignal<<4|gdextension.SizeCallable<<8, unsafe.Pointer(&struct {
		gdextension.Callable
	}{
		gdextension.Callable(pointers.Get(callable)),
	}))
}
func (s Signal) GetName() StringName {
	var ptr = gdextension.Signal(pointers.Get(s))
	return pointers.New[StringName](callBuiltinMethod[gdextension.StringName](unsafe.Pointer(&ptr), Global.builtin.Signal.get_name, gdextension.SizeStringName|gdextension.SizeSignal<<4, nil))
}
func (s Signal) GetConnections() Array {
	var ptr = gdextension.Signal(pointers.Get(s))
	return pointers.New[Array](callBuiltinMethod[gdextension.Array](unsafe.Pointer(&ptr), Global.builtin.Signal.get_connections, gdextension.SizeArray|gdextension.SizeSignal<<4, nil))
}
func (s Signal) GetObject() Object {
	var ptr = gdextension.Signal(pointers.Get(s))
	return pointers.New[Object]([3]uint64{uint64(callBuiltinMethod[gdextension.Object](unsafe.Pointer(&ptr), Global.builtin.Signal.get_object, gdextension.SizeObject|gdextension.SizeSignal<<4, nil))})
}
func (s String) Length() int64 {
	var ptr = pointers.Get(s)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.String.length, gdextension.SizeInt|gdextension.SizeString<<4, nil)
}
func (s String) Substr(begin, end int64) String {
	var ptr = pointers.Get(s)
	return pointers.New[String](callBuiltinMethod[gdextension.String](unsafe.Pointer(&ptr), Global.builtin.String.substr, gdextension.SizeString|gdextension.SizeString<<4|gdextension.SizeInt<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		Begin int64
		End   int64
	}{
		begin, end,
	})))
}
func (s String) CasecmpTo(other String) int64 {
	var ptr = pointers.Get(s)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.String.casecmp_to, gdextension.SizeInt|gdextension.SizeString<<4|gdextension.SizeString<<8, unsafe.Pointer(&struct {
		Other gdextension.String
	}{
		pointers.Get(other),
	}))
}
func (s StringName) Length() int64 {
	var ptr = pointers.Get(s)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.StringName.length, gdextension.SizeInt|gdextension.SizeStringName<<4, nil)
}
func (s StringName) Substr(begin, end int64) String {
	var ptr = pointers.Get(s)
	return pointers.New[String](callBuiltinMethod[gdextension.String](unsafe.Pointer(&ptr), Global.builtin.StringName.substr, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeInt<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		Begin int64
		End   int64
	}{
		begin, end,
	})))
}
func (s StringName) CasecmpTo(other String) int64 {
	var ptr = pointers.Get(s)
	return callBuiltinMethod[int64](unsafe.Pointer(&ptr), Global.builtin.StringName.casecmp_to, gdextension.SizeInt|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8, unsafe.Pointer(&struct {
		Other gdextension.String
	}{
		pointers.Get(other),
	}))
}

func (o Object) Get(name StringName) Variant {
	return pointers.New[Variant]([3]uint64(gdextension.Call[gdextension.Variant](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_get, gdextension.SizeVariant|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdextension.StringName
	}{
		pointers.Get(name),
	}))))
}
func (o Object) Set(name StringName, value Variant) {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_set, 0|gdextension.SizeStringName<<4|gdextension.SizeVariant<<8, unsafe.Pointer(&struct {
		Name  gdextension.StringName
		Value gdextension.Variant
	}{
		pointers.Get(name), gdextension.Variant(pointers.Get(value)),
	}))
}
func (o Object) HasMethod(name StringName) bool {
	return gdextension.Call[bool](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_has_method, gdextension.SizeBool|gdextension.SizeStringName<<4, unsafe.Pointer(&struct {
		Name gdextension.StringName
	}{
		pointers.Get(name),
	}))
}
func (o Object) Call(method StringName, args ...Variant) (Variant, error) {
	self := gdextension.Object(pointers.Get(o.AsObject()[0])[0])
	name := pointers.Get(method)
	if gdextension.Host.Objects.Script.DefinesMethod(self, name) {
		var converted []gdextension.Variant
		for _, arg := range args {
			converted = append(converted, gdextension.Variant(pointers.Get(arg)))
		}
		var err gdextension.CallError
		var result gdextension.Variant
		gdextension.Host.Objects.Script.Call(self, name,
			gdextension.CallReturns[gdextension.Variant](&result),
			len(args),
			gdextension.CallAccepts[gdextension.Variant](unsafe.SliceData(converted)),
			gdextension.CallReturns[gdextension.CallError](&err),
		)
		return pointers.New[Variant]([3]uint64(result)), err.Err()
	}
	return NewVariant(o).Call(method, args...) // FIXME is this ok?
}

func (o Object) CanTranslateMessages() bool {
	return gdextension.Call[bool](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_can_translate_messages, gdextension.SizeBool, nil)
}
func (o Object) GetScript() Variant {
	return pointers.New[Variant]([3]uint64(gdextension.Call[gdextension.Variant](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_get_script, gdextension.SizeVariant, nil)))
}
func (o Object) NotifyPropertyListChanged() {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_notify_property_list_changed, 0, nil)
}
func (o Object) SetBlockSignals(blocking bool) {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_set_block_signals, 0|gdextension.SizeBool<<4, unsafe.Pointer(&struct {
		Blocking bool
	}{
		blocking,
	}))
}
func (o Object) SetScript(script Variant) {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_set_script, 0|gdextension.SizeVariant<<4, unsafe.Pointer(&struct {
		Script gdextension.Variant
	}{
		gdextension.Variant(pointers.Get(script)),
	}))
}
func (o Object) ToString() String {
	return pointers.New[String](gdextension.Call[gdextension.String](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_to_string, gdextension.SizeString, nil))
}
func (o Object) Tr(message StringName, context StringName) String {
	return pointers.New[String](gdextension.Call[gdextension.String](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_tr, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8, unsafe.Pointer(&struct {
		Message gdextension.StringName
		Context gdextension.StringName
	}{
		pointers.Get(message), pointers.Get(context),
	})))
}
func (o Object) TrN(message StringName, plural StringName, n int64, context StringName) String {
	return pointers.New[String](gdextension.Call[gdextension.String](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_tr_n, gdextension.SizeString|gdextension.SizeStringName<<4|gdextension.SizeStringName<<8|gdextension.SizeInt<<12|gdextension.SizeStringName<<16, unsafe.Pointer(&struct {
		Message gdextension.StringName
		Plural  gdextension.StringName
		N       int64
		Context gdextension.StringName
	}{
		pointers.Get(message), pointers.Get(plural), n, pointers.Get(context),
	})))
}
func (o Object) SetMessageTranslation(enable bool) {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_set_message_translation, 0|gdextension.SizeBool<<4, unsafe.Pointer(&struct {
		Enable bool
	}{
		enable,
	}))
}
func (o Object) IsBlockingSignals() bool {
	return gdextension.Call[bool](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_is_blocking_signals, gdextension.SizeBool, nil)
}
func (o Object) GetClass() String {
	return pointers.New[String](gdextension.Call[gdextension.String](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_get_class, gdextension.SizeString, nil))
}
func (o Object) Connect(signal StringName, callable Callable, flags int64) int64 {
	return gdextension.Call[int64](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_connect, gdextension.SizeInt|gdextension.SizeStringName<<4|gdextension.SizeCallable<<8|gdextension.SizeInt<<12, unsafe.Pointer(&struct {
		Signal   gdextension.StringName
		Callable gdextension.Callable
		Flags    int64
	}{
		pointers.Get(signal), gdextension.Callable(pointers.Get(callable)), flags,
	}))
}
func (o Object) IsConnected(signal StringName, callable Callable) bool {
	return gdextension.Call[bool](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_is_connected, gdextension.SizeBool|gdextension.SizeStringName<<4|gdextension.SizeCallable<<8, unsafe.Pointer(&struct {
		Signal   gdextension.StringName
		Callable gdextension.Callable
	}{
		pointers.Get(signal), gdextension.Callable(pointers.Get(callable)),
	}))
}
func (o Object) Disconnect(signal StringName, callable Callable) {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_disconnect, 0|gdextension.SizeStringName<<4|gdextension.SizeCallable<<8, unsafe.Pointer(&struct {
		Signal   gdextension.StringName
		Callable gdextension.Callable
	}{
		pointers.Get(signal), gdextension.Callable(pointers.Get(callable)),
	}))
}
func (o Object) IsQueuedForDeletion() bool {
	return gdextension.Call[bool](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_is_queued_for_deletion, gdextension.SizeBool, nil)
}
func (o Object) Notification(what Int, reversed bool) {
	gdextension.Call[struct{}](ObjectChecked(o.AsObject()), Global.Methods.Object.Bind_notification, 0|gdextension.SizeInt<<4|gdextension.SizeBool<<8, unsafe.Pointer(&struct {
		What     int64
		Reversed bool
	}{
		int64(what), reversed,
	}))
}

func (rc RefCounted) Reference() {
	gdextension.Call[struct{}](ObjectChecked(rc.AsObject()), Global.Methods.RefCounted.Bind_reference, 0, nil)
}
func (rc RefCounted) Unreference() bool {
	return gdextension.Call[bool](ObjectChecked(rc.AsObject()), Global.Methods.RefCounted.Bind_unreference, gdextension.SizeBool, nil)
}
