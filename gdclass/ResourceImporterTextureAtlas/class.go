package ResourceImporterTextureAtlas

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ResourceImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This imports a collection of textures from a PNG image into an [AtlasTexture] or 2D [ArrayMesh]. This can be used to save memory when importing 2D animations from spritesheets. Texture atlases are only supported in 2D rendering, not 3D. See also [ResourceImporterTexture] and [ResourceImporterLayeredTexture].
[b]Note:[/b] [ResourceImporterTextureAtlas] does not handle importing [TileSetAtlasSource], which is created using the [TileSet] editor instead.

*/
type Go [1]classdb.ResourceImporterTextureAtlas
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ResourceImporterTextureAtlas
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ResourceImporterTextureAtlas"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsResourceImporterTextureAtlas() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporterTextureAtlas() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResourceImporter() ResourceImporter.GD { return *((*ResourceImporter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporter() ResourceImporter.Go { return *((*ResourceImporter.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}
func init() {classdb.Register("ResourceImporterTextureAtlas", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
