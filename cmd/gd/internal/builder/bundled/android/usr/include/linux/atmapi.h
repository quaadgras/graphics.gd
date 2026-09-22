/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_ATMAPI_H
#define _LINUX_ATMAPI_H
#if defined(__sparc__) || defined(__ia64__)
#define __ATM_API_ALIGN __attribute__((aligned(8)))
#else
#define __ATM_API_ALIGN
#endif
typedef struct {
  unsigned char _[8];
} __ATM_API_ALIGN atm_kptr_t;
#endif
