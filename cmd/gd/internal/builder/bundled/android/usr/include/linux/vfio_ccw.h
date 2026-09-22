/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _VFIO_CCW_H_
#define _VFIO_CCW_H_
#include <linux/types.h>
struct ccw_io_region {
#define ORB_AREA_SIZE 12
  __u8 orb_area[ORB_AREA_SIZE];
#define SCSW_AREA_SIZE 12
  __u8 scsw_area[SCSW_AREA_SIZE];
#define IRB_AREA_SIZE 96
  __u8 irb_area[IRB_AREA_SIZE];
  __u32 ret_code;
} __attribute__((__packed__));
#define VFIO_CCW_ASYNC_CMD_HSCH (1 << 0)
#define VFIO_CCW_ASYNC_CMD_CSCH (1 << 1)
struct ccw_cmd_region {
  __u32 command;
  __u32 ret_code;
} __attribute__((__packed__));
struct ccw_schib_region {
#define SCHIB_AREA_SIZE 52
  __u8 schib_area[SCHIB_AREA_SIZE];
} __attribute__((__packed__));
struct ccw_crw_region {
  __u32 crw;
  __u32 pad;
} __attribute__((__packed__));
#endif
