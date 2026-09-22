// AppKit declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_APPKIT_H
#define GD_APPKIT_H

#import <Foundation/Foundation.h>
#import <QuartzCore/QuartzCore.h>
#import <CoreGraphics/CoreGraphics.h>
#import <ApplicationServices/ApplicationServices.h>
#import <UniformTypeIdentifiers/UniformTypeIdentifiers.h>
#import <OpenGL/OpenGL.h>

#if defined(__cplusplus)
#define APPKIT_EXTERN extern "C"
#else
#define APPKIT_EXTERN extern
#endif
#define IBOutlet
#define IBAction void

NS_ASSUME_NONNULL_BEGIN

@class NSView, NSWindow, NSScreen, NSEvent, NSMenu, NSMenuItem, NSImage, NSColor, NSFont, NSCursor, NSResponder, NSApplication, NSRunningApplication, NSPasteboard, NSPasteboardItem, NSTrackingArea, NSAppearance, NSWindowTab, NSTextInputContext, NSTouch, NSGraphicsContext, NSBezierPath, NSLayoutConstraint, NSLayoutXAxisAnchor, NSLayoutYAxisAnchor, NSLayoutDimension, NSOpenGLContext, NSOpenGLPixelFormat, NSSpeechSynthesizer, NSStatusItem, NSStatusBar, NSButton, NSAlert, NSSavePanel, NSPopUpButton, NSTextField, NSCharacterSet, NSColorSpace, NSSet;

// Geometry

typedef CGPoint NSPoint;
typedef CGSize NSSize;
typedef CGRect NSRect;
typedef NSPoint *NSPointPointer;
typedef NSRect *NSRectPointer;
typedef struct NSEdgeInsets {
	CGFloat top, left, bottom, right;
} NSEdgeInsets;
#define NSZeroPoint CGPointZero
#define NSZeroSize CGSizeZero
#define NSZeroRect CGRectZero
NS_INLINE NSPoint NSMakePoint(CGFloat x, CGFloat y) {
	return CGPointMake(x, y);
}
NS_INLINE NSSize NSMakeSize(CGFloat w, CGFloat h) {
	return CGSizeMake(w, h);
}
NS_INLINE NSRect NSMakeRect(CGFloat x, CGFloat y, CGFloat w, CGFloat h) {
	return CGRectMake(x, y, w, h);
}
NS_INLINE CGFloat NSMinX(NSRect r) { return r.origin.x; }
NS_INLINE CGFloat NSMinY(NSRect r) { return r.origin.y; }
NS_INLINE CGFloat NSMaxX(NSRect r) { return r.origin.x + r.size.width; }
NS_INLINE CGFloat NSMaxY(NSRect r) { return r.origin.y + r.size.height; }
NS_INLINE CGFloat NSWidth(NSRect r) { return r.size.width; }
NS_INLINE CGFloat NSHeight(NSRect r) { return r.size.height; }
NS_INLINE BOOL NSEqualRects(NSRect a, NSRect b) { return CGRectEqualToRect(a, b); }
NS_INLINE BOOL NSEqualSizes(NSSize a, NSSize b) { return CGSizeEqualToSize(a, b); }
NS_INLINE BOOL NSEqualPoints(NSPoint a, NSPoint b) { return CGPointEqualToPoint(a, b); }
NS_INLINE BOOL NSPointInRect(NSPoint p, NSRect r) { return CGRectContainsPoint(r, p); }
NS_INLINE BOOL NSMouseInRect(NSPoint aPoint, NSRect aRect, BOOL flipped) {
	if (flipped) {
		return aPoint.x >= NSMinX(aRect) && aPoint.x < NSMaxX(aRect) && aPoint.y > NSMinY(aRect) && aPoint.y <= NSMaxY(aRect);
	}
	return aPoint.x >= NSMinX(aRect) && aPoint.x < NSMaxX(aRect) && aPoint.y >= NSMinY(aRect) && aPoint.y < NSMaxY(aRect);
}
NS_INLINE NSRect NSIntersectionRect(NSRect a, NSRect b) { return CGRectIntersection(a, b); }
NS_INLINE BOOL NSIsEmptyRect(NSRect r) { return CGRectIsEmpty(r); }

@interface NSValue (NSValueGeometryExtensions)
+ (NSValue *)valueWithPoint:(NSPoint)point;
+ (NSValue *)valueWithSize:(NSSize)size;
+ (NSValue *)valueWithRect:(NSRect)rect;
@property(readonly) NSPoint pointValue;
@property(readonly) NSSize sizeValue;
@property(readonly) NSRect rectValue;
@end

// Foundation additions AppKit's callers count on

@interface NSCharacterSet : NSObject <NSCopying, NSMutableCopying, NSSecureCoding>
@property(class, readonly, copy) NSCharacterSet *controlCharacterSet;
@property(class, readonly, copy) NSCharacterSet *whitespaceCharacterSet;
@property(class, readonly, copy) NSCharacterSet *whitespaceAndNewlineCharacterSet;
@property(class, readonly, copy) NSCharacterSet *newlineCharacterSet;
@property(class, readonly, copy) NSCharacterSet *alphanumericCharacterSet;
@property(class, readonly, copy) NSCharacterSet *URLFragmentAllowedCharacterSet;
@property(class, readonly, copy) NSCharacterSet *URLPathAllowedCharacterSet;
@property(class, readonly, copy) NSCharacterSet *URLQueryAllowedCharacterSet;
+ (NSCharacterSet *)characterSetWithCharactersInString:(NSString *)aString;
- (BOOL)characterIsMember:(unichar)aCharacter;
@end

@interface NSString (NSCharacterSetAdditions)
- (NSRange)rangeOfCharacterFromSet:(NSCharacterSet *)searchSet;
- (NSString *)stringByTrimmingCharactersInSet:(NSCharacterSet *)set;
@property(readonly, copy) NSString *precomposedStringWithCanonicalMapping;
@property(readonly, copy) NSString *decomposedStringWithCanonicalMapping;
@end

typedef NS_OPTIONS(NSUInteger, NSURLBookmarkResolutionOptions) {
	NSURLBookmarkResolutionWithoutUI = (1UL << 8),
	NSURLBookmarkResolutionWithoutMounting = (1UL << 9),
	NSURLBookmarkResolutionWithSecurityScope = (1UL << 10),
};
typedef NS_OPTIONS(NSUInteger, NSURLBookmarkCreationOptions) {
	NSURLBookmarkCreationMinimalBookmark = (1UL << 9),
	NSURLBookmarkCreationSuitableForBookmarkFile = (1UL << 10),
	NSURLBookmarkCreationWithSecurityScope = (1UL << 11),
	NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess = (1UL << 12),
};
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeNameKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeLocalizedNameKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeIsRemovableKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeIsEjectableKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLIsDirectoryKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLIsPackageKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLIsHiddenKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLLocalizedNameKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeURLKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLIsSystemImmutableKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLIsUserImmutableKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeSupportsCaseSensitiveNamesKey;
FOUNDATION_EXPORT NSURLResourceKey const NSURLVolumeSupportsCasePreservedNamesKey;

@interface NSURL (NSURLBookmarks)
- (nullable NSData *)bookmarkDataWithOptions:(NSURLBookmarkCreationOptions)options includingResourceValuesForKeys:(nullable NSArray<NSURLResourceKey> *)keys relativeToURL:(nullable NSURL *)relativeURL error:(NSError **)error;
+ (nullable instancetype)URLByResolvingBookmarkData:(NSData *)bookmarkData options:(NSURLBookmarkResolutionOptions)options relativeToURL:(nullable NSURL *)relativeURL bookmarkDataIsStale:(BOOL *_Nullable)isStale error:(NSError **)error;
- (BOOL)getResourceValue:(out id _Nullable *_Nonnull)value forKey:(NSURLResourceKey)key error:(out NSError **)error;
- (nullable NSDictionary<NSURLResourceKey, id> *)resourceValuesForKeys:(NSArray<NSURLResourceKey> *)keys error:(NSError **)error;
@property(nullable, readonly, copy) NSURL *URLByStandardizingPath;
@property(nullable, readonly, copy) NSURL *URLByResolvingSymlinksInPath;
@property(nullable, readonly, copy) NSURL *URLByDeletingLastPathComponent;
@end

typedef NS_OPTIONS(NSUInteger, NSVolumeEnumerationOptions) {
	NSVolumeEnumerationSkipHiddenVolumes = 1UL << 1,
	NSVolumeEnumerationProduceFileReferenceURLs = 1UL << 2,
};

@interface NSFileManager (NSFileManagerMacOS)
- (nullable NSArray<NSURL *> *)mountedVolumeURLsIncludingResourceValuesForKeys:(nullable NSArray<NSURLResourceKey> *)propertyKeys options:(NSVolumeEnumerationOptions)options;
- (BOOL)trashItemAtURL:(NSURL *)url resultingItemURL:(NSURL *_Nullable *_Nullable)outResultingURL error:(NSError **)error;
@end

@interface NSDistributedNotificationCenter : NSNotificationCenter
@property(class, readonly, strong) NSDistributedNotificationCenter *defaultCenter;
@end

@interface NSBundle (NSBundleMacOS)
@property(nullable, readonly, copy) NSURL *executableURL;
@end

#define NSLocalizedString(key, comment) [[NSBundle mainBundle] localizedStringForKey:(key) value:@"" table:nil]

// Responder & application

typedef NS_OPTIONS(NSUInteger, NSEventModifierFlags) {
	NSEventModifierFlagCapsLock = 1 << 16,
	NSEventModifierFlagShift = 1 << 17,
	NSEventModifierFlagControl = 1 << 18,
	NSEventModifierFlagOption = 1 << 19,
	NSEventModifierFlagCommand = 1 << 20,
	NSEventModifierFlagNumericPad = 1 << 21,
	NSEventModifierFlagHelp = 1 << 22,
	NSEventModifierFlagFunction = 1 << 23,
	NSEventModifierFlagDeviceIndependentFlagsMask = 0xffff0000UL,
};

@interface NSResponder : NSObject <NSCoding>
@property(nullable, assign) NSResponder *nextResponder;
@property(readonly) BOOL acceptsFirstResponder;
- (BOOL)becomeFirstResponder;
- (BOOL)resignFirstResponder;
- (BOOL)tryToPerform:(SEL)action with:(nullable id)object;
- (BOOL)performKeyEquivalent:(NSEvent *)event;
- (void)mouseDown:(NSEvent *)event;
- (void)rightMouseDown:(NSEvent *)event;
- (void)otherMouseDown:(NSEvent *)event;
- (void)mouseUp:(NSEvent *)event;
- (void)rightMouseUp:(NSEvent *)event;
- (void)otherMouseUp:(NSEvent *)event;
- (void)mouseMoved:(NSEvent *)event;
- (void)mouseDragged:(NSEvent *)event;
- (void)scrollWheel:(NSEvent *)event;
- (void)rightMouseDragged:(NSEvent *)event;
- (void)otherMouseDragged:(NSEvent *)event;
- (void)mouseEntered:(NSEvent *)event;
- (void)mouseExited:(NSEvent *)event;
- (void)keyDown:(NSEvent *)event;
- (void)keyUp:(NSEvent *)event;
- (void)flagsChanged:(NSEvent *)event;
- (void)tabletPoint:(NSEvent *)event;
- (void)tabletProximity:(NSEvent *)event;
- (void)cursorUpdate:(NSEvent *)event;
- (void)magnifyWithEvent:(NSEvent *)event;
- (void)rotateWithEvent:(NSEvent *)event;
- (void)swipeWithEvent:(NSEvent *)event;
- (void)beginGestureWithEvent:(NSEvent *)event;
- (void)endGestureWithEvent:(NSEvent *)event;
- (void)smartMagnifyWithEvent:(NSEvent *)event;
- (void)pressureChangeWithEvent:(NSEvent *)event;
- (void)noResponderFor:(SEL)eventSelector;
- (void)interpretKeyEvents:(NSArray<NSEvent *> *)eventArray;
- (void)doCommandBySelector:(SEL)selector;
- (void)insertText:(id)insertString;
@property(nullable, strong) NSMenu *menu;
@end

typedef NS_ENUM(NSInteger, NSApplicationActivationPolicy) {
	NSApplicationActivationPolicyRegular,
	NSApplicationActivationPolicyAccessory,
	NSApplicationActivationPolicyProhibited,
};
typedef NS_OPTIONS(NSUInteger, NSApplicationActivationOptions) {
	NSApplicationActivateAllWindows = 1 << 0,
	NSApplicationActivateIgnoringOtherApps = 1 << 1,
};
typedef NS_OPTIONS(NSUInteger, NSApplicationPresentationOptions) {
	NSApplicationPresentationDefault = 0,
	NSApplicationPresentationAutoHideDock = (1 << 0),
	NSApplicationPresentationHideDock = (1 << 1),
	NSApplicationPresentationAutoHideMenuBar = (1 << 2),
	NSApplicationPresentationHideMenuBar = (1 << 3),
	NSApplicationPresentationDisableAppleMenu = (1 << 4),
	NSApplicationPresentationDisableProcessSwitching = (1 << 5),
	NSApplicationPresentationDisableForceQuit = (1 << 6),
	NSApplicationPresentationDisableSessionTermination = (1 << 7),
	NSApplicationPresentationDisableHideApplication = (1 << 8),
	NSApplicationPresentationDisableMenuBarTransparency = (1 << 9),
	NSApplicationPresentationFullScreen = (1 << 10),
	NSApplicationPresentationAutoHideToolbar = (1 << 11),
	NSApplicationPresentationDisableCursorLocationAssistance = (1 << 12),
};
typedef NS_ENUM(NSUInteger, NSApplicationTerminateReply) {
	NSTerminateCancel = 0,
	NSTerminateNow = 1,
	NSTerminateLater = 2,
};
typedef NS_ENUM(NSInteger, NSRequestUserAttentionType) {
	NSCriticalRequest = 0,
	NSInformationalRequest = 10,
};
typedef NS_ENUM(NSInteger, NSUserInterfaceLayoutDirection) {
	NSUserInterfaceLayoutDirectionLeftToRight = 0,
	NSUserInterfaceLayoutDirectionRightToLeft = 1,
};
typedef NSInteger NSModalResponse;
static const NSModalResponse NSModalResponseStop = (-1000);
static const NSModalResponse NSModalResponseAbort = (-1001);
static const NSModalResponse NSModalResponseContinue = (-1002);
static const NSModalResponse NSModalResponseOK = 1;
static const NSModalResponse NSModalResponseCancel = 0;

APPKIT_EXTERN const double NSAppKitVersionNumber;
#define NSAppKitVersionNumber10_15 1894
#define NSAppKitVersionNumber11_0 2022
APPKIT_EXTERN NSRunLoopMode const NSModalPanelRunLoopMode;
APPKIT_EXTERN NSRunLoopMode const NSEventTrackingRunLoopMode;
APPKIT_EXTERN NSNotificationName const NSApplicationDidChangeScreenParametersNotification;
APPKIT_EXTERN NSNotificationName const NSApplicationDidBecomeActiveNotification;
APPKIT_EXTERN NSNotificationName const NSApplicationDidResignActiveNotification;
APPKIT_EXTERN NSNotificationName const NSApplicationDidFinishLaunchingNotification;
APPKIT_EXTERN NSNotificationName const NSApplicationWillTerminateNotification;
APPKIT_EXTERN NSNotificationName const NSApplicationDidHideNotification;
APPKIT_EXTERN NSNotificationName const NSApplicationDidUnhideNotification;

@protocol NSApplicationDelegate <NSObject>
@optional
- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender;
- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)sender;
- (void)applicationWillFinishLaunching:(NSNotification *)notification;
- (void)applicationDidFinishLaunching:(NSNotification *)notification;
- (void)applicationWillTerminate:(NSNotification *)notification;
- (void)applicationDidBecomeActive:(NSNotification *)notification;
- (void)applicationDidResignActive:(NSNotification *)notification;
- (void)applicationWillHide:(NSNotification *)notification;
- (void)applicationDidHide:(NSNotification *)notification;
- (void)applicationDidUnhide:(NSNotification *)notification;
- (void)applicationDidChangeScreenParameters:(NSNotification *)notification;
- (BOOL)application:(NSApplication *)sender openFile:(NSString *)filename;
- (void)application:(NSApplication *)sender openFiles:(NSArray<NSString *> *)filenames;
- (void)application:(NSApplication *)application openURLs:(NSArray<NSURL *> *)urls;
- (BOOL)applicationSupportsSecureRestorableState:(NSApplication *)app;
- (nullable NSMenu *)applicationDockMenu:(NSApplication *)sender;
- (BOOL)applicationShouldHandleReopen:(NSApplication *)sender hasVisibleWindows:(BOOL)flag;
@end

@protocol NSUserInterfaceItemSearching <NSObject>
- (void)searchForItemsWithSearchString:(NSString *)searchString resultLimit:(NSInteger)resultLimit matchedItemHandler:(void (^)(NSArray *items))handleMatchedItems;
- (NSArray<NSString *> *)localizedTitlesForItem:(id)item;
@optional
- (void)performActionForItem:(id)item;
- (BOOL)showAllHelpTopicsForSearchString:(NSString *)searchString;
@end

@interface NSDockTile : NSObject
@property(readonly) NSSize size;
@property(nullable, strong) NSView *contentView;
@property(nullable, copy) NSString *badgeLabel;
@property BOOL showsApplicationBadge;
- (void)display;
@end

@interface NSApplication : NSResponder
@property(class, readonly, strong) __kindof NSApplication *sharedApplication;
@property(readonly, strong) NSDockTile *dockTile;
@property(nullable, weak) id<NSApplicationDelegate> delegate;
@property(nullable, readonly, weak) NSWindow *mainWindow;
@property(nullable, readonly, weak) NSWindow *keyWindow;
@property(readonly, copy) NSArray<NSWindow *> *windows;
@property(readonly, copy) NSArray<NSWindow *> *orderedWindows;
@property(readonly, getter=isActive) BOOL active;
@property(readonly, getter=isHidden) BOOL hidden;
@property(readonly, getter=isRunning) BOOL running;
@property(nullable, readonly, strong) NSEvent *currentEvent;
@property(nullable, strong) NSMenu *mainMenu;
@property(nullable, strong) NSMenu *helpMenu;
@property(nullable, strong) NSMenu *servicesMenu;
@property(nullable, strong) NSMenu *windowsMenu;
@property(nullable, strong) NSImage *applicationIconImage;
@property(readonly) NSUserInterfaceLayoutDirection userInterfaceLayoutDirection;
@property(nullable, strong) NSAppearance *appearance;
@property(readonly, strong) NSAppearance *effectiveAppearance;
@property NSApplicationPresentationOptions presentationOptions;
@property(readonly) NSApplicationPresentationOptions currentSystemPresentationOptions;
- (void)run;
- (void)finishLaunching;
- (void)stop:(nullable id)sender;
- (void)terminate:(nullable id)sender;
- (void)activateIgnoringOtherApps:(BOOL)flag;
- (void)activate;
- (void)hide:(nullable id)sender;
- (void)unhide:(nullable id)sender;
- (void)sendEvent:(NSEvent *)event;
- (void)postEvent:(NSEvent *)event atStart:(BOOL)flag;
- (nullable NSEvent *)nextEventMatchingMask:(NSUInteger)mask untilDate:(nullable NSDate *)expiration inMode:(NSRunLoopMode)mode dequeue:(BOOL)deqFlag;
- (void)updateWindows;
- (NSInteger)requestUserAttention:(NSRequestUserAttentionType)requestType;
- (void)cancelUserAttentionRequest:(NSInteger)request;
- (BOOL)setActivationPolicy:(NSApplicationActivationPolicy)activationPolicy;
- (NSApplicationActivationPolicy)activationPolicy;
- (NSModalResponse)runModalForWindow:(NSWindow *)window;
- (void)stopModal;
- (void)stopModalWithCode:(NSModalResponse)returnCode;
- (void)abortModal;
- (nullable NSWindow *)windowWithWindowNumber:(NSInteger)windowNum;
- (void)registerUserInterfaceItemSearchHandler:(id<NSUserInterfaceItemSearching>)handler;
- (void)unregisterUserInterfaceItemSearchHandler:(id<NSUserInterfaceItemSearching>)handler;
- (BOOL)sendAction:(SEL)action to:(nullable id)target from:(nullable id)sender;
- (void)setWindowsNeedUpdate:(BOOL)needUpdate;
- (void)orderFrontCharacterPalette:(nullable id)sender;
- (void)orderFrontStandardAboutPanel:(nullable id)sender;
- (void)hideOtherApplications:(nullable id)sender;
- (void)unhideAllApplications:(nullable id)sender;
@end

APPKIT_EXTERN __kindof NSApplication *_Null_unspecified NSApp;
APPKIT_EXTERN int NSApplicationMain(int argc, const char *_Nonnull argv[_Nonnull]);
APPKIT_EXTERN void NSBeep(void);
APPKIT_EXTERN void NSRectFill(NSRect rect);
APPKIT_EXTERN void NSRectFillUsingOperation(NSRect rect, NSInteger op);
APPKIT_EXTERN void NSFrameRect(NSRect rect);

@interface NSRunningApplication : NSObject
@property(class, readonly, strong) NSRunningApplication *currentApplication;
+ (NSArray<NSRunningApplication *> *)runningApplicationsWithBundleIdentifier:(NSString *)bundleIdentifier;
+ (nullable instancetype)runningApplicationWithProcessIdentifier:(pid_t)pid;
@property(readonly, getter=isTerminated) BOOL terminated;
@property(readonly, getter=isFinishedLaunching) BOOL finishedLaunching;
@property(readonly, getter=isHidden) BOOL hidden;
@property(readonly, getter=isActive) BOOL active;
@property(nullable, readonly, copy) NSString *localizedName;
@property(nullable, readonly, copy) NSString *bundleIdentifier;
@property(nullable, readonly, copy) NSURL *bundleURL;
@property(nullable, readonly, copy) NSURL *executableURL;
@property(readonly) pid_t processIdentifier;
@property(readonly) NSApplicationActivationPolicy activationPolicy;
- (BOOL)activateWithOptions:(NSApplicationActivationOptions)options;
- (BOOL)hide;
- (BOOL)unhide;
- (BOOL)terminate;
- (BOOL)forceTerminate;
@end

// Workspace

typedef NS_OPTIONS(NSUInteger, NSWorkspaceLaunchOptions) {
	NSWorkspaceLaunchAndPrint = 0x00000002,
	NSWorkspaceLaunchWithErrorPresentation = 0x00000040,
	NSWorkspaceLaunchInhibitingBackgroundOnly = 0x00000080,
	NSWorkspaceLaunchWithoutAddingToRecents = 0x00000100,
	NSWorkspaceLaunchWithoutActivation = 0x00000200,
	NSWorkspaceLaunchAsync = 0x00010000,
	NSWorkspaceLaunchNewInstance = 0x00080000,
	NSWorkspaceLaunchAndHide = 0x00100000,
	NSWorkspaceLaunchAndHideOthers = 0x00200000,
	NSWorkspaceLaunchDefault = NSWorkspaceLaunchAsync,
};
typedef NSString *NSWorkspaceLaunchConfigurationKey NS_TYPED_ENUM;
APPKIT_EXTERN NSWorkspaceLaunchConfigurationKey const NSWorkspaceLaunchConfigurationAppleEvent;
APPKIT_EXTERN NSWorkspaceLaunchConfigurationKey const NSWorkspaceLaunchConfigurationArguments;
APPKIT_EXTERN NSWorkspaceLaunchConfigurationKey const NSWorkspaceLaunchConfigurationEnvironment;
APPKIT_EXTERN NSWorkspaceLaunchConfigurationKey const NSWorkspaceLaunchConfigurationArchitecture;
APPKIT_EXTERN NSNotificationName const NSWorkspaceAccessibilityDisplayOptionsDidChangeNotification;
APPKIT_EXTERN NSNotificationName const NSWorkspaceDidActivateApplicationNotification;
APPKIT_EXTERN NSNotificationName const NSWorkspaceDidTerminateApplicationNotification;

@interface NSWorkspaceOpenConfiguration : NSObject <NSCopying>
+ (instancetype)configuration;
@property BOOL promptsUserIfNeeded;
@property BOOL addsToRecentItems;
@property BOOL activates;
@property BOOL hides;
@property BOOL hidesOthers;
@property(getter=isForPrinting) BOOL forPrinting;
@property BOOL createsNewApplicationInstance;
@property BOOL allowsRunningApplicationSubstitution;
@property(copy) NSArray<NSString *> *arguments;
@property(copy) NSDictionary<NSString *, NSString *> *environment;
@property BOOL requiresUniversalLinks;
@end

@interface NSWorkspace : NSObject
@property(class, readonly, strong) NSWorkspace *sharedWorkspace;
@property(readonly, strong) NSNotificationCenter *notificationCenter;
@property(readonly, copy) NSArray<NSRunningApplication *> *runningApplications;
@property(nullable, readonly, strong) NSRunningApplication *frontmostApplication;
@property(readonly) BOOL accessibilityDisplayShouldReduceMotion;
@property(readonly) BOOL accessibilityDisplayShouldIncreaseContrast;
@property(readonly) BOOL accessibilityDisplayShouldReduceTransparency;
@property(readonly) BOOL accessibilityDisplayShouldDifferentiateWithoutColor;
@property(readonly, getter=isVoiceOverEnabled) BOOL voiceOverEnabled;
@property(readonly, getter=isSwitchControlEnabled) BOOL switchControlEnabled;
- (BOOL)openURL:(NSURL *)url;
- (void)openURL:(NSURL *)url configuration:(NSWorkspaceOpenConfiguration *)configuration completionHandler:(void (^_Nullable)(NSRunningApplication *_Nullable app, NSError *_Nullable error))completionHandler;
- (void)openURLs:(NSArray<NSURL *> *)urls withApplicationAtURL:(NSURL *)applicationURL configuration:(NSWorkspaceOpenConfiguration *)configuration completionHandler:(void (^_Nullable)(NSRunningApplication *_Nullable app, NSError *_Nullable error))completionHandler;
- (void)openApplicationAtURL:(NSURL *)applicationURL configuration:(NSWorkspaceOpenConfiguration *)configuration completionHandler:(void (^_Nullable)(NSRunningApplication *_Nullable app, NSError *_Nullable error))completionHandler;
- (nullable NSRunningApplication *)launchApplicationAtURL:(NSURL *)url options:(NSWorkspaceLaunchOptions)options configuration:(NSDictionary<NSWorkspaceLaunchConfigurationKey, id> *)configuration error:(NSError **)error;
- (BOOL)openURLs:(NSArray<NSURL *> *)urls withApplicationAtURL:(NSURL *)applicationURL options:(NSWorkspaceLaunchOptions)options configuration:(NSDictionary<NSWorkspaceLaunchConfigurationKey, id> *)configuration error:(NSError **)error;
- (BOOL)launchApplication:(NSString *)appName;
- (BOOL)selectFile:(nullable NSString *)fullPath inFileViewerRootedAtPath:(NSString *)rootFullPath;
- (void)activateFileViewerSelectingURLs:(NSArray<NSURL *> *)fileURLs;
- (nullable NSURL *)URLForApplicationWithBundleIdentifier:(NSString *)bundleIdentifier;
- (NSImage *)iconForFile:(NSString *)fullPath;
- (BOOL)isFilePackageAtPath:(NSString *)fullPath;
@end

// Events

typedef NS_ENUM(NSUInteger, NSEventType) {
	NSEventTypeLeftMouseDown = 1,
	NSEventTypeLeftMouseUp = 2,
	NSEventTypeRightMouseDown = 3,
	NSEventTypeRightMouseUp = 4,
	NSEventTypeMouseMoved = 5,
	NSEventTypeLeftMouseDragged = 6,
	NSEventTypeRightMouseDragged = 7,
	NSEventTypeMouseEntered = 8,
	NSEventTypeMouseExited = 9,
	NSEventTypeKeyDown = 10,
	NSEventTypeKeyUp = 11,
	NSEventTypeFlagsChanged = 12,
	NSEventTypeAppKitDefined = 13,
	NSEventTypeSystemDefined = 14,
	NSEventTypeApplicationDefined = 15,
	NSEventTypePeriodic = 16,
	NSEventTypeCursorUpdate = 17,
	NSEventTypeScrollWheel = 22,
	NSEventTypeTabletPoint = 23,
	NSEventTypeTabletProximity = 24,
	NSEventTypeOtherMouseDown = 25,
	NSEventTypeOtherMouseUp = 26,
	NSEventTypeOtherMouseDragged = 27,
	NSEventTypeGesture = 29,
	NSEventTypeMagnify = 30,
	NSEventTypeSwipe = 31,
	NSEventTypeRotate = 18,
	NSEventTypeBeginGesture = 19,
	NSEventTypeEndGesture = 20,
	NSEventTypeSmartMagnify = 32,
	NSEventTypeQuickLook = 33,
	NSEventTypePressure = 34,
	NSEventTypeDirectTouch = 37,
	NSEventTypeChangeMode = 38,
};
typedef NS_OPTIONS(unsigned long long, NSEventMask) {
	NSEventMaskLeftMouseDown = 1ULL << NSEventTypeLeftMouseDown,
	NSEventMaskLeftMouseUp = 1ULL << NSEventTypeLeftMouseUp,
	NSEventMaskRightMouseDown = 1ULL << NSEventTypeRightMouseDown,
	NSEventMaskRightMouseUp = 1ULL << NSEventTypeRightMouseUp,
	NSEventMaskMouseMoved = 1ULL << NSEventTypeMouseMoved,
	NSEventMaskLeftMouseDragged = 1ULL << NSEventTypeLeftMouseDragged,
	NSEventMaskRightMouseDragged = 1ULL << NSEventTypeRightMouseDragged,
	NSEventMaskMouseEntered = 1ULL << NSEventTypeMouseEntered,
	NSEventMaskMouseExited = 1ULL << NSEventTypeMouseExited,
	NSEventMaskKeyDown = 1ULL << NSEventTypeKeyDown,
	NSEventMaskKeyUp = 1ULL << NSEventTypeKeyUp,
	NSEventMaskFlagsChanged = 1ULL << NSEventTypeFlagsChanged,
	NSEventMaskAppKitDefined = 1ULL << NSEventTypeAppKitDefined,
	NSEventMaskSystemDefined = 1ULL << NSEventTypeSystemDefined,
	NSEventMaskApplicationDefined = 1ULL << NSEventTypeApplicationDefined,
	NSEventMaskPeriodic = 1ULL << NSEventTypePeriodic,
	NSEventMaskCursorUpdate = 1ULL << NSEventTypeCursorUpdate,
	NSEventMaskScrollWheel = 1ULL << NSEventTypeScrollWheel,
	NSEventMaskTabletPoint = 1ULL << NSEventTypeTabletPoint,
	NSEventMaskTabletProximity = 1ULL << NSEventTypeTabletProximity,
	NSEventMaskOtherMouseDown = 1ULL << NSEventTypeOtherMouseDown,
	NSEventMaskOtherMouseUp = 1ULL << NSEventTypeOtherMouseUp,
	NSEventMaskOtherMouseDragged = 1ULL << NSEventTypeOtherMouseDragged,
	NSEventMaskAny = 0xffffffffffffffffULL,
};
typedef NS_ENUM(short, NSEventSubtype) {
	NSEventSubtypeWindowExposed = 0,
	NSEventSubtypeApplicationActivated = 1,
	NSEventSubtypeApplicationDeactivated = 2,
	NSEventSubtypeWindowMoved = 4,
	NSEventSubtypeScreenChanged = 8,
	NSEventSubtypePowerOff = 1,
	NSEventSubtypeMouseEvent = 0,
	NSEventSubtypeTabletPoint = 1,
	NSEventSubtypeTabletProximity = 2,
	NSEventSubtypeTouch = 3,
};
typedef NS_OPTIONS(NSUInteger, NSEventPhase) {
	NSEventPhaseNone = 0,
	NSEventPhaseBegan = 0x1 << 0,
	NSEventPhaseStationary = 0x1 << 1,
	NSEventPhaseChanged = 0x1 << 2,
	NSEventPhaseEnded = 0x1 << 3,
	NSEventPhaseCancelled = 0x1 << 4,
	NSEventPhaseMayBegin = 0x1 << 5,
};
typedef NS_ENUM(NSUInteger, NSPointingDeviceType) {
	NSPointingDeviceTypeUnknown = 0,
	NSPointingDeviceTypePen = 1,
	NSPointingDeviceTypeCursor = 2,
	NSPointingDeviceTypeEraser = 3,
};
typedef NS_OPTIONS(NSUInteger, NSEventButtonMask) {
	NSEventButtonMaskPenTip = 1,
	NSEventButtonMaskPenLowerSide = 2,
	NSEventButtonMaskPenUpperSide = 4,
};

enum {
	NSUpArrowFunctionKey = 0xF700,
	NSDownArrowFunctionKey = 0xF701,
	NSLeftArrowFunctionKey = 0xF702,
	NSRightArrowFunctionKey = 0xF703,
	NSF1FunctionKey = 0xF704,
	NSF2FunctionKey = 0xF705,
	NSF3FunctionKey = 0xF706,
	NSF4FunctionKey = 0xF707,
	NSF5FunctionKey = 0xF708,
	NSF6FunctionKey = 0xF709,
	NSF7FunctionKey = 0xF70A,
	NSF8FunctionKey = 0xF70B,
	NSF9FunctionKey = 0xF70C,
	NSF10FunctionKey = 0xF70D,
	NSF11FunctionKey = 0xF70E,
	NSF12FunctionKey = 0xF70F,
	NSF13FunctionKey = 0xF710,
	NSF14FunctionKey = 0xF711,
	NSF15FunctionKey = 0xF712,
	NSF16FunctionKey = 0xF713,
	NSF17FunctionKey = 0xF714,
	NSF18FunctionKey = 0xF715,
	NSF19FunctionKey = 0xF716,
	NSF20FunctionKey = 0xF717,
	NSF21FunctionKey = 0xF718,
	NSF22FunctionKey = 0xF719,
	NSF23FunctionKey = 0xF71A,
	NSF24FunctionKey = 0xF71B,
	NSF25FunctionKey = 0xF71C,
	NSF26FunctionKey = 0xF71D,
	NSF27FunctionKey = 0xF71E,
	NSF28FunctionKey = 0xF71F,
	NSF29FunctionKey = 0xF720,
	NSF30FunctionKey = 0xF721,
	NSF31FunctionKey = 0xF722,
	NSF32FunctionKey = 0xF723,
	NSF33FunctionKey = 0xF724,
	NSF34FunctionKey = 0xF725,
	NSF35FunctionKey = 0xF726,
	NSInsertFunctionKey = 0xF727,
	NSDeleteFunctionKey = 0xF728,
	NSHomeFunctionKey = 0xF729,
	NSBeginFunctionKey = 0xF72A,
	NSEndFunctionKey = 0xF72B,
	NSPageUpFunctionKey = 0xF72C,
	NSPageDownFunctionKey = 0xF72D,
	NSPrintScreenFunctionKey = 0xF72E,
	NSScrollLockFunctionKey = 0xF72F,
	NSPauseFunctionKey = 0xF730,
	NSSysReqFunctionKey = 0xF731,
	NSBreakFunctionKey = 0xF732,
	NSResetFunctionKey = 0xF733,
	NSStopFunctionKey = 0xF734,
	NSMenuFunctionKey = 0xF735,
	NSUserFunctionKey = 0xF736,
	NSSystemFunctionKey = 0xF737,
	NSPrintFunctionKey = 0xF738,
	NSClearLineFunctionKey = 0xF739,
	NSClearDisplayFunctionKey = 0xF73A,
	NSInsertLineFunctionKey = 0xF73B,
	NSDeleteLineFunctionKey = 0xF73C,
	NSInsertCharFunctionKey = 0xF73D,
	NSDeleteCharFunctionKey = 0xF73E,
	NSPrevFunctionKey = 0xF73F,
	NSNextFunctionKey = 0xF740,
	NSSelectFunctionKey = 0xF741,
	NSExecuteFunctionKey = 0xF742,
	NSUndoFunctionKey = 0xF743,
	NSRedoFunctionKey = 0xF744,
	NSFindFunctionKey = 0xF745,
	NSHelpFunctionKey = 0xF746,
	NSModeSwitchFunctionKey = 0xF747,
};

@interface NSEvent : NSObject <NSCopying, NSCoding>
@property(readonly) NSEventType type;
@property(readonly) NSEventModifierFlags modifierFlags;
@property(readonly) NSTimeInterval timestamp;
@property(nullable, readonly, weak) NSWindow *window;
@property(readonly) NSInteger windowNumber;
@property(readonly) NSInteger clickCount;
@property(readonly) NSInteger buttonNumber;
@property(readonly) NSInteger eventNumber;
@property(readonly) float pressure;
@property(readonly) NSPoint locationInWindow;
@property(readonly) CGFloat deltaX;
@property(readonly) CGFloat deltaY;
@property(readonly) CGFloat deltaZ;
@property(readonly) BOOL hasPreciseScrollingDeltas;
@property(readonly) CGFloat scrollingDeltaX;
@property(readonly) CGFloat scrollingDeltaY;
@property(readonly) NSEventPhase momentumPhase;
@property(readonly) NSEventPhase phase;
@property(readonly) CGFloat magnification;
@property(readonly) CGFloat rotation;
@property(nullable, readonly, copy) NSString *characters;
@property(nullable, readonly, copy) NSString *charactersIgnoringModifiers;
@property(readonly, getter=isARepeat) BOOL ARepeat;
@property(readonly) unsigned short keyCode;
@property(readonly) NSInteger trackingNumber;
@property(nullable, readonly) void *userData;
@property(nullable, readonly, strong) NSTrackingArea *trackingArea;
@property(readonly) NSEventSubtype subtype;
@property(readonly) NSInteger data1;
@property(readonly) NSInteger data2;
@property(nullable, readonly) CGEventRef CGEvent;
@property(readonly) NSPoint tilt;
@property(readonly) NSPointingDeviceType pointingDeviceType;
@property(readonly) NSUInteger pointingDeviceID;
@property(readonly) NSUInteger deviceID;
@property(readonly) NSUInteger buttonMask;
@property(readonly) float tangentialPressure;
@property(readonly, getter=isEnteringProximity) BOOL enteringProximity;
@property(readonly, getter=isDirectionInvertedFromDevice) BOOL directionInvertedFromDevice;
@property(class, readonly) NSUInteger pressedMouseButtons;
@property(class, readonly) NSPoint mouseLocation;
@property(class, readonly) NSEventModifierFlags modifierFlags;
@property(class, readonly) NSTimeInterval doubleClickInterval;
+ (nullable NSEvent *)eventWithCGEvent:(CGEventRef)cgEvent;
+ (nullable NSEvent *)mouseEventWithType:(NSEventType)type location:(NSPoint)location modifierFlags:(NSEventModifierFlags)flags timestamp:(NSTimeInterval)time windowNumber:(NSInteger)wNum context:(nullable NSGraphicsContext *)unusedPassNil eventNumber:(NSInteger)eNum clickCount:(NSInteger)cNum pressure:(float)pressure;
+ (nullable NSEvent *)otherEventWithType:(NSEventType)type location:(NSPoint)location modifierFlags:(NSEventModifierFlags)flags timestamp:(NSTimeInterval)time windowNumber:(NSInteger)wNum context:(nullable NSGraphicsContext *)unusedPassNil subtype:(short)subtype data1:(NSInteger)d1 data2:(NSInteger)d2;
+ (nullable id)addLocalMonitorForEventsMatchingMask:(NSEventMask)mask handler:(NSEvent *_Nullable (^)(NSEvent *event))block;
+ (nullable id)addGlobalMonitorForEventsMatchingMask:(NSEventMask)mask handler:(void (^)(NSEvent *event))block;
+ (void)removeMonitor:(id)eventMonitor;
@end

// Screens & appearance

APPKIT_EXTERN NSString *const NSScreenNumber; // key of deviceDescription.
APPKIT_EXTERN NSString *const NSDeviceSize;
APPKIT_EXTERN NSString *const NSDeviceResolution;
APPKIT_EXTERN NSString *const NSDeviceBitsPerSample;
APPKIT_EXTERN NSString *const NSDeviceIsScreen;
typedef NSString *NSDeviceDescriptionKey NS_TYPED_EXTENSIBLE_ENUM;

@interface NSScreen : NSObject
@property(class, readonly, copy) NSArray<NSScreen *> *screens;
@property(class, readonly, nullable, strong) NSScreen *mainScreen;
@property(class, readonly, nullable, strong) NSScreen *deepestScreen;
@property(readonly) NSRect frame;
@property(readonly) NSRect visibleFrame;
@property(readonly) NSEdgeInsets safeAreaInsets;
@property(readonly, copy) NSDictionary<NSDeviceDescriptionKey, id> *deviceDescription;
@property(readonly) CGFloat backingScaleFactor;
@property(readonly, copy) NSString *localizedName;
@property(readonly) NSInteger maximumFramesPerSecond;
@property(readonly) NSTimeInterval minimumRefreshInterval;
@property(readonly) NSTimeInterval maximumRefreshInterval;
@property(readonly) CGFloat maximumPotentialExtendedDynamicRangeColorComponentValue;
@property(readonly) CGFloat maximumExtendedDynamicRangeColorComponentValue;
@property(readonly) CGFloat maximumReferenceExtendedDynamicRangeColorComponentValue;
@property(nullable, readonly, strong) NSColorSpace *colorSpace;
- (NSRect)convertRectToBacking:(NSRect)rect;
- (NSRect)convertRectFromBacking:(NSRect)rect;
- (CADisplayLink *)displayLinkWithTarget:(id)target selector:(SEL)selector;
@end

typedef NSString *NSAppearanceName NS_TYPED_EXTENSIBLE_ENUM;
APPKIT_EXTERN NSAppearanceName const NSAppearanceNameAqua;
APPKIT_EXTERN NSAppearanceName const NSAppearanceNameDarkAqua;
APPKIT_EXTERN NSAppearanceName const NSAppearanceNameVibrantLight;
APPKIT_EXTERN NSAppearanceName const NSAppearanceNameVibrantDark;

@interface NSAppearance : NSObject <NSSecureCoding>
@property(readonly, copy) NSAppearanceName name;
@property(class, null_resettable, strong) NSAppearance *currentAppearance;
@property(class, readonly, strong) NSAppearance *currentDrawingAppearance;
+ (nullable NSAppearance *)appearanceNamed:(NSAppearanceName)name;
- (nullable NSAppearanceName)bestMatchFromAppearancesWithNames:(NSArray<NSAppearanceName> *)appearances;
- (void)performAsCurrentDrawingAppearance:(void (NS_NOESCAPE ^)(void))block;
@end

@protocol NSAppearanceCustomization <NSObject>
@property(nullable, strong) NSAppearance *appearance;
@property(readonly, strong) NSAppearance *effectiveAppearance;
@end

// Colors & images

typedef NSString *NSColorSpaceName NS_TYPED_ENUM;
APPKIT_EXTERN NSColorSpaceName const NSCalibratedRGBColorSpace;
APPKIT_EXTERN NSColorSpaceName const NSDeviceRGBColorSpace;
APPKIT_EXTERN NSColorSpaceName const NSCalibratedWhiteColorSpace;

@interface NSColorSpace : NSObject <NSSecureCoding>
@property(class, readonly, strong) NSColorSpace *sRGBColorSpace;
@property(class, readonly, strong) NSColorSpace *genericRGBColorSpace;
@property(class, readonly, strong) NSColorSpace *deviceRGBColorSpace;
@property(class, readonly, strong) NSColorSpace *displayP3ColorSpace;
@property(class, readonly, strong) NSColorSpace *extendedSRGBColorSpace;
@property(class, readonly, strong) NSColorSpace *extendedGenericGamma22GrayColorSpace;
@property(class, readonly, strong) NSColorSpace *genericGamma22GrayColorSpace;
- (nullable instancetype)initWithCGColorSpace:(CGColorSpaceRef)cgColorSpace;
@property(nullable, readonly) CGColorSpaceRef CGColorSpace;
@property(nullable, readonly, copy) NSString *localizedName;
@end

@interface NSColor : NSObject <NSCopying, NSSecureCoding>
+ (NSColor *)colorWithCalibratedRed:(CGFloat)red green:(CGFloat)green blue:(CGFloat)blue alpha:(CGFloat)alpha;
+ (NSColor *)colorWithDeviceRed:(CGFloat)red green:(CGFloat)green blue:(CGFloat)blue alpha:(CGFloat)alpha;
+ (NSColor *)colorWithSRGBRed:(CGFloat)red green:(CGFloat)green blue:(CGFloat)blue alpha:(CGFloat)alpha;
+ (NSColor *)colorWithRed:(CGFloat)red green:(CGFloat)green blue:(CGFloat)blue alpha:(CGFloat)alpha;
+ (NSColor *)colorWithColorSpace:(NSColorSpace *)space components:(const CGFloat *)components count:(NSInteger)numberOfComponents;
+ (nullable NSColor *)colorWithCGColor:(CGColorRef)cgColor;
@property(class, strong, readonly) NSColor *blackColor;
@property(class, strong, readonly) NSColor *whiteColor;
@property(class, strong, readonly) NSColor *grayColor;
@property(class, strong, readonly) NSColor *clearColor;
@property(class, strong, readonly) NSColor *redColor;
@property(class, strong, readonly) NSColor *controlColor;
@property(class, strong, readonly) NSColor *controlAccentColor;
@property(class, strong, readonly) NSColor *controlTextColor;
@property(class, strong, readonly) NSColor *windowBackgroundColor;
@property(class, strong, readonly) NSColor *textColor;
@property(class, strong, readonly) NSColor *labelColor;
@property(class, strong, readonly) NSColor *secondaryLabelColor;
@property(class, strong, readonly) NSColor *tertiaryLabelColor;
@property(class, strong, readonly) NSColor *textBackgroundColor;
@property(class, strong, readonly) NSColor *selectedContentBackgroundColor;
@property(readonly) CGColorRef CGColor;
- (nullable NSColor *)colorUsingColorSpace:(NSColorSpace *)space;
- (nullable NSColor *)colorUsingColorSpaceName:(NSColorSpaceName)name;
- (void)getRed:(nullable CGFloat *)red green:(nullable CGFloat *)green blue:(nullable CGFloat *)blue alpha:(nullable CGFloat *)alpha;
- (void)set;
- (void)setFill;
- (void)setStroke;
@property(readonly) CGFloat redComponent;
@property(readonly) CGFloat greenComponent;
@property(readonly) CGFloat blueComponent;
@property(readonly) CGFloat alphaComponent;
@end

@interface NSImageRep : NSObject <NSCopying, NSCoding>
@property NSSize size;
@property NSInteger pixelsWide;
@property NSInteger pixelsHigh;
@property NSInteger bitsPerSample;
@property(nullable, copy) NSColorSpaceName colorSpaceName;
@property(getter=hasAlpha) BOOL alpha;
- (BOOL)drawInRect:(NSRect)rect;
@end

typedef NS_OPTIONS(NSUInteger, NSBitmapFormat) {
	NSBitmapFormatAlphaFirst = 1 << 0,
	NSBitmapFormatAlphaNonpremultiplied = 1 << 1,
	NSBitmapFormatFloatingPointSamples = 1 << 2,
	NSBitmapFormatSixteenBitLittleEndian = (1 << 8),
	NSBitmapFormatThirtyTwoBitLittleEndian = (1 << 9),
	NSBitmapFormatSixteenBitBigEndian = (1 << 10),
	NSBitmapFormatThirtyTwoBitBigEndian = (1 << 11),
};
typedef NS_ENUM(NSUInteger, NSBitmapImageFileType) {
	NSBitmapImageFileTypeTIFF,
	NSBitmapImageFileTypeBMP,
	NSBitmapImageFileTypeGIF,
	NSBitmapImageFileTypeJPEG,
	NSBitmapImageFileTypePNG,
	NSBitmapImageFileTypeJPEG2000,
};
typedef NSString *NSBitmapImageRepPropertyKey NS_TYPED_ENUM;
APPKIT_EXTERN NSBitmapImageRepPropertyKey const NSImageCompressionFactor;
APPKIT_EXTERN NSBitmapImageRepPropertyKey const NSImageInterlaced;

@interface NSBitmapImageRep : NSImageRep
- (nullable instancetype)initWithBitmapDataPlanes:(unsigned char *_Nullable *_Nullable)planes pixelsWide:(NSInteger)width pixelsHigh:(NSInteger)height bitsPerSample:(NSInteger)bps samplesPerPixel:(NSInteger)spp hasAlpha:(BOOL)alpha isPlanar:(BOOL)isPlanar colorSpaceName:(NSColorSpaceName)colorSpaceName bytesPerRow:(NSInteger)rBytes bitsPerPixel:(NSInteger)pBits;
- (nullable instancetype)initWithBitmapDataPlanes:(unsigned char *_Nullable *_Nullable)planes pixelsWide:(NSInteger)width pixelsHigh:(NSInteger)height bitsPerSample:(NSInteger)bps samplesPerPixel:(NSInteger)spp hasAlpha:(BOOL)alpha isPlanar:(BOOL)isPlanar colorSpaceName:(NSColorSpaceName)colorSpaceName bitmapFormat:(NSBitmapFormat)bitmapFormat bytesPerRow:(NSInteger)rBytes bitsPerPixel:(NSInteger)pBits;
- (nullable instancetype)initWithCGImage:(CGImageRef)cgImage;
- (nullable instancetype)initWithData:(NSData *)data;
+ (nullable instancetype)imageRepWithData:(NSData *)data;
@property(nullable, readonly) unsigned char *bitmapData;
@property(readonly) NSInteger bytesPerPlane;
@property(readonly) NSInteger bytesPerRow;
@property(readonly) NSInteger bitsPerPixel;
@property(readonly) NSInteger samplesPerPixel;
@property(readonly, getter=isPlanar) BOOL planar;
@property(readonly) NSBitmapFormat bitmapFormat;
@property(nullable, readonly) CGImageRef CGImage;
- (nullable NSData *)representationUsingType:(NSBitmapImageFileType)storageType properties:(NSDictionary<NSBitmapImageRepPropertyKey, id> *)properties;
- (nullable NSColor *)colorAtX:(NSInteger)x y:(NSInteger)y;
- (void)getPixel:(NSUInteger[_Nonnull])p atX:(NSInteger)x y:(NSInteger)y;
@end

typedef NSString *NSImageName NS_TYPED_EXTENSIBLE_ENUM;

@interface NSImage : NSObject <NSCopying, NSSecureCoding>
+ (nullable NSImage *)imageNamed:(NSImageName)name;
+ (nullable NSImage *)imageWithSystemSymbolName:(NSString *)symbolName accessibilityDescription:(nullable NSString *)description;
- (instancetype)initWithSize:(NSSize)size;
- (nullable instancetype)initWithData:(NSData *)data;
- (nullable instancetype)initWithContentsOfFile:(NSString *)fileName;
- (nullable instancetype)initWithContentsOfURL:(NSURL *)url;
- (instancetype)initWithCGImage:(CGImageRef)cgImage size:(NSSize)size;
@property NSSize size;
@property(readonly, copy) NSArray<NSImageRep *> *representations;
@property(nullable, copy) NSString *name;
- (void)addRepresentation:(NSImageRep *)imageRep;
- (void)removeRepresentation:(NSImageRep *)imageRep;
- (void)lockFocus;
- (void)unlockFocus;
- (nullable CGImageRef)CGImageForProposedRect:(nullable NSRect *)proposedDestRect context:(nullable NSGraphicsContext *)referenceContext hints:(nullable NSDictionary *)hints;
- (nullable NSData *)TIFFRepresentation;
- (void)drawInRect:(NSRect)rect;
- (void)drawInRect:(NSRect)dstRect fromRect:(NSRect)srcRect operation:(NSInteger)op fraction:(CGFloat)delta;
@end

typedef NS_ENUM(NSUInteger, NSImageInterpolation) {
	NSImageInterpolationDefault = 0,
	NSImageInterpolationNone = 1,
	NSImageInterpolationLow = 2,
	NSImageInterpolationMedium = 4,
	NSImageInterpolationHigh = 3,
};

@interface NSGraphicsContext : NSObject
@property(class, nullable, strong) NSGraphicsContext *currentContext;
+ (NSGraphicsContext *)graphicsContextWithBitmapImageRep:(NSBitmapImageRep *)bitmapRep;
+ (NSGraphicsContext *)graphicsContextWithCGContext:(CGContextRef)graphicsPort flipped:(BOOL)initialFlippedState;
+ (void)saveGraphicsState;
+ (void)restoreGraphicsState;
- (void)saveGraphicsState;
- (void)restoreGraphicsState;
- (void)flushGraphics;
@property NSImageInterpolation imageInterpolation;
@property(readonly) CGContextRef CGContext;
@property(readonly, getter=isFlipped) BOOL flipped;
@end

@interface NSBezierPath : NSObject <NSCopying, NSSecureCoding>
+ (NSBezierPath *)bezierPath;
+ (NSBezierPath *)bezierPathWithRect:(NSRect)rect;
+ (NSBezierPath *)bezierPathWithOvalInRect:(NSRect)rect;
+ (NSBezierPath *)bezierPathWithRoundedRect:(NSRect)rect xRadius:(CGFloat)xRadius yRadius:(CGFloat)yRadius;
@property CGFloat lineWidth;
- (void)moveToPoint:(NSPoint)point;
- (void)lineToPoint:(NSPoint)point;
- (void)closePath;
- (void)stroke;
- (void)fill;
- (void)addClip;
- (void)setClip;
@end

// Cursor

@interface NSCursor : NSObject <NSCoding>
@property(class, readonly, strong) NSCursor *currentCursor;
@property(class, readonly, strong) NSCursor *arrowCursor;
@property(class, readonly, strong) NSCursor *IBeamCursor;
@property(class, readonly, strong) NSCursor *pointingHandCursor;
@property(class, readonly, strong) NSCursor *closedHandCursor;
@property(class, readonly, strong) NSCursor *openHandCursor;
@property(class, readonly, strong) NSCursor *resizeLeftCursor;
@property(class, readonly, strong) NSCursor *resizeRightCursor;
@property(class, readonly, strong) NSCursor *resizeLeftRightCursor;
@property(class, readonly, strong) NSCursor *resizeUpCursor;
@property(class, readonly, strong) NSCursor *resizeDownCursor;
@property(class, readonly, strong) NSCursor *resizeUpDownCursor;
@property(class, readonly, strong) NSCursor *crosshairCursor;
@property(class, readonly, strong) NSCursor *disappearingItemCursor;
@property(class, readonly, strong) NSCursor *operationNotAllowedCursor;
@property(class, readonly, strong) NSCursor *dragLinkCursor;
@property(class, readonly, strong) NSCursor *dragCopyCursor;
@property(class, readonly, strong) NSCursor *contextualMenuCursor;
@property(class, readonly, strong) NSCursor *IBeamCursorForVerticalLayout;
- (instancetype)initWithImage:(NSImage *)newImage hotSpot:(NSPoint)point;
@property(readonly, strong) NSImage *image;
@property(readonly) NSPoint hotSpot;
+ (void)hide;
+ (void)unhide;
+ (void)setHiddenUntilMouseMoves:(BOOL)flag;
+ (void)pop;
- (void)push;
- (void)pop;
- (void)set;
@end

// Views

typedef NS_OPTIONS(NSUInteger, NSAutoresizingMaskOptions) {
	NSViewNotSizable = 0,
	NSViewMinXMargin = 1,
	NSViewWidthSizable = 2,
	NSViewMaxXMargin = 4,
	NSViewMinYMargin = 8,
	NSViewHeightSizable = 16,
	NSViewMaxYMargin = 32,
};
typedef NS_ENUM(NSInteger, NSViewLayerContentsRedrawPolicy) {
	NSViewLayerContentsRedrawNever = 0,
	NSViewLayerContentsRedrawOnSetNeedsDisplay = 1,
	NSViewLayerContentsRedrawDuringViewResize = 2,
	NSViewLayerContentsRedrawBeforeViewResize = 3,
	NSViewLayerContentsRedrawCrossfade = 4,
};
typedef NS_ENUM(NSInteger, NSViewLayerContentsPlacement) {
	NSViewLayerContentsPlacementScaleAxesIndependently = 0,
	NSViewLayerContentsPlacementScaleProportionallyToFit = 1,
	NSViewLayerContentsPlacementScaleProportionallyToFill = 2,
	NSViewLayerContentsPlacementCenter = 3,
	NSViewLayerContentsPlacementTop = 4,
	NSViewLayerContentsPlacementTopRight = 5,
	NSViewLayerContentsPlacementRight = 6,
	NSViewLayerContentsPlacementBottomRight = 7,
	NSViewLayerContentsPlacementBottom = 8,
	NSViewLayerContentsPlacementBottomLeft = 9,
	NSViewLayerContentsPlacementLeft = 10,
	NSViewLayerContentsPlacementTopLeft = 11,
};
typedef NS_OPTIONS(NSUInteger, NSTrackingAreaOptions) {
	NSTrackingMouseEnteredAndExited = 0x01,
	NSTrackingMouseMoved = 0x02,
	NSTrackingCursorUpdate = 0x04,
	NSTrackingActiveWhenFirstResponder = 0x10,
	NSTrackingActiveInKeyWindow = 0x20,
	NSTrackingActiveInActiveApp = 0x40,
	NSTrackingActiveAlways = 0x80,
	NSTrackingAssumeInside = 0x100,
	NSTrackingInVisibleRect = 0x200,
	NSTrackingEnabledDuringMouseDrag = 0x400,
};
typedef NS_OPTIONS(NSUInteger, NSDragOperation) {
	NSDragOperationNone = 0,
	NSDragOperationCopy = 1,
	NSDragOperationLink = 2,
	NSDragOperationGeneric = 4,
	NSDragOperationPrivate = 8,
	NSDragOperationMove = 16,
	NSDragOperationDelete = 32,
	NSDragOperationEvery = NSUIntegerMax,
};
typedef NSString *NSPasteboardType NS_TYPED_EXTENSIBLE_ENUM;
typedef NSString *NSPasteboardName NS_TYPED_EXTENSIBLE_ENUM;
typedef NSString *NSPasteboardReadingOptionKey NS_TYPED_ENUM;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeString;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypePDF;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeTIFF;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypePNG;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeRTF;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeHTML;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeURL;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeFileURL;
APPKIT_EXTERN NSPasteboardType const NSPasteboardTypeColor;
APPKIT_EXTERN NSPasteboardName const NSPasteboardNameGeneral;
APPKIT_EXTERN NSPasteboardName const NSPasteboardNameDrag;
APPKIT_EXTERN NSPasteboardReadingOptionKey const NSPasteboardURLReadingFileURLsOnlyKey;
APPKIT_EXTERN NSPasteboardReadingOptionKey const NSPasteboardURLReadingContentsConformToTypesKey;

@interface NSTrackingArea : NSObject <NSCopying, NSCoding>
- (instancetype)initWithRect:(NSRect)rect options:(NSTrackingAreaOptions)options owner:(nullable id)owner userInfo:(nullable NSDictionary<id, id> *)userInfo;
@property(readonly) NSRect rect;
@property(readonly) NSTrackingAreaOptions options;
@property(nullable, readonly, weak) id owner;
@property(nullable, readonly, copy) NSDictionary<id, id> *userInfo;
@end

@protocol NSPasteboardWriting <NSObject>
@end
@protocol NSPasteboardReading <NSObject>
@end

@interface NSPasteboardItem : NSObject <NSPasteboardWriting, NSPasteboardReading>
@property(readonly, copy) NSArray<NSPasteboardType> *types;
- (nullable NSPasteboardType)availableTypeFromArray:(NSArray<NSPasteboardType> *)types;
- (BOOL)setData:(NSData *)data forType:(NSPasteboardType)type;
- (BOOL)setString:(NSString *)string forType:(NSPasteboardType)type;
- (BOOL)setPropertyList:(id)propertyList forType:(NSPasteboardType)type;
- (nullable NSData *)dataForType:(NSPasteboardType)type;
- (nullable NSString *)stringForType:(NSPasteboardType)type;
- (nullable id)propertyListForType:(NSPasteboardType)type;
@end

@interface NSPasteboard : NSObject
@property(class, readonly, strong) NSPasteboard *generalPasteboard;
+ (NSPasteboard *)pasteboardWithName:(NSPasteboardName)name;
+ (NSPasteboard *)pasteboardWithUniqueName;
@property(readonly) NSPasteboardName name;
@property(readonly) NSInteger changeCount;
@property(nullable, readonly, copy) NSArray<NSPasteboardType> *types;
@property(nullable, readonly, copy) NSArray<NSPasteboardItem *> *pasteboardItems;
- (NSInteger)clearContents;
- (BOOL)writeObjects:(NSArray<id<NSPasteboardWriting>> *)objects;
- (nullable NSArray *)readObjectsForClasses:(NSArray<Class> *)classArray options:(nullable NSDictionary<NSPasteboardReadingOptionKey, id> *)options;
- (BOOL)canReadObjectForClasses:(NSArray<Class> *)classArray options:(nullable NSDictionary<NSPasteboardReadingOptionKey, id> *)options;
- (nullable NSPasteboardType)availableTypeFromArray:(NSArray<NSPasteboardType> *)types;
- (nullable NSData *)dataForType:(NSPasteboardType)dataType;
- (nullable NSString *)stringForType:(NSPasteboardType)dataType;
- (nullable id)propertyListForType:(NSPasteboardType)dataType;
- (BOOL)setData:(nullable NSData *)data forType:(NSPasteboardType)dataType;
- (BOOL)setString:(NSString *)string forType:(NSPasteboardType)dataType;
- (BOOL)setPropertyList:(id)plist forType:(NSPasteboardType)dataType;
- (NSInteger)declareTypes:(NSArray<NSPasteboardType> *)newTypes owner:(nullable id)newOwner;
@end

@interface NSURL (NSPasteboardSupport) <NSPasteboardWriting, NSPasteboardReading>
@end
@interface NSString (NSPasteboardSupport) <NSPasteboardWriting, NSPasteboardReading>
@end
@interface NSImage (NSPasteboardSupport) <NSPasteboardWriting, NSPasteboardReading>
@end

@protocol NSDraggingInfo <NSObject>
@property(nullable, readonly) NSWindow *draggingDestinationWindow;
@property(readonly) NSDragOperation draggingSourceOperationMask;
@property(readonly) NSPoint draggingLocation;
@property(readonly) NSPasteboard *draggingPasteboard;
@property(nullable, readonly) id draggingSource;
@property(readonly) NSInteger draggingSequenceNumber;
@end

@protocol NSDraggingDestination <NSObject>
@optional
- (NSDragOperation)draggingEntered:(id<NSDraggingInfo>)sender;
- (NSDragOperation)draggingUpdated:(id<NSDraggingInfo>)sender;
- (void)draggingExited:(nullable id<NSDraggingInfo>)sender;
- (BOOL)prepareForDragOperation:(id<NSDraggingInfo>)sender;
- (BOOL)performDragOperation:(id<NSDraggingInfo>)sender;
- (void)concludeDragOperation:(nullable id<NSDraggingInfo>)sender;
- (BOOL)wantsPeriodicDraggingUpdates;
@end

@interface NSLayoutAnchor<AnchorType> : NSObject <NSCopying, NSCoding>
- (NSLayoutConstraint *)constraintEqualToAnchor:(NSLayoutAnchor<AnchorType> *)anchor;
- (NSLayoutConstraint *)constraintEqualToAnchor:(NSLayoutAnchor<AnchorType> *)anchor constant:(CGFloat)c;
- (NSLayoutConstraint *)constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor<AnchorType> *)anchor;
- (NSLayoutConstraint *)constraintLessThanOrEqualToAnchor:(NSLayoutAnchor<AnchorType> *)anchor;
@end
@interface NSLayoutXAxisAnchor : NSLayoutAnchor <NSLayoutXAxisAnchor *>
@end
@interface NSLayoutYAxisAnchor : NSLayoutAnchor <NSLayoutYAxisAnchor *>
@end
@interface NSLayoutDimension : NSLayoutAnchor <NSLayoutDimension *>
- (NSLayoutConstraint *)constraintEqualToConstant:(CGFloat)c;
@end
@interface NSLayoutConstraint : NSObject
+ (void)activateConstraints:(NSArray<NSLayoutConstraint *> *)constraints;
+ (void)deactivateConstraints:(NSArray<NSLayoutConstraint *> *)constraints;
@property(getter=isActive) BOOL active;
@property CGFloat constant;
@end

@interface NSView : NSResponder <NSDraggingDestination, NSAppearanceCustomization>
- (instancetype)initWithFrame:(NSRect)frameRect;
@property(nullable, readonly, assign) NSWindow *window;
@property(nullable, readonly, assign) NSView *superview;
@property(copy) NSArray<__kindof NSView *> *subviews;
- (void)addSubview:(NSView *)view;
- (void)removeFromSuperview;
- (void)viewDidMoveToWindow;
- (void)viewDidMoveToSuperview;
- (void)viewDidChangeBackingProperties;
- (void)viewDidChangeEffectiveAppearance;
@property NSRect frame;
@property NSRect bounds;
@property(readonly, getter=isFlipped) BOOL flipped;
@property(getter=isHidden) BOOL hidden;
@property(readonly, getter=isHiddenOrHasHiddenAncestor) BOOL hiddenOrHasHiddenAncestor;
@property(readonly) BOOL canBecomeKeyView;
@property NSAutoresizingMaskOptions autoresizingMask;
@property BOOL translatesAutoresizingMaskIntoConstraints;
@property BOOL needsDisplay;
@property BOOL needsLayout;
@property(nullable, strong) CALayer *layer;
@property BOOL wantsLayer;
@property(readonly) BOOL wantsUpdateLayer;
@property NSViewLayerContentsRedrawPolicy layerContentsRedrawPolicy;
@property NSViewLayerContentsPlacement layerContentsPlacement;
@property(readonly) NSEdgeInsets safeAreaInsets;
@property(readonly, copy) NSArray<NSTrackingArea *> *trackingAreas;
@property(nullable, copy) NSString *toolTip;
@property(nullable, readonly, strong) NSTextInputContext *inputContext;
@property(readonly, strong) NSLayoutXAxisAnchor *leadingAnchor;
@property(readonly, strong) NSLayoutXAxisAnchor *trailingAnchor;
@property(readonly, strong) NSLayoutXAxisAnchor *leftAnchor;
@property(readonly, strong) NSLayoutXAxisAnchor *rightAnchor;
@property(readonly, strong) NSLayoutYAxisAnchor *topAnchor;
@property(readonly, strong) NSLayoutYAxisAnchor *bottomAnchor;
@property(readonly, strong) NSLayoutDimension *widthAnchor;
@property(readonly, strong) NSLayoutDimension *heightAnchor;
@property(readonly, strong) NSLayoutXAxisAnchor *centerXAnchor;
@property(readonly, strong) NSLayoutYAxisAnchor *centerYAnchor;
- (void)setFrameOrigin:(NSPoint)newOrigin;
- (void)setFrameSize:(NSSize)newSize;
- (void)setNeedsDisplayInRect:(NSRect)invalidRect;
- (void)display;
- (void)displayIfNeeded;
- (void)drawRect:(NSRect)dirtyRect;
- (void)updateLayer;
- (void)layout;
- (void)addTrackingArea:(NSTrackingArea *)trackingArea;
- (void)removeTrackingArea:(NSTrackingArea *)trackingArea;
- (void)updateTrackingAreas;
- (NSPoint)convertPoint:(NSPoint)point fromView:(nullable NSView *)view;
- (NSPoint)convertPoint:(NSPoint)point toView:(nullable NSView *)view;
- (NSRect)convertRect:(NSRect)rect fromView:(nullable NSView *)view;
- (NSRect)convertRect:(NSRect)rect toView:(nullable NSView *)view;
- (NSRect)convertRectToBacking:(NSRect)rect;
- (NSRect)convertRectFromBacking:(NSRect)rect;
- (NSPoint)convertPointToBacking:(NSPoint)point;
- (NSPoint)convertPointFromBacking:(NSPoint)point;
- (NSSize)convertSizeToBacking:(NSSize)size;
- (NSSize)convertSizeFromBacking:(NSSize)size;
- (void)registerForDraggedTypes:(NSArray<NSPasteboardType> *)newTypes;
- (void)unregisterDraggedTypes;
- (nullable NSView *)hitTest:(NSPoint)point;
- (BOOL)mouse:(NSPoint)point inRect:(NSRect)rect;
- (void)scrollWheel:(NSEvent *)event;
- (NSInteger)tag;
@end

// Windows

typedef NS_OPTIONS(NSUInteger, NSWindowStyleMask) {
	NSWindowStyleMaskBorderless = 0,
	NSWindowStyleMaskTitled = 1 << 0,
	NSWindowStyleMaskClosable = 1 << 1,
	NSWindowStyleMaskMiniaturizable = 1 << 2,
	NSWindowStyleMaskResizable = 1 << 3,
	NSWindowStyleMaskUtilityWindow = 1 << 4,
	NSWindowStyleMaskDocModalWindow = 1 << 6,
	NSWindowStyleMaskNonactivatingPanel = 1 << 7,
	NSWindowStyleMaskUnifiedTitleAndToolbar = 1 << 12,
	NSWindowStyleMaskFullScreen = 1 << 14,
	NSWindowStyleMaskFullSizeContentView = 1 << 15,
	NSWindowStyleMaskHUDWindow = 1 << 13,
};
typedef NS_ENUM(NSUInteger, NSBackingStoreType) {
	NSBackingStoreRetained = 0,
	NSBackingStoreNonretained = 1,
	NSBackingStoreBuffered = 2,
};
typedef NS_ENUM(NSInteger, NSWindowOrderingMode) {
	NSWindowAbove = 1,
	NSWindowBelow = -1,
	NSWindowOut = 0,
};
typedef NS_ENUM(NSInteger, NSWindowSharingType) {
	NSWindowSharingNone = 0,
	NSWindowSharingReadOnly = 1,
	NSWindowSharingReadWrite = 2,
};
typedef NS_OPTIONS(NSUInteger, NSWindowCollectionBehavior) {
	NSWindowCollectionBehaviorDefault = 0,
	NSWindowCollectionBehaviorCanJoinAllSpaces = 1 << 0,
	NSWindowCollectionBehaviorMoveToActiveSpace = 1 << 1,
	NSWindowCollectionBehaviorManaged = 1 << 2,
	NSWindowCollectionBehaviorTransient = 1 << 3,
	NSWindowCollectionBehaviorStationary = 1 << 4,
	NSWindowCollectionBehaviorParticipatesInCycle = 1 << 5,
	NSWindowCollectionBehaviorIgnoresCycle = 1 << 6,
	NSWindowCollectionBehaviorFullScreenPrimary = 1 << 7,
	NSWindowCollectionBehaviorFullScreenAuxiliary = 1 << 8,
	NSWindowCollectionBehaviorFullScreenNone = 1 << 9,
	NSWindowCollectionBehaviorFullScreenAllowsTiling = 1 << 11,
	NSWindowCollectionBehaviorFullScreenDisallowsTiling = 1 << 12,
};
typedef NS_ENUM(NSInteger, NSWindowTabbingMode) {
	NSWindowTabbingModeAutomatic,
	NSWindowTabbingModePreferred,
	NSWindowTabbingModeDisallowed,
};
typedef NS_ENUM(NSInteger, NSWindowTitleVisibility) {
	NSWindowTitleVisible = 0,
	NSWindowTitleHidden = 1,
};
typedef NS_ENUM(NSUInteger, NSWindowButton) {
	NSWindowCloseButton,
	NSWindowMiniaturizeButton,
	NSWindowZoomButton,
	NSWindowToolbarButton,
	NSWindowDocumentIconButton,
	NSWindowDocumentVersionsButton = 6,
};
typedef NS_OPTIONS(NSUInteger, NSWindowOcclusionState) {
	NSWindowOcclusionStateVisible = 1UL << 1,
};
typedef NSInteger NSWindowLevel;
static const NSWindowLevel NSNormalWindowLevel = 0;
static const NSWindowLevel NSFloatingWindowLevel = 3;
static const NSWindowLevel NSSubmenuWindowLevel = 3;
static const NSWindowLevel NSTornOffMenuWindowLevel = 3;
static const NSWindowLevel NSMainMenuWindowLevel = 24;
static const NSWindowLevel NSStatusWindowLevel = 25;
static const NSWindowLevel NSModalPanelWindowLevel = 8;
static const NSWindowLevel NSPopUpMenuWindowLevel = 101;
static const NSWindowLevel NSScreenSaverWindowLevel = 1000;
APPKIT_EXTERN NSString *const NSBackingPropertyOldScaleFactorKey;
APPKIT_EXTERN NSString *const NSBackingPropertyOldColorSpaceKey;
APPKIT_EXTERN NSNotificationName const NSWindowDidBecomeKeyNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidResignKeyNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidResizeNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidMoveNotification;
APPKIT_EXTERN NSNotificationName const NSWindowWillCloseNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidChangeScreenNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidChangeBackingPropertiesNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidEnterFullScreenNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidExitFullScreenNotification;
APPKIT_EXTERN NSNotificationName const NSWindowDidChangeOcclusionStateNotification;

@protocol NSWindowDelegate <NSObject>
@optional
- (BOOL)windowShouldClose:(NSWindow *)sender;
- (void)windowWillClose:(NSNotification *)notification;
- (void)windowDidResize:(NSNotification *)notification;
- (void)windowDidMove:(NSNotification *)notification;
- (void)windowDidBecomeKey:(NSNotification *)notification;
- (void)windowDidResignKey:(NSNotification *)notification;
- (void)windowDidBecomeMain:(NSNotification *)notification;
- (void)windowDidResignMain:(NSNotification *)notification;
- (void)windowDidMiniaturize:(NSNotification *)notification;
- (void)windowDidDeminiaturize:(NSNotification *)notification;
- (void)windowWillEnterFullScreen:(NSNotification *)notification;
- (void)windowDidEnterFullScreen:(NSNotification *)notification;
- (void)windowWillExitFullScreen:(NSNotification *)notification;
- (void)windowDidExitFullScreen:(NSNotification *)notification;
- (void)windowDidFailToEnterFullScreen:(NSWindow *)window;
- (void)windowDidFailToExitFullScreen:(NSWindow *)window;
- (void)windowDidChangeScreen:(NSNotification *)notification;
- (void)windowDidChangeBackingProperties:(NSNotification *)notification;
- (void)windowDidChangeOcclusionState:(NSNotification *)notification;
- (void)windowDidUpdate:(NSNotification *)notification;
- (void)windowDidExpose:(NSNotification *)notification;
- (NSSize)windowWillResize:(NSWindow *)sender toSize:(NSSize)frameSize;
- (BOOL)windowShouldZoom:(NSWindow *)window toFrame:(NSRect)newFrame;
- (NSRect)windowWillUseStandardFrame:(NSWindow *)window defaultFrame:(NSRect)newFrame;
- (NSApplicationPresentationOptions)window:(NSWindow *)window willUseFullScreenPresentationOptions:(NSApplicationPresentationOptions)proposedOptions;
- (NSSize)window:(NSWindow *)window willUseFullScreenContentSize:(NSSize)proposedSize;
@end

@interface NSWindow : NSResponder <NSAppearanceCustomization>
- (instancetype)initWithContentRect:(NSRect)contentRect styleMask:(NSWindowStyleMask)style backing:(NSBackingStoreType)backingStoreType defer:(BOOL)flag;
+ (NSRect)frameRectForContentRect:(NSRect)cRect styleMask:(NSWindowStyleMask)style;
+ (NSRect)contentRectForFrameRect:(NSRect)fRect styleMask:(NSWindowStyleMask)style;
- (NSRect)frameRectForContentRect:(NSRect)contentRect;
- (NSRect)contentRectForFrameRect:(NSRect)frameRect;
+ (NSInteger)windowNumberAtPoint:(NSPoint)point belowWindowWithWindowNumber:(NSInteger)windowNumber;
@property(copy) NSString *title;
@property NSWindowTitleVisibility titleVisibility;
@property BOOL titlebarAppearsTransparent;
@property(nullable, weak) id<NSWindowDelegate> delegate;
@property(readonly) NSInteger windowNumber;
@property NSWindowStyleMask styleMask;
@property(readonly) NSRect frame;
@property(nullable, readonly, strong) NSScreen *screen;
@property(nullable, readonly, strong) NSScreen *deepestScreen;
@property(nullable, strong) __kindof NSView *contentView;
@property(readonly, getter=isVisible) BOOL visible;
@property(readonly, getter=isKeyWindow) BOOL keyWindow;
@property(readonly, getter=isMainWindow) BOOL mainWindow;
@property(readonly, getter=isMiniaturized) BOOL miniaturized;
@property(readonly, getter=isZoomed) BOOL zoomed;
@property(readonly, getter=isOnActiveSpace) BOOL onActiveSpace;
@property(readonly) BOOL canBecomeKeyWindow;
@property(readonly) BOOL canBecomeMainWindow;
@property(readonly) BOOL worksWhenModal;
@property NSWindowLevel level;
@property(nullable, readonly, weak) NSWindow *parentWindow;
@property(nullable, readonly, copy) NSArray<__kindof NSWindow *> *childWindows;
@property(nullable, readonly, weak) NSResponder *firstResponder;
@property(readonly) CGFloat backingScaleFactor;
@property(readonly) NSWindowOcclusionState occlusionState;
@property NSWindowSharingType sharingType;
@property NSWindowCollectionBehavior collectionBehavior;
@property NSWindowTabbingMode tabbingMode;
@property(nullable, strong) NSColorSpace *colorSpace;
@property BOOL hidesOnDeactivate;
@property(nullable, copy) NSColor *backgroundColor;
@property(getter=isOpaque) BOOL opaque;
@property CGFloat alphaValue;
@property BOOL hasShadow;
@property BOOL ignoresMouseEvents;
@property BOOL acceptsMouseMovedEvents;
@property(getter=isMovable) BOOL movable;
@property(getter=isMovableByWindowBackground) BOOL movableByWindowBackground;
@property(getter=isReleasedWhenClosed) BOOL releasedWhenClosed;
@property(getter=isRestorable) BOOL restorable;
@property(getter=isExcludedFromWindowsMenu) BOOL excludedFromWindowsMenu;
@property NSSize minSize;
@property NSSize maxSize;
@property NSSize contentMinSize;
@property NSSize contentMaxSize;
@property NSSize contentAspectRatio;
@property NSSize aspectRatio;
@property(readonly) NSUserInterfaceLayoutDirection windowTitlebarLayoutDirection;
@property(nullable, readonly, strong) NSGraphicsContext *graphicsContext;
@property(nullable, readonly, strong) NSWindowTab *tab;
- (void)setFrame:(NSRect)frameRect display:(BOOL)flag;
- (void)setFrame:(NSRect)frameRect display:(BOOL)displayFlag animate:(BOOL)animateFlag;
- (void)setFrameOrigin:(NSPoint)point;
- (void)setFrameTopLeftPoint:(NSPoint)point;
- (void)setContentSize:(NSSize)size;
- (NSTimeInterval)animationResizeTime:(NSRect)newFrame;
- (void)makeKeyAndOrderFront:(nullable id)sender;
- (void)makeKeyWindow;
- (void)makeMainWindow;
- (void)orderFront:(nullable id)sender;
- (void)orderBack:(nullable id)sender;
- (void)orderOut:(nullable id)sender;
- (void)orderWindow:(NSWindowOrderingMode)place relativeTo:(NSInteger)otherWin;
- (void)orderFrontRegardless;
- (void)close;
- (void)performClose:(nullable id)sender;
- (void)performMiniaturize:(nullable id)sender;
- (void)performZoom:(nullable id)sender;
- (void)miniaturize:(nullable id)sender;
- (void)deminiaturize:(nullable id)sender;
- (void)zoom:(nullable id)sender;
- (void)toggleFullScreen:(nullable id)sender;
- (void)center;
- (void)display;
- (void)update;
- (void)invalidateShadow;
- (void)disableCursorRects;
- (void)enableCursorRects;
- (void)invalidateCursorRectsForView:(NSView *)view;
- (BOOL)makeFirstResponder:(nullable NSResponder *)responder;
- (void)sendEvent:(NSEvent *)event;
- (void)addChildWindow:(NSWindow *)childWin ordered:(NSWindowOrderingMode)place;
- (void)removeChildWindow:(NSWindow *)childWin;
- (nullable NSButton *)standardWindowButton:(NSWindowButton)b;
- (NSPoint)convertPointToScreen:(NSPoint)point;
- (NSPoint)convertPointFromScreen:(NSPoint)point;
- (NSRect)convertRectToScreen:(NSRect)rect;
- (NSRect)convertRectFromScreen:(NSRect)rect;
- (NSRect)convertRectToBacking:(NSRect)rect;
- (NSRect)convertRectFromBacking:(NSRect)rect;
- (NSPoint)mouseLocationOutsideOfEventStream;
- (void)performWindowDragWithEvent:(NSEvent *)event;
- (void)setIsVisible:(BOOL)flag;
- (void)setIsMiniaturized:(BOOL)flag;
- (void)setIsZoomed:(BOOL)flag;
- (void)setDocumentEdited:(BOOL)dirtyFlag;
- (void)setTitleWithRepresentedFilename:(NSString *)filename;
- (void)registerForDraggedTypes:(NSArray<NSPasteboardType> *)newTypes;
@property(readonly, getter=isSheet) BOOL sheet;
@property(nullable, readonly, strong) NSWindow *sheetParent;
@property(nullable, readonly, strong) NSWindow *attachedSheet;
@property(readonly, copy) NSArray<NSWindow *> *sheets;
+ (nullable NSButton *)standardWindowButton:(NSWindowButton)b forStyleMask:(NSWindowStyleMask)styleMask;
- (void)beginSheet:(NSWindow *)sheetWindow completionHandler:(void (^_Nullable)(NSModalResponse returnCode))handler;
- (void)endSheet:(NSWindow *)sheetWindow returnCode:(NSModalResponse)returnCode;
- (void)setLevel:(NSWindowLevel)level;
@end

@interface NSWindowTab : NSObject
@property(copy) NSString *title;
@end

@interface NSPanel : NSWindow
@property(getter=isFloatingPanel) BOOL floatingPanel;
@property BOOL becomesKeyOnlyIfNeeded;
@property BOOL worksWhenModal;
@end

// Controls

typedef NS_ENUM(NSInteger, NSControlStateValue) {
	NSControlStateValueMixed = -1,
	NSControlStateValueOff = 0,
	NSControlStateValueOn = 1,
};
typedef NS_ENUM(NSUInteger, NSTextAlignment) {
	NSTextAlignmentLeft = 0,
	NSTextAlignmentCenter = 1, // the other way around to iOS.
	NSTextAlignmentRight = 2,
	NSTextAlignmentJustified = 3,
	NSTextAlignmentNatural = 4,
};
typedef NS_ENUM(NSUInteger, NSCellImagePosition) {
	NSNoImage = 0,
	NSImageOnly = 1,
	NSImageLeft = 2,
	NSImageRight = 3,
	NSImageBelow = 4,
	NSImageAbove = 5,
	NSImageOverlaps = 6,
	NSImageLeading = 7,
	NSImageTrailing = 8,
};
typedef NS_ENUM(NSUInteger, NSImageScaling) {
	NSImageScaleProportionallyDown = 0,
	NSImageScaleAxesIndependently,
	NSImageScaleNone,
	NSImageScaleProportionallyUpOrDown,
};

@interface NSControl : NSView
@property(nullable, weak) id target;
@property(nullable) SEL action;
@property NSInteger tag;
@property(getter=isEnabled) BOOL enabled;
@property(getter=isHighlighted) BOOL highlighted;
@property(copy) NSString *stringValue;
@property(copy) NSAttributedString *attributedStringValue;
@property NSInteger integerValue;
@property int intValue;
@property double doubleValue;
@property float floatValue;
@property NSTextAlignment alignment;
@property(nullable, strong) NSFont *font;
@property(nullable, strong) id cell;
- (NSInteger)sendActionOn:(NSEventMask)mask;
- (BOOL)sendAction:(nullable SEL)action to:(nullable id)target;
- (void)sizeToFit;
@end

@interface NSButton : NSControl
+ (instancetype)buttonWithTitle:(NSString *)title target:(nullable id)target action:(nullable SEL)action;
+ (instancetype)checkboxWithTitle:(NSString *)title target:(nullable id)target action:(nullable SEL)action;
+ (instancetype)buttonWithImage:(NSImage *)image target:(nullable id)target action:(nullable SEL)action;
@property(copy) NSString *title;
@property(copy) NSString *alternateTitle;
@property(nullable, strong) NSImage *image;
@property NSCellImagePosition imagePosition;
@property NSImageScaling imageScaling;
@property NSControlStateValue state;
@property(copy) NSString *keyEquivalent;
@property NSEventModifierFlags keyEquivalentModifierMask;
@property BOOL allowsMixedState;
@property(getter=isBordered) BOOL bordered;
@property(getter=isTransparent) BOOL transparent;
- (void)setButtonType:(NSUInteger)type;
@end

@interface NSPopUpButton : NSButton
- (instancetype)initWithFrame:(NSRect)buttonFrame pullsDown:(BOOL)flag;
- (void)addItemWithTitle:(NSString *)title;
- (void)addItemsWithTitles:(NSArray<NSString *> *)itemTitles;
- (void)insertItemWithTitle:(NSString *)title atIndex:(NSInteger)index;
- (void)removeItemWithTitle:(NSString *)title;
- (void)removeItemAtIndex:(NSInteger)index;
- (void)removeAllItems;
- (void)selectItemAtIndex:(NSInteger)index;
- (void)selectItemWithTitle:(NSString *)title;
@property(readonly) NSInteger indexOfSelectedItem;
@property(nullable, readonly, strong) NSMenuItem *selectedItem;
@property(readonly, copy) NSArray<NSMenuItem *> *itemArray;
@property(readonly) NSInteger numberOfItems;
- (nullable NSMenuItem *)itemAtIndex:(NSInteger)index;
- (nullable NSString *)titleOfSelectedItem;
@end

@interface NSTextField : NSControl
+ (instancetype)labelWithString:(NSString *)stringValue;
+ (instancetype)wrappingLabelWithString:(NSString *)stringValue;
+ (instancetype)textFieldWithString:(NSString *)stringValue;
@property(nullable, copy) NSString *placeholderString;
@property(nullable, copy) NSColor *textColor;
@property(getter=isEditable) BOOL editable;
@property(getter=isSelectable) BOOL selectable;
@property(getter=isBordered) BOOL bordered;
@property(getter=isBezeled) BOOL bezeled;
@property BOOL drawsBackground;
@property NSInteger maximumNumberOfLines;
@property CGFloat preferredMaxLayoutWidth;
@end

typedef NS_ENUM(NSInteger, NSGridCellPlacement) {
	NSGridCellPlacementInherited = 0,
	NSGridCellPlacementNone,
	NSGridCellPlacementLeading,
	NSGridCellPlacementTop = NSGridCellPlacementLeading,
	NSGridCellPlacementTrailing,
	NSGridCellPlacementBottom = NSGridCellPlacementTrailing,
	NSGridCellPlacementCenter,
	NSGridCellPlacementFill,
};
typedef NS_ENUM(NSInteger, NSGridRowAlignment) {
	NSGridRowAlignmentInherited = 0,
	NSGridRowAlignmentNone,
	NSGridRowAlignmentFirstBaseline,
	NSGridRowAlignmentLastBaseline,
};

@interface NSGridRow : NSObject
@property CGFloat height;
@property CGFloat topPadding;
@property CGFloat bottomPadding;
@property NSGridCellPlacement yPlacement;
@property NSGridRowAlignment rowAlignment;
@property(getter=isHidden) BOOL hidden;
@end
@interface NSGridColumn : NSObject
@property CGFloat width;
@property CGFloat leadingPadding;
@property CGFloat trailingPadding;
@property NSGridCellPlacement xPlacement;
@property(getter=isHidden) BOOL hidden;
@end
@interface NSGridView : NSView
+ (instancetype)gridViewWithNumberOfColumns:(NSInteger)columnCount rows:(NSInteger)rowCount;
+ (instancetype)gridViewWithViews:(NSArray<NSArray<NSView *> *> *)rows;
@property(readonly) NSInteger numberOfRows;
@property(readonly) NSInteger numberOfColumns;
@property CGFloat rowSpacing;
@property CGFloat columnSpacing;
@property NSGridCellPlacement xPlacement;
@property NSGridCellPlacement yPlacement;
@property NSGridRowAlignment rowAlignment;
- (NSGridRow *)rowAtIndex:(NSInteger)index;
- (NSGridColumn *)columnAtIndex:(NSInteger)index;
- (NSGridRow *)addRowWithViews:(NSArray<NSView *> *)views;
- (NSGridColumn *)addColumnWithViews:(NSArray<NSView *> *)views;
@end

@interface NSAttributedString (NSAttributedStringAppKitAdditions)
@property(readonly) NSSize size;
- (void)drawAtPoint:(NSPoint)point;
- (void)drawInRect:(NSRect)rect;
@end

@interface NSFont : NSObject <NSCopying, NSSecureCoding>
+ (NSFont *)systemFontOfSize:(CGFloat)fontSize;
+ (NSFont *)boldSystemFontOfSize:(CGFloat)fontSize;
+ (NSFont *)titleBarFontOfSize:(CGFloat)fontSize;
+ (NSFont *)menuFontOfSize:(CGFloat)fontSize;
+ (NSFont *)labelFontOfSize:(CGFloat)fontSize;
+ (nullable NSFont *)fontWithName:(NSString *)fontName size:(CGFloat)fontSize;
@property(class, readonly) CGFloat systemFontSize;
@property(class, readonly) CGFloat smallSystemFontSize;
@property(class, readonly) CGFloat labelFontSize;
@property(readonly, copy) NSString *fontName;
@property(readonly, copy) NSString *familyName;
@property(nullable, readonly, copy) NSString *displayName;
@property(readonly) CGFloat pointSize;
@end

// Menus

typedef NS_ENUM(NSInteger, NSMenuItemBadgeType) {
	NSMenuItemBadgeTypeNone = 0,
};

@protocol NSMenuItemValidation <NSObject>
- (BOOL)validateMenuItem:(NSMenuItem *)menuItem;
@end

@interface NSMenuItem : NSObject <NSCopying, NSCoding>
- (instancetype)initWithTitle:(NSString *)string action:(nullable SEL)selector keyEquivalent:(NSString *)charCode;
@property(class, readonly, strong) NSMenuItem *separatorItem;
@property(readonly) BOOL isSeparatorItem;
@property(nullable, assign) NSMenu *menu;
@property(nullable, strong) NSMenu *submenu;
@property(readonly) BOOL hasSubmenu;
@property(nullable, readonly, assign) NSMenuItem *parentItem;
@property(copy) NSString *title;
@property(nullable, copy) NSAttributedString *attributedTitle;
@property(copy) NSString *keyEquivalent;
@property NSEventModifierFlags keyEquivalentModifierMask;
@property(nullable, strong) NSImage *image;
@property NSControlStateValue state;
@property(null_resettable, strong) NSImage *onStateImage;
@property(null_resettable, strong) NSImage *offStateImage;
@property(null_resettable, strong) NSImage *mixedStateImage;
@property(getter=isEnabled) BOOL enabled;
@property(getter=isHidden) BOOL hidden;
@property(getter=isHighlighted, readonly) BOOL highlighted;
@property(getter=isAlternate) BOOL alternate;
@property NSInteger indentationLevel;
@property(nullable, weak) id target;
@property(nullable) SEL action;
@property NSInteger tag;
@property(nullable, strong) id representedObject;
@property(nullable, strong) NSView *view;
@property(nullable, copy) NSString *toolTip;
@property BOOL allowsKeyEquivalentWhenHidden;
@end

@protocol NSMenuDelegate <NSObject>
@optional
- (void)menuNeedsUpdate:(NSMenu *)menu;
- (NSInteger)numberOfItemsInMenu:(NSMenu *)menu;
- (BOOL)menu:(NSMenu *)menu updateItem:(NSMenuItem *)item atIndex:(NSInteger)index shouldCancel:(BOOL)shouldCancel;
- (BOOL)menuHasKeyEquivalent:(NSMenu *)menu forEvent:(NSEvent *)event target:(id _Nullable *_Nonnull)target action:(SEL _Nullable *_Nonnull)action;
- (void)menuWillOpen:(NSMenu *)menu;
- (void)menuDidClose:(NSMenu *)menu;
- (void)menu:(NSMenu *)menu willHighlightItem:(nullable NSMenuItem *)item;
- (NSRect)confinementRectForMenu:(NSMenu *)menu onScreen:(nullable NSScreen *)screen;
@end

@interface NSMenu : NSObject <NSCopying, NSCoding>
- (instancetype)initWithTitle:(NSString *)title;
@property(copy) NSString *title;
@property(nullable, weak) id<NSMenuDelegate> delegate;
@property(nullable, readonly, assign) NSMenu *supermenu;
@property(readonly, copy) NSArray<NSMenuItem *> *itemArray;
@property(readonly) NSInteger numberOfItems;
@property BOOL autoenablesItems;
@property(nullable, readonly, strong) NSMenuItem *highlightedItem;
@property(readonly) NSSize size;
@property CGFloat minimumWidth;
@property(nullable, weak) NSFont *font;
@property BOOL allowsContextMenuPlugIns;
@property BOOL showsStateColumn;
@property NSUserInterfaceLayoutDirection userInterfaceLayoutDirection;
- (NSMenuItem *)addItemWithTitle:(NSString *)string action:(nullable SEL)selector keyEquivalent:(NSString *)charCode;
- (NSMenuItem *)insertItemWithTitle:(NSString *)string action:(nullable SEL)selector keyEquivalent:(NSString *)charCode atIndex:(NSInteger)index;
- (void)addItem:(NSMenuItem *)newItem;
- (void)insertItem:(NSMenuItem *)newItem atIndex:(NSInteger)index;
- (void)removeItem:(NSMenuItem *)item;
- (void)removeItemAtIndex:(NSInteger)index;
- (void)removeAllItems;
- (void)setSubmenu:(nullable NSMenu *)menu forItem:(NSMenuItem *)item;
- (nullable NSMenuItem *)itemAtIndex:(NSInteger)index;
- (nullable NSMenuItem *)itemWithTag:(NSInteger)tag;
- (nullable NSMenuItem *)itemWithTitle:(NSString *)title;
- (NSInteger)indexOfItem:(NSMenuItem *)item;
- (NSInteger)indexOfItemWithTag:(NSInteger)tag;
- (NSInteger)indexOfItemWithSubmenu:(nullable NSMenu *)submenu;
- (NSInteger)indexOfItemWithTitle:(NSString *)title;
- (NSInteger)indexOfItemWithRepresentedObject:(nullable id)object;
- (void)update;
- (void)cancelTracking;
- (void)cancelTrackingWithoutAnimation;
- (BOOL)performKeyEquivalent:(NSEvent *)event;
- (void)performActionForItemAtIndex:(NSInteger)index;
- (BOOL)popUpMenuPositioningItem:(nullable NSMenuItem *)item atLocation:(NSPoint)location inView:(nullable NSView *)view;
+ (void)popUpContextMenu:(NSMenu *)menu withEvent:(NSEvent *)event forView:(NSView *)view;
+ (void)setMenuBarVisible:(BOOL)visible;
+ (BOOL)menuBarVisible;
@end

static const CGFloat NSVariableStatusItemLength = -1.0;
static const CGFloat NSSquareStatusItemLength = -2.0;

@interface NSStatusBarButton : NSButton
@end
@interface NSStatusItem : NSObject
@property(nullable, readonly, weak) NSStatusBar *statusBar;
@property CGFloat length;
@property(nullable, strong) NSMenu *menu;
@property(nullable, readonly, strong) NSStatusBarButton *button;
@property(getter=isVisible) BOOL visible;
@property(nullable, copy) NSString *autosaveName;
@end
@interface NSStatusBar : NSObject
@property(class, readonly, strong) NSStatusBar *systemStatusBar;
- (NSStatusItem *)statusItemWithLength:(CGFloat)length;
- (void)removeStatusItem:(NSStatusItem *)item;
@property(readonly, getter=isVertical) BOOL vertical;
@property(readonly) CGFloat thickness;
@end

// Alerts & panels

typedef NS_ENUM(NSUInteger, NSAlertStyle) {
	NSAlertStyleWarning = 0,
	NSAlertStyleInformational = 1,
	NSAlertStyleCritical = 2,
};
static const NSModalResponse NSAlertFirstButtonReturn = 1000;
static const NSModalResponse NSAlertSecondButtonReturn = 1001;
static const NSModalResponse NSAlertThirdButtonReturn = 1002;

@interface NSAlert : NSObject
+ (NSAlert *)alertWithError:(NSError *)error;
@property(copy) NSString *messageText;
@property(copy) NSString *informativeText;
@property(null_resettable, strong) NSImage *icon;
- (NSButton *)addButtonWithTitle:(NSString *)title;
@property(readonly, copy) NSArray<NSButton *> *buttons;
@property BOOL showsHelp;
@property(nullable, copy) NSString *helpAnchor;
@property NSAlertStyle alertStyle;
@property BOOL showsSuppressionButton;
@property(nullable, readonly, strong) NSButton *suppressionButton;
@property(nullable, strong) NSView *accessoryView;
@property(readonly, strong) NSWindow *window;
- (void)layout;
- (NSModalResponse)runModal;
- (void)beginSheetModalForWindow:(NSWindow *)sheetWindow completionHandler:(void (^_Nullable)(NSModalResponse returnCode))handler;
@end

@protocol NSOpenSavePanelDelegate <NSObject>
@optional
- (BOOL)panel:(id)sender shouldEnableURL:(NSURL *)url;
- (BOOL)panel:(id)sender validateURL:(NSURL *)url error:(NSError **)outError;
- (void)panel:(id)sender didChangeToDirectoryURL:(nullable NSURL *)url;
- (nullable NSString *)panel:(id)sender userEnteredFilename:(NSString *)filename confirmed:(BOOL)okFlag;
- (void)panelSelectionDidChange:(nullable id)sender;
@end

@interface NSSavePanel : NSPanel
@property(class, readonly, strong) NSSavePanel *savePanel;
+ (NSSavePanel *)savePanel;
@property(nullable, readonly, copy) NSURL *URL;
@property(nullable, copy) NSURL *directoryURL;
@property(copy) NSArray<UTType *> *allowedContentTypes;
@property(nullable, copy) NSArray<NSString *> *allowedFileTypes;
@property BOOL allowsOtherFileTypes;
@property(nullable, strong) NSView *accessoryView;
@property(nullable, weak) id<NSOpenSavePanelDelegate> delegate;
@property(getter=isExpanded, readonly) BOOL expanded;
@property BOOL canCreateDirectories;
@property BOOL canSelectHiddenExtension;
@property(getter=isExtensionHidden) BOOL extensionHidden;
@property BOOL treatsFilePackagesAsDirectories;
@property(copy) NSString *prompt;
@property(copy) NSString *title;
@property(copy) NSString *nameFieldLabel;
@property(copy) NSString *nameFieldStringValue;
@property(copy) NSString *message;
@property BOOL showsHiddenFiles;
@property BOOL showsTagField;
- (void)beginSheetModalForWindow:(NSWindow *)window completionHandler:(void (^)(NSModalResponse result))handler;
- (void)beginWithCompletionHandler:(void (^)(NSModalResponse result))handler;
- (NSModalResponse)runModal;
- (void)validateVisibleColumns;
@end

@interface NSOpenPanel : NSSavePanel
@property(class, readonly, strong) NSOpenPanel *openPanel;
+ (NSOpenPanel *)openPanel;
@property(readonly, copy) NSArray<NSURL *> *URLs;
@property BOOL resolvesAliases;
@property BOOL canChooseDirectories;
@property BOOL allowsMultipleSelection;
@property BOOL canChooseFiles;
@property BOOL canDownloadUbiquitousContents;
@property BOOL canResolveUbiquitousConflicts;
@end

// Text input

@interface NSTextInputContext : NSObject
@property(class, nullable, readonly, strong) NSTextInputContext *currentInputContext;
- (instancetype)initWithClient:(id)client;
@property(readonly, strong) id client;
@property BOOL acceptsGlyphInfo;
@property(nullable, copy) NSArray<NSString *> *allowedInputSourceLocales;
@property(nullable, readonly, copy) NSString *selectedKeyboardInputSource;
+ (nullable NSString *)localizedNameForInputSource:(NSString *)inputSourceIdentifier;
- (void)activate;
- (void)deactivate;
- (BOOL)handleEvent:(NSEvent *)event;
- (void)discardMarkedText;
- (void)invalidateCharacterCoordinates;
@end

@protocol NSTextInputClient
- (void)insertText:(id)string replacementRange:(NSRange)replacementRange;
- (void)doCommandBySelector:(SEL)selector;
- (void)setMarkedText:(id)string selectedRange:(NSRange)selectedRange replacementRange:(NSRange)replacementRange;
- (void)unmarkText;
- (NSRange)selectedRange;
- (NSRange)markedRange;
- (BOOL)hasMarkedText;
- (nullable NSAttributedString *)attributedSubstringForProposedRange:(NSRange)range actualRange:(nullable NSRangePointer)actualRange;
- (NSArray<NSAttributedStringKey> *)validAttributesForMarkedText;
- (NSRect)firstRectForCharacterRange:(NSRange)range actualRange:(nullable NSRangePointer)actualRange;
- (NSUInteger)characterIndexForPoint:(NSPoint)point;
@optional
- (NSAttributedString *)attributedString;
- (CGFloat)fractionOfDistanceThroughGlyphForPoint:(NSPoint)point;
- (CGFloat)baselineDeltaForCharacterAtIndex:(NSUInteger)anIndex;
- (NSInteger)windowLevel;
- (BOOL)drawsVerticallyForCharacterAtIndex:(NSUInteger)charIndex;
@end

// OpenGL (deprecated on macOS, but the engine still offers it)

typedef uint32_t NSOpenGLPixelFormatAttribute;
enum {
	NSOpenGLPFAAllRenderers = 1,
	NSOpenGLPFATripleBuffer = 3,
	NSOpenGLPFADoubleBuffer = 5,
	NSOpenGLPFAAuxBuffers = 7,
	NSOpenGLPFAColorSize = 8,
	NSOpenGLPFAAlphaSize = 11,
	NSOpenGLPFADepthSize = 12,
	NSOpenGLPFAStencilSize = 13,
	NSOpenGLPFAAccumSize = 14,
	NSOpenGLPFAMinimumPolicy = 51,
	NSOpenGLPFAMaximumPolicy = 52,
	NSOpenGLPFASampleBuffers = 55,
	NSOpenGLPFASamples = 56,
	NSOpenGLPFAAuxDepthStencil = 57,
	NSOpenGLPFAColorFloat = 58,
	NSOpenGLPFAMultisample = 59,
	NSOpenGLPFASupersample = 60,
	NSOpenGLPFASampleAlpha = 61,
	NSOpenGLPFARendererID = 70,
	NSOpenGLPFANoRecovery = 72,
	NSOpenGLPFAAccelerated = 73,
	NSOpenGLPFAClosestPolicy = 74,
	NSOpenGLPFABackingStore = 76,
	NSOpenGLPFAScreenMask = 84,
	NSOpenGLPFAAllowOfflineRenderers = 96,
	NSOpenGLPFAAcceleratedCompute = 97,
	NSOpenGLPFAOpenGLProfile = 99,
	NSOpenGLPFAVirtualScreenCount = 128,
};
enum {
	NSOpenGLProfileVersionLegacy = 0x1000,
	NSOpenGLProfileVersion3_2Core = 0x3200,
	NSOpenGLProfileVersion4_1Core = 0x4100,
};
typedef NS_ENUM(NSInteger, NSOpenGLContextParameter) {
	NSOpenGLContextParameterSwapInterval = 222,
	NSOpenGLContextParameterSurfaceOrder = 235,
	NSOpenGLContextParameterSurfaceOpacity = 236,
	NSOpenGLContextParameterSurfaceBackingSize = 304,
	NSOpenGLContextParameterReclaimResources = 308,
	NSOpenGLContextParameterCurrentRendererID = 309,
	NSOpenGLContextParameterGPUVertexProcessing = 310,
	NSOpenGLContextParameterGPUFragmentProcessing = 311,
	NSOpenGLContextParameterHasDrawable = 314,
	NSOpenGLContextParameterMPSwapsInFlight = 315,
};

@interface NSOpenGLPixelFormat : NSObject <NSCoding>
- (nullable instancetype)initWithAttributes:(const NSOpenGLPixelFormatAttribute *)attribs;
@property(readonly) CGLPixelFormatObj CGLPixelFormatObj;
@end

@interface NSOpenGLContext : NSObject
- (nullable instancetype)initWithFormat:(NSOpenGLPixelFormat *)format shareContext:(nullable NSOpenGLContext *)share;
@property(nullable, strong) NSView *view;
@property(class, nullable, strong) NSOpenGLContext *currentContext;
+ (void)clearCurrentContext;
- (void)makeCurrentContext;
- (void)update;
- (void)flushBuffer;
- (void)setValues:(const GLint *)vals forParameter:(NSOpenGLContextParameter)param;
- (void)getValues:(GLint *)vals forParameter:(NSOpenGLContextParameter)param;
@property(readonly) CGLContextObj CGLContextObj;
@end

@interface NSOpenGLView : NSView
- (nullable instancetype)initWithFrame:(NSRect)frameRect pixelFormat:(nullable NSOpenGLPixelFormat *)format;
@property(nullable, strong) NSOpenGLContext *openGLContext;
@property(nullable, strong) NSOpenGLPixelFormat *pixelFormat;
@property BOOL wantsBestResolutionOpenGLSurface;
- (void)prepareOpenGL;
- (void)reshape;
- (void)update;
@end

// Speech

typedef NSString *NSSpeechSynthesizerVoiceName NS_TYPED_EXTENSIBLE_ENUM;
typedef NSString *NSVoiceAttributeKey NS_TYPED_ENUM;
typedef NSString *NSSpeechPropertyKey NS_TYPED_ENUM;
APPKIT_EXTERN NSVoiceAttributeKey const NSVoiceName;
APPKIT_EXTERN NSVoiceAttributeKey const NSVoiceIdentifier;
APPKIT_EXTERN NSVoiceAttributeKey const NSVoiceAge;
APPKIT_EXTERN NSVoiceAttributeKey const NSVoiceGender;
APPKIT_EXTERN NSVoiceAttributeKey const NSVoiceDemoText;
APPKIT_EXTERN NSVoiceAttributeKey const NSVoiceLocaleIdentifier;
APPKIT_EXTERN NSSpeechPropertyKey const NSSpeechPitchBaseProperty;
APPKIT_EXTERN NSSpeechPropertyKey const NSSpeechVolumeProperty;
APPKIT_EXTERN NSSpeechPropertyKey const NSSpeechRateProperty;
APPKIT_EXTERN NSSpeechPropertyKey const NSSpeechResetProperty;
typedef NS_ENUM(NSUInteger, NSSpeechBoundary) {
	NSSpeechImmediateBoundary = 0,
	NSSpeechWordBoundary,
	NSSpeechSentenceBoundary,
};

@class NSSpeechSynthesizer;
@protocol NSSpeechSynthesizerDelegate <NSObject>
@optional
- (void)speechSynthesizer:(NSSpeechSynthesizer *)sender didFinishSpeaking:(BOOL)finishedSpeaking;
- (void)speechSynthesizer:(NSSpeechSynthesizer *)sender willSpeakWord:(NSRange)characterRange ofString:(NSString *)string;
- (void)speechSynthesizer:(NSSpeechSynthesizer *)sender willSpeakPhoneme:(short)phonemeOpcode;
@end

@interface NSSpeechSynthesizer : NSObject
- (nullable instancetype)initWithVoice:(nullable NSSpeechSynthesizerVoiceName)voice;
- (BOOL)startSpeakingString:(NSString *)string;
- (BOOL)startSpeakingString:(NSString *)string toURL:(NSURL *)url;
@property(getter=isSpeaking, readonly) BOOL speaking;
- (void)stopSpeaking;
- (void)stopSpeakingAtBoundary:(NSSpeechBoundary)boundary;
- (void)pauseSpeakingAtBoundary:(NSSpeechBoundary)boundary;
- (void)continueSpeaking;
@property(nullable, weak) id<NSSpeechSynthesizerDelegate> delegate;
- (nullable NSSpeechSynthesizerVoiceName)voice;
- (BOOL)setVoice:(nullable NSSpeechSynthesizerVoiceName)voice;
@property float rate;
@property float volume;
@property BOOL usesFeedbackWindow;
- (nullable id)objectForProperty:(NSSpeechPropertyKey)property error:(NSError **)outError;
- (BOOL)setObject:(nullable id)object forProperty:(NSSpeechPropertyKey)property error:(NSError **)outError;
@property(class, readonly, getter=isAnyApplicationSpeaking) BOOL anyApplicationSpeaking;
@property(class, readonly) NSSpeechSynthesizerVoiceName defaultVoice;
@property(class, readonly, copy) NSArray<NSSpeechSynthesizerVoiceName> *availableVoices;
+ (NSDictionary<NSVoiceAttributeKey, id> *)attributesForVoice:(NSSpeechSynthesizerVoiceName)voice;
@end

NS_ASSUME_NONNULL_END

#endif
