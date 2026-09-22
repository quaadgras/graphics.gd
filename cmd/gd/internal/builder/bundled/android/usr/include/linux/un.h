/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_UN_H
#define _LINUX_UN_H
#include <linux/socket.h>
#define UNIX_PATH_MAX 108
struct sockaddr_un {
  __kernel_sa_family_t sun_family;
  char sun_path[UNIX_PATH_MAX];
};
#define SIOCUNIXFILE (SIOCPROTOPRIVATE + 0)
#endif
