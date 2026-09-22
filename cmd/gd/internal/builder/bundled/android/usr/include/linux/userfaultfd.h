/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_USERFAULTFD_H
#define _LINUX_USERFAULTFD_H
#include <linux/types.h>
#define USERFAULTFD_IOC 0xAA
#define USERFAULTFD_IOC_NEW _IO(USERFAULTFD_IOC, 0x00)
#define UFFD_API ((__u64) 0xAA)
#define UFFD_API_REGISTER_MODES (UFFDIO_REGISTER_MODE_MISSING | UFFDIO_REGISTER_MODE_WP | UFFDIO_REGISTER_MODE_MINOR)
#define UFFD_API_FEATURES (UFFD_FEATURE_PAGEFAULT_FLAG_WP | UFFD_FEATURE_EVENT_FORK | UFFD_FEATURE_EVENT_REMAP | UFFD_FEATURE_EVENT_REMOVE | UFFD_FEATURE_EVENT_UNMAP | UFFD_FEATURE_MISSING_HUGETLBFS | UFFD_FEATURE_MISSING_SHMEM | UFFD_FEATURE_SIGBUS | UFFD_FEATURE_THREAD_ID | UFFD_FEATURE_MINOR_HUGETLBFS | UFFD_FEATURE_MINOR_SHMEM | UFFD_FEATURE_EXACT_ADDRESS | UFFD_FEATURE_WP_HUGETLBFS_SHMEM | UFFD_FEATURE_WP_UNPOPULATED | UFFD_FEATURE_POISON | UFFD_FEATURE_WP_ASYNC | UFFD_FEATURE_MOVE)
#define UFFD_API_IOCTLS ((__u64) 1 << _UFFDIO_REGISTER | (__u64) 1 << _UFFDIO_UNREGISTER | (__u64) 1 << _UFFDIO_API)
#define UFFD_API_RANGE_IOCTLS ((__u64) 1 << _UFFDIO_WAKE | (__u64) 1 << _UFFDIO_COPY | (__u64) 1 << _UFFDIO_ZEROPAGE | (__u64) 1 << _UFFDIO_MOVE | (__u64) 1 << _UFFDIO_WRITEPROTECT | (__u64) 1 << _UFFDIO_CONTINUE | (__u64) 1 << _UFFDIO_POISON)
#define UFFD_API_RANGE_IOCTLS_BASIC ((__u64) 1 << _UFFDIO_WAKE | (__u64) 1 << _UFFDIO_COPY | (__u64) 1 << _UFFDIO_WRITEPROTECT | (__u64) 1 << _UFFDIO_CONTINUE | (__u64) 1 << _UFFDIO_POISON)
#define _UFFDIO_REGISTER (0x00)
#define _UFFDIO_UNREGISTER (0x01)
#define _UFFDIO_WAKE (0x02)
#define _UFFDIO_COPY (0x03)
#define _UFFDIO_ZEROPAGE (0x04)
#define _UFFDIO_MOVE (0x05)
#define _UFFDIO_WRITEPROTECT (0x06)
#define _UFFDIO_CONTINUE (0x07)
#define _UFFDIO_POISON (0x08)
#define _UFFDIO_API (0x3F)
#define UFFDIO 0xAA
#define UFFDIO_API _IOWR(UFFDIO, _UFFDIO_API, struct uffdio_api)
#define UFFDIO_REGISTER _IOWR(UFFDIO, _UFFDIO_REGISTER, struct uffdio_register)
#define UFFDIO_UNREGISTER _IOR(UFFDIO, _UFFDIO_UNREGISTER, struct uffdio_range)
#define UFFDIO_WAKE _IOR(UFFDIO, _UFFDIO_WAKE, struct uffdio_range)
#define UFFDIO_COPY _IOWR(UFFDIO, _UFFDIO_COPY, struct uffdio_copy)
#define UFFDIO_ZEROPAGE _IOWR(UFFDIO, _UFFDIO_ZEROPAGE, struct uffdio_zeropage)
#define UFFDIO_MOVE _IOWR(UFFDIO, _UFFDIO_MOVE, struct uffdio_move)
#define UFFDIO_WRITEPROTECT _IOWR(UFFDIO, _UFFDIO_WRITEPROTECT, struct uffdio_writeprotect)
#define UFFDIO_CONTINUE _IOWR(UFFDIO, _UFFDIO_CONTINUE, struct uffdio_continue)
#define UFFDIO_POISON _IOWR(UFFDIO, _UFFDIO_POISON, struct uffdio_poison)
struct uffd_msg {
  __u8 event;
  __u8 reserved1;
  __u16 reserved2;
  __u32 reserved3;
  union {
    struct {
      __u64 flags;
      __u64 address;
      union {
        __u32 ptid;
      } feat;
    } pagefault;
    struct {
      __u32 ufd;
    } fork;
    struct {
      __u64 from;
      __u64 to;
      __u64 len;
    } remap;
    struct {
      __u64 start;
      __u64 end;
    } remove;
    struct {
      __u64 reserved1;
      __u64 reserved2;
      __u64 reserved3;
    } reserved;
  } arg;
} __attribute__((__packed__));
#define UFFD_EVENT_PAGEFAULT 0x12
#define UFFD_EVENT_FORK 0x13
#define UFFD_EVENT_REMAP 0x14
#define UFFD_EVENT_REMOVE 0x15
#define UFFD_EVENT_UNMAP 0x16
#define UFFD_PAGEFAULT_FLAG_WRITE (1 << 0)
#define UFFD_PAGEFAULT_FLAG_WP (1 << 1)
#define UFFD_PAGEFAULT_FLAG_MINOR (1 << 2)
struct uffdio_api {
  __u64 api;
#define UFFD_FEATURE_PAGEFAULT_FLAG_WP (1 << 0)
#define UFFD_FEATURE_EVENT_FORK (1 << 1)
#define UFFD_FEATURE_EVENT_REMAP (1 << 2)
#define UFFD_FEATURE_EVENT_REMOVE (1 << 3)
#define UFFD_FEATURE_MISSING_HUGETLBFS (1 << 4)
#define UFFD_FEATURE_MISSING_SHMEM (1 << 5)
#define UFFD_FEATURE_EVENT_UNMAP (1 << 6)
#define UFFD_FEATURE_SIGBUS (1 << 7)
#define UFFD_FEATURE_THREAD_ID (1 << 8)
#define UFFD_FEATURE_MINOR_HUGETLBFS (1 << 9)
#define UFFD_FEATURE_MINOR_SHMEM (1 << 10)
#define UFFD_FEATURE_EXACT_ADDRESS (1 << 11)
#define UFFD_FEATURE_WP_HUGETLBFS_SHMEM (1 << 12)
#define UFFD_FEATURE_WP_UNPOPULATED (1 << 13)
#define UFFD_FEATURE_POISON (1 << 14)
#define UFFD_FEATURE_WP_ASYNC (1 << 15)
#define UFFD_FEATURE_MOVE (1 << 16)
  __u64 features;
  __u64 ioctls;
};
struct uffdio_range {
  __u64 start;
  __u64 len;
};
struct uffdio_register {
  struct uffdio_range range;
#define UFFDIO_REGISTER_MODE_MISSING ((__u64) 1 << 0)
#define UFFDIO_REGISTER_MODE_WP ((__u64) 1 << 1)
#define UFFDIO_REGISTER_MODE_MINOR ((__u64) 1 << 2)
  __u64 mode;
  __u64 ioctls;
};
struct uffdio_copy {
  __u64 dst;
  __u64 src;
  __u64 len;
#define UFFDIO_COPY_MODE_DONTWAKE ((__u64) 1 << 0)
#define UFFDIO_COPY_MODE_WP ((__u64) 1 << 1)
  __u64 mode;
  __s64 copy;
};
struct uffdio_zeropage {
  struct uffdio_range range;
#define UFFDIO_ZEROPAGE_MODE_DONTWAKE ((__u64) 1 << 0)
  __u64 mode;
  __s64 zeropage;
};
struct uffdio_writeprotect {
  struct uffdio_range range;
#define UFFDIO_WRITEPROTECT_MODE_WP ((__u64) 1 << 0)
#define UFFDIO_WRITEPROTECT_MODE_DONTWAKE ((__u64) 1 << 1)
  __u64 mode;
};
struct uffdio_continue {
  struct uffdio_range range;
#define UFFDIO_CONTINUE_MODE_DONTWAKE ((__u64) 1 << 0)
#define UFFDIO_CONTINUE_MODE_WP ((__u64) 1 << 1)
  __u64 mode;
  __s64 mapped;
};
struct uffdio_poison {
  struct uffdio_range range;
#define UFFDIO_POISON_MODE_DONTWAKE ((__u64) 1 << 0)
  __u64 mode;
  __s64 updated;
};
struct uffdio_move {
  __u64 dst;
  __u64 src;
  __u64 len;
#define UFFDIO_MOVE_MODE_DONTWAKE ((__u64) 1 << 0)
#define UFFDIO_MOVE_MODE_ALLOW_SRC_HOLES ((__u64) 1 << 1)
  __u64 mode;
  __s64 move;
};
#define UFFD_USER_MODE_ONLY 1
#endif
