/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_IF_ALG_H
#define _LINUX_IF_ALG_H
#include <linux/types.h>
struct sockaddr_alg {
  __u16 salg_family;
  __u8 salg_type[14];
  __u32 salg_feat;
  __u32 salg_mask;
  __u8 salg_name[64];
};
struct sockaddr_alg_new {
  __u16 salg_family;
  __u8 salg_type[14];
  __u32 salg_feat;
  __u32 salg_mask;
  __u8 salg_name[];
};
struct af_alg_iv {
  __u32 ivlen;
  __u8 iv[];
};
#define ALG_SET_KEY 1
#define ALG_SET_IV 2
#define ALG_SET_OP 3
#define ALG_SET_AEAD_ASSOCLEN 4
#define ALG_SET_AEAD_AUTHSIZE 5
#define ALG_SET_DRBG_ENTROPY 6
#define ALG_SET_KEY_BY_KEY_SERIAL 7
#define ALG_OP_DECRYPT 0
#define ALG_OP_ENCRYPT 1
#endif
