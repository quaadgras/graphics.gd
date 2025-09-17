package gdjson

import (
	"reflect"

	"graphics.gd/variant/Float"
)

var Returnables = map[string]map[string]reflect.Type{
	"OS.execute":                              {"output": reflect.TypeFor[[]string]()},
	"EditorExportPlatform.ssh_run_on_remote":  {"output": reflect.TypeFor[[]string]()},
	"ResourceLoader.load_threaded_get_status": {"progress": reflect.TypeFor[Float.X]()},
}
