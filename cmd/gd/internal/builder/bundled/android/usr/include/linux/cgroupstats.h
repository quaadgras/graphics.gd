/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_CGROUPSTATS_H
#define _LINUX_CGROUPSTATS_H
#include <linux/types.h>
#include <linux/taskstats.h>
struct cgroupstats {
  __u64 nr_sleeping;
  __u64 nr_running;
  __u64 nr_stopped;
  __u64 nr_uninterruptible;
  __u64 nr_io_wait;
};
enum {
  CGROUPSTATS_CMD_UNSPEC = __TASKSTATS_CMD_MAX,
  CGROUPSTATS_CMD_GET,
  CGROUPSTATS_CMD_NEW,
  __CGROUPSTATS_CMD_MAX,
};
#define CGROUPSTATS_CMD_MAX (__CGROUPSTATS_CMD_MAX - 1)
enum {
  CGROUPSTATS_TYPE_UNSPEC = 0,
  CGROUPSTATS_TYPE_CGROUP_STATS,
  __CGROUPSTATS_TYPE_MAX,
};
#define CGROUPSTATS_TYPE_MAX (__CGROUPSTATS_TYPE_MAX - 1)
enum {
  CGROUPSTATS_CMD_ATTR_UNSPEC = 0,
  CGROUPSTATS_CMD_ATTR_FD,
  __CGROUPSTATS_CMD_ATTR_MAX,
};
#define CGROUPSTATS_CMD_ATTR_MAX (__CGROUPSTATS_CMD_ATTR_MAX - 1)
#endif
