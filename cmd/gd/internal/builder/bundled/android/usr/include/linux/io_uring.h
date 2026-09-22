/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef LINUX_IO_URING_H
#define LINUX_IO_URING_H
#include <linux/fs.h>
#include <linux/types.h>
#ifndef UAPI_LINUX_IO_URING_H_SKIP_LINUX_TIME_TYPES_H
#include <linux/time_types.h>
#endif
#ifdef __cplusplus
extern "C" {
#endif
struct io_uring_sqe {
  __u8 opcode;
  __u8 flags;
  __u16 ioprio;
  __s32 fd;
  union {
    __u64 off;
    __u64 addr2;
    struct {
      __u32 cmd_op;
      __u32 __pad1;
    };
  };
  union {
    __u64 addr;
    __u64 splice_off_in;
    struct {
      __u32 level;
      __u32 optname;
    };
  };
  __u32 len;
  union {
    __kernel_rwf_t rw_flags;
    __u32 fsync_flags;
    __u16 poll_events;
    __u32 poll32_events;
    __u32 sync_range_flags;
    __u32 msg_flags;
    __u32 timeout_flags;
    __u32 accept_flags;
    __u32 cancel_flags;
    __u32 open_flags;
    __u32 statx_flags;
    __u32 fadvise_advice;
    __u32 splice_flags;
    __u32 rename_flags;
    __u32 unlink_flags;
    __u32 hardlink_flags;
    __u32 xattr_flags;
    __u32 msg_ring_flags;
    __u32 uring_cmd_flags;
    __u32 waitid_flags;
    __u32 futex_flags;
    __u32 install_fd_flags;
    __u32 nop_flags;
  };
  __u64 user_data;
  union {
    __u16 buf_index;
    __u16 buf_group;
  } __attribute__((packed));
  __u16 personality;
  union {
    __s32 splice_fd_in;
    __u32 file_index;
    __u32 optlen;
    struct {
      __u16 addr_len;
      __u16 __pad3[1];
    };
  };
  union {
    struct {
      __u64 addr3;
      __u64 __pad2[1];
    };
    __u64 optval;
    __u8 cmd[0];
  };
};
#define IORING_FILE_INDEX_ALLOC (~0U)
enum io_uring_sqe_flags_bit {
  IOSQE_FIXED_FILE_BIT,
  IOSQE_IO_DRAIN_BIT,
  IOSQE_IO_LINK_BIT,
  IOSQE_IO_HARDLINK_BIT,
  IOSQE_ASYNC_BIT,
  IOSQE_BUFFER_SELECT_BIT,
  IOSQE_CQE_SKIP_SUCCESS_BIT,
};
#define IOSQE_FIXED_FILE (1U << IOSQE_FIXED_FILE_BIT)
#define IOSQE_IO_DRAIN (1U << IOSQE_IO_DRAIN_BIT)
#define IOSQE_IO_LINK (1U << IOSQE_IO_LINK_BIT)
#define IOSQE_IO_HARDLINK (1U << IOSQE_IO_HARDLINK_BIT)
#define IOSQE_ASYNC (1U << IOSQE_ASYNC_BIT)
#define IOSQE_BUFFER_SELECT (1U << IOSQE_BUFFER_SELECT_BIT)
#define IOSQE_CQE_SKIP_SUCCESS (1U << IOSQE_CQE_SKIP_SUCCESS_BIT)
#define IORING_SETUP_IOPOLL (1U << 0)
#define IORING_SETUP_SQPOLL (1U << 1)
#define IORING_SETUP_SQ_AFF (1U << 2)
#define IORING_SETUP_CQSIZE (1U << 3)
#define IORING_SETUP_CLAMP (1U << 4)
#define IORING_SETUP_ATTACH_WQ (1U << 5)
#define IORING_SETUP_R_DISABLED (1U << 6)
#define IORING_SETUP_SUBMIT_ALL (1U << 7)
#define IORING_SETUP_COOP_TASKRUN (1U << 8)
#define IORING_SETUP_TASKRUN_FLAG (1U << 9)
#define IORING_SETUP_SQE128 (1U << 10)
#define IORING_SETUP_CQE32 (1U << 11)
#define IORING_SETUP_SINGLE_ISSUER (1U << 12)
#define IORING_SETUP_DEFER_TASKRUN (1U << 13)
#define IORING_SETUP_NO_MMAP (1U << 14)
#define IORING_SETUP_REGISTERED_FD_ONLY (1U << 15)
#define IORING_SETUP_NO_SQARRAY (1U << 16)
enum io_uring_op {
  IORING_OP_NOP,
  IORING_OP_READV,
  IORING_OP_WRITEV,
  IORING_OP_FSYNC,
  IORING_OP_READ_FIXED,
  IORING_OP_WRITE_FIXED,
  IORING_OP_POLL_ADD,
  IORING_OP_POLL_REMOVE,
  IORING_OP_SYNC_FILE_RANGE,
  IORING_OP_SENDMSG,
  IORING_OP_RECVMSG,
  IORING_OP_TIMEOUT,
  IORING_OP_TIMEOUT_REMOVE,
  IORING_OP_ACCEPT,
  IORING_OP_ASYNC_CANCEL,
  IORING_OP_LINK_TIMEOUT,
  IORING_OP_CONNECT,
  IORING_OP_FALLOCATE,
  IORING_OP_OPENAT,
  IORING_OP_CLOSE,
  IORING_OP_FILES_UPDATE,
  IORING_OP_STATX,
  IORING_OP_READ,
  IORING_OP_WRITE,
  IORING_OP_FADVISE,
  IORING_OP_MADVISE,
  IORING_OP_SEND,
  IORING_OP_RECV,
  IORING_OP_OPENAT2,
  IORING_OP_EPOLL_CTL,
  IORING_OP_SPLICE,
  IORING_OP_PROVIDE_BUFFERS,
  IORING_OP_REMOVE_BUFFERS,
  IORING_OP_TEE,
  IORING_OP_SHUTDOWN,
  IORING_OP_RENAMEAT,
  IORING_OP_UNLINKAT,
  IORING_OP_MKDIRAT,
  IORING_OP_SYMLINKAT,
  IORING_OP_LINKAT,
  IORING_OP_MSG_RING,
  IORING_OP_FSETXATTR,
  IORING_OP_SETXATTR,
  IORING_OP_FGETXATTR,
  IORING_OP_GETXATTR,
  IORING_OP_SOCKET,
  IORING_OP_URING_CMD,
  IORING_OP_SEND_ZC,
  IORING_OP_SENDMSG_ZC,
  IORING_OP_READ_MULTISHOT,
  IORING_OP_WAITID,
  IORING_OP_FUTEX_WAIT,
  IORING_OP_FUTEX_WAKE,
  IORING_OP_FUTEX_WAITV,
  IORING_OP_FIXED_FD_INSTALL,
  IORING_OP_FTRUNCATE,
  IORING_OP_BIND,
  IORING_OP_LISTEN,
  IORING_OP_LAST,
};
#define IORING_URING_CMD_FIXED (1U << 0)
#define IORING_URING_CMD_MASK IORING_URING_CMD_FIXED
#define IORING_FSYNC_DATASYNC (1U << 0)
#define IORING_TIMEOUT_ABS (1U << 0)
#define IORING_TIMEOUT_UPDATE (1U << 1)
#define IORING_TIMEOUT_BOOTTIME (1U << 2)
#define IORING_TIMEOUT_REALTIME (1U << 3)
#define IORING_LINK_TIMEOUT_UPDATE (1U << 4)
#define IORING_TIMEOUT_ETIME_SUCCESS (1U << 5)
#define IORING_TIMEOUT_MULTISHOT (1U << 6)
#define IORING_TIMEOUT_CLOCK_MASK (IORING_TIMEOUT_BOOTTIME | IORING_TIMEOUT_REALTIME)
#define IORING_TIMEOUT_UPDATE_MASK (IORING_TIMEOUT_UPDATE | IORING_LINK_TIMEOUT_UPDATE)
#define SPLICE_F_FD_IN_FIXED (1U << 31)
#define IORING_POLL_ADD_MULTI (1U << 0)
#define IORING_POLL_UPDATE_EVENTS (1U << 1)
#define IORING_POLL_UPDATE_USER_DATA (1U << 2)
#define IORING_POLL_ADD_LEVEL (1U << 3)
#define IORING_ASYNC_CANCEL_ALL (1U << 0)
#define IORING_ASYNC_CANCEL_FD (1U << 1)
#define IORING_ASYNC_CANCEL_ANY (1U << 2)
#define IORING_ASYNC_CANCEL_FD_FIXED (1U << 3)
#define IORING_ASYNC_CANCEL_USERDATA (1U << 4)
#define IORING_ASYNC_CANCEL_OP (1U << 5)
#define IORING_RECVSEND_POLL_FIRST (1U << 0)
#define IORING_RECV_MULTISHOT (1U << 1)
#define IORING_RECVSEND_FIXED_BUF (1U << 2)
#define IORING_SEND_ZC_REPORT_USAGE (1U << 3)
#define IORING_RECVSEND_BUNDLE (1U << 4)
#define IORING_NOTIF_USAGE_ZC_COPIED (1U << 31)
#define IORING_ACCEPT_MULTISHOT (1U << 0)
#define IORING_ACCEPT_DONTWAIT (1U << 1)
#define IORING_ACCEPT_POLL_FIRST (1U << 2)
enum io_uring_msg_ring_flags {
  IORING_MSG_DATA,
  IORING_MSG_SEND_FD,
};
#define IORING_MSG_RING_CQE_SKIP (1U << 0)
#define IORING_MSG_RING_FLAGS_PASS (1U << 1)
#define IORING_FIXED_FD_NO_CLOEXEC (1U << 0)
#define IORING_NOP_INJECT_RESULT (1U << 0)
struct io_uring_cqe {
  __u64 user_data;
  __s32 res;
  __u32 flags;
  __u64 big_cqe[];
};
#define IORING_CQE_F_BUFFER (1U << 0)
#define IORING_CQE_F_MORE (1U << 1)
#define IORING_CQE_F_SOCK_NONEMPTY (1U << 2)
#define IORING_CQE_F_NOTIF (1U << 3)
#define IORING_CQE_F_BUF_MORE (1U << 4)
#define IORING_CQE_BUFFER_SHIFT 16
#define IORING_OFF_SQ_RING 0ULL
#define IORING_OFF_CQ_RING 0x8000000ULL
#define IORING_OFF_SQES 0x10000000ULL
#define IORING_OFF_PBUF_RING 0x80000000ULL
#define IORING_OFF_PBUF_SHIFT 16
#define IORING_OFF_MMAP_MASK 0xf8000000ULL
struct io_sqring_offsets {
  __u32 head;
  __u32 tail;
  __u32 ring_mask;
  __u32 ring_entries;
  __u32 flags;
  __u32 dropped;
  __u32 array;
  __u32 resv1;
  __u64 user_addr;
};
#define IORING_SQ_NEED_WAKEUP (1U << 0)
#define IORING_SQ_CQ_OVERFLOW (1U << 1)
#define IORING_SQ_TASKRUN (1U << 2)
struct io_cqring_offsets {
  __u32 head;
  __u32 tail;
  __u32 ring_mask;
  __u32 ring_entries;
  __u32 overflow;
  __u32 cqes;
  __u32 flags;
  __u32 resv1;
  __u64 user_addr;
};
#define IORING_CQ_EVENTFD_DISABLED (1U << 0)
#define IORING_ENTER_GETEVENTS (1U << 0)
#define IORING_ENTER_SQ_WAKEUP (1U << 1)
#define IORING_ENTER_SQ_WAIT (1U << 2)
#define IORING_ENTER_EXT_ARG (1U << 3)
#define IORING_ENTER_REGISTERED_RING (1U << 4)
#define IORING_ENTER_ABS_TIMER (1U << 5)
struct io_uring_params {
  __u32 sq_entries;
  __u32 cq_entries;
  __u32 flags;
  __u32 sq_thread_cpu;
  __u32 sq_thread_idle;
  __u32 features;
  __u32 wq_fd;
  __u32 resv[3];
  struct io_sqring_offsets sq_off;
  struct io_cqring_offsets cq_off;
};
#define IORING_FEAT_SINGLE_MMAP (1U << 0)
#define IORING_FEAT_NODROP (1U << 1)
#define IORING_FEAT_SUBMIT_STABLE (1U << 2)
#define IORING_FEAT_RW_CUR_POS (1U << 3)
#define IORING_FEAT_CUR_PERSONALITY (1U << 4)
#define IORING_FEAT_FAST_POLL (1U << 5)
#define IORING_FEAT_POLL_32BITS (1U << 6)
#define IORING_FEAT_SQPOLL_NONFIXED (1U << 7)
#define IORING_FEAT_EXT_ARG (1U << 8)
#define IORING_FEAT_NATIVE_WORKERS (1U << 9)
#define IORING_FEAT_RSRC_TAGS (1U << 10)
#define IORING_FEAT_CQE_SKIP (1U << 11)
#define IORING_FEAT_LINKED_FILE (1U << 12)
#define IORING_FEAT_REG_REG_RING (1U << 13)
#define IORING_FEAT_RECVSEND_BUNDLE (1U << 14)
#define IORING_FEAT_MIN_TIMEOUT (1U << 15)
enum io_uring_register_op {
  IORING_REGISTER_BUFFERS = 0,
  IORING_UNREGISTER_BUFFERS = 1,
  IORING_REGISTER_FILES = 2,
  IORING_UNREGISTER_FILES = 3,
  IORING_REGISTER_EVENTFD = 4,
  IORING_UNREGISTER_EVENTFD = 5,
  IORING_REGISTER_FILES_UPDATE = 6,
  IORING_REGISTER_EVENTFD_ASYNC = 7,
  IORING_REGISTER_PROBE = 8,
  IORING_REGISTER_PERSONALITY = 9,
  IORING_UNREGISTER_PERSONALITY = 10,
  IORING_REGISTER_RESTRICTIONS = 11,
  IORING_REGISTER_ENABLE_RINGS = 12,
  IORING_REGISTER_FILES2 = 13,
  IORING_REGISTER_FILES_UPDATE2 = 14,
  IORING_REGISTER_BUFFERS2 = 15,
  IORING_REGISTER_BUFFERS_UPDATE = 16,
  IORING_REGISTER_IOWQ_AFF = 17,
  IORING_UNREGISTER_IOWQ_AFF = 18,
  IORING_REGISTER_IOWQ_MAX_WORKERS = 19,
  IORING_REGISTER_RING_FDS = 20,
  IORING_UNREGISTER_RING_FDS = 21,
  IORING_REGISTER_PBUF_RING = 22,
  IORING_UNREGISTER_PBUF_RING = 23,
  IORING_REGISTER_SYNC_CANCEL = 24,
  IORING_REGISTER_FILE_ALLOC_RANGE = 25,
  IORING_REGISTER_PBUF_STATUS = 26,
  IORING_REGISTER_NAPI = 27,
  IORING_UNREGISTER_NAPI = 28,
  IORING_REGISTER_CLOCK = 29,
  IORING_REGISTER_CLONE_BUFFERS = 30,
  IORING_REGISTER_LAST,
  IORING_REGISTER_USE_REGISTERED_RING = 1U << 31
};
enum io_wq_type {
  IO_WQ_BOUND,
  IO_WQ_UNBOUND,
};
struct io_uring_files_update {
  __u32 offset;
  __u32 resv;
  __aligned_u64 fds;
};
#define IORING_RSRC_REGISTER_SPARSE (1U << 0)
struct io_uring_rsrc_register {
  __u32 nr;
  __u32 flags;
  __u64 resv2;
  __aligned_u64 data;
  __aligned_u64 tags;
};
struct io_uring_rsrc_update {
  __u32 offset;
  __u32 resv;
  __aligned_u64 data;
};
struct io_uring_rsrc_update2 {
  __u32 offset;
  __u32 resv;
  __aligned_u64 data;
  __aligned_u64 tags;
  __u32 nr;
  __u32 resv2;
};
#define IORING_REGISTER_FILES_SKIP (- 2)
#define IO_URING_OP_SUPPORTED (1U << 0)
struct io_uring_probe_op {
  __u8 op;
  __u8 resv;
  __u16 flags;
  __u32 resv2;
};
struct io_uring_probe {
  __u8 last_op;
  __u8 ops_len;
  __u16 resv;
  __u32 resv2[3];
  struct io_uring_probe_op ops[];
};
struct io_uring_restriction {
  __u16 opcode;
  union {
    __u8 register_op;
    __u8 sqe_op;
    __u8 sqe_flags;
  };
  __u8 resv;
  __u32 resv2[3];
};
struct io_uring_clock_register {
  __u32 clockid;
  __u32 __resv[3];
};
enum {
  IORING_REGISTER_SRC_REGISTERED = 1,
};
struct io_uring_clone_buffers {
  __u32 src_fd;
  __u32 flags;
  __u32 pad[6];
};
struct io_uring_buf {
  __u64 addr;
  __u32 len;
  __u16 bid;
  __u16 resv;
};
struct io_uring_buf_ring {
  union {
    struct {
      __u64 resv1;
      __u32 resv2;
      __u16 resv3;
      __u16 tail;
    };
    __DECLARE_FLEX_ARRAY(struct io_uring_buf, bufs);
  };
};
enum io_uring_register_pbuf_ring_flags {
  IOU_PBUF_RING_MMAP = 1,
  IOU_PBUF_RING_INC = 2,
};
struct io_uring_buf_reg {
  __u64 ring_addr;
  __u32 ring_entries;
  __u16 bgid;
  __u16 flags;
  __u64 resv[3];
};
struct io_uring_buf_status {
  __u32 buf_group;
  __u32 head;
  __u32 resv[8];
};
struct io_uring_napi {
  __u32 busy_poll_to;
  __u8 prefer_busy_poll;
  __u8 pad[3];
  __u64 resv;
};
enum io_uring_register_restriction_op {
  IORING_RESTRICTION_REGISTER_OP = 0,
  IORING_RESTRICTION_SQE_OP = 1,
  IORING_RESTRICTION_SQE_FLAGS_ALLOWED = 2,
  IORING_RESTRICTION_SQE_FLAGS_REQUIRED = 3,
  IORING_RESTRICTION_LAST
};
struct io_uring_getevents_arg {
  __u64 sigmask;
  __u32 sigmask_sz;
  __u32 min_wait_usec;
  __u64 ts;
};
struct io_uring_sync_cancel_reg {
  __u64 addr;
  __s32 fd;
  __u32 flags;
  struct __kernel_timespec timeout;
  __u8 opcode;
  __u8 pad[7];
  __u64 pad2[3];
};
struct io_uring_file_index_range {
  __u32 off;
  __u32 len;
  __u64 resv;
};
struct io_uring_recvmsg_out {
  __u32 namelen;
  __u32 controllen;
  __u32 payloadlen;
  __u32 flags;
};
enum io_uring_socket_op {
  SOCKET_URING_OP_SIOCINQ = 0,
  SOCKET_URING_OP_SIOCOUTQ,
  SOCKET_URING_OP_GETSOCKOPT,
  SOCKET_URING_OP_SETSOCKOPT,
};
#ifdef __cplusplus
}
#endif
#endif
