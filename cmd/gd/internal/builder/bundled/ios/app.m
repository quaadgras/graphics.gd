// The entry point of a Godot iOS app, for engines built by gd.
//
// Godot's own is written in Swift (drivers/apple_embedded/app.swift), which
// zig cannot compile: a SwiftUI App whose only scene hosts the engine's
// GDTViewController, with GDTApplicationDelegate adapted as its delegate.
// This is that, in terms of UIKit: gd's stand-in for swift-frontend compiles
// this file in place of the Swift it is asked for.

#import <UIKit/UIKit.h>

@interface GDTViewController : UIViewController
@end

@interface GDTAppDelegateService : NSObject
+ (void)setViewController:(GDTViewController *)viewController;
@end

@interface GDTApplicationDelegate : NSObject <UIApplicationDelegate, UIWindowSceneDelegate>
@end

// GDSceneDelegate is what a SwiftUI WindowGroup amounts to: a window with the
// engine's view controller at its root. Everything else about the scene is
// left to the engine's delegate, as it is when SwiftUI adapts it.
@interface GDSceneDelegate : GDTApplicationDelegate
@property(strong, nonatomic) UIWindow *window;
@end

@implementation GDSceneDelegate

- (void)scene:(UIScene *)scene willConnectToSession:(UISceneSession *)session options:(UISceneConnectionOptions *)connectionOptions {
	[super scene:scene willConnectToSession:session options:connectionOptions];
	if (![scene isKindOfClass:[UIWindowScene class]]) {
		return;
	}
	GDTViewController *viewController = [[GDTViewController alloc] init];
	[GDTAppDelegateService setViewController:viewController];
	self.window = [[UIWindow alloc] initWithWindowScene:(UIWindowScene *)scene];
	self.window.rootViewController = viewController;
	[self.window makeKeyAndVisible];
}

@end

@interface GDApplicationDelegate : GDTApplicationDelegate
@end

@implementation GDApplicationDelegate

- (UISceneConfiguration *)application:(UIApplication *)application configurationForConnectingSceneSession:(UISceneSession *)connectingSceneSession options:(UISceneConnectionOptions *)options {
	UISceneConfiguration *configuration = [super application:application configurationForConnectingSceneSession:connectingSceneSession options:options];
	configuration.delegateClass = [GDSceneDelegate class];
	return configuration;
}

@end

int main(int argc, char *argv[]) {
	@autoreleasepool {
		return UIApplicationMain(argc, argv, nil, NSStringFromClass([GDApplicationDelegate class]));
	}
}
