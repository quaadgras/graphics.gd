/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_IOAM6_H
#define _UAPI_LINUX_IOAM6_H
#include <asm/byteorder.h>
#include <linux/types.h>
#define IOAM6_U16_UNAVAILABLE U16_MAX
#define IOAM6_U32_UNAVAILABLE U32_MAX
#define IOAM6_U64_UNAVAILABLE U64_MAX
#define IOAM6_DEFAULT_ID (IOAM6_U32_UNAVAILABLE >> 8)
#define IOAM6_DEFAULT_ID_WIDE (IOAM6_U64_UNAVAILABLE >> 8)
#define IOAM6_DEFAULT_IF_ID IOAM6_U16_UNAVAILABLE
#define IOAM6_DEFAULT_IF_ID_WIDE IOAM6_U32_UNAVAILABLE
struct ioam6_hdr {
  __u8 opt_type;
  __u8 opt_len;
  __u8 : 8;
#define IOAM6_TYPE_PREALLOC 0
  __u8 type;
} __attribute__((packed));
struct ioam6_trace_hdr {
  __be16 namespace_id;
#ifdef __LITTLE_ENDIAN_BITFIELD
  __u8 : 1, : 1, overflow : 1, nodelen : 5;
  __u8 remlen : 7, : 1;
  union {
    __be32 type_be32;
    struct {
      __u32 bit7 : 1, bit6 : 1, bit5 : 1, bit4 : 1, bit3 : 1, bit2 : 1, bit1 : 1, bit0 : 1, bit15 : 1, bit14 : 1, bit13 : 1, bit12 : 1, bit11 : 1, bit10 : 1, bit9 : 1, bit8 : 1, bit23 : 1, bit22 : 1, bit21 : 1, bit20 : 1, bit19 : 1, bit18 : 1, bit17 : 1, bit16 : 1, : 8;
    } type;
  };
#elif defined(__BIG_ENDIAN_BITFIELD)
  __u8 nodelen : 5, overflow : 1, : 1, : 1;
  __u8 : 1, remlen : 7;
  union {
    __be32 type_be32;
    struct {
      __u32 bit0 : 1, bit1 : 1, bit2 : 1, bit3 : 1, bit4 : 1, bit5 : 1, bit6 : 1, bit7 : 1, bit8 : 1, bit9 : 1, bit10 : 1, bit11 : 1, bit12 : 1, bit13 : 1, bit14 : 1, bit15 : 1, bit16 : 1, bit17 : 1, bit18 : 1, bit19 : 1, bit20 : 1, bit21 : 1, bit22 : 1, bit23 : 1, : 8;
    } type;
  };
#else
#error "Please fix <asm/byteorder.h>"
#endif
#define IOAM6_TRACE_DATA_SIZE_MAX 244
  __u8 data[];
} __attribute__((packed));
#endif
