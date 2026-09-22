/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPICN_PROC_H
#define _UAPICN_PROC_H
#include <linux/types.h>
enum proc_cn_mcast_op {
  PROC_CN_MCAST_LISTEN = 1,
  PROC_CN_MCAST_IGNORE = 2
};
#define PROC_EVENT_ALL (PROC_EVENT_FORK | PROC_EVENT_EXEC | PROC_EVENT_UID | PROC_EVENT_GID | PROC_EVENT_SID | PROC_EVENT_PTRACE | PROC_EVENT_COMM | PROC_EVENT_NONZERO_EXIT | PROC_EVENT_COREDUMP | PROC_EVENT_EXIT)
enum proc_cn_event {
  PROC_EVENT_NONE = 0x00000000,
  PROC_EVENT_FORK = 0x00000001,
  PROC_EVENT_EXEC = 0x00000002,
  PROC_EVENT_UID = 0x00000004,
  PROC_EVENT_GID = 0x00000040,
  PROC_EVENT_SID = 0x00000080,
  PROC_EVENT_PTRACE = 0x00000100,
  PROC_EVENT_COMM = 0x00000200,
  PROC_EVENT_NONZERO_EXIT = 0x20000000,
  PROC_EVENT_COREDUMP = 0x40000000,
  PROC_EVENT_EXIT = 0x80000000
};
struct proc_input {
  enum proc_cn_mcast_op mcast_op;
  enum proc_cn_event event_type;
};
struct proc_event {
  enum proc_cn_event what;
  __u32 cpu;
  __u64 __attribute__((aligned(8))) timestamp_ns;
  union {
    struct {
      __u32 err;
    } ack;
    struct fork_proc_event {
      __kernel_pid_t parent_pid;
      __kernel_pid_t parent_tgid;
      __kernel_pid_t child_pid;
      __kernel_pid_t child_tgid;
    } fork;
    struct exec_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
    } exec;
    struct id_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
      union {
        __u32 ruid;
        __u32 rgid;
      } r;
      union {
        __u32 euid;
        __u32 egid;
      } e;
    } id;
    struct sid_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
    } sid;
    struct ptrace_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
      __kernel_pid_t tracer_pid;
      __kernel_pid_t tracer_tgid;
    } ptrace;
    struct comm_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
      char comm[16];
    } comm;
    struct coredump_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
      __kernel_pid_t parent_pid;
      __kernel_pid_t parent_tgid;
    } coredump;
    struct exit_proc_event {
      __kernel_pid_t process_pid;
      __kernel_pid_t process_tgid;
      __u32 exit_code, exit_signal;
      __kernel_pid_t parent_pid;
      __kernel_pid_t parent_tgid;
    } exit;
  } event_data;
};
#endif
