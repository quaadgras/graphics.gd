/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_TLS_H
#define _UAPI_LINUX_TLS_H
#include <linux/types.h>
#define TLS_TX 1
#define TLS_RX 2
#define TLS_TX_ZEROCOPY_RO 3
#define TLS_RX_EXPECT_NO_PAD 4
#define TLS_VERSION_MINOR(ver) ((ver) & 0xFF)
#define TLS_VERSION_MAJOR(ver) (((ver) >> 8) & 0xFF)
#define TLS_VERSION_NUMBER(id) ((((id ##_VERSION_MAJOR) & 0xFF) << 8) | ((id ##_VERSION_MINOR) & 0xFF))
#define TLS_1_2_VERSION_MAJOR 0x3
#define TLS_1_2_VERSION_MINOR 0x3
#define TLS_1_2_VERSION TLS_VERSION_NUMBER(TLS_1_2)
#define TLS_1_3_VERSION_MAJOR 0x3
#define TLS_1_3_VERSION_MINOR 0x4
#define TLS_1_3_VERSION TLS_VERSION_NUMBER(TLS_1_3)
#define TLS_CIPHER_AES_GCM_128 51
#define TLS_CIPHER_AES_GCM_128_IV_SIZE 8
#define TLS_CIPHER_AES_GCM_128_KEY_SIZE 16
#define TLS_CIPHER_AES_GCM_128_SALT_SIZE 4
#define TLS_CIPHER_AES_GCM_128_TAG_SIZE 16
#define TLS_CIPHER_AES_GCM_128_REC_SEQ_SIZE 8
#define TLS_CIPHER_AES_GCM_256 52
#define TLS_CIPHER_AES_GCM_256_IV_SIZE 8
#define TLS_CIPHER_AES_GCM_256_KEY_SIZE 32
#define TLS_CIPHER_AES_GCM_256_SALT_SIZE 4
#define TLS_CIPHER_AES_GCM_256_TAG_SIZE 16
#define TLS_CIPHER_AES_GCM_256_REC_SEQ_SIZE 8
#define TLS_CIPHER_AES_CCM_128 53
#define TLS_CIPHER_AES_CCM_128_IV_SIZE 8
#define TLS_CIPHER_AES_CCM_128_KEY_SIZE 16
#define TLS_CIPHER_AES_CCM_128_SALT_SIZE 4
#define TLS_CIPHER_AES_CCM_128_TAG_SIZE 16
#define TLS_CIPHER_AES_CCM_128_REC_SEQ_SIZE 8
#define TLS_CIPHER_CHACHA20_POLY1305 54
#define TLS_CIPHER_CHACHA20_POLY1305_IV_SIZE 12
#define TLS_CIPHER_CHACHA20_POLY1305_KEY_SIZE 32
#define TLS_CIPHER_CHACHA20_POLY1305_SALT_SIZE 0
#define TLS_CIPHER_CHACHA20_POLY1305_TAG_SIZE 16
#define TLS_CIPHER_CHACHA20_POLY1305_REC_SEQ_SIZE 8
#define TLS_CIPHER_SM4_GCM 55
#define TLS_CIPHER_SM4_GCM_IV_SIZE 8
#define TLS_CIPHER_SM4_GCM_KEY_SIZE 16
#define TLS_CIPHER_SM4_GCM_SALT_SIZE 4
#define TLS_CIPHER_SM4_GCM_TAG_SIZE 16
#define TLS_CIPHER_SM4_GCM_REC_SEQ_SIZE 8
#define TLS_CIPHER_SM4_CCM 56
#define TLS_CIPHER_SM4_CCM_IV_SIZE 8
#define TLS_CIPHER_SM4_CCM_KEY_SIZE 16
#define TLS_CIPHER_SM4_CCM_SALT_SIZE 4
#define TLS_CIPHER_SM4_CCM_TAG_SIZE 16
#define TLS_CIPHER_SM4_CCM_REC_SEQ_SIZE 8
#define TLS_CIPHER_ARIA_GCM_128 57
#define TLS_CIPHER_ARIA_GCM_128_IV_SIZE 8
#define TLS_CIPHER_ARIA_GCM_128_KEY_SIZE 16
#define TLS_CIPHER_ARIA_GCM_128_SALT_SIZE 4
#define TLS_CIPHER_ARIA_GCM_128_TAG_SIZE 16
#define TLS_CIPHER_ARIA_GCM_128_REC_SEQ_SIZE 8
#define TLS_CIPHER_ARIA_GCM_256 58
#define TLS_CIPHER_ARIA_GCM_256_IV_SIZE 8
#define TLS_CIPHER_ARIA_GCM_256_KEY_SIZE 32
#define TLS_CIPHER_ARIA_GCM_256_SALT_SIZE 4
#define TLS_CIPHER_ARIA_GCM_256_TAG_SIZE 16
#define TLS_CIPHER_ARIA_GCM_256_REC_SEQ_SIZE 8
#define TLS_SET_RECORD_TYPE 1
#define TLS_GET_RECORD_TYPE 2
struct tls_crypto_info {
  __u16 version;
  __u16 cipher_type;
};
struct tls12_crypto_info_aes_gcm_128 {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_AES_GCM_128_IV_SIZE];
  unsigned char key[TLS_CIPHER_AES_GCM_128_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_AES_GCM_128_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_AES_GCM_128_REC_SEQ_SIZE];
};
struct tls12_crypto_info_aes_gcm_256 {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_AES_GCM_256_IV_SIZE];
  unsigned char key[TLS_CIPHER_AES_GCM_256_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_AES_GCM_256_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_AES_GCM_256_REC_SEQ_SIZE];
};
struct tls12_crypto_info_aes_ccm_128 {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_AES_CCM_128_IV_SIZE];
  unsigned char key[TLS_CIPHER_AES_CCM_128_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_AES_CCM_128_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_AES_CCM_128_REC_SEQ_SIZE];
};
struct tls12_crypto_info_chacha20_poly1305 {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_CHACHA20_POLY1305_IV_SIZE];
  unsigned char key[TLS_CIPHER_CHACHA20_POLY1305_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_CHACHA20_POLY1305_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_CHACHA20_POLY1305_REC_SEQ_SIZE];
};
struct tls12_crypto_info_sm4_gcm {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_SM4_GCM_IV_SIZE];
  unsigned char key[TLS_CIPHER_SM4_GCM_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_SM4_GCM_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_SM4_GCM_REC_SEQ_SIZE];
};
struct tls12_crypto_info_sm4_ccm {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_SM4_CCM_IV_SIZE];
  unsigned char key[TLS_CIPHER_SM4_CCM_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_SM4_CCM_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_SM4_CCM_REC_SEQ_SIZE];
};
struct tls12_crypto_info_aria_gcm_128 {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_ARIA_GCM_128_IV_SIZE];
  unsigned char key[TLS_CIPHER_ARIA_GCM_128_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_ARIA_GCM_128_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_ARIA_GCM_128_REC_SEQ_SIZE];
};
struct tls12_crypto_info_aria_gcm_256 {
  struct tls_crypto_info info;
  unsigned char iv[TLS_CIPHER_ARIA_GCM_256_IV_SIZE];
  unsigned char key[TLS_CIPHER_ARIA_GCM_256_KEY_SIZE];
  unsigned char salt[TLS_CIPHER_ARIA_GCM_256_SALT_SIZE];
  unsigned char rec_seq[TLS_CIPHER_ARIA_GCM_256_REC_SEQ_SIZE];
};
enum {
  TLS_INFO_UNSPEC,
  TLS_INFO_VERSION,
  TLS_INFO_CIPHER,
  TLS_INFO_TXCONF,
  TLS_INFO_RXCONF,
  TLS_INFO_ZC_RO_TX,
  TLS_INFO_RX_NO_PAD,
  __TLS_INFO_MAX,
};
#define TLS_INFO_MAX (__TLS_INFO_MAX - 1)
#define TLS_CONF_BASE 1
#define TLS_CONF_SW 2
#define TLS_CONF_HW 3
#define TLS_CONF_HW_RECORD 4
#endif
