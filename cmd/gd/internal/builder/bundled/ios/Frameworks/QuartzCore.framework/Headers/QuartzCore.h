// QuartzCore declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_QUARTZCORE_H
#define GD_QUARTZCORE_H

#import <Foundation/Foundation.h>

#if defined(__OBJC__)

NS_ASSUME_NONNULL_BEGIN

FOUNDATION_EXTERN CFTimeInterval CACurrentMediaTime(void);

typedef struct CATransform3D {
	CGFloat m11, m12, m13, m14;
	CGFloat m21, m22, m23, m24;
	CGFloat m31, m32, m33, m34;
	CGFloat m41, m42, m43, m44;
} CATransform3D;

typedef struct CAFrameRateRange {
	float minimum;
	float maximum;
	float preferred;
} CAFrameRateRange;

@protocol CAMediaTiming
@property CFTimeInterval beginTime;
@property CFTimeInterval duration;
@property float speed;
@end

@protocol CALayerDelegate <NSObject>
@end

@interface CALayer : NSObject <NSSecureCoding, CAMediaTiming>
+ (instancetype)layer;
- (instancetype)init;
@property CGRect bounds;
@property CGPoint position;
@property CGPoint anchorPoint;
@property CGRect frame;
@property CATransform3D transform;
@property(getter=isHidden) BOOL hidden;
@property(getter=isOpaque) BOOL opaque;
@property float opacity;
@property CGFloat contentsScale;
@property CGFloat cornerRadius;
@property BOOL masksToBounds;
@property(nullable) CGColorRef backgroundColor;
@property(nullable, strong) id contents;
@property(nullable, copy) NSString *name;
@property(nullable, weak) id<CALayerDelegate> delegate;
@property(nullable, readonly) CALayer *superlayer;
@property(nullable, copy) NSArray<__kindof CALayer *> *sublayers;
- (void)addSublayer:(CALayer *)layer;
- (void)insertSublayer:(CALayer *)layer atIndex:(unsigned)idx;
- (void)removeFromSuperlayer;
- (void)setNeedsDisplay;
- (void)setNeedsLayout;
- (void)layoutIfNeeded;
- (void)layoutSublayers;
- (void)display;
- (void)removeAllAnimations;
@end

@interface CATransaction : NSObject
+ (void)begin;
+ (void)commit;
+ (void)flush;
+ (void)setDisableActions:(BOOL)flag;
@end

@interface CADisplayLink : NSObject
+ (CADisplayLink *)displayLinkWithTarget:(id)target selector:(SEL)sel;
- (void)addToRunLoop:(NSRunLoop *)runloop forMode:(NSRunLoopMode)mode;
- (void)removeFromRunLoop:(NSRunLoop *)runloop forMode:(NSRunLoopMode)mode;
- (void)invalidate;
@property(readonly, nonatomic) CFTimeInterval timestamp;
@property(readonly, nonatomic) CFTimeInterval duration;
@property(readonly, nonatomic) CFTimeInterval targetTimestamp;
@property(getter=isPaused, nonatomic) BOOL paused;
@property(nonatomic) NSInteger frameInterval;
@property(nonatomic) NSInteger preferredFramesPerSecond;
@property(nonatomic) CAFrameRateRange preferredFrameRateRange;
@end

NS_ASSUME_NONNULL_END

#import <QuartzCore/CAEAGLLayer.h>
#import <QuartzCore/CAMetalLayer.h>

#endif // __OBJC__

#endif
