/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_KCMP_H
#define _UAPI_LINUX_KCMP_H
#include <linux/types.h>
enum kcmp_type {
  KCMP_FILE,
  KCMP_VM,
  KCMP_FILES,
  KCMP_FS,
  KCMP_SIGHAND,
  KCMP_IO,
  KCMP_SYSVSEM,
  KCMP_EPOLL_TFD,
  KCMP_TYPES,
};
struct kcmp_epoll_slot {
  __u32 efd;
  __u32 tfd;
  __u32 toff;
};
#endif
