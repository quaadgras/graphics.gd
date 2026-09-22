// CoreServices declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_CORESERVICES_H
#define GD_CORESERVICES_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef SInt32 CGError;
enum {
	kCGErrorSuccess = 0,
	kCGErrorFailure = 1000,
	kCGErrorIllegalArgument = 1001,
	kCGErrorInvalidConnection = 1002,
	kCGErrorInvalidContext = 1003,
	kCGErrorCannotComplete = 1004,
	kCGErrorNotImplemented = 1006,
	kCGErrorRangeCheck = 1007,
	kCGErrorTypeCheck = 1008,
	kCGErrorInvalidOperation = 1010,
	kCGErrorNoneAvailable = 1011,
};

// Launch Services
typedef CF_OPTIONS(OptionBits, LSLaunchFlags) {
	kLSLaunchDefaults = 0x00000001,
	kLSLaunchAndPrint = 0x00000002,
	kLSLaunchAndDisplayErrors = 0x00000040,
	kLSLaunchDontAddToRecents = 0x00000100,
	kLSLaunchDontSwitch = 0x00000200,
	kLSLaunchAsync = 0x00010000,
	kLSLaunchNewInstance = 0x00080000,
	kLSLaunchAndHide = 0x00100000,
	kLSLaunchAndHideOthers = 0x00200000,
};
CF_EXPORT OSStatus LSOpenCFURLRef(CFURLRef inURL, CFURLRef *outLaunchedURL);

CF_EXTERN_C_END

#endif
