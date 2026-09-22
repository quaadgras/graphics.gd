/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __TPS6594_PFSM_H
#define __TPS6594_PFSM_H
#include <linux/const.h>
#include <linux/ioctl.h>
#include <linux/types.h>
struct pmic_state_opt {
  __u8 gpio_retention;
  __u8 ddr_retention;
  __u8 mcu_only_startup_dest;
};
#define PMIC_BASE 'P'
#define PMIC_GOTO_STANDBY _IO(PMIC_BASE, 0)
#define PMIC_GOTO_LP_STANDBY _IO(PMIC_BASE, 1)
#define PMIC_UPDATE_PGM _IO(PMIC_BASE, 2)
#define PMIC_SET_ACTIVE_STATE _IO(PMIC_BASE, 3)
#define PMIC_SET_MCU_ONLY_STATE _IOW(PMIC_BASE, 4, struct pmic_state_opt)
#define PMIC_SET_RETENTION_STATE _IOW(PMIC_BASE, 5, struct pmic_state_opt)
#endif
