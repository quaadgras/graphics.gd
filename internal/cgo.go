//go:build !generate

package gd

import (
	"sync"
	"unsafe"
)

// cache is responsible for keeping a local copy for the various
// function pointers that are looked up at runtime and used for
// calling methods on classes.
type cache struct {
	utility utility
	builtin builtin
	typeset typeset
	variant variant
	Methods methods
}

type variant struct {
	FromType [TypeMax]func(CallFrameBack, CallFrameArgs)
	IntoType [TypeMax]func(CallFrameBack, CallFrameArgs)
}

const maxMethodArgs = 14
const maxArgWords = 6

var callFrames sync.Pool

type CallFrameArgs uintptr
type CallFrameBack uintptr

type unsafeCallFrame struct {
	ptrs [maxMethodArgs]uintptr
	back [maxArgWords]uintptr
	data [maxMethodArgs * maxArgWords]uintptr
}

type callFrame struct {
	ptr unsafe.Pointer
}

func (frame callFrame) Get(i int) CallFrameArgs {
	unsafeFrame := (*unsafeCallFrame)(frame.ptr)
	return CallFrameArgs(unsafeFrame.ptrs[i])
}

func (frame callFrame) Args() CallFrameArgs {
	return CallFrameArgs(frame.ptr)
}

func (frame callFrame) ArgsOffset(i Int) CallFrameArgs {
	return CallFrameArgs(frame.ptr) + CallFrameArgs(i)
}

func (frame callFrame) Back() CallFrameBack {
	unsafeFrame := (*unsafeCallFrame)(frame.ptr)
	return CallFrameBack(unsafe.Pointer(&unsafeFrame.back[0]))
}

func (Godot API) NewFrame() callFrame {
	frame, ok := callFrames.Get().(callFrame)
	if !ok {
		frame = callFrame{
			ptr: Godot.Allocate(unsafe.Sizeof(unsafeCallFrame{})),
		}
		unsafeFrame := (*unsafeCallFrame)(frame.ptr)
		*unsafeFrame = unsafeCallFrame{}
		for i := 0; i < maxMethodArgs; i++ {
			unsafeFrame.ptrs[i] = uintptr(unsafe.Pointer(&unsafeFrame.data[i*maxArgWords]))
		}
	}
	unsafeFrame := (*unsafeCallFrame)(frame.ptr)
	unsafeFrame.back = [maxArgWords]uintptr{}
	return frame
}

func (frame callFrame) Free() { callFrames.Put(frame) }

func FrameSet[T any](index int, frame callFrame, value T) {
	unsafeFrame := (*unsafeCallFrame)(frame.ptr)
	*(*T)(unsafe.Pointer(&unsafeFrame.data[index*maxArgWords])) = value
}

func FrameGet[T any](frame callFrame) T {
	unsafeFrame := (*unsafeCallFrame)(frame.ptr)
	return *(*T)(unsafe.Pointer(&unsafeFrame.back))
}

/*func (self classAESContext) Update(ctx context.Context, src PackedByteArray) PackedByteArray {
	var frame = makeCallFrame()
	setFrame[[2]uintptr](0, frame, src.Pointer())
	self.API.UnsafeCall(self.API.Methods.AESContext_update, self.Pointer(), frame.Args, frame.Ret)
	ret := getFrame[[2]uintptr](frame)
	frame.Free()
	return mmm.Make[gd.API, PackedByteArray](ctx, self.API, ret)
}*/

type UnsafeArgs uintptr

func UnsafeGet[T any](frame UnsafeArgs, index int) T {
	return *unsafe.Slice((**T)(unsafe.Pointer(frame)), index+1)[index]
}

type UnsafeBack uintptr

func UnsafeSet[T any](frame UnsafeBack, value T) {
	*(*T)(unsafe.Pointer(frame)) = value
}
