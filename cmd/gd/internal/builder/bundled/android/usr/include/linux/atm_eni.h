/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef LINUX_ATM_ENI_H
#define LINUX_ATM_ENI_H
#include <linux/atmioc.h>
struct eni_multipliers {
  int tx, rx;
};
#define ENI_MEMDUMP _IOW('a', ATMIOC_SARPRV, struct atmif_sioc)
#define ENI_SETMULT _IOW('a', ATMIOC_SARPRV + 7, struct atmif_sioc)
#endif
