/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPINFS_IDMAP_H
#define _UAPINFS_IDMAP_H
#include <linux/types.h>
#define IDMAP_NAMESZ 128
#define IDMAP_TYPE_USER 0
#define IDMAP_TYPE_GROUP 1
#define IDMAP_CONV_IDTONAME 0
#define IDMAP_CONV_NAMETOID 1
#define IDMAP_STATUS_INVALIDMSG 0x01
#define IDMAP_STATUS_AGAIN 0x02
#define IDMAP_STATUS_LOOKUPFAIL 0x04
#define IDMAP_STATUS_SUCCESS 0x08
struct idmap_msg {
  __u8 im_type;
  __u8 im_conv;
  char im_name[IDMAP_NAMESZ];
  __u32 im_id;
  __u8 im_status;
};
#endif
