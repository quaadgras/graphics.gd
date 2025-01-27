// Package LightmapProbe provides methods for working with LightmapProbe object instances.
package LightmapProbe

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

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

/*
[LightmapProbe] represents the position of a single manually placed probe for dynamic object lighting with [LightmapGI]. Lightmap probes affect the lighting of [GeometryInstance3D]-derived nodes that have their [member GeometryInstance3D.gi_mode] set to [constant GeometryInstance3D.GI_MODE_DYNAMIC].
Typically, [LightmapGI] probes are placed automatically by setting [member LightmapGI.generate_probes_subdiv] to a value other than [constant LightmapGI.GENERATE_PROBES_DISABLED]. By creating [LightmapProbe] nodes before baking lightmaps, you can add more probes in specific areas for greater detail, or disable automatic generation and rely only on manually placed probes instead.
[b]Note:[/b] [LightmapProbe] nodes that are placed after baking lightmaps are ignored by dynamic objects. You must bake lightmaps again after creating or modifying [LightmapProbe]s for the probes to be effective.
*/
type Instance [1]gdclass.LightmapProbe

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLightmapProbe() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.LightmapProbe

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LightmapProbe"))
	casted := Instance{*(*gdclass.LightmapProbe)(unsafe.Pointer(&object))}
	return casted
}

func (self class) AsLightmapProbe() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLightmapProbe() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("LightmapProbe", func(ptr gd.Object) any {
		return [1]gdclass.LightmapProbe{*(*gdclass.LightmapProbe)(unsafe.Pointer(&ptr))}
	})
}
