// GameController declarations that only macOS has, for graphics.gd's SDK.
//
// Written from the public documentation of the API. Not derived from
// Apple's SDK headers.
#ifndef GD_GCCONTROLLER_MACOS_H
#define GD_GCCONTROLLER_MACOS_H

#import <GameController/GameController.h>
#include <IOKit/hid/IOHIDLib.h>

NS_ASSUME_NONNULL_BEGIN

@interface GCController (macOS)
+ (BOOL)supportsHIDDevice:(IOHIDDeviceRef)device API_AVAILABLE(macos(11.0));
@end

NS_ASSUME_NONNULL_END

#endif
