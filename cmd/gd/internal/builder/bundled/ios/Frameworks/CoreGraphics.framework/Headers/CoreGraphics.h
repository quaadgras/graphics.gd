// CoreGraphics declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot,
// SDL and metal-cpp make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREGRAPHICS_H
#define GD_COREGRAPHICS_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

#if defined(__LP64__) && __LP64__
typedef double CGFloat;
#define CGFLOAT_IS_DOUBLE 1
#define CGFLOAT_MIN DBL_MIN
#define CGFLOAT_MAX DBL_MAX
#else
typedef float CGFloat;
#define CGFLOAT_IS_DOUBLE 0
#define CGFLOAT_MIN FLT_MIN
#define CGFLOAT_MAX FLT_MAX
#endif
#define CGFLOAT_DEFINED 1

typedef struct CGPoint {
	CGFloat x;
	CGFloat y;
} CGPoint;
typedef struct CGSize {
	CGFloat width;
	CGFloat height;
} CGSize;
typedef struct CGVector {
	CGFloat dx;
	CGFloat dy;
} CGVector;
typedef struct CGRect {
	CGPoint origin;
	CGSize size;
} CGRect;
typedef struct CGAffineTransform {
	CGFloat a, b, c, d;
	CGFloat tx, ty;
} CGAffineTransform;

CF_EXPORT const CGPoint CGPointZero;
CF_EXPORT const CGSize CGSizeZero;
CF_EXPORT const CGRect CGRectZero;
CF_EXPORT const CGRect CGRectNull;
CF_EXPORT const CGAffineTransform CGAffineTransformIdentity;

CF_INLINE CGPoint CGPointMake(CGFloat x, CGFloat y) {
	CGPoint p;
	p.x = x;
	p.y = y;
	return p;
}
CF_INLINE CGSize CGSizeMake(CGFloat width, CGFloat height) {
	CGSize size;
	size.width = width;
	size.height = height;
	return size;
}
CF_INLINE CGVector CGVectorMake(CGFloat dx, CGFloat dy) {
	CGVector vector;
	vector.dx = dx;
	vector.dy = dy;
	return vector;
}
CF_INLINE CGRect CGRectMake(CGFloat x, CGFloat y, CGFloat width, CGFloat height) {
	CGRect rect;
	rect.origin.x = x;
	rect.origin.y = y;
	rect.size.width = width;
	rect.size.height = height;
	return rect;
}
CF_INLINE bool CGSizeEqualToSize(CGSize size1, CGSize size2) {
	return size1.width == size2.width && size1.height == size2.height;
}
CF_INLINE bool CGPointEqualToPoint(CGPoint point1, CGPoint point2) {
	return point1.x == point2.x && point1.y == point2.y;
}

CF_EXPORT CGFloat CGRectGetMinX(CGRect rect);
CF_EXPORT CGFloat CGRectGetMidX(CGRect rect);
CF_EXPORT CGFloat CGRectGetMaxX(CGRect rect);
CF_EXPORT CGFloat CGRectGetMinY(CGRect rect);
CF_EXPORT CGFloat CGRectGetMidY(CGRect rect);
CF_EXPORT CGFloat CGRectGetMaxY(CGRect rect);
CF_EXPORT CGFloat CGRectGetWidth(CGRect rect);
CF_EXPORT CGFloat CGRectGetHeight(CGRect rect);
CF_EXPORT bool CGRectEqualToRect(CGRect rect1, CGRect rect2);
CF_EXPORT bool CGRectIsEmpty(CGRect rect);
CF_EXPORT bool CGRectContainsPoint(CGRect rect, CGPoint point);
CF_EXPORT bool CGRectIntersectsRect(CGRect rect1, CGRect rect2);
CF_EXPORT CGAffineTransform CGAffineTransformMakeRotation(CGFloat angle);
CF_EXPORT CGAffineTransform CGAffineTransformMakeScale(CGFloat sx, CGFloat sy);

typedef struct CF_BRIDGED_TYPE(id) CGColorSpace *CGColorSpaceRef;
typedef struct CF_BRIDGED_TYPE(id) CGColor *CGColorRef;
typedef struct CF_BRIDGED_TYPE(id) CGImage *CGImageRef;
typedef struct CF_BRIDGED_TYPE(id) CGContext *CGContextRef;

CF_EXPORT const CFStringRef kCGColorSpaceSRGB;
CF_EXPORT const CFStringRef kCGColorSpaceLinearSRGB;
CF_EXPORT const CFStringRef kCGColorSpaceExtendedSRGB;
CF_EXPORT const CFStringRef kCGColorSpaceExtendedLinearSRGB;
CF_EXPORT const CFStringRef kCGColorSpaceDisplayP3;
CF_EXPORT const CFStringRef kCGColorSpaceExtendedLinearDisplayP3;
CF_EXPORT const CFStringRef kCGColorSpaceITUR_2100_PQ;
CF_EXPORT const CFStringRef kCGColorSpaceITUR_2020;
CF_EXPORT const CFStringRef kCGColorSpaceExtendedLinearITUR_2020;
CF_EXPORT CGColorSpaceRef CGColorSpaceCreateWithName(CFStringRef name) CF_RETURNS_RETAINED;
CF_EXPORT CGColorSpaceRef CGColorSpaceCreateDeviceRGB(void) CF_RETURNS_RETAINED;
CF_EXPORT void CGColorSpaceRelease(CGColorSpaceRef space);
CF_EXPORT CGColorSpaceRef CGColorSpaceRetain(CGColorSpaceRef space);

CF_EXTERN_C_END

#endif
