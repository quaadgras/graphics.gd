// wasi-libc (no threads) omits __cxa_thread_atexit; thread_local
// destructors in a single-instance wazero guest run at module teardown,
// so registering them is a no-op.
int __cxa_thread_atexit(void (*func)(void *), void *obj, void *dso) {
    (void)func; (void)obj; (void)dso;
    return 0;
}

// posix_madvise is advisory; wasi has no madvise, so accept and ignore.
int posix_madvise(void *addr, unsigned long len, int advice) {
    (void)addr; (void)len; (void)advice;
    return 0;
}
