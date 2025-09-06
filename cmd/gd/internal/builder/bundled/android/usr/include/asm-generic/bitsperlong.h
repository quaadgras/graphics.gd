/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__ASM_GENERIC_BITS_PER_LONG
#define _UAPI__ASM_GENERIC_BITS_PER_LONG
#ifndef __BITS_PER_LONG
#if defined(__CHAR_BIT__) && defined(__SIZEOF_LONG__)
#define __BITS_PER_LONG (__CHAR_BIT__ * __SIZEOF_LONG__)
#else
#define __BITS_PER_LONG 32
#endif
#endif
#ifndef __BITS_PER_LONG_LONG
#define __BITS_PER_LONG_LONG 64
#endif
#endif
