// UIKit declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_UIKIT_H
#define GD_UIKIT_H

#import <Foundation/Foundation.h>
#import <QuartzCore/QuartzCore.h>
#import <UIKit/UIKeyConstants.h>

#if defined(__cplusplus)
#define UIKIT_EXTERN extern "C"
#else
#define UIKIT_EXTERN extern
#endif
#define UIKIT_STATIC_INLINE static inline
#define IBOutlet
#define IBAction void
#define IBInspectable

NS_ASSUME_NONNULL_BEGIN

@class INIntent, CKShareMetadata; // from the Intents & CloudKit frameworks.
@class UIView, UIWindow, UIViewController, UIScreen, UIEvent, UITouch, UIPress, UIPressesEvent, UIKey, UIColor, UIImage, UIFont, UIScene, UISceneSession, UISceneConfiguration, UISceneConnectionOptions, UIWindowScene, UIOpenURLContext, UIApplication, UIApplicationShortcutItem, UITraitCollection, UIGestureRecognizer, UIStoryboard, UITextPosition, UITextRange, UIAlertAction, UIPasteboard, UIScreenMode, UIResponder, UILayoutGuide;

// Geometry

typedef struct UIEdgeInsets {
	CGFloat top, left, bottom, right;
} UIEdgeInsets;
UIKIT_EXTERN const UIEdgeInsets UIEdgeInsetsZero;
UIKIT_STATIC_INLINE UIEdgeInsets UIEdgeInsetsMake(CGFloat top, CGFloat left, CGFloat bottom, CGFloat right) {
	UIEdgeInsets insets = { top, left, bottom, right };
	return insets;
}

typedef NS_OPTIONS(NSUInteger, UIRectEdge) {
	UIRectEdgeNone = 0,
	UIRectEdgeTop = 1 << 0,
	UIRectEdgeLeft = 1 << 1,
	UIRectEdgeBottom = 1 << 2,
	UIRectEdgeRight = 1 << 3,
	UIRectEdgeAll = UIRectEdgeTop | UIRectEdgeLeft | UIRectEdgeBottom | UIRectEdgeRight,
};

@interface NSValue (NSValueUIGeometryExtensions)
+ (NSValue *)valueWithCGPoint:(CGPoint)point;
+ (NSValue *)valueWithCGSize:(CGSize)size;
+ (NSValue *)valueWithCGRect:(CGRect)rect;
@property(nonatomic, readonly) CGPoint CGPointValue;
@property(nonatomic, readonly) CGSize CGSizeValue;
@property(nonatomic, readonly) CGRect CGRectValue;
@end

UIKIT_EXTERN NSString *NSStringFromCGRect(CGRect rect);
UIKIT_EXTERN NSString *NSStringFromCGSize(CGSize size);
UIKIT_EXTERN NSString *NSStringFromCGPoint(CGPoint point);

// Orientation & idiom

typedef NS_ENUM(NSInteger, UIDeviceOrientation) {
	UIDeviceOrientationUnknown,
	UIDeviceOrientationPortrait,
	UIDeviceOrientationPortraitUpsideDown,
	UIDeviceOrientationLandscapeLeft,
	UIDeviceOrientationLandscapeRight,
	UIDeviceOrientationFaceUp,
	UIDeviceOrientationFaceDown,
};
typedef NS_ENUM(NSInteger, UIInterfaceOrientation) {
	UIInterfaceOrientationUnknown = UIDeviceOrientationUnknown,
	UIInterfaceOrientationPortrait = UIDeviceOrientationPortrait,
	UIInterfaceOrientationPortraitUpsideDown = UIDeviceOrientationPortraitUpsideDown,
	UIInterfaceOrientationLandscapeLeft = UIDeviceOrientationLandscapeRight,
	UIInterfaceOrientationLandscapeRight = UIDeviceOrientationLandscapeLeft,
};
typedef NS_OPTIONS(NSUInteger, UIInterfaceOrientationMask) {
	UIInterfaceOrientationMaskPortrait = (1 << UIInterfaceOrientationPortrait),
	UIInterfaceOrientationMaskLandscapeLeft = (1 << UIInterfaceOrientationLandscapeLeft),
	UIInterfaceOrientationMaskLandscapeRight = (1 << UIInterfaceOrientationLandscapeRight),
	UIInterfaceOrientationMaskPortraitUpsideDown = (1 << UIInterfaceOrientationPortraitUpsideDown),
	UIInterfaceOrientationMaskLandscape = (UIInterfaceOrientationMaskLandscapeLeft | UIInterfaceOrientationMaskLandscapeRight),
	UIInterfaceOrientationMaskAll = (UIInterfaceOrientationMaskPortrait | UIInterfaceOrientationMaskLandscapeLeft | UIInterfaceOrientationMaskLandscapeRight | UIInterfaceOrientationMaskPortraitUpsideDown),
	UIInterfaceOrientationMaskAllButUpsideDown = (UIInterfaceOrientationMaskPortrait | UIInterfaceOrientationMaskLandscapeLeft | UIInterfaceOrientationMaskLandscapeRight),
};
UIKIT_STATIC_INLINE BOOL UIInterfaceOrientationIsPortrait(UIInterfaceOrientation orientation) {
	return orientation == UIInterfaceOrientationPortrait || orientation == UIInterfaceOrientationPortraitUpsideDown;
}
UIKIT_STATIC_INLINE BOOL UIInterfaceOrientationIsLandscape(UIInterfaceOrientation orientation) {
	return orientation == UIInterfaceOrientationLandscapeLeft || orientation == UIInterfaceOrientationLandscapeRight;
}
UIKIT_STATIC_INLINE BOOL UIDeviceOrientationIsLandscape(UIDeviceOrientation orientation) {
	return orientation == UIDeviceOrientationLandscapeLeft || orientation == UIDeviceOrientationLandscapeRight;
}
UIKIT_STATIC_INLINE BOOL UIDeviceOrientationIsPortrait(UIDeviceOrientation orientation) {
	return orientation == UIDeviceOrientationPortrait || orientation == UIDeviceOrientationPortraitUpsideDown;
}

typedef NS_ENUM(NSInteger, UIUserInterfaceIdiom) {
	UIUserInterfaceIdiomUnspecified = -1,
	UIUserInterfaceIdiomPhone,
	UIUserInterfaceIdiomPad,
	UIUserInterfaceIdiomTV,
	UIUserInterfaceIdiomCarPlay,
	UIUserInterfaceIdiomMac = 5,
	UIUserInterfaceIdiomVision = 6,
};
typedef NS_ENUM(NSInteger, UIUserInterfaceStyle) {
	UIUserInterfaceStyleUnspecified,
	UIUserInterfaceStyleLight,
	UIUserInterfaceStyleDark,
};
typedef NS_ENUM(NSInteger, UIUserInterfaceSizeClass) {
	UIUserInterfaceSizeClassUnspecified = 0,
	UIUserInterfaceSizeClassCompact = 1,
	UIUserInterfaceSizeClassRegular = 2,
};

// UIColor, UIImage & UIFont

@interface UIColor : NSObject <NSSecureCoding, NSCopying>
+ (UIColor *)colorWithWhite:(CGFloat)white alpha:(CGFloat)alpha;
+ (UIColor *)colorWithRed:(CGFloat)red green:(CGFloat)green blue:(CGFloat)blue alpha:(CGFloat)alpha;
@property(class, nonatomic, readonly) UIColor *blackColor;
@property(class, nonatomic, readonly) UIColor *whiteColor;
@property(class, nonatomic, readonly) UIColor *clearColor;
@property(class, nonatomic, readonly) UIColor *systemBackgroundColor;
@property(nonatomic, readonly) CGColorRef CGColor;
- (BOOL)getRed:(nullable CGFloat *)red green:(nullable CGFloat *)green blue:(nullable CGFloat *)blue alpha:(nullable CGFloat *)alpha;
@end

@interface UIImage : NSObject <NSSecureCoding>
+ (nullable UIImage *)imageNamed:(NSString *)name;
+ (nullable UIImage *)imageWithContentsOfFile:(NSString *)path;
+ (nullable UIImage *)imageWithData:(NSData *)data;
@property(nonatomic, readonly) CGSize size;
@property(nonatomic, readonly) CGFloat scale;
@property(nullable, nonatomic, readonly) CGImageRef CGImage;
@end

@interface UIFont : NSObject <NSCopying, NSSecureCoding>
+ (UIFont *)systemFontOfSize:(CGFloat)fontSize;
@end

// UITraitCollection

@protocol UITraitDefinition
@end
typedef Class<UITraitDefinition> UITrait;
@interface UITraitUserInterfaceStyle : NSObject <UITraitDefinition>
@end
@interface UITraitUserInterfaceIdiom : NSObject <UITraitDefinition>
@end

@interface UITraitCollection : NSObject <NSCopying, NSSecureCoding>
@property(class, nonatomic, readonly) UITraitCollection *currentTraitCollection;
@property(nonatomic, readonly) UIUserInterfaceIdiom userInterfaceIdiom;
@property(nonatomic, readonly) UIUserInterfaceStyle userInterfaceStyle;
@property(nonatomic, readonly) UIUserInterfaceSizeClass horizontalSizeClass;
@property(nonatomic, readonly) UIUserInterfaceSizeClass verticalSizeClass;
@property(nonatomic, readonly) CGFloat displayScale;
- (BOOL)hasDifferentColorAppearanceComparedToTraitCollection:(nullable UITraitCollection *)traitCollection;
@end

@protocol UITraitEnvironment <NSObject>
@property(nonatomic, readonly) UITraitCollection *traitCollection;
- (void)traitCollectionDidChange:(nullable UITraitCollection *)previousTraitCollection;
@end

@protocol UITraitChangeRegistration <NSObject, NSCopying>
@end
@protocol UITraitChangeObservable
- (id<UITraitChangeRegistration>)registerForTraitChanges:(NSArray<UITrait> *)traits withTarget:(id)target action:(SEL)action;
- (id<UITraitChangeRegistration>)registerForTraitChanges:(NSArray<UITrait> *)traits withAction:(SEL)action;
- (void)unregisterForTraitChanges:(id<UITraitChangeRegistration>)registration;
@end

// UIResponder, touches & presses

typedef NS_ENUM(NSInteger, UITouchPhase) {
	UITouchPhaseBegan,
	UITouchPhaseMoved,
	UITouchPhaseStationary,
	UITouchPhaseEnded,
	UITouchPhaseCancelled,
	UITouchPhaseRegionEntered,
	UITouchPhaseRegionMoved,
	UITouchPhaseRegionExited,
};
typedef NS_ENUM(NSInteger, UITouchType) {
	UITouchTypeDirect,
	UITouchTypeIndirect,
	UITouchTypePencil,
	UITouchTypeStylus = UITouchTypePencil,
	UITouchTypeIndirectPointer,
};
typedef NS_ENUM(NSInteger, UIEventType) {
	UIEventTypeTouches,
	UIEventTypeMotion,
	UIEventTypeRemoteControl,
	UIEventTypePresses,
	UIEventTypeScroll = 10,
	UIEventTypeHover = 11,
	UIEventTypeTransform = 14,
};
typedef NS_ENUM(NSInteger, UIEventSubtype) {
	UIEventSubtypeNone = 0,
	UIEventSubtypeMotionShake = 1,
};
typedef NS_OPTIONS(NSInteger, UIEventButtonMask) {
	UIEventButtonMaskPrimary = 1 << 0,
	UIEventButtonMaskSecondary = 1 << 1,
};
typedef NS_OPTIONS(NSInteger, UIKeyModifierFlags) {
	UIKeyModifierAlphaShift = 1 << 16,
	UIKeyModifierShift = 1 << 17,
	UIKeyModifierControl = 1 << 18,
	UIKeyModifierAlternate = 1 << 19,
	UIKeyModifierCommand = 1 << 20,
	UIKeyModifierNumericPad = 1 << 21,
};
typedef NS_ENUM(NSInteger, UIPressPhase) {
	UIPressPhaseBegan,
	UIPressPhaseChanged,
	UIPressPhaseStationary,
	UIPressPhaseEnded,
	UIPressPhaseCancelled,
};
typedef NS_ENUM(NSInteger, UIPressType) {
	UIPressTypeUpArrow,
	UIPressTypeDownArrow,
	UIPressTypeLeftArrow,
	UIPressTypeRightArrow,
	UIPressTypeSelect,
	UIPressTypeMenu,
	UIPressTypePlayPause,
};

@interface UITouch : NSObject
@property(nonatomic, readonly) NSTimeInterval timestamp;
@property(nonatomic, readonly) UITouchPhase phase;
@property(nonatomic, readonly) NSUInteger tapCount;
@property(nonatomic, readonly) UITouchType type;
@property(nonatomic, readonly) CGFloat majorRadius;
@property(nonatomic, readonly) CGFloat force;
@property(nonatomic, readonly) CGFloat maximumPossibleForce;
@property(nonatomic, readonly) CGFloat altitudeAngle;
@property(nullable, nonatomic, readonly, strong) UIWindow *window;
@property(nullable, nonatomic, readonly, strong) UIView *view;
- (CGPoint)locationInView:(nullable UIView *)view;
- (CGPoint)previousLocationInView:(nullable UIView *)view;
- (CGPoint)preciseLocationInView:(nullable UIView *)view;
- (CGFloat)azimuthAngleInView:(nullable UIView *)view;
- (CGVector)azimuthUnitVectorInView:(nullable UIView *)view;
@end

@interface UIEvent : NSObject
@property(nonatomic, readonly) UIEventType type;
@property(nonatomic, readonly) UIEventSubtype subtype;
@property(nonatomic, readonly) NSTimeInterval timestamp;
@property(nonatomic, readonly) UIKeyModifierFlags modifierFlags;
@property(nonatomic, readonly) UIEventButtonMask buttonMask;
@property(nonatomic, readonly, nullable) NSSet<UITouch *> *allTouches;
- (nullable NSSet<UITouch *> *)touchesForWindow:(UIWindow *)window;
- (nullable NSSet<UITouch *> *)touchesForView:(UIView *)view;
- (nullable NSArray<UITouch *> *)coalescedTouchesForTouch:(UITouch *)touch;
@end

@interface UIKey : NSObject <NSCopying, NSCoding>
@property(nonatomic, readonly) NSString *characters;
@property(nonatomic, readonly) NSString *charactersIgnoringModifiers;
@property(nonatomic, readonly) UIKeyModifierFlags modifierFlags;
@property(nonatomic, readonly) UIKeyboardHIDUsage keyCode;
@end

@interface UIPress : NSObject
@property(nonatomic, readonly) NSTimeInterval timestamp;
@property(nonatomic, readonly) UIPressPhase phase;
@property(nonatomic, readonly) UIPressType type;
@property(nullable, nonatomic, readonly, strong) UIWindow *window;
@property(nullable, nonatomic, readonly, strong) UIResponder *responder;
@property(nonatomic, readonly) CGFloat force;
@property(nonatomic, nullable, readonly) UIKey *key;
@end

@interface UIPressesEvent : UIEvent
@property(nonatomic, readonly) NSSet<UIPress *> *allPresses;
@end

@interface UIKeyCommand : NSObject
@end

@interface UIResponder : NSObject
@property(nonatomic, readonly, nullable) UIResponder *nextResponder;
@property(nonatomic, readonly) BOOL canBecomeFirstResponder;
- (BOOL)becomeFirstResponder;
@property(nonatomic, readonly) BOOL canResignFirstResponder;
- (BOOL)resignFirstResponder;
@property(nonatomic, readonly) BOOL isFirstResponder;
- (void)touchesBegan:(NSSet<UITouch *> *)touches withEvent:(nullable UIEvent *)event;
- (void)touchesMoved:(NSSet<UITouch *> *)touches withEvent:(nullable UIEvent *)event;
- (void)touchesEnded:(NSSet<UITouch *> *)touches withEvent:(nullable UIEvent *)event;
- (void)touchesCancelled:(NSSet<UITouch *> *)touches withEvent:(nullable UIEvent *)event;
- (void)pressesBegan:(NSSet<UIPress *> *)presses withEvent:(nullable UIPressesEvent *)event;
- (void)pressesChanged:(NSSet<UIPress *> *)presses withEvent:(nullable UIPressesEvent *)event;
- (void)pressesEnded:(NSSet<UIPress *> *)presses withEvent:(nullable UIPressesEvent *)event;
- (void)pressesCancelled:(NSSet<UIPress *> *)presses withEvent:(nullable UIPressesEvent *)event;
- (void)motionBegan:(UIEventSubtype)motion withEvent:(nullable UIEvent *)event;
- (void)motionEnded:(UIEventSubtype)motion withEvent:(nullable UIEvent *)event;
- (void)motionCancelled:(UIEventSubtype)motion withEvent:(nullable UIEvent *)event;
@property(nullable, nonatomic, readonly) NSArray<UIKeyCommand *> *keyCommands;
@property(nullable, nonatomic, readonly, strong) __kindof UIView *inputView;
@property(nullable, nonatomic, readonly, strong) __kindof UIView *inputAccessoryView;
- (void)reloadInputViews;
@end

// UIView

typedef NS_OPTIONS(NSUInteger, UIViewAutoresizing) {
	UIViewAutoresizingNone = 0,
	UIViewAutoresizingFlexibleLeftMargin = 1 << 0,
	UIViewAutoresizingFlexibleWidth = 1 << 1,
	UIViewAutoresizingFlexibleRightMargin = 1 << 2,
	UIViewAutoresizingFlexibleTopMargin = 1 << 3,
	UIViewAutoresizingFlexibleHeight = 1 << 4,
	UIViewAutoresizingFlexibleBottomMargin = 1 << 5,
};
typedef NS_ENUM(NSInteger, UIViewContentMode) {
	UIViewContentModeScaleToFill,
	UIViewContentModeScaleAspectFit,
	UIViewContentModeScaleAspectFill,
	UIViewContentModeRedraw,
	UIViewContentModeCenter,
};

@interface UILayoutGuide : NSObject <NSCoding>
@property(nonatomic, readonly) CGRect layoutFrame;
@end

@interface UIView : UIResponder <NSCoding, UITraitEnvironment, UITraitChangeObservable, CALayerDelegate>
@property(class, nonatomic, readonly) Class layerClass;
- (instancetype)initWithFrame:(CGRect)frame;
- (nullable instancetype)initWithCoder:(NSCoder *)coder;
@property(nonatomic, getter=isUserInteractionEnabled) BOOL userInteractionEnabled;
@property(nonatomic) NSInteger tag;
@property(nonatomic, readonly, strong) CALayer *layer;
@property(nonatomic) CGRect frame;
@property(nonatomic) CGRect bounds;
@property(nonatomic) CGPoint center;
@property(nonatomic) CGAffineTransform transform;
@property(nonatomic) CGFloat contentScaleFactor;
@property(nonatomic, getter=isMultipleTouchEnabled) BOOL multipleTouchEnabled;
@property(nonatomic, getter=isExclusiveTouch) BOOL exclusiveTouch;
@property(nonatomic) BOOL autoresizesSubviews;
@property(nonatomic) UIViewAutoresizing autoresizingMask;
@property(nonatomic) UIViewContentMode contentMode;
@property(nullable, nonatomic, readonly) UIView *superview;
@property(nonatomic, readonly, copy) NSArray<__kindof UIView *> *subviews;
@property(nullable, nonatomic, readonly) UIWindow *window;
@property(nonatomic, readonly) UIEdgeInsets safeAreaInsets;
@property(nonatomic, readonly, strong) UILayoutGuide *safeAreaLayoutGuide;
@property(nullable, nonatomic, copy) UIColor *backgroundColor;
@property(nonatomic) CGFloat alpha;
@property(nonatomic, getter=isOpaque) BOOL opaque;
@property(nonatomic, getter=isHidden) BOOL hidden;
@property(nonatomic) BOOL clipsToBounds;
@property(nonatomic) BOOL translatesAutoresizingMaskIntoConstraints;
@property(nullable, nonatomic, copy) NSArray<__kindof UIGestureRecognizer *> *gestureRecognizers;
- (void)removeFromSuperview;
- (void)insertSubview:(UIView *)view atIndex:(NSInteger)index;
- (void)addSubview:(UIView *)view;
- (void)bringSubviewToFront:(UIView *)view;
- (void)sendSubviewToBack:(UIView *)view;
- (void)didMoveToWindow;
- (void)didMoveToSuperview;
- (void)safeAreaInsetsDidChange;
- (BOOL)isDescendantOfView:(UIView *)view;
- (nullable __kindof UIView *)viewWithTag:(NSInteger)tag;
- (void)setNeedsLayout;
- (void)layoutIfNeeded;
- (void)layoutSubviews;
- (void)setNeedsDisplay;
- (void)drawRect:(CGRect)rect;
- (CGPoint)convertPoint:(CGPoint)point toView:(nullable UIView *)view;
- (CGPoint)convertPoint:(CGPoint)point fromView:(nullable UIView *)view;
- (CGRect)convertRect:(CGRect)rect toView:(nullable UIView *)view;
- (CGRect)convertRect:(CGRect)rect fromView:(nullable UIView *)view;
- (nullable UIView *)hitTest:(CGPoint)point withEvent:(nullable UIEvent *)event;
- (BOOL)pointInside:(CGPoint)point withEvent:(nullable UIEvent *)event;
- (void)addGestureRecognizer:(UIGestureRecognizer *)gestureRecognizer;
- (void)removeGestureRecognizer:(UIGestureRecognizer *)gestureRecognizer;
+ (void)animateWithDuration:(NSTimeInterval)duration animations:(void (^)(void))animations;
+ (void)animateWithDuration:(NSTimeInterval)duration animations:(void (^)(void))animations completion:(void (^_Nullable)(BOOL finished))completion;
@end

// UIGestureRecognizer

typedef NS_ENUM(NSInteger, UIGestureRecognizerState) {
	UIGestureRecognizerStatePossible,
	UIGestureRecognizerStateBegan,
	UIGestureRecognizerStateChanged,
	UIGestureRecognizerStateEnded,
	UIGestureRecognizerStateCancelled,
	UIGestureRecognizerStateFailed,
	UIGestureRecognizerStateRecognized = UIGestureRecognizerStateEnded,
};

@protocol UIGestureRecognizerDelegate <NSObject>
@optional
- (BOOL)gestureRecognizerShouldBegin:(UIGestureRecognizer *)gestureRecognizer;
- (BOOL)gestureRecognizer:(UIGestureRecognizer *)gestureRecognizer shouldRecognizeSimultaneouslyWithGestureRecognizer:(UIGestureRecognizer *)otherGestureRecognizer;
- (BOOL)gestureRecognizer:(UIGestureRecognizer *)gestureRecognizer shouldReceiveTouch:(UITouch *)touch;
@end

@interface UIGestureRecognizer : NSObject
- (instancetype)initWithTarget:(nullable id)target action:(nullable SEL)action;
- (void)addTarget:(id)target action:(SEL)action;
@property(nonatomic, readonly) UIGestureRecognizerState state;
@property(nullable, nonatomic, weak) id<UIGestureRecognizerDelegate> delegate;
@property(nonatomic, getter=isEnabled) BOOL enabled;
@property(nullable, nonatomic, readonly) UIView *view;
@property(nonatomic) BOOL cancelsTouchesInView;
@property(nonatomic) BOOL delaysTouchesBegan;
@property(nonatomic) BOOL delaysTouchesEnded;
@property(nonatomic, copy) NSArray<NSNumber *> *allowedTouchTypes;
- (CGPoint)locationInView:(nullable UIView *)view;
@property(nonatomic, readonly) NSUInteger numberOfTouches;
@end

@interface UIPanGestureRecognizer : UIGestureRecognizer
@property(nonatomic) NSUInteger minimumNumberOfTouches;
@property(nonatomic) NSUInteger maximumNumberOfTouches;
@property(nonatomic) NSInteger allowedScrollTypesMask;
- (CGPoint)translationInView:(nullable UIView *)view;
- (void)setTranslation:(CGPoint)translation inView:(nullable UIView *)view;
- (CGPoint)velocityInView:(nullable UIView *)view;
@end

@interface UIHoverGestureRecognizer : UIGestureRecognizer
@end

@interface UITapGestureRecognizer : UIGestureRecognizer
@property(nonatomic) NSUInteger numberOfTapsRequired;
@end

// UIScreen & UIDevice

@interface UIScreenMode : NSObject
@property(readonly, nonatomic) CGSize size;
@end

@interface UIScreen : NSObject <UITraitEnvironment>
@property(class, nonatomic, readonly) UIScreen *mainScreen;
@property(class, nonatomic, readonly) NSArray<UIScreen *> *screens;
@property(nonatomic, readonly) CGRect bounds;
@property(nonatomic, readonly) CGFloat scale;
@property(nonatomic, readonly) CGRect nativeBounds;
@property(nonatomic, readonly) CGFloat nativeScale;
@property(nonatomic, readonly) NSInteger maximumFramesPerSecond;
@property(nonatomic) CGFloat brightness;
@property(nullable, nonatomic, strong) UIScreenMode *currentMode;
@property(nullable, nonatomic, readonly, strong) UIScreenMode *preferredMode;
@property(nonatomic, readonly) CGFloat potentialEDRHeadroom;
@property(nonatomic, readonly) CGFloat currentEDRHeadroom;
@end

typedef NS_ENUM(NSInteger, UIDeviceBatteryState) {
	UIDeviceBatteryStateUnknown,
	UIDeviceBatteryStateUnplugged,
	UIDeviceBatteryStateCharging,
	UIDeviceBatteryStateFull,
};

UIKIT_EXTERN NSNotificationName const UIDeviceOrientationDidChangeNotification;
UIKIT_EXTERN NSNotificationName const UIDeviceBatteryStateDidChangeNotification;
UIKIT_EXTERN NSNotificationName const UIDeviceBatteryLevelDidChangeNotification;

@interface UIDevice : NSObject
@property(class, nonatomic, readonly) UIDevice *currentDevice;
@property(nonatomic, readonly, strong) NSString *name;
@property(nonatomic, readonly, strong) NSString *model;
@property(nonatomic, readonly, strong) NSString *localizedModel;
@property(nonatomic, readonly, strong) NSString *systemName;
@property(nonatomic, readonly, strong) NSString *systemVersion;
@property(nonatomic, readonly) UIDeviceOrientation orientation;
@property(nullable, nonatomic, readonly, strong) NSUUID *identifierForVendor;
@property(nonatomic, readonly, getter=isGeneratingDeviceOrientationNotifications) BOOL generatesDeviceOrientationNotifications;
- (void)beginGeneratingDeviceOrientationNotifications;
- (void)endGeneratingDeviceOrientationNotifications;
@property(nonatomic, getter=isBatteryMonitoringEnabled) BOOL batteryMonitoringEnabled;
@property(nonatomic, readonly) UIDeviceBatteryState batteryState;
@property(nonatomic, readonly) float batteryLevel;
@property(nonatomic, readonly) UIUserInterfaceIdiom userInterfaceIdiom;
@end
#define UI_USER_INTERFACE_IDIOM() ([[UIDevice currentDevice] userInterfaceIdiom])

// UIViewController

typedef NS_ENUM(NSInteger, UIModalPresentationStyle) {
	UIModalPresentationFullScreen = 0,
	UIModalPresentationPageSheet,
	UIModalPresentationFormSheet,
	UIModalPresentationCurrentContext,
	UIModalPresentationCustom,
	UIModalPresentationOverFullScreen,
	UIModalPresentationOverCurrentContext,
	UIModalPresentationPopover,
	UIModalPresentationNone = -1,
	UIModalPresentationAutomatic = -2,
};
typedef NS_ENUM(NSInteger, UIStatusBarStyle) {
	UIStatusBarStyleDefault = 0,
	UIStatusBarStyleLightContent = 1,
	UIStatusBarStyleDarkContent = 3,
};

@protocol UIViewControllerTransitionCoordinatorContext <NSObject>
@property(nonatomic, readonly, getter=isAnimated) BOOL animated;
@property(nonatomic, readonly) NSTimeInterval transitionDuration;
@property(nonatomic, readonly) UIView *containerView;
@end

@protocol UIViewControllerTransitionCoordinator <UIViewControllerTransitionCoordinatorContext>
- (BOOL)animateAlongsideTransition:(void (^_Nullable)(id<UIViewControllerTransitionCoordinatorContext> context))animation completion:(void (^_Nullable)(id<UIViewControllerTransitionCoordinatorContext> context))completion;
@end

@protocol UIContentContainer <NSObject>
- (void)viewWillTransitionToSize:(CGSize)size withTransitionCoordinator:(id<UIViewControllerTransitionCoordinator>)coordinator;
@end

@interface UIViewController : UIResponder <NSCoding, UITraitEnvironment, UITraitChangeObservable, UIContentContainer>
- (instancetype)initWithNibName:(nullable NSString *)nibNameOrNil bundle:(nullable NSBundle *)nibBundleOrNil;
- (nullable instancetype)initWithCoder:(NSCoder *)coder;
@property(null_resettable, nonatomic, strong) UIView *view;
- (void)loadView;
- (void)loadViewIfNeeded;
@property(nullable, nonatomic, readonly, strong) UIView *viewIfLoaded;
@property(nonatomic, readonly, getter=isViewLoaded) BOOL viewLoaded;
- (void)viewDidLoad;
- (void)viewWillAppear:(BOOL)animated;
- (void)viewDidAppear:(BOOL)animated;
- (void)viewWillDisappear:(BOOL)animated;
- (void)viewDidDisappear:(BOOL)animated;
- (void)viewWillLayoutSubviews;
- (void)viewDidLayoutSubviews;
- (void)viewSafeAreaInsetsDidChange;
- (void)didReceiveMemoryWarning;
@property(nullable, nonatomic, copy) NSString *title;
@property(nullable, nonatomic, weak, readonly) UIViewController *parentViewController;
@property(nullable, nonatomic, readonly) UIViewController *presentedViewController;
@property(nullable, nonatomic, readonly) UIViewController *presentingViewController;
@property(nonatomic, assign) UIModalPresentationStyle modalPresentationStyle;
@property(nonatomic, getter=isModalInPresentation) BOOL modalInPresentation;
- (void)presentViewController:(UIViewController *)viewControllerToPresent animated:(BOOL)flag completion:(void (^_Nullable)(void))completion;
- (void)dismissViewControllerAnimated:(BOOL)flag completion:(void (^_Nullable)(void))completion;
@property(nonatomic, readonly) BOOL shouldAutorotate;
@property(nonatomic, readonly) UIInterfaceOrientationMask supportedInterfaceOrientations;
@property(nonatomic, readonly) UIInterfaceOrientation preferredInterfaceOrientationForPresentation;
+ (void)attemptRotationToDeviceOrientation;
- (void)setNeedsUpdateOfSupportedInterfaceOrientations;
@property(nonatomic, readonly) BOOL prefersStatusBarHidden;
@property(nonatomic, readonly) UIStatusBarStyle preferredStatusBarStyle;
- (void)setNeedsStatusBarAppearanceUpdate;
@property(nonatomic, readonly) BOOL prefersHomeIndicatorAutoHidden;
- (void)setNeedsUpdateOfHomeIndicatorAutoHidden;
@property(nonatomic, readonly) UIRectEdge preferredScreenEdgesDeferringSystemGestures;
- (void)setNeedsUpdateOfScreenEdgesDeferringSystemGestures;
@property(nonatomic, readonly) BOOL prefersPointerLocked;
- (void)setNeedsUpdateOfPrefersPointerLocked;
@property(nullable, nonatomic, readonly, strong) UIStoryboard *storyboard;
- (void)addChildViewController:(UIViewController *)childController;
- (void)removeFromParentViewController;
- (void)willMoveToParentViewController:(nullable UIViewController *)parent;
- (void)didMoveToParentViewController:(nullable UIViewController *)parent;
@end

@interface UIStoryboard : NSObject
+ (UIStoryboard *)storyboardWithName:(NSString *)name bundle:(nullable NSBundle *)storyboardBundleOrNil;
- (nullable __kindof UIViewController *)instantiateInitialViewController;
- (__kindof UIViewController *)instantiateViewControllerWithIdentifier:(NSString *)identifier;
@end

// UIWindow

typedef CGFloat UIWindowLevel NS_TYPED_EXTENSIBLE_ENUM;
UIKIT_EXTERN const UIWindowLevel UIWindowLevelNormal;
UIKIT_EXTERN const UIWindowLevel UIWindowLevelAlert;

@interface UIWindow : UIView
- (instancetype)initWithWindowScene:(UIWindowScene *)windowScene;
- (instancetype)initWithFrame:(CGRect)frame;
@property(nullable, nonatomic, weak) UIWindowScene *windowScene;
@property(nonatomic, strong) UIScreen *screen;
@property(nonatomic) UIWindowLevel windowLevel;
@property(nonatomic, readonly, getter=isKeyWindow) BOOL keyWindow;
@property(nullable, nonatomic, strong) UIViewController *rootViewController;
- (void)makeKeyWindow;
- (void)makeKeyAndVisible;
- (void)sendEvent:(UIEvent *)event;
@end

// Keyboard & text input

UIKIT_EXTERN NSNotificationName const UIKeyboardWillShowNotification;
UIKIT_EXTERN NSNotificationName const UIKeyboardDidShowNotification;
UIKIT_EXTERN NSNotificationName const UIKeyboardWillHideNotification;
UIKIT_EXTERN NSNotificationName const UIKeyboardDidHideNotification;
UIKIT_EXTERN NSNotificationName const UIKeyboardWillChangeFrameNotification;
UIKIT_EXTERN NSNotificationName const UIKeyboardDidChangeFrameNotification;
UIKIT_EXTERN NSString *const UIKeyboardFrameBeginUserInfoKey;
UIKIT_EXTERN NSString *const UIKeyboardFrameEndUserInfoKey;
UIKIT_EXTERN NSString *const UIKeyboardAnimationDurationUserInfoKey;
UIKIT_EXTERN NSString *const UIKeyboardAnimationCurveUserInfoKey;

typedef NS_ENUM(NSInteger, UIKeyboardType) {
	UIKeyboardTypeDefault,
	UIKeyboardTypeASCIICapable,
	UIKeyboardTypeNumbersAndPunctuation,
	UIKeyboardTypeURL,
	UIKeyboardTypeNumberPad,
	UIKeyboardTypePhonePad,
	UIKeyboardTypeNamePhonePad,
	UIKeyboardTypeEmailAddress,
	UIKeyboardTypeDecimalPad,
	UIKeyboardTypeTwitter,
	UIKeyboardTypeWebSearch,
	UIKeyboardTypeASCIICapableNumberPad,
	UIKeyboardTypeAlphabet = UIKeyboardTypeASCIICapable,
};
typedef NS_ENUM(NSInteger, UITextAutocapitalizationType) {
	UITextAutocapitalizationTypeNone,
	UITextAutocapitalizationTypeWords,
	UITextAutocapitalizationTypeSentences,
	UITextAutocapitalizationTypeAllCharacters,
};
typedef NS_ENUM(NSInteger, UITextAutocorrectionType) {
	UITextAutocorrectionTypeDefault,
	UITextAutocorrectionTypeNo,
	UITextAutocorrectionTypeYes,
};
typedef NS_ENUM(NSInteger, UITextSpellCheckingType) {
	UITextSpellCheckingTypeDefault,
	UITextSpellCheckingTypeNo,
	UITextSpellCheckingTypeYes,
};
typedef NS_ENUM(NSInteger, UIKeyboardAppearance) {
	UIKeyboardAppearanceDefault,
	UIKeyboardAppearanceDark,
	UIKeyboardAppearanceLight,
};
typedef NS_ENUM(NSInteger, UIReturnKeyType) {
	UIReturnKeyDefault,
	UIReturnKeyGo,
	UIReturnKeyGoogle,
	UIReturnKeyJoin,
	UIReturnKeyNext,
	UIReturnKeyRoute,
	UIReturnKeySearch,
	UIReturnKeySend,
	UIReturnKeyYahoo,
	UIReturnKeyDone,
	UIReturnKeyEmergencyCall,
	UIReturnKeyContinue,
};

typedef NSString *UITextContentType NS_TYPED_ENUM;
UIKIT_EXTERN UITextContentType const UITextContentTypeName;
UIKIT_EXTERN UITextContentType const UITextContentTypeURL;
UIKIT_EXTERN UITextContentType const UITextContentTypeEmailAddress;
UIKIT_EXTERN UITextContentType const UITextContentTypeTelephoneNumber;
UIKIT_EXTERN UITextContentType const UITextContentTypeUsername;
UIKIT_EXTERN UITextContentType const UITextContentTypePassword;
UIKIT_EXTERN UITextContentType const UITextContentTypeNewPassword;
UIKIT_EXTERN UITextContentType const UITextContentTypeOneTimeCode;

@protocol UITextInputTraits <NSObject>
@optional
@property(nonatomic) UITextAutocapitalizationType autocapitalizationType;
@property(nonatomic) UITextAutocorrectionType autocorrectionType;
@property(nonatomic) UITextSpellCheckingType spellCheckingType;
@property(nonatomic) UIKeyboardType keyboardType;
@property(nonatomic) UIKeyboardAppearance keyboardAppearance;
@property(nonatomic) UIReturnKeyType returnKeyType;
@property(nonatomic) BOOL enablesReturnKeyAutomatically;
@property(nonatomic, getter=isSecureTextEntry) BOOL secureTextEntry;
@property(null_unspecified, nonatomic, copy) UITextContentType textContentType;
@end

@protocol UIKeyInput <UITextInputTraits>
@property(nonatomic, readonly) BOOL hasText;
- (void)insertText:(NSString *)text;
- (void)deleteBackward;
@end

@interface UITextPosition : NSObject
@end
@interface UITextRange : NSObject
@property(nonatomic, readonly, getter=isEmpty) BOOL empty;
@property(nonatomic, readonly) UITextPosition *start;
@property(nonatomic, readonly) UITextPosition *end;
@end

@protocol UIScrollViewDelegate <NSObject>
@end

@interface UIScrollView : UIView <NSCoding>
@property(nonatomic) CGPoint contentOffset;
@property(nonatomic) CGSize contentSize;
@property(nonatomic) UIEdgeInsets contentInset;
@property(nonatomic, getter=isScrollEnabled) BOOL scrollEnabled;
@end

@class UITextView;
@protocol UITextViewDelegate <NSObject, UIScrollViewDelegate>
@optional
- (BOOL)textViewShouldBeginEditing:(UITextView *)textView;
- (BOOL)textViewShouldEndEditing:(UITextView *)textView;
- (void)textViewDidBeginEditing:(UITextView *)textView;
- (void)textViewDidEndEditing:(UITextView *)textView;
- (BOOL)textView:(UITextView *)textView shouldChangeTextInRange:(NSRange)range replacementText:(NSString *)text;
- (void)textViewDidChange:(UITextView *)textView;
- (void)textViewDidChangeSelection:(UITextView *)textView;
@end

UIKIT_EXTERN NSNotificationName const UITextViewTextDidBeginEditingNotification;
UIKIT_EXTERN NSNotificationName const UITextViewTextDidChangeNotification;
UIKIT_EXTERN NSNotificationName const UITextViewTextDidEndEditingNotification;

@interface NSTextContainer : NSObject <NSSecureCoding>
@end

@interface UITextView : UIScrollView <UIKeyInput>
- (instancetype)initWithFrame:(CGRect)frame textContainer:(nullable NSTextContainer *)textContainer;
@property(nullable, nonatomic, weak) id<UITextViewDelegate> delegate;
@property(null_resettable, nonatomic, copy) NSString *text;
@property(nullable, nonatomic, strong) UIFont *font;
@property(nullable, nonatomic, strong) UIColor *textColor;
@property(nonatomic) NSRange selectedRange;
@property(nullable, readwrite, copy) UITextRange *selectedTextRange;
@property(nonatomic, readonly) UITextPosition *beginningOfDocument;
@property(nonatomic, readonly) UITextPosition *endOfDocument;
@property(nonatomic, getter=isEditable) BOOL editable;
@property(nonatomic, getter=isSelectable) BOOL selectable;
@property(nullable, readwrite, strong) UIView *inputView;
@property(nullable, readwrite, strong) UIView *inputAccessoryView;
- (nullable UITextRange *)textRangeFromPosition:(UITextPosition *)fromPosition toPosition:(UITextPosition *)toPosition;
- (nullable UITextPosition *)positionFromPosition:(UITextPosition *)position offset:(NSInteger)offset;
- (NSInteger)offsetFromPosition:(UITextPosition *)from toPosition:(UITextPosition *)toPosition;
@end

// UIPasteboard

@interface UIPasteboard : NSObject
@property(class, nonatomic, readonly) UIPasteboard *generalPasteboard;
@property(nullable, nonatomic, copy) NSString *string;
@property(nullable, nonatomic, copy) NSArray<NSString *> *strings;
@property(nullable, nonatomic, copy) UIImage *image;
@property(nullable, nonatomic, copy) NSURL *URL;
@property(nonatomic, readonly) BOOL hasStrings;
@property(nonatomic, readonly) BOOL hasImages;
@property(nonatomic, readonly) BOOL hasURLs;
@end

// UIAlertController

typedef NS_ENUM(NSInteger, UIAlertActionStyle) {
	UIAlertActionStyleDefault = 0,
	UIAlertActionStyleCancel,
	UIAlertActionStyleDestructive,
};
typedef NS_ENUM(NSInteger, UIAlertControllerStyle) {
	UIAlertControllerStyleActionSheet = 0,
	UIAlertControllerStyleAlert,
};

@interface UIAlertAction : NSObject <NSCopying>
+ (instancetype)actionWithTitle:(nullable NSString *)title style:(UIAlertActionStyle)style handler:(void (^_Nullable)(UIAlertAction *action))handler;
@property(nullable, nonatomic, readonly) NSString *title;
@property(nonatomic, readonly) UIAlertActionStyle style;
@property(nonatomic, getter=isEnabled) BOOL enabled;
@end

@interface UIAlertController : UIViewController
+ (instancetype)alertControllerWithTitle:(nullable NSString *)title message:(nullable NSString *)message preferredStyle:(UIAlertControllerStyle)preferredStyle;
- (void)addAction:(UIAlertAction *)action;
@property(nonatomic, readonly) NSArray<UIAlertAction *> *actions;
@property(nullable, nonatomic, copy) NSString *message;
@end

// UIScene

typedef NSString *UISceneSessionRole NS_TYPED_ENUM;
UIKIT_EXTERN UISceneSessionRole const UIWindowSceneSessionRoleApplication;
UIKIT_EXTERN UISceneSessionRole const UIWindowSceneSessionRoleExternalDisplayNonInteractive;

typedef NS_ENUM(NSInteger, UISceneActivationState) {
	UISceneActivationStateUnattached = -1,
	UISceneActivationStateForegroundActive,
	UISceneActivationStateForegroundInactive,
	UISceneActivationStateBackground,
};

@interface UISceneConfiguration : NSObject <NSCopying, NSSecureCoding>
+ (instancetype)configurationWithName:(nullable NSString *)name sessionRole:(UISceneSessionRole)sessionRole;
- (instancetype)initWithName:(nullable NSString *)name sessionRole:(UISceneSessionRole)sessionRole;
@property(nonatomic, nullable, readonly) NSString *name;
@property(nonatomic, readonly) UISceneSessionRole role;
@property(nonatomic, nullable) Class sceneClass;
@property(nonatomic, nullable) Class delegateClass;
@property(nonatomic, nullable, strong) UIStoryboard *storyboard;
@end

@interface UISceneSession : NSObject <NSSecureCoding>
@property(nonatomic, readonly, nullable) UIScene *scene;
@property(nonatomic, readonly) UISceneSessionRole role;
@property(nonatomic, readonly, copy) UISceneConfiguration *configuration;
@property(nonatomic, readonly) NSString *persistentIdentifier;
@end

@interface UIOpenURLContext : NSObject
@property(nonatomic, readonly, copy) NSURL *URL;
@end

@interface UIApplicationShortcutItem : NSObject <NSCopying, NSMutableCopying>
@property(nonatomic, copy, readonly) NSString *type;
@property(nonatomic, copy, readonly) NSString *localizedTitle;
@end

@interface UISceneConnectionOptions : NSObject
@property(nonatomic, readonly, copy) NSSet<UIOpenURLContext *> *URLContexts;
@property(nullable, nonatomic, readonly) NSString *sourceApplication;
@property(nonatomic, readonly, copy) NSSet<NSUserActivity *> *userActivities;
@property(nullable, nonatomic, readonly) UIApplicationShortcutItem *shortcutItem;
@end

@protocol UISceneDelegate <NSObject>
@optional
- (void)scene:(UIScene *)scene willConnectToSession:(UISceneSession *)session options:(UISceneConnectionOptions *)connectionOptions;
- (void)sceneDidDisconnect:(UIScene *)scene;
- (void)sceneDidBecomeActive:(UIScene *)scene;
- (void)sceneWillResignActive:(UIScene *)scene;
- (void)sceneWillEnterForeground:(UIScene *)scene;
- (void)sceneDidEnterBackground:(UIScene *)scene;
- (void)scene:(UIScene *)scene openURLContexts:(NSSet<UIOpenURLContext *> *)URLContexts;
- (void)scene:(UIScene *)scene continueUserActivity:(NSUserActivity *)userActivity;
@end

@interface UIScene : UIResponder
@property(nonatomic, readonly) UISceneSession *session;
@property(nullable, nonatomic, strong) id<UISceneDelegate> delegate;
@property(nonatomic, readonly) UISceneActivationState activationState;
@property(null_resettable, nonatomic, copy) NSString *title;
@end

@interface UIWindowSceneGeometryPreferences : NSObject
@end
@interface UIWindowSceneGeometryPreferencesIOS : UIWindowSceneGeometryPreferences
- (instancetype)initWithInterfaceOrientations:(UIInterfaceOrientationMask)interfaceOrientations;
@property(nonatomic, assign) UIInterfaceOrientationMask interfaceOrientations;
@end

@interface UIWindowSceneGeometry : NSObject <NSCopying>
@property(nonatomic, readonly) UIInterfaceOrientation interfaceOrientation;
@end

@protocol UIWindowSceneDelegate <UISceneDelegate>
@optional
@property(nullable, nonatomic, strong) UIWindow *window;
- (void)windowScene:(UIWindowScene *)windowScene performActionForShortcutItem:(UIApplicationShortcutItem *)shortcutItem completionHandler:(void (^)(BOOL succeeded))completionHandler;
@end

@interface UIWindowScene : UIScene
@property(nonatomic, readonly) UIScreen *screen;
@property(nonatomic, readonly) UIInterfaceOrientation interfaceOrientation;
@property(nonatomic, readonly) UITraitCollection *traitCollection;
@property(nonatomic, readonly) NSArray<UIWindow *> *windows;
@property(nullable, nonatomic, readonly) UIWindow *keyWindow;
@property(nonatomic, readonly) UIWindowSceneGeometry *effectiveGeometry;
- (void)requestGeometryUpdateWithPreferences:(UIWindowSceneGeometryPreferences *)geometryPreferences errorHandler:(nullable void (^)(NSError *error))errorHandler;
@end

// UIApplication

typedef NSString *UIApplicationLaunchOptionsKey NS_TYPED_ENUM;
typedef NSString *UIApplicationOpenURLOptionsKey NS_TYPED_ENUM;
typedef NSString *UIApplicationOpenExternalURLOptionsKey NS_TYPED_ENUM;
typedef NSString *UIApplicationExtensionPointIdentifier NS_TYPED_ENUM;
UIKIT_EXTERN UIApplicationLaunchOptionsKey const UIApplicationLaunchOptionsURLKey;
UIKIT_EXTERN UIApplicationExtensionPointIdentifier const UIApplicationKeyboardExtensionPointIdentifier;
UIKIT_EXTERN NSString *const UIApplicationOpenSettingsURLString;

typedef NS_ENUM(NSInteger, UIApplicationState) {
	UIApplicationStateActive,
	UIApplicationStateInactive,
	UIApplicationStateBackground,
};

UIKIT_EXTERN NSNotificationName const UIApplicationDidEnterBackgroundNotification;
UIKIT_EXTERN NSNotificationName const UIApplicationWillEnterForegroundNotification;
UIKIT_EXTERN NSNotificationName const UIApplicationDidFinishLaunchingNotification;
UIKIT_EXTERN NSNotificationName const UIApplicationDidBecomeActiveNotification;
UIKIT_EXTERN NSNotificationName const UIApplicationWillResignActiveNotification;
UIKIT_EXTERN NSNotificationName const UIApplicationDidReceiveMemoryWarningNotification;
UIKIT_EXTERN NSNotificationName const UIApplicationWillTerminateNotification;

@protocol UIUserActivityRestoring <NSObject>
- (void)restoreUserActivityState:(NSUserActivity *)userActivity;
@end

@protocol UIApplicationDelegate <NSObject>
@optional
- (void)applicationDidFinishLaunching:(UIApplication *)application;
- (BOOL)application:(UIApplication *)application willFinishLaunchingWithOptions:(nullable NSDictionary<UIApplicationLaunchOptionsKey, id> *)launchOptions;
- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(nullable NSDictionary<UIApplicationLaunchOptionsKey, id> *)launchOptions;
- (void)applicationDidBecomeActive:(UIApplication *)application;
- (void)applicationWillResignActive:(UIApplication *)application;
- (void)applicationDidEnterBackground:(UIApplication *)application;
- (void)applicationWillEnterForeground:(UIApplication *)application;
- (void)applicationDidReceiveMemoryWarning:(UIApplication *)application;
- (void)applicationWillTerminate:(UIApplication *)application;
- (void)applicationSignificantTimeChange:(UIApplication *)application;
- (BOOL)application:(UIApplication *)app openURL:(NSURL *)url options:(NSDictionary<UIApplicationOpenURLOptionsKey, id> *)options;
- (void)application:(UIApplication *)application didRegisterForRemoteNotificationsWithDeviceToken:(NSData *)deviceToken;
- (void)application:(UIApplication *)application didFailToRegisterForRemoteNotificationsWithError:(NSError *)error;
- (void)application:(UIApplication *)application didReceiveRemoteNotification:(NSDictionary *)userInfo fetchCompletionHandler:(void (^)(NSUInteger result))completionHandler;
- (void)application:(UIApplication *)application performActionForShortcutItem:(UIApplicationShortcutItem *)shortcutItem completionHandler:(void (^)(BOOL succeeded))completionHandler;
- (void)application:(UIApplication *)application handleEventsForBackgroundURLSession:(NSString *)identifier completionHandler:(void (^)(void))completionHandler;
- (BOOL)application:(UIApplication *)application shouldAllowExtensionPointIdentifier:(UIApplicationExtensionPointIdentifier)extensionPointIdentifier;
- (BOOL)application:(UIApplication *)application willContinueUserActivityWithType:(NSString *)userActivityType;
- (BOOL)application:(UIApplication *)application continueUserActivity:(NSUserActivity *)userActivity restorationHandler:(void (^)(NSArray<id<UIUserActivityRestoring>> *_Nullable restorableObjects))restorationHandler;
- (void)application:(UIApplication *)application didFailToContinueUserActivityWithType:(NSString *)userActivityType error:(NSError *)error;
- (void)application:(UIApplication *)application didUpdateUserActivity:(NSUserActivity *)userActivity;
- (UIInterfaceOrientationMask)application:(UIApplication *)application supportedInterfaceOrientationsForWindow:(nullable UIWindow *)window;
- (UISceneConfiguration *)application:(UIApplication *)application configurationForConnectingSceneSession:(UISceneSession *)connectingSceneSession options:(UISceneConnectionOptions *)options;
- (void)application:(UIApplication *)application didDiscardSceneSessions:(NSSet<UISceneSession *> *)sceneSessions;
- (void)applicationProtectedDataWillBecomeUnavailable:(UIApplication *)application;
- (void)applicationProtectedDataDidBecomeAvailable:(UIApplication *)application;
- (BOOL)application:(UIApplication *)application shouldSaveSecureApplicationState:(NSCoder *)coder;
- (BOOL)application:(UIApplication *)application shouldRestoreSecureApplicationState:(NSCoder *)coder;
- (nullable UIViewController *)application:(UIApplication *)application viewControllerWithRestorationIdentifierPath:(NSArray<NSString *> *)identifierComponents coder:(NSCoder *)coder;
- (void)application:(UIApplication *)application willEncodeRestorableStateWithCoder:(NSCoder *)coder;
- (void)application:(UIApplication *)application didDecodeRestorableStateWithCoder:(NSCoder *)coder;
- (void)application:(UIApplication *)application handleWatchKitExtensionRequest:(nullable NSDictionary *)userInfo reply:(void (^)(NSDictionary *_Nullable replyInfo))reply;
- (void)applicationShouldRequestHealthAuthorization:(UIApplication *)application;
- (nullable id)application:(UIApplication *)application handlerForIntent:(INIntent *)intent;
- (void)application:(UIApplication *)application userDidAcceptCloudKitShareWithMetadata:(CKShareMetadata *)cloudKitShareMetadata;
@property(nullable, nonatomic, strong) UIWindow *window;
@end

@interface UIApplication : UIResponder
@property(class, nonatomic, readonly) UIApplication *sharedApplication;
@property(nullable, nonatomic, assign) id<UIApplicationDelegate> delegate;
@property(nonatomic, getter=isIdleTimerDisabled) BOOL idleTimerDisabled;
@property(nonatomic, readonly) UIApplicationState applicationState;
@property(nonatomic, readonly) NSSet<UIScene *> *connectedScenes;
@property(nonatomic, readonly) NSSet<UISceneSession *> *openSessions;
@property(nonatomic, readonly) BOOL supportsMultipleScenes;
@property(nullable, nonatomic, readonly) UIWindow *keyWindow;
@property(nonatomic, readonly) NSArray<__kindof UIWindow *> *windows;
@property(nonatomic, readonly) UIInterfaceOrientation statusBarOrientation;
@property(nonatomic, readonly) CGRect statusBarFrame;
@property(nonatomic) NSInteger applicationIconBadgeNumber;
- (BOOL)canOpenURL:(NSURL *)url;
- (BOOL)openURL:(NSURL *)url;
- (void)openURL:(NSURL *)url options:(NSDictionary<UIApplicationOpenExternalURLOptionsKey, id> *)options completionHandler:(void (^_Nullable)(BOOL success))completion;
- (void)sendEvent:(UIEvent *)event;
- (BOOL)sendAction:(SEL)action to:(nullable id)target from:(nullable id)sender forEvent:(nullable UIEvent *)event;
- (void)registerForRemoteNotifications;
@end

UIKIT_EXTERN int UIApplicationMain(int argc, char *_Nullable argv[_Nonnull], NSString *_Nullable principalClassName, NSString *_Nullable delegateClassName);

// UIFeedbackGenerator

typedef NS_ENUM(NSInteger, UIImpactFeedbackStyle) {
	UIImpactFeedbackStyleLight,
	UIImpactFeedbackStyleMedium,
	UIImpactFeedbackStyleHeavy,
	UIImpactFeedbackStyleSoft,
	UIImpactFeedbackStyleRigid,
};

@interface UIFeedbackGenerator : NSObject
- (void)prepare;
@end

@interface UIImpactFeedbackGenerator : UIFeedbackGenerator
- (instancetype)initWithStyle:(UIImpactFeedbackStyle)style;
- (void)impactOccurred;
- (void)impactOccurredWithIntensity:(CGFloat)intensity;
@end

NS_ASSUME_NONNULL_END

#endif
