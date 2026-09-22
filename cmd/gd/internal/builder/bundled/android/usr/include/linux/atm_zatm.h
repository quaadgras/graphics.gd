/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef LINUX_ATM_ZATM_H
#define LINUX_ATM_ZATM_H
#include <linux/atmapi.h>
#include <linux/atmioc.h>
#define ZATM_GETPOOL _IOW('a', ATMIOC_SARPRV + 1, struct atmif_sioc)
#define ZATM_GETPOOLZ _IOW('a', ATMIOC_SARPRV + 2, struct atmif_sioc)
#define ZATM_SETPOOL _IOW('a', ATMIOC_SARPRV + 3, struct atmif_sioc)
struct zatm_pool_info {
  int ref_count;
  int low_water, high_water;
  int rqa_count, rqu_count;
  int offset, next_off;
  int next_cnt, next_thres;
};
struct zatm_pool_req {
  int pool_num;
  struct zatm_pool_info info;
};
#define ZATM_OAM_POOL 0
#define ZATM_AAL0_POOL 1
#define ZATM_AAL5_POOL_BASE 2
#define ZATM_LAST_POOL ZATM_AAL5_POOL_BASE + 10
#define ZATM_TIMER_HISTORY_SIZE 16
#endif
