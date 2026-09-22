/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_CAN_H
#define _UAPI_CAN_H
#include <linux/types.h>
#include <linux/socket.h>
#include <linux/stddef.h>
#define CAN_EFF_FLAG 0x80000000U
#define CAN_RTR_FLAG 0x40000000U
#define CAN_ERR_FLAG 0x20000000U
#define CAN_SFF_MASK 0x000007FFU
#define CAN_EFF_MASK 0x1FFFFFFFU
#define CAN_ERR_MASK 0x1FFFFFFFU
#define CANXL_PRIO_MASK CAN_SFF_MASK
typedef __u32 canid_t;
#define CAN_SFF_ID_BITS 11
#define CAN_EFF_ID_BITS 29
#define CANXL_PRIO_BITS CAN_SFF_ID_BITS
typedef __u32 can_err_mask_t;
#define CAN_MAX_DLC 8
#define CAN_MAX_RAW_DLC 15
#define CAN_MAX_DLEN 8
#define CANFD_MAX_DLC 15
#define CANFD_MAX_DLEN 64
#define CANXL_MIN_DLC 0
#define CANXL_MAX_DLC 2047
#define CANXL_MAX_DLC_MASK 0x07FF
#define CANXL_MIN_DLEN 1
#define CANXL_MAX_DLEN 2048
struct can_frame {
  canid_t can_id;
  union {
    __u8 len;
    __u8 can_dlc;
  } __attribute__((packed));
  __u8 __pad;
  __u8 __res0;
  __u8 len8_dlc;
  __u8 data[CAN_MAX_DLEN] __attribute__((aligned(8)));
};
#define CANFD_BRS 0x01
#define CANFD_ESI 0x02
#define CANFD_FDF 0x04
struct canfd_frame {
  canid_t can_id;
  __u8 len;
  __u8 flags;
  __u8 __res0;
  __u8 __res1;
  __u8 data[CANFD_MAX_DLEN] __attribute__((aligned(8)));
};
#define CANXL_XLF 0x80
#define CANXL_SEC 0x01
#define CANXL_VCID_OFFSET 16
#define CANXL_VCID_VAL_MASK 0xFFUL
#define CANXL_VCID_MASK (CANXL_VCID_VAL_MASK << CANXL_VCID_OFFSET)
struct canxl_frame {
  canid_t prio;
  __u8 flags;
  __u8 sdt;
  __u16 len;
  __u32 af;
  __u8 data[CANXL_MAX_DLEN];
};
#define CAN_MTU (sizeof(struct can_frame))
#define CANFD_MTU (sizeof(struct canfd_frame))
#define CANXL_MTU (sizeof(struct canxl_frame))
#define CANXL_HDR_SIZE (offsetof(struct canxl_frame, data))
#define CANXL_MIN_MTU (CANXL_HDR_SIZE + 64)
#define CANXL_MAX_MTU CANXL_MTU
#define CAN_RAW 1
#define CAN_BCM 2
#define CAN_TP16 3
#define CAN_TP20 4
#define CAN_MCNET 5
#define CAN_ISOTP 6
#define CAN_J1939 7
#define CAN_NPROTO 8
#define SOL_CAN_BASE 100
struct sockaddr_can {
  __kernel_sa_family_t can_family;
  int can_ifindex;
  union {
    struct {
      canid_t rx_id, tx_id;
    } tp;
    struct {
      __u64 name;
      __u32 pgn;
      __u8 addr;
    } j1939;
  } can_addr;
};
struct can_filter {
  canid_t can_id;
  canid_t can_mask;
};
#define CAN_INV_FILTER 0x20000000U
#endif
