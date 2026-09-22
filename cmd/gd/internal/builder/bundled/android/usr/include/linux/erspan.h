/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ERSPAN_H
#define _UAPI_ERSPAN_H
#include <linux/types.h>
#include <asm/byteorder.h>
struct erspan_md2 {
  __be32 timestamp;
  __be16 sgt;
#ifdef __LITTLE_ENDIAN_BITFIELD
  __u8 hwid_upper : 2, ft : 5, p : 1;
  __u8 o : 1, gra : 2, dir : 1, hwid : 4;
#elif defined(__BIG_ENDIAN_BITFIELD)
  __u8 p : 1, ft : 5, hwid_upper : 2;
  __u8 hwid : 4, dir : 1, gra : 2, o : 1;
#else
#error "Please fix <asm/byteorder.h>"
#endif
};
struct erspan_metadata {
  int version;
  union {
    __be32 index;
    struct erspan_md2 md2;
  } u;
};
#endif
