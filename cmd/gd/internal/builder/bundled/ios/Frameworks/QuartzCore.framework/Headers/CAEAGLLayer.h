// CAEAGLLayer declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_CAEAGLLAYER_H
#define GD_CAEAGLLAYER_H

#import <OpenGLES/EAGLDrawable.h>
#import <QuartzCore/QuartzCore.h>

NS_ASSUME_NONNULL_BEGIN

@interface CAEAGLLayer : CALayer <EAGLDrawable>
@property BOOL presentsWithTransaction;
@end

NS_ASSUME_NONNULL_END

#endif
