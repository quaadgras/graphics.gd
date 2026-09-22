/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef LINUX_ATM_IDT77105_H
#define LINUX_ATM_IDT77105_H
#include <linux/types.h>
#include <linux/atmioc.h>
#include <linux/atmdev.h>
struct idt77105_stats {
  __u32 symbol_errors;
  __u32 tx_cells;
  __u32 rx_cells;
  __u32 rx_hec_errors;
};
#define IDT77105_GETSTAT _IOW('a', ATMIOC_PHYPRV + 2, struct atmif_sioc)
#define IDT77105_GETSTATZ _IOW('a', ATMIOC_PHYPRV + 3, struct atmif_sioc)
#endif
