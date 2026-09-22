/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __TARGET_CORE_USER_H
#define __TARGET_CORE_USER_H
#include <linux/types.h>
#include <linux/uio.h>
#define TCMU_VERSION "2.0"
#define TCMU_MAILBOX_VERSION 2
#define ALIGN_SIZE 64
#define TCMU_MAILBOX_FLAG_CAP_OOOC (1 << 0)
#define TCMU_MAILBOX_FLAG_CAP_READ_LEN (1 << 1)
#define TCMU_MAILBOX_FLAG_CAP_TMR (1 << 2)
#define TCMU_MAILBOX_FLAG_CAP_KEEP_BUF (1 << 3)
struct tcmu_mailbox {
  __u16 version;
  __u16 flags;
  __u32 cmdr_off;
  __u32 cmdr_size;
  __u32 cmd_head;
  __u32 cmd_tail __attribute__((__aligned__(ALIGN_SIZE)));
} __attribute__((__packed__));
enum tcmu_opcode {
  TCMU_OP_PAD = 0,
  TCMU_OP_CMD,
  TCMU_OP_TMR,
};
struct tcmu_cmd_entry_hdr {
  __u32 len_op;
  __u16 cmd_id;
  __u8 kflags;
#define TCMU_UFLAG_UNKNOWN_OP 0x1
#define TCMU_UFLAG_READ_LEN 0x2
#define TCMU_UFLAG_KEEP_BUF 0x4
  __u8 uflags;
} __attribute__((__packed__));
#define TCMU_OP_MASK 0x7
#define TCMU_SENSE_BUFFERSIZE 96
struct tcmu_cmd_entry {
  struct tcmu_cmd_entry_hdr hdr;
  union {
    struct {
      __u32 iov_cnt;
      __u32 iov_bidi_cnt;
      __u32 iov_dif_cnt;
      __u64 cdb_off;
      __u64 __pad1;
      __u64 __pad2;
      __DECLARE_FLEX_ARRAY(struct iovec, iov);
    } req;
    struct {
      __u8 scsi_status;
      __u8 __pad1;
      __u16 __pad2;
      __u32 read_len;
      char sense_buffer[TCMU_SENSE_BUFFERSIZE];
    } rsp;
  };
} __attribute__((__packed__));
struct tcmu_tmr_entry {
  struct tcmu_cmd_entry_hdr hdr;
#define TCMU_TMR_UNKNOWN 0
#define TCMU_TMR_ABORT_TASK 1
#define TCMU_TMR_ABORT_TASK_SET 2
#define TCMU_TMR_CLEAR_ACA 3
#define TCMU_TMR_CLEAR_TASK_SET 4
#define TCMU_TMR_LUN_RESET 5
#define TCMU_TMR_TARGET_WARM_RESET 6
#define TCMU_TMR_TARGET_COLD_RESET 7
#define TCMU_TMR_LUN_RESET_PRO 128
  __u8 tmr_type;
  __u8 __pad1;
  __u16 __pad2;
  __u32 cmd_cnt;
  __u64 __pad3;
  __u64 __pad4;
  __u16 cmd_ids[];
} __attribute__((__packed__));
#define TCMU_OP_ALIGN_SIZE sizeof(__u64)
enum tcmu_genl_cmd {
  TCMU_CMD_UNSPEC,
  TCMU_CMD_ADDED_DEVICE,
  TCMU_CMD_REMOVED_DEVICE,
  TCMU_CMD_RECONFIG_DEVICE,
  TCMU_CMD_ADDED_DEVICE_DONE,
  TCMU_CMD_REMOVED_DEVICE_DONE,
  TCMU_CMD_RECONFIG_DEVICE_DONE,
  TCMU_CMD_SET_FEATURES,
  __TCMU_CMD_MAX,
};
#define TCMU_CMD_MAX (__TCMU_CMD_MAX - 1)
enum tcmu_genl_attr {
  TCMU_ATTR_UNSPEC,
  TCMU_ATTR_DEVICE,
  TCMU_ATTR_MINOR,
  TCMU_ATTR_PAD,
  TCMU_ATTR_DEV_CFG,
  TCMU_ATTR_DEV_SIZE,
  TCMU_ATTR_WRITECACHE,
  TCMU_ATTR_CMD_STATUS,
  TCMU_ATTR_DEVICE_ID,
  TCMU_ATTR_SUPP_KERN_CMD_REPLY,
  __TCMU_ATTR_MAX,
};
#define TCMU_ATTR_MAX (__TCMU_ATTR_MAX - 1)
#endif
