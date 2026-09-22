// OpenGL (CGL) declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_OPENGL_OPENGL_H
#define GD_OPENGL_OPENGL_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef int GLint;
typedef unsigned int GLuint;
typedef unsigned int GLenum;
typedef struct _CGLContextObject *CGLContextObj;
typedef struct _CGLPixelFormatObject *CGLPixelFormatObj;
typedef int CGLError;
enum {
	kCGLNoError = 0,
	kCGLBadAttribute = 10000,
	kCGLBadProperty = 10001,
	kCGLBadPixelFormat = 10002,
	kCGLBadContext = 10004,
};
typedef enum _CGLContextParameter {
	kCGLCPSwapRectangle = 200,
	kCGLCPSwapInterval = 222,
	kCGLCPDispatchTableSize = 224,
	kCGLCPClientStorage = 226,
	kCGLCPSurfaceTexture = 228,
	kCGLCPSurfaceOrder = 235,
	kCGLCPSurfaceOpacity = 236,
	kCGLCPSurfaceBackingSize = 304,
	kCGLCPSurfaceSurfaceVolatile = 306,
	kCGLCPReclaimResources = 308,
	kCGLCPCurrentRendererID = 309,
	kCGLCPGPUVertexProcessing = 310,
	kCGLCPGPUFragmentProcessing = 311,
	kCGLCPHasDrawable = 314,
	kCGLCPMPSwapsInFlight = 315,
} CGLContextParameter;
typedef enum _CGLContextEnable {
	kCGLCESwapRectangle = 201,
	kCGLCESwapLimit = 203,
	kCGLCERasterization = 221,
	kCGLCEStateValidation = 301,
	kCGLCESurfaceBackingSize = 305,
	kCGLCEDisplayListOptimization = 307,
	kCGLCEMPEngine = 313,
	kCGLCECrashOnRemovedFunctions = 316,
} CGLContextEnable;

CF_EXPORT CGLContextObj CGLGetCurrentContext(void);
CF_EXPORT CGLError CGLSetCurrentContext(CGLContextObj ctx);
CF_EXPORT CGLError CGLSetParameter(CGLContextObj ctx, CGLContextParameter pname, const GLint *params);
CF_EXPORT CGLError CGLGetParameter(CGLContextObj ctx, CGLContextParameter pname, GLint *params);
CF_EXPORT CGLError CGLEnable(CGLContextObj ctx, CGLContextEnable pname);
CF_EXPORT CGLError CGLDisable(CGLContextObj ctx, CGLContextEnable pname);
CF_EXPORT CGLError CGLFlushDrawable(CGLContextObj ctx);
CF_EXPORT CGLError CGLLockContext(CGLContextObj ctx);
CF_EXPORT CGLError CGLUnlockContext(CGLContextObj ctx);
CF_EXPORT CGLError CGLUpdateContext(CGLContextObj ctx);
CF_EXPORT const char *CGLErrorString(CGLError error);

CF_EXTERN_C_END

#endif
