#ifndef _SYS_SYSTEM_PROPERTIES_H
#define _SYS_SYSTEM_PROPERTIES_H
#define PROP_VALUE_MAX 92
static inline int __system_property_get(const char* name, char* value) { value[0] = 0; return 0; }
#endif
