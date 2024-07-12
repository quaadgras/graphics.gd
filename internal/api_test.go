//go:build !generate

package gd_test

import (
	"fmt"
	"os"
	"testing"
	"unsafe"

	"grow.graphics/gd/gdextension"
	gd "grow.graphics/gd/internal"
)

var API *gd.API

var godot gd.Lifetime

func TestMain(M *testing.M) {
	var ok bool
	godot, ok = gdextension.Link()
	if ok {
		API = godot.API
		os.Exit(M.Run())
	}
}

func TestGetGodotVersion(t *testing.T) {
	version := API.GetGodotVersion()
	if version.Major != 4 {
		t.Fail()
	}
	if version.Minor < 2 {
		t.Fail()
	}
	if version.Patch < 1 {
		t.Fail()
	}
	if version.String() == "" {
		t.Fail()
	}
}

func TestNativeStructSize(t *testing.T) {
	godot := gd.NewLifetime(API)
	defer godot.End()
	for name, expectation := range map[string]uintptr{
		"ObjectID":                                unsafe.Sizeof(gd.ObjectID(0)),
		"AudioFrame":                              unsafe.Sizeof(gd.AudioFrame{}),
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
		if API.GetNativeStructSize(godot.StringName(name)) != expectation {
			t.Fatalf("Our size of %v is %v, but Godot's is %v", name, expectation, API.GetNativeStructSize(godot.StringName(name)))
		}
	}
}

func TestGetLibraryPath(t *testing.T) {
	godot := gd.NewLifetime(API)
	defer godot.End()
	fmt.Println(godot.GetLibraryPath())
}
