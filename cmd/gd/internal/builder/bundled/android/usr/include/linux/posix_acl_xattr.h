/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __UAPI_POSIX_ACL_XATTR_H
#define __UAPI_POSIX_ACL_XATTR_H
#include <linux/types.h>
#define POSIX_ACL_XATTR_VERSION 0x0002
#define ACL_UNDEFINED_ID (- 1)
struct posix_acl_xattr_entry {
  __le16 e_tag;
  __le16 e_perm;
  __le32 e_id;
};
struct posix_acl_xattr_header {
  __le32 a_version;
};
#endif
