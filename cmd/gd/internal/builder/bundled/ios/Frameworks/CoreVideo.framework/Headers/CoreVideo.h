// CoreVideo declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREVIDEO_H
#define GD_COREVIDEO_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef SInt32 CVReturn;
typedef UInt64 CVOptionFlags;
enum {
	kCVReturnSuccess = 0,
};

typedef struct CF_BRIDGED_TYPE(id) __CVBuffer *CVBufferRef;
typedef CVBufferRef CVImageBufferRef;
typedef CVImageBufferRef CVPixelBufferRef;

typedef CF_OPTIONS(CVOptionFlags, CVPixelBufferLockFlags) {
	kCVPixelBufferLock_ReadOnly = 0x00000001,
};

enum {
	kCVPixelFormatType_24RGB = 0x00000018,
	kCVPixelFormatType_32ARGB = 0x00000020,
	kCVPixelFormatType_32BGRA = 'BGRA',
	kCVPixelFormatType_32RGBA = 'RGBA',
	kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange = '420v',
	kCVPixelFormatType_420YpCbCr8BiPlanarFullRange = '420f',
	kCVPixelFormatType_422YpCbCr8 = '2vuy',
};

CF_EXPORT const CFStringRef kCVPixelBufferPixelFormatTypeKey;
CF_EXPORT const CFStringRef kCVPixelBufferWidthKey;
CF_EXPORT const CFStringRef kCVPixelBufferHeightKey;

CF_EXPORT CVReturn CVPixelBufferLockBaseAddress(CVPixelBufferRef pixelBuffer, CVPixelBufferLockFlags lockFlags);
CF_EXPORT CVReturn CVPixelBufferUnlockBaseAddress(CVPixelBufferRef pixelBuffer, CVPixelBufferLockFlags unlockFlags);
CF_EXPORT size_t CVPixelBufferGetWidth(CVPixelBufferRef pixelBuffer);
CF_EXPORT size_t CVPixelBufferGetHeight(CVPixelBufferRef pixelBuffer);
CF_EXPORT size_t CVPixelBufferGetBytesPerRow(CVPixelBufferRef pixelBuffer);
CF_EXPORT void *CVPixelBufferGetBaseAddress(CVPixelBufferRef pixelBuffer);
CF_EXPORT OSType CVPixelBufferGetPixelFormatType(CVPixelBufferRef pixelBuffer);
CF_EXPORT Boolean CVPixelBufferIsPlanar(CVPixelBufferRef pixelBuffer);
CF_EXPORT size_t CVPixelBufferGetPlaneCount(CVPixelBufferRef pixelBuffer);
CF_EXPORT size_t CVPixelBufferGetWidthOfPlane(CVPixelBufferRef pixelBuffer, size_t planeIndex);
CF_EXPORT size_t CVPixelBufferGetHeightOfPlane(CVPixelBufferRef pixelBuffer, size_t planeIndex);
CF_EXPORT size_t CVPixelBufferGetBytesPerRowOfPlane(CVPixelBufferRef pixelBuffer, size_t planeIndex);
CF_EXPORT void *CVPixelBufferGetBaseAddressOfPlane(CVPixelBufferRef pixelBuffer, size_t planeIndex);

// Display links (macOS)
typedef struct CF_BRIDGED_TYPE(id) __CVDisplayLink *CVDisplayLinkRef;
typedef uint32_t CVDirectDisplayID_;
typedef struct {
	uint32_t version;
	int32_t videoTimeScale;
	int64_t videoTime;
	uint64_t hostTime;
	double rateScalar;
	int64_t videoRefreshPeriod;
	uint8_t smpteTime[24];
	uint64_t flags;
	uint64_t reserved;
} CVTimeStamp;
typedef CVReturn (*CVDisplayLinkOutputCallback)(CVDisplayLinkRef displayLink, const CVTimeStamp *inNow, const CVTimeStamp *inOutputTime, CVOptionFlags flagsIn, CVOptionFlags *flagsOut, void *displayLinkContext);
typedef CVReturn (^CVDisplayLinkOutputHandler)(CVDisplayLinkRef displayLink, const CVTimeStamp *inNow, const CVTimeStamp *inOutputTime, CVOptionFlags flagsIn, CVOptionFlags *flagsOut);
CF_EXPORT CVReturn CVDisplayLinkCreateWithCGDisplay(uint32_t displayID, CVDisplayLinkRef *displayLinkOut);
CF_EXPORT CVReturn CVDisplayLinkCreateWithActiveCGDisplays(CVDisplayLinkRef *displayLinkOut);
CF_EXPORT CVReturn CVDisplayLinkSetCurrentCGDisplay(CVDisplayLinkRef displayLink, uint32_t displayID);
CF_EXPORT CVReturn CVDisplayLinkSetOutputCallback(CVDisplayLinkRef displayLink, CVDisplayLinkOutputCallback callback, void *userInfo);
CF_EXPORT CVReturn CVDisplayLinkSetOutputHandler(CVDisplayLinkRef displayLink, CVDisplayLinkOutputHandler handler);
CF_EXPORT CVReturn CVDisplayLinkStart(CVDisplayLinkRef displayLink);
CF_EXPORT CVReturn CVDisplayLinkStop(CVDisplayLinkRef displayLink);
CF_EXPORT Boolean CVDisplayLinkIsRunning(CVDisplayLinkRef displayLink);
CF_EXPORT void CVDisplayLinkRelease(CVDisplayLinkRef displayLink);
CF_EXPORT CVDisplayLinkRef CVDisplayLinkRetain(CVDisplayLinkRef displayLink);

CF_EXTERN_C_END

#endif
