package gd

import (
	"grow.graphics/gd/internal/callframe"
	"runtime.link/mmm"
)

type Array mmm.Pointer[API, Array, uintptr]

func (a Array) Index(ctx Context, index Int) Variant {
	return mmm.API(a).Array.Index(ctx, a, index)
}

func (a Array) SetIndex(index Int, value Variant) {
	mmm.API(a).Array.SetIndex(a, index, value)
}

func (a Array) isArray() {}

func (a Array) end() { mmm.End(a) }

func (a Array) Free() {
	var frame = callframe.New()
	mmm.API(a).typeset.destruct.Array(callframe.Arg(frame, mmm.End(a)).Uintptr())
	frame.Free()
}

type IsArray interface {
	mmm.ManagedPointer[uintptr]

	end()
	isArray()
}

type IsArrayType[T any] interface {
	Typed(TypedArray[T])
	Array() Array
}

type ArrayOf[T any] interface {
	IsArray
	IsArrayType[T]

	All(method Callable) bool
	Any(method Callable) bool
	Append(value T)
	AppendArray(array ArrayOf[T])
	Assign(array ArrayOf[T])
	Back(ctx Context) T
	Bsearch(value T, before bool) int64
	BsearchCustom(value T, fn Callable, before bool) int64
	Clear()
	Count(value T) int64
	Duplicate(ctx Context, deep bool) ArrayOf[T]
	Erase(value T)
	Fill(value T)
	Filter(ctx Context, method Callable) ArrayOf[T]
	Find(what T, from int64) int64
	Free()
	Front(ctx Context) T
	GetTypedBuiltin() int64
	GetTypedClassName(ctx Context) StringName
	GetTypedScript(ctx Context) Variant
	Has(value T) bool
	Hash() int64
	Index(ctx Context, index int64) T
	Insert(position int64, value T) int64
	IsEmpty() bool
	IsReadOnly() bool
	IsSameTyped(array Array) bool
	IsTyped() bool
	MakeReadOnly()
	Map(ctx Context, method Callable) ArrayOf[T]
	Max(ctx Context) T
	Min(ctx Context) T
	PickRandom(ctx Context) T
	PopAt(ctx Context, position int64) T
	PopBack(ctx Context) T
	PopFront(ctx Context) T
	PushBack(value T)
	PushFront(value T)
	Reduce(ctx Context, method Callable, accum T) T
	RemoveAt(position int64)
	Resize(size int64) int64
	Reverse()
	Rfind(what T, from int64) int64
	SetIndex(index int64, value T)
	Shuffle()
	Size() int64
	Slice(ctx Context, begin int64, end int64, step int64, deep bool) ArrayOf[T]
	Sort()
	SortCustom(fn Callable)
}

func (godot Context) Array() Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	godot.API.typeset.creation.Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return mmm.New[Array](godot.Lifetime, godot.API, raw)
}

type TypedArray[T any] Array

func (a TypedArray[T]) Typed(TypedArray[T]) {}
func (a TypedArray[T]) isArray()            {}
func (a TypedArray[T]) Array() Array        { return Array(a) }

func (a TypedArray[T]) All(method Callable) bool { return Array(a).All(method) }
func (a TypedArray[T]) Any(method Callable) bool { return Array(a).Any(method) }
func (a TypedArray[T]) Append(value T) {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	Array(a).Append(godot.Variant(value))
}
func (a TypedArray[T]) AppendArray(array ArrayOf[T]) {
	Array(a).AppendArray(Array(array.(TypedArray[T])))
}
func (a TypedArray[T]) Assign(array ArrayOf[T]) {
	Array(a).Assign(Array(array.(TypedArray[T])))
}
func (a TypedArray[T]) Back(godot Context) T {
	return Array(a).Back(godot).Interface(godot).(T)
}
func (a TypedArray[T]) Bsearch(value T, before bool) int64 {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).Bsearch(godot.Variant(value), before)
}
func (a TypedArray[T]) BsearchCustom(value T, fn Callable, before bool) int64 {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).BsearchCustom(godot.Variant(value), fn, before)
}
func (a TypedArray[T]) Clear() { Array(a).Clear() }
func (a TypedArray[T]) Count(value T) int64 {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).Count(godot.Variant(value))
}
func (a TypedArray[T]) Duplicate(ctx Context, deep bool) ArrayOf[T] {
	return TypedArray[T](Array(a).Duplicate(ctx, deep))
}
func (a TypedArray[T]) Erase(value T) {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	Array(a).Erase(godot.Variant(value))
}
func (a TypedArray[T]) Fill(value T) {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	Array(a).Fill(godot.Variant(value))
}
func (a TypedArray[T]) Filter(ctx Context, method Callable) ArrayOf[T] {
	return TypedArray[T](Array(a).Filter(ctx, method))
}
func (a TypedArray[T]) Find(what T, from int64) int64 {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).Find(godot.Variant(what), from)
}
func (a TypedArray[T]) Free() {
	Array(a).Free()
}
func (a TypedArray[T]) Front(ctx Context) T {
	return Array(a).Front(ctx).Interface(ctx).(T)
}
func (a TypedArray[T]) GetTypedBuiltin() int64 {
	return Array(a).GetTypedBuiltin()
}
func (a TypedArray[T]) GetTypedClassName(ctx Context) StringName {
	return Array(a).GetTypedClassName(ctx)
}
func (a TypedArray[T]) GetTypedScript(ctx Context) Variant {
	return Array(a).GetTypedScript(ctx)
}
func (a TypedArray[T]) Has(value T) bool {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).Has(godot.Variant(value))
}
func (a TypedArray[T]) Hash() int64 {
	return Array(a).Hash()
}
func (a TypedArray[T]) Index(ctx Context, index int64) T {
	return Array(a).Index(ctx, index).Interface(ctx).(T)
}
func (a TypedArray[T]) Insert(position int64, value T) int64 {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).Insert(position, godot.Variant(value))
}
func (a TypedArray[T]) IsEmpty() bool {
	return Array(a).IsEmpty()
}
func (a TypedArray[T]) IsReadOnly() bool {
	return Array(a).IsReadOnly()
}
func (a TypedArray[T]) IsSameTyped(array Array) bool {
	return Array(a).IsSameTyped(array)
}
func (a TypedArray[T]) IsTyped() bool {
	return Array(a).IsTyped()
}
func (a TypedArray[T]) MakeReadOnly() {
	Array(a).MakeReadOnly()
}
func (a TypedArray[T]) Map(ctx Context, method Callable) ArrayOf[T] {
	return TypedArray[T](Array(a).Map(ctx, method))
}
func (a TypedArray[T]) Max(ctx Context) T {
	return Array(a).Max(ctx).Interface(ctx).(T)
}
func (a TypedArray[T]) Min(ctx Context) T {
	return Array(a).Min(ctx).Interface(ctx).(T)
}
func (a TypedArray[T]) PickRandom(ctx Context) T {
	return Array(a).PickRandom(ctx).Interface(ctx).(T)
}
func (a TypedArray[T]) PopAt(ctx Context, position int64) T {
	return Array(a).PopAt(ctx, position).Interface(ctx).(T)
}
func (a TypedArray[T]) PopBack(ctx Context) T {
	return Array(a).PopBack(ctx).Interface(ctx).(T)
}
func (a TypedArray[T]) PopFront(ctx Context) T {
	return Array(a).PopFront(ctx).Interface(ctx).(T)
}
func (a TypedArray[T]) PushBack(value T) {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	Array(a).PushBack(godot.Variant(value))
}
func (a TypedArray[T]) PushFront(value T) {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	Array(a).PushFront(godot.Variant(value))
}
func (a TypedArray[T]) Reduce(ctx Context, method Callable, accum T) T {
	return Array(a).Reduce(ctx, method, ctx.Variant(accum)).Interface(ctx).(T)
}
func (a TypedArray[T]) RemoveAt(position int64) {
	Array(a).RemoveAt(position)
}
func (a TypedArray[T]) Resize(size int64) int64 {
	return Array(a).Resize(size)
}
func (a TypedArray[T]) Reverse() {
	Array(a).Reverse()
}
func (a TypedArray[T]) Rfind(what T, from int64) int64 {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	return Array(a).Rfind(godot.Variant(what), from)
}
func (a TypedArray[T]) SetIndex(index int64, value T) {
	godot := NewContext(mmm.API(Array(a)))
	defer godot.End()
	Array(a).SetIndex(index, godot.Variant(value))
}
func (a TypedArray[T]) Shuffle() {
	Array(a).Shuffle()
}
func (a TypedArray[T]) Size() int64 {
	return Array(a).Size()
}
func (a TypedArray[T]) Slice(ctx Context, begin int64, end int64, step int64, deep bool) ArrayOf[T] {
	return TypedArray[T](Array(a).Slice(ctx, begin, end, step, deep))
}
func (a TypedArray[T]) Sort() {
	Array(a).Sort()

}
func (a TypedArray[T]) SortCustom(fn Callable) {
	Array(a).SortCustom(fn)

}
func (a TypedArray[T]) end() { Array(a).end() }
