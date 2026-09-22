graphics.gd's iOS SDK
=====================

Enough of an iOS SDK for zig to compile and link Go extensions along with
the Godot engine itself for iOS, without Xcode. gd lays it out the way the
tools expect (see ios.go and ios_engine.go), and ../ios-conformance checks it against
the real thing.

lib/ and Frameworks/*/*.tbd
	Link stubs, naming the symbols that the system's libraries export.
	Put together by hand from the symbols that linking reports missing.

Frameworks/*/Headers and include/
	Declarations of the system's API, written from its public
	documentation and only as far as Godot, SDL and metal-cpp make use
	of it. None of it is derived from the headers of Apple's SDK: when a
	build reports something missing, declare it here from the
	documentation rather than copying it across. A declaration has to
	agree with the system on names, types and the values of constants,
	as nothing checks them at link time.

	The C library's headers are not here, gd takes those from zig
	(lib/libc/include/any-macos-any, under the APSL and BSD licenses).

Frameworks/OpenGLES.framework/Headers/{ES1,ES2,ES3,KHR}
	The OpenGL ES API as published by Khronos under the MIT license,
	taken from platform/frameworks/native/opengl/include of the Android
	Open Source Project at 4f463a6b1de9198963dc6aff74154a504ba3f8f6, with
	their includes of each other made relative. gl.h and glext.h are the
	names iOS knows them by.

Frameworks/UIKit.framework/Headers/UIKeyConstants.h
	Generated from the keyboard/keypad page (0x07) of the USB HID Usage
	Tables.

app.m
	The entry point of an app. Godot's is written in Swift, which zig
	cannot compile, so gd's stand-in for swift-frontend compiles this in
	its place when building an engine.
