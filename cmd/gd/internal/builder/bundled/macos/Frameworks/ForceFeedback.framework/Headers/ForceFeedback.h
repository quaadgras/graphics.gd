// ForceFeedback declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what SDL makes
// use of. Not derived from Apple's SDK headers.
#ifndef GD_FORCEFEEDBACK_H
#define GD_FORCEFEEDBACK_H

#include <CoreFoundation/CoreFoundation.h>
#include <IOKit/IOTypes.h>

CF_EXTERN_C_BEGIN

typedef UInt32 FFCapabilitiesEffectType;
typedef UInt32 FFCapabilitiesEffectSubType;
typedef UInt32 FFCommandFlag;
typedef UInt32 FFCooperativeLevelFlag;
typedef UInt32 FFCoordinateSystemFlag;
typedef UInt32 FFEffectParameterFlag;
typedef UInt32 FFEffectStartFlag;
typedef UInt32 FFEffectStatusFlag;
typedef UInt32 FFProperty;
typedef UInt32 FFState;
// Windows types, which the API is modelled on.
typedef UInt32 DWORD, *LPDWORD, *LPDWORD_FF;
typedef SInt32 LONG, *LPLONG, *LPLONG_FF;
typedef SInt32 HRESULT;
typedef UInt32 IOByteCount;
typedef struct _FFDevice *FFDeviceObjectReference;
typedef struct _FFEffect *FFEffectObjectReference;

typedef struct {
	UInt32 dwSize;
	UInt32 dwAttackLevel;
	UInt32 dwAttackTime;
	UInt32 dwFadeLevel;
	UInt32 dwFadeTime;
} FFENVELOPE, *PFFENVELOPE;

typedef struct FFEFFECT {
	UInt32 dwSize;
	UInt32 dwFlags;
	UInt32 dwDuration;
	UInt32 dwSamplePeriod;
	UInt32 dwGain;
	UInt32 dwTriggerButton;
	UInt32 dwTriggerRepeatInterval;
	UInt32 cAxes;
	LPDWORD_FF rgdwAxes;
	LPLONG_FF rglDirection;
	PFFENVELOPE lpEnvelope;
	UInt32 cbTypeSpecificParams;
	void *lpvTypeSpecificParams;
	UInt32 dwStartDelay;
} FFEFFECT, *PFFEFFECT;

typedef struct {
	SInt32 lMagnitude;
} FFCONSTANTFORCE, *PFFCONSTANTFORCE;

typedef struct {
	SInt32 lStart;
	SInt32 lEnd;
} FFRAMPFORCE, *PFFRAMPFORCE;

typedef struct {
	UInt32 dwMagnitude;
	SInt32 lOffset;
	UInt32 dwPhase;
	UInt32 dwPeriod;
} FFPERIODIC, *PFFPERIODIC;

typedef struct {
	SInt32 lOffset;
	SInt32 lPositiveCoefficient;
	SInt32 lNegativeCoefficient;
	UInt32 dwPositiveSaturation;
	UInt32 dwNegativeSaturation;
	SInt32 lDeadBand;
} FFCONDITION, *PFFCONDITION;

typedef struct {
	UInt32 cChannels;
	UInt32 dwSamplePeriod;
	UInt32 cSamples;
	LPLONG_FF rglForceData;
} FFCUSTOMFORCE, *PFFCUSTOMFORCE;

typedef struct {
	UInt32 ffSpecVer[4]; // NumVersion, of the specification.
	FFCapabilitiesEffectType supportedEffects;
	FFCapabilitiesEffectType emulatedEffects;
	FFCapabilitiesEffectSubType subType;
	UInt32 numFfAxes;
	UInt8 ffAxes[32];
	UInt32 storageCapacity;
	UInt32 playbackCapacity;
	UInt32 firmwareVersion[4];
	UInt32 driverVersion[4];
} FFCAPABILITIES, *PFFCAPABILITIES;

CF_EXPORT HRESULT FFCreateDevice(io_service_t hidDevice, FFDeviceObjectReference *pDeviceReference);
CF_EXPORT HRESULT FFReleaseDevice(FFDeviceObjectReference deviceReference);
CF_EXPORT HRESULT FFIsForceFeedback(io_service_t hidDevice);
CF_EXPORT HRESULT FFDeviceCreateEffect(FFDeviceObjectReference deviceReference, CFUUIDRef uuidRef, FFEFFECT *pEffectDefinition, FFEffectObjectReference *pEffectReference);
CF_EXPORT HRESULT FFDeviceReleaseEffect(FFDeviceObjectReference deviceReference, FFEffectObjectReference effectReference);
CF_EXPORT HRESULT FFDeviceGetForceFeedbackCapabilities(FFDeviceObjectReference deviceReference, FFCAPABILITIES *pFFCapabilities);
CF_EXPORT HRESULT FFDeviceSendForceFeedbackCommand(FFDeviceObjectReference deviceReference, FFCommandFlag flags);
CF_EXPORT HRESULT FFDeviceSetForceFeedbackProperty(FFDeviceObjectReference deviceReference, FFProperty property, void *pValue);
CF_EXPORT HRESULT FFDeviceGetForceFeedbackProperty(FFDeviceObjectReference deviceReference, FFProperty property, void *pValue, IOByteCount valueSize);
CF_EXPORT HRESULT FFDeviceSetCooperativeLevel(FFDeviceObjectReference deviceReference, void *taskIdentifier, FFCooperativeLevelFlag flags);
CF_EXPORT HRESULT FFEffectStart(FFEffectObjectReference effectReference, UInt32 iterations, FFEffectStartFlag flags);
CF_EXPORT HRESULT FFEffectStop(FFEffectObjectReference effectReference);
CF_EXPORT HRESULT FFEffectGetEffectStatus(FFEffectObjectReference effectReference, FFEffectStatusFlag *pFlags);
CF_EXPORT HRESULT FFEffectSetParameters(FFEffectObjectReference effectReference, FFEFFECT *pFFEffect, FFEffectParameterFlag flags);
CF_EXPORT HRESULT FFEffectGetParameters(FFEffectObjectReference effectReference, FFEFFECT *pFFEffect, FFEffectParameterFlag flags);

CF_EXTERN_C_END

#endif
