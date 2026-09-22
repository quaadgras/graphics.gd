/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __LINUX_IF_ADDRLABEL_H
#define __LINUX_IF_ADDRLABEL_H
#include <linux/types.h>
struct ifaddrlblmsg {
  __u8 ifal_family;
  __u8 __ifal_reserved;
  __u8 ifal_prefixlen;
  __u8 ifal_flags;
  __u32 ifal_index;
  __u32 ifal_seq;
};
enum {
  IFAL_ADDRESS = 1,
  IFAL_LABEL = 2,
  __IFAL_MAX
};
#define IFAL_MAX (__IFAL_MAX - 1)
#endif
