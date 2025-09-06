/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_TIME_TYPES_H
#define _UAPI_LINUX_TIME_TYPES_H
#include <linux/time.h>
#include <linux/types.h>
struct __kernel_timespec {
  __kernel_time64_t tv_sec;
  long long tv_nsec;
};
struct __kernel_itimerspec {
  struct __kernel_timespec it_interval;
  struct __kernel_timespec it_value;
};
struct __kernel_old_timespec {
  __kernel_old_time_t tv_sec;
  long tv_nsec;
};
struct __kernel_sock_timeval {
  __s64 tv_sec;
  __s64 tv_usec;
};
#endif
