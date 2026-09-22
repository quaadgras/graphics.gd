/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _ASM_X86_SEMBUF_H
#define _ASM_X86_SEMBUF_H
#include <asm/ipcbuf.h>
struct semid64_ds {
  struct ipc64_perm sem_perm;
#ifdef __i386__
  unsigned long sem_otime;
  unsigned long sem_otime_high;
  unsigned long sem_ctime;
  unsigned long sem_ctime_high;
#else
  __kernel_long_t sem_otime;
  __kernel_ulong_t __unused1;
  __kernel_long_t sem_ctime;
  __kernel_ulong_t __unused2;
#endif
  __kernel_ulong_t sem_nsems;
  __kernel_ulong_t __unused3;
  __kernel_ulong_t __unused4;
};
#endif
