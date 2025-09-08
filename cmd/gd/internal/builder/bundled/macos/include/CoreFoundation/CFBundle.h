#include <stddef.h>

typedef void *CFBundleRef;
typedef void *CFURLRef;
typedef void *CFStringRef;
typedef unsigned char Boolean;
typedef void *CFTypeRef;

const int kCFStringEncodingUTF8 = 0x08000100;

CFBundleRef CFBundleGetMainBundle();
CFURLRef CFBundleCopyResourceURL(CFBundleRef bundle, CFStringRef info, CFStringRef plist, void*);

#define CFSTR(cStr) ((CFStringRef)__builtin___CFStringMakeConstantString("" cStr ""))

CFStringRef CFURLGetString(CFURLRef url);
Boolean CFStringGetCString(CFStringRef theString, char *buffer, size_t bufferSize, int encoding);

void CFRelease(CFURLRef cf);

CFTypeRef CFBundleGetValueForInfoDictionaryKey(CFBundleRef bundle, CFStringRef key);