/*
 * This file is auto-generated. Modifications will be lost.
 *
 * See https://android.googlesource.com/platform/bionic/+/master/libc/kernel/
 * for more information.
 */
#ifndef _UAPI_LINUX_DNS_RESOLVER_H
#define _UAPI_LINUX_DNS_RESOLVER_H
#include <linux/types.h>
enum dns_payload_content_type {
  DNS_PAYLOAD_IS_SERVER_LIST = 0,
};
enum dns_payload_address_type {
  DNS_ADDRESS_IS_IPV4 = 0,
  DNS_ADDRESS_IS_IPV6 = 1,
};
enum dns_payload_protocol_type {
  DNS_SERVER_PROTOCOL_UNSPECIFIED = 0,
  DNS_SERVER_PROTOCOL_UDP = 1,
  DNS_SERVER_PROTOCOL_TCP = 2,
};
enum dns_record_source {
  DNS_RECORD_UNAVAILABLE = 0,
  DNS_RECORD_FROM_CONFIG = 1,
  DNS_RECORD_FROM_DNS_A = 2,
  DNS_RECORD_FROM_DNS_AFSDB = 3,
  DNS_RECORD_FROM_DNS_SRV = 4,
  DNS_RECORD_FROM_NSS = 5,
  NR__dns_record_source
};
enum dns_lookup_status {
  DNS_LOOKUP_NOT_DONE = 0,
  DNS_LOOKUP_GOOD = 1,
  DNS_LOOKUP_GOOD_WITH_BAD = 2,
  DNS_LOOKUP_BAD = 3,
  DNS_LOOKUP_GOT_NOT_FOUND = 4,
  DNS_LOOKUP_GOT_LOCAL_FAILURE = 5,
  DNS_LOOKUP_GOT_TEMP_FAILURE = 6,
  DNS_LOOKUP_GOT_NS_FAILURE = 7,
  NR__dns_lookup_status
};
struct dns_payload_header {
  __u8 zero;
  __u8 content;
  __u8 version;
} __attribute__((__packed__));
struct dns_server_list_v1_header {
  struct dns_payload_header hdr;
  __u8 source;
  __u8 status;
  __u8 nr_servers;
} __attribute__((__packed__));
struct dns_server_list_v1_server {
  __u16 name_len;
  __u16 priority;
  __u16 weight;
  __u16 port;
  __u8 source;
  __u8 status;
  __u8 protocol;
  __u8 nr_addrs;
} __attribute__((__packed__));
struct dns_server_list_v1_address {
  __u8 address_type;
} __attribute__((__packed__));
#endif
