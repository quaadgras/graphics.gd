#include <cstdarg>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include "android/log.h"

int __android_log_write(int prio, const char* tag, const char* msg) {
    return fprintf(stderr, "%s: %s\n", tag, msg);
}

int __android_log_print(int prio, const char* tag, const char* fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "%s: ", tag);
    int ret = vfprintf(stderr, fmt, ap);
    fprintf(stderr, "\n");
    va_end(ap);
    return ret;
}

int __android_log_vprint(int prio, const char* tag, const char* fmt, va_list ap) {
    fprintf(stderr, "%s: ", tag);
    int ret = vfprintf(stderr, fmt, ap);
    fprintf(stderr, "\n");
    return ret;
}

void __android_log_assert(const char* cond, const char* tag, const char* fmt, ...) {
    if (fmt) {
        va_list ap;
        va_start(ap, fmt);
        fprintf(stderr, "ASSERT %s: %s: ", tag, cond ? cond : "(null)");
        vfprintf(stderr, fmt, ap);
        fprintf(stderr, "\n");
        va_end(ap);
    }
    __builtin_trap();
}

int __android_log_is_loggable(int prio, const char* tag, int default_prio) {
    return 1;
}

int __android_log_is_loggable_len(int prio, const char* tag, size_t len, int default_prio) {
    return 1;
}

int __android_log_buf_print(int bufID, int prio, const char* tag, const char* fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    int ret = __android_log_vprint(prio, tag, fmt, ap);
    va_end(ap);
    return ret;
}

void __android_log_set_logger(void (*logger)(const struct __android_log_message*)) {
    // Stub: no-op for host build
}

void __android_log_set_aborter(void (*aborter)(const char* abort_message)) {
    // Stub: no-op for host build
}

void __android_log_call_aborter(const char* msg) {
    fprintf(stderr, "%s\n", msg);
    exit(1);
}

void __android_log_set_default_tag(const char* tag) {
    // Stub: no-op for host build
}

void __android_log_logd_logger(const struct __android_log_message* log_message) {
    __android_log_print(log_message->priority, log_message->tag, "%s", log_message->message);
}

int __android_log_set_minimum_priority(int priority) {
    return 0;
}

char *program_invocation_short_name = const_cast<char *>("aapt2");

void __android_log_write_log_message(struct __android_log_message* log_message) {
    __android_log_print(log_message->priority, log_message->tag, "%s", log_message->message);
}

int __android_log_get_minimum_priority(void) {
    return ANDROID_LOG_DEFAULT;
}

void android_errorWriteLog(int tag, const char *subTag) {
    __android_log_print(ANDROID_LOG_ERROR, "ERROR", "tag 0x%x sub %s", tag, subTag);
}
