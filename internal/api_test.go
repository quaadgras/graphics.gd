//go:build !generate

package gd_test

import (
	"testing"
	"unsafe"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioEffectInstance"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
)

func TestMain(m *testing.M) {
	classdb.Register[Converter]()
	classdb.Register[CustomConverterObject]()
	classdb.Register[CustomStringSignals]()
	classdb.Register[CustomSignal]()

	startup.LoadingScene()
	m.Run()
	if tree, ok := Object.As[SceneTree.Instance](Engine.GetMainLoop()); ok {
		tree.Quit()
	}
}

func TestGetGodotVersion(t *testing.T) {
	if gdextension.Host.Version.Major() != 4 {
		t.Fail()
	}
	if gdextension.Host.Version.Major() < 3 {
		t.Fail()
	}
	if gdextension.Host.Version.String() == (gdextension.String{}) {
		t.Fail()
	}
}

func TestUtilities(t *testing.T) {
	id := Resource.AllocateID()
	if id != Resource.AllocateID()-1 {
		t.Fatal("Resource.AllocateID did not return the expected value")
	}
}

func TestNativeStructSize(t *testing.T) {
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
		if size := gdextension.Host.Memory.Sizeof(pointers.Get(gd.NewStringName(name))); uintptr(size) != expectation {
			t.Fatalf("Our size of %v is %v, but Godot's is %v", name, expectation, size)
		}
	}
}

func TestLog(t *testing.T) {
	gdextension.Host.Log.Error("This is a test error message", "go", "gd_test.TestLog", "gd_test.go", 42, true)
	gdextension.Host.Log.Warning("This is a test warning message", "go", "gd_test.TestLog", "gd_test.go", 43, true)
}
