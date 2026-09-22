Checks ../ios and ../macos (graphics.gd's iOS and macOS SDKs) against
Apple's, on a Mac with Xcode.

The headers of ../ios are written from documentation, so nothing about a
build tells a wrong constant, struct layout, selector or prototype apart
from a right one: it shows up as a bug on the device. These scripts turn
the headers into two source files that only name what they declare, to be
compiled against both SDKs.

	python3 values_source.py [ios|macos]        # writes conformance.m
	python3 declarations_source.py [ios|macos]  # writes conformance2.m (after the above)

conformance.m stores every constant, struct size and field offset. Compile
it to assembly with each SDK and compare the values:

	zig cc -target aarch64-ios.15.0 -fobjc-arc -fblocks -isystem <sdk>/usr/include \
		-iframework <sdk>/System/Library/Frameworks -S -o mine.s conformance.m
	xcrun --sdk iphoneos clang -target arm64-apple-ios15.0 -fobjc-arc -fblocks \
		-ferror-limit=0 -S -o real.s conformance.m                  # on the Mac
	python3 values_of.py mine.s > mine.txt; python3 values_of.py real.s > real.txt
	diff mine.txt real.txt

(<sdk> is what gd lays out in $GDPATH/lib/ios/sdk). An error from the Mac's
compile is a name that the real SDK does not have.

conformance2.m redeclares every C function and constant, and references
every selector. Compile it on the Mac only: a conflicting type is an error,
and -Wundeclared-selector reports selectors that nothing in the SDK declares
(`void` and `setVoid:` come from block properties and can be ignored, as
can nullability warnings, the file assumes nonnull throughout).

	xcrun --sdk iphoneos clang -target arm64-apple-ios15.0 -fobjc-arc -fblocks \
		-Wundeclared-selector -ferror-limit=0 -c -o /dev/null conformance2.m

For macOS the target is arm64-apple-macos11.0 and the Mac's compile needs
-fmodules, as Apple's CoreHaptics only compiles that way there.

Last runs: the iOS 26.5 SDK, 819 values, 331 prototypes and 1199 selectors;
the macOS 26.5 SDK, 1369 values, 608 prototypes and 1754 selectors; all in
agreement.
