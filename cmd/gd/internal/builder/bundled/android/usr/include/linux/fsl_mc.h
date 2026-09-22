/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_FSL_MC_H_
#define _UAPI_FSL_MC_H_
#include <linux/types.h>
#define MC_CMD_NUM_OF_PARAMS 7
struct fsl_mc_command {
  __le64 header;
  __le64 params[MC_CMD_NUM_OF_PARAMS];
};
#define FSL_MC_SEND_CMD_IOCTL_TYPE 'R'
#define FSL_MC_SEND_CMD_IOCTL_SEQ 0xE0
#define FSL_MC_SEND_MC_COMMAND _IOWR(FSL_MC_SEND_CMD_IOCTL_TYPE, FSL_MC_SEND_CMD_IOCTL_SEQ, struct fsl_mc_command)
#endif
