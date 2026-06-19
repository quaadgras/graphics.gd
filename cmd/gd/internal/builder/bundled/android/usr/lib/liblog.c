// Stub for android's liblog, providing the symbols cgo's runtime
// expects at link time. The dynamic linker resolves them to the
// device's real liblog.so at runtime, so these definitions are never
// executed — they only exist so `ld` is satisfied.
//
// Kept in sync with the symbol set in the bundled aarch64 liblog.so
// (run `nm -D --defined-only` to enumerate). When NDK adds new entry
// points we may need to extend this file, but the bundled binary stub
// has been stable for years.

#include <stdarg.h>

void __android_log_assert(const char* a, const char* b, const char* c, ...) { (void)a; (void)b; (void)c; }
int  __android_log_buf_print(int a, int b, const char* c, const char* d, ...) { (void)a; (void)b; (void)c; (void)d; return 0; }
int  __android_log_buf_write(int a, int b, const char* c, const char* d) { (void)a; (void)b; (void)c; (void)d; return 0; }
void __android_log_call_aborter(const char* a) { (void)a; }
void __android_log_default_aborter(const char* a) { (void)a; }
int  __android_log_get_minimum_priority(void) { return 0; }
int  __android_log_is_loggable(int a, const char* b, int c) { (void)a; (void)b; (void)c; return 0; }
int  __android_log_is_loggable_len(int a, const char* b, unsigned long c, int d) { (void)a; (void)b; (void)c; (void)d; return 0; }
void __android_log_logd_logger(const void* a) { (void)a; }
int  __android_log_print(int a, const char* b, const char* c, ...) { (void)a; (void)b; (void)c; return 0; }
int  __android_log_vprint(int a, const char* b, const char* c, va_list d) { (void)a; (void)b; (void)c; (void)d; return 0; }
int  __android_log_write(int a, const char* b, const char* c) { (void)a; (void)b; (void)c; return 0; }
void __android_log_set_aborter(void* a) { (void)a; }
void __android_log_set_default_tag(const char* a) { (void)a; }
void __android_log_set_logger(void* a) { (void)a; }
void __android_log_set_minimum_priority(int a) { (void)a; }
void __android_log_stderr_logger(const void* a) { (void)a; }
void __android_log_write_log_message(const void* a) { (void)a; }
