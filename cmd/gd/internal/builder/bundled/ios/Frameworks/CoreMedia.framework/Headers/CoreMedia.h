// CoreMedia declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREMEDIA_H
#define GD_COREMEDIA_H

#include <CoreVideo/CoreVideo.h>

CF_EXTERN_C_BEGIN

typedef int64_t CMTimeValue;
typedef int32_t CMTimeScale;
typedef int64_t CMTimeEpoch;
typedef uint32_t CMTimeFlags;
typedef struct {
	CMTimeValue value;
	CMTimeScale timescale;
	CMTimeFlags flags;
	CMTimeEpoch epoch;
} CMTime;

typedef struct CF_BRIDGED_TYPE(id) opaqueCMSampleBuffer *CMSampleBufferRef;
typedef const struct CF_BRIDGED_TYPE(id) opaqueCMFormatDescription *CMFormatDescriptionRef;

CF_EXPORT CVImageBufferRef CMSampleBufferGetImageBuffer(CMSampleBufferRef sbuf);
CF_EXPORT CMTime CMSampleBufferGetPresentationTimeStamp(CMSampleBufferRef sbuf);
CF_EXPORT CMTime CMTimeMake(int64_t value, int32_t timescale);

CF_EXTERN_C_END

#endif
