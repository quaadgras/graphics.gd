// CAMetalLayer declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_CAMETALLAYER_H
#define GD_CAMETALLAYER_H

#import <Metal/Metal.h>
#import <QuartzCore/QuartzCore.h>

#if defined(__OBJC__)

NS_ASSUME_NONNULL_BEGIN

@class CAMetalLayer;

@protocol CAMetalDrawable <MTLDrawable>
@property(readonly) id<MTLTexture> texture;
@property(readonly) CAMetalLayer *layer;
@end

@interface CAMetalLayer : CALayer
@property(nullable, retain) id<MTLDevice> device;
@property(nullable, readonly) id<MTLDevice> preferredDevice;
@property MTLPixelFormat pixelFormat;
@property BOOL framebufferOnly;
@property CGSize drawableSize;
@property NSUInteger maximumDrawableCount;
@property BOOL presentsWithTransaction;
@property(nullable) CGColorSpaceRef colorspace;
@property BOOL wantsExtendedDynamicRangeContent;
@property BOOL allowsNextDrawableTimeout;
- (nullable id<CAMetalDrawable>)nextDrawable;
@end

NS_ASSUME_NONNULL_END

#endif // __OBJC__

#endif
