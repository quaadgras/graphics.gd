/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VIRTIO_NET_H
#define _UAPI_LINUX_VIRTIO_NET_H
#include <linux/types.h>
#include <linux/virtio_ids.h>
#include <linux/virtio_config.h>
#include <linux/virtio_types.h>
#include <linux/if_ether.h>
#define VIRTIO_NET_F_CSUM 0
#define VIRTIO_NET_F_GUEST_CSUM 1
#define VIRTIO_NET_F_CTRL_GUEST_OFFLOADS 2
#define VIRTIO_NET_F_MTU 3
#define VIRTIO_NET_F_MAC 5
#define VIRTIO_NET_F_GUEST_TSO4 7
#define VIRTIO_NET_F_GUEST_TSO6 8
#define VIRTIO_NET_F_GUEST_ECN 9
#define VIRTIO_NET_F_GUEST_UFO 10
#define VIRTIO_NET_F_HOST_TSO4 11
#define VIRTIO_NET_F_HOST_TSO6 12
#define VIRTIO_NET_F_HOST_ECN 13
#define VIRTIO_NET_F_HOST_UFO 14
#define VIRTIO_NET_F_MRG_RXBUF 15
#define VIRTIO_NET_F_STATUS 16
#define VIRTIO_NET_F_CTRL_VQ 17
#define VIRTIO_NET_F_CTRL_RX 18
#define VIRTIO_NET_F_CTRL_VLAN 19
#define VIRTIO_NET_F_CTRL_RX_EXTRA 20
#define VIRTIO_NET_F_GUEST_ANNOUNCE 21
#define VIRTIO_NET_F_MQ 22
#define VIRTIO_NET_F_CTRL_MAC_ADDR 23
#define VIRTIO_NET_F_DEVICE_STATS 50
#define VIRTIO_NET_F_VQ_NOTF_COAL 52
#define VIRTIO_NET_F_NOTF_COAL 53
#define VIRTIO_NET_F_GUEST_USO4 54
#define VIRTIO_NET_F_GUEST_USO6 55
#define VIRTIO_NET_F_HOST_USO 56
#define VIRTIO_NET_F_HASH_REPORT 57
#define VIRTIO_NET_F_GUEST_HDRLEN 59
#define VIRTIO_NET_F_RSS 60
#define VIRTIO_NET_F_RSC_EXT 61
#define VIRTIO_NET_F_STANDBY 62
#define VIRTIO_NET_F_SPEED_DUPLEX 63
#ifndef VIRTIO_NET_NO_LEGACY
#define VIRTIO_NET_F_GSO 6
#endif
#define VIRTIO_NET_S_LINK_UP 1
#define VIRTIO_NET_S_ANNOUNCE 2
#define VIRTIO_NET_RSS_HASH_TYPE_IPv4 (1 << 0)
#define VIRTIO_NET_RSS_HASH_TYPE_TCPv4 (1 << 1)
#define VIRTIO_NET_RSS_HASH_TYPE_UDPv4 (1 << 2)
#define VIRTIO_NET_RSS_HASH_TYPE_IPv6 (1 << 3)
#define VIRTIO_NET_RSS_HASH_TYPE_TCPv6 (1 << 4)
#define VIRTIO_NET_RSS_HASH_TYPE_UDPv6 (1 << 5)
#define VIRTIO_NET_RSS_HASH_TYPE_IP_EX (1 << 6)
#define VIRTIO_NET_RSS_HASH_TYPE_TCP_EX (1 << 7)
#define VIRTIO_NET_RSS_HASH_TYPE_UDP_EX (1 << 8)
struct virtio_net_config {
  __u8 mac[ETH_ALEN];
  __virtio16 status;
  __virtio16 max_virtqueue_pairs;
  __virtio16 mtu;
  __le32 speed;
  __u8 duplex;
  __u8 rss_max_key_size;
  __le16 rss_max_indirection_table_length;
  __le32 supported_hash_types;
} __attribute__((packed));
struct virtio_net_hdr_v1 {
#define VIRTIO_NET_HDR_F_NEEDS_CSUM 1
#define VIRTIO_NET_HDR_F_DATA_VALID 2
#define VIRTIO_NET_HDR_F_RSC_INFO 4
  __u8 flags;
#define VIRTIO_NET_HDR_GSO_NONE 0
#define VIRTIO_NET_HDR_GSO_TCPV4 1
#define VIRTIO_NET_HDR_GSO_UDP 3
#define VIRTIO_NET_HDR_GSO_TCPV6 4
#define VIRTIO_NET_HDR_GSO_UDP_L4 5
#define VIRTIO_NET_HDR_GSO_ECN 0x80
  __u8 gso_type;
  __virtio16 hdr_len;
  __virtio16 gso_size;
  union {
    struct {
      __virtio16 csum_start;
      __virtio16 csum_offset;
    };
    struct {
      __virtio16 start;
      __virtio16 offset;
    } csum;
    struct {
      __le16 segments;
      __le16 dup_acks;
    } rsc;
  };
  __virtio16 num_buffers;
};
struct virtio_net_hdr_v1_hash {
  struct virtio_net_hdr_v1 hdr;
  __le32 hash_value;
#define VIRTIO_NET_HASH_REPORT_NONE 0
#define VIRTIO_NET_HASH_REPORT_IPv4 1
#define VIRTIO_NET_HASH_REPORT_TCPv4 2
#define VIRTIO_NET_HASH_REPORT_UDPv4 3
#define VIRTIO_NET_HASH_REPORT_IPv6 4
#define VIRTIO_NET_HASH_REPORT_TCPv6 5
#define VIRTIO_NET_HASH_REPORT_UDPv6 6
#define VIRTIO_NET_HASH_REPORT_IPv6_EX 7
#define VIRTIO_NET_HASH_REPORT_TCPv6_EX 8
#define VIRTIO_NET_HASH_REPORT_UDPv6_EX 9
  __le16 hash_report;
  __le16 padding;
};
#ifndef VIRTIO_NET_NO_LEGACY
struct virtio_net_hdr {
  __u8 flags;
  __u8 gso_type;
  __virtio16 hdr_len;
  __virtio16 gso_size;
  __virtio16 csum_start;
  __virtio16 csum_offset;
};
struct virtio_net_hdr_mrg_rxbuf {
  struct virtio_net_hdr hdr;
  __virtio16 num_buffers;
};
#endif
struct virtio_net_ctrl_hdr {
  __u8 __linux_class;
  __u8 cmd;
} __attribute__((packed));
typedef __u8 virtio_net_ctrl_ack;
#define VIRTIO_NET_OK 0
#define VIRTIO_NET_ERR 1
#define VIRTIO_NET_CTRL_RX 0
#define VIRTIO_NET_CTRL_RX_PROMISC 0
#define VIRTIO_NET_CTRL_RX_ALLMULTI 1
#define VIRTIO_NET_CTRL_RX_ALLUNI 2
#define VIRTIO_NET_CTRL_RX_NOMULTI 3
#define VIRTIO_NET_CTRL_RX_NOUNI 4
#define VIRTIO_NET_CTRL_RX_NOBCAST 5
struct virtio_net_ctrl_mac {
  __virtio32 entries;
  __u8 macs[][ETH_ALEN];
} __attribute__((packed));
#define VIRTIO_NET_CTRL_MAC 1
#define VIRTIO_NET_CTRL_MAC_TABLE_SET 0
#define VIRTIO_NET_CTRL_MAC_ADDR_SET 1
#define VIRTIO_NET_CTRL_VLAN 2
#define VIRTIO_NET_CTRL_VLAN_ADD 0
#define VIRTIO_NET_CTRL_VLAN_DEL 1
#define VIRTIO_NET_CTRL_ANNOUNCE 3
#define VIRTIO_NET_CTRL_ANNOUNCE_ACK 0
#define VIRTIO_NET_CTRL_MQ 4
struct virtio_net_ctrl_mq {
  __virtio16 virtqueue_pairs;
};
#define VIRTIO_NET_CTRL_MQ_VQ_PAIRS_SET 0
#define VIRTIO_NET_CTRL_MQ_VQ_PAIRS_MIN 1
#define VIRTIO_NET_CTRL_MQ_VQ_PAIRS_MAX 0x8000
struct virtio_net_rss_config {
  __le32 hash_types;
  __le16 indirection_table_mask;
  __le16 unclassified_queue;
  __le16 indirection_table[1];
  __le16 max_tx_vq;
  __u8 hash_key_length;
  __u8 hash_key_data[];
};
#define VIRTIO_NET_CTRL_MQ_RSS_CONFIG 1
struct virtio_net_hash_config {
  __le32 hash_types;
  __le16 reserved[4];
  __u8 hash_key_length;
  __u8 hash_key_data[];
};
#define VIRTIO_NET_CTRL_MQ_HASH_CONFIG 2
#define VIRTIO_NET_CTRL_GUEST_OFFLOADS 5
#define VIRTIO_NET_CTRL_GUEST_OFFLOADS_SET 0
#define VIRTIO_NET_CTRL_NOTF_COAL 6
struct virtio_net_ctrl_coal_tx {
  __le32 tx_max_packets;
  __le32 tx_usecs;
};
#define VIRTIO_NET_CTRL_NOTF_COAL_TX_SET 0
struct virtio_net_ctrl_coal_rx {
  __le32 rx_max_packets;
  __le32 rx_usecs;
};
#define VIRTIO_NET_CTRL_NOTF_COAL_RX_SET 1
#define VIRTIO_NET_CTRL_NOTF_COAL_VQ_SET 2
#define VIRTIO_NET_CTRL_NOTF_COAL_VQ_GET 3
struct virtio_net_ctrl_coal {
  __le32 max_packets;
  __le32 max_usecs;
};
struct virtio_net_ctrl_coal_vq {
  __le16 vqn;
  __le16 reserved;
  struct virtio_net_ctrl_coal coal;
};
#define VIRTIO_NET_CTRL_STATS 8
#define VIRTIO_NET_CTRL_STATS_QUERY 0
#define VIRTIO_NET_CTRL_STATS_GET 1
struct virtio_net_stats_capabilities {
#define VIRTIO_NET_STATS_TYPE_CVQ (1ULL << 32)
#define VIRTIO_NET_STATS_TYPE_RX_BASIC (1ULL << 0)
#define VIRTIO_NET_STATS_TYPE_RX_CSUM (1ULL << 1)
#define VIRTIO_NET_STATS_TYPE_RX_GSO (1ULL << 2)
#define VIRTIO_NET_STATS_TYPE_RX_SPEED (1ULL << 3)
#define VIRTIO_NET_STATS_TYPE_TX_BASIC (1ULL << 16)
#define VIRTIO_NET_STATS_TYPE_TX_CSUM (1ULL << 17)
#define VIRTIO_NET_STATS_TYPE_TX_GSO (1ULL << 18)
#define VIRTIO_NET_STATS_TYPE_TX_SPEED (1ULL << 19)
  __le64 supported_stats_types[1];
};
struct virtio_net_ctrl_queue_stats {
  struct {
    __le16 vq_index;
    __le16 reserved[3];
    __le64 types_bitmap[1];
  } stats[1];
};
struct virtio_net_stats_reply_hdr {
#define VIRTIO_NET_STATS_TYPE_REPLY_CVQ 32
#define VIRTIO_NET_STATS_TYPE_REPLY_RX_BASIC 0
#define VIRTIO_NET_STATS_TYPE_REPLY_RX_CSUM 1
#define VIRTIO_NET_STATS_TYPE_REPLY_RX_GSO 2
#define VIRTIO_NET_STATS_TYPE_REPLY_RX_SPEED 3
#define VIRTIO_NET_STATS_TYPE_REPLY_TX_BASIC 16
#define VIRTIO_NET_STATS_TYPE_REPLY_TX_CSUM 17
#define VIRTIO_NET_STATS_TYPE_REPLY_TX_GSO 18
#define VIRTIO_NET_STATS_TYPE_REPLY_TX_SPEED 19
  __u8 type;
  __u8 reserved;
  __le16 vq_index;
  __le16 reserved1;
  __le16 size;
};
struct virtio_net_stats_cvq {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 command_num;
  __le64 ok_num;
};
struct virtio_net_stats_rx_basic {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 rx_notifications;
  __le64 rx_packets;
  __le64 rx_bytes;
  __le64 rx_interrupts;
  __le64 rx_drops;
  __le64 rx_drop_overruns;
};
struct virtio_net_stats_tx_basic {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 tx_notifications;
  __le64 tx_packets;
  __le64 tx_bytes;
  __le64 tx_interrupts;
  __le64 tx_drops;
  __le64 tx_drop_malformed;
};
struct virtio_net_stats_rx_csum {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 rx_csum_valid;
  __le64 rx_needs_csum;
  __le64 rx_csum_none;
  __le64 rx_csum_bad;
};
struct virtio_net_stats_tx_csum {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 tx_csum_none;
  __le64 tx_needs_csum;
};
struct virtio_net_stats_rx_gso {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 rx_gso_packets;
  __le64 rx_gso_bytes;
  __le64 rx_gso_packets_coalesced;
  __le64 rx_gso_bytes_coalesced;
};
struct virtio_net_stats_tx_gso {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 tx_gso_packets;
  __le64 tx_gso_bytes;
  __le64 tx_gso_segments;
  __le64 tx_gso_segments_bytes;
  __le64 tx_gso_packets_noseg;
  __le64 tx_gso_bytes_noseg;
};
struct virtio_net_stats_rx_speed {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 rx_ratelimit_packets;
  __le64 rx_ratelimit_bytes;
};
struct virtio_net_stats_tx_speed {
  struct virtio_net_stats_reply_hdr hdr;
  __le64 tx_ratelimit_packets;
  __le64 tx_ratelimit_bytes;
};
#endif
