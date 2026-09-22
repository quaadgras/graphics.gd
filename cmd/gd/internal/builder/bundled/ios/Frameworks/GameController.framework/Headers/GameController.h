// GameController declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_GAMECONTROLLER_H
#define GD_GAMECONTROLLER_H

#import <CoreHaptics/CoreHaptics.h>
#import <Foundation/Foundation.h>
#include <dispatch/dispatch.h>

#if defined(__cplusplus)
#define GAMECONTROLLER_EXTERN extern "C"
#else
#define GAMECONTROLLER_EXTERN extern
#endif
#define GAMECONTROLLER_EXPORT GAMECONTROLLER_EXTERN

NS_ASSUME_NONNULL_BEGIN

@class GCController, GCControllerElement, GCControllerButtonInput, GCControllerAxisInput, GCControllerDirectionPad, GCExtendedGamepad, GCMicroGamepad, GCGamepad, GCMotion, GCPhysicalInputProfile, GCDeviceHaptics, GCDeviceLight, GCDeviceBattery, GCKeyboard, GCMouse, GCKeyboardInput, GCMouseInput, GCControllerTouchpad;

GAMECONTROLLER_EXPORT NSNotificationName const GCControllerDidConnectNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCControllerDidDisconnectNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCControllerDidBecomeCurrentNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCControllerDidStopBeingCurrentNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCKeyboardDidConnectNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCKeyboardDidDisconnectNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCMouseDidConnectNotification;
GAMECONTROLLER_EXPORT NSNotificationName const GCMouseDidDisconnectNotification;

// Element names

GAMECONTROLLER_EXPORT NSString *const GCInputButtonA;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonB;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonX;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonY;
GAMECONTROLLER_EXPORT NSString *const GCInputDirectionPad;
GAMECONTROLLER_EXPORT NSString *const GCInputLeftThumbstick;
GAMECONTROLLER_EXPORT NSString *const GCInputRightThumbstick;
GAMECONTROLLER_EXPORT NSString *const GCInputLeftShoulder;
GAMECONTROLLER_EXPORT NSString *const GCInputRightShoulder;
GAMECONTROLLER_EXPORT NSString *const GCInputLeftTrigger;
GAMECONTROLLER_EXPORT NSString *const GCInputRightTrigger;
GAMECONTROLLER_EXPORT NSString *const GCInputLeftThumbstickButton;
GAMECONTROLLER_EXPORT NSString *const GCInputRightThumbstickButton;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonHome;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonMenu;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonOptions;
GAMECONTROLLER_EXPORT NSString *const GCInputButtonShare;
GAMECONTROLLER_EXPORT NSString *const GCInputXboxPaddleOne;
GAMECONTROLLER_EXPORT NSString *const GCInputXboxPaddleTwo;
GAMECONTROLLER_EXPORT NSString *const GCInputXboxPaddleThree;
GAMECONTROLLER_EXPORT NSString *const GCInputXboxPaddleFour;
GAMECONTROLLER_EXPORT NSString *const GCInputDualShockTouchpadOne;
GAMECONTROLLER_EXPORT NSString *const GCInputDualShockTouchpadTwo;
GAMECONTROLLER_EXPORT NSString *const GCInputDualShockTouchpadButton;

// Elements

typedef NS_ENUM(NSInteger, GCSystemGestureState) {
	GCSystemGestureStateEnabled = 0,
	GCSystemGestureStateAlwaysReceive,
	GCSystemGestureStateDisabled,
};

@interface GCControllerElement : NSObject
@property(nonatomic, weak, readonly, nullable) GCControllerElement *collection;
@property(nonatomic, readonly, getter=isAnalog) BOOL analog;
@property(nonatomic, readonly, getter=isBoundToSystemGesture) BOOL boundToSystemGesture;
@property(nonatomic, readwrite) GCSystemGestureState preferredSystemGestureState;
@property(nonatomic, strong, nullable) NSString *sfSymbolsName;
@property(nonatomic, strong, nullable) NSString *localizedName;
@property(nonatomic, strong, nullable) NSString *unmappedLocalizedName;
@property(nonatomic, readonly) NSSet<NSString *> *aliases;
@end

typedef void (^GCControllerButtonValueChangedHandler)(GCControllerButtonInput *button, float value, BOOL pressed);
typedef void (^GCControllerButtonTouchedChangedHandler)(GCControllerButtonInput *button, float value, BOOL pressed, BOOL touched);

@interface GCControllerButtonInput : GCControllerElement
@property(nonatomic, copy, nullable) GCControllerButtonValueChangedHandler valueChangedHandler;
@property(nonatomic, copy, nullable) GCControllerButtonValueChangedHandler pressedChangedHandler;
@property(nonatomic, copy, nullable) GCControllerButtonTouchedChangedHandler touchedChangedHandler;
@property(nonatomic, readonly) float value;
@property(nonatomic, readonly, getter=isPressed) BOOL pressed;
@property(nonatomic, readonly, getter=isTouched) BOOL touched;
@end

typedef void (^GCControllerAxisValueChangedHandler)(GCControllerAxisInput *axis, float value);

@interface GCControllerAxisInput : GCControllerElement
@property(nonatomic, copy, nullable) GCControllerAxisValueChangedHandler valueChangedHandler;
@property(nonatomic, readonly) float value;
@end

typedef void (^GCControllerDirectionPadValueChangedHandler)(GCControllerDirectionPad *dpad, float xValue, float yValue);

@interface GCControllerDirectionPad : GCControllerElement
@property(nonatomic, copy, nullable) GCControllerDirectionPadValueChangedHandler valueChangedHandler;
@property(nonatomic, readonly) GCControllerAxisInput *xAxis;
@property(nonatomic, readonly) GCControllerAxisInput *yAxis;
@property(nonatomic, readonly) GCControllerButtonInput *up;
@property(nonatomic, readonly) GCControllerButtonInput *down;
@property(nonatomic, readonly) GCControllerButtonInput *left;
@property(nonatomic, readonly) GCControllerButtonInput *right;
@end

@interface GCControllerTouchpad : GCControllerElement
@property(nonatomic, readonly) GCControllerButtonInput *button;
@property(nonatomic, readonly) GCControllerDirectionPad *touchSurface;
@end

// Profiles

@protocol GCDevice <NSObject>
@property(nonatomic, strong) dispatch_queue_t handlerQueue;
@property(nonatomic, readonly, copy, nullable) NSString *vendorName;
@property(nonatomic, readonly) NSString *productCategory;
@property(nonatomic, strong, readonly) GCPhysicalInputProfile *physicalInputProfile;
@end

@interface GCPhysicalInputProfile : NSObject
@property(nonatomic, readonly, weak, nullable) id<GCDevice> device;
@property(atomic, readonly) NSTimeInterval lastEventTimestamp;
@property(atomic, readonly) BOOL hasRemappedElements;
@property(nonatomic, copy, nullable) void (^valueDidChangeHandler)(__kindof GCPhysicalInputProfile *profile, GCControllerElement *element);
@property(nonatomic, readonly, strong) NSDictionary<NSString *, GCControllerElement *> *elements;
@property(nonatomic, readonly, strong) NSDictionary<NSString *, GCControllerButtonInput *> *buttons;
@property(nonatomic, readonly, strong) NSDictionary<NSString *, GCControllerAxisInput *> *axes;
@property(nonatomic, readonly, strong) NSDictionary<NSString *, GCControllerDirectionPad *> *dpads;
@property(nonatomic, readonly, strong) NSSet<GCControllerElement *> *allElements;
@property(nonatomic, readonly, strong) NSSet<GCControllerButtonInput *> *allButtons;
@property(nonatomic, readonly, strong) NSSet<GCControllerAxisInput *> *allAxes;
@property(nonatomic, readonly, strong) NSSet<GCControllerDirectionPad *> *allDpads;
- (__kindof GCControllerElement *_Nullable)objectForKeyedSubscript:(NSString *)key;
@end

typedef void (^GCExtendedGamepadValueChangedHandler)(GCExtendedGamepad *gamepad, GCControllerElement *element);

@interface GCExtendedGamepad : GCPhysicalInputProfile
@property(nonatomic, readonly, weak, nullable) GCController *controller;
@property(nonatomic, copy, nullable) GCExtendedGamepadValueChangedHandler valueChangedHandler;
@property(nonatomic, readonly) GCControllerDirectionPad *dpad;
@property(nonatomic, readonly) GCControllerButtonInput *buttonA;
@property(nonatomic, readonly) GCControllerButtonInput *buttonB;
@property(nonatomic, readonly) GCControllerButtonInput *buttonX;
@property(nonatomic, readonly) GCControllerButtonInput *buttonY;
@property(nonatomic, readonly) GCControllerButtonInput *buttonMenu;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *buttonOptions;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *buttonHome;
@property(nonatomic, readonly) GCControllerDirectionPad *leftThumbstick;
@property(nonatomic, readonly) GCControllerDirectionPad *rightThumbstick;
@property(nonatomic, readonly) GCControllerButtonInput *leftShoulder;
@property(nonatomic, readonly) GCControllerButtonInput *rightShoulder;
@property(nonatomic, readonly) GCControllerButtonInput *leftTrigger;
@property(nonatomic, readonly) GCControllerButtonInput *rightTrigger;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *leftThumbstickButton;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *rightThumbstickButton;
@end

@interface GCDualShockGamepad : GCExtendedGamepad
@property(nonatomic, readonly) GCControllerButtonInput *touchpadButton;
@property(nonatomic, readonly) GCControllerDirectionPad *touchpadPrimary;
@property(nonatomic, readonly) GCControllerDirectionPad *touchpadSecondary;
@end

@interface GCDualSenseGamepad : GCExtendedGamepad
@property(nonatomic, readonly) GCControllerButtonInput *touchpadButton;
@property(nonatomic, readonly) GCControllerDirectionPad *touchpadPrimary;
@property(nonatomic, readonly) GCControllerDirectionPad *touchpadSecondary;
@end

@interface GCXboxGamepad : GCExtendedGamepad
@property(nonatomic, readonly, nullable) GCControllerButtonInput *paddleButton1;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *paddleButton2;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *paddleButton3;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *paddleButton4;
@property(nonatomic, readonly, nullable) GCControllerButtonInput *buttonShare;
@end

typedef void (^GCMicroGamepadValueChangedHandler)(GCMicroGamepad *gamepad, GCControllerElement *element);

@interface GCMicroGamepad : GCPhysicalInputProfile
@property(nonatomic, readonly, weak, nullable) GCController *controller;
@property(nonatomic, copy, nullable) GCMicroGamepadValueChangedHandler valueChangedHandler;
@property(nonatomic, readonly, retain) GCControllerDirectionPad *dpad;
@property(nonatomic, readonly, retain) GCControllerButtonInput *buttonA;
@property(nonatomic, readonly, retain) GCControllerButtonInput *buttonX;
@property(nonatomic, readonly) GCControllerButtonInput *buttonMenu;
@property(nonatomic, assign) BOOL reportsAbsoluteDpadValues;
@property(nonatomic, assign) BOOL allowsRotation;
@end

// Motion

typedef struct {
	double x, y, z;
} GCAcceleration;
typedef struct {
	double x, y, z;
} GCRotationRate;
typedef struct {
	double pitch, yaw, roll;
} GCEulerAngles;
typedef struct GCQuaternion {
	double x, y, z, w;
} GCQuaternion;

typedef void (^GCMotionValueChangedHandler)(GCMotion *motion);

@interface GCMotion : NSObject
@property(nonatomic, assign, readonly, nullable) GCController *controller;
@property(nonatomic, copy, nullable) GCMotionValueChangedHandler valueChangedHandler;
@property(nonatomic, readonly) BOOL sensorsRequireManualActivation;
@property(nonatomic) BOOL sensorsActive;
@property(nonatomic, readonly) BOOL hasGravityAndUserAcceleration;
@property(nonatomic, readonly) BOOL hasAttitudeAndRotationRate;
@property(nonatomic, readonly) BOOL hasAttitude;
@property(nonatomic, readonly) BOOL hasRotationRate;
@property(nonatomic, readonly) GCAcceleration gravity;
@property(nonatomic, readonly) GCAcceleration userAcceleration;
@property(nonatomic, readonly) GCAcceleration acceleration;
@property(nonatomic, readonly) GCQuaternion attitude;
@property(nonatomic, readonly) GCRotationRate rotationRate;
@end

// Haptics, light & battery

typedef NSString *GCHapticsLocality NS_TYPED_ENUM;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityDefault;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityAll;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityHandles;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityLeftHandle;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityRightHandle;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityTriggers;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityLeftTrigger;
GAMECONTROLLER_EXPORT GCHapticsLocality const GCHapticsLocalityRightTrigger;
GAMECONTROLLER_EXPORT const float GCHapticDurationInfinite;

@interface GCDeviceHaptics : NSObject
@property(nonatomic, readonly, copy) NSSet<GCHapticsLocality> *supportedLocalities;
- (nullable CHHapticEngine *)createEngineWithLocality:(GCHapticsLocality)locality;
@end

@interface GCColor : NSObject <NSCopying, NSSecureCoding>
@property(readonly) float red;
@property(readonly) float green;
@property(readonly) float blue;
- (instancetype)initWithRed:(float)red green:(float)green blue:(float)blue;
@end

@interface GCDeviceLight : NSObject
@property(nonatomic, copy) GCColor *color;
@end

typedef NS_ENUM(NSInteger, GCDeviceBatteryState) {
	GCDeviceBatteryStateUnknown = -1,
	GCDeviceBatteryStateDischarging,
	GCDeviceBatteryStateCharging,
	GCDeviceBatteryStateFull,
};

@interface GCDeviceBattery : NSObject
@property(nonatomic, readonly) float batteryLevel;
@property(nonatomic, readonly) GCDeviceBatteryState batteryState;
@end

// GCController

typedef NS_ENUM(NSInteger, GCControllerPlayerIndex) {
	GCControllerPlayerIndexUnset = -1,
	GCControllerPlayerIndex1 = 0,
	GCControllerPlayerIndex2,
	GCControllerPlayerIndex3,
	GCControllerPlayerIndex4,
};

@interface GCController : NSObject <GCDevice>
@property(nonatomic, copy, nullable) void (^controllerPausedHandler)(GCController *controller);
@property(class, atomic, strong, readonly, nullable) GCController *current;
@property(class, nonatomic, readwrite) BOOL shouldMonitorBackgroundEvents;
@property(nonatomic, readonly, getter=isAttachedToDevice) BOOL attachedToDevice;
@property(nonatomic, readonly, getter=isSnapshot) BOOL snapshot;
@property(nonatomic) GCControllerPlayerIndex playerIndex;
@property(nonatomic, copy, readonly, nullable) GCDeviceBattery *battery;
@property(nonatomic, retain, readonly, nullable) GCDeviceHaptics *haptics;
@property(nonatomic, retain, readonly, nullable) GCDeviceLight *light;
@property(nonatomic, retain, readonly, nullable) GCGamepad *gamepad;
@property(nonatomic, retain, readonly, nullable) GCMicroGamepad *microGamepad;
@property(nonatomic, retain, readonly, nullable) GCExtendedGamepad *extendedGamepad;
@property(nonatomic, retain, readonly, nullable) GCMotion *motion;
+ (NSArray<GCController *> *)controllers;
+ (void)startWirelessControllerDiscoveryWithCompletionHandler:(nullable void (^)(void))completionHandler;
+ (void)stopWirelessControllerDiscovery;
@end

// Keyboard & mouse

typedef NSString *GCKeyCode NS_TYPED_EXTENSIBLE_ENUM;

@interface GCKeyboardInput : GCPhysicalInputProfile
@property(nonatomic, readonly, getter=isAnyKeyPressed) BOOL anyKeyPressed;
@end

@interface GCKeyboard : NSObject <GCDevice>
@property(nonatomic, strong, readonly, nullable) GCKeyboardInput *keyboardInput;
@property(class, atomic, strong, readonly, nullable) GCKeyboard *coalescedKeyboard;
@end

@interface GCMouseInput : GCPhysicalInputProfile
@end

@interface GCMouse : NSObject <GCDevice>
@property(nonatomic, strong, readonly, nullable) GCMouseInput *mouseInput;
@property(class, atomic, strong, readonly, nullable) GCMouse *current;
+ (NSArray<GCMouse *> *)mice;
@end

// GCEventViewController

NS_ASSUME_NONNULL_END

#if __has_include(<GameController/GCController+macOS.h>)
#import <GameController/GCController+macOS.h>
#endif

#endif
