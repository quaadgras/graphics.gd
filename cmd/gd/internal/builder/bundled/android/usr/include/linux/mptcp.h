/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_MPTCP_H
#define _UAPI_MPTCP_H
#include <netinet/in.h>
#include <sys/socket.h>
#include <linux/const.h>
#include <linux/types.h>
#include <linux/in.h>
#include <linux/in6.h>
#include <linux/socket.h>
#define MPTCP_SUBFLOW_FLAG_MCAP_REM _BITUL(0)
#define MPTCP_SUBFLOW_FLAG_MCAP_LOC _BITUL(1)
#define MPTCP_SUBFLOW_FLAG_JOIN_REM _BITUL(2)
#define MPTCP_SUBFLOW_FLAG_JOIN_LOC _BITUL(3)
#define MPTCP_SUBFLOW_FLAG_BKUP_REM _BITUL(4)
#define MPTCP_SUBFLOW_FLAG_BKUP_LOC _BITUL(5)
#define MPTCP_SUBFLOW_FLAG_FULLY_ESTABLISHED _BITUL(6)
#define MPTCP_SUBFLOW_FLAG_CONNECTED _BITUL(7)
#define MPTCP_SUBFLOW_FLAG_MAPVALID _BITUL(8)
#define MPTCP_PM_CMD_GRP_NAME "mptcp_pm_cmds"
#define MPTCP_PM_EV_GRP_NAME "mptcp_pm_events"
#include <linux/mptcp_pm.h>
#define MPTCP_INFO_FLAG_FALLBACK _BITUL(0)
#define MPTCP_INFO_FLAG_REMOTE_KEY_RECEIVED _BITUL(1)
#define MPTCP_PM_ADDR_FLAG_SIGNAL (1 << 0)
#define MPTCP_PM_ADDR_FLAG_SUBFLOW (1 << 1)
#define MPTCP_PM_ADDR_FLAG_BACKUP (1 << 2)
#define MPTCP_PM_ADDR_FLAG_FULLMESH (1 << 3)
#define MPTCP_PM_ADDR_FLAG_IMPLICIT (1 << 4)
struct mptcp_info {
  __u8 mptcpi_subflows;
  __u8 mptcpi_add_addr_signal;
  __u8 mptcpi_add_addr_accepted;
  __u8 mptcpi_subflows_max;
  __u8 mptcpi_add_addr_signal_max;
  __u8 mptcpi_add_addr_accepted_max;
  __u32 mptcpi_flags;
  __u32 mptcpi_token;
  __u64 mptcpi_write_seq;
  __u64 mptcpi_snd_una;
  __u64 mptcpi_rcv_nxt;
  __u8 mptcpi_local_addr_used;
  __u8 mptcpi_local_addr_max;
  __u8 mptcpi_csum_enabled;
  __u32 mptcpi_retransmits;
  __u64 mptcpi_bytes_retrans;
  __u64 mptcpi_bytes_sent;
  __u64 mptcpi_bytes_received;
  __u64 mptcpi_bytes_acked;
  __u8 mptcpi_subflows_total;
  __u8 reserved[3];
  __u32 mptcpi_last_data_sent;
  __u32 mptcpi_last_data_recv;
  __u32 mptcpi_last_ack_recv;
};
#define MPTCP_RST_EUNSPEC 0
#define MPTCP_RST_EMPTCP 1
#define MPTCP_RST_ERESOURCE 2
#define MPTCP_RST_EPROHIBIT 3
#define MPTCP_RST_EWQ2BIG 4
#define MPTCP_RST_EBADPERF 5
#define MPTCP_RST_EMIDDLEBOX 6
struct mptcp_subflow_data {
  __u32 size_subflow_data;
  __u32 num_subflows;
  __u32 size_kernel;
  __u32 size_user;
} __attribute__((aligned(8)));
struct mptcp_subflow_addrs {
  union {
    __kernel_sa_family_t sa_family;
    struct sockaddr sa_local;
    struct sockaddr_in sin_local;
    struct sockaddr_in6 sin6_local;
    struct sockaddr_storage ss_local;
  };
  union {
    struct sockaddr sa_remote;
    struct sockaddr_in sin_remote;
    struct sockaddr_in6 sin6_remote;
    struct sockaddr_storage ss_remote;
  };
};
struct mptcp_subflow_info {
  __u32 id;
  struct mptcp_subflow_addrs addrs;
};
struct mptcp_full_info {
  __u32 size_tcpinfo_kernel;
  __u32 size_tcpinfo_user;
  __u32 size_sfinfo_kernel;
  __u32 size_sfinfo_user;
  __u32 num_subflows;
  __u32 size_arrays_user;
  __aligned_u64 subflow_info;
  __aligned_u64 tcp_info;
  struct mptcp_info mptcp_info;
};
#define MPTCP_INFO 1
#define MPTCP_TCPINFO 2
#define MPTCP_SUBFLOW_ADDRS 3
#define MPTCP_FULL_INFO 4
#endif
