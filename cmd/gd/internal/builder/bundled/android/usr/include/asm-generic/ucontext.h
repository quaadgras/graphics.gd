/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __ASM_GENERIC_UCONTEXT_H
#define __ASM_GENERIC_UCONTEXT_H
struct ucontext {
  unsigned long uc_flags;
  struct ucontext * uc_link;
  stack_t uc_stack;
  struct sigcontext uc_mcontext;
  sigset_t uc_sigmask;
};
#endif
