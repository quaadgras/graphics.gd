// ApplicationServices (CoreGraphics display, event and window services)
// declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_APPLICATIONSERVICES_H
#define GD_APPLICATIONSERVICES_H

#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <CoreText/CoreText.h>
#include <CoreServices/CoreServices.h>
#include <mach/boolean.h>
#include <dispatch/dispatch.h>
#include <IOKit/IOKitLib.h>
#include <Security/Security.h>

CF_EXTERN_C_BEGIN

// Displays

typedef uint32_t CGDirectDisplayID;
typedef uint32_t CGDisplayCount;
typedef uint32_t CGWindowID;
typedef CGError CGDisplayErr;
typedef struct CF_BRIDGED_TYPE(id) CGDisplayMode *CGDisplayModeRef;
#define kCGNullDirectDisplay ((CGDirectDisplayID)0)
#define kCGNullWindowID ((CGWindowID)0)
#define kCGDirectMainDisplay CGMainDisplayID()

typedef CF_OPTIONS(uint32_t, CGDisplayChangeSummaryFlags) {
	kCGDisplayBeginConfigurationFlag = (1 << 0),
	kCGDisplayMovedFlag = (1 << 1),
	kCGDisplaySetMainFlag = (1 << 2),
	kCGDisplaySetModeFlag = (1 << 3),
	kCGDisplayAddFlag = (1 << 4),
	kCGDisplayRemoveFlag = (1 << 5),
	kCGDisplayEnabledFlag = (1 << 8),
	kCGDisplayDisabledFlag = (1 << 9),
	kCGDisplayMirrorFlag = (1 << 10),
	kCGDisplayUnMirrorFlag = (1 << 11),
	kCGDisplayDesktopShapeChangedFlag = (1 << 12),
};
typedef void (*CGDisplayReconfigurationCallBack)(CGDirectDisplayID display, CGDisplayChangeSummaryFlags flags, void *userInfo);

CF_EXPORT CGDirectDisplayID CGMainDisplayID(void);
CF_EXPORT CGError CGGetActiveDisplayList(uint32_t maxDisplays, CGDirectDisplayID *activeDisplays, uint32_t *displayCount);
CF_EXPORT CGError CGGetOnlineDisplayList(uint32_t maxDisplays, CGDirectDisplayID *onlineDisplays, uint32_t *displayCount);
CF_EXPORT CGRect CGDisplayBounds(CGDirectDisplayID display);
CF_EXPORT size_t CGDisplayPixelsWide(CGDirectDisplayID display);
CF_EXPORT size_t CGDisplayPixelsHigh(CGDirectDisplayID display);
CF_EXPORT CGSize CGDisplayScreenSize(CGDirectDisplayID display);
CF_EXPORT boolean_t CGDisplayIsMain(CGDirectDisplayID display);
CF_EXPORT boolean_t CGDisplayIsBuiltin(CGDirectDisplayID display);
CF_EXPORT uint32_t CGDisplayVendorNumber(CGDirectDisplayID display);
CF_EXPORT uint32_t CGDisplayModelNumber(CGDirectDisplayID display);
CF_EXPORT CGDisplayModeRef CGDisplayCopyDisplayMode(CGDirectDisplayID display) CF_RETURNS_RETAINED;
CF_EXPORT CFArrayRef CGDisplayCopyAllDisplayModes(CGDirectDisplayID display, CFDictionaryRef options) CF_RETURNS_RETAINED;
CF_EXPORT void CGDisplayModeRelease(CGDisplayModeRef mode);
CF_EXPORT double CGDisplayModeGetRefreshRate(CGDisplayModeRef mode);
CF_EXPORT size_t CGDisplayModeGetWidth(CGDisplayModeRef mode);
CF_EXPORT size_t CGDisplayModeGetHeight(CGDisplayModeRef mode);
CF_EXPORT size_t CGDisplayModeGetPixelWidth(CGDisplayModeRef mode);
CF_EXPORT size_t CGDisplayModeGetPixelHeight(CGDisplayModeRef mode);
CF_EXPORT CGError CGDisplayRegisterReconfigurationCallback(CGDisplayReconfigurationCallBack callback, void *userInfo);
CF_EXPORT CGError CGDisplayRemoveReconfigurationCallback(CGDisplayReconfigurationCallBack callback, void *userInfo);
CF_EXPORT CGError CGDisplayHideCursor(CGDirectDisplayID display);
CF_EXPORT CGError CGDisplayShowCursor(CGDirectDisplayID display);
CF_EXPORT CGError CGDisplayMoveCursorToPoint(CGDirectDisplayID display, CGPoint point);
CF_EXPORT CGError CGWarpMouseCursorPosition(CGPoint newCursorPosition);
CF_EXPORT CGError CGAssociateMouseAndMouseCursorPosition(boolean_t connected);
CF_EXPORT CGError CGDisplayCapture(CGDirectDisplayID display);
CF_EXPORT CGError CGDisplayRelease(CGDirectDisplayID display);

// Events

typedef struct CF_BRIDGED_TYPE(id) __CGEvent *CGEventRef;
typedef struct CF_BRIDGED_TYPE(id) __CGEventSource *CGEventSourceRef;
typedef CF_ENUM(uint32_t, CGEventType) {
	kCGEventNull = 0,
	kCGEventLeftMouseDown = 1,
	kCGEventLeftMouseUp = 2,
	kCGEventRightMouseDown = 3,
	kCGEventRightMouseUp = 4,
	kCGEventMouseMoved = 5,
	kCGEventLeftMouseDragged = 6,
	kCGEventRightMouseDragged = 7,
	kCGEventKeyDown = 10,
	kCGEventKeyUp = 11,
	kCGEventFlagsChanged = 12,
	kCGEventScrollWheel = 22,
	kCGEventTabletPointer = 23,
	kCGEventTabletProximity = 24,
	kCGEventOtherMouseDown = 25,
	kCGEventOtherMouseUp = 26,
	kCGEventOtherMouseDragged = 27,
};
typedef CF_ENUM(int32_t, CGEventSourceStateID) {
	kCGEventSourceStatePrivate = -1,
	kCGEventSourceStateCombinedSessionState = 0,
	kCGEventSourceStateHIDSystemState = 1,
};
typedef uint16_t CGKeyCode;
typedef CF_OPTIONS(uint64_t, CGEventFlags) {
	kCGEventFlagMaskAlphaShift = 0x00010000,
	kCGEventFlagMaskShift = 0x00020000,
	kCGEventFlagMaskControl = 0x00040000,
	kCGEventFlagMaskAlternate = 0x00080000,
	kCGEventFlagMaskCommand = 0x00100000,
	kCGEventFlagMaskHelp = 0x00400000,
	kCGEventFlagMaskSecondaryFn = 0x00800000,
	kCGEventFlagMaskNumericPad = 0x00200000,
	kCGEventFlagMaskNonCoalesced = 0x00000100,
};
CF_EXPORT CGEventRef CGEventCreate(CGEventSourceRef source) CF_RETURNS_RETAINED;
CF_EXPORT CGPoint CGEventGetLocation(CGEventRef event);
CF_EXPORT CGEventType CGEventGetType(CGEventRef event);
CF_EXPORT CGEventFlags CGEventGetFlags(CGEventRef event);
CF_EXPORT CGEventSourceRef CGEventSourceCreate(CGEventSourceStateID stateID) CF_RETURNS_RETAINED;
CF_EXPORT bool CGEventSourceKeyState(CGEventSourceStateID stateID, CGKeyCode key);
CF_EXPORT bool CGEventSourceButtonState(CGEventSourceStateID stateID, uint32_t button);
CF_EXPORT void CGEventSourceSetLocalEventsSuppressionInterval(CGEventSourceRef source, CFTimeInterval seconds);
CF_EXPORT CFTimeInterval CGEventSourceGetLocalEventsSuppressionInterval(CGEventSourceRef source);

// Windows

typedef CF_OPTIONS(uint32_t, CGWindowListOption) {
	kCGWindowListOptionAll = 0,
	kCGWindowListOptionOnScreenOnly = (1 << 0),
	kCGWindowListOptionOnScreenAboveWindow = (1 << 1),
	kCGWindowListOptionOnScreenBelowWindow = (1 << 2),
	kCGWindowListOptionIncludingWindow = (1 << 3),
	kCGWindowListExcludeDesktopElements = (1 << 4),
};
typedef CF_OPTIONS(uint32_t, CGWindowImageOption) {
	kCGWindowImageDefault = 0,
	kCGWindowImageBoundsIgnoreFraming = (1 << 0),
	kCGWindowImageShouldBeOpaque = (1 << 1),
	kCGWindowImageOnlyShadows = (1 << 2),
	kCGWindowImageBestResolution = (1 << 3),
	kCGWindowImageNominalResolution = (1 << 4),
};
CF_EXPORT CFArrayRef CGWindowListCreate(CGWindowListOption option, CGWindowID relativeToWindow) CF_RETURNS_RETAINED;
CF_EXPORT CFArrayRef CGWindowListCopyWindowInfo(CGWindowListOption option, CGWindowID relativeToWindow) CF_RETURNS_RETAINED;
CF_EXPORT CGImageRef CGWindowListCreateImage(CGRect screenBounds, CGWindowListOption listOption, CGWindowID windowID, CGWindowImageOption imageOption) CF_RETURNS_RETAINED;
CF_EXPORT CGImageRef CGWindowListCreateImageFromArray(CGRect screenBounds, CFArrayRef windowArray, CGWindowImageOption imageOption) CF_RETURNS_RETAINED;
CF_EXPORT bool CGPreflightScreenCaptureAccess(void);
CF_EXPORT bool CGRequestScreenCaptureAccess(void);

// Bitmap contexts & images

typedef CF_ENUM(uint32_t, CGImageAlphaInfo) {
	kCGImageAlphaNone,
	kCGImageAlphaPremultipliedLast,
	kCGImageAlphaPremultipliedFirst,
	kCGImageAlphaLast,
	kCGImageAlphaFirst,
	kCGImageAlphaNoneSkipLast,
	kCGImageAlphaNoneSkipFirst,
	kCGImageAlphaOnly,
};
typedef CF_OPTIONS(uint32_t, CGBitmapInfo) {
	kCGBitmapAlphaInfoMask = 0x1F,
	kCGBitmapFloatInfoMask = 0xF00,
	kCGBitmapFloatComponents = (1 << 8),
	kCGBitmapByteOrderMask = 0x7000,
	kCGBitmapByteOrderDefault = (0 << 12),
	kCGBitmapByteOrder16Little = (1 << 12),
	kCGBitmapByteOrder32Little = (2 << 12),
	kCGBitmapByteOrder16Big = (3 << 12),
	kCGBitmapByteOrder32Big = (4 << 12),
};
typedef CF_ENUM(int32_t, CGInterpolationQuality) {
	kCGInterpolationDefault = 0,
	kCGInterpolationNone = 1,
	kCGInterpolationLow = 2,
	kCGInterpolationMedium = 4,
	kCGInterpolationHigh = 3,
};
typedef struct CF_BRIDGED_TYPE(id) CGDataProvider *CGDataProviderRef;
CF_EXPORT CGContextRef CGBitmapContextCreate(void *data, size_t width, size_t height, size_t bitsPerComponent, size_t bytesPerRow, CGColorSpaceRef space, uint32_t bitmapInfo) CF_RETURNS_RETAINED;
CF_EXPORT CGImageRef CGBitmapContextCreateImage(CGContextRef context) CF_RETURNS_RETAINED;
CF_EXPORT void *CGBitmapContextGetData(CGContextRef context);
CF_EXPORT void CGContextRelease(CGContextRef c);
CF_EXPORT void CGContextDrawImage(CGContextRef c, CGRect rect, CGImageRef image);
CF_EXPORT void CGContextSetInterpolationQuality(CGContextRef c, CGInterpolationQuality quality);
CF_EXPORT void CGContextSetBlendMode(CGContextRef c, int32_t mode);
CF_EXPORT void CGContextScaleCTM(CGContextRef c, CGFloat sx, CGFloat sy);
CF_EXPORT void CGContextTranslateCTM(CGContextRef c, CGFloat tx, CGFloat ty);
CF_EXPORT size_t CGImageGetWidth(CGImageRef image);
CF_EXPORT size_t CGImageGetHeight(CGImageRef image);
CF_EXPORT size_t CGImageGetBitsPerComponent(CGImageRef image);
CF_EXPORT size_t CGImageGetBitsPerPixel(CGImageRef image);
CF_EXPORT size_t CGImageGetBytesPerRow(CGImageRef image);
CF_EXPORT CGBitmapInfo CGImageGetBitmapInfo(CGImageRef image);
CF_EXPORT CGImageAlphaInfo CGImageGetAlphaInfo(CGImageRef image);
CF_EXPORT CGDataProviderRef CGImageGetDataProvider(CGImageRef image);
CF_EXPORT CFDataRef CGDataProviderCopyData(CGDataProviderRef provider) CF_RETURNS_RETAINED;
CF_EXPORT CGImageRef CGImageRetain(CGImageRef image);
CF_EXPORT void CGImageRelease(CGImageRef image);
CF_EXPORT CGRect CGRectApplyAffineTransform(CGRect rect, CGAffineTransform t);
CF_EXPORT CGRect CGRectIntersection(CGRect r1, CGRect r2);
CF_EXPORT CGRect CGRectUnion(CGRect r1, CGRect r2);

// Processes (HIServices)
typedef UInt32 ProcessApplicationTransformState;
enum {
	kNoProcess = 0,
	kSystemProcess = 1,
	kCurrentProcess = 2,
};
enum {
	kProcessTransformToForegroundApplication = 1,
	kProcessTransformToBackgroundApplication = 2,
	kProcessTransformToUIElementApplication = 4,
};
CF_EXPORT OSStatus TransformProcessType(const ProcessSerialNumber *psn, ProcessApplicationTransformState transformState);
CF_EXPORT OSErr GetCurrentProcess(ProcessSerialNumber *PSN);

CF_EXTERN_C_END

#endif
