// CoreFoundation declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot,
// SDL and metal-cpp make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREFOUNDATION_H
#define GD_COREFOUNDATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <sys/types.h>
#include <TargetConditionals.h>

// The umbrella header is documented to bring the C standard library with it.
#include <assert.h>
#include <ctype.h>
#include <errno.h>
#include <float.h>
#include <limits.h>
#include <locale.h>
#include <math.h>
#include <setjmp.h>
#include <signal.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#if defined(__cplusplus)
#define CF_EXTERN_C_BEGIN extern "C" {
#define CF_EXTERN_C_END }
#else
#define CF_EXTERN_C_BEGIN
#define CF_EXTERN_C_END
#endif
#define CF_EXPORT extern
#define CF_INLINE static inline
#define CF_RETURNS_RETAINED __attribute__((cf_returns_retained))
#define CF_BRIDGED_TYPE(T) __attribute__((objc_bridge(T)))
#define CF_BRIDGED_MUTABLE_TYPE(T) __attribute__((objc_bridge_mutable(T)))
#define CF_ENUM(type, name) enum name : type name; enum name : type
#define CF_OPTIONS(type, name) type name; enum : type
#define CF_ASSUME_NONNULL_BEGIN _Pragma("clang assume_nonnull begin")
#define CF_ASSUME_NONNULL_END _Pragma("clang assume_nonnull end")

CF_EXTERN_C_BEGIN

typedef unsigned char Boolean;
typedef unsigned char UInt8;
typedef signed char SInt8;
typedef unsigned short UInt16;
typedef signed short SInt16;
typedef unsigned int UInt32;
typedef signed int SInt32;
typedef uint64_t UInt64;
typedef int64_t SInt64;
typedef SInt32 OSStatus;
typedef SInt16 OSErr;
typedef float Float32;
typedef double Float64;
typedef UInt16 UniChar;
typedef UInt32 UTF32Char;
typedef UInt32 FourCharCode;
typedef FourCharCode OSType;
#ifndef TRUE
#define TRUE 1
#endif
#ifndef FALSE
#define FALSE 0
#endif
enum { noErr = 0 };

typedef unsigned long CFTypeID;
typedef unsigned long CFOptionFlags;
typedef unsigned long CFHashCode;
typedef signed long CFIndex;
typedef const void *CFTypeRef;
typedef double CFTimeInterval;
typedef double CFAbsoluteTime;

typedef struct {
	CFIndex location;
	CFIndex length;
} CFRange;

CF_INLINE CFRange CFRangeMake(CFIndex loc, CFIndex len) {
	CFRange range;
	range.location = loc;
	range.length = len;
	return range;
}

typedef const struct CF_BRIDGED_TYPE(NSString) __CFString *CFStringRef;
typedef struct CF_BRIDGED_MUTABLE_TYPE(NSMutableString) __CFString *CFMutableStringRef;
typedef const struct CF_BRIDGED_TYPE(NSArray) __CFArray *CFArrayRef;
typedef struct CF_BRIDGED_MUTABLE_TYPE(NSMutableArray) __CFArray *CFMutableArrayRef;
typedef const struct CF_BRIDGED_TYPE(NSDictionary) __CFDictionary *CFDictionaryRef;
typedef struct CF_BRIDGED_MUTABLE_TYPE(NSMutableDictionary) __CFDictionary *CFMutableDictionaryRef;
typedef const struct CF_BRIDGED_TYPE(NSNumber) __CFNumber *CFNumberRef;
typedef const struct CF_BRIDGED_TYPE(NSNumber) __CFBoolean *CFBooleanRef;
typedef const struct CF_BRIDGED_TYPE(NSData) __CFData *CFDataRef;
typedef const struct CF_BRIDGED_TYPE(NSURL) __CFURL *CFURLRef;
typedef const struct CF_BRIDGED_TYPE(NSSet) __CFSet *CFSetRef;
typedef const struct CF_BRIDGED_TYPE(NSLocale) __CFLocale *CFLocaleRef;
typedef const struct CF_BRIDGED_TYPE(id) __CFAllocator *CFAllocatorRef;
typedef struct CF_BRIDGED_TYPE(id) __CFBundle *CFBundleRef;
typedef struct CF_BRIDGED_TYPE(id) __CFRunLoop *CFRunLoopRef;
typedef struct CF_BRIDGED_TYPE(id) __CFRunLoopSource *CFRunLoopSourceRef;
typedef struct CF_BRIDGED_TYPE(id) __CFUUID *CFUUIDRef;
typedef CFStringRef CFRunLoopMode;
typedef CFStringRef CFNotificationName;

typedef CF_ENUM(CFIndex, CFComparisonResult) {
	kCFCompareLessThan = -1,
	kCFCompareEqualTo = 0,
	kCFCompareGreaterThan = 1,
};

CF_EXPORT const CFAllocatorRef kCFAllocatorDefault;
CF_EXPORT const CFAllocatorRef kCFAllocatorNull;
CF_EXPORT const CFBooleanRef kCFBooleanTrue;
CF_EXPORT const CFBooleanRef kCFBooleanFalse;

CF_EXPORT CFTypeRef CFRetain(CFTypeRef cf);
CF_EXPORT void CFRelease(CFTypeRef cf);
CF_EXPORT CFTypeID CFGetTypeID(CFTypeRef cf);
CF_EXPORT Boolean CFEqual(CFTypeRef cf1, CFTypeRef cf2);
CF_EXPORT CFStringRef CFCopyDescription(CFTypeRef cf) CF_RETURNS_RETAINED;
CF_EXPORT CFAbsoluteTime CFAbsoluteTimeGetCurrent(void);

// CFString
typedef UInt32 CFStringEncoding;
typedef CF_ENUM(CFStringEncoding, CFStringBuiltInEncodings) {
	kCFStringEncodingMacRoman = 0,
	kCFStringEncodingWindowsLatin1 = 0x0500,
	kCFStringEncodingISOLatin1 = 0x0201,
	kCFStringEncodingNextStepLatin = 0x0B01,
	kCFStringEncodingASCII = 0x0600,
	kCFStringEncodingUnicode = 0x0100,
	kCFStringEncodingUTF8 = 0x08000100,
	kCFStringEncodingNonLossyASCII = 0x0BFF,
	kCFStringEncodingUTF16 = 0x0100,
	kCFStringEncodingUTF16BE = 0x10000100,
	kCFStringEncodingUTF16LE = 0x14000100,
	kCFStringEncodingUTF32 = 0x0c000100,
	kCFStringEncodingUTF32BE = 0x18000100,
	kCFStringEncodingUTF32LE = 0x1c000100,
};
#define kCFStringEncodingInvalidId (0xffffffffU)
typedef CF_OPTIONS(CFOptionFlags, CFStringCompareFlags) {
	kCFCompareCaseInsensitive = 1,
	kCFCompareBackwards = 4,
	kCFCompareAnchored = 8,
	kCFCompareNonliteral = 16,
	kCFCompareLocalized = 32,
	kCFCompareNumerically = 64,
};

CF_EXPORT CFStringRef __CFStringMakeConstantString(const char *cStr);
#define CFSTR(cStr) ((CFStringRef)__builtin___CFStringMakeConstantString("" cStr ""))

CF_EXPORT CFTypeID CFStringGetTypeID(void);
CF_EXPORT CFStringRef CFStringCreateWithCString(CFAllocatorRef alloc, const char *cStr, CFStringEncoding encoding) CF_RETURNS_RETAINED;
CF_EXPORT CFStringRef CFStringCreateWithBytes(CFAllocatorRef alloc, const UInt8 *bytes, CFIndex numBytes, CFStringEncoding encoding, Boolean isExternalRepresentation) CF_RETURNS_RETAINED;
CF_EXPORT CFStringRef CFStringCreateWithCharacters(CFAllocatorRef alloc, const UniChar *chars, CFIndex numChars) CF_RETURNS_RETAINED;
CF_EXPORT CFStringRef CFStringCreateCopy(CFAllocatorRef alloc, CFStringRef theString) CF_RETURNS_RETAINED;
CF_EXPORT CFIndex CFStringGetLength(CFStringRef theString);
CF_EXPORT UniChar CFStringGetCharacterAtIndex(CFStringRef theString, CFIndex idx);
CF_EXPORT void CFStringGetCharacters(CFStringRef theString, CFRange range, UniChar *buffer);
CF_EXPORT Boolean CFStringGetCString(CFStringRef theString, char *buffer, CFIndex bufferSize, CFStringEncoding encoding);
CF_EXPORT const char *CFStringGetCStringPtr(CFStringRef theString, CFStringEncoding encoding);
CF_EXPORT CFIndex CFStringGetBytes(CFStringRef theString, CFRange range, CFStringEncoding encoding, UInt8 lossByte, Boolean isExternalRepresentation, UInt8 *buffer, CFIndex maxBufLen, CFIndex *usedBufLen);
CF_EXPORT CFStringEncoding CFStringGetFastestEncoding(CFStringRef theString);
CF_EXPORT CFStringEncoding CFStringGetSystemEncoding(void);
CF_EXPORT CFIndex CFStringGetMaximumSizeForEncoding(CFIndex length, CFStringEncoding encoding);
CF_EXPORT CFComparisonResult CFStringCompare(CFStringRef theString1, CFStringRef theString2, CFStringCompareFlags compareOptions);

// CFNumber
typedef CF_ENUM(CFIndex, CFNumberType) {
	kCFNumberSInt8Type = 1,
	kCFNumberSInt16Type = 2,
	kCFNumberSInt32Type = 3,
	kCFNumberSInt64Type = 4,
	kCFNumberFloat32Type = 5,
	kCFNumberFloat64Type = 6,
	kCFNumberCharType = 7,
	kCFNumberShortType = 8,
	kCFNumberIntType = 9,
	kCFNumberLongType = 10,
	kCFNumberLongLongType = 11,
	kCFNumberFloatType = 12,
	kCFNumberDoubleType = 13,
	kCFNumberCFIndexType = 14,
	kCFNumberNSIntegerType = 15,
	kCFNumberCGFloatType = 16,
};
CF_EXPORT CFTypeID CFNumberGetTypeID(void);
CF_EXPORT CFNumberRef CFNumberCreate(CFAllocatorRef allocator, CFNumberType theType, const void *valuePtr) CF_RETURNS_RETAINED;
CF_EXPORT Boolean CFNumberGetValue(CFNumberRef number, CFNumberType theType, void *valuePtr);
CF_EXPORT Boolean CFBooleanGetValue(CFBooleanRef boolean);

// CFData
CF_EXPORT CFDataRef CFDataCreate(CFAllocatorRef allocator, const UInt8 *bytes, CFIndex length) CF_RETURNS_RETAINED;
CF_EXPORT CFIndex CFDataGetLength(CFDataRef theData);
CF_EXPORT const UInt8 *CFDataGetBytePtr(CFDataRef theData);

// CFArray
typedef struct {
	CFIndex version;
	const void *(*retain)(CFAllocatorRef allocator, const void *value);
	void (*release)(CFAllocatorRef allocator, const void *value);
	CFStringRef (*copyDescription)(const void *value);
	Boolean (*equal)(const void *value1, const void *value2);
} CFArrayCallBacks;
CF_EXPORT const CFArrayCallBacks kCFTypeArrayCallBacks;
CF_EXPORT CFArrayRef CFArrayCreate(CFAllocatorRef allocator, const void **values, CFIndex numValues, const CFArrayCallBacks *callBacks) CF_RETURNS_RETAINED;
CF_EXPORT CFIndex CFArrayGetCount(CFArrayRef theArray);
CF_EXPORT const void *CFArrayGetValueAtIndex(CFArrayRef theArray, CFIndex idx);

// CFDictionary
typedef struct {
	CFIndex version;
	const void *(*retain)(CFAllocatorRef allocator, const void *value);
	void (*release)(CFAllocatorRef allocator, const void *value);
	CFStringRef (*copyDescription)(const void *value);
	Boolean (*equal)(const void *value1, const void *value2);
	CFHashCode (*hash)(const void *value);
} CFDictionaryKeyCallBacks;
typedef struct {
	CFIndex version;
	const void *(*retain)(CFAllocatorRef allocator, const void *value);
	void (*release)(CFAllocatorRef allocator, const void *value);
	CFStringRef (*copyDescription)(const void *value);
	Boolean (*equal)(const void *value1, const void *value2);
} CFDictionaryValueCallBacks;
CF_EXPORT const CFDictionaryKeyCallBacks kCFTypeDictionaryKeyCallBacks;
CF_EXPORT const CFDictionaryValueCallBacks kCFTypeDictionaryValueCallBacks;
CF_EXPORT CFDictionaryRef CFDictionaryCreate(CFAllocatorRef allocator, const void **keys, const void **values, CFIndex numValues, const CFDictionaryKeyCallBacks *keyCallBacks, const CFDictionaryValueCallBacks *valueCallBacks) CF_RETURNS_RETAINED;
CF_EXPORT CFMutableDictionaryRef CFDictionaryCreateMutable(CFAllocatorRef allocator, CFIndex capacity, const CFDictionaryKeyCallBacks *keyCallBacks, const CFDictionaryValueCallBacks *valueCallBacks) CF_RETURNS_RETAINED;
CF_EXPORT CFIndex CFDictionaryGetCount(CFDictionaryRef theDict);
CF_EXPORT const void *CFDictionaryGetValue(CFDictionaryRef theDict, const void *key);
CF_EXPORT Boolean CFDictionaryGetValueIfPresent(CFDictionaryRef theDict, const void *key, const void **value);
CF_EXPORT void CFDictionaryAddValue(CFMutableDictionaryRef theDict, const void *key, const void *value);
CF_EXPORT void CFDictionarySetValue(CFMutableDictionaryRef theDict, const void *key, const void *value);

// CFURL & CFBundle
CF_EXPORT Boolean CFURLGetFileSystemRepresentation(CFURLRef url, Boolean resolveAgainstBase, UInt8 *buffer, CFIndex maxBufLen);
CF_EXPORT CFBundleRef CFBundleGetMainBundle(void);
CF_EXPORT CFURLRef CFBundleCopyResourcesDirectoryURL(CFBundleRef bundle) CF_RETURNS_RETAINED;
CF_EXPORT CFURLRef CFBundleCopyBundleURL(CFBundleRef bundle) CF_RETURNS_RETAINED;
CF_EXPORT CFTypeRef CFBundleGetValueForInfoDictionaryKey(CFBundleRef bundle, CFStringRef key);

// CFLocale
CF_EXPORT CFLocaleRef CFLocaleCopyCurrent(void) CF_RETURNS_RETAINED;
CF_EXPORT CFArrayRef CFLocaleCopyPreferredLanguages(void) CF_RETURNS_RETAINED;
CF_EXPORT CFStringRef CFLocaleGetIdentifier(CFLocaleRef locale);

// CFRunLoop
typedef CF_ENUM(SInt32, CFRunLoopRunResult) {
	kCFRunLoopRunFinished = 1,
	kCFRunLoopRunStopped = 2,
	kCFRunLoopRunTimedOut = 3,
	kCFRunLoopRunHandledSource = 4,
};
CF_EXPORT const CFRunLoopMode kCFRunLoopDefaultMode;
CF_EXPORT const CFRunLoopMode kCFRunLoopCommonModes;
CF_EXPORT CFRunLoopRef CFRunLoopGetCurrent(void);
CF_EXPORT CFRunLoopRef CFRunLoopGetMain(void);
CF_EXPORT void CFRunLoopRun(void);
CF_EXPORT CFRunLoopRunResult CFRunLoopRunInMode(CFRunLoopMode mode, CFTimeInterval seconds, Boolean returnAfterSourceHandled);
CF_EXPORT void CFRunLoopStop(CFRunLoopRef rl);

CF_EXTERN_C_END

#endif
