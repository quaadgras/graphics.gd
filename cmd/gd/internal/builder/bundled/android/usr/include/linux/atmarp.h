/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_ATMARP_H
#define _LINUX_ATMARP_H
#include <linux/types.h>
#include <linux/atmapi.h>
#include <linux/atmioc.h>
#define ATMARP_RETRY_DELAY 30
#define ATMARP_MAX_UNRES_PACKETS 5
#define ATMARPD_CTRL _IO('a', ATMIOC_CLIP + 1)
#define ATMARP_MKIP _IO('a', ATMIOC_CLIP + 2)
#define ATMARP_SETENTRY _IO('a', ATMIOC_CLIP + 3)
#define ATMARP_ENCAP _IO('a', ATMIOC_CLIP + 5)
enum atmarp_ctrl_type {
  act_invalid,
  act_need,
  act_up,
  act_down,
  act_change
};
struct atmarp_ctrl {
  enum atmarp_ctrl_type type;
  int itf_num;
  __be32 ip;
};
#endif
