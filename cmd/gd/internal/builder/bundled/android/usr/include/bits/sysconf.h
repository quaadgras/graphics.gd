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

#include <sys/cdefs.h>

/** sysconf() query for the maximum number of bytes of exec() arguments. */
#define _SC_ARG_MAX 0x0000
/** sysconf() query for bc(1) behavior equivalent to _POSIX2_BC_BASE_MAX. */
#define _SC_BC_BASE_MAX 0x0001
/** sysconf() query for bc(1) behavior equivalent to _POSIX2_BC_DIM_MAX. */
#define _SC_BC_DIM_MAX 0x0002
/** sysconf() query for bc(1) behavior equivalent to _POSIX2_BC_SCALE_MAX. */
#define _SC_BC_SCALE_MAX 0x0003
/** sysconf() query for bc(1) behavior equivalent to _POSIX2_BC_STRING_MAX. */
#define _SC_BC_STRING_MAX 0x0004
/** sysconf() query equivalent to RLIMIT_NPROC. */
#define _SC_CHILD_MAX 0x0005
/** sysconf() query equivalent to AT_CLKTCK. */
#define _SC_CLK_TCK 0x0006
/** sysconf() query for collation behavior equivalent to _POSIX2_COLL_WEIGHTS_MAX. */
#define _SC_COLL_WEIGHTS_MAX 0x0007
/** sysconf() query for expr(1) behavior equivalent to _POSIX2_EXPR_NEST_MAX. */
#define _SC_EXPR_NEST_MAX 0x0008
/** sysconf() query for command-line tool behavior equivalent to _POSIX2_LINE_MAX. */
#define _SC_LINE_MAX 0x0009
/** sysconf() query equivalent to NGROUPS_MAX. */
#define _SC_NGROUPS_MAX 0x000a
/** sysconf() query equivalent to RLIMIT_NOFILE. */
#define _SC_OPEN_MAX 0x000b
/** sysconf() query equivalent to PASS_MAX. */
#define _SC_PASS_MAX 0x000c
/** sysconf() query equivalent to _POSIX2_C_BIND. */
#define _SC_2_C_BIND 0x000d
/** sysconf() query equivalent to _POSIX2_C_DEV. */
#define _SC_2_C_DEV 0x000e
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_C_VERSION 0x000f
/** sysconf() query equivalent to _POSIX2_CHAR_TERM. */
#define _SC_2_CHAR_TERM 0x0010
/** sysconf() query equivalent to _POSIX2_FORT_DEV. */
#define _SC_2_FORT_DEV 0x0011
/** sysconf() query equivalent to _POSIX2_FORT_RUN. */
#define _SC_2_FORT_RUN 0x0012
/** sysconf() query equivalent to _POSIX2_LOCALEDEF. */
#define _SC_2_LOCALEDEF 0x0013
/** sysconf() query equivalent to _POSIX2_SW_DEV. */
#define _SC_2_SW_DEV 0x0014
/** sysconf() query equivalent to _POSIX2_UPE. */
#define _SC_2_UPE 0x0015
/** sysconf() query equivalent to _POSIX2_VERSION. */
#define _SC_2_VERSION 0x0016
/** sysconf() query equivalent to _POSIX_JOB_CONTROL. */
#define _SC_JOB_CONTROL 0x0017
/** sysconf() query equivalent to _POSIX_SAVED_IDS. */
#define _SC_SAVED_IDS 0x0018
/** sysconf() query equivalent to _POSIX_VERSION. */
#define _SC_VERSION 0x0019
/** sysconf() query equivalent to _POSIX_RE_DUP_MAX. */
#define _SC_RE_DUP_MAX 0x001a
/** sysconf() query equivalent to FOPEN_MAX. */
#define _SC_STREAM_MAX 0x001b
/** sysconf() query equivalent to _POSIX_TZNAME_MAX. */
#define _SC_TZNAME_MAX 0x001c
/** sysconf() query equivalent to _XOPEN_CRYPT. */
#define _SC_XOPEN_CRYPT 0x001d
/** sysconf() query equivalent to _XOPEN_ENH_I18N. */
#define _SC_XOPEN_ENH_I18N 0x001e
/** sysconf() query equivalent to _XOPEN_SHM. */
#define _SC_XOPEN_SHM 0x001f
/** sysconf() query equivalent to _XOPEN_VERSION. */
#define _SC_XOPEN_VERSION 0x0020
/** Obsolescent in POSIX.1-2008. */
#define _SC_XOPEN_XCU_VERSION 0x0021
/** sysconf() query equivalent to _XOPEN_REALTIME. */
#define _SC_XOPEN_REALTIME 0x0022
/** sysconf() query equivalent to _XOPEN_REALTIME_THREADS. */
#define _SC_XOPEN_REALTIME_THREADS 0x0023
/** sysconf() query equivalent to _XOPEN_LEGACY. */
#define _SC_XOPEN_LEGACY 0x0024
/** sysconf() query for the maximum number of atexit() handlers. Unlimited on Android. */
#define _SC_ATEXIT_MAX 0x0025
/** sysconf() query equivalent to IOV_MAX. */
#define _SC_IOV_MAX 0x0026
/** Same as _SC_IOV_MAX. */
#define _SC_UIO_MAXIOV _SC_IOV_MAX
/** Same as _SC_PAGE_SIZE. */
#define _SC_PAGESIZE 0x0027
/** sysconf() query equivalent to getpagesize(). */
#define _SC_PAGE_SIZE 0x0028
/** sysconf() query equivalent to _XOPEN_UNIX. */
#define _SC_XOPEN_UNIX 0x0029
/** Obsolescent in POSIX.1-2008. */
#define _SC_XBS5_ILP32_OFF32 0x002a
/** Obsolescent in POSIX.1-2008. */
#define _SC_XBS5_ILP32_OFFBIG 0x002b
/** Obsolescent in POSIX.1-2008. */
#define _SC_XBS5_LP64_OFF64 0x002c
/** Obsolescent in POSIX.1-2008. */
#define _SC_XBS5_LPBIG_OFFBIG 0x002d
/** sysconf() query equivalent to _POSIX_AIO_LISTIO_MAX. */
#define _SC_AIO_LISTIO_MAX 0x002e
/** sysconf() query equivalent to _POSIX_AIO_MAX. */
#define _SC_AIO_MAX 0x002f
/** Unimplemented on Android. */
#define _SC_AIO_PRIO_DELTA_MAX  0x0030
/** sysconf() query equivalent to _POSIX_DELAYTIMER_MAX. */
#define _SC_DELAYTIMER_MAX 0x0031
/** sysconf() query equivalent to _POSIX_MQ_OPEN_MAX. */
#define _SC_MQ_OPEN_MAX 0x0032
/** sysconf() query equivalent to _POSIX_MQ_PRIO_MAX. */
#define _SC_MQ_PRIO_MAX 0x0033
/** sysconf() query equivalent to RTSIG_MAX. Constant on Android. */
#define _SC_RTSIG_MAX 0x0034
/** sysconf() query equivalent to _POSIX_SEM_NSEMS_MAX. Constant on Android. */
#define _SC_SEM_NSEMS_MAX 0x0035
/** sysconf() query equivalent to SEM_VALUE_MAX. Constant on Android. */
#define _SC_SEM_VALUE_MAX 0x0036
/** sysconf() query equivalent to _POSIX_SIGQUEUE_MAX. */
#define _SC_SIGQUEUE_MAX 0x0037
/** sysconf() query equivalent to _POSIX_TIMER_MAX. */
#define _SC_TIMER_MAX 0x0038
/** sysconf() query equivalent to _POSIX_ASYNCHRONOUS_IO. */
#define _SC_ASYNCHRONOUS_IO 0x0039
/** sysconf() query equivalent to _POSIX_FSYNC. */
#define _SC_FSYNC 0x003a
/** sysconf() query equivalent to _POSIX_MAPPED_FILES. */
#define _SC_MAPPED_FILES 0x003b
/** sysconf() query equivalent to _POSIX_MEMLOCK. */
#define _SC_MEMLOCK 0x003c
/** sysconf() query equivalent to _POSIX_MEMLOCK_RANGE. */
#define _SC_MEMLOCK_RANGE 0x003d
/** sysconf() query equivalent to _POSIX_MEMORY_PROTECTION. */
#define _SC_MEMORY_PROTECTION 0x003e
/** sysconf() query equivalent to _POSIX_MESSAGE_PASSING. */
#define _SC_MESSAGE_PASSING 0x003f
/** sysconf() query equivalent to _POSIX_PRIORITIZED_IO. */
#define _SC_PRIORITIZED_IO 0x0040
/** sysconf() query equivalent to _POSIX_PRIORITY_SCHEDULING. */
#define _SC_PRIORITY_SCHEDULING 0x0041
/** sysconf() query equivalent to _POSIX_REALTIME_SIGNALS. */
#define _SC_REALTIME_SIGNALS 0x0042
/** sysconf() query equivalent to _POSIX_SEMAPHORES. */
#define _SC_SEMAPHORES 0x0043
/** sysconf() query equivalent to _POSIX_SHARED_MEMORY_OBJECTS. */
#define _SC_SHARED_MEMORY_OBJECTS 0x0044
/** sysconf() query equivalent to _POSIX_SYNCHRONIZED_IO. */
#define _SC_SYNCHRONIZED_IO 0x0045
/** sysconf() query equivalent to _POSIX_TIMERS. */
#define _SC_TIMERS 0x0046
/** sysconf() query for an initial size for getgrgid_r() and getgrnam_r() buffers. */
#define _SC_GETGR_R_SIZE_MAX 0x0047
/** sysconf() query for an initial size for getpwuid_r() and getpwnam_r() buffers. */
#define _SC_GETPW_R_SIZE_MAX 0x0048
/** sysconf() query equivalent to LOGIN_NAME_MAX. */
#define _SC_LOGIN_NAME_MAX 0x0049
/** sysconf() query equivalent to PTHREAD_DESTRUCTOR_ITERATIONS. */
#define _SC_THREAD_DESTRUCTOR_ITERATIONS 0x004a
/** sysconf() query equivalent to PTHREAD_KEYS_MAX. */
#define _SC_THREAD_KEYS_MAX 0x004b
/** sysconf() query equivalent to PTHREAD_STACK_MIN. */
#define _SC_THREAD_STACK_MIN 0x004c
/** sysconf() query for a maximum number of threads. Unlimited on Android. */
#define _SC_THREAD_THREADS_MAX 0x004d
/** sysconf() query equivalent to TTY_NAME_MAX. */
#define _SC_TTY_NAME_MAX 0x004e
/** sysconf() query equivalent to _POSIX_THREADS. */
#define _SC_THREADS 0x004f
/** sysconf() query equivalent to _POSIX_THREAD_ATTR_STACKADDR. */
#define _SC_THREAD_ATTR_STACKADDR 0x0050
/** sysconf() query equivalent to _POSIX_THREAD_ATTR_STACKSIZE. */
#define _SC_THREAD_ATTR_STACKSIZE 0x0051
/** sysconf() query equivalent to _POSIX_THREAD_PRIORITY_SCHEDULING. */
#define _SC_THREAD_PRIORITY_SCHEDULING 0x0052
/** sysconf() query equivalent to _POSIX_THREAD_PRIO_INHERIT. */
#define _SC_THREAD_PRIO_INHERIT 0x0053
/** sysconf() query equivalent to _POSIX_THREAD_PRIO_PROTECT. */
#define _SC_THREAD_PRIO_PROTECT 0x0054
/** sysconf() query equivalent to _POSIX_THREAD_SAFE_FUNCTIONS. */
#define _SC_THREAD_SAFE_FUNCTIONS 0x0055
/** sysconf() query equivalent to get_nprocs_conf(). */
#define _SC_NPROCESSORS_CONF 0x0060
/** sysconf() query equivalent to get_nprocs(). */
#define _SC_NPROCESSORS_ONLN 0x0061
/** sysconf() query equivalent to get_phys_pages(). */
#define _SC_PHYS_PAGES 0x0062
/** sysconf() query equivalent to get_avphys_pages(). */
#define _SC_AVPHYS_PAGES 0x0063
/** sysconf() query equivalent to _POSIX_MONOTONIC_CLOCK. */
#define _SC_MONOTONIC_CLOCK 0x0064
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_PBS 0x0065
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_PBS_ACCOUNTING 0x0066
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_PBS_CHECKPOINT 0x0067
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_PBS_LOCATE 0x0068
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_PBS_MESSAGE 0x0069
/** Obsolescent in POSIX.1-2008. */
#define _SC_2_PBS_TRACK 0x006a
/** sysconf() query equivalent to _POSIX_ADVISORY_INFO. */
#define _SC_ADVISORY_INFO 0x006b
/** sysconf() query equivalent to _POSIX_BARRIERS. */
#define _SC_BARRIERS 0x006c
/** sysconf() query equivalent to _POSIX_CLOCK_SELECTION. */
#define _SC_CLOCK_SELECTION 0x006d
/** sysconf() query equivalent to _POSIX_CPUTIME. */
#define _SC_CPUTIME 0x006e
/** sysconf() query equivalent to _POSIX_HOST_NAME_MAX. */
#define _SC_HOST_NAME_MAX 0x006f
/** sysconf() query equivalent to _POSIX_IPV6. */
#define _SC_IPV6 0x0070
/** sysconf() query equivalent to _POSIX_RAW_SOCKETS. */
#define _SC_RAW_SOCKETS 0x0071
/** sysconf() query equivalent to _POSIX_READER_WRITER_LOCKS. */
#define _SC_READER_WRITER_LOCKS 0x0072
/** sysconf() query equivalent to _POSIX_REGEXP. */
#define _SC_REGEXP 0x0073
/** sysconf() query equivalent to _POSIX_SHELL. */
#define _SC_SHELL 0x0074
/** sysconf() query equivalent to _POSIX_SPAWN. */
#define _SC_SPAWN 0x0075
/** sysconf() query equivalent to _POSIX_SPIN_LOCKS. */
#define _SC_SPIN_LOCKS 0x0076
/** sysconf() query equivalent to _POSIX_SPORADIC_SERVER. */
#define _SC_SPORADIC_SERVER 0x0077
/** sysconf() query equivalent to _POSIX_SS_REPL_MAX. */
#define _SC_SS_REPL_MAX 0x0078
/** sysconf() query equivalent to _POSIX_SYMLOOP_MAX. */
#define _SC_SYMLOOP_MAX 0x0079
/** sysconf() query equivalent to _POSIX_THREAD_CPUTIME. */
#define _SC_THREAD_CPUTIME 0x007a
/** sysconf() query equivalent to _POSIX_THREAD_PROCESS_SHARED. */
#define _SC_THREAD_PROCESS_SHARED 0x007b
/** sysconf() query equivalent to _POSIX_THREAD_ROBUST_PRIO_INHERIT. */
#define _SC_THREAD_ROBUST_PRIO_INHERIT 0x007c
/** sysconf() query equivalent to _POSIX_THREAD_ROBUST_PRIO_PROTECT. */
#define _SC_THREAD_ROBUST_PRIO_PROTECT 0x007d
/** sysconf() query equivalent to _POSIX_THREAD_SPORADIC_SERVER. */
#define _SC_THREAD_SPORADIC_SERVER 0x007e
/** sysconf() query equivalent to _POSIX_TIMEOUTS. */
#define _SC_TIMEOUTS 0x007f
/** Unimplemented. */
#define _SC_TRACE 0x0080
/** Unimplemented. */
#define _SC_TRACE_EVENT_FILTER 0x0081
/** Unimplemented. */
#define _SC_TRACE_EVENT_NAME_MAX 0x0082
/** Unimplemented. */
#define _SC_TRACE_INHERIT 0x0083
/** Unimplemented. */
#define _SC_TRACE_LOG 0x0084
/** Unimplemented. */
#define _SC_TRACE_NAME_MAX 0x0085
/** Unimplemented. */
#define _SC_TRACE_SYS_MAX 0x0086
/** Unimplemented. */
#define _SC_TRACE_USER_EVENT_MAX 0x0087
/** sysconf() query equivalent to _POSIX_TYPED_MEMORY_OBJECTS. */
#define _SC_TYPED_MEMORY_OBJECTS 0x0088
/** sysconf() query equivalent to _POSIX_V7_ILP32_OFF32. */
#define _SC_V7_ILP32_OFF32 0x0089
/** sysconf() query equivalent to _POSIX_V7_ILP32_OFFBIG. */
#define _SC_V7_ILP32_OFFBIG 0x008a
/** sysconf() query equivalent to _POSIX_V7_ILP64_OFF64. */
#define _SC_V7_LP64_OFF64 0x008b
/** sysconf() query equivalent to _POSIX_V7_ILP64_OFFBIG. */
#define _SC_V7_LPBIG_OFFBIG 0x008c
/** Unimplemented. */
#define _SC_XOPEN_STREAMS 0x008d
/** Meaningless in Android, unsupported in every other libc (but defined by POSIX). */
#define _SC_XOPEN_UUCP 0x008e
/** sysconf() query for the L1 instruction cache size. Not available on all architectures. */
#define _SC_LEVEL1_ICACHE_SIZE 0x008f
/** sysconf() query for the L1 instruction cache associativity. Not available on all architectures. */
#define _SC_LEVEL1_ICACHE_ASSOC 0x0090
/** sysconf() query for the L1 instruction cache line size. Not available on all architectures. */
#define _SC_LEVEL1_ICACHE_LINESIZE 0x0091
/** sysconf() query for the L1 data cache size. Not available on all architectures. */
#define _SC_LEVEL1_DCACHE_SIZE 0x0092
/** sysconf() query for the L1 data cache associativity. Not available on all architectures. */
#define _SC_LEVEL1_DCACHE_ASSOC 0x0093
/** sysconf() query for the L1 data cache line size. Not available on all architectures. */
#define _SC_LEVEL1_DCACHE_LINESIZE 0x0094
/** sysconf() query for the L2 cache size. Not available on all architectures. */
#define _SC_LEVEL2_CACHE_SIZE 0x0095
/** sysconf() query for the L2 cache associativity. Not available on all architectures. */
#define _SC_LEVEL2_CACHE_ASSOC 0x0096
/** sysconf() query for the L2 cache line size. Not available on all architectures. */
#define _SC_LEVEL2_CACHE_LINESIZE 0x0097
/** sysconf() query for the L3 cache size. Not available on all architectures. */
#define _SC_LEVEL3_CACHE_SIZE 0x0098
/** sysconf() query for the L3 cache associativity. Not available on all architectures. */
#define _SC_LEVEL3_CACHE_ASSOC 0x0099
/** sysconf() query for the L3 cache line size. Not available on all architectures. */
#define _SC_LEVEL3_CACHE_LINESIZE 0x009a
/** sysconf() query for the L4 cache size. Not available on all architectures. */
#define _SC_LEVEL4_CACHE_SIZE 0x009b
/** sysconf() query for the L4 cache associativity. Not available on all architectures. */
#define _SC_LEVEL4_CACHE_ASSOC 0x009c
/** sysconf() query for the L4 cache line size. Not available on all architectures. */
#define _SC_LEVEL4_CACHE_LINESIZE 0x009d
/** sysconf() query equivalent to NSIG. Available from API level 37. */
#define _SC_NSIG 0x009e

__BEGIN_DECLS

/**
 * [sysconf(3)](https://man7.org/linux/man-pages/man3/sysconf.3.html)
 * gets system configuration at runtime, corresponding to the given
 * `_SC_` constant. See the man page for details on how to interpret
 * the results.
 *
 * For `_SC_` constants where an equivalent is given, it's cheaper on Android
 * to go straight to that function call --- sysconf() is just a multiplexer.
 * This may not be true on other systems, and other systems may not support the
 * direct function, so sysconf() can be useful for portability, though despite
 * POSIX's best efforts, the exact set of constants that return useful results
 * will also vary by system.
 */
long sysconf(int __name);

__END_DECLS
