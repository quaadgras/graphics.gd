// AudioToolbox declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_AUDIOTOOLBOX_H
#define GD_AUDIOTOOLBOX_H

#include <CoreAudioTypes/CoreAudioTypes.h>
#include <CoreGraphics/CoreGraphics.h> // as the umbrella brings in CoreFoundation and with it CoreGraphics.

CF_EXTERN_C_BEGIN

// AudioComponent

typedef struct OpaqueAudioComponent *AudioComponent;
typedef struct OpaqueAudioComponentInstance *AudioComponentInstance;

struct AudioComponentDescription {
	OSType componentType;
	OSType componentSubType;
	OSType componentManufacturer;
	UInt32 componentFlags;
	UInt32 componentFlagsMask;
};
typedef struct AudioComponentDescription AudioComponentDescription;

CF_EXPORT AudioComponent AudioComponentFindNext(AudioComponent inComponent, const AudioComponentDescription *inDesc);
CF_EXPORT UInt32 AudioComponentCount(const AudioComponentDescription *inDesc);
CF_EXPORT OSStatus AudioComponentInstanceNew(AudioComponent inComponent, AudioComponentInstance *outInstance);
CF_EXPORT OSStatus AudioComponentInstanceDispose(AudioComponentInstance inInstance);

// AudioUnit

typedef AudioComponentInstance AudioUnit;
typedef UInt32 AudioUnitPropertyID;
typedef UInt32 AudioUnitScope;
typedef UInt32 AudioUnitElement;
typedef UInt32 AudioUnitParameterID;
typedef Float32 AudioUnitParameterValue;

enum {
	kAudioUnitType_Output = 'auou',
	kAudioUnitType_MusicDevice = 'aumu',
	kAudioUnitType_MusicEffect = 'aumf',
	kAudioUnitType_FormatConverter = 'aufc',
	kAudioUnitType_Effect = 'aufx',
	kAudioUnitType_Mixer = 'aumx',
	kAudioUnitType_Generator = 'augn',
};
enum {
	kAudioUnitManufacturer_Apple = 'appl',
};
enum {
	kAudioUnitSubType_GenericOutput = 'genr',
	kAudioUnitSubType_VoiceProcessingIO = 'vpio',
#if TARGET_OS_IPHONE
	kAudioUnitSubType_RemoteIO = 'rioc',
#endif
	kAudioUnitSubType_HALOutput = 'ahal',
	kAudioUnitSubType_DefaultOutput = 'def ',
	kAudioUnitSubType_SystemOutput = 'sys ',
};
enum {
	kAudioUnitScope_Global = 0,
	kAudioUnitScope_Input = 1,
	kAudioUnitScope_Output = 2,
	kAudioUnitScope_Group = 3,
	kAudioUnitScope_Part = 4,
	kAudioUnitScope_Note = 5,
};
enum {
	kAudioUnitProperty_ClassInfo = 0,
	kAudioUnitProperty_MakeConnection = 1,
	kAudioUnitProperty_SampleRate = 2,
	kAudioUnitProperty_ParameterList = 3,
	kAudioUnitProperty_ParameterInfo = 4,
	kAudioUnitProperty_StreamFormat = 8,
	kAudioUnitProperty_ElementCount = 11,
	kAudioUnitProperty_Latency = 12,
	kAudioUnitProperty_MaximumFramesPerSlice = 14,
	kAudioUnitProperty_SetRenderCallback = 23,
	kAudioUnitProperty_ShouldAllocateBuffer = 51,
};
enum {
	kAudioOutputUnitProperty_CurrentDevice = 2000,
	kAudioOutputUnitProperty_IsRunning = 2001,
	kAudioOutputUnitProperty_ChannelMap = 2002,
	kAudioOutputUnitProperty_EnableIO = 2003,
	kAudioOutputUnitProperty_StartTime = 2004,
	kAudioOutputUnitProperty_SetInputCallback = 2005,
	kAudioOutputUnitProperty_HasIO = 2006,
};

typedef UInt32 AudioUnitRenderActionFlags;
enum {
	kAudioUnitRenderAction_PreRender = (1U << 2),
	kAudioUnitRenderAction_PostRender = (1U << 3),
	kAudioUnitRenderAction_OutputIsSilence = (1U << 4),
	kAudioOfflineUnitRenderAction_Preflight = (1U << 5),
	kAudioOfflineUnitRenderAction_Render = (1U << 6),
	kAudioOfflineUnitRenderAction_Complete = (1U << 7),
	kAudioUnitRenderAction_PostRenderError = (1U << 8),
	kAudioUnitRenderAction_DoNotCheckRenderArgs = (1U << 9),
};

typedef OSStatus (*AURenderCallback)(void *inRefCon, AudioUnitRenderActionFlags *ioActionFlags, const AudioTimeStamp *inTimeStamp, UInt32 inBusNumber, UInt32 inNumberFrames, AudioBufferList *ioData);

struct AURenderCallbackStruct {
	AURenderCallback inputProc;
	void *inputProcRefCon;
};
typedef struct AURenderCallbackStruct AURenderCallbackStruct;

CF_EXPORT OSStatus AudioUnitInitialize(AudioUnit inUnit);
CF_EXPORT OSStatus AudioUnitUninitialize(AudioUnit inUnit);
CF_EXPORT OSStatus AudioUnitGetPropertyInfo(AudioUnit inUnit, AudioUnitPropertyID inID, AudioUnitScope inScope, AudioUnitElement inElement, UInt32 *outDataSize, Boolean *outWritable);
CF_EXPORT OSStatus AudioUnitGetProperty(AudioUnit inUnit, AudioUnitPropertyID inID, AudioUnitScope inScope, AudioUnitElement inElement, void *outData, UInt32 *ioDataSize);
CF_EXPORT OSStatus AudioUnitSetProperty(AudioUnit inUnit, AudioUnitPropertyID inID, AudioUnitScope inScope, AudioUnitElement inElement, const void *inData, UInt32 inDataSize);
CF_EXPORT OSStatus AudioUnitRender(AudioUnit inUnit, AudioUnitRenderActionFlags *ioActionFlags, const AudioTimeStamp *inTimeStamp, UInt32 inOutputBusNumber, UInt32 inNumberFrames, AudioBufferList *ioData);
CF_EXPORT OSStatus AudioUnitReset(AudioUnit inUnit, AudioUnitScope inScope, AudioUnitElement inElement);
CF_EXPORT OSStatus AudioOutputUnitStart(AudioUnit ci);
CF_EXPORT OSStatus AudioOutputUnitStop(AudioUnit ci);

// AudioServices

typedef UInt32 SystemSoundID;
enum {
	kSystemSoundID_Vibrate = 0x00000FFF,
};
CF_EXPORT void AudioServicesPlaySystemSound(SystemSoundID inSystemSoundID);
CF_EXPORT void AudioServicesPlayAlertSound(SystemSoundID inSystemSoundID);
CF_EXPORT void AudioServicesPlaySystemSoundWithCompletion(SystemSoundID inSystemSoundID, void (^inCompletionBlock)(void));

CF_EXTERN_C_END

#endif
