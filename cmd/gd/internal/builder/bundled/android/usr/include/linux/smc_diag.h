/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_SMC_DIAG_H_
#define _UAPI_SMC_DIAG_H_
#include <linux/types.h>
#include <linux/inet_diag.h>
#include <rdma/ib_user_verbs.h>
struct smc_diag_req {
  __u8 diag_family;
  __u8 pad[2];
  __u8 diag_ext;
  struct inet_diag_sockid id;
};
struct smc_diag_msg {
  __u8 diag_family;
  __u8 diag_state;
  union {
    __u8 diag_mode;
    __u8 diag_fallback;
  };
  __u8 diag_shutdown;
  struct inet_diag_sockid id;
  __u32 diag_uid;
  __aligned_u64 diag_inode;
};
enum {
  SMC_DIAG_MODE_SMCR,
  SMC_DIAG_MODE_FALLBACK_TCP,
  SMC_DIAG_MODE_SMCD,
};
enum {
  SMC_DIAG_NONE,
  SMC_DIAG_CONNINFO,
  SMC_DIAG_LGRINFO,
  SMC_DIAG_SHUTDOWN,
  SMC_DIAG_DMBINFO,
  SMC_DIAG_FALLBACK,
  __SMC_DIAG_MAX,
};
#define SMC_DIAG_MAX (__SMC_DIAG_MAX - 1)
struct smc_diag_cursor {
  __u16 reserved;
  __u16 wrap;
  __u32 count;
};
struct smc_diag_conninfo {
  __u32 token;
  __u32 sndbuf_size;
  __u32 rmbe_size;
  __u32 peer_rmbe_size;
  struct smc_diag_cursor rx_prod;
  struct smc_diag_cursor rx_cons;
  struct smc_diag_cursor tx_prod;
  struct smc_diag_cursor tx_cons;
  __u8 rx_prod_flags;
  __u8 rx_conn_state_flags;
  __u8 tx_prod_flags;
  __u8 tx_conn_state_flags;
  struct smc_diag_cursor tx_prep;
  struct smc_diag_cursor tx_sent;
  struct smc_diag_cursor tx_fin;
};
struct smc_diag_linkinfo {
  __u8 link_id;
  __u8 ibname[IB_DEVICE_NAME_MAX];
  __u8 ibport;
  __u8 gid[40];
  __u8 peer_gid[40];
};
struct smc_diag_lgrinfo {
  struct smc_diag_linkinfo lnk[1];
  __u8 role;
};
struct smc_diag_fallback {
  __u32 reason;
  __u32 peer_diagnosis;
};
struct smcd_diag_dmbinfo {
  __u32 linkid;
  __aligned_u64 peer_gid;
  __aligned_u64 my_gid;
  __aligned_u64 token;
  __aligned_u64 peer_token;
  __aligned_u64 peer_gid_ext;
  __aligned_u64 my_gid_ext;
};
#endif
