/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_NEXTHOP_H
#define _UAPI_LINUX_NEXTHOP_H
#include <linux/types.h>
struct nhmsg {
  unsigned char nh_family;
  unsigned char nh_scope;
  unsigned char nh_protocol;
  unsigned char resvd;
  unsigned int nh_flags;
};
struct nexthop_grp {
  __u32 id;
  __u8 weight;
  __u8 weight_high;
  __u16 resvd2;
};
enum {
  NEXTHOP_GRP_TYPE_MPATH,
  NEXTHOP_GRP_TYPE_RES,
  __NEXTHOP_GRP_TYPE_MAX,
};
#define NEXTHOP_GRP_TYPE_MAX (__NEXTHOP_GRP_TYPE_MAX - 1)
#define NHA_OP_FLAG_DUMP_STATS BIT(0)
#define NHA_OP_FLAG_DUMP_HW_STATS BIT(1)
#define NHA_OP_FLAG_RESP_GRP_RESVD_0 BIT(31)
enum {
  NHA_UNSPEC,
  NHA_ID,
  NHA_GROUP,
  NHA_GROUP_TYPE,
  NHA_BLACKHOLE,
  NHA_OIF,
  NHA_GATEWAY,
  NHA_ENCAP_TYPE,
  NHA_ENCAP,
  NHA_GROUPS,
  NHA_MASTER,
  NHA_FDB,
  NHA_RES_GROUP,
  NHA_RES_BUCKET,
  NHA_OP_FLAGS,
  NHA_GROUP_STATS,
  NHA_HW_STATS_ENABLE,
  NHA_HW_STATS_USED,
  __NHA_MAX,
};
#define NHA_MAX (__NHA_MAX - 1)
enum {
  NHA_RES_GROUP_UNSPEC,
  NHA_RES_GROUP_PAD = NHA_RES_GROUP_UNSPEC,
  NHA_RES_GROUP_BUCKETS,
  NHA_RES_GROUP_IDLE_TIMER,
  NHA_RES_GROUP_UNBALANCED_TIMER,
  NHA_RES_GROUP_UNBALANCED_TIME,
  __NHA_RES_GROUP_MAX,
};
#define NHA_RES_GROUP_MAX (__NHA_RES_GROUP_MAX - 1)
enum {
  NHA_RES_BUCKET_UNSPEC,
  NHA_RES_BUCKET_PAD = NHA_RES_BUCKET_UNSPEC,
  NHA_RES_BUCKET_INDEX,
  NHA_RES_BUCKET_IDLE_TIME,
  NHA_RES_BUCKET_NH_ID,
  __NHA_RES_BUCKET_MAX,
};
#define NHA_RES_BUCKET_MAX (__NHA_RES_BUCKET_MAX - 1)
enum {
  NHA_GROUP_STATS_UNSPEC,
  NHA_GROUP_STATS_ENTRY,
  __NHA_GROUP_STATS_MAX,
};
#define NHA_GROUP_STATS_MAX (__NHA_GROUP_STATS_MAX - 1)
enum {
  NHA_GROUP_STATS_ENTRY_UNSPEC,
  NHA_GROUP_STATS_ENTRY_ID,
  NHA_GROUP_STATS_ENTRY_PACKETS,
  NHA_GROUP_STATS_ENTRY_PACKETS_HW,
  __NHA_GROUP_STATS_ENTRY_MAX,
};
#define NHA_GROUP_STATS_ENTRY_MAX (__NHA_GROUP_STATS_ENTRY_MAX - 1)
#endif
