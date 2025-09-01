/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_SOCKET_H
#define _UAPI_LINUX_SOCKET_H
#include <bits/sockaddr_storage.h>
#define _K_SS_MAXSIZE 128
typedef unsigned short __kernel_sa_family_t;
#define SOCK_SNDBUF_LOCK 1
#define SOCK_RCVBUF_LOCK 2
#define SOCK_BUF_LOCK_MASK (SOCK_SNDBUF_LOCK | SOCK_RCVBUF_LOCK)
#define SOCK_TXREHASH_DEFAULT 255
#define SOCK_TXREHASH_DISABLED 0
#define SOCK_TXREHASH_ENABLED 1
#endif
