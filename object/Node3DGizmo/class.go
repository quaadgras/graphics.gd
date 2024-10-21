package Node3DGizmo

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This abstract class helps connect the [Node3D] scene with the editor-specific [EditorNode3DGizmo] class.
[Node3DGizmo] by itself has no exposed API, refer to [method Node3D.add_gizmo] and pass it an [EditorNode3DGizmo] instance.

*/
type Simple [1]classdb.Node3DGizmo
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Node3DGizmo
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsNode3DGizmo() Expert { return self[0].AsNode3DGizmo() }


//go:nosplit
func (self Simple) AsNode3DGizmo() Simple { return self[0].AsNode3DGizmo() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Node3DGizmo", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
