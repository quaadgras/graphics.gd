// EAGLDrawable declarations for graphics.gd's iOS SDK, written from the
// public documentation of the API. Not derived from Apple's SDK headers.
#ifndef GD_OPENGLES_EAGLDRAWABLE_H
#define GD_OPENGLES_EAGLDRAWABLE_H

#import <OpenGLES/EAGL.h>

NS_ASSUME_NONNULL_BEGIN

extern NSString *const kEAGLDrawablePropertyRetainedBacking;
extern NSString *const kEAGLDrawablePropertyColorFormat;
extern NSString *const kEAGLColorFormatRGBA8;
extern NSString *const kEAGLColorFormatRGB565;
extern NSString *const kEAGLColorFormatSRGBA8;

@protocol EAGLDrawable
@property(nullable, copy) NSDictionary<NSString *, id> *drawableProperties;
@end

@interface EAGLContext (EAGLContextDrawableAdditions)
- (BOOL)renderbufferStorage:(NSUInteger)target fromDrawable:(nullable id<EAGLDrawable>)drawable;
- (BOOL)presentRenderbuffer:(NSUInteger)target;
@end

NS_ASSUME_NONNULL_END

#endif
