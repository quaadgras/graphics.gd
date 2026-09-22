// IOKitLib declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_IOKIT_IOKITLIB_H
#define GD_IOKIT_IOKITLIB_H

#include <IOKit/IOTypes.h>

CF_EXTERN_C_BEGIN

CF_EXPORT kern_return_t IOMasterPort(mach_port_t bootstrapPort, mach_port_t *masterPort);
CF_EXPORT kern_return_t IOMainPort(mach_port_t bootstrapPort, mach_port_t *mainPort);
CF_EXPORT kern_return_t IOObjectRelease(io_object_t object);
CF_EXPORT kern_return_t IOObjectRetain(io_object_t object);
CF_EXPORT boolean_t IOObjectIsEqualTo(io_object_t object, io_object_t anObject);
CF_EXPORT boolean_t IOObjectConformsTo(io_object_t object, const io_name_t className);
CF_EXPORT io_object_t IOIteratorNext(io_iterator_t iterator);
CF_EXPORT void IOIteratorReset(io_iterator_t iterator);
CF_EXPORT boolean_t IOIteratorIsValid(io_iterator_t iterator);
CF_EXPORT CFMutableDictionaryRef IOServiceMatching(const char *name) CF_RETURNS_RETAINED;
CF_EXPORT CFMutableDictionaryRef IOServiceNameMatching(const char *name) CF_RETURNS_RETAINED;
CF_EXPORT CFMutableDictionaryRef IORegistryEntryIDMatching(uint64_t entryID) CF_RETURNS_RETAINED;
CF_EXPORT io_service_t IOServiceGetMatchingService(mach_port_t mainPort, CFDictionaryRef matching CF_CONSUMED);
CF_EXPORT kern_return_t IOServiceGetMatchingServices(mach_port_t mainPort, CFDictionaryRef matching CF_CONSUMED, io_iterator_t *existing);
CF_EXPORT io_registry_entry_t IORegistryEntryFromPath(mach_port_t mainPort, const io_string_t path);
CF_EXPORT kern_return_t IORegistryEntryGetPath(io_registry_entry_t entry, const io_name_t plane, io_string_t path);
CF_EXPORT kern_return_t IORegistryEntryGetName(io_registry_entry_t entry, io_name_t name);
CF_EXPORT kern_return_t IORegistryEntryGetRegistryEntryID(io_registry_entry_t entry, uint64_t *entryID);
CF_EXPORT kern_return_t IORegistryEntryGetParentEntry(io_registry_entry_t entry, const io_name_t plane, io_registry_entry_t *parent);
CF_EXPORT kern_return_t IORegistryEntryGetChildIterator(io_registry_entry_t entry, const io_name_t plane, io_iterator_t *iterator);
CF_EXPORT CFTypeRef IORegistryEntryCreateCFProperty(io_registry_entry_t entry, CFStringRef key, CFAllocatorRef allocator, IOOptionBits options) CF_RETURNS_RETAINED;
CF_EXPORT kern_return_t IORegistryEntryCreateCFProperties(io_registry_entry_t entry, CFMutableDictionaryRef *properties, CFAllocatorRef allocator, IOOptionBits options);
CF_EXPORT kern_return_t IORegistryEntrySetCFProperty(io_registry_entry_t entry, CFStringRef propertyName, CFTypeRef property);

// Notifications
typedef struct IONotificationPort *IONotificationPortRef;
typedef void (*IOServiceMatchingCallback)(void *refcon, io_iterator_t iterator);
#define kIOPublishNotification "IOServicePublish"
#define kIOFirstPublishNotification "IOServiceFirstPublish"
#define kIOMatchedNotification "IOServiceMatched"
#define kIOFirstMatchNotification "IOServiceFirstMatch"
#define kIOTerminatedNotification "IOServiceTerminate"
CF_EXPORT IONotificationPortRef IONotificationPortCreate(mach_port_t mainPort);
CF_EXPORT void IONotificationPortDestroy(IONotificationPortRef notify);
CF_EXPORT CFRunLoopSourceRef IONotificationPortGetRunLoopSource(IONotificationPortRef notify);
CF_EXPORT mach_port_t IONotificationPortGetMachPort(IONotificationPortRef notify);
CF_EXPORT void IODispatchCalloutFromMessage(void *unused, mach_msg_header_t *msg, void *reference);
CF_EXPORT kern_return_t IOServiceAddMatchingNotification(IONotificationPortRef notifyPort, const io_name_t notificationType, CFDictionaryRef matching CF_CONSUMED, IOServiceMatchingCallback callback, void *refCon, io_iterator_t *notification);

CF_EXTERN_C_END

#endif
