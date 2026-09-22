/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_VMCORE_H
#define _UAPI_VMCORE_H
#include <linux/types.h>
#define VMCOREDD_NOTE_NAME "LINUX"
#define VMCOREDD_MAX_NAME_BYTES 44
struct vmcoredd_header {
  __u32 n_namesz;
  __u32 n_descsz;
  __u32 n_type;
  __u8 name[8];
  __u8 dump_name[VMCOREDD_MAX_NAME_BYTES];
};
#endif
