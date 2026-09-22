// CoreAudio type declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREAUDIOTYPES_H
#define GD_COREAUDIOTYPES_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef UInt32 AudioFormatID;
typedef UInt32 AudioFormatFlags;
typedef UInt32 AudioChannelLabel;
typedef UInt32 AudioChannelLayoutTag;

enum {
	kAudio_UnimplementedError = -4,
	kAudio_FileNotFoundError = -43,
	kAudio_ParamError = -50,
	kAudio_MemFullError = -108,
};

enum {
	kAudioFormatLinearPCM = 'lpcm',
	kAudioFormatAC3 = 'ac-3',
	kAudioFormatAppleIMA4 = 'ima4',
	kAudioFormatMPEG4AAC = 'aac ',
	kAudioFormatMPEGLayer3 = '.mp3',
	kAudioFormatAppleLossless = 'alac',
	kAudioFormatOpus = 'opus',
};

enum {
	kAudioFormatFlagIsFloat = (1U << 0),
	kAudioFormatFlagIsBigEndian = (1U << 1),
	kAudioFormatFlagIsSignedInteger = (1U << 2),
	kAudioFormatFlagIsPacked = (1U << 3),
	kAudioFormatFlagIsAlignedHigh = (1U << 4),
	kAudioFormatFlagIsNonInterleaved = (1U << 5),
	kAudioFormatFlagIsNonMixable = (1U << 6),
	kAudioFormatFlagsAreAllClear = 0x80000000,
	kLinearPCMFormatFlagIsFloat = kAudioFormatFlagIsFloat,
	kLinearPCMFormatFlagIsBigEndian = kAudioFormatFlagIsBigEndian,
	kLinearPCMFormatFlagIsSignedInteger = kAudioFormatFlagIsSignedInteger,
	kLinearPCMFormatFlagIsPacked = kAudioFormatFlagIsPacked,
	kLinearPCMFormatFlagIsAlignedHigh = kAudioFormatFlagIsAlignedHigh,
	kLinearPCMFormatFlagIsNonInterleaved = kAudioFormatFlagIsNonInterleaved,
	kLinearPCMFormatFlagIsNonMixable = kAudioFormatFlagIsNonMixable,
	kLinearPCMFormatFlagsAreAllClear = kAudioFormatFlagsAreAllClear,
	kAudioFormatFlagsNativeEndian = 0,
	kAudioFormatFlagsNativeFloatPacked = kAudioFormatFlagIsFloat | kAudioFormatFlagsNativeEndian | kAudioFormatFlagIsPacked,
};

struct AudioStreamBasicDescription {
	Float64 mSampleRate;
	AudioFormatID mFormatID;
	AudioFormatFlags mFormatFlags;
	UInt32 mBytesPerPacket;
	UInt32 mFramesPerPacket;
	UInt32 mBytesPerFrame;
	UInt32 mChannelsPerFrame;
	UInt32 mBitsPerChannel;
	UInt32 mReserved;
};
typedef struct AudioStreamBasicDescription AudioStreamBasicDescription;

struct AudioBuffer {
	UInt32 mNumberChannels;
	UInt32 mDataByteSize;
	void *mData;
};
typedef struct AudioBuffer AudioBuffer;

struct AudioBufferList {
	UInt32 mNumberBuffers;
	AudioBuffer mBuffers[1]; // variable length, mNumberBuffers elements.
};
typedef struct AudioBufferList AudioBufferList;

struct SMPTETime {
	SInt16 mSubframes;
	SInt16 mSubframeDivisor;
	UInt32 mCounter;
	UInt32 mType;
	UInt32 mFlags;
	SInt16 mHours;
	SInt16 mMinutes;
	SInt16 mSeconds;
	SInt16 mFrames;
};
typedef struct SMPTETime SMPTETime;

struct AudioTimeStamp {
	Float64 mSampleTime;
	UInt64 mHostTime;
	Float64 mRateScalar;
	UInt64 mWordClockTime;
	SMPTETime mSMPTETime;
	UInt32 mFlags;
	UInt32 mReserved;
};
typedef struct AudioTimeStamp AudioTimeStamp;

enum {
	kAudioTimeStampNothingValid = 0,
	kAudioTimeStampSampleTimeValid = (1U << 0),
	kAudioTimeStampHostTimeValid = (1U << 1),
	kAudioTimeStampRateScalarValid = (1U << 2),
	kAudioTimeStampWordClockTimeValid = (1U << 3),
	kAudioTimeStampSMPTETimeValid = (1U << 4),
};

CF_EXTERN_C_END

#endif
