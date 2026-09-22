// CoreMIDI declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_COREMIDI_H
#define GD_COREMIDI_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef UInt32 MIDIObjectRef;
typedef MIDIObjectRef MIDIClientRef;
typedef MIDIObjectRef MIDIPortRef;
typedef MIDIObjectRef MIDIDeviceRef;
typedef MIDIObjectRef MIDIEntityRef;
typedef MIDIObjectRef MIDIEndpointRef;
typedef UInt64 MIDITimeStamp;
typedef UInt32 MIDIUniqueID;
typedef unsigned long ItemCount;

// Packets are packed on 4 byte boundaries, as the data follows the header
// without padding.
#pragma pack(push, 4)
typedef struct MIDIPacket {
	MIDITimeStamp timeStamp;
	UInt16 length;
	Byte data[256];
} MIDIPacket;
typedef struct MIDIPacketList {
	UInt32 numPackets;
	MIDIPacket packet[1];
} MIDIPacketList;
#pragma pack(pop)

typedef struct MIDINotification {
	SInt32 messageID;
	UInt32 messageSize;
} MIDINotification;

typedef void (*MIDINotifyProc)(const MIDINotification *message, void *refCon);
typedef void (*MIDIReadProc)(const MIDIPacketList *pktlist, void *readProcRefCon, void *srcConnRefCon);

CF_EXPORT const CFStringRef kMIDIPropertyName;
CF_EXPORT const CFStringRef kMIDIPropertyDisplayName;
CF_EXPORT const CFStringRef kMIDIPropertyManufacturer;
CF_EXPORT const CFStringRef kMIDIPropertyModel;

CF_EXPORT OSStatus MIDIClientCreate(CFStringRef name, MIDINotifyProc notifyProc, void *notifyRefCon, MIDIClientRef *outClient);
CF_EXPORT OSStatus MIDIClientDispose(MIDIClientRef client);
CF_EXPORT OSStatus MIDIInputPortCreate(MIDIClientRef client, CFStringRef portName, MIDIReadProc readProc, void *refCon, MIDIPortRef *outPort);
CF_EXPORT OSStatus MIDIPortDispose(MIDIPortRef port);
CF_EXPORT OSStatus MIDIPortConnectSource(MIDIPortRef port, MIDIEndpointRef source, void *connRefCon);
CF_EXPORT OSStatus MIDIPortDisconnectSource(MIDIPortRef port, MIDIEndpointRef source);
CF_EXPORT ItemCount MIDIGetNumberOfSources(void);
CF_EXPORT MIDIEndpointRef MIDIGetSource(ItemCount sourceIndex0);
CF_EXPORT OSStatus MIDIObjectGetStringProperty(MIDIObjectRef obj, CFStringRef propertyID, CFStringRef *str);

CF_INLINE const MIDIPacket *MIDIPacketNext(const MIDIPacket *pkt) {
	// packets follow each other unaligned, the next begins after the data.
	return (const MIDIPacket *)(((const Byte *)pkt->data) + pkt->length);
}

CF_EXTERN_C_END

#endif
