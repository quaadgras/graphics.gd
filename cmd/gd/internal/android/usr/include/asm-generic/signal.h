/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__ASM_GENERIC_SIGNAL_H
#define _UAPI__ASM_GENERIC_SIGNAL_H
#include <linux/types.h>
#define _KERNEL__NSIG 64
#define _NSIG_BPW __BITS_PER_LONG
#define _NSIG_WORDS (_KERNEL__NSIG / _NSIG_BPW)
#define SIGHUP 1
#define SIGINT 2
#define SIGQUIT 3
#define SIGILL 4
#define SIGTRAP 5
#define SIGABRT 6
#define SIGIOT 6
#define SIGBUS 7
#define SIGFPE 8
#define SIGKILL 9
#define SIGUSR1 10
#define SIGSEGV 11
#define SIGUSR2 12
#define SIGPIPE 13
#define SIGALRM 14
#define SIGTERM 15
#define SIGSTKFLT 16
#define SIGCHLD 17
#define SIGCONT 18
#define SIGSTOP 19
#define SIGTSTP 20
#define SIGTTIN 21
#define SIGTTOU 22
#define SIGURG 23
#define SIGXCPU 24
#define SIGXFSZ 25
#define SIGVTALRM 26
#define SIGPROF 27
#define SIGWINCH 28
#define SIGIO 29
#define SIGPOLL SIGIO
#define SIGPWR 30
#define SIGSYS 31
#define SIGUNUSED 31
#define __SIGRTMIN 32
#ifndef __SIGRTMAX
#define __SIGRTMAX _KERNEL__NSIG
#endif
#if !defined(MINSIGSTKSZ) || !defined(SIGSTKSZ)
#define MINSIGSTKSZ 2048
#define SIGSTKSZ 8192
#endif
#ifndef __ASSEMBLY__
typedef struct {
  unsigned long sig[_NSIG_WORDS];
} sigset_t;
typedef unsigned long old_sigset_t;
#include <asm-generic/signal-defs.h>
#ifdef SA_RESTORER
#define __ARCH_HAS_SA_RESTORER
#endif
struct __kernel_sigaction {
  __sighandler_t sa_handler;
  unsigned long sa_flags;
#ifdef SA_RESTORER
  __sigrestore_t sa_restorer;
#endif
  sigset_t sa_mask;
};
typedef struct sigaltstack {
  void  * ss_sp;
  int ss_flags;
  __kernel_size_t ss_size;
} stack_t;
#endif
#endif
