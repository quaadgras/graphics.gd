// Metal (Objective-C) declarations for graphics.gd's iOS SDK. The engine drives
// Metal through metal-cpp, this only covers what crosses into Objective-C.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_METAL_H
#define GD_METAL_H

#import <Foundation/Foundation.h>

#if defined(__OBJC__)

NS_ASSUME_NONNULL_BEGIN

typedef NS_ENUM(NSUInteger, MTLPixelFormat) {
	MTLPixelFormatInvalid = 0,
	MTLPixelFormatR8Unorm = 10,
	MTLPixelFormatRG8Unorm = 30,
	MTLPixelFormatRGBA8Uint = 73,
	MTLPixelFormatRGBA8Unorm = 70,
	MTLPixelFormatRGBA8Unorm_sRGB = 71,
	MTLPixelFormatBGRA8Unorm = 80,
	MTLPixelFormatBGRA8Unorm_sRGB = 81,
	MTLPixelFormatRGB10A2Unorm = 90,
	MTLPixelFormatBGR10A2Unorm = 94,
	MTLPixelFormatRGBA16Float = 115,
	MTLPixelFormatBGR10_XR = 554,
	MTLPixelFormatBGR10_XR_sRGB = 555,
	MTLPixelFormatDepth16Unorm = 250,
	MTLPixelFormatDepth32Float = 252,
	MTLPixelFormatStencil8 = 253,
	MTLPixelFormatDepth24Unorm_Stencil8 = 255,
	MTLPixelFormatDepth32Float_Stencil8 = 260,
};

@protocol MTLDevice <NSObject>
@property(readonly) NSString *name;
@end

@protocol MTLResource <NSObject>
@property(nullable, copy, atomic) NSString *label;
@property(readonly) id<MTLDevice> device;
@end

@protocol MTLTexture <MTLResource>
@property(readonly) MTLPixelFormat pixelFormat;
@property(readonly) NSUInteger width;
@property(readonly) NSUInteger height;
@end

@protocol MTLDrawable <NSObject>
- (void)present;
- (void)presentAtTime:(CFTimeInterval)presentationTime;
@property(nonatomic, readonly) CFTimeInterval presentedTime;
@property(nonatomic, readonly) NSUInteger drawableID;
@end

FOUNDATION_EXTERN id<MTLDevice> _Nullable MTLCreateSystemDefaultDevice(void) NS_RETURNS_RETAINED;

NS_ASSUME_NONNULL_END

#endif // __OBJC__

#endif
