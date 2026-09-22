// IOSurface declarations for graphics.gd's iOS SDK, written from the public
// documentation of the API. Not derived from Apple's SDK headers.
#ifndef GD_IOSURFACEREF_H
#define GD_IOSURFACEREF_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef struct CF_BRIDGED_TYPE(id) __IOSurface *IOSurfaceRef;

CF_EXPORT size_t IOSurfaceGetWidth(IOSurfaceRef buffer);
CF_EXPORT size_t IOSurfaceGetHeight(IOSurfaceRef buffer);

CF_EXTERN_C_END

#endif
