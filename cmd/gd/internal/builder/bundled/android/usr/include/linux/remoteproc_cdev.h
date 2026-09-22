/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_REMOTEPROC_CDEV_H_
#define _UAPI_REMOTEPROC_CDEV_H_
#include <linux/ioctl.h>
#include <linux/types.h>
#define RPROC_MAGIC 0xB7
#define RPROC_SET_SHUTDOWN_ON_RELEASE _IOW(RPROC_MAGIC, 1, __s32)
#define RPROC_GET_SHUTDOWN_ON_RELEASE _IOR(RPROC_MAGIC, 2, __s32)
#endif
