/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_EVENTPOLL_H
#define _UAPI_LINUX_EVENTPOLL_H
#include <bits/epoll_event.h>
#include <linux/fcntl.h>
#include <linux/types.h>
#define EPOLL_CLOEXEC O_CLOEXEC
#define EPOLL_CTL_ADD 1
#define EPOLL_CTL_DEL 2
#define EPOLL_CTL_MOD 3
#define EPOLLIN ( __poll_t) 0x00000001
#define EPOLLPRI ( __poll_t) 0x00000002
#define EPOLLOUT ( __poll_t) 0x00000004
#define EPOLLERR ( __poll_t) 0x00000008
#define EPOLLHUP ( __poll_t) 0x00000010
#define EPOLLNVAL ( __poll_t) 0x00000020
#define EPOLLRDNORM ( __poll_t) 0x00000040
#define EPOLLRDBAND ( __poll_t) 0x00000080
#define EPOLLWRNORM ( __poll_t) 0x00000100
#define EPOLLWRBAND ( __poll_t) 0x00000200
#define EPOLLMSG ( __poll_t) 0x00000400
#define EPOLLRDHUP ( __poll_t) 0x00002000
#define EPOLL_URING_WAKE (( __poll_t) (1U << 27))
#define EPOLLEXCLUSIVE (( __poll_t) (1U << 28))
#define EPOLLWAKEUP (( __poll_t) (1U << 29))
#define EPOLLONESHOT (( __poll_t) (1U << 30))
#define EPOLLET (( __poll_t) (1U << 31))
#ifdef __x86_64__
#define EPOLL_PACKED __attribute__((packed))
#else
#define EPOLL_PACKED
#endif
struct epoll_params {
  __u32 busy_poll_usecs;
  __u16 busy_poll_budget;
  __u8 prefer_busy_poll;
  __u8 __pad;
};
#define EPOLL_IOC_TYPE 0x8A
#define EPIOCSPARAMS _IOW(EPOLL_IOC_TYPE, 0x01, struct epoll_params)
#define EPIOCGPARAMS _IOR(EPOLL_IOC_TYPE, 0x02, struct epoll_params)
#endif
