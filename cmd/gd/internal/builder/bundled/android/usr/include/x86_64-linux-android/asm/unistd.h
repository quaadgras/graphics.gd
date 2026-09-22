/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ASM_X86_UNISTD_H
#define _UAPI_ASM_X86_UNISTD_H
#define __X32_SYSCALL_BIT 0x40000000
#ifdef __i386__
#include <asm/unistd_32.h>
#elif defined(__ILP32__)
#include <asm/unistd_x32.h>
#else
#include <asm/unistd_64.h>
#endif
#endif
