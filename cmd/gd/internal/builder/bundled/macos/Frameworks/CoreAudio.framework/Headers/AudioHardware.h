// CoreAudio (AudioHardware) declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_AUDIOHARDWARE_H
#define GD_AUDIOHARDWARE_H

#include <CoreAudioTypes/CoreAudioTypes.h>
#include <dispatch/dispatch.h>

CF_EXTERN_C_BEGIN

typedef UInt32 AudioObjectID;
typedef AudioObjectID AudioDeviceID;
typedef UInt32 AudioObjectPropertySelector;
typedef UInt32 AudioObjectPropertyScope;
typedef UInt32 AudioObjectPropertyElement;

struct AudioObjectPropertyAddress {
	AudioObjectPropertySelector mSelector;
	AudioObjectPropertyScope mScope;
	AudioObjectPropertyElement mElement;
};
typedef struct AudioObjectPropertyAddress AudioObjectPropertyAddress;

enum {
	kAudioObjectSystemObject = 1,
	kAudioObjectUnknown = 0,
};
enum {
	kAudioObjectPropertyScopeGlobal = 'glob',
	kAudioObjectPropertyScopeInput = 'inpt',
	kAudioObjectPropertyScopeOutput = 'outp',
	kAudioObjectPropertyScopePlayThrough = 'ptru',
	kAudioObjectPropertyElementMain = 0,
	kAudioObjectPropertyElementMaster = 0,
};
enum {
	kAudioObjectPropertyName = 'lnam',
	kAudioObjectPropertyManufacturer = 'lmak',
};
enum {
	kAudioHardwarePropertyDevices = 'dev#',
	kAudioHardwarePropertyDefaultInputDevice = 'dIn ',
	kAudioHardwarePropertyDefaultOutputDevice = 'dOut',
};
enum {
	kAudioDevicePropertyScopeInput = kAudioObjectPropertyScopeInput,
	kAudioDevicePropertyScopeOutput = kAudioObjectPropertyScopeOutput,
	kAudioDevicePropertyDeviceUID = 'uid ',
	kAudioDevicePropertyStreamConfiguration = 'slay',
	kAudioDevicePropertyNominalSampleRate = 'nsrt',
	kAudioDevicePropertyBufferFrameSize = 'fsiz',
	kAudioDevicePropertyDeviceIsAlive = 'livn',
};
enum {
	kAudioHardwareNoError = 0,
	kAudioHardwareNotRunningError = 'stop',
	kAudioHardwareUnspecifiedError = 'what',
	kAudioHardwareUnknownPropertyError = 'who?',
	kAudioHardwareBadPropertySizeError = '!siz',
	kAudioHardwareBadObjectError = '!obj',
	kAudioHardwareBadDeviceError = '!dev',
};

typedef OSStatus (*AudioObjectPropertyListenerProc)(AudioObjectID inObjectID, UInt32 inNumberAddresses, const AudioObjectPropertyAddress *inAddresses, void *inClientData);

CF_EXPORT Boolean AudioObjectHasProperty(AudioObjectID inObjectID, const AudioObjectPropertyAddress *inAddress);
CF_EXPORT OSStatus AudioObjectGetPropertyDataSize(AudioObjectID inObjectID, const AudioObjectPropertyAddress *inAddress, UInt32 inQualifierDataSize, const void *inQualifierData, UInt32 *outDataSize);
CF_EXPORT OSStatus AudioObjectGetPropertyData(AudioObjectID inObjectID, const AudioObjectPropertyAddress *inAddress, UInt32 inQualifierDataSize, const void *inQualifierData, UInt32 *ioDataSize, void *outData);
CF_EXPORT OSStatus AudioObjectSetPropertyData(AudioObjectID inObjectID, const AudioObjectPropertyAddress *inAddress, UInt32 inQualifierDataSize, const void *inQualifierData, UInt32 inDataSize, const void *inData);
CF_EXPORT OSStatus AudioObjectAddPropertyListener(AudioObjectID inObjectID, const AudioObjectPropertyAddress *inAddress, AudioObjectPropertyListenerProc inListener, void *inClientData);
CF_EXPORT OSStatus AudioObjectRemovePropertyListener(AudioObjectID inObjectID, const AudioObjectPropertyAddress *inAddress, AudioObjectPropertyListenerProc inListener, void *inClientData);

CF_EXTERN_C_END

#endif
