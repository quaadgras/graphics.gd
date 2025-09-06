/*
 * Copyright (C) 2008 The Android Open Source Project
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *  * Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in
 *    the documentation and/or other materials provided with the
 *    distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
 * OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
 * OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#pragma once

/**
 * @file sched.h
 * @brief Thread execution scheduling.
 */

#include <sys/cdefs.h>

#include <bits/timespec.h>
#include <linux/sched.h>
#include <linux/sched/types.h>

__BEGIN_DECLS

/*
 * @def SCHED_NORMAL
 * The standard (as opposed to real-time) round-robin scheduling policy.
 *
 * (Linux's name for POSIX's SCHED_OTHER.)
 *
 * See [sched(7)](https://man7.org/linux/man-pages/man7/sched.7.html)
 */

/*
 * @def SCHED_FIFO
 * The real-time first-in/first-out scheduling policy.
 *
 * See [sched(7)](https://man7.org/linux/man-pages/man7/sched.7.html)
 */

/*
 * @def SCHED_RR
 * The real-time round-robin policy. (See also SCHED_NORMAL/SCHED_OTHER.)
 *
 * See [sched(7)](https://man7.org/linux/man-pages/man7/sched.7.html)
 */

/*
 * @def SCHED_BATCH
 * The batch scheduling policy.
 *
 * See [sched(7)](https://man7.org/linux/man-pages/man7/sched.7.html)
 */

/*
 * @def SCHED_IDLE
 * The low priority "only when otherwise idle" scheduling priority.
 *
 * See [sched(7)](https://man7.org/linux/man-pages/man7/sched.7.html)
 */

/*
 * @def SCHED_DEADLINE
 * The deadline scheduling policy.
 *
 * See [sched(7)](https://man7.org/linux/man-pages/man7/sched.7.html)
 */

/*
 * The standard (as opposed to real-time) round-robin scheduling policy.
 *
 * (POSIX's name for Linux's SCHED_NORMAL.)
 */
#define SCHED_OTHER SCHED_NORMAL

/**
 * See sched_getparam()/sched_setparam() and
 * sched_getscheduler()/sched_setscheduler().
 */
struct sched_param {
  int sched_priority;
};

/**
 * [sched_setscheduler(2)](https://man7.org/linux/man-pages/man2/sched_setscheduler.2.html)
 * sets the scheduling policy and associated parameters for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_setscheduler(pid_t __pid, int __policy, const struct sched_param* _Nonnull __param);

/**
 * [sched_getscheduler(2)](https://man7.org/linux/man-pages/man2/sched_getscheduler.2)
 * gets the scheduling policy for the given thread.
 *
 * Returns a non-negative thread policy on success and returns -1 and sets
 * `errno` on failure.
 */
int sched_getscheduler(pid_t __pid);

/**
 * [sched_yield(2)](https://man7.org/linux/man-pages/man2/sched_yield.2.html)
 * voluntarily gives up using the CPU so that another thread can run.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_yield(void);

/**
 * [sched_get_priority_max(2)](https://man7.org/linux/man-pages/man2/sched_get_priority_max.2.html)
 * gets the maximum priority value allowed for the given scheduling policy.
 *
 * Returns a priority on success and returns -1 and sets `errno` on failure.
 */
int sched_get_priority_max(int __policy);

/**
 * [sched_get_priority_min(2)](https://man7.org/linux/man-pages/man2/sched_get_priority_min.2.html)
 * gets the minimum priority value allowed for the given scheduling policy.
 *
 * Returns a priority on success and returns -1 and sets `errno` on failure.
 */
int sched_get_priority_min(int __policy);

/**
 * [sched_setparam(2)](https://man7.org/linux/man-pages/man2/sched_setparam.2.html)
 * sets the scheduling parameters for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_setparam(pid_t __pid, const struct sched_param* _Nonnull __param);

/**
 * [sched_getparam(2)](https://man7.org/linux/man-pages/man2/sched_getparam.2.html)
 * gets the scheduling parameters for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_getparam(pid_t __pid, struct sched_param* _Nonnull __param);

/**
 * [sched_rr_get_interval(2)](https://man7.org/linux/man-pages/man2/sched_rr_get_interval.2.html)
 * queries the round-robin time quantum for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_rr_get_interval(pid_t __pid, struct timespec* _Nonnull __quantum);

#if defined(__USE_GNU)

/**
 * [clone(2)](https://man7.org/linux/man-pages/man2/clone.2.html)
 * creates a new child process.
 *
 * Returns the pid of the child to the caller on success and
 * returns -1 and sets `errno` on failure.
 */
int clone(int (* __BIONIC_COMPLICATED_NULLNESS __fn)(void* __BIONIC_COMPLICATED_NULLNESS ), void* __BIONIC_COMPLICATED_NULLNESS __child_stack, int __flags, void* _Nullable __arg, ...);

/**
 * [unshare(2)](https://man7.org/linux/man-pages/man2/unshare.2.html)
 * disassociates part of the caller's execution context.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int unshare(int __flags);

/**
 * [setns(2)](https://man7.org/linux/man-pages/man2/setns.2.html)
 * reassociates a thread with a different namespace.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int setns(int __fd, int __ns_type);

/**
 * [sched_getcpu(3)](https://man7.org/linux/man-pages/man3/sched_getcpu.3.html)
 * reports which CPU the caller is running on.
 *
 * Returns a non-negative CPU number on success and returns -1 and sets
 * `errno` on failure.
 */
int sched_getcpu(void);

#ifdef __LP64__
#define CPU_SETSIZE 1024
#else
#define CPU_SETSIZE 32
#endif

#define __CPU_BITTYPE  unsigned long int  /* mandated by the kernel  */
#define __CPU_BITS     (8 * sizeof(__CPU_BITTYPE))
#define __CPU_ELT(x)   ((x) / __CPU_BITS)
#define __CPU_MASK(x)  ((__CPU_BITTYPE)1 << ((x) & (__CPU_BITS - 1)))

/**
 * [cpu_set_t](https://man7.org/linux/man-pages/man3/CPU_SET.3.html) is a
 * statically-sized CPU set. See `CPU_ALLOC` for dynamically-sized CPU sets.
 */
typedef struct {
  __CPU_BITTYPE  __bits[ CPU_SETSIZE / __CPU_BITS ];
} cpu_set_t;

/**
 * [sched_setaffinity(2)](https://man7.org/linux/man-pages/man2/sched_setaffinity.2.html)
 * sets the CPU affinity mask for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_setaffinity(pid_t __pid, size_t __set_size, const cpu_set_t* _Nonnull __set);

/**
 * [sched_getaffinity(2)](https://man7.org/linux/man-pages/man2/sched_getaffinity.2.html)
 * gets the CPU affinity mask for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_getaffinity(pid_t __pid, size_t __set_size, cpu_set_t* _Nonnull __set);

/**
 * [sched_setattr(2)](https://man7.org/linux/man-pages/man2/sched_setattr.2.html)
 * sets the scheduling attributes for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_setattr(pid_t __pid, struct sched_attr* _Nonnull __attr, unsigned __flags) __INTRODUCED_IN(37);

/**
 * [sched_getattr(2)](https://man7.org/linux/man-pages/man2/sched_getattr.2.html)
 * gets the scheduling attributes for the given thread.
 *
 * Returns 0 on success and returns -1 and sets `errno` on failure.
 */
int sched_getattr(pid_t __pid, struct sched_attr* _Nonnull __attr, unsigned __size, unsigned __flags) __INTRODUCED_IN(37);

/**
 * [CPU_ZERO](https://man7.org/linux/man-pages/man3/CPU_ZERO.3.html) clears all
 * bits in a static CPU set.
 */
#define CPU_ZERO(set)          CPU_ZERO_S(sizeof(cpu_set_t), set)
/**
 * [CPU_ZERO_S](https://man7.org/linux/man-pages/man3/CPU_ZERO_S.3.html) clears all
 * bits in a dynamic CPU set allocated by `CPU_ALLOC`.
 */
#define CPU_ZERO_S(setsize, set)  __builtin_memset(set, 0, setsize)

/**
 * [CPU_SET](https://man7.org/linux/man-pages/man3/CPU_SET.3.html) sets one
 * bit in a static CPU set.
 */
#define CPU_SET(cpu, set)      CPU_SET_S(cpu, sizeof(cpu_set_t), set)
/**
 * [CPU_SET_S](https://man7.org/linux/man-pages/man3/CPU_SET_S.3.html) sets one
 * bit in a dynamic CPU set allocated by `CPU_ALLOC`.
 */
#define CPU_SET_S(cpu, setsize, set) \
  do { \
    size_t __cpu = (cpu); \
    if (__cpu < 8 * (setsize)) \
      (set)->__bits[__CPU_ELT(__cpu)] |= __CPU_MASK(__cpu); \
  } while (0)

/**
 * [CPU_CLR](https://man7.org/linux/man-pages/man3/CPU_CLR.3.html) clears one
 * bit in a static CPU set.
 */
#define CPU_CLR(cpu, set)      CPU_CLR_S(cpu, sizeof(cpu_set_t), set)
/**
 * [CPU_CLR_S](https://man7.org/linux/man-pages/man3/CPU_CLR_S.3.html) clears one
 * bit in a dynamic CPU set allocated by `CPU_ALLOC`.
 */
#define CPU_CLR_S(cpu, setsize, set) \
  do { \
    size_t __cpu = (cpu); \
    if (__cpu < 8 * (setsize)) \
      (set)->__bits[__CPU_ELT(__cpu)] &= ~__CPU_MASK(__cpu); \
  } while (0)

/**
 * [CPU_ISSET](https://man7.org/linux/man-pages/man3/CPU_ISSET.3.html) tests
 * whether the given bit is set in a static CPU set.
 */
#define CPU_ISSET(cpu, set)    CPU_ISSET_S(cpu, sizeof(cpu_set_t), set)
/**
 * [CPU_ISSET_S](https://man7.org/linux/man-pages/man3/CPU_ISSET_S.3.html) tests
 * whether the given bit is set in a dynamic CPU set allocated by `CPU_ALLOC`.
 */
#define CPU_ISSET_S(cpu, setsize, set) \
  (__extension__ ({ \
    size_t __cpu = (cpu); \
    (__cpu < 8 * (setsize)) \
      ? ((set)->__bits[__CPU_ELT(__cpu)] & __CPU_MASK(__cpu)) != 0 \
      : 0; \
  }))

/**
 * [CPU_COUNT](https://man7.org/linux/man-pages/man3/CPU_COUNT.3.html) counts
 * how many bits are set in a static CPU set.
 */
#define CPU_COUNT(set)         CPU_COUNT_S(sizeof(cpu_set_t), set)
/**
 * [CPU_COUNT_S](https://man7.org/linux/man-pages/man3/CPU_COUNT_S.3.html) counts
 * how many bits are set in a dynamic CPU set allocated by `CPU_ALLOC`.
 */
#define CPU_COUNT_S(setsize, set)  __sched_cpucount((setsize), (set))
int __sched_cpucount(size_t __set_size, const cpu_set_t* _Nonnull __set);

/**
 * [CPU_EQUAL](https://man7.org/linux/man-pages/man3/CPU_EQUAL.3.html) tests
 * whether two static CPU sets have the same bits set and cleared as each other.
 */
#define CPU_EQUAL(set1, set2)  CPU_EQUAL_S(sizeof(cpu_set_t), set1, set2)
/**
 * [CPU_EQUAL_S](https://man7.org/linux/man-pages/man3/CPU_EQUAL_S.3.html) tests
 * whether two dynamic CPU sets allocated by `CPU_ALLOC` have the same bits
 * set and cleared as each other.
 */
#define CPU_EQUAL_S(setsize, set1, set2)  (__builtin_memcmp(set1, set2, setsize) == 0)

/**
 * [CPU_AND](https://man7.org/linux/man-pages/man3/CPU_AND.3.html) ands two
 * static CPU sets.
 */
#define CPU_AND(dst, set1, set2)  __CPU_OP(dst, set1, set2, &)
/**
 * [CPU_AND_S](https://man7.org/linux/man-pages/man3/CPU_AND_S.3.html) ands two
 * dynamic CPU sets allocated by `CPU_ALLOC`.
 */
#define CPU_AND_S(setsize, dst, set1, set2)  __CPU_OP_S(setsize, dst, set1, set2, &)

/**
 * [CPU_OR](https://man7.org/linux/man-pages/man3/CPU_OR.3.html) ors two
 * static CPU sets.
 */
#define CPU_OR(dst, set1, set2)   __CPU_OP(dst, set1, set2, |)
/**
 * [CPU_OR_S](https://man7.org/linux/man-pages/man3/CPU_OR_S.3.html) ors two
 * dynamic CPU sets allocated by `CPU_ALLOC`.
 */
#define CPU_OR_S(setsize, dst, set1, set2)   __CPU_OP_S(setsize, dst, set1, set2, |)

/**
 * [CPU_XOR](https://man7.org/linux/man-pages/man3/CPU_XOR.3.html)
 * exclusive-ors two static CPU sets.
 */
#define CPU_XOR(dst, set1, set2)  __CPU_OP(dst, set1, set2, ^)
/**
 * [CPU_XOR_S](https://man7.org/linux/man-pages/man3/CPU_XOR_S.3.html)
 * exclusive-ors two dynamic CPU sets allocated by `CPU_ALLOC`.
 */
#define CPU_XOR_S(setsize, dst, set1, set2)  __CPU_OP_S(setsize, dst, set1, set2, ^)

#define __CPU_OP(dst, set1, set2, op)  __CPU_OP_S(sizeof(cpu_set_t), dst, set1, set2, op)

#define __CPU_OP_S(setsize, dstset, srcset1, srcset2, op) \
  do { \
    cpu_set_t* __dst = (dstset); \
    const __CPU_BITTYPE* __src1 = (srcset1)->__bits; \
    const __CPU_BITTYPE* __src2 = (srcset2)->__bits; \
    size_t __nn = 0, __nn_max = (setsize)/sizeof(__CPU_BITTYPE); \
    for (; __nn < __nn_max; __nn++) \
      (__dst)->__bits[__nn] = __src1[__nn] op __src2[__nn]; \
  } while (0)

/**
 * [CPU_ALLOC_SIZE](https://man7.org/linux/man-pages/man3/CPU_ALLOC_SIZE.3.html)
 * returns the size of a CPU set large enough for CPUs in the range 0..count-1.
 */
#define CPU_ALLOC_SIZE(count) \
  __CPU_ELT((count) + (__CPU_BITS - 1)) * sizeof(__CPU_BITTYPE)

/**
 * [CPU_ALLOC](https://man7.org/linux/man-pages/man3/CPU_ALLOC.3.html)
 * allocates a CPU set large enough for CPUs in the range 0..count-1.
 */
#define CPU_ALLOC(count)  __sched_cpualloc((count))
cpu_set_t* _Nullable __sched_cpualloc(size_t __count);

/**
 * [CPU_FREE](https://man7.org/linux/man-pages/man3/CPU_FREE.3.html)
 * deallocates a CPU set allocated by `CPU_ALLOC`.
 */
#define CPU_FREE(set)     __sched_cpufree((set))
void __sched_cpufree(cpu_set_t* _Nonnull __set);

#endif /* __USE_GNU */

__END_DECLS
