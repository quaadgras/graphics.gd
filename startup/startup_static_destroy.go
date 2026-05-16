//go:build musl || archive

package startup

import "C"

//export go_destroy_engine
func go_destroy_engine() {
	if engine, ok := startup.(*engineAsStaticLibrary); ok {
		if engine.destroy != nil {
			destroy := engine.destroy
			engine.destroy = nil
			destroy()
		}
	}
}
