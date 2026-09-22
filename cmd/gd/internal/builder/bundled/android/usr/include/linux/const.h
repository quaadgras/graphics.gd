/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_CONST_H
#define _UAPI_LINUX_CONST_H
#ifdef __ASSEMBLY__
#define _AC(X,Y) X
#define _AT(T,X) X
#else
#define __AC(X,Y) (X ##Y)
#define _AC(X,Y) __AC(X, Y)
#define _AT(T,X) ((T) (X))
#endif
#define _UL(x) (_AC(x, UL))
#define _ULL(x) (_AC(x, ULL))
#define _BITUL(x) (_UL(1) << (x))
#define _BITULL(x) (_ULL(1) << (x))
#ifndef __ASSEMBLY__
#define _BIT128(x) ((unsigned __int128) (1) << (x))
#endif
#define __ALIGN_KERNEL(x,a) __ALIGN_KERNEL_MASK(x, (__typeof__(x)) (a) - 1)
#define __ALIGN_KERNEL_MASK(x,mask) (((x) + (mask)) & ~(mask))
#define __KERNEL_DIV_ROUND_UP(n,d) (((n) + (d) - 1) / (d))
#endif
