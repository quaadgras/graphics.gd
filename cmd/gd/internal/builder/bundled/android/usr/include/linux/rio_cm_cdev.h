/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _RIO_CM_CDEV_H_
#define _RIO_CM_CDEV_H_
#include <linux/types.h>
struct rio_cm_channel {
  __u16 id;
  __u16 remote_channel;
  __u16 remote_destid;
  __u8 mport_id;
};
struct rio_cm_msg {
  __u16 ch_num;
  __u16 size;
  __u32 rxto;
  __u64 msg;
};
struct rio_cm_accept {
  __u16 ch_num;
  __u16 pad0;
  __u32 wait_to;
};
#define RIO_CM_IOC_MAGIC 'c'
#define RIO_CM_EP_GET_LIST_SIZE _IOWR(RIO_CM_IOC_MAGIC, 1, __u32)
#define RIO_CM_EP_GET_LIST _IOWR(RIO_CM_IOC_MAGIC, 2, __u32)
#define RIO_CM_CHAN_CREATE _IOWR(RIO_CM_IOC_MAGIC, 3, __u16)
#define RIO_CM_CHAN_CLOSE _IOW(RIO_CM_IOC_MAGIC, 4, __u16)
#define RIO_CM_CHAN_BIND _IOW(RIO_CM_IOC_MAGIC, 5, struct rio_cm_channel)
#define RIO_CM_CHAN_LISTEN _IOW(RIO_CM_IOC_MAGIC, 6, __u16)
#define RIO_CM_CHAN_ACCEPT _IOWR(RIO_CM_IOC_MAGIC, 7, struct rio_cm_accept)
#define RIO_CM_CHAN_CONNECT _IOW(RIO_CM_IOC_MAGIC, 8, struct rio_cm_channel)
#define RIO_CM_CHAN_SEND _IOW(RIO_CM_IOC_MAGIC, 9, struct rio_cm_msg)
#define RIO_CM_CHAN_RECEIVE _IOWR(RIO_CM_IOC_MAGIC, 10, struct rio_cm_msg)
#define RIO_CM_MPORT_GET_LIST _IOWR(RIO_CM_IOC_MAGIC, 11, __u32)
#endif
