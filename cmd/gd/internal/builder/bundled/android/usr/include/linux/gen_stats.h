/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __LINUX_GEN_STATS_H
#define __LINUX_GEN_STATS_H
#include <linux/types.h>
enum {
  TCA_STATS_UNSPEC,
  TCA_STATS_BASIC,
  TCA_STATS_RATE_EST,
  TCA_STATS_QUEUE,
  TCA_STATS_APP,
  TCA_STATS_RATE_EST64,
  TCA_STATS_PAD,
  TCA_STATS_BASIC_HW,
  TCA_STATS_PKT64,
  __TCA_STATS_MAX,
};
#define TCA_STATS_MAX (__TCA_STATS_MAX - 1)
struct gnet_stats_basic {
  __u64 bytes;
  __u32 packets;
};
struct gnet_stats_rate_est {
  __u32 bps;
  __u32 pps;
};
struct gnet_stats_rate_est64 {
  __u64 bps;
  __u64 pps;
};
struct gnet_stats_queue {
  __u32 qlen;
  __u32 backlog;
  __u32 drops;
  __u32 requeues;
  __u32 overlimits;
};
struct gnet_estimator {
  signed char interval;
  unsigned char ewma_log;
};
#endif
