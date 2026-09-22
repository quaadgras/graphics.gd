/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_AMT_H_
#define _UAPI_AMT_H_
enum ifla_amt_mode {
  AMT_MODE_GATEWAY = 0,
  AMT_MODE_RELAY,
  __AMT_MODE_MAX,
};
#define AMT_MODE_MAX (__AMT_MODE_MAX - 1)
enum {
  IFLA_AMT_UNSPEC,
  IFLA_AMT_MODE,
  IFLA_AMT_RELAY_PORT,
  IFLA_AMT_GATEWAY_PORT,
  IFLA_AMT_LINK,
  IFLA_AMT_LOCAL_IP,
  IFLA_AMT_REMOTE_IP,
  IFLA_AMT_DISCOVERY_IP,
  IFLA_AMT_MAX_TUNNELS,
  __IFLA_AMT_MAX,
};
#define IFLA_AMT_MAX (__IFLA_AMT_MAX - 1)
#endif
