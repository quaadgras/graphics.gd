#ifndef _ANDROID_LOG_H
#define _ANDROID_LOG_H

#include <stdarg.h>
#include <stdint.h>

typedef enum android_LogPriority {
    ANDROID_LOG_UNKNOWN = 0,
    ANDROID_LOG_DEFAULT,
    ANDROID_LOG_VERBOSE,
    ANDROID_LOG_DEBUG,
    ANDROID_LOG_INFO,
    ANDROID_LOG_WARN,
    ANDROID_LOG_ERROR,
    ANDROID_LOG_FATAL,
    ANDROID_LOG_SILENT,
} android_LogPriority;

typedef enum log_id {
    LOG_ID_MIN = 0,
    LOG_ID_MAIN = 0,
    LOG_ID_RADIO = 1,
    LOG_ID_EVENTS = 2,
    LOG_ID_SYSTEM = 3,
    LOG_ID_CRASH = 4,
    LOG_ID_STATS = 5,
    LOG_ID_SECURITY = 6,
    LOG_ID_KERNEL = 7,
    LOG_ID_MAX,
    LOG_ID_DEFAULT = 0x7FFFFFFF
} log_id_t;

struct __android_log_message {
    size_t sizeof_struct;
    int32_t buffer_id;
    int32_t priority;
    const char* tag;
    const char* file;
    uint32_t line;
    const char* message;
};

int __android_log_write(int prio, const char* tag, const char* msg);
int __android_log_print(int prio, const char* tag, const char* fmt, ...);
int __android_log_vprint(int prio, const char* tag, const char* fmt, va_list ap);
void __android_log_assert(const char* cond, const char* tag, const char* fmt = NULL, ...);
int __android_log_is_loggable(int prio, const char* tag, int default_prio);
int __android_log_is_loggable_len(int prio, const char* tag, size_t len, int default_prio);
int __android_log_buf_print(int bufID, int prio, const char* tag, const char* fmt, ...);
void __android_log_set_logger(void (*logger)(const struct __android_log_message*));
void __android_log_set_aborter(void (*aborter)(const char* abort_message));
void __android_log_call_aborter(const char* msg);
void __android_log_set_default_tag(const char* tag);
void __android_log_logd_logger(const struct __android_log_message* log_message);
int __android_log_set_minimum_priority(int priority);
void __android_log_write_log_message(struct __android_log_message* log_message);
int __android_log_get_minimum_priority(void);

#define ALOGW_IF(cond, ...) do { if (cond) ALOGW(__VA_ARGS__); } while(0)
#define LOG_WARN ANDROID_LOG_WARN
#define LOG_ALWAYS_FATAL_IF(cond, ...) do { if (cond) __android_log_assert(#cond, LOG_TAG, ##__VA_ARGS__); } while(0)

void android_errorWriteLog(int tag, const char *subTag);

#ifndef LOG_TAG
#define LOG_TAG NULL
#endif
#define ALOG(priority, tag, ...) __android_log_print(priority, tag, __VA_ARGS__)
#define ALOGV(...) ((void)0)
#define ALOGD(...) ((void)0)
#define ALOGI(...) ((void)0)
#define ALOGW(...) __android_log_print(ANDROID_LOG_WARN, LOG_TAG, __VA_ARGS__)
#define ALOGE(...) __android_log_print(ANDROID_LOG_ERROR, LOG_TAG, __VA_ARGS__)
#define ALOGF(...) __android_log_print(ANDROID_LOG_FATAL, LOG_TAG, __VA_ARGS__)
#define LOG_ALWAYS_FATAL(...) __android_log_assert(NULL, LOG_TAG, __VA_ARGS__)
#define LOG_FATAL_IF(cond, ...) do { if (cond) __android_log_assert(#cond, LOG_TAG, __VA_ARGS__); } while(0)
#define ALOG_ASSERT(cond, ...) LOG_FATAL_IF(!(cond), __VA_ARGS__)

/* TEMP_FAILURE_RETRY for musl */
#ifndef TEMP_FAILURE_RETRY
#define TEMP_FAILURE_RETRY(exp) ({         \
    decltype(exp) _rc;                   \
    do {                                   \
        _rc = (exp);                       \
    } while (_rc == -1 && errno == EINTR); \
    _rc; })
#endif

/* O_BINARY doesn't exist on Linux */
#ifndef O_BINARY
#define O_BINARY 0
#endif

#endif /* _ANDROID_LOG_H */
