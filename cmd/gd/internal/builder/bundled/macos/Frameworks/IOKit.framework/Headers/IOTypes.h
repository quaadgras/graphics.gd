// IOKit types and IOReturn for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_IOKIT_IOTYPES_H
#define GD_IOKIT_IOTYPES_H

#include <CoreFoundation/CoreFoundation.h>
#include <mach/mach.h>

CF_EXTERN_C_BEGIN

typedef UInt32 IOOptionBits;
typedef kern_return_t IOReturn;
typedef mach_port_t io_object_t;
typedef io_object_t io_iterator_t;
typedef io_object_t io_registry_entry_t;
typedef io_object_t io_service_t;
typedef io_object_t io_connect_t;
typedef char io_name_t[128];
typedef char io_string_t[512];

#define IO_OBJECT_NULL ((io_object_t)0)
CF_EXPORT const mach_port_t kIOMasterPortDefault;
CF_EXPORT const mach_port_t kIOMainPortDefault;

#define kIOServicePlane "IOService"
#define kIOPlatformSerialNumberKey "IOPlatformSerialNumber"
#define kIOPlatformUUIDKey "IOPlatformUUID"

// IOReturn values are Mach error codes of the IOKit system (0x38) and the
// sub-system 0: sys_iokit | sub_iokit_common | code.
#define iokit_common_err(code) ((IOReturn)(0xe0000000 | (code)))
#define kIOReturnSuccess 0
#define kIOReturnError iokit_common_err(0x2bc)
#define kIOReturnNoMemory iokit_common_err(0x2bd)
#define kIOReturnNoResources iokit_common_err(0x2be)
#define kIOReturnBadArgument iokit_common_err(0x2c2)
#define kIOReturnNotOpen iokit_common_err(0x2cd)
#define kIOReturnNoDevice iokit_common_err(0x2c0)
#define kIOReturnNotPermitted iokit_common_err(0x2e2)
#define kIOReturnExclusiveAccess iokit_common_err(0x2c5)
#define kIOReturnUnsupported iokit_common_err(0x2c7)
#define kIOReturnTimeout iokit_common_err(0x2d6)
#define kIOReturnInvalid iokit_common_err(0x1)

CF_EXTERN_C_END

#endif
