/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_NFS4_MOUNT_H
#define _LINUX_NFS4_MOUNT_H
#define NFS4_MOUNT_VERSION 1
struct nfs_string {
  unsigned int len;
  const char  * data;
};
struct nfs4_mount_data {
  int version;
  int flags;
  int rsize;
  int wsize;
  int timeo;
  int retrans;
  int acregmin;
  int acregmax;
  int acdirmin;
  int acdirmax;
  struct nfs_string client_addr;
  struct nfs_string mnt_path;
  struct nfs_string hostname;
  unsigned int host_addrlen;
  struct sockaddr  * host_addr;
  int proto;
  int auth_flavourlen;
  int  * auth_flavours;
};
#define NFS4_MOUNT_SOFT 0x0001
#define NFS4_MOUNT_INTR 0x0002
#define NFS4_MOUNT_NOCTO 0x0010
#define NFS4_MOUNT_NOAC 0x0020
#define NFS4_MOUNT_STRICTLOCK 0x1000
#define NFS4_MOUNT_UNSHARED 0x8000
#define NFS4_MOUNT_FLAGMASK 0x9033
#endif
