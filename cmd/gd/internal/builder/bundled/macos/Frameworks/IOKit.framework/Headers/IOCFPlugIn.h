// IOCFPlugIn declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_IOKIT_IOCFPLUGIN_H
#define GD_IOKIT_IOCFPLUGIN_H

#include <CoreFoundation/CoreFoundation.h>
#include <IOKit/IOKitLib.h>

CF_EXTERN_C_BEGIN

// COM style interfaces, as in CFPlugInCOM.h.
typedef SInt32 HRESULT;
typedef UInt32 ULONG;
typedef void *LPVOID;
typedef CFUUIDRef REFIID;
#define S_OK ((HRESULT)0x00000000L)
// CFPlugInCOM's own values, not those of Windows' COM.
#define E_UNEXPECTED ((HRESULT)0x8000FFFFL)
#define E_NOTIMPL ((HRESULT)0x80000001L)
#define E_OUTOFMEMORY ((HRESULT)0x80000002L)
#define E_INVALIDARG ((HRESULT)0x80000003L)
#define E_NOINTERFACE ((HRESULT)0x80000004L)
#define E_POINTER ((HRESULT)0x80000005L)
#define E_HANDLE ((HRESULT)0x80000006L)
#define E_ABORT ((HRESULT)0x80000007L)
#define E_FAIL ((HRESULT)0x80000008L)
#define E_ACCESSDENIED ((HRESULT)0x80000009L)
#define REGDB_E_CLASSNOTREG ((HRESULT)0x80040154L)
typedef struct IUnknownVTbl {
	void *_reserved;
	HRESULT (*QueryInterface)(void *thisPointer, REFIID iid, LPVOID *ppv);
	ULONG (*AddRef)(void *thisPointer);
	ULONG (*Release)(void *thisPointer);
} IUnknownVTbl;

#define IUNKNOWN_C_GUTS \
	void *_reserved;    \
	HRESULT (*QueryInterface)(void *thisPointer, REFIID iid, LPVOID *ppv); \
	ULONG (*AddRef)(void *thisPointer); \
	ULONG (*Release)(void *thisPointer)

typedef struct IOCFPlugInInterfaceStruct {
	IUNKNOWN_C_GUTS;
	UInt16 version;
	UInt16 revision;
	IOReturn (*Probe)(void *thisPointer, CFDictionaryRef propertyTable, io_service_t service, SInt32 *order);
	IOReturn (*Start)(void *thisPointer, CFDictionaryRef propertyTable, io_service_t service);
	IOReturn (*Stop)(void *thisPointer);
} IOCFPlugInInterface;

CF_EXPORT kern_return_t IOCreatePlugInInterfaceForService(io_service_t service, CFUUIDRef pluginType, CFUUIDRef interfaceType, IOCFPlugInInterface ***theInterface, SInt32 *theScore);
CF_EXPORT kern_return_t IODestroyPlugInInterface(IOCFPlugInInterface **interface);

CF_EXTERN_C_END

#endif
