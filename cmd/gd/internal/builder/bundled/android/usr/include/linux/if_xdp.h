/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_IF_XDP_H
#define _LINUX_IF_XDP_H
#include <linux/types.h>
#define XDP_SHARED_UMEM (1 << 0)
#define XDP_COPY (1 << 1)
#define XDP_ZEROCOPY (1 << 2)
#define XDP_USE_NEED_WAKEUP (1 << 3)
#define XDP_USE_SG (1 << 4)
#define XDP_UMEM_UNALIGNED_CHUNK_FLAG (1 << 0)
#define XDP_UMEM_TX_SW_CSUM (1 << 1)
#define XDP_UMEM_TX_METADATA_LEN (1 << 2)
struct sockaddr_xdp {
  __u16 sxdp_family;
  __u16 sxdp_flags;
  __u32 sxdp_ifindex;
  __u32 sxdp_queue_id;
  __u32 sxdp_shared_umem_fd;
};
#define XDP_RING_NEED_WAKEUP (1 << 0)
struct xdp_ring_offset {
  __u64 producer;
  __u64 consumer;
  __u64 desc;
  __u64 flags;
};
struct xdp_mmap_offsets {
  struct xdp_ring_offset rx;
  struct xdp_ring_offset tx;
  struct xdp_ring_offset fr;
  struct xdp_ring_offset cr;
};
#define XDP_MMAP_OFFSETS 1
#define XDP_RX_RING 2
#define XDP_TX_RING 3
#define XDP_UMEM_REG 4
#define XDP_UMEM_FILL_RING 5
#define XDP_UMEM_COMPLETION_RING 6
#define XDP_STATISTICS 7
#define XDP_OPTIONS 8
struct xdp_umem_reg {
  __u64 addr;
  __u64 len;
  __u32 chunk_size;
  __u32 headroom;
  __u32 flags;
  __u32 tx_metadata_len;
};
struct xdp_statistics {
  __u64 rx_dropped;
  __u64 rx_invalid_descs;
  __u64 tx_invalid_descs;
  __u64 rx_ring_full;
  __u64 rx_fill_ring_empty_descs;
  __u64 tx_ring_empty_descs;
};
struct xdp_options {
  __u32 flags;
};
#define XDP_OPTIONS_ZEROCOPY (1 << 0)
#define XDP_PGOFF_RX_RING 0
#define XDP_PGOFF_TX_RING 0x80000000
#define XDP_UMEM_PGOFF_FILL_RING 0x100000000ULL
#define XDP_UMEM_PGOFF_COMPLETION_RING 0x180000000ULL
#define XSK_UNALIGNED_BUF_OFFSET_SHIFT 48
#define XSK_UNALIGNED_BUF_ADDR_MASK ((1ULL << XSK_UNALIGNED_BUF_OFFSET_SHIFT) - 1)
#define XDP_TXMD_FLAGS_TIMESTAMP (1 << 0)
#define XDP_TXMD_FLAGS_CHECKSUM (1 << 1)
struct xsk_tx_metadata {
  __u64 flags;
  union {
    struct {
      __u16 csum_start;
      __u16 csum_offset;
    } request;
    struct {
      __u64 tx_timestamp;
    } completion;
  };
};
struct xdp_desc {
  __u64 addr;
  __u32 len;
  __u32 options;
};
#define XDP_PKT_CONTD (1 << 0)
#define XDP_TX_METADATA (1 << 1)
#endif
