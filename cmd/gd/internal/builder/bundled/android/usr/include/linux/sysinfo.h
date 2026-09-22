/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_SYSINFO_H
#define _LINUX_SYSINFO_H
#include <linux/types.h>
#define SI_LOAD_SHIFT 16
struct sysinfo {
  __kernel_long_t uptime;
  __kernel_ulong_t loads[3];
  __kernel_ulong_t totalram;
  __kernel_ulong_t freeram;
  __kernel_ulong_t sharedram;
  __kernel_ulong_t bufferram;
  __kernel_ulong_t totalswap;
  __kernel_ulong_t freeswap;
  __u16 procs;
  __u16 pad;
  __kernel_ulong_t totalhigh;
  __kernel_ulong_t freehigh;
  __u32 mem_unit;
  char _f[20 - 2 * sizeof(__kernel_ulong_t) - sizeof(__u32)];
};
#endif
