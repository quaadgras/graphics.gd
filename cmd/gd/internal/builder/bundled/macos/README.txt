graphics.gd's macOS SDK
=======================

Enough of a macOS SDK for zig to compile and link Go extensions along
with the Godot engine itself for macOS, without Xcode. gd lays it out the
way Xcode does one (see macos.go and macos_engine.go), on top of the iOS
SDK next to it: the frameworks the two platforms share (Foundation,
QuartzCore, Metal, AVFoundation, GameController...) are declared once,
in ../ios, and only what macOS has of its own is here.

lib/ and Frameworks/*/*.tbd
	Link stubs, naming the symbols that the system's libraries export.
	Written from the symbols that linking reports missing, each told
	apart by the framework that exports it (a Mac with Xcode has the
	real stubs to look that up in, see ../ios-conformance).

Frameworks/*/Headers
	Declarations of the system's API, written from its public
	documentation and only as far as Godot and SDL make use of it:
	AppKit, Carbon, IOKit, ForceFeedback, CoreAudio, CoreMIDI, the
	display, event and window services of CoreGraphics, Security,
	OpenGL (CGL) and UniformTypeIdentifiers. None of it is derived from
	the headers of Apple's SDK. A declaration has to agree with the
	system on names, types and the values of constants, which
	../ios-conformance checks against the real SDK: the last run found
	every one of its 1369 values, 608 prototypes and 1754 selectors in
	agreement (after it caught six that were not).

include/
	The Go build's, for the most part: gd's engine builds take the C
	library's headers from zig and only mach-o/ from here.
