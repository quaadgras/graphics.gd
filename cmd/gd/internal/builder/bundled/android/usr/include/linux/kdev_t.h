/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_KDEV_T_H
#define _UAPI_LINUX_KDEV_T_H
#define MAJOR(dev) ((dev) >> 8)
#define MINOR(dev) ((dev) & 0xff)
#define MKDEV(ma,mi) ((ma) << 8 | (mi))
#endif
