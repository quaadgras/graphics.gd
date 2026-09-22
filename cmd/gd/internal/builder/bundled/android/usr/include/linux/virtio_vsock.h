/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_VIRTIO_VSOCK_H
#define _UAPI_LINUX_VIRTIO_VSOCK_H
#include <linux/types.h>
#include <linux/virtio_ids.h>
#include <linux/virtio_config.h>
#define VIRTIO_VSOCK_F_SEQPACKET 1
struct virtio_vsock_config {
  __le64 guest_cid;
} __attribute__((packed));
enum virtio_vsock_event_id {
  VIRTIO_VSOCK_EVENT_TRANSPORT_RESET = 0,
};
struct virtio_vsock_event {
  __le32 id;
} __attribute__((packed));
struct virtio_vsock_hdr {
  __le64 src_cid;
  __le64 dst_cid;
  __le32 src_port;
  __le32 dst_port;
  __le32 len;
  __le16 type;
  __le16 op;
  __le32 flags;
  __le32 buf_alloc;
  __le32 fwd_cnt;
} __attribute__((packed));
enum virtio_vsock_type {
  VIRTIO_VSOCK_TYPE_STREAM = 1,
  VIRTIO_VSOCK_TYPE_SEQPACKET = 2,
};
enum virtio_vsock_op {
  VIRTIO_VSOCK_OP_INVALID = 0,
  VIRTIO_VSOCK_OP_REQUEST = 1,
  VIRTIO_VSOCK_OP_RESPONSE = 2,
  VIRTIO_VSOCK_OP_RST = 3,
  VIRTIO_VSOCK_OP_SHUTDOWN = 4,
  VIRTIO_VSOCK_OP_RW = 5,
  VIRTIO_VSOCK_OP_CREDIT_UPDATE = 6,
  VIRTIO_VSOCK_OP_CREDIT_REQUEST = 7,
};
enum virtio_vsock_shutdown {
  VIRTIO_VSOCK_SHUTDOWN_RCV = 1,
  VIRTIO_VSOCK_SHUTDOWN_SEND = 2,
};
enum virtio_vsock_rw {
  VIRTIO_VSOCK_SEQ_EOM = 1,
  VIRTIO_VSOCK_SEQ_EOR = 2,
};
#endif
