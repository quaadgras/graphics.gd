package gd

import (
	"iter"

	"grow.graphics/gd/internal/callframe"
	"grow.graphics/gd/internal/pointers"
)

func (a Array) Index(index Int) Variant {
	return Global.Array.Index(a, index)
}

func (a Array) SetIndex(index Int, value Variant) {
	Global.Array.SetIndex(a, index, value)
}

func (a Array) isArray() {}

func (a Array) end() { pointers.End(a) }

func (a Array) Free() {
	if ptr, ok := pointers.End(a); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Array(callframe.Arg(frame, ptr).Uintptr())
		frame.Free()
	}
}

func (a Array) Iter() iter.Seq2[Int, Variant] {
	return func(yield func(Int, Variant) bool) {
		for i := Int(0); i < a.Size(); i++ {
			if !yield(i, a.Index(i)) {
				break
			}
		}
	}
}

func NewArray() Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	Global.typeset.creation.Array[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Array]([1]uintptr{raw})
}
