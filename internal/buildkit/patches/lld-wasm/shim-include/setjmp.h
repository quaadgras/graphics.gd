#ifndef _WAZERO_FAKE_SETJMP_H
#define _WAZERO_FAKE_SETJMP_H
/* wasi under wazero: no exception-handling proposal, so no real
 * setjmp/longjmp. Crash isolation comes from the host running every
 * tool invocation in a fresh module instance, so recovery paths are
 * dead code: setjmp always takes the first-return path and longjmp
 * traps out to the host. */
typedef int jmp_buf[1];
typedef int sigjmp_buf[1];
#define setjmp(env) ((void)(env), 0)
#define sigsetjmp(env, save) ((void)(env), (void)(save), 0)
#ifdef __cplusplus
extern "C" {
#endif
__attribute__((noreturn)) static inline void longjmp(jmp_buf env, int val) {
    (void)env; (void)val; __builtin_trap();
}
__attribute__((noreturn)) static inline void siglongjmp(sigjmp_buf env, int val) {
    (void)env; (void)val; __builtin_trap();
}
#ifdef __cplusplus
}
#endif
#endif
