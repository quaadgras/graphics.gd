/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_IF_PLIP_H
#define _LINUX_IF_PLIP_H
#include <linux/sockios.h>
#define SIOCDEVPLIP SIOCDEVPRIVATE
struct plipconf {
  unsigned short pcmd;
  unsigned long nibble;
  unsigned long trigger;
};
#define PLIP_GET_TIMEOUT 0x1
#define PLIP_SET_TIMEOUT 0x2
#endif
