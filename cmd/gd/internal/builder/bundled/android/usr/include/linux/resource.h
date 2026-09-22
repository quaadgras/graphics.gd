/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_RESOURCE_H
#define _UAPI_LINUX_RESOURCE_H
#include <linux/time_types.h>
#include <linux/types.h>
#define RUSAGE_SELF 0
#define RUSAGE_CHILDREN (- 1)
#define RUSAGE_BOTH (- 2)
#define RUSAGE_THREAD 1
struct rusage {
  struct timeval ru_utime;
  struct timeval ru_stime;
  __kernel_long_t ru_maxrss;
  __kernel_long_t ru_ixrss;
  __kernel_long_t ru_idrss;
  __kernel_long_t ru_isrss;
  __kernel_long_t ru_minflt;
  __kernel_long_t ru_majflt;
  __kernel_long_t ru_nswap;
  __kernel_long_t ru_inblock;
  __kernel_long_t ru_oublock;
  __kernel_long_t ru_msgsnd;
  __kernel_long_t ru_msgrcv;
  __kernel_long_t ru_nsignals;
  __kernel_long_t ru_nvcsw;
  __kernel_long_t ru_nivcsw;
};
struct rlimit {
  __kernel_ulong_t rlim_cur;
  __kernel_ulong_t rlim_max;
};
#define RLIM64_INFINITY (~0ULL)
struct rlimit64 {
  __u64 rlim_cur;
  __u64 rlim_max;
};
#define PRIO_MIN (- 20)
#define PRIO_MAX 20
#define PRIO_PROCESS 0
#define PRIO_PGRP 1
#define PRIO_USER 2
#define _STK_LIM (8 * 1024 * 1024)
#define MLOCK_LIMIT (8 * 1024 * 1024)
#include <asm/resource.h>
#endif
