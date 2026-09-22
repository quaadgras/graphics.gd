/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef USER_BLK_DRV_CMD_INC_H
#define USER_BLK_DRV_CMD_INC_H
#include <linux/types.h>
#define UBLK_CMD_GET_QUEUE_AFFINITY 0x01
#define UBLK_CMD_GET_DEV_INFO 0x02
#define UBLK_CMD_ADD_DEV 0x04
#define UBLK_CMD_DEL_DEV 0x05
#define UBLK_CMD_START_DEV 0x06
#define UBLK_CMD_STOP_DEV 0x07
#define UBLK_CMD_SET_PARAMS 0x08
#define UBLK_CMD_GET_PARAMS 0x09
#define UBLK_CMD_START_USER_RECOVERY 0x10
#define UBLK_CMD_END_USER_RECOVERY 0x11
#define UBLK_CMD_GET_DEV_INFO2 0x12
#define UBLK_U_CMD_GET_QUEUE_AFFINITY _IOR('u', UBLK_CMD_GET_QUEUE_AFFINITY, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_GET_DEV_INFO _IOR('u', UBLK_CMD_GET_DEV_INFO, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_ADD_DEV _IOWR('u', UBLK_CMD_ADD_DEV, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_DEL_DEV _IOWR('u', UBLK_CMD_DEL_DEV, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_START_DEV _IOWR('u', UBLK_CMD_START_DEV, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_STOP_DEV _IOWR('u', UBLK_CMD_STOP_DEV, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_SET_PARAMS _IOWR('u', UBLK_CMD_SET_PARAMS, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_GET_PARAMS _IOR('u', UBLK_CMD_GET_PARAMS, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_START_USER_RECOVERY _IOWR('u', UBLK_CMD_START_USER_RECOVERY, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_END_USER_RECOVERY _IOWR('u', UBLK_CMD_END_USER_RECOVERY, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_GET_DEV_INFO2 _IOR('u', UBLK_CMD_GET_DEV_INFO2, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_GET_FEATURES _IOR('u', 0x13, struct ublksrv_ctrl_cmd)
#define UBLK_U_CMD_DEL_DEV_ASYNC _IOR('u', 0x14, struct ublksrv_ctrl_cmd)
#define UBLK_FEATURES_LEN 8
#define UBLK_IO_FETCH_REQ 0x20
#define UBLK_IO_COMMIT_AND_FETCH_REQ 0x21
#define UBLK_IO_NEED_GET_DATA 0x22
#define UBLK_U_IO_FETCH_REQ _IOWR('u', UBLK_IO_FETCH_REQ, struct ublksrv_io_cmd)
#define UBLK_U_IO_COMMIT_AND_FETCH_REQ _IOWR('u', UBLK_IO_COMMIT_AND_FETCH_REQ, struct ublksrv_io_cmd)
#define UBLK_U_IO_NEED_GET_DATA _IOWR('u', UBLK_IO_NEED_GET_DATA, struct ublksrv_io_cmd)
#define UBLK_IO_RES_OK 0
#define UBLK_IO_RES_NEED_GET_DATA 1
#define UBLK_IO_RES_ABORT (- ENODEV)
#define UBLKSRV_CMD_BUF_OFFSET 0
#define UBLKSRV_IO_BUF_OFFSET 0x80000000
#define UBLK_MAX_QUEUE_DEPTH 4096
#define UBLK_IO_BUF_OFF 0
#define UBLK_IO_BUF_BITS 25
#define UBLK_IO_BUF_BITS_MASK ((1ULL << UBLK_IO_BUF_BITS) - 1)
#define UBLK_TAG_OFF UBLK_IO_BUF_BITS
#define UBLK_TAG_BITS 16
#define UBLK_TAG_BITS_MASK ((1ULL << UBLK_TAG_BITS) - 1)
#define UBLK_QID_OFF (UBLK_TAG_OFF + UBLK_TAG_BITS)
#define UBLK_QID_BITS 12
#define UBLK_QID_BITS_MASK ((1ULL << UBLK_QID_BITS) - 1)
#define UBLK_MAX_NR_QUEUES (1U << UBLK_QID_BITS)
#define UBLKSRV_IO_BUF_TOTAL_BITS (UBLK_QID_OFF + UBLK_QID_BITS)
#define UBLKSRV_IO_BUF_TOTAL_SIZE (1ULL << UBLKSRV_IO_BUF_TOTAL_BITS)
#define UBLK_F_SUPPORT_ZERO_COPY (1ULL << 0)
#define UBLK_F_URING_CMD_COMP_IN_TASK (1ULL << 1)
#define UBLK_F_NEED_GET_DATA (1UL << 2)
#define UBLK_F_USER_RECOVERY (1UL << 3)
#define UBLK_F_USER_RECOVERY_REISSUE (1UL << 4)
#define UBLK_F_UNPRIVILEGED_DEV (1UL << 5)
#define UBLK_F_CMD_IOCTL_ENCODE (1UL << 6)
#define UBLK_F_USER_COPY (1UL << 7)
#define UBLK_F_ZONED (1ULL << 8)
#define UBLK_S_DEV_DEAD 0
#define UBLK_S_DEV_LIVE 1
#define UBLK_S_DEV_QUIESCED 2
struct ublksrv_ctrl_cmd {
  __u32 dev_id;
  __u16 queue_id;
  __u16 len;
  __u64 addr;
  __u64 data[1];
  __u16 dev_path_len;
  __u16 pad;
  __u32 reserved;
};
struct ublksrv_ctrl_dev_info {
  __u16 nr_hw_queues;
  __u16 queue_depth;
  __u16 state;
  __u16 pad0;
  __u32 max_io_buf_bytes;
  __u32 dev_id;
  __s32 ublksrv_pid;
  __u32 pad1;
  __u64 flags;
  __u64 ublksrv_flags;
  __u32 owner_uid;
  __u32 owner_gid;
  __u64 reserved1;
  __u64 reserved2;
};
#define UBLK_IO_OP_READ 0
#define UBLK_IO_OP_WRITE 1
#define UBLK_IO_OP_FLUSH 2
#define UBLK_IO_OP_DISCARD 3
#define UBLK_IO_OP_WRITE_SAME 4
#define UBLK_IO_OP_WRITE_ZEROES 5
#define UBLK_IO_OP_ZONE_OPEN 10
#define UBLK_IO_OP_ZONE_CLOSE 11
#define UBLK_IO_OP_ZONE_FINISH 12
#define UBLK_IO_OP_ZONE_APPEND 13
#define UBLK_IO_OP_ZONE_RESET_ALL 14
#define UBLK_IO_OP_ZONE_RESET 15
#define UBLK_IO_OP_REPORT_ZONES 18
#define UBLK_IO_F_FAILFAST_DEV (1U << 8)
#define UBLK_IO_F_FAILFAST_TRANSPORT (1U << 9)
#define UBLK_IO_F_FAILFAST_DRIVER (1U << 10)
#define UBLK_IO_F_META (1U << 11)
#define UBLK_IO_F_FUA (1U << 13)
#define UBLK_IO_F_NOUNMAP (1U << 15)
#define UBLK_IO_F_SWAP (1U << 16)
struct ublksrv_io_desc {
  __u32 op_flags;
  union {
    __u32 nr_sectors;
    __u32 nr_zones;
  };
  __u64 start_sector;
  __u64 addr;
};
struct ublksrv_io_cmd {
  __u16 q_id;
  __u16 tag;
  __s32 result;
  union {
    __u64 addr;
    __u64 zone_append_lba;
  };
};
struct ublk_param_basic {
#define UBLK_ATTR_READ_ONLY (1 << 0)
#define UBLK_ATTR_ROTATIONAL (1 << 1)
#define UBLK_ATTR_VOLATILE_CACHE (1 << 2)
#define UBLK_ATTR_FUA (1 << 3)
  __u32 attrs;
  __u8 logical_bs_shift;
  __u8 physical_bs_shift;
  __u8 io_opt_shift;
  __u8 io_min_shift;
  __u32 max_sectors;
  __u32 chunk_sectors;
  __u64 dev_sectors;
  __u64 virt_boundary_mask;
};
struct ublk_param_discard {
  __u32 discard_alignment;
  __u32 discard_granularity;
  __u32 max_discard_sectors;
  __u32 max_write_zeroes_sectors;
  __u16 max_discard_segments;
  __u16 reserved0;
};
struct ublk_param_devt {
  __u32 char_major;
  __u32 char_minor;
  __u32 disk_major;
  __u32 disk_minor;
};
struct ublk_param_zoned {
  __u32 max_open_zones;
  __u32 max_active_zones;
  __u32 max_zone_append_sectors;
  __u8 reserved[20];
};
struct ublk_params {
  __u32 len;
#define UBLK_PARAM_TYPE_BASIC (1 << 0)
#define UBLK_PARAM_TYPE_DISCARD (1 << 1)
#define UBLK_PARAM_TYPE_DEVT (1 << 2)
#define UBLK_PARAM_TYPE_ZONED (1 << 3)
  __u32 types;
  struct ublk_param_basic basic;
  struct ublk_param_discard discard;
  struct ublk_param_devt devt;
  struct ublk_param_zoned zoned;
};
#endif
