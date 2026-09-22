/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_TIMES_H
#define _LINUX_TIMES_H
#include <linux/types.h>
struct tms {
  __kernel_clock_t tms_utime;
  __kernel_clock_t tms_stime;
  __kernel_clock_t tms_cutime;
  __kernel_clock_t tms_cstime;
};
#endif
