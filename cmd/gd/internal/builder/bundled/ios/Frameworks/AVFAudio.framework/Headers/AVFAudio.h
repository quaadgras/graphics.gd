// AVFAudio declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_AVFAUDIO_H
#define GD_AVFAUDIO_H

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

// AVAudioSession

typedef NSString *AVAudioSessionCategory NS_STRING_ENUM;
FOUNDATION_EXTERN AVAudioSessionCategory const AVAudioSessionCategoryAmbient;
FOUNDATION_EXTERN AVAudioSessionCategory const AVAudioSessionCategorySoloAmbient;
FOUNDATION_EXTERN AVAudioSessionCategory const AVAudioSessionCategoryPlayback;
FOUNDATION_EXTERN AVAudioSessionCategory const AVAudioSessionCategoryRecord;
FOUNDATION_EXTERN AVAudioSessionCategory const AVAudioSessionCategoryPlayAndRecord;
FOUNDATION_EXTERN AVAudioSessionCategory const AVAudioSessionCategoryMultiRoute;

typedef NSString *AVAudioSessionMode NS_STRING_ENUM;
FOUNDATION_EXTERN AVAudioSessionMode const AVAudioSessionModeDefault;
FOUNDATION_EXTERN AVAudioSessionMode const AVAudioSessionModeGameChat;
FOUNDATION_EXTERN AVAudioSessionMode const AVAudioSessionModeVoiceChat;
FOUNDATION_EXTERN AVAudioSessionMode const AVAudioSessionModeMeasurement;

typedef NS_OPTIONS(NSUInteger, AVAudioSessionCategoryOptions) {
	AVAudioSessionCategoryOptionMixWithOthers = 0x1,
	AVAudioSessionCategoryOptionDuckOthers = 0x2,
	AVAudioSessionCategoryOptionAllowBluetooth = 0x4,
	AVAudioSessionCategoryOptionDefaultToSpeaker = 0x8,
	AVAudioSessionCategoryOptionInterruptSpokenAudioAndMixWithOthers = 0x11,
	AVAudioSessionCategoryOptionAllowBluetoothA2DP = 0x20,
	AVAudioSessionCategoryOptionAllowAirPlay = 0x40,
	AVAudioSessionCategoryOptionOverrideMutedMicrophoneInterruption = 0x80,
};
typedef NS_OPTIONS(NSUInteger, AVAudioSessionSetActiveOptions) {
	AVAudioSessionSetActiveOptionNotifyOthersOnDeactivation = 1,
};
typedef NS_ENUM(NSUInteger, AVAudioSessionInterruptionType) {
	AVAudioSessionInterruptionTypeBegan = 1,
	AVAudioSessionInterruptionTypeEnded = 0,
};
typedef NS_OPTIONS(NSUInteger, AVAudioSessionInterruptionOptions) {
	AVAudioSessionInterruptionOptionShouldResume = 1,
};
typedef NS_ENUM(NSUInteger, AVAudioSessionRouteChangeReason) {
	AVAudioSessionRouteChangeReasonUnknown = 0,
	AVAudioSessionRouteChangeReasonNewDeviceAvailable = 1,
	AVAudioSessionRouteChangeReasonOldDeviceUnavailable = 2,
	AVAudioSessionRouteChangeReasonCategoryChange = 3,
	AVAudioSessionRouteChangeReasonOverride = 4,
	AVAudioSessionRouteChangeReasonWakeFromSleep = 6,
	AVAudioSessionRouteChangeReasonNoSuitableRouteForCategory = 7,
	AVAudioSessionRouteChangeReasonRouteConfigurationChange = 8,
};
typedef NS_ENUM(NSUInteger, AVAudioSessionRecordPermission) {
	AVAudioSessionRecordPermissionUndetermined = 'undt',
	AVAudioSessionRecordPermissionDenied = 'deny',
	AVAudioSessionRecordPermissionGranted = 'grnt',
};

FOUNDATION_EXTERN NSNotificationName const AVAudioSessionInterruptionNotification;
FOUNDATION_EXTERN NSNotificationName const AVAudioSessionRouteChangeNotification;
FOUNDATION_EXTERN NSNotificationName const AVAudioSessionMediaServicesWereLostNotification;
FOUNDATION_EXTERN NSNotificationName const AVAudioSessionMediaServicesWereResetNotification;
FOUNDATION_EXTERN NSString *const AVAudioSessionInterruptionTypeKey;
FOUNDATION_EXTERN NSString *const AVAudioSessionInterruptionOptionKey;
FOUNDATION_EXTERN NSString *const AVAudioSessionRouteChangeReasonKey;

@interface AVAudioSession : NSObject
+ (AVAudioSession *)sharedInstance;
@property(readonly) AVAudioSessionCategory category;
@property(readonly) AVAudioSessionCategoryOptions categoryOptions;
@property(readonly) AVAudioSessionMode mode;
- (BOOL)setCategory:(AVAudioSessionCategory)category error:(NSError **)outError;
- (BOOL)setCategory:(AVAudioSessionCategory)category withOptions:(AVAudioSessionCategoryOptions)options error:(NSError **)outError;
- (BOOL)setCategory:(AVAudioSessionCategory)category mode:(AVAudioSessionMode)mode options:(AVAudioSessionCategoryOptions)options error:(NSError **)outError;
- (BOOL)setMode:(AVAudioSessionMode)mode error:(NSError **)outError;
- (BOOL)setActive:(BOOL)active error:(NSError **)outError;
- (BOOL)setActive:(BOOL)active withOptions:(AVAudioSessionSetActiveOptions)options error:(NSError **)outError;
- (BOOL)setPreferredSampleRate:(double)sampleRate error:(NSError **)outError;
- (BOOL)setPreferredIOBufferDuration:(NSTimeInterval)duration error:(NSError **)outError;
@property(readonly) double sampleRate;
@property(readonly) double preferredSampleRate;
@property(readonly) NSTimeInterval IOBufferDuration;
@property(readonly) NSTimeInterval inputLatency;
@property(readonly) NSTimeInterval outputLatency;
@property(readonly) NSInteger inputNumberOfChannels;
@property(readonly) NSInteger outputNumberOfChannels;
@property(readonly) float outputVolume;
@property(readonly, getter=isInputAvailable) BOOL inputAvailable;
@property(readonly, getter=isOtherAudioPlaying) BOOL otherAudioPlaying;
@property(readonly) AVAudioSessionRecordPermission recordPermission;
- (void)requestRecordPermission:(void (^)(BOOL granted))response;
@end

// AVAudioApplication

typedef NS_ENUM(NSInteger, AVAudioApplicationRecordPermission) {
	AVAudioApplicationRecordPermissionUndetermined = 'undt',
	AVAudioApplicationRecordPermissionDenied = 'deny',
	AVAudioApplicationRecordPermissionGranted = 'grnt',
};

API_AVAILABLE(ios(17.0))
@interface AVAudioApplication : NSObject
@property(class, readonly) AVAudioApplication *sharedInstance;
@property(readonly) AVAudioApplicationRecordPermission recordPermission;
+ (void)requestRecordPermissionWithCompletionHandler:(void (^)(BOOL granted))response;
@end

// AVSpeechSynthesis

typedef NS_ENUM(NSInteger, AVSpeechBoundary) {
	AVSpeechBoundaryImmediate,
	AVSpeechBoundaryWord,
};
typedef NS_ENUM(NSInteger, AVSpeechSynthesisVoiceQuality) {
	AVSpeechSynthesisVoiceQualityDefault = 1,
	AVSpeechSynthesisVoiceQualityEnhanced,
	AVSpeechSynthesisVoiceQualityPremium,
};

FOUNDATION_EXTERN const float AVSpeechUtteranceMinimumSpeechRate;
FOUNDATION_EXTERN const float AVSpeechUtteranceMaximumSpeechRate;
FOUNDATION_EXTERN const float AVSpeechUtteranceDefaultSpeechRate;

@protocol AVSpeechSynthesizerDelegate;

@interface AVSpeechSynthesisVoice : NSObject <NSSecureCoding>
+ (NSArray<AVSpeechSynthesisVoice *> *)speechVoices;
+ (NSString *)currentLanguageCode;
+ (nullable AVSpeechSynthesisVoice *)voiceWithLanguage:(nullable NSString *)languageCode;
+ (nullable AVSpeechSynthesisVoice *)voiceWithIdentifier:(NSString *)identifier;
@property(nonatomic, readonly) NSString *language;
@property(nonatomic, readonly) NSString *identifier;
@property(nonatomic, readonly) NSString *name;
@property(nonatomic, readonly) AVSpeechSynthesisVoiceQuality quality;
@end

@interface AVSpeechUtterance : NSObject <NSCopying, NSSecureCoding>
+ (instancetype)speechUtteranceWithString:(NSString *)string;
- (instancetype)initWithString:(NSString *)string;
@property(nonatomic, retain, nullable) AVSpeechSynthesisVoice *voice;
@property(nonatomic, readonly) NSString *speechString;
@property(nonatomic) float rate;
@property(nonatomic) float pitchMultiplier;
@property(nonatomic) float volume;
@property(nonatomic) NSTimeInterval preUtteranceDelay;
@property(nonatomic) NSTimeInterval postUtteranceDelay;
@end

@interface AVSpeechSynthesizer : NSObject
@property(nonatomic, weak, nullable) id<AVSpeechSynthesizerDelegate> delegate;
@property(nonatomic, readonly, getter=isSpeaking) BOOL speaking;
@property(nonatomic, readonly, getter=isPaused) BOOL paused;
- (void)speakUtterance:(AVSpeechUtterance *)utterance;
- (BOOL)stopSpeakingAtBoundary:(AVSpeechBoundary)boundary;
- (BOOL)pauseSpeakingAtBoundary:(AVSpeechBoundary)boundary;
- (BOOL)continueSpeaking;
@end

@protocol AVSpeechSynthesizerDelegate <NSObject>
@optional
- (void)speechSynthesizer:(AVSpeechSynthesizer *)synthesizer didStartSpeechUtterance:(AVSpeechUtterance *)utterance;
- (void)speechSynthesizer:(AVSpeechSynthesizer *)synthesizer didFinishSpeechUtterance:(AVSpeechUtterance *)utterance;
- (void)speechSynthesizer:(AVSpeechSynthesizer *)synthesizer didPauseSpeechUtterance:(AVSpeechUtterance *)utterance;
- (void)speechSynthesizer:(AVSpeechSynthesizer *)synthesizer didContinueSpeechUtterance:(AVSpeechUtterance *)utterance;
- (void)speechSynthesizer:(AVSpeechSynthesizer *)synthesizer didCancelSpeechUtterance:(AVSpeechUtterance *)utterance;
- (void)speechSynthesizer:(AVSpeechSynthesizer *)synthesizer willSpeakRangeOfSpeechString:(NSRange)characterRange utterance:(AVSpeechUtterance *)utterance;
@end

NS_ASSUME_NONNULL_END

#endif
