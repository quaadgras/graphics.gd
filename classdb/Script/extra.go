package Script

import (
	"graphics.gd/variant/Object"
)

// Get returns the object's Script instance, or false if no script is attached.
func Get(obj Object.Any) (Instance, bool) { //gd:Object.get_script
	script, ok := obj.AsObject()[0].GetScript().Interface().(Instance)
	return script, ok
}
