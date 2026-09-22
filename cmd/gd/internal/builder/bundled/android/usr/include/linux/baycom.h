/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _BAYCOM_H
#define _BAYCOM_H
struct baycom_debug_data {
  unsigned long debug1;
  unsigned long debug2;
  long debug3;
};
struct baycom_ioctl {
  int cmd;
  union {
    struct baycom_debug_data dbg;
  } data;
};
#define BAYCOMCTL_GETDEBUG 0x92
#endif
