// AVFoundation declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_AVFOUNDATION_H
#define GD_AVFOUNDATION_H

#import <AVFAudio/AVFAudio.h>
#import <CoreMedia/CoreMedia.h>
#import <Foundation/Foundation.h>
#include <dispatch/dispatch.h>

NS_ASSUME_NONNULL_BEGIN

typedef NSString *AVMediaType NS_EXTENSIBLE_STRING_ENUM;
FOUNDATION_EXTERN AVMediaType const AVMediaTypeVideo;
FOUNDATION_EXTERN AVMediaType const AVMediaTypeAudio;

typedef NS_ENUM(NSInteger, AVAuthorizationStatus) {
	AVAuthorizationStatusNotDetermined = 0,
	AVAuthorizationStatusRestricted = 1,
	AVAuthorizationStatusDenied = 2,
	AVAuthorizationStatusAuthorized = 3,
};
typedef NS_ENUM(NSInteger, AVCaptureDevicePosition) {
	AVCaptureDevicePositionUnspecified = 0,
	AVCaptureDevicePositionBack = 1,
	AVCaptureDevicePositionFront = 2,
};
typedef NS_ENUM(NSInteger, AVCaptureFocusMode) {
	AVCaptureFocusModeLocked = 0,
	AVCaptureFocusModeAutoFocus = 1,
	AVCaptureFocusModeContinuousAutoFocus = 2,
};
typedef NS_ENUM(NSInteger, AVCaptureExposureMode) {
	AVCaptureExposureModeLocked = 0,
	AVCaptureExposureModeAutoExpose = 1,
	AVCaptureExposureModeContinuousAutoExposure = 2,
	AVCaptureExposureModeCustom = 3,
};
typedef NS_ENUM(NSInteger, AVCaptureWhiteBalanceMode) {
	AVCaptureWhiteBalanceModeLocked = 0,
	AVCaptureWhiteBalanceModeAutoWhiteBalance = 1,
	AVCaptureWhiteBalanceModeContinuousAutoWhiteBalance = 2,
};
typedef NS_ENUM(NSInteger, AVCaptureVideoOrientation) {
	AVCaptureVideoOrientationPortrait = 1,
	AVCaptureVideoOrientationPortraitUpsideDown = 2,
	AVCaptureVideoOrientationLandscapeRight = 3,
	AVCaptureVideoOrientationLandscapeLeft = 4,
};

typedef NSString *AVCaptureDeviceType NS_TYPED_ENUM;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeExternal API_AVAILABLE(ios(17.0));
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeExternalUnknown API_AVAILABLE(macos(10.15));
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeContinuityCamera API_AVAILABLE(ios(17.0));
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInWideAngleCamera;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInTelephotoCamera;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInUltraWideCamera;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInDualCamera;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInDualWideCamera;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInTripleCamera;
FOUNDATION_EXTERN AVCaptureDeviceType const AVCaptureDeviceTypeBuiltInTrueDepthCamera;

typedef NSString *AVCaptureSessionPreset NS_TYPED_EXTENSIBLE_ENUM;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPresetPhoto;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPresetHigh;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPresetMedium;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPresetLow;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPreset640x480;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPreset1280x720;
FOUNDATION_EXTERN AVCaptureSessionPreset const AVCaptureSessionPreset1920x1080;

FOUNDATION_EXTERN NSNotificationName const AVCaptureDeviceWasConnectedNotification;
FOUNDATION_EXTERN NSNotificationName const AVCaptureDeviceWasDisconnectedNotification;

@class AVCaptureOutput, AVCaptureConnection;

@interface AVCaptureDevice : NSObject
+ (NSArray<AVCaptureDevice *> *)devices;
+ (NSArray<AVCaptureDevice *> *)devicesWithMediaType:(AVMediaType)mediaType;
+ (nullable AVCaptureDevice *)defaultDeviceWithMediaType:(AVMediaType)mediaType;
+ (nullable AVCaptureDevice *)deviceWithUniqueID:(NSString *)deviceUniqueID;
+ (AVAuthorizationStatus)authorizationStatusForMediaType:(AVMediaType)mediaType;
+ (void)requestAccessForMediaType:(AVMediaType)mediaType completionHandler:(void (^)(BOOL granted))handler;
@property(nonatomic, readonly) NSString *uniqueID;
@property(nonatomic, readonly) NSString *modelID;
@property(nonatomic, readonly) NSString *localizedName;
@property(nonatomic, readonly) AVCaptureDevicePosition position;
@property(nonatomic, readonly) AVCaptureDeviceType deviceType;
@property(nonatomic, readonly, getter=isConnected) BOOL connected;
- (BOOL)hasMediaType:(AVMediaType)mediaType;
- (BOOL)lockForConfiguration:(NSError **)outError;
- (void)unlockForConfiguration;
- (BOOL)isFocusModeSupported:(AVCaptureFocusMode)focusMode;
@property(nonatomic) AVCaptureFocusMode focusMode;
- (BOOL)isExposureModeSupported:(AVCaptureExposureMode)exposureMode;
@property(nonatomic) AVCaptureExposureMode exposureMode;
- (BOOL)isWhiteBalanceModeSupported:(AVCaptureWhiteBalanceMode)whiteBalanceMode;
@property(nonatomic) AVCaptureWhiteBalanceMode whiteBalanceMode;
@end

@interface AVCaptureDeviceDiscoverySession : NSObject
+ (instancetype)discoverySessionWithDeviceTypes:(NSArray<AVCaptureDeviceType> *)deviceTypes mediaType:(nullable AVMediaType)mediaType position:(AVCaptureDevicePosition)position;
@property(nonatomic, readonly) NSArray<AVCaptureDevice *> *devices;
@end

@interface AVCaptureInput : NSObject
@end

@interface AVCaptureDeviceInput : AVCaptureInput
+ (nullable instancetype)deviceInputWithDevice:(AVCaptureDevice *)device error:(NSError **)outError;
- (nullable instancetype)initWithDevice:(AVCaptureDevice *)device error:(NSError **)outError;
@property(nonatomic, readonly) AVCaptureDevice *device;
@end

@interface AVCaptureConnection : NSObject
@property(nonatomic, getter=isEnabled) BOOL enabled;
@property(nonatomic, readonly, getter=isVideoOrientationSupported) BOOL supportsVideoOrientation;
@property(nonatomic) AVCaptureVideoOrientation videoOrientation;
@property(nonatomic, getter=isVideoMirrored) BOOL videoMirrored;
@end

@interface AVCaptureOutput : NSObject
@property(nonatomic, readonly) NSArray<AVCaptureConnection *> *connections;
- (nullable AVCaptureConnection *)connectionWithMediaType:(AVMediaType)mediaType;
@end

@protocol AVCaptureVideoDataOutputSampleBufferDelegate <NSObject>
@optional
- (void)captureOutput:(AVCaptureOutput *)output didOutputSampleBuffer:(CMSampleBufferRef)sampleBuffer fromConnection:(AVCaptureConnection *)connection;
- (void)captureOutput:(AVCaptureOutput *)output didDropSampleBuffer:(CMSampleBufferRef)sampleBuffer fromConnection:(AVCaptureConnection *)connection;
@end

@interface AVCaptureVideoDataOutput : AVCaptureOutput
- (void)setSampleBufferDelegate:(nullable id<AVCaptureVideoDataOutputSampleBufferDelegate>)sampleBufferDelegate queue:(nullable dispatch_queue_t)sampleBufferCallbackQueue;
@property(nonatomic, readonly, nullable) id<AVCaptureVideoDataOutputSampleBufferDelegate> sampleBufferDelegate;
@property(nonatomic, copy, null_resettable) NSDictionary<NSString *, id> *videoSettings;
@property(nonatomic, readonly) NSArray<NSNumber *> *availableVideoCVPixelFormatTypes;
@property(nonatomic) BOOL alwaysDiscardsLateVideoFrames;
@end

@interface AVCaptureSession : NSObject
- (BOOL)canSetSessionPreset:(AVCaptureSessionPreset)preset;
@property(nonatomic, copy) AVCaptureSessionPreset sessionPreset;
@property(nonatomic, readonly) NSArray<__kindof AVCaptureInput *> *inputs;
@property(nonatomic, readonly) NSArray<__kindof AVCaptureOutput *> *outputs;
- (BOOL)canAddInput:(AVCaptureInput *)input;
- (void)addInput:(AVCaptureInput *)input;
- (void)removeInput:(AVCaptureInput *)input;
- (BOOL)canAddOutput:(AVCaptureOutput *)output;
- (void)addOutput:(AVCaptureOutput *)output;
- (void)removeOutput:(AVCaptureOutput *)output;
- (void)beginConfiguration;
- (void)commitConfiguration;
@property(nonatomic, readonly, getter=isRunning) BOOL running;
- (void)startRunning;
- (void)stopRunning;
@end

NS_ASSUME_NONNULL_END

#endif
