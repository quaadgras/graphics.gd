//go:build !generate

package gd_test

import (
	"testing"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioEffectInstance"
	"graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
	"graphics.gd/internal/threadcheck"

	_ "graphics.gd/startup"
)

//go:linkname keep_reachable_instances_alive graphics.gd/classdb.keep_reachable_instances_alive
func keep_reachable_instances_alive()

func init() {
	classdb.Register[Converter]()
	classdb.Register[CustomConverterObject]()
	classdb.Register[CustomStringSignals]()
	classdb.Register[CustomSignal]()
}

func TestThreadCheck(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		if !threadcheck.Main() {
			t.Fail()
		}
	})
}

func TestGetGodotVersion(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		if gdunsafe.VersionMajor() != 4 {
			t.Fail()
		}
		if gdunsafe.VersionMajor() < 3 {
			t.Fail()
		}
		if gdunsafe.Version() == (gdunsafe.String{}) {
			t.Fail()
		}
	})
}

func TestUtilities(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		id := Resource.AllocateID()
		if id != Resource.AllocateID()-1 {
			t.Fatal("Resource.AllocateID did not return the expected value")
		}
	})
}

func TestNativeStructSize(t *testing.T) {
	t.Skip()
	runOnMain(t, func(t testing.TB) {
		for name, expectation := range map[string]uintptr{
			"ObjectID":                                unsafe.Sizeof(gd.ObjectID(0)),
			"AudioFrame":                              unsafe.Sizeof(AudioEffectInstance.AudioFrame{}),
			"ScriptLanguageExtensionProfilingInfo":    unsafe.Sizeof(gd.ScriptLanguageExtensionProfilingInfo{}),
			"Glyph":                                   unsafe.Sizeof(gd.Glyph{}),
			"CaretInfo":                               unsafe.Sizeof(gd.CaretInfo{}),
			"PhysicsServer2DExtensionRayResult":       unsafe.Sizeof(gd.PhysicsServer2DExtensionRayResult{}),
			"PhysicsServer2DExtensionShapeResult":     unsafe.Sizeof(gd.PhysicsServer2DExtensionShapeResult{}),
			"PhysicsServer2DExtensionShapeRestInfo":   unsafe.Sizeof(gd.PhysicsServer2DExtensionShapeRestInfo{}),
			"PhysicsServer2DExtensionMotionResult":    unsafe.Sizeof(gd.PhysicsServer2DExtensionMotionResult{}),
			"PhysicsServer3DExtensionRayResult":       unsafe.Sizeof(gd.PhysicsServer3DExtensionRayResult{}),
			"PhysicsServer3DExtensionShapeResult":     unsafe.Sizeof(gd.PhysicsServer3DExtensionShapeResult{}),
			"PhysicsServer3DExtensionShapeRestInfo":   unsafe.Sizeof(gd.PhysicsServer3DExtensionShapeRestInfo{}),
			"PhysicsServer3DExtensionMotionCollision": unsafe.Sizeof(gd.PhysicsServer3DExtensionMotionCollision{}),
			"PhysicsServer3DExtensionMotionResult":    unsafe.Sizeof(gd.PhysicsServer3DExtensionMotionResult{}),
		} {
			if size := gdunsafe.TypeByName(name).Size(); uintptr(size) != expectation {
				t.Fatalf("Our size of %v is %v, but Godot's is %v", name, expectation, size)
			}
		}
	})
}

func TestLog(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		gdunsafe.Log(gdunsafe.LogError, "This is a test error message", "go", "gd_test.TestLog", "gd_test.go", 42, true)
		gdunsafe.Log(gdunsafe.LogWarning, "This is a test warning message", "go", "gd_test.TestLog", "gd_test.go", 43, true)
	})
}
