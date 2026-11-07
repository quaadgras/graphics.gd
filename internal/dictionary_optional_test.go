package gd_test

import (
	"strings"
	"sync"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Logger"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/PhysicsDirectSpaceState3D"
	"graphics.gd/classdb/ScriptBacktrace"
	"graphics.gd/variant/Object"
)

// ErrorCollector records the errors the engine logs.
type ErrorCollector struct {
	Logger.Extension[ErrorCollector]

	mu     sync.Mutex
	errors []string
}

func (l *ErrorCollector) LogError(function string, file string, line int, code string, rationale string, editor_notify bool, error_type int, script_backtraces []ScriptBacktrace.Instance) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.errors = append(l.errors, code+" "+rationale)
}

func (l *ErrorCollector) LogMessage(message string, error bool) {}

// OptionalKeys takes a struct whose keys are not all in the dictionary
// it is given, like an intersect_ray result read from intersect_shape.
type OptionalKeys struct {
	Object.Extension[OptionalKeys]

	got PhysicsDirectSpaceState3D.PhysicsDirectSpaceState3D_Intersection
}

func (o *OptionalKeys) Take(hit PhysicsDirectSpaceState3D.PhysicsDirectSpaceState3D_Intersection) {
	o.got = hit
}

func init() {
	classdb.Register[ErrorCollector]()
	classdb.Register[OptionalKeys]()
}

// TestDictionaryOptionalKeys checks that a struct is read from a dictionary
// that leaves out some of its fields without the engine reporting an error
// for each missing key, and that those fields are left zero.
func TestDictionaryOptionalKeys(t *testing.T) {
	const source string = `extends OptionalKeys

func send() -> void:
	take({"collider_id": 7, "shape": 3, "rid": RID()})
`
	logger := new(ErrorCollector)
	OS.AddLogger(logger.AsLogger())
	defer OS.RemoveLogger(logger.AsLogger())

	var runner = new(OptionalKeys)
	var script = GDScript.New().AsScript()
	script.SetSourceCode(source)
	script.Reload()
	Object.Instance(runner.AsObject()).SetScript(script)
	Object.Call(runner, "send")

	if runner.got.ColliderID != 7 || runner.got.Shape != 3 {
		t.Fatalf("expected collider_id 7 and shape 3, got %+v", runner.got)
	}
	if runner.got.FaceIndex != 0 || runner.got.Normal.X != 0 || runner.got.Position.X != 0 {
		t.Fatalf("expected missing fields to be zero, got %+v", runner.got)
	}
	logger.mu.Lock()
	defer logger.mu.Unlock()
	for _, err := range logger.errors {
		if strings.Contains(err, "Dictionary") {
			t.Errorf("engine reported: %s", err)
		}
	}
}
