/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _USERIO_H
#define _USERIO_H
#include <linux/types.h>
enum userio_cmd_type {
  USERIO_CMD_REGISTER = 0,
  USERIO_CMD_SET_PORT_TYPE = 1,
  USERIO_CMD_SEND_INTERRUPT = 2
};
struct userio_cmd {
  __u8 type;
  __u8 data;
} __attribute__((__packed__));
#endif
