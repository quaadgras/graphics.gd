/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _ASM_X86_SIGINFO_H
#define _ASM_X86_SIGINFO_H
#ifdef __x86_64__
#ifdef __ILP32__
typedef long long __kernel_si_clock_t __attribute__((aligned(4)));
#define __ARCH_SI_CLOCK_T __kernel_si_clock_t
#define __ARCH_SI_ATTRIBUTES __attribute__((aligned(8)))
#endif
#endif
#include <asm-generic/siginfo.h>
#endif
