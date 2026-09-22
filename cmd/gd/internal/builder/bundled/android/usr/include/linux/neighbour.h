/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __LINUX_NEIGHBOUR_H
#define __LINUX_NEIGHBOUR_H
#include <linux/types.h>
#include <linux/netlink.h>
struct ndmsg {
  __u8 ndm_family;
  __u8 ndm_pad1;
  __u16 ndm_pad2;
  __s32 ndm_ifindex;
  __u16 ndm_state;
  __u8 ndm_flags;
  __u8 ndm_type;
};
enum {
  NDA_UNSPEC,
  NDA_DST,
  NDA_LLADDR,
  NDA_CACHEINFO,
  NDA_PROBES,
  NDA_VLAN,
  NDA_PORT,
  NDA_VNI,
  NDA_IFINDEX,
  NDA_MASTER,
  NDA_LINK_NETNSID,
  NDA_SRC_VNI,
  NDA_PROTOCOL,
  NDA_NH_ID,
  NDA_FDB_EXT_ATTRS,
  NDA_FLAGS_EXT,
  NDA_NDM_STATE_MASK,
  NDA_NDM_FLAGS_MASK,
  __NDA_MAX
};
#define NDA_MAX (__NDA_MAX - 1)
#define NTF_USE (1 << 0)
#define NTF_SELF (1 << 1)
#define NTF_MASTER (1 << 2)
#define NTF_PROXY (1 << 3)
#define NTF_EXT_LEARNED (1 << 4)
#define NTF_OFFLOADED (1 << 5)
#define NTF_STICKY (1 << 6)
#define NTF_ROUTER (1 << 7)
#define NTF_EXT_MANAGED (1 << 0)
#define NTF_EXT_LOCKED (1 << 1)
#define NUD_INCOMPLETE 0x01
#define NUD_REACHABLE 0x02
#define NUD_STALE 0x04
#define NUD_DELAY 0x08
#define NUD_PROBE 0x10
#define NUD_FAILED 0x20
#define NUD_NOARP 0x40
#define NUD_PERMANENT 0x80
#define NUD_NONE 0x00
struct nda_cacheinfo {
  __u32 ndm_confirmed;
  __u32 ndm_used;
  __u32 ndm_updated;
  __u32 ndm_refcnt;
};
struct ndt_stats {
  __u64 ndts_allocs;
  __u64 ndts_destroys;
  __u64 ndts_hash_grows;
  __u64 ndts_res_failed;
  __u64 ndts_lookups;
  __u64 ndts_hits;
  __u64 ndts_rcv_probes_mcast;
  __u64 ndts_rcv_probes_ucast;
  __u64 ndts_periodic_gc_runs;
  __u64 ndts_forced_gc_runs;
  __u64 ndts_table_fulls;
};
enum {
  NDTPA_UNSPEC,
  NDTPA_IFINDEX,
  NDTPA_REFCNT,
  NDTPA_REACHABLE_TIME,
  NDTPA_BASE_REACHABLE_TIME,
  NDTPA_RETRANS_TIME,
  NDTPA_GC_STALETIME,
  NDTPA_DELAY_PROBE_TIME,
  NDTPA_QUEUE_LEN,
  NDTPA_APP_PROBES,
  NDTPA_UCAST_PROBES,
  NDTPA_MCAST_PROBES,
  NDTPA_ANYCAST_DELAY,
  NDTPA_PROXY_DELAY,
  NDTPA_PROXY_QLEN,
  NDTPA_LOCKTIME,
  NDTPA_QUEUE_LENBYTES,
  NDTPA_MCAST_REPROBES,
  NDTPA_PAD,
  NDTPA_INTERVAL_PROBE_TIME_MS,
  __NDTPA_MAX
};
#define NDTPA_MAX (__NDTPA_MAX - 1)
struct ndtmsg {
  __u8 ndtm_family;
  __u8 ndtm_pad1;
  __u16 ndtm_pad2;
};
struct ndt_config {
  __u16 ndtc_key_len;
  __u16 ndtc_entry_size;
  __u32 ndtc_entries;
  __u32 ndtc_last_flush;
  __u32 ndtc_last_rand;
  __u32 ndtc_hash_rnd;
  __u32 ndtc_hash_mask;
  __u32 ndtc_hash_chain_gc;
  __u32 ndtc_proxy_qlen;
};
enum {
  NDTA_UNSPEC,
  NDTA_NAME,
  NDTA_THRESH1,
  NDTA_THRESH2,
  NDTA_THRESH3,
  NDTA_CONFIG,
  NDTA_PARMS,
  NDTA_STATS,
  NDTA_GC_INTERVAL,
  NDTA_PAD,
  __NDTA_MAX
};
#define NDTA_MAX (__NDTA_MAX - 1)
enum {
  FDB_NOTIFY_BIT = (1 << 0),
  FDB_NOTIFY_INACTIVE_BIT = (1 << 1)
};
enum {
  NFEA_UNSPEC,
  NFEA_ACTIVITY_NOTIFY,
  NFEA_DONT_REFRESH,
  __NFEA_MAX
};
#define NFEA_MAX (__NFEA_MAX - 1)
#endif
