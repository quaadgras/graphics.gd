// IOKit power management declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_IOKIT_IOPMLIB_H
#define GD_IOKIT_IOPMLIB_H

#include <IOKit/IOTypes.h>

CF_EXTERN_C_BEGIN

typedef UInt32 IOPMAssertionID;
typedef UInt32 IOPMAssertionLevel;
enum {
	kIOPMNullAssertionID = 0,
	kIOPMAssertionLevelOff = 0,
	kIOPMAssertionLevelOn = 255,
};
#define kIOPMAssertionTypeNoDisplaySleep CFSTR("NoDisplaySleepAssertion")
#define kIOPMAssertionTypeNoIdleSleep CFSTR("NoIdleSleepAssertion")
#define kIOPMAssertPreventUserIdleDisplaySleep CFSTR("PreventUserIdleDisplaySleep")
#define kIOPMAssertPreventUserIdleSystemSleep CFSTR("PreventUserIdleSystemSleep")

CF_EXPORT IOReturn IOPMAssertionCreateWithName(CFStringRef AssertionType, IOPMAssertionLevel AssertionLevel, CFStringRef AssertionName, IOPMAssertionID *AssertionID);
CF_EXPORT IOReturn IOPMAssertionCreateWithDescription(CFStringRef AssertionType, CFStringRef Name, CFStringRef Details, CFStringRef HumanReadableReason, CFStringRef LocalizationBundlePath, CFTimeInterval Timeout, CFStringRef TimeoutAction, IOPMAssertionID *AssertionID);
CF_EXPORT IOReturn IOPMAssertionRelease(IOPMAssertionID AssertionID);

CF_EXTERN_C_END

#endif
