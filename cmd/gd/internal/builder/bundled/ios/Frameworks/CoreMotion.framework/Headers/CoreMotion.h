// CoreMotion declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREMOTION_H
#define GD_COREMOTION_H

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

typedef struct {
	double x;
	double y;
	double z;
} CMAcceleration;
typedef struct {
	double x;
	double y;
	double z;
} CMRotationRate;
typedef struct {
	double x;
	double y;
	double z;
} CMMagneticField;
typedef struct {
	double m11, m12, m13;
	double m21, m22, m23;
	double m31, m32, m33;
} CMRotationMatrix;
typedef struct {
	double x, y, z, w;
} CMQuaternion;

typedef NS_ENUM(int, CMMagneticFieldCalibrationAccuracy) {
	CMMagneticFieldCalibrationAccuracyUncalibrated = -1,
	CMMagneticFieldCalibrationAccuracyLow,
	CMMagneticFieldCalibrationAccuracyMedium,
	CMMagneticFieldCalibrationAccuracyHigh,
};
typedef struct {
	CMMagneticField field;
	CMMagneticFieldCalibrationAccuracy accuracy;
} CMCalibratedMagneticField;

typedef NS_OPTIONS(NSUInteger, CMAttitudeReferenceFrame) {
	CMAttitudeReferenceFrameXArbitraryZVertical = 1 << 0,
	CMAttitudeReferenceFrameXArbitraryCorrectedZVertical = 1 << 1,
	CMAttitudeReferenceFrameXMagneticNorthZVertical = 1 << 2,
	CMAttitudeReferenceFrameXTrueNorthZVertical = 1 << 3,
};

@interface CMLogItem : NSObject <NSSecureCoding, NSCopying>
@property(readonly, nonatomic) NSTimeInterval timestamp;
@end

@interface CMAccelerometerData : CMLogItem
@property(readonly, nonatomic) CMAcceleration acceleration;
@end

@interface CMGyroData : CMLogItem
@property(readonly, nonatomic) CMRotationRate rotationRate;
@end

@interface CMMagnetometerData : CMLogItem
@property(readonly, nonatomic) CMMagneticField magneticField;
@end

@interface CMAttitude : NSObject <NSCopying, NSSecureCoding>
@property(readonly, nonatomic) double roll;
@property(readonly, nonatomic) double pitch;
@property(readonly, nonatomic) double yaw;
@property(readonly, nonatomic) CMRotationMatrix rotationMatrix;
@property(readonly, nonatomic) CMQuaternion quaternion;
@end

@interface CMDeviceMotion : CMLogItem
@property(readonly, nonatomic) CMAttitude *attitude;
@property(readonly, nonatomic) CMRotationRate rotationRate;
@property(readonly, nonatomic) CMAcceleration gravity;
@property(readonly, nonatomic) CMAcceleration userAcceleration;
@property(readonly, nonatomic) CMCalibratedMagneticField magneticField;
@end

@interface CMMotionManager : NSObject
@property(assign, nonatomic) NSTimeInterval accelerometerUpdateInterval;
@property(readonly, nonatomic, getter=isAccelerometerAvailable) BOOL accelerometerAvailable;
@property(readonly, nonatomic, getter=isAccelerometerActive) BOOL accelerometerActive;
@property(readonly, nullable) CMAccelerometerData *accelerometerData;
- (void)startAccelerometerUpdates;
- (void)stopAccelerometerUpdates;
@property(assign, nonatomic) NSTimeInterval gyroUpdateInterval;
@property(readonly, nonatomic, getter=isGyroAvailable) BOOL gyroAvailable;
@property(readonly, nonatomic, getter=isGyroActive) BOOL gyroActive;
@property(readonly, nullable) CMGyroData *gyroData;
- (void)startGyroUpdates;
- (void)stopGyroUpdates;
@property(assign, nonatomic) NSTimeInterval magnetometerUpdateInterval;
@property(readonly, nonatomic, getter=isMagnetometerAvailable) BOOL magnetometerAvailable;
@property(readonly, nonatomic, getter=isMagnetometerActive) BOOL magnetometerActive;
@property(readonly, nullable) CMMagnetometerData *magnetometerData;
- (void)startMagnetometerUpdates;
- (void)stopMagnetometerUpdates;
@property(assign, nonatomic) NSTimeInterval deviceMotionUpdateInterval;
@property(readonly, nonatomic, getter=isDeviceMotionAvailable) BOOL deviceMotionAvailable;
@property(readonly, nonatomic, getter=isDeviceMotionActive) BOOL deviceMotionActive;
@property(readonly, nullable) CMDeviceMotion *deviceMotion;
+ (CMAttitudeReferenceFrame)availableAttitudeReferenceFrames;
- (void)startDeviceMotionUpdates;
- (void)startDeviceMotionUpdatesUsingReferenceFrame:(CMAttitudeReferenceFrame)referenceFrame;
- (void)stopDeviceMotionUpdates;
@end

NS_ASSUME_NONNULL_END

#endif
