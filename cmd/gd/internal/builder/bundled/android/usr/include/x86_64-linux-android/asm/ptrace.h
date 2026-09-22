/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_ASM_X86_PTRACE_H
#define _UAPI_ASM_X86_PTRACE_H
#include <linux/compiler.h>
#include <asm/ptrace-abi.h>
#include <asm/processor-flags.h>
#ifndef __ASSEMBLY__
#ifdef __i386__
struct pt_regs {
  long ebx;
  long ecx;
  long edx;
  long esi;
  long edi;
  long ebp;
  long eax;
  int xds;
  int xes;
  int xfs;
  int xgs;
  long orig_eax;
  long eip;
  int xcs;
  long eflags;
  long esp;
  int xss;
};
#else
struct pt_regs {
  unsigned long r15;
  unsigned long r14;
  unsigned long r13;
  unsigned long r12;
  unsigned long rbp;
  unsigned long rbx;
  unsigned long r11;
  unsigned long r10;
  unsigned long r9;
  unsigned long r8;
  unsigned long rax;
  unsigned long rcx;
  unsigned long rdx;
  unsigned long rsi;
  unsigned long rdi;
  unsigned long orig_rax;
  unsigned long rip;
  unsigned long cs;
  unsigned long eflags;
  unsigned long rsp;
  unsigned long ss;
};
#endif
#endif
#endif
