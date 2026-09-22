/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef __PSP_DBC_USER_H__
#define __PSP_DBC_USER_H__
#include <linux/types.h>
#define DBC_NONCE_SIZE 16
#define DBC_SIG_SIZE 32
#define DBC_UID_SIZE 16
struct dbc_user_nonce {
  __u32 auth_needed;
  __u8 nonce[DBC_NONCE_SIZE];
  __u8 signature[DBC_SIG_SIZE];
} __attribute__((__packed__));
struct dbc_user_setuid {
  __u8 uid[DBC_UID_SIZE];
  __u8 signature[DBC_SIG_SIZE];
} __attribute__((__packed__));
struct dbc_user_param {
  __u32 msg_index;
  __u32 param;
  __u8 signature[DBC_SIG_SIZE];
} __attribute__((__packed__));
#define DBC_IOC_TYPE 'D'
#define DBCIOCNONCE _IOWR(DBC_IOC_TYPE, 0x1, struct dbc_user_nonce)
#define DBCIOCUID _IOW(DBC_IOC_TYPE, 0x2, struct dbc_user_setuid)
#define DBCIOCPARAM _IOWR(DBC_IOC_TYPE, 0x3, struct dbc_user_param)
enum dbc_cmd_msg {
  PARAM_GET_FMAX_CAP = 0x3,
  PARAM_SET_FMAX_CAP = 0x4,
  PARAM_GET_PWR_CAP = 0x5,
  PARAM_SET_PWR_CAP = 0x6,
  PARAM_GET_GFX_MODE = 0x7,
  PARAM_SET_GFX_MODE = 0x8,
  PARAM_GET_CURR_TEMP = 0x9,
  PARAM_GET_FMAX_MAX = 0xA,
  PARAM_GET_FMAX_MIN = 0xB,
  PARAM_GET_SOC_PWR_MAX = 0xC,
  PARAM_GET_SOC_PWR_MIN = 0xD,
  PARAM_GET_SOC_PWR_CUR = 0xE,
};
#endif
