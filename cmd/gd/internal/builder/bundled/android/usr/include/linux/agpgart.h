/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_AGP_H
#define _UAPI_AGP_H
#define AGPIOC_BASE 'A'
#define AGPIOC_INFO _IOR(AGPIOC_BASE, 0, struct agp_info *)
#define AGPIOC_ACQUIRE _IO(AGPIOC_BASE, 1)
#define AGPIOC_RELEASE _IO(AGPIOC_BASE, 2)
#define AGPIOC_SETUP _IOW(AGPIOC_BASE, 3, struct agp_setup *)
#define AGPIOC_RESERVE _IOW(AGPIOC_BASE, 4, struct agp_region *)
#define AGPIOC_PROTECT _IOW(AGPIOC_BASE, 5, struct agp_region *)
#define AGPIOC_ALLOCATE _IOWR(AGPIOC_BASE, 6, struct agp_allocate *)
#define AGPIOC_DEALLOCATE _IOW(AGPIOC_BASE, 7, int)
#define AGPIOC_BIND _IOW(AGPIOC_BASE, 8, struct agp_bind *)
#define AGPIOC_UNBIND _IOW(AGPIOC_BASE, 9, struct agp_unbind *)
#define AGPIOC_CHIPSET_FLUSH _IO(AGPIOC_BASE, 10)
#define AGP_DEVICE "/dev/agpgart"
#ifndef TRUE
#define TRUE 1
#endif
#ifndef FALSE
#define FALSE 0
#endif
#include <linux/types.h>
struct agp_version {
  __u16 major;
  __u16 minor;
};
typedef struct _agp_info {
  struct agp_version version;
  __u32 bridge_id;
  __u32 agp_mode;
  unsigned long aper_base;
  __kernel_size_t aper_size;
  __kernel_size_t pg_total;
  __kernel_size_t pg_system;
  __kernel_size_t pg_used;
} agp_info;
typedef struct _agp_setup {
  __u32 agp_mode;
} agp_setup;
typedef struct _agp_segment {
  __kernel_off_t pg_start;
  __kernel_size_t pg_count;
  int prot;
} agp_segment;
typedef struct _agp_region {
  __kernel_pid_t pid;
  __kernel_size_t seg_count;
  struct _agp_segment * seg_list;
} agp_region;
typedef struct _agp_allocate {
  int key;
  __kernel_size_t pg_count;
  __u32 type;
  __u32 physical;
} agp_allocate;
typedef struct _agp_bind {
  int key;
  __kernel_off_t pg_start;
} agp_bind;
typedef struct _agp_unbind {
  int key;
  __u32 priority;
} agp_unbind;
#endif
