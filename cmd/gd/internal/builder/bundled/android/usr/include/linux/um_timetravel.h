/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_UM_TIMETRAVEL_H
#define _UAPI_LINUX_UM_TIMETRAVEL_H
#include <linux/types.h>
struct um_timetravel_msg {
  __u32 op;
  __u32 seq;
  __u64 time;
};
#define UM_TIMETRAVEL_MAX_FDS 2
enum um_timetravel_shared_mem_fds {
  UM_TIMETRAVEL_SHARED_MEMFD,
  UM_TIMETRAVEL_SHARED_LOGFD,
  UM_TIMETRAVEL_SHARED_MAX_FDS,
};
enum um_timetravel_start_ack {
  UM_TIMETRAVEL_START_ACK_ID = 0xffff,
};
enum um_timetravel_ops {
  UM_TIMETRAVEL_ACK = 0,
  UM_TIMETRAVEL_START = 1,
  UM_TIMETRAVEL_REQUEST = 2,
  UM_TIMETRAVEL_WAIT = 3,
  UM_TIMETRAVEL_GET = 4,
  UM_TIMETRAVEL_UPDATE = 5,
  UM_TIMETRAVEL_RUN = 6,
  UM_TIMETRAVEL_FREE_UNTIL = 7,
  UM_TIMETRAVEL_GET_TOD = 8,
  UM_TIMETRAVEL_BROADCAST = 9,
};
#define UM_TIMETRAVEL_SCHEDSHM_VERSION 2
enum um_timetravel_schedshm_cap {
  UM_TIMETRAVEL_SCHEDSHM_CAP_TIME_SHARE = 0x1,
};
enum um_timetravel_schedshm_flags {
  UM_TIMETRAVEL_SCHEDSHM_FLAGS_REQ_RUN = 0x1,
};
union um_timetravel_schedshm_client {
  struct {
    __u32 capa;
    __u32 flags;
    __u64 req_time;
    __u64 name;
  };
  char reserve[128];
};
struct um_timetravel_schedshm {
  union {
    struct {
      __u32 version;
      __u32 len;
      __u64 free_until;
      __u64 current_time;
      __u16 running_id;
      __u16 max_clients;
    };
    char hdr[4096];
  };
  union um_timetravel_schedshm_client clients[];
};
#endif
