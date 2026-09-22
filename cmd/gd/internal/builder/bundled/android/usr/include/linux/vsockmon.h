/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_VSOCKMON_H
#define _UAPI_VSOCKMON_H
#include <linux/virtio_vsock.h>
struct af_vsockmon_hdr {
  __le64 src_cid;
  __le64 dst_cid;
  __le32 src_port;
  __le32 dst_port;
  __le16 op;
  __le16 transport;
  __le16 len;
  __u8 reserved[2];
};
enum af_vsockmon_op {
  AF_VSOCK_OP_UNKNOWN = 0,
  AF_VSOCK_OP_CONNECT = 1,
  AF_VSOCK_OP_DISCONNECT = 2,
  AF_VSOCK_OP_CONTROL = 3,
  AF_VSOCK_OP_PAYLOAD = 4,
};
enum af_vsockmon_transport {
  AF_VSOCK_TRANSPORT_UNKNOWN = 0,
  AF_VSOCK_TRANSPORT_NO_INFO = 1,
  AF_VSOCK_TRANSPORT_VIRTIO = 2,
};
#endif
