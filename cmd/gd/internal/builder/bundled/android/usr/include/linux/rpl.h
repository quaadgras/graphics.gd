/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_RPL_H
#define _UAPI_LINUX_RPL_H
#include <asm/byteorder.h>
#include <linux/types.h>
#include <linux/in6.h>
struct ipv6_rpl_sr_hdr {
  __u8 nexthdr;
  __u8 hdrlen;
  __u8 type;
  __u8 segments_left;
#ifdef __LITTLE_ENDIAN_BITFIELD
  __u32 cmpre : 4, cmpri : 4, reserved : 4, pad : 4, reserved1 : 16;
#elif defined(__BIG_ENDIAN_BITFIELD)
  __u32 cmpri : 4, cmpre : 4, pad : 4, reserved : 20;
#else
#error "Please fix <asm/byteorder.h>"
#endif
  union {
    __DECLARE_FLEX_ARRAY(struct in6_addr, addr);
    __DECLARE_FLEX_ARRAY(__u8, data);
  } segments;
} __attribute__((packed));
#define rpl_segaddr segments.addr
#define rpl_segdata segments.data
#endif
