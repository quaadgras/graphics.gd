//go:build !generate

package internal

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
	variant variant
	methods methods
}

type AudioFrame struct {
	Left, Right float32
}

type PhysicsServer2DExtensionMotionResult struct {
	Travel, Remainder, CollisionPoint, CollisionNormal, ColliderVelocity Vector2
	CollisionDepth, CollisionSafeFraction, CollisionUnsafeFraction       float32
	CollisionLocalShape                                                  int32
	ColliderID                                                           uint64
	ColliderRID                                                          RID
	ColliderShape                                                        int32
}

type PhysicsServer2DExtensionRayResult struct {
	Position, Normal Vector2
	RID              RID
	ColliderID       uint64
	Collider         *Object
	Shape            int32
}

type PhysicsServer2DExtensionShapeRestInfo struct {
	Point, Normal, LinearVelocity Vector2
	RID                           RID
	ColliderID                    uint64
	Shape                         int32
}

type PhysicsServer2DExtensionShapeResult struct {
	RID        RID
	ColliderID uint64
	Collider   *Object
	Shape      int32
}

type PhysicsServer3DExtensionRayResult struct {
	Position, Normal Vector3
	RID              RID
	ColliderID       uint64
	Collider         uintptr
	Shape            int32
}
type PhysicsServer3DExtensionShapeResult struct {
	RID        RID
	ColliderID uint64
	Collider   uintptr
	Shape      int32
}
type PhysicsServer3DExtensionShapeRestInfo struct {
	Point, Normal  Vector3
	RID            RID
	ColliderID     uint64
	Shape          int32
	LinearVelocity Vector3
}
type PhysicsServer3DExtensionMotionCollision struct {
	Position, Normal Vector3
	ColliderVelocity Vector3
	Depth            float32
	LocalShape       int32
	ColliderID       uint64
	Collider         RID
	ColliderShape    int32
}
type PhysicsServer3DExtensionMotionResult struct {
	Travel                  Vector3
	Remainder               Vector3
	CollisionSafeFraction   float32
	CollisionUnsafeFraction float32
	Collisions              [32]PhysicsServer3DExtensionMotionCollision
	CollisionCount          int32
}
type ScriptLanguageExtensionProfilingInfo struct {
	Signature  StringName
	Call_count uint64
	TotalTime  uint64
	SelfTime   uint64
}
type Glyph struct {
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
	LeadingCaret, TrailingCaret Rect2

	LeadingDirection, TrailingDirection int64
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

func (Godot API) newFrame() callFrame {
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

func (frame callFrame) free() { callFrames.Put(frame) }

func frameSet[T any](index int, frame callFrame, value T) {
	unsafeFrame := (*unsafeCallFrame)(frame.ptr)
	*(*T)(unsafe.Pointer(&unsafeFrame.data[index*maxArgWords])) = value
}

func frameGet[T any](frame callFrame) T {
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

type godotArgs uintptr

func godotGet[T any](frame godotArgs, index int) T {
	return unsafe.Slice((*T)(unsafe.Pointer(frame)), index+1)[index]
}

type godotBack uintptr

func godotSet[T any](frame godotBack, value T) {
	*(*T)(unsafe.Pointer(frame)) = value
}
