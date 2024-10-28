package ShaderGlobalsOverride

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Similar to how a [WorldEnvironment] node can be used to override the environment while a specific scene is loaded, [ShaderGlobalsOverride] can be used to override global shader parameters temporarily. Once the node is removed, the project-wide values for the global shader parameters are restored. See the [RenderingServer] [code]global_shader_parameter_*[/code] methods for more information.
[b]Note:[/b] Only one [ShaderGlobalsOverride] can be used per scene. If there is more than one [ShaderGlobalsOverride] node in the scene tree, only the first node (in tree order) will be taken into account.
[b]Note:[/b] All [ShaderGlobalsOverride] nodes are made part of a [code]"shader_overrides_group"[/code] group when they are added to the scene tree. The currently active [ShaderGlobalsOverride] node also has a [code]"shader_overrides_group_active"[/code] group added to it. You can use this to check which [ShaderGlobalsOverride] node is currently active.

*/
type Go [1]classdb.ShaderGlobalsOverride
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ShaderGlobalsOverride
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ShaderGlobalsOverride"))
	return Go{classdb.ShaderGlobalsOverride(object)}
}

func (self class) AsShaderGlobalsOverride() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShaderGlobalsOverride() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("ShaderGlobalsOverride", func(ptr gd.Object) any { return classdb.ShaderGlobalsOverride(ptr) })}
