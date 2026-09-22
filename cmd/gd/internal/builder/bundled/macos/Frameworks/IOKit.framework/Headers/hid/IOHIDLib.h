// IOKit HID declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_IOKIT_IOHIDLIB_H
#define GD_IOKIT_IOHIDLIB_H

#include <IOKit/IOKitLib.h>
#include <IOKit/hid/IOHIDKeys.h>
#include <IOKit/hid/IOHIDUsageTables.h>

CF_EXTERN_C_BEGIN

typedef struct CF_BRIDGED_TYPE(id) __IOHIDManager *IOHIDManagerRef;
typedef struct CF_BRIDGED_TYPE(id) __IOHIDDevice *IOHIDDeviceRef;
typedef struct CF_BRIDGED_TYPE(id) __IOHIDElement *IOHIDElementRef;
typedef struct CF_BRIDGED_TYPE(id) __IOHIDValue *IOHIDValueRef;
typedef uint32_t IOHIDElementCookie;

enum {
	kIOHIDOptionsTypeNone = 0x00,
	kIOHIDOptionsTypeSeizeDevice = 0x01,
};
typedef enum {
	kIOHIDReportTypeInput = 0,
	kIOHIDReportTypeOutput,
	kIOHIDReportTypeFeature,
	kIOHIDReportTypeCount,
} IOHIDReportType;
typedef enum {
	kIOHIDElementTypeInput_Misc = 1,
	kIOHIDElementTypeInput_Button = 2,
	kIOHIDElementTypeInput_Axis = 3,
	kIOHIDElementTypeInput_ScanCodes = 4,
	kIOHIDElementTypeInput_NULL = 5,
	kIOHIDElementTypeOutput = 129,
	kIOHIDElementTypeFeature = 257,
	kIOHIDElementTypeCollection = 513,
} IOHIDElementType;

typedef void (*IOHIDCallback)(void *context, IOReturn result, void *sender);
typedef void (*IOHIDReportCallback)(void *context, IOReturn result, void *sender, IOHIDReportType type, uint32_t reportID, uint8_t *report, CFIndex reportLength);
typedef void (*IOHIDValueCallback)(void *context, IOReturn result, void *sender, IOHIDValueRef value);
typedef void (*IOHIDDeviceCallback)(void *context, IOReturn result, void *sender, IOHIDDeviceRef device);

// IOHIDManager
CF_EXPORT CFTypeID IOHIDManagerGetTypeID(void);
CF_EXPORT IOHIDManagerRef IOHIDManagerCreate(CFAllocatorRef allocator, IOOptionBits options) CF_RETURNS_RETAINED;
CF_EXPORT IOReturn IOHIDManagerOpen(IOHIDManagerRef manager, IOOptionBits options);
CF_EXPORT IOReturn IOHIDManagerClose(IOHIDManagerRef manager, IOOptionBits options);
CF_EXPORT CFSetRef IOHIDManagerCopyDevices(IOHIDManagerRef manager) CF_RETURNS_RETAINED;
CF_EXPORT void IOHIDManagerSetDeviceMatching(IOHIDManagerRef manager, CFDictionaryRef matching);
CF_EXPORT void IOHIDManagerSetDeviceMatchingMultiple(IOHIDManagerRef manager, CFArrayRef multiple);
CF_EXPORT void IOHIDManagerRegisterDeviceMatchingCallback(IOHIDManagerRef manager, IOHIDDeviceCallback callback, void *context);
CF_EXPORT void IOHIDManagerRegisterDeviceRemovalCallback(IOHIDManagerRef manager, IOHIDDeviceCallback callback, void *context);
CF_EXPORT void IOHIDManagerRegisterInputValueCallback(IOHIDManagerRef manager, IOHIDValueCallback callback, void *context);
CF_EXPORT void IOHIDManagerScheduleWithRunLoop(IOHIDManagerRef manager, CFRunLoopRef runLoop, CFStringRef runLoopMode);
CF_EXPORT void IOHIDManagerUnscheduleFromRunLoop(IOHIDManagerRef manager, CFRunLoopRef runLoop, CFStringRef runLoopMode);

// IOHIDDevice
CF_EXPORT CFTypeID IOHIDDeviceGetTypeID(void);
CF_EXPORT IOHIDDeviceRef IOHIDDeviceCreate(CFAllocatorRef allocator, io_service_t service) CF_RETURNS_RETAINED;
CF_EXPORT io_service_t IOHIDDeviceGetService(IOHIDDeviceRef device);
CF_EXPORT IOReturn IOHIDDeviceOpen(IOHIDDeviceRef device, IOOptionBits options);
CF_EXPORT IOReturn IOHIDDeviceClose(IOHIDDeviceRef device, IOOptionBits options);
CF_EXPORT Boolean IOHIDDeviceConformsTo(IOHIDDeviceRef device, uint32_t usagePage, uint32_t usage);
CF_EXPORT CFTypeRef IOHIDDeviceGetProperty(IOHIDDeviceRef device, CFStringRef key);
CF_EXPORT Boolean IOHIDDeviceSetProperty(IOHIDDeviceRef device, CFStringRef key, CFTypeRef property);
CF_EXPORT CFArrayRef IOHIDDeviceCopyMatchingElements(IOHIDDeviceRef device, CFDictionaryRef matching, IOOptionBits options) CF_RETURNS_RETAINED;
CF_EXPORT void IOHIDDeviceScheduleWithRunLoop(IOHIDDeviceRef device, CFRunLoopRef runLoop, CFStringRef runLoopMode);
CF_EXPORT void IOHIDDeviceUnscheduleFromRunLoop(IOHIDDeviceRef device, CFRunLoopRef runLoop, CFStringRef runLoopMode);
CF_EXPORT void IOHIDDeviceRegisterRemovalCallback(IOHIDDeviceRef device, IOHIDCallback callback, void *context);
CF_EXPORT void IOHIDDeviceRegisterInputValueCallback(IOHIDDeviceRef device, IOHIDValueCallback callback, void *context);
CF_EXPORT void IOHIDDeviceRegisterInputReportCallback(IOHIDDeviceRef device, uint8_t *report, CFIndex reportLength, IOHIDReportCallback callback, void *context);
CF_EXPORT IOReturn IOHIDDeviceSetReport(IOHIDDeviceRef device, IOHIDReportType reportType, CFIndex reportID, const uint8_t *report, CFIndex reportLength);
CF_EXPORT IOReturn IOHIDDeviceGetReport(IOHIDDeviceRef device, IOHIDReportType reportType, CFIndex reportID, uint8_t *report, CFIndex *pReportLength);
CF_EXPORT IOReturn IOHIDDeviceGetValue(IOHIDDeviceRef device, IOHIDElementRef element, IOHIDValueRef *pValue);
CF_EXPORT IOReturn IOHIDDeviceSetValue(IOHIDDeviceRef device, IOHIDElementRef element, IOHIDValueRef value);

// IOHIDElement
CF_EXPORT CFTypeID IOHIDElementGetTypeID(void);
CF_EXPORT IOHIDDeviceRef IOHIDElementGetDevice(IOHIDElementRef element);
CF_EXPORT IOHIDElementRef IOHIDElementGetParent(IOHIDElementRef element);
CF_EXPORT CFArrayRef IOHIDElementGetChildren(IOHIDElementRef element);
CF_EXPORT IOHIDElementCookie IOHIDElementGetCookie(IOHIDElementRef element);
CF_EXPORT IOHIDElementType IOHIDElementGetType(IOHIDElementRef element);
CF_EXPORT uint32_t IOHIDElementGetUsagePage(IOHIDElementRef element);
CF_EXPORT uint32_t IOHIDElementGetUsage(IOHIDElementRef element);
CF_EXPORT uint32_t IOHIDElementGetReportID(IOHIDElementRef element);
CF_EXPORT uint32_t IOHIDElementGetReportSize(IOHIDElementRef element);
CF_EXPORT uint32_t IOHIDElementGetReportCount(IOHIDElementRef element);
CF_EXPORT CFIndex IOHIDElementGetLogicalMin(IOHIDElementRef element);
CF_EXPORT CFIndex IOHIDElementGetLogicalMax(IOHIDElementRef element);
CF_EXPORT CFIndex IOHIDElementGetPhysicalMin(IOHIDElementRef element);
CF_EXPORT CFIndex IOHIDElementGetPhysicalMax(IOHIDElementRef element);
CF_EXPORT Boolean IOHIDElementIsRelative(IOHIDElementRef element);
CF_EXPORT Boolean IOHIDElementHasNullState(IOHIDElementRef element);

// IOHIDValue
CF_EXPORT CFTypeID IOHIDValueGetTypeID(void);
CF_EXPORT IOHIDElementRef IOHIDValueGetElement(IOHIDValueRef value);
CF_EXPORT uint64_t IOHIDValueGetTimeStamp(IOHIDValueRef value);
CF_EXPORT CFIndex IOHIDValueGetLength(IOHIDValueRef value);
CF_EXPORT const uint8_t *IOHIDValueGetBytePtr(IOHIDValueRef value);
CF_EXPORT CFIndex IOHIDValueGetIntegerValue(IOHIDValueRef value);
CF_EXPORT double_t IOHIDValueGetScaledValue(IOHIDValueRef value, uint32_t type);

CF_EXTERN_C_END

#endif
