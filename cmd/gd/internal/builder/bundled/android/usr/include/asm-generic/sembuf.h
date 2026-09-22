/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_GENERIC_SEMBUF_H
#define __ASM_GENERIC_SEMBUF_H
#include <asm/bitsperlong.h>
#include <asm/ipcbuf.h>
struct semid64_ds {
  struct ipc64_perm sem_perm;
#if __BITS_PER_LONG == 64
  long sem_otime;
  long sem_ctime;
#else
  unsigned long sem_otime;
  unsigned long sem_otime_high;
  unsigned long sem_ctime;
  unsigned long sem_ctime_high;
#endif
  unsigned long sem_nsems;
  unsigned long __unused3;
  unsigned long __unused4;
};
#endif
