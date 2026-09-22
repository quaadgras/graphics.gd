/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _LINUX_ATMSVC_H
#define _LINUX_ATMSVC_H
#include <linux/atmapi.h>
#include <linux/atm.h>
#include <linux/atmioc.h>
#define ATMSIGD_CTRL _IO('a', ATMIOC_SPECIAL)
enum atmsvc_msg_type {
  as_catch_null,
  as_bind,
  as_connect,
  as_accept,
  as_reject,
  as_listen,
  as_okay,
  as_error,
  as_indicate,
  as_close,
  as_itf_notify,
  as_modify,
  as_identify,
  as_terminate,
  as_addparty,
  as_dropparty
};
struct atmsvc_msg {
  enum atmsvc_msg_type type;
  atm_kptr_t vcc;
  atm_kptr_t listen_vcc;
  int reply;
  struct sockaddr_atmpvc pvc;
  struct sockaddr_atmsvc local;
  struct atm_qos qos;
  struct atm_sap sap;
  unsigned int session;
  struct sockaddr_atmsvc svc;
} __ATM_API_ALIGN;
#define SELECT_TOP_PCR(tp) ((tp).pcr ? (tp).pcr : (tp).max_pcr && (tp).max_pcr != ATM_MAX_PCR ? (tp).max_pcr : (tp).min_pcr ? (tp).min_pcr : ATM_MAX_PCR)
#endif
