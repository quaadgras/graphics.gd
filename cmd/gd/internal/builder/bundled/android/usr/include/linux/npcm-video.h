/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_NPCM_VIDEO_H
#define _UAPI_LINUX_NPCM_VIDEO_H
#include <linux/v4l2-controls.h>
#define V4L2_CID_NPCM_CAPTURE_MODE (V4L2_CID_USER_NPCM_BASE + 0)
enum v4l2_npcm_capture_mode {
  V4L2_NPCM_CAPTURE_MODE_COMPLETE = 0,
  V4L2_NPCM_CAPTURE_MODE_DIFF = 1,
};
#define V4L2_CID_NPCM_RECT_COUNT (V4L2_CID_USER_NPCM_BASE + 1)
#endif
