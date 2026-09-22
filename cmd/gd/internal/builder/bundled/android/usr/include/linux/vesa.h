/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VESA_H
#define _UAPI_LINUX_VESA_H
enum vesa_blank_mode {
  VESA_NO_BLANKING = 0,
#define VESA_NO_BLANKING VESA_NO_BLANKING
  VESA_VSYNC_SUSPEND = 1,
#define VESA_VSYNC_SUSPEND VESA_VSYNC_SUSPEND
  VESA_HSYNC_SUSPEND = 2,
#define VESA_HSYNC_SUSPEND VESA_HSYNC_SUSPEND
  VESA_POWERDOWN = VESA_VSYNC_SUSPEND | VESA_HSYNC_SUSPEND,
#define VESA_POWERDOWN VESA_POWERDOWN
  VESA_BLANK_MAX = VESA_POWERDOWN,
};
#endif
