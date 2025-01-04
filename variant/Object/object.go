package Object

import (
	"reflect"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
)

type ID int64

type Instance [1]gdclass.Object

func New() Instance {
	return Instance{gd.Global.ClassDB.ConstructObject(gd.NewStringName("Object"))}
}

func (obj Instance) AsObject() gd.Object             { return obj[0] }
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

func (obj Instance) Virtual(name string) reflect.Value { return obj[0].Virtual(name) }

func (obj Instance) ClassName() string {
	return obj[0].GetClass().String()
}
