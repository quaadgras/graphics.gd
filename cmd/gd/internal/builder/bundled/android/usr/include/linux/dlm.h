/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI__DLM_DOT_H__
#define _UAPI__DLM_DOT_H__
#include <linux/dlmconstants.h>
#include <linux/types.h>
typedef void dlm_lockspace_t;
#define DLM_SBF_DEMOTED 0x01
#define DLM_SBF_VALNOTVALID 0x02
#define DLM_SBF_ALTMODE 0x04
struct dlm_lksb {
  int sb_status;
  __u32 sb_lkid;
  char sb_flags;
  char * sb_lvbptr;
};
#define DLM_LSFL_TIMEWARN 0x00000002
#define DLM_LSFL_NEWEXCL 0x00000008
#define __DLM_LSFL_RESERVED0 0x00000010
#endif
