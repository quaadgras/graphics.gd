/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_MOUNT_H
#define _UAPI_LINUX_MOUNT_H
#include <linux/types.h>
#define MS_RDONLY 1
#define MS_NOSUID 2
#define MS_NODEV 4
#define MS_NOEXEC 8
#define MS_SYNCHRONOUS 16
#define MS_REMOUNT 32
#define MS_MANDLOCK 64
#define MS_DIRSYNC 128
#define MS_NOSYMFOLLOW 256
#define MS_NOATIME 1024
#define MS_NODIRATIME 2048
#define MS_BIND 4096
#define MS_MOVE 8192
#define MS_REC 16384
#define MS_VERBOSE 32768
#define MS_SILENT 32768
#define MS_POSIXACL (1 << 16)
#define MS_UNBINDABLE (1 << 17)
#define MS_PRIVATE (1 << 18)
#define MS_SLAVE (1 << 19)
#define MS_SHARED (1 << 20)
#define MS_RELATIME (1 << 21)
#define MS_KERNMOUNT (1 << 22)
#define MS_I_VERSION (1 << 23)
#define MS_STRICTATIME (1 << 24)
#define MS_LAZYTIME (1 << 25)
#define MS_SUBMOUNT (1 << 26)
#define MS_NOREMOTELOCK (1 << 27)
#define MS_NOSEC (1 << 28)
#define MS_BORN (1 << 29)
#define MS_ACTIVE (1 << 30)
#define MS_NOUSER (1 << 31)
#define MS_RMT_MASK (MS_RDONLY | MS_SYNCHRONOUS | MS_MANDLOCK | MS_I_VERSION | MS_LAZYTIME)
#define MS_MGC_VAL 0xC0ED0000
#define MS_MGC_MSK 0xffff0000
#define OPEN_TREE_CLONE 1
#define OPEN_TREE_CLOEXEC O_CLOEXEC
#define MOVE_MOUNT_F_SYMLINKS 0x00000001
#define MOVE_MOUNT_F_AUTOMOUNTS 0x00000002
#define MOVE_MOUNT_F_EMPTY_PATH 0x00000004
#define MOVE_MOUNT_T_SYMLINKS 0x00000010
#define MOVE_MOUNT_T_AUTOMOUNTS 0x00000020
#define MOVE_MOUNT_T_EMPTY_PATH 0x00000040
#define MOVE_MOUNT_SET_GROUP 0x00000100
#define MOVE_MOUNT_BENEATH 0x00000200
#define MOVE_MOUNT__MASK 0x00000377
#define FSOPEN_CLOEXEC 0x00000001
#define FSPICK_CLOEXEC 0x00000001
#define FSPICK_SYMLINK_NOFOLLOW 0x00000002
#define FSPICK_NO_AUTOMOUNT 0x00000004
#define FSPICK_EMPTY_PATH 0x00000008
enum fsconfig_command {
  FSCONFIG_SET_FLAG = 0,
  FSCONFIG_SET_STRING = 1,
  FSCONFIG_SET_BINARY = 2,
  FSCONFIG_SET_PATH = 3,
  FSCONFIG_SET_PATH_EMPTY = 4,
  FSCONFIG_SET_FD = 5,
  FSCONFIG_CMD_CREATE = 6,
  FSCONFIG_CMD_RECONFIGURE = 7,
  FSCONFIG_CMD_CREATE_EXCL = 8,
};
#define FSMOUNT_CLOEXEC 0x00000001
#define MOUNT_ATTR_RDONLY 0x00000001
#define MOUNT_ATTR_NOSUID 0x00000002
#define MOUNT_ATTR_NODEV 0x00000004
#define MOUNT_ATTR_NOEXEC 0x00000008
#define MOUNT_ATTR__ATIME 0x00000070
#define MOUNT_ATTR_RELATIME 0x00000000
#define MOUNT_ATTR_NOATIME 0x00000010
#define MOUNT_ATTR_STRICTATIME 0x00000020
#define MOUNT_ATTR_NODIRATIME 0x00000080
#define MOUNT_ATTR_IDMAP 0x00100000
#define MOUNT_ATTR_NOSYMFOLLOW 0x00200000
struct mount_attr {
  __u64 attr_set;
  __u64 attr_clr;
  __u64 propagation;
  __u64 userns_fd;
};
#define MOUNT_ATTR_SIZE_VER0 32
struct statmount {
  __u32 size;
  __u32 mnt_opts;
  __u64 mask;
  __u32 sb_dev_major;
  __u32 sb_dev_minor;
  __u64 sb_magic;
  __u32 sb_flags;
  __u32 fs_type;
  __u64 mnt_id;
  __u64 mnt_parent_id;
  __u32 mnt_id_old;
  __u32 mnt_parent_id_old;
  __u64 mnt_attr;
  __u64 mnt_propagation;
  __u64 mnt_peer_group;
  __u64 mnt_master;
  __u64 propagate_from;
  __u32 mnt_root;
  __u32 mnt_point;
  __u64 mnt_ns_id;
  __u64 __spare2[49];
  char str[];
};
struct mnt_id_req {
  __u32 size;
  __u32 spare;
  __u64 mnt_id;
  __u64 param;
  __u64 mnt_ns_id;
};
#define MNT_ID_REQ_SIZE_VER0 24
#define MNT_ID_REQ_SIZE_VER1 32
#define STATMOUNT_SB_BASIC 0x00000001U
#define STATMOUNT_MNT_BASIC 0x00000002U
#define STATMOUNT_PROPAGATE_FROM 0x00000004U
#define STATMOUNT_MNT_ROOT 0x00000008U
#define STATMOUNT_MNT_POINT 0x00000010U
#define STATMOUNT_FS_TYPE 0x00000020U
#define STATMOUNT_MNT_NS_ID 0x00000040U
#define STATMOUNT_MNT_OPTS 0x00000080U
#define LSMT_ROOT 0xffffffffffffffff
#define LISTMOUNT_REVERSE (1 << 0)
#endif
