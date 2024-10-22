package LightmapProbe

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[LightmapProbe] represents the position of a single manually placed probe for dynamic object lighting with [LightmapGI]. Lightmap probes affect the lighting of [GeometryInstance3D]-derived nodes that have their [member GeometryInstance3D.gi_mode] set to [constant GeometryInstance3D.GI_MODE_DYNAMIC].
Typically, [LightmapGI] probes are placed automatically by setting [member LightmapGI.generate_probes_subdiv] to a value other than [constant LightmapGI.GENERATE_PROBES_DISABLED]. By creating [LightmapProbe] nodes before baking lightmaps, you can add more probes in specific areas for greater detail, or disable automatic generation and rely only on manually placed probes instead.
[b]Note:[/b] [LightmapProbe] nodes that are placed after baking lightmaps are ignored by dynamic objects. You must bake lightmaps again after creating or modifying [LightmapProbe]s for the probes to be effective.

*/
type Go [1]classdb.LightmapProbe
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.LightmapProbe
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("LightmapProbe"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsLightmapProbe() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLightmapProbe() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("LightmapProbe", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
