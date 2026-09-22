// os/log.h declarations for graphics.gd's iOS SDK, written from the public
// documentation of the API. Not derived from Apple's SDK headers.
#ifndef GD_OS_LOG_H
#define GD_OS_LOG_H

#include <os/base.h>
#include <stdbool.h>
#include <stdint.h>

#if defined(__cplusplus)
extern "C" {
#endif

#if defined(__OBJC__)
@protocol OS_os_log;
typedef NSObject<OS_os_log> *os_log_t;
#else
typedef struct os_log_s *os_log_t;
#endif

typedef uint8_t os_log_type_t;
enum {
	OS_LOG_TYPE_DEFAULT = 0x00,
	OS_LOG_TYPE_INFO = 0x01,
	OS_LOG_TYPE_DEBUG = 0x02,
	OS_LOG_TYPE_ERROR = 0x10,
	OS_LOG_TYPE_FAULT = 0x11,
};

extern struct os_log_s _os_log_default;
#define OS_LOG_DEFAULT ((os_log_t)&_os_log_default)
extern struct os_log_s _os_log_disabled;
#define OS_LOG_DISABLED ((os_log_t)&_os_log_disabled)
extern void *__dso_handle;

os_log_t os_log_create(const char *subsystem, const char *category);
bool os_log_type_enabled(os_log_t oslog, os_log_type_t type);
void _os_log_impl(void *dso, os_log_t log, os_log_type_t type, const char *format, uint8_t *buf, uint32_t size);

#define OS_LOG_CATEGORY_POINTS_OF_INTEREST "PointsOfInterest"
#define OS_LOG_CATEGORY_DYNAMIC_TRACING "DynamicTracing"
#define OS_LOG_CATEGORY_DYNAMIC_STACK_TRACING "DynamicStackTracing"

// Format strings are kept apart from other strings, where the logging system
// (which records them by reference) expects to find them.
#define OS_LOG_STRING(name, string) \
	static const char name[] __attribute__((section("__TEXT,__oslogstring,cstring_literals"))) = string

// The arguments of a log are encoded into a buffer by the compiler.
#define os_log_with_type(log, type, format, ...) __extension__({ \
	os_log_t _gd_log = (log); \
	os_log_type_t _gd_type = (type); \
	if (os_log_type_enabled(_gd_log, _gd_type)) { \
		OS_LOG_STRING(_gd_format, format); \
		uint8_t _gd_buf[__builtin_os_log_format_buffer_size(format, ##__VA_ARGS__)]; \
		_os_log_impl(&__dso_handle, _gd_log, _gd_type, _gd_format, \
				(uint8_t *)__builtin_os_log_format(_gd_buf, format, ##__VA_ARGS__), (uint32_t)sizeof(_gd_buf)); \
	} \
})
#define os_log(log, format, ...) os_log_with_type(log, OS_LOG_TYPE_DEFAULT, format, ##__VA_ARGS__)
#define os_log_info(log, format, ...) os_log_with_type(log, OS_LOG_TYPE_INFO, format, ##__VA_ARGS__)
#define os_log_debug(log, format, ...) os_log_with_type(log, OS_LOG_TYPE_DEBUG, format, ##__VA_ARGS__)
#define os_log_error(log, format, ...) os_log_with_type(log, OS_LOG_TYPE_ERROR, format, ##__VA_ARGS__)
#define os_log_fault(log, format, ...) os_log_with_type(log, OS_LOG_TYPE_FAULT, format, ##__VA_ARGS__)

#if defined(__cplusplus)
}
#endif

#endif
