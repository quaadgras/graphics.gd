// Security declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot
// makes use of. Not derived from Apple's SDK headers.
#ifndef GD_SECURITY_H
#define GD_SECURITY_H

#include <CoreFoundation/CoreFoundation.h>

CF_EXTERN_C_BEGIN

typedef struct CF_BRIDGED_TYPE(id) __SecCertificate *SecCertificateRef;
typedef struct CF_BRIDGED_TYPE(id) __SecTrust *SecTrustRef;
typedef struct CF_BRIDGED_TYPE(id) __SecKey *SecKeyRef;
typedef struct CF_BRIDGED_TYPE(id) __SecPolicy *SecPolicyRef;

enum {
	errSecSuccess = 0,
	errSecUnimplemented = -4,
	errSecParam = -50,
	errSecAllocate = -108,
};

CF_EXPORT OSStatus SecTrustCopyAnchorCertificates(CFArrayRef *anchors);
CF_EXPORT SecCertificateRef SecCertificateCreateWithData(CFAllocatorRef allocator, CFDataRef data) CF_RETURNS_RETAINED;
CF_EXPORT CFDataRef SecCertificateCopyData(SecCertificateRef certificate) CF_RETURNS_RETAINED;
CF_EXPORT CFStringRef SecCertificateCopySubjectSummary(SecCertificateRef certificate) CF_RETURNS_RETAINED;

CF_EXTERN_C_END

#endif
