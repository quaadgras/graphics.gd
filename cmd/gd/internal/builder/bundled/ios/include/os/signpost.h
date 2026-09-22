// os/signpost.h declarations for graphics.gd's iOS SDK, written from the
// public documentation of the API. Not derived from Apple's SDK headers.
#ifndef GD_OS_SIGNPOST_H
#define GD_OS_SIGNPOST_H

#include <os/log.h>

#if defined(__cplusplus)
extern "C" {
#endif

typedef uint64_t os_signpost_id_t;
#define OS_SIGNPOST_ID_NULL ((os_signpost_id_t)0)
#define OS_SIGNPOST_ID_INVALID ((os_signpost_id_t)~0)
#define OS_SIGNPOST_ID_EXCLUSIVE ((os_signpost_id_t)0xEEEEB0B5B2B2EEEE)

typedef uint8_t os_signpost_type_t;
enum {
	OS_SIGNPOST_EVENT = 0x00,
	OS_SIGNPOST_INTERVAL_BEGIN = 0x01,
	OS_SIGNPOST_INTERVAL_END = 0x02,
};

os_signpost_id_t os_signpost_id_generate(os_log_t log);
os_signpost_id_t os_signpost_id_make_with_pointer(os_log_t log, const void *ptr);
bool os_signpost_enabled(os_log_t log);
void _os_signpost_emit_with_name_impl(void *dso, os_log_t log, os_signpost_type_t type, os_signpost_id_t spid, const char *name, const char *format, uint8_t *buf, uint32_t size);

// The format (and its arguments) are optional, "" adjoins a format if there is one.
#define _gd_signpost_format(format, ...) format
#define os_signpost_emit_with_type(log, type, spid, name, ...) __extension__({ \
	os_log_t _gd_log = (log); \
	os_signpost_id_t _gd_spid = (spid); \
	if (_gd_spid != OS_SIGNPOST_ID_NULL && _gd_spid != OS_SIGNPOST_ID_INVALID && os_signpost_enabled(_gd_log)) { \
		OS_LOG_STRING(_gd_name, name); \
		OS_LOG_STRING(_gd_format, _gd_signpost_format("" __VA_ARGS__)); \
		uint8_t _gd_buf[__builtin_os_log_format_buffer_size("" __VA_ARGS__)]; \
		_os_signpost_emit_with_name_impl(&__dso_handle, _gd_log, (type), _gd_spid, _gd_name, _gd_format, \
				(uint8_t *)__builtin_os_log_format(_gd_buf, "" __VA_ARGS__), (uint32_t)sizeof(_gd_buf)); \
	} \
})
#define os_signpost_event_emit(log, spid, name, ...) os_signpost_emit_with_type(log, OS_SIGNPOST_EVENT, spid, name, ##__VA_ARGS__)
#define os_signpost_interval_begin(log, spid, name, ...) os_signpost_emit_with_type(log, OS_SIGNPOST_INTERVAL_BEGIN, spid, name, ##__VA_ARGS__)
#define os_signpost_interval_end(log, spid, name, ...) os_signpost_emit_with_type(log, OS_SIGNPOST_INTERVAL_END, spid, name, ##__VA_ARGS__)

#if defined(__cplusplus)
}
#endif

#endif
