/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_SWAB_H
#define _UAPI_LINUX_SWAB_H
#include <linux/types.h>
#include <linux/stddef.h>
#include <asm/bitsperlong.h>
#include <asm/swab.h>
#define ___constant_swab16(x) ((__u16) ((((__u16) (x) & (__u16) 0x00ffU) << 8) | (((__u16) (x) & (__u16) 0xff00U) >> 8)))
#define ___constant_swab32(x) ((__u32) ((((__u32) (x) & (__u32) 0x000000ffUL) << 24) | (((__u32) (x) & (__u32) 0x0000ff00UL) << 8) | (((__u32) (x) & (__u32) 0x00ff0000UL) >> 8) | (((__u32) (x) & (__u32) 0xff000000UL) >> 24)))
#define ___constant_swab64(x) ((__u64) ((((__u64) (x) & (__u64) 0x00000000000000ffULL) << 56) | (((__u64) (x) & (__u64) 0x000000000000ff00ULL) << 40) | (((__u64) (x) & (__u64) 0x0000000000ff0000ULL) << 24) | (((__u64) (x) & (__u64) 0x00000000ff000000ULL) << 8) | (((__u64) (x) & (__u64) 0x000000ff00000000ULL) >> 8) | (((__u64) (x) & (__u64) 0x0000ff0000000000ULL) >> 24) | (((__u64) (x) & (__u64) 0x00ff000000000000ULL) >> 40) | (((__u64) (x) & (__u64) 0xff00000000000000ULL) >> 56)))
#define ___constant_swahw32(x) ((__u32) ((((__u32) (x) & (__u32) 0x0000ffffUL) << 16) | (((__u32) (x) & (__u32) 0xffff0000UL) >> 16)))
#define ___constant_swahb32(x) ((__u32) ((((__u32) (x) & (__u32) 0x00ff00ffUL) << 8) | (((__u32) (x) & (__u32) 0xff00ff00UL) >> 8)))
#ifdef __SWAB_64_THRU_32__
#else
#endif
static inline __attribute__((__const__)) __u32 __fswahw32(__u32 val) {
  return ___constant_swahw32(val);
}
static inline __attribute__((__const__)) __u32 __fswahb32(__u32 val) {
  return ___constant_swahb32(val);
}
#define __swab16(x) (__u16) __builtin_bswap16((__u16) (x))
#define __swab32(x) (__u32) __builtin_bswap32((__u32) (x))
#define __swab64(x) (__u64) __builtin_bswap64((__u64) (x))
#if __BITS_PER_LONG == 64
#else
#endif
#define __swahw32(x) (__builtin_constant_p((__u32) (x)) ? ___constant_swahw32(x) : __fswahw32(x))
#define __swahb32(x) (__builtin_constant_p((__u32) (x)) ? ___constant_swahb32(x) : __fswahb32(x))
static __always_inline __u16 __swab16p(const __u16 * p) {
  return __swab16(* p);
}
static __always_inline __u32 __swab32p(const __u32 * p) {
  return __swab32(* p);
}
static __always_inline __u64 __swab64p(const __u64 * p) {
  return __swab64(* p);
}
static inline __u32 __swahw32p(const __u32 * p) {
  return __swahw32(* p);
}
static inline __u32 __swahb32p(const __u32 * p) {
  return __swahb32(* p);
}
static inline void __swab16s(__u16 * p) {
  * p = __swab16p(p);
}
static __always_inline void __swab32s(__u32 * p) {
  * p = __swab32p(p);
}
static __always_inline void __swab64s(__u64 * p) {
  * p = __swab64p(p);
}
static inline void __swahw32s(__u32 * p) {
  * p = __swahw32p(p);
}
static inline void __swahb32s(__u32 * p) {
  * p = __swahb32p(p);
}
#endif
