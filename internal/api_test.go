//go:build !generate

package gd_test

import (
	"os"
	"testing"
	"unsafe"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioEffectInstance"
	gd "graphics.gd/internal"
	"graphics.gd/startup"
)

func TestMain(m *testing.M) {
	classdb.Register[Converter]()
	classdb.Register[CustomConverterObject]()
	startup.LoadingScene()
	os.Exit(m.Run())
}

func TestGetGodotVersion(t *testing.T) {
	version := gd.Global.GetGodotVersion()
	if version.Major != 4 {
		t.Fail()
	}
	if version.Minor < 3 {
		t.Fail()
	}
	if version.Patch < 0 {
		t.Fail()
	}
	if version.String() == "" {
		t.Fail()
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
		if size := gd.Global.GetNativeStructSize(gd.NewStringName(name)); size != expectation {
			t.Fatalf("Our size of %v is %v, but Godot's is %v", name, expectation, size)
		}
	}
}
