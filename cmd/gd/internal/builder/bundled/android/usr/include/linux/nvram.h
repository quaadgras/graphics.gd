/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_NVRAM_H
#define _UAPI_LINUX_NVRAM_H
#include <linux/ioctl.h>
#define NVRAM_INIT _IO('p', 0x40)
#define NVRAM_SETCKS _IO('p', 0x41)
#define NVRAM_FIRST_BYTE 14
#define NVRAM_OFFSET(x) ((x) - NVRAM_FIRST_BYTE)
#endif
