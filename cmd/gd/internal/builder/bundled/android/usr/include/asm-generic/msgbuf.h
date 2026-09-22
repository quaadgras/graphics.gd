/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_GENERIC_MSGBUF_H
#define __ASM_GENERIC_MSGBUF_H
#include <asm/bitsperlong.h>
#include <asm/ipcbuf.h>
struct msqid64_ds {
  struct ipc64_perm msg_perm;
#if __BITS_PER_LONG == 64
  long msg_stime;
  long msg_rtime;
  long msg_ctime;
#else
  unsigned long msg_stime;
  unsigned long msg_stime_high;
  unsigned long msg_rtime;
  unsigned long msg_rtime_high;
  unsigned long msg_ctime;
  unsigned long msg_ctime_high;
#endif
  unsigned long msg_cbytes;
  unsigned long msg_qnum;
  unsigned long msg_qbytes;
  __kernel_pid_t msg_lspid;
  __kernel_pid_t msg_lrpid;
  unsigned long __unused4;
  unsigned long __unused5;
};
#endif
