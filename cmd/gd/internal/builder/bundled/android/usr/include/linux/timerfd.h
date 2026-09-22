/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_TIMERFD_H
#define _UAPI_LINUX_TIMERFD_H
#include <linux/types.h>
#include <linux/fcntl.h>
#include <linux/ioctl.h>
#define TFD_TIMER_ABSTIME (1 << 0)
#define TFD_TIMER_CANCEL_ON_SET (1 << 1)
#define TFD_CLOEXEC O_CLOEXEC
#define TFD_NONBLOCK O_NONBLOCK
#define TFD_IOC_SET_TICKS _IOW('T', 0, __u64)
#endif
