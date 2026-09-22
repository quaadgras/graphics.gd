// EAGL declarations for graphics.gd's iOS SDK, written from the public
// documentation of the API. Not derived from Apple's SDK headers.
#ifndef GD_OPENGLES_EAGL_H
#define GD_OPENGLES_EAGL_H

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

typedef NS_ENUM(NSUInteger, EAGLRenderingAPI) {
	kEAGLRenderingAPIOpenGLES1 = 1,
	kEAGLRenderingAPIOpenGLES2 = 2,
	kEAGLRenderingAPIOpenGLES3 = 3,
};

@interface EAGLSharegroup : NSObject
@property(nullable, copy, nonatomic) NSString *debugLabel;
@end

@interface EAGLContext : NSObject
- (nullable instancetype)initWithAPI:(EAGLRenderingAPI)api;
- (nullable instancetype)initWithAPI:(EAGLRenderingAPI)api sharegroup:(EAGLSharegroup *)sharegroup;
+ (BOOL)setCurrentContext:(nullable EAGLContext *)context;
+ (nullable EAGLContext *)currentContext;
@property(readonly) EAGLRenderingAPI API;
@property(nonnull, readonly) EAGLSharegroup *sharegroup;
@property(nullable, copy, nonatomic) NSString *debugLabel;
@property(getter=isMultiThreaded, nonatomic) BOOL multiThreaded;
@end

NS_ASSUME_NONNULL_END

#endif
