//go:build !generate

package gd

import (
	"structs"
	"sync"
	"unsafe"

	"graphics.gd/internal/callframe"
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
	FromType [TypeMax]func(ret callframe.Ptr[[3]uintptr], arg uintptr)
	IntoType [TypeMax]func(ret uintptr, arg callframe.Ptr[[3]uintptr])
}

type ObjectID uint64

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

func (Godot API) NewFrame() (frame callFrame) {
	var ok bool
	frame.ptr, ok = callFrames.Get().(unsafe.Pointer)
	if !ok {
		frame = callFrame{
			ptr: Godot.Memory.Allocate(unsafe.Sizeof(unsafeCallFrame{})),
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

func (frame callFrame) Free() { callFrames.Put(frame.ptr) }

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
	return mmm.Make[API, PackedByteArray](ctx, self.API, ret)
}*/

type UnsafeArgs unsafe.Pointer

func UnsafeGet[T any](frame UnsafeArgs, index int) T {
	return *unsafe.Slice((**T)(unsafe.Pointer(frame)), index+1)[index]
}

type UnsafeBack unsafe.Pointer

func UnsafeSet[T any](frame UnsafeBack, value T) {
	*(*T)(unsafe.Pointer(frame)) = value
}

type AudioFrame struct {
	_ structs.HostLayout

	Left, Right float32
}

type PhysicsServer2DExtensionMotionResult struct {
	_ structs.HostLayout

	Travel, Remainder, CollisionPoint, CollisionNormal, ColliderVelocity Vector2
	CollisionDepth, CollisionSafeFraction, CollisionUnsafeFraction       float32
	CollisionLocalShape                                                  int32
	ColliderID                                                           uint64
	ColliderRID                                                          RID
	ColliderShape                                                        int32
}

type PhysicsServer2DExtensionRayResult struct {
	_ structs.HostLayout

	Position, Normal Vector2
	RID              RID
	ColliderID       uint64
	Collider         *Object
	Shape            int32
}

type PhysicsServer2DExtensionShapeRestInfo struct {
	_ structs.HostLayout

	Point, Normal, LinearVelocity Vector2
	RID                           RID
	ColliderID                    uint64
	Shape                         int32
}

type PhysicsServer2DExtensionShapeResult struct {
	_ structs.HostLayout

	RID        RID
	ColliderID uint64
	Collider   *Object
	Shape      int32
}

type PhysicsServer3DExtensionRayResult struct {
	_ structs.HostLayout

	Position, Normal Vector3
	RID              RID
	ColliderID       uint64
	Collider         uintptr
	Shape            int32
}
type PhysicsServer3DExtensionShapeResult struct {
	_ structs.HostLayout

	RID        RID
	ColliderID uint64
	Collider   uintptr
	Shape      int32
}
type PhysicsServer3DExtensionShapeRestInfo struct {
	_ structs.HostLayout

	Point, Normal  Vector3
	RID            RID
	ColliderID     uint64
	Shape          int32
	LinearVelocity Vector3
}
type PhysicsServer3DExtensionMotionCollision struct {
	_ structs.HostLayout

	Position, Normal        Vector3
	ColliderVelocity        Vector3
	ColliderAngularVelocity Vector3
	Depth                   float32
	LocalShape              int32
	ColliderID              ObjectID
	Collider                RID
	ColliderShape           int32
}
type PhysicsServer3DExtensionMotionResult struct {
	_ structs.HostLayout

	Travel                  Vector3
	Remainder               Vector3
	CollisionDepth          float32
	CollisionSafeFraction   float32
	CollisionUnsafeFraction float32
	Collisions              [32]PhysicsServer3DExtensionMotionCollision
	CollisionCount          int32
}
type ScriptLanguageExtensionProfilingInfo struct {
	_ structs.HostLayout

	Signature    uintptr // StringName.Pointer()
	CallCount    uint64
	TotalTime    uint64
	SelfTime     uint64
	InternalTime uint64
}
type Glyph struct {
	_ structs.HostLayout

	Start    int32
	End      int32
	Count    uint8
	Repeat   uint8
	Flags    uint16
	Xoffset  float32
	Yoffset  float32
	Advance  float32
	FontRID  RID
	FontSize int32
	Index    int32
}
type CaretInfo struct {
	_ structs.HostLayout

	LeadingCaret, TrailingCaret Rect2

	LeadingDirection, TrailingDirection uint32
}
