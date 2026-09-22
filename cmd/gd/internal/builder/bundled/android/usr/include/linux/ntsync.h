/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __LINUX_NTSYNC_H
#define __LINUX_NTSYNC_H
#include <linux/types.h>
struct ntsync_sem_args {
  __u32 sem;
  __u32 count;
  __u32 max;
};
#define NTSYNC_IOC_CREATE_SEM _IOWR('N', 0x80, struct ntsync_sem_args)
#define NTSYNC_IOC_SEM_POST _IOWR('N', 0x81, __u32)
#endif
