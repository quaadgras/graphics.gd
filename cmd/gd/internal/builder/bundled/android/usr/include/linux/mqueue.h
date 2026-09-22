/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_MQUEUE_H
#define _LINUX_MQUEUE_H
#include <linux/types.h>
#define MQ_PRIO_MAX 32768
#define MQ_BYTES_MAX 819200
struct mq_attr {
  __kernel_long_t mq_flags;
  __kernel_long_t mq_maxmsg;
  __kernel_long_t mq_msgsize;
  __kernel_long_t mq_curmsgs;
  __kernel_long_t __reserved[4];
};
#define NOTIFY_NONE 0
#define NOTIFY_WOKENUP 1
#define NOTIFY_REMOVED 2
#define NOTIFY_COOKIE_LEN 32
#endif
