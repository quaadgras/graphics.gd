/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_VM_SOCKETS_H
#define _UAPI_VM_SOCKETS_H
#include <linux/socket.h>
#include <linux/types.h>
#define SO_VM_SOCKETS_BUFFER_SIZE 0
#define SO_VM_SOCKETS_BUFFER_MIN_SIZE 1
#define SO_VM_SOCKETS_BUFFER_MAX_SIZE 2
#define SO_VM_SOCKETS_PEER_HOST_VM_ID 3
#define SO_VM_SOCKETS_TRUSTED 5
#define SO_VM_SOCKETS_CONNECT_TIMEOUT_OLD 6
#define SO_VM_SOCKETS_NONBLOCK_TXRX 7
#define SO_VM_SOCKETS_CONNECT_TIMEOUT_NEW 8
#if __BITS_PER_LONG == 64 || defined(__x86_64__) && defined(__ILP32__)
#define SO_VM_SOCKETS_CONNECT_TIMEOUT SO_VM_SOCKETS_CONNECT_TIMEOUT_OLD
#else
#define SO_VM_SOCKETS_CONNECT_TIMEOUT (sizeof(time_t) == sizeof(__kernel_long_t) ? SO_VM_SOCKETS_CONNECT_TIMEOUT_OLD : SO_VM_SOCKETS_CONNECT_TIMEOUT_NEW)
#endif
#define VMADDR_CID_ANY - 1U
#define VMADDR_PORT_ANY - 1U
#define VMADDR_CID_HYPERVISOR 0
#define VMADDR_CID_LOCAL 1
#define VMADDR_CID_HOST 2
#define VMADDR_FLAG_TO_HOST 0x01
#define VM_SOCKETS_INVALID_VERSION - 1U
#define VM_SOCKETS_VERSION_EPOCH(_v) (((_v) & 0xFF000000) >> 24)
#define VM_SOCKETS_VERSION_MAJOR(_v) (((_v) & 0x00FF0000) >> 16)
#define VM_SOCKETS_VERSION_MINOR(_v) (((_v) & 0x0000FFFF))
struct sockaddr_vm {
  __kernel_sa_family_t svm_family;
  unsigned short svm_reserved1;
  unsigned int svm_port;
  unsigned int svm_cid;
  __u8 svm_flags;
  unsigned char svm_zero[sizeof(struct sockaddr) - sizeof(sa_family_t) - sizeof(unsigned short) - sizeof(unsigned int) - sizeof(unsigned int) - sizeof(__u8)];
};
#define IOCTL_VM_SOCKETS_GET_LOCAL_CID _IO(7, 0xb9)
#define SOL_VSOCK 287
#define VSOCK_RECVERR 1
#endif
