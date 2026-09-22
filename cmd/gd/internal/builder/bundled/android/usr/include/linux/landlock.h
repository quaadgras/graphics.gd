/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_LANDLOCK_H
#define _UAPI_LINUX_LANDLOCK_H
#include <linux/types.h>
struct landlock_ruleset_attr {
  __u64 handled_access_fs;
  __u64 handled_access_net;
  __u64 scoped;
};
#define LANDLOCK_CREATE_RULESET_VERSION (1U << 0)
enum landlock_rule_type {
  LANDLOCK_RULE_PATH_BENEATH = 1,
  LANDLOCK_RULE_NET_PORT,
};
struct landlock_path_beneath_attr {
  __u64 allowed_access;
  __s32 parent_fd;
} __attribute__((packed));
struct landlock_net_port_attr {
  __u64 allowed_access;
  __u64 port;
};
#define LANDLOCK_ACCESS_FS_EXECUTE (1ULL << 0)
#define LANDLOCK_ACCESS_FS_WRITE_FILE (1ULL << 1)
#define LANDLOCK_ACCESS_FS_READ_FILE (1ULL << 2)
#define LANDLOCK_ACCESS_FS_READ_DIR (1ULL << 3)
#define LANDLOCK_ACCESS_FS_REMOVE_DIR (1ULL << 4)
#define LANDLOCK_ACCESS_FS_REMOVE_FILE (1ULL << 5)
#define LANDLOCK_ACCESS_FS_MAKE_CHAR (1ULL << 6)
#define LANDLOCK_ACCESS_FS_MAKE_DIR (1ULL << 7)
#define LANDLOCK_ACCESS_FS_MAKE_REG (1ULL << 8)
#define LANDLOCK_ACCESS_FS_MAKE_SOCK (1ULL << 9)
#define LANDLOCK_ACCESS_FS_MAKE_FIFO (1ULL << 10)
#define LANDLOCK_ACCESS_FS_MAKE_BLOCK (1ULL << 11)
#define LANDLOCK_ACCESS_FS_MAKE_SYM (1ULL << 12)
#define LANDLOCK_ACCESS_FS_REFER (1ULL << 13)
#define LANDLOCK_ACCESS_FS_TRUNCATE (1ULL << 14)
#define LANDLOCK_ACCESS_FS_IOCTL_DEV (1ULL << 15)
#define LANDLOCK_ACCESS_NET_BIND_TCP (1ULL << 0)
#define LANDLOCK_ACCESS_NET_CONNECT_TCP (1ULL << 1)
#define LANDLOCK_SCOPE_ABSTRACT_UNIX_SOCKET (1ULL << 0)
#define LANDLOCK_SCOPE_SIGNAL (1ULL << 1)
#endif
