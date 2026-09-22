/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_EXFAT_H
#define _UAPI_LINUX_EXFAT_H
#include <linux/types.h>
#include <linux/ioctl.h>
#define EXFAT_IOC_SHUTDOWN _IOR('X', 125, __u32)
#define EXFAT_GOING_DOWN_DEFAULT 0x0
#define EXFAT_GOING_DOWN_FULLSYNC 0x1
#define EXFAT_GOING_DOWN_NOSYNC 0x2
#endif
