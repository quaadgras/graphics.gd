// CoreHaptics declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREHAPTICS_H
#define GD_COREHAPTICS_H

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

#define CHHapticTimeImmediate 0.0f

typedef NSString *CHHapticPatternKey NS_TYPED_ENUM;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyVersion;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyPattern;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyEvent;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyEventType;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyTime;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyEventDuration;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyEventWaveformPath;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyEventParameters;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyParameter;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyParameterID;
FOUNDATION_EXTERN CHHapticPatternKey CHHapticPatternKeyParameterValue;

typedef NSString *CHHapticEventType NS_TYPED_ENUM;
FOUNDATION_EXTERN CHHapticEventType CHHapticEventTypeHapticTransient;
FOUNDATION_EXTERN CHHapticEventType CHHapticEventTypeHapticContinuous;
FOUNDATION_EXTERN CHHapticEventType CHHapticEventTypeAudioContinuous;
FOUNDATION_EXTERN CHHapticEventType CHHapticEventTypeAudioCustom;

typedef NSString *CHHapticEventParameterID NS_TYPED_ENUM;
FOUNDATION_EXTERN CHHapticEventParameterID CHHapticEventParameterIDHapticIntensity;
FOUNDATION_EXTERN CHHapticEventParameterID CHHapticEventParameterIDHapticSharpness;
FOUNDATION_EXTERN CHHapticEventParameterID CHHapticEventParameterIDAttackTime;
FOUNDATION_EXTERN CHHapticEventParameterID CHHapticEventParameterIDDecayTime;
FOUNDATION_EXTERN CHHapticEventParameterID CHHapticEventParameterIDReleaseTime;
FOUNDATION_EXTERN CHHapticEventParameterID CHHapticEventParameterIDSustained;

typedef NSString *CHHapticDynamicParameterID NS_TYPED_ENUM;
FOUNDATION_EXTERN CHHapticDynamicParameterID CHHapticDynamicParameterIDHapticIntensityControl;
FOUNDATION_EXTERN CHHapticDynamicParameterID CHHapticDynamicParameterIDHapticSharpnessControl;

typedef NS_ENUM(NSInteger, CHHapticEngineStoppedReason) {
	CHHapticEngineStoppedReasonAudioSessionInterrupt = 1,
	CHHapticEngineStoppedReasonApplicationSuspended = 2,
	CHHapticEngineStoppedReasonIdleTimeout = 3,
	CHHapticEngineStoppedReasonNotifyWhenFinished = 4,
	CHHapticEngineStoppedReasonEngineDestroyed = 5,
	CHHapticEngineStoppedReasonGameControllerDisconnect = 6,
	CHHapticEngineStoppedReasonSystemError = -1,
};
typedef NS_ENUM(NSInteger, CHHapticEngineFinishedAction) {
	CHHapticEngineFinishedActionStopEngine = 1,
	CHHapticEngineFinishedActionLeaveEngineRunning = 2,
};

typedef void (^CHHapticCompletionHandler)(NSError *_Nullable error);
typedef void (^CHHapticEngineStoppedHandler)(CHHapticEngineStoppedReason stoppedReason);
typedef void (^CHHapticEngineResetHandler)(void);
typedef CHHapticEngineFinishedAction (^CHHapticEngineFinishedHandler)(NSError *_Nullable error);

@protocol CHHapticDeviceCapability
@property(readonly) BOOL supportsHaptics;
@property(readonly) BOOL supportsAudio;
@end

@interface CHHapticEventParameter : NSObject
@property(readonly) CHHapticEventParameterID parameterID;
@property float value;
- (instancetype)initWithParameterID:(CHHapticEventParameterID)parameterID value:(float)value;
@end

@interface CHHapticDynamicParameter : NSObject
@property(readonly) CHHapticDynamicParameterID parameterID;
@property float value;
@property NSTimeInterval relativeTime;
- (instancetype)initWithParameterID:(CHHapticDynamicParameterID)parameterID value:(float)value relativeTime:(NSTimeInterval)time;
@end

@interface CHHapticEvent : NSObject
@property(readonly) CHHapticEventType type;
@property NSTimeInterval relativeTime;
@property NSTimeInterval duration;
- (instancetype)initWithEventType:(CHHapticEventType)type parameters:(NSArray<CHHapticEventParameter *> *)eventParams relativeTime:(NSTimeInterval)time;
- (instancetype)initWithEventType:(CHHapticEventType)type parameters:(NSArray<CHHapticEventParameter *> *)eventParams relativeTime:(NSTimeInterval)time duration:(NSTimeInterval)duration;
@end

@interface CHHapticPattern : NSObject
@property(readonly) NSTimeInterval duration;
- (nullable instancetype)initWithEvents:(NSArray<CHHapticEvent *> *)events parameters:(NSArray<CHHapticDynamicParameter *> *)parameters error:(NSError **)outError;
- (nullable instancetype)initWithDictionary:(NSDictionary<CHHapticPatternKey, id> *)patternDict error:(NSError **)outError;
@end

@protocol CHHapticPatternPlayer <NSObject>
- (BOOL)startAtTime:(NSTimeInterval)time error:(NSError **)outError;
- (BOOL)stopAtTime:(NSTimeInterval)time error:(NSError **)outError;
- (BOOL)sendParameters:(NSArray<CHHapticDynamicParameter *> *)parameters atTime:(NSTimeInterval)time error:(NSError **)outError;
- (BOOL)cancelAndReturnError:(NSError **)outError;
@property(readwrite) BOOL isMuted;
@end

@interface CHHapticEngine : NSObject
+ (id<CHHapticDeviceCapability>)capabilitiesForHardware;
@property(readonly) NSTimeInterval currentTime;
@property(readwrite, atomic) CHHapticEngineStoppedHandler stoppedHandler;
@property(readwrite, atomic) CHHapticEngineResetHandler resetHandler;
@property(readwrite, nonatomic) BOOL playsHapticsOnly;
@property(readwrite, nonatomic) BOOL isMutedForAudio;
@property(readwrite, nonatomic) BOOL isMutedForHaptics;
@property(readwrite, nonatomic, getter=isAutoShutdownEnabled) BOOL autoShutdownEnabled;
- (nullable instancetype)initAndReturnError:(NSError **)error;
- (void)startWithCompletionHandler:(nullable CHHapticCompletionHandler)completionHandler;
- (BOOL)startAndReturnError:(NSError **)outError;
- (void)stopWithCompletionHandler:(nullable CHHapticCompletionHandler)completionHandler;
- (void)notifyWhenPlayersFinished:(CHHapticEngineFinishedHandler)finishedHandler;
- (nullable id<CHHapticPatternPlayer>)createPlayerWithPattern:(CHHapticPattern *)pattern error:(NSError **)outError;
@end

NS_ASSUME_NONNULL_END

#endif
