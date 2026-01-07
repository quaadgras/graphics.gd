#ifndef _CUTILS_TRACE_H
#define _CUTILS_TRACE_H
#define ATRACE_TAG_ALWAYS 0
#define ATRACE_TAG_RESOURCES 0
#define ATRACE_BEGIN(name) ((void)0)
#define ATRACE_END() ((void)0)
#define ATRACE_INT(name, value) ((void)0)
#define ATRACE_ENABLED() 0
static inline void atrace_begin(uint64_t tag, const char* name) {}
static inline void atrace_end(uint64_t tag) {}
#endif
