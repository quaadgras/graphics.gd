/* dlopen.c - Cosmopolitan-style dlopen port for static musl binaries
 *
 * Supports loading both musl-built and glibc-built shared libraries on Linux.
 * A static musl binary cannot dlopen() glibc libraries directly, so we run a
 * tiny glibc "helper" program in-process (borrowing the host's ld.so + glibc)
 * that hands us the real dlopen/dlsym and a set of glibc TLS blocks. Every call
 * into a foreign library then switches the CPU TLS register (%fs / tpidr_el0)
 * to a glibc TLS block for the duration of the call (see foreign_tramp.S).
 *
 * The helper comes from one of two sources, in order:
 *   1. An embedded, prebuilt ELF loaded from an in-memory fd: helper_bin_*.c,
 *      cross-built by gen_helper.sh against an old glibc, when a real glibc is
 *      detected on the host, else helper_musl_bin_*.c, linked against musl,
 *      when musl's loader is. This needs no compiler and is immune to noexec
 *      mounts, so a stock glibc or musl desktop/server "just works".
 *   2. helper_src.h compiled on the fly with cc/gcc/clang and cached per-user.
 *      Used as a fallback if neither embedded helper suits the host.
 *
 * This file is only compiled when targeting musl (not glibc).
 */

#ifndef __GLIBC__

#define _GNU_SOURCE
#include <assert.h>
#include <dlfcn.h>
#include <elf.h>
#include <errno.h>
#include <fcntl.h>
#include <limits.h>
#include <pthread.h>
#include <setjmp.h>
#include <signal.h>
#include <spawn.h>
#include <stdatomic.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/auxv.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <unistd.h>

#include "helper_src.h"  /* HELPER: the borrowed-loader program source */

#define RTLD_LOCAL 0
#define RTLD_LAZY  1
#define RTLD_NOW   2
#define RTLD_GLOBAL 256

/* Auxiliary vector keys we copy through from the host to the helper's glibc
 * ld.so/libc. Not all are defined in every <elf.h>, so guard them. */
#ifndef AT_PLATFORM
#define AT_PLATFORM 15
#endif
#ifndef AT_HWCAP
#define AT_HWCAP 16
#endif
#ifndef AT_HWCAP2
#define AT_HWCAP2 26
#endif
#ifndef AT_EXECFN
#define AT_EXECFN 31
#endif
#ifndef AT_SYSINFO_EHDR
#define AT_SYSINFO_EHDR 33
#endif
#ifndef AT_MINSIGSTKSZ
#define AT_MINSIGSTKSZ 51
#endif

/* Uncomment to enable TLS debug tracing */
//#define DLOPEN_DEBUG 0

#define READ32LE(p) \
  ((uint32_t)(((const uint8_t *)(p))[0]) | \
   ((uint32_t)(((const uint8_t *)(p))[1]) << 8) | \
   ((uint32_t)(((const uint8_t *)(p))[2]) << 16) | \
   ((uint32_t)(((const uint8_t *)(p))[3]) << 24))

#define WRITE32LE(p, v) do { \
  uint8_t *_p = (uint8_t *)(p); \
  uint32_t _v = (v); \
  _p[0] = _v; _p[1] = _v >> 8; _p[2] = _v >> 16; _p[3] = _v >> 24; \
} while (0)

#define WRITE64LE(p, v) do { \
  uint8_t *_p = (uint8_t *)(p); \
  uint64_t _v = (v); \
  _p[0] = _v; _p[1] = _v >> 8; _p[2] = _v >> 16; _p[3] = _v >> 24; \
  _p[4] = _v >> 32; _p[5] = _v >> 40; _p[6] = _v >> 48; _p[7] = _v >> 56; \
} while (0)

#define TLS_POOL_SIZE 64

struct Loaded {
  char *base;
  char *entry;
  Elf64_Ehdr eh;
  Elf64_Phdr ph[25];
};

static _Thread_local char dlerror_buf[128];

/* TLS pool structure (must match helper_src.h's struct tls_pool). Only the
 * tls_ptrs array is read on the musl side; the semaphores/count/lock that
 * follow it in the helper are init-only and never touched here. */
struct tls_pool {
  void *tls_ptrs[TLS_POOL_SIZE];      /* glibc TLS pointers */
};

/* Shared state populated by the helper's callback (foreign_helper). The order of
 * the callback pointer array is an ABI contract with helper_src.h's main(). */
struct {
  struct tls_pool *pool;   /* pre-spawned glibc TLS pool */
  void *foreign_tls;       /* main thread's glibc TLS (bootstrap for on-demand) */
  void *native_tls;        /* main thread's musl TLS */
  bool is_supported;
  void *(*dlopen_real)(const char *, int);
  void *(*dlsym_real)(void *, const char *);
  int (*dlclose_real)(void *);
  char *(*dlerror_real)(void);
  jmp_buf jb;
  atomic_int next_slot;
  /* Factory (in the helper's glibc context) that spawns a parked glibc thread
   * and returns its TLS pointer, used to grow past the pre-spawned pool. */
  void *(*glibc_tcb_create)(void);
} __foreign;

/* --- async-signal blocking around the TLS-switched window --------------------
 *
 * While the CPU TLS register points at a glibc TLS block, no Go signal handler
 * may run: every Go handler reads the current goroutine `g` from %fs/tpidr-
 * relative TLS and would dereference garbage. The trampoline (foreign_tramp.S)
 * blocks signals itself on the hot path; the cold C entry points below use these
 * helpers. We block all catchable signals EXCEPT the synchronous faults
 * (SIGSEGV/SIGBUS/SIGFPE/SIGILL/SIGTRAP/SIGABRT) so a genuine fault inside a
 * foreign library still terminates promptly instead of being deferred. */
static void block_foreign_signals(sigset_t *old) {
  sigset_t s;
  sigfillset(&s);
  sigdelset(&s, SIGSEGV);
  sigdelset(&s, SIGBUS);
  sigdelset(&s, SIGFPE);
  sigdelset(&s, SIGILL);
  sigdelset(&s, SIGTRAP);
  sigdelset(&s, SIGABRT);
  pthread_sigmask(SIG_BLOCK, &s, old);
}
static void restore_signals(const sigset_t *old) {
  pthread_sigmask(SIG_SETMASK, old, NULL);
}

/* Map a CPU TLS pointer -> this thread's assigned glibc TLS pointer.
 *
 * Keyed by the value of the CPU TLS register (%fs / tpidr_el0), which the
 * trampoline already reads on entry and passes in -- so the lookup needs no
 * syscall (the previous design did a gettid() per foreign call). It is correct
 * regardless of which TLS is active: an OUTER call enters on the thread's native
 * (musl) TLS, a NESTED call (foreign lib -> callback -> wrapped fn) enters on the
 * thread's foreign (glibc) TLS, so we insert BOTH `native -> foreign` and
 * `foreign -> foreign`; either key resolves to the same foreign TLS.
 *
 * Each thread only ever reads its own two entries (no other thread looks up its
 * unique TLS pointers), so `ftls` needs no cross-thread synchronisation -- only
 * the `key` CAS is atomic. Threads beyond the pre-spawned pool get a freshly
 * created glibc TCB on demand (create_donated_tls), so there is no fixed cap.
 * Entries are never removed (donated glibc threads parked at exit are leaked) --
 * bounded for engine thread pools; thread-exit cleanup is a later refinement. */
#define FTLS_TABLE_SIZE 1024
static struct ftls_ent {
  _Atomic(uintptr_t) key;   /* CPU TLS pointer (0 = empty) */
  void *ftls;
} ftls_table[FTLS_TABLE_SIZE];

/* Defined later; needed by create_donated_tls below. */
static void *get_current_tls(void);
static void set_current_tls(void *tls);

static void ftls_abort(const char *msg) {
  write(2, msg, strlen(msg));
  abort();
}

/* Create a fresh, dedicated glibc TLS for a thread beyond the pre-spawned pool.
 * We ask the helper (glibc context) to spawn a parked glibc thread and donate
 * its TCB. That factory call must run under a valid glibc TLS, so we briefly
 * switch to the main thread's foreign TLS -- serialized by the lock, and that
 * TLS is otherwise unused (slot 0 is never handed out). The result is a glibc
 * TCB owned 1:1 by this thread (no borrowing/sharing, no cap). */
static pthread_mutex_t tcb_create_lock = PTHREAD_MUTEX_INITIALIZER;
static void *create_donated_tls(void) {
  if (!__foreign.glibc_tcb_create)
    ftls_abort("graphics.gd dlopen: on-demand glibc TLS unavailable "
               "(helper too old); aborting\n");
  pthread_mutex_lock(&tcb_create_lock);
  void *saved = get_current_tls();
  set_current_tls(__foreign.foreign_tls);
  void *tcb = __foreign.glibc_tcb_create();
  set_current_tls(saved);
  pthread_mutex_unlock(&tcb_create_lock);
  if (!tcb)
    ftls_abort("graphics.gd dlopen: glibc_tcb_create failed (out of resources); aborting\n");
  return tcb;
}

/* Insert key -> ftls (linear probe; idempotent on the same key). Only the owning
 * thread inserts its own keys, so the ftls store needs no publish barrier. */
static void ftls_insert(uintptr_t key, void *ftls) {
  unsigned start = (unsigned)((key >> 4) * 2654435761u) % FTLS_TABLE_SIZE;
  for (unsigned i = 0; i < FTLS_TABLE_SIZE; i++) {
    struct ftls_ent *e = &ftls_table[(start + i) % FTLS_TABLE_SIZE];
    uintptr_t expected = 0;
    if (atomic_compare_exchange_strong_explicit(
            &e->key, &expected, key, memory_order_acq_rel, memory_order_acquire)) {
      e->ftls = ftls;
      return;
    }
    if (expected == key) return;  /* already present */
  }
  ftls_abort("graphics.gd dlopen: foreign TLS table full; aborting\n");
}

/* Resolve (and lazily assign) the current thread's glibc TLS pointer, keyed on
 * `entry_tls` (the CPU TLS pointer active on entry, passed by foreign_tramp.S).
 * Non-static: the trampoline calls it by name. No syscall on the hot path. */
void *get_thread_foreign_tls(void *entry_tls) {
  uintptr_t key = (uintptr_t)entry_tls;
  unsigned start = (unsigned)((key >> 4) * 2654435761u) % FTLS_TABLE_SIZE;
  /* Fast path: this TLS pointer is already mapped (outer: native; nested: foreign). */
  for (unsigned i = 0; i < FTLS_TABLE_SIZE; i++) {
    struct ftls_ent *e = &ftls_table[(start + i) % FTLS_TABLE_SIZE];
    uintptr_t k = atomic_load_explicit(&e->key, memory_order_acquire);
    if (k == key) return e->ftls;
    if (k == 0) break;  /* no tombstones, so an empty slot ends the chain */
  }
  /* Miss: first foreign call by this thread (entry_tls is its native TLS).
   * Assign a glibc TLS from the pool, or create a dedicated one on demand. */
  void *ftls;
  int slot = atomic_fetch_add(&__foreign.next_slot, 1);
  if (slot < TLS_POOL_SIZE) {
    ftls = __foreign.pool ? __foreign.pool->tls_ptrs[slot] : NULL;
    if (!ftls)
      ftls_abort("graphics.gd dlopen: glibc TLS slot uninitialised "
                 "(helper thread pool incomplete); aborting\n");
  } else {
    ftls = create_donated_tls();
  }
  /* Map both the native key and the foreign TLS itself, so a later nested call
   * (which enters on the foreign TLS) also hits the fast path above. */
  ftls_insert(key, ftls);
  ftls_insert((uintptr_t)ftls, ftls);
  return ftls;
}

/* Forward declarations for assembly functions */
extern void *foreign_tramp(void);  // Assembly trampoline for TLS switching

/* Get current TLS pointer */
static void *get_current_tls(void) {
#ifdef __x86_64__
  void *tls;
  asm volatile ("mov %%fs:0, %0" : "=r"(tls));
  return tls;
#elif defined(__aarch64__)
  void *tls;
  asm volatile ("mrs %0, tpidr_el0" : "=r"(tls));
  return tls;
#else
#error "unsupported architecture"
#endif
}

/* Set current TLS pointer */
static void set_current_tls(void *tls) {
#ifdef DLOPEN_DEBUG
  void *old = get_current_tls();
  const char *from = (old == __foreign.native_tls) ? "native" :
                     (old == __foreign.foreign_tls) ? "foreign" : "unknown";
  const char *to = (tls == __foreign.native_tls) ? "native" :
                   (tls == __foreign.foreign_tls) ? "foreign" : "unknown";
  fprintf(stderr, "[TLS] %s -> %s (from %p to %p)\n", from, to, old, tls);
#endif
#ifdef __x86_64__
  asm volatile (
    "mov $0x1002, %%edi\n"  /* ARCH_SET_FS */
    "mov %0, %%rsi\n"
    "mov $158, %%eax\n"     /* __NR_arch_prctl */
    "syscall"
    : : "r"(tls) : "rdi", "rsi", "rax", "rcx", "r11", "memory"
  );
#elif defined(__aarch64__)
  asm volatile ("msr tpidr_el0, %0" : : "r"(tls) : "memory");
#else
#error "unsupported architecture"
#endif
}

static size_t my_strlcpy(char *dst, const char *src, size_t siz) {
  size_t len = strlen(src);
  if (siz) {
    size_t n = len < siz - 1 ? len : siz - 1;
    memcpy(dst, src, n);
    dst[n] = '\0';
  }
  return len;
}

static const char *get_program_executable_name(void) {
  static char buf[PATH_MAX];
  static bool initialized = false;
  if (!initialized) {
    ssize_t len = readlink("/proc/self/exe", buf, sizeof(buf) - 1);
    if (len > 0) {
      buf[len] = '\0';
    } else {
      strcpy(buf, "unknown");
    }
    initialized = true;
  }
  return buf;
}

static unsigned elf2prot(unsigned x) {
  unsigned r = 0;
  if (x & PF_R) r |= PROT_READ;
  if (x & PF_W) r |= PROT_WRITE;
  if (x & PF_X) r |= PROT_EXEC;
  return r;
}

static int get_host_elf_machine(void) {
#ifdef __x86_64__
  return EM_X86_64;
#elif defined(__aarch64__)
  return EM_AARCH64;
#else
#error "unsupported architecture"
#endif
}

static bool is_elf64(const Elf64_Ehdr *eh) {
  return memcmp(eh->e_ident, ELFMAG, SELFMAG) == 0 &&
         eh->e_ident[EI_CLASS] == ELFCLASS64;
}

static char *elf_map(int fd, const Elf64_Ehdr *ehdr, Elf64_Phdr *phdr, long pagesz,
                     char *interp_path, size_t interp_size) {
  Elf64_Addr minva = -1ULL, maxva = 0;
  for (int i = 0; i < ehdr->e_phnum; i++) {
    Elf64_Phdr *p = &phdr[i];
    if (p->p_type == PT_LOAD) {
      Elf64_Addr start = p->p_vaddr & -pagesz;
      if (start < minva) minva = start;
      Elf64_Addr end = p->p_vaddr + p->p_memsz;
      if (end > maxva) maxva = end;
    }
  }
  char *base = mmap(NULL, maxva - minva, PROT_NONE,
                    MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
  if (base == MAP_FAILED) return MAP_FAILED;

  bool failed = false;
  for (int i = 0; i < ehdr->e_phnum; i++) {
    Elf64_Phdr *p = &phdr[i];
    if (p->p_type == PT_LOAD) {
      Elf64_Addr skew = p->p_vaddr & (pagesz - 1);
      Elf64_Off off = p->p_offset - skew;
      Elf64_Addr a = p->p_vaddr + p->p_filesz;
      Elf64_Addr b = (a + pagesz - 1) & -pagesz;
      Elf64_Addr c = p->p_vaddr + p->p_memsz;
      int prot1 = elf2prot(p->p_flags);
      int prot2 = prot1;
      if (b > a) { prot1 |= PROT_WRITE; prot1 &= ~PROT_EXEC; }
      if (mmap(base + p->p_vaddr - skew, skew + p->p_filesz, prot1,
               MAP_FIXED | MAP_PRIVATE, fd, off) == MAP_FAILED) { failed = true; break; }
      if (b > a) memset(base + a, 0, b - a);
      if (c > b && mmap(base + b, c - b, prot2,
                        MAP_FIXED | MAP_PRIVATE | MAP_ANONYMOUS, -1, 0) == MAP_FAILED) { failed = true; break; }
      if (prot1 != prot2 && mprotect(base + p->p_vaddr - skew, skew + p->p_filesz, prot2)) { failed = true; break; }
    } else if (p->p_type == PT_INTERP && interp_size && interp_path) {
      if (pread(fd, interp_path, p->p_filesz, p->p_offset) != (ssize_t)p->p_filesz) { failed = true; break; }
      interp_path[p->p_filesz] = '\0';
    }
  }
  if (failed) {
    munmap(base, maxva - minva);
    return MAP_FAILED;
  }
  return base;
}

/* Load an ELF from an already-open, seekable fd (does not take ownership). */
static bool elf_load_fd(struct Loaded *l, int fd, long pagesz,
                        char *interp_path, size_t interp_size) {
  if (pread(fd, &l->eh, sizeof(l->eh), 0) != sizeof(l->eh) ||
      !is_elf64(&l->eh) ||
      l->eh.e_phnum > sizeof(l->ph)/sizeof(l->ph[0]) ||
      l->eh.e_machine != get_host_elf_machine()) {
    errno = ENOEXEC;
    return false;
  }

  if (pread(fd, l->ph, l->eh.e_phnum * sizeof(l->ph[0]), l->eh.e_phoff) !=
      (ssize_t)(l->eh.e_phnum * sizeof(l->ph[0]))) {
    return false;
  }

  l->base = elf_map(fd, &l->eh, l->ph, pagesz, interp_path, interp_size);
  if (l->base == MAP_FAILED) return false;

  l->entry = l->base + l->eh.e_entry;
  return true;
}

static bool elf_load(struct Loaded *l, const char *file, long pagesz,
                     char *interp_path, size_t interp_size) {
  int fd = open(file, O_RDONLY | O_CLOEXEC);
  if (fd == -1) return false;
  bool ok = elf_load_fd(l, fd, pagesz, interp_path, interp_size);
  close(fd);
  return ok;
}

static void foreign_helper(void **p) {
  __foreign.dlopen_real = p[0];
  __foreign.dlsym_real = p[1];
  __foreign.dlclose_real = p[2];
  __foreign.dlerror_real = p[3];
  __foreign.pool = p[4];  /* TLS pool for multi-threaded support */
  __foreign.glibc_tcb_create = p[5];  /* on-demand glibc TCB factory */

  /* Capture the foreign TLS - we're running in foreign context now */
  __foreign.foreign_tls = get_current_tls();

  /* Initialize slot assignment - slot 0 is main thread */
  __foreign.next_slot = 1;

  longjmp(__foreign.jb, 1);
}

/* Map the helper (from prog_fd) and its interpreter in-process, build a proper
 * ELF initial stack, and jump into the interpreter. On success control returns
 * via foreign_helper's longjmp (never through this function). Returns to the
 * caller only on failure; the caller owns prog_fd. */
static void elf_exec_fd(int prog_fd, char **envp) {
  long pagesz = sysconf(_SC_PAGESIZE);
  if (pagesz <= 0) pagesz = 4096;

  struct Loaded prog;
  char interp_path[256] = {0};
  if (!elf_load_fd(&prog, prog_fd, pagesz, interp_path, sizeof(interp_path))) return;

  struct Loaded interp;
  if (!elf_load(&interp, interp_path, pagesz, NULL, 0)) return;

  char *map = mmap(NULL, 128 << 10, PROT_READ | PROT_WRITE,
                   MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
  if (map == MAP_FAILED) return;

  /* Build proper ELF initial stack:
   * - argc
   * - argv[0..argc-1]
   * - NULL (argv terminator)
   * - envp[0..n-1]
   * - NULL (envp terminator)
   * - auxv pairs (key, value, ..., AT_NULL, 0)
   */
  long *stack_top = (long *)(map + (128 << 10));
  long *sp = stack_top;

  char addr_str[32];
  snprintf(addr_str, sizeof(addr_str), "%lu", (unsigned long)(uintptr_t)foreign_helper);

  /* Random bytes for AT_RANDOM (16 bytes required by glibc) */
  static char random_bytes[16] = {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16};

  /* Count environment variables */
  int envc = 0;
  if (envp) {
    while (envp[envc]) envc++;
  }

  /* Auxiliary vectors (pushed in reverse order) */
  *--sp = 0;                              /* AT_NULL value */
  *--sp = 0;                              /* AT_NULL key */

  /* Copy through host auxv entries that modern glibc ld.so/libc consult but
   * that we'd otherwise omit. Missing AT_SYSINFO_EHDR (vDSO) / AT_HWCAP /
   * AT_PLATFORM etc. cause glibc-version- and CPU-dependent misbehaviour in
   * the in-process helper. These pointers/values are valid in our own address
   * space, which the helper shares. */
  unsigned long auxval;
#define PUSH_HOST_AUXV(key) \
  do { if ((auxval = getauxval(key))) { *--sp = (long)auxval; *--sp = (key); } } while (0)
  PUSH_HOST_AUXV(AT_SYSINFO_EHDR);
  PUSH_HOST_AUXV(AT_HWCAP);
  PUSH_HOST_AUXV(AT_HWCAP2);
  PUSH_HOST_AUXV(AT_PLATFORM);
  PUSH_HOST_AUXV(AT_EXECFN);
  PUSH_HOST_AUXV(AT_MINSIGSTKSZ);
#undef PUSH_HOST_AUXV

  *--sp = 0;                              /* AT_SECURE value */
  *--sp = 23;                             /* AT_SECURE key */
  /* Prefer the host's real AT_RANDOM (glibc derives the stack canary and
   * pointer guard from these 16 bytes); fall back to our fixed bytes. */
  { unsigned long r = getauxval(AT_RANDOM);
    *--sp = r ? (long)r : (long)random_bytes; }
  *--sp = 25;                             /* AT_RANDOM key */
  *--sp = 100;                            /* AT_CLKTCK value */
  *--sp = 17;                             /* AT_CLKTCK key */
  *--sp = (long)getegid();                /* AT_EGID value */
  *--sp = 14;                             /* AT_EGID key */
  *--sp = (long)getgid();                 /* AT_GID value */
  *--sp = 13;                             /* AT_GID key */
  *--sp = (long)geteuid();                /* AT_EUID value */
  *--sp = 12;                             /* AT_EUID key */
  *--sp = (long)getuid();                 /* AT_UID value */
  *--sp = 11;                             /* AT_UID key */
  *--sp = (long)pagesz;                   /* AT_PAGESZ value */
  *--sp = 6;                              /* AT_PAGESZ key */
  *--sp = (long)interp.base;              /* AT_BASE value (interpreter base) */
  *--sp = 7;                              /* AT_BASE key */
  *--sp = 0;                              /* AT_FLAGS value */
  *--sp = 8;                              /* AT_FLAGS key */
  *--sp = (long)prog.entry;               /* AT_ENTRY value (program entry) */
  *--sp = 9;                              /* AT_ENTRY key */
  *--sp = (long)prog.eh.e_phnum;          /* AT_PHNUM value */
  *--sp = 5;                              /* AT_PHNUM key */
  *--sp = (long)prog.eh.e_phentsize;      /* AT_PHENT value */
  *--sp = 4;                              /* AT_PHENT key */
  *--sp = (long)(prog.base + prog.eh.e_phoff); /* AT_PHDR value */
  *--sp = 3;                              /* AT_PHDR key */

  /* envp terminator */
  *--sp = 0;

  /* Environment variables (in reverse order) */
  for (int i = envc - 1; i >= 0; i--) {
    *--sp = (long)envp[i];
  }

  /* argv terminator */
  *--sp = 0;

  /* argv[1] = callback address */
  *--sp = (long)addr_str;

  /* argv[0] = program name */
  *--sp = (long)get_program_executable_name();

  /* argc */
  *--sp = 2;

  /* Ensure 16-byte stack alignment as required by x86_64 ABI.
   * The stack must be 16-byte aligned at process entry, with argc at sp.
   * If currently misaligned, we need to shift everything down by 8 bytes.
   */
  if ((uintptr_t)sp & 8) {
    /* Shift all stack data down by 8 bytes to align */
    size_t count = stack_top - sp;
    memmove(sp - 1, sp, count * sizeof(long));
    sp--;
  }

#ifdef __x86_64__
  asm volatile (
    "mov %0, %%rsp\n"
    "jmp *%1"
    : : "r"(sp), "r"(interp.entry) : "memory"
  );
#elif defined(__aarch64__)
  register long x0 asm("x0") = 0;
  register long x9 asm("x9") = (long)sp;
  register long x16 asm("x16") = (long)interp.entry;
  asm volatile ("mov sp, x9\n br x16" : : "r"(x0), "r"(x9), "r"(x16) : "memory");
#endif
  __builtin_unreachable();
}

static char *dlerror_set(const char *s) {
  my_strlcpy(dlerror_buf, s ? s : "Unknown error", sizeof(dlerror_buf));
  return dlerror_buf;
}

/* Allocate a writable buffer to assemble a JIT stub into. The buffer is mapped
 * read/write only (never RWX), preserving W^X so this works on hardened kernels
 * (SELinux execmem, PaX MPROTECT, prctl PR_SET_MDWE). Each stub gets its own
 * page-aligned mapping; stubs are few, so the per-page slack is negligible.
 * Call foreign_seal() once the stub bytes are written to make it executable. */
static void *foreign_alloc(size_t n) {
  long pagesz = sysconf(_SC_PAGESIZE);
  if (pagesz <= 0) pagesz = 4096;
  size_t len = (n + (size_t)pagesz - 1) & ~((size_t)pagesz - 1);
  void *p = mmap(NULL, len, PROT_READ | PROT_WRITE,
                 MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
  if (p == MAP_FAILED) {
    dlerror_set("dlopen: failed to allocate JIT stub memory");
    return NULL;
  }
  return p;
}

/* Make a freshly-written stub executable: drop write, add execute (W^X) and
 * flush the instruction cache (required on aarch64, harmless on x86_64).
 * Returns false and leaves a dlerror if the kernel refuses PROT_EXEC, which
 * is then surfaced to the caller as a NULL function pointer rather than a
 * crash through unwritable/unexecutable memory. */
static bool foreign_seal(void *p, size_t n) {
  if (!p) return false;
  long pagesz = sysconf(_SC_PAGESIZE);
  if (pagesz <= 0) pagesz = 4096;
  uintptr_t start = (uintptr_t)p & ~((uintptr_t)pagesz - 1);
  size_t len = ((uintptr_t)p + n - start + (size_t)pagesz - 1) & ~((size_t)pagesz - 1);
  if (mprotect((void *)start, len, PROT_READ | PROT_EXEC) != 0) {
    dlerror_set("dlopen: kernel denied executable JIT memory (hardened W^X policy?)");
    return false;
  }
  __builtin___clear_cache((char *)p, (char *)p + n);
  return true;
}

/* Generate a tiny trampoline stub for a wrapped foreign function.
 *
 * The stub does nothing but hand control to the shared foreign_tramp with the
 * real function pointer in a scratch register; foreign_tramp resolves this
 * thread's glibc TLS, saves all argument/return registers around the switch,
 * calls the real function, and restores the caller's TLS. All argument
 * registers, the variadic count (%al / none on aarch64) and the struct-return
 * pointer flow through untouched. */
__attribute__((noinline))
static void *foreign_wrap(void *real_func) {
  if (!real_func) return NULL;

#ifdef DLOPEN_DEBUG
  fprintf(stderr, "[TRAMP] wrapping function at %p, foreign_tramp at %p\n",
          real_func, (void*)foreign_tramp);
#endif

#ifdef __x86_64__
  /* movabs $real_func,%r11 ; movabs $foreign_tramp,%r10 ; jmp *%r10
   * r10/r11 are caller-saved and not argument registers, so %rax (the variadic
   * vector count), all integer/SSE args and the return address are preserved. */
  unsigned char *stub = foreign_alloc(32);
  if (!stub) return NULL;
  int i = 0;
  stub[i++] = 0x49; stub[i++] = 0xbb;                 /* movabs imm64, %r11 */
  WRITE64LE(stub + i, (uintptr_t)real_func); i += 8;
  stub[i++] = 0x49; stub[i++] = 0xba;                 /* movabs imm64, %r10 */
  WRITE64LE(stub + i, (uintptr_t)foreign_tramp); i += 8;
  stub[i++] = 0x41; stub[i++] = 0xff; stub[i++] = 0xe2; /* jmp *%r10 */
  if (!foreign_seal(stub, i)) return NULL;
  return stub;

#elif defined(__aarch64__)
  /* ldr x9,real_func ; ldr x16,foreign_tramp ; br x16
   * x9/x16 are caller-saved temporaries; x0-x7, v0-v7 and x8 (indirect result)
   * are untouched. Literals are 8-byte aligned. */
  unsigned char *stub = foreign_alloc(32);
  if (!stub) return NULL;
  WRITE32LE(stub + 0,  0x58000089);  /* ldr x9,  [pc, #16] -> real_func @16 */
  WRITE32LE(stub + 4,  0x580000b0);  /* ldr x16, [pc, #20] -> foreign_tramp @24 */
  WRITE32LE(stub + 8,  0xd61f0200);  /* br x16 */
  WRITE32LE(stub + 12, 0xd503201f);  /* nop (pad literals to 8-byte alignment) */
  WRITE64LE(stub + 16, (uintptr_t)real_func);
  WRITE64LE(stub + 24, (uintptr_t)foreign_tramp);
  if (!foreign_seal(stub, 32)) return NULL;
  return stub;
#else
#error "unsupported architecture"
#endif
}

/* --- proc-addr getter interposition -------------------------------------------
 *
 * Loaders like Vulkan and GLX hand out further entry points through their own
 * getter functions (vkGetInstanceProcAddr, vkGetDeviceProcAddr,
 * glXGetProcAddress, eglGetProcAddress). Those pointers cross the musl<->glibc
 * boundary as plain data, so dlsym's trampoline wrapping never sees them and
 * the caller ends up invoking raw glibc code under the musl TLS register. The
 * first TLS-dependent thing that code touches then reads a bogus TCB - e.g.
 * glibc's __ctype_tolower_loc() returning a junk locale table inside the
 * Vulkan loader's strcasecmp, crashing on a NULL-based table index.
 *
 * Fix: when dlsym is asked for one of these getters, hand back a native shim
 * instead. The shim calls the real getter through its foreign trampoline (so
 * the lookup itself runs under glibc TLS) and passes every function pointer it
 * returns through foreign_wrap before it re-enters native code. Self
 * references (asking a getter for itself, or for the device-level getter) are
 * answered with the shims again so no raw getter ever escapes. */

static void *(*vk_gipa_real)(void *instance, const char *name); /* trampolined */
static void *(*vk_gdpa_real)(void *device, const char *name);   /* trampolined */
static void *(*gl_gpa_real)(const char *name);                  /* trampolined */

static void *gd_vkGetDeviceProcAddr(void *device, const char *name);

static void *gd_vkGetInstanceProcAddr(void *instance, const char *name) {
  if (!vk_gipa_real || !name) return NULL;
  if (!strcmp(name, "vkGetInstanceProcAddr"))
    return (void *)gd_vkGetInstanceProcAddr;
  void *raw = vk_gipa_real(instance, name);
  if (!raw) return NULL;
  if (!strcmp(name, "vkGetDeviceProcAddr")) {
    if (!vk_gdpa_real)
      vk_gdpa_real = (void *(*)(void *, const char *))foreign_wrap(raw);
    return (void *)gd_vkGetDeviceProcAddr;
  }
  return foreign_wrap(raw);
}

static void *gd_vkGetDeviceProcAddr(void *device, const char *name) {
  if (!vk_gdpa_real || !name) return NULL;
  if (!strcmp(name, "vkGetDeviceProcAddr"))
    return (void *)gd_vkGetDeviceProcAddr;
  void *raw = vk_gdpa_real(device, name);
  if (!raw) return NULL;
  return foreign_wrap(raw);
}

static void *gd_glGetProcAddress(const char *name) {
  if (!gl_gpa_real || !name) return NULL;
  /* Asking the getter for a getter (glad reloads the whole GLX symbol set
   * through it, glXGetProcAddressARB included) must hand the shims back:
   * wrapping the raw getter with a plain trampoline would make its RETURNED
   * pointers escape unwrapped, and the first GL call under the native TLS
   * register faults in libGLdispatch's TLS access. */
  if (!strcmp(name, "glXGetProcAddress") ||
      !strcmp(name, "glXGetProcAddressARB") ||
      !strcmp(name, "eglGetProcAddress") ||
      !strcmp(name, "wlEglGetProcAddress") ||
      !strcmp(name, "SDL_GL_GetProcAddress"))
    return (void *)gd_glGetProcAddress;
  if (!strcmp(name, "vkGetInstanceProcAddr") && vk_gipa_real)
    return (void *)gd_vkGetInstanceProcAddr;
  void *raw = gl_gpa_real(name);
  if (!raw) return NULL;
  return foreign_wrap(raw);
}

/* Called from dlsym with the requested symbol name and its already-wrapped
 * trampoline. Returns the shim to hand out instead, or NULL to let dlsym
 * return `wrapped` unchanged. */
static void *interpose_proc_getter(const char *name, void *wrapped) {
  if (!name) return NULL;
  if (!strcmp(name, "vkGetInstanceProcAddr")) {
    vk_gipa_real = (void *(*)(void *, const char *))wrapped;
    return (void *)gd_vkGetInstanceProcAddr;
  }
  if (!strcmp(name, "vkGetDeviceProcAddr")) {
    vk_gdpa_real = (void *(*)(void *, const char *))wrapped;
    return (void *)gd_vkGetDeviceProcAddr;
  }
  if (!strcmp(name, "glXGetProcAddress") ||
      !strcmp(name, "glXGetProcAddressARB") ||
      !strcmp(name, "eglGetProcAddress") ||
      !strcmp(name, "wlEglGetProcAddress") ||
      !strcmp(name, "SDL_GL_GetProcAddress")) {
    gl_gpa_real = (void *(*)(const char *))wrapped;
    return (void *)gd_glGetProcAddress;
  }
  return NULL;
}

/* --- embedded prebuilt helpers ------------------------------------------------
 *
 * helper_bin_<arch>.c and helper_musl_bin_<arch>.c define these (guarded by
 * arch) with the ELFs cross-built by gen_helper.sh: one against an old glibc,
 * one dynamically linked against musl. */
#if defined(__x86_64__) || defined(__aarch64__)
extern const unsigned char helper_bin[];
extern const unsigned int helper_bin_len;
extern const unsigned char helper_musl_bin[];
extern const unsigned int helper_musl_bin_len;
#define HAVE_EMBEDDED_HELPER 1
#endif

#if defined(__x86_64__)
#define GLIBC_LDSO "/lib64/ld-linux-x86-64.so.2"
#define GLIBC_TRIPLE "x86_64-linux-gnu"
#define MUSL_LDSO "/lib/ld-musl-x86_64.so.1"
#elif defined(__aarch64__)
#define GLIBC_LDSO "/lib/ld-linux-aarch64.so.1"
#define GLIBC_TRIPLE "aarch64-linux-gnu"
#define MUSL_LDSO "/lib/ld-musl-aarch64.so.1"
#endif

/* Is a real glibc runtime present on this host? The embedded glibc helper is a
 * glibc binary; jumping into a non-glibc loader (e.g. a musl host, or a gcompat
 * stub masquerading as /lib64/ld-linux-*) could hard-exit us with no chance to
 * fall back. Requiring libc.so.6 on disk keeps that helper to genuine glibc
 * hosts. libdl.so.2 must sit next to it: the helper is built against glibc
 * 2.17, where dlopen lives in libdl, and the loader _exit(127)s the whole
 * process when a needed library is missing (busybox:glibc ships libc.so.6
 * alone). glibc >= 2.34 still installs libdl.so.2 as a stub. */
static bool has_glibc_libc(void) {
  static const char *const dirs[] = {
    "/lib/" GLIBC_TRIPLE,
    "/usr/lib/" GLIBC_TRIPLE,
    "/lib64",
    "/usr/lib64",
    "/usr/lib",
    "/lib",
  };
  char path[64];
  for (size_t i = 0; i < sizeof(dirs)/sizeof(dirs[0]); i++) {
    snprintf(path, sizeof(path), "%s/libc.so.6", dirs[i]);
    if (access(path, F_OK) != 0) continue;
    snprintf(path, sizeof(path), "%s/libdl.so.2", dirs[i]);
    if (access(path, F_OK) == 0) return true;
  }
  /* NixOS keeps glibc in /nix/store, out of the paths above. With nix-ld
   * enabled (NixOS's own way to run binaries built elsewhere) the standard
   * loader path is nix-ld's shim, which hands over to the real glibc loader
   * and adds NIX_LD_LIBRARY_PATH to its search path. */
  if (access("/etc/NIXOS", F_OK) == 0 && access(GLIBC_LDSO, F_OK) == 0)
    return true;
  return false;
}

/* Is musl's dynamic loader (which is also its libc) present? Checked after
 * glibc, so a glibc host that happens to have musl installed beside it (e.g.
 * Debian's musl package) still loads its libraries through glibc. */
static bool has_musl_loader(void) {
  return access(MUSL_LDSO, F_OK) == 0;
}

/* Return an fd holding the embedded helper ELF that suits this host, or -1 if
 * there is none. Uses an anonymous in-memory file (memfd), so it never touches
 * disk and is immune to noexec mounts -- elf_exec_fd maps it PROT_EXEC directly
 * from the fd. */
static int embedded_helper_fd(void) {
#ifdef HAVE_EMBEDDED_HELPER
  const unsigned char *p;
  size_t left;
  if (has_glibc_libc()) {
    p = helper_bin;
    left = helper_bin_len;
  } else if (has_musl_loader()) {
    p = helper_musl_bin;
    left = helper_musl_bin_len;
  } else {
    return -1;
  }
  int fd = memfd_create("gdhelper", MFD_CLOEXEC);
  if (fd < 0) return -1;
  while (left) {
    ssize_t n = write(fd, p, left);
    if (n <= 0) { close(fd); return -1; }
    p += (size_t)n;
    left -= (size_t)n;
  }
  return fd;
#else
  return -1;
#endif
}

/* --- compiled-helper fallback (per-user, ownership-verified cache) ------------
 *
 * On hosts where no embedded helper can run (no glibc or musl at the standard
 * paths, or glibc older than the embedded baseline) we compile helper_src.h with the system C compiler. The
 * cache is keyed by a hash of the helper source, so an ABI change to the helper
 * never reuses a stale binary, and lives in a per-user 0700 directory that we
 * verify we own before executing anything from it in-process. */

static uint64_t fnv1a(const char *s) {
  uint64_t h = 1469598103934665603ULL;
  while (*s) { h ^= (unsigned char)*s++; h *= 1099511628211ULL; }
  return h;
}

/* True if `d` is a directory we own with no group/other write bit. */
static bool path_is_own_secure(const char *d) {
  struct stat st;
  if (lstat(d, &st)) return false;
  return st.st_uid == getuid() && !(st.st_mode & (S_IWGRP | S_IWOTH));
}

/* Pick (creating if needed) a per-user, 0700, we-own-it cache directory that is
 * suitable for holding an executable. Tries XDG_CACHE_HOME, ~/.cache, TMPDIR,
 * then /tmp. Returns true and fills `out` (>= PATH_MAX), or false. */
static bool make_cache_dir(char *out) {
  unsigned uid = (unsigned)getuid();
  char cand[PATH_MAX];
  const char *xdg = getenv("XDG_CACHE_HOME");
  const char *home = getenv("HOME");
  const char *tmp = getenv("TMPDIR");
  for (int which = 0; which < 4; which++) {
    switch (which) {
      case 0: if (!xdg || !*xdg) continue;
        snprintf(cand, sizeof(cand), "%s/graphics.gd-dlopen-%u", xdg, uid); break;
      case 1: if (!home || !*home) continue;
        snprintf(cand, sizeof(cand), "%s/.cache/graphics.gd-dlopen-%u", home, uid); break;
      case 2: if (!tmp || !*tmp) continue;
        snprintf(cand, sizeof(cand), "%s/graphics.gd-dlopen-%u", tmp, uid); break;
      default:
        snprintf(cand, sizeof(cand), "/tmp/graphics.gd-dlopen-%u", uid); break;
    }
    if (mkdir(cand, 0700) != 0 && errno != EEXIST) continue;
    if (!path_is_own_secure(cand)) continue;  /* hijacked or wrong owner */
    my_strlcpy(out, cand, PATH_MAX);
    return true;
  }
  return false;
}

/* Compile `src` to `out` with `compiler`, capturing the compiler's stderr into
 * errbuf on failure. Returns true on a clean build, false if the compiler was
 * missing or the compile failed. */
static bool run_compiler(const char *compiler, const char *src, const char *out,
                         char *errbuf, size_t errn) {
  char errfile[PATH_MAX];
  snprintf(errfile, sizeof(errfile), "%s.err", out);

  posix_spawn_file_actions_t fa;
  posix_spawn_file_actions_init(&fa);
  posix_spawn_file_actions_addopen(&fa, 2, errfile, O_WRONLY | O_CREAT | O_TRUNC, 0600);

  char *args[] = {(char *)compiler, "-pie", "-fPIC", "-O2",
                  (char *)src, "-o", (char *)out, "-ldl", "-lpthread", NULL};
  pid_t pid;
  int rc = posix_spawnp(&pid, compiler, &fa, NULL, args, environ);
  posix_spawn_file_actions_destroy(&fa);
  if (rc != 0) { unlink(errfile); return false; }  /* compiler not found */

  int status;
  waitpid(pid, &status, 0);
  bool ok = WIFEXITED(status) && WEXITSTATUS(status) == 0;
  if (!ok && errbuf && errn) {
    int f = open(errfile, O_RDONLY);
    if (f >= 0) {
      ssize_t n = read(f, errbuf, errn - 1);
      if (n > 0) errbuf[n] = '\0';
      close(f);
    }
  }
  unlink(errfile);
  return ok;
}

static bool foreign_compile(char exe[PATH_MAX]) {
  char dir[PATH_MAX];
  if (!make_cache_dir(dir)) {
    dlerror_set("dlopen: no writable per-user cache directory for helper");
    return false;
  }

  /* Name the cached binary by a hash of the helper source: an ABI change to the
   * helper yields a new name, so we never execute a stale/mismatched helper. */
  snprintf(exe, PATH_MAX, "%s/helper-%016llx",
           dir, (unsigned long long)fnv1a(HELPER));

  /* Reuse a cached helper only if we own it and it isn't group/other-writable. */
  struct stat st;
  if (stat(exe, &st) == 0 && st.st_uid == getuid() &&
      !(st.st_mode & (S_IWGRP | S_IWOTH)) && (st.st_mode & S_IXUSR)) {
    return true;
  }

  char src[PATH_MAX];
  snprintf(src, sizeof(src), "%s/helper.c", dir);
  int fd = open(src, O_WRONLY | O_CREAT | O_TRUNC, 0600);
  if (fd == -1) { dlerror_set("dlopen: cannot write helper source"); return false; }
  if (write(fd, HELPER, sizeof(HELPER)-1) != sizeof(HELPER)-1) {
    close(fd); unlink(src); dlerror_set("dlopen: short write of helper source"); return false;
  }
  close(fd);

  char tmp[PATH_MAX];
  snprintf(tmp, sizeof(tmp), "%s/helper.tmpXXXXXX", dir);
  int tmpfd = mkstemp(tmp);
  if (tmpfd == -1) { unlink(src); dlerror_set("dlopen: mkstemp failed"); return false; }
  close(tmpfd);

  /* Try the common compiler front-ends in order; capture the last error. */
  static const char *const compilers[] = {"cc", "gcc", "clang"};
  char errbuf[128] = {0};
  bool built = false;
  for (size_t i = 0; i < sizeof(compilers)/sizeof(compilers[0]); i++) {
    if (run_compiler(compilers[i], src, tmp, errbuf, sizeof(errbuf))) { built = true; break; }
  }
  unlink(src);
  if (!built) {
    unlink(tmp);
    if (errbuf[0]) {
      char msg[128];
      snprintf(msg, sizeof(msg), "dlopen: helper compile failed: %s", errbuf);
      dlerror_set(msg);
    } else {
      dlerror_set("dlopen: no C compiler found (cc/gcc/clang) to build helper");
    }
    return false;
  }
  if (rename(tmp, exe) == -1) { unlink(tmp); dlerror_set("dlopen: helper rename failed"); return false; }
  return true;
}

/* Mask saved across the helper-exec window. File scope, not a foreign_setup
 * local, because the helper returns via longjmp: a non-volatile local written
 * between setjmp and longjmp is indeterminate afterwards, and the restore must
 * run in the success branch (longjmp does not restore the signal mask). Setup
 * runs once under pthread_once, so a single static is safe. */
static sigset_t foreign_setup_mask;

static void foreign_setup(void) {
  /* Save our native TLS before executing the helper (it will change TLS). */
  __foreign.native_tls = get_current_tls();

  if (setjmp(__foreign.jb) != 0) {
    /* A helper jumped back to us (foreign_helper): success. Restore the signal
     * mask the longjmp skipped over, then finish. */
    set_current_tls(__foreign.native_tls);
    restore_signals(&foreign_setup_mask);
    if (!__foreign.foreign_tls) {
      dlerror_set("Failed to capture foreign TLS pointer");
      return;
    }
    __foreign.is_supported = true;
    return;
  }

  /* GRAPHICS_GD_DLOPEN_HELPER lets a user/operator pin the helper source when
   * diagnosing a host: "embed" forces the prebuilt helper, "compile" forces a
   * runtime build. Unset (default) tries embedded first, then compiled. */
  const char *force = getenv("GRAPHICS_GD_DLOPEN_HELPER");
  bool allow_embed = !force || strcmp(force, "compile") != 0;
  bool allow_compile = !force || strcmp(force, "embed") != 0;

  /* Running the helper maps ld.so + glibc in-process under foreign TLS, and any
   * async signal whose Go handler runs during that window reads a bogus g. Block
   * them for the whole (cold, one-time) setup; restored in the success branch
   * above, or below if every attempt fails. */
  block_foreign_signals(&foreign_setup_mask);

  /* Attempt 1: an embedded prebuilt helper via an in-memory fd (glibc or musl;
   * no compiler needed, immune to noexec). elf_exec_fd only returns on failure;
   * on success it longjmps to the setjmp above. */
  if (allow_embed) {
    int fd = embedded_helper_fd();
    if (fd >= 0) {
      elf_exec_fd(fd, environ);
      close(fd);
    }
  }

  /* Attempt 2: compile a host-native helper and run it from disk. */
  char exe[PATH_MAX];
  if (allow_compile && foreign_compile(exe)) {
    int cfd = open(exe, O_RDONLY | O_CLOEXEC);
    if (cfd >= 0) {
      elf_exec_fd(cfd, environ);
      close(cfd);
    }
  }

  /* Reached only if every attempt failed (success longjmps past this). */
  restore_signals(&foreign_setup_mask);
  if (!__foreign.is_supported && !dlerror_buf[0])
    dlerror_set("Failed to set up dlopen helper (embedded and compiled paths failed)");
}

static pthread_once_t foreign_once_control = PTHREAD_ONCE_INIT;
static void foreign_once(void) { foreign_setup(); }

static bool foreign_init(void) {
  pthread_once(&foreign_once_control, foreign_once);
  return __foreign.is_supported;
}

/* Public dlfcn API
 *
 * These functions run entirely under the calling thread's glibc TLS. They save
 * and restore the current TLS (rather than unconditionally switching to native
 * at exit) so that if a foreign function calls back into code that calls
 * dlsym(), the outer foreign TLS context is preserved. Async signals are blocked
 * across the switched window so no Go signal handler runs under foreign TLS. */
__attribute__((noinline))
void *dlopen(const char *path, int mode) {
  if (!foreign_init()) return NULL;
  sigset_t old;
  block_foreign_signals(&old);
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(saved_tls));
  void *result = __foreign.dlopen_real(path, mode);
  set_current_tls(saved_tls);
  restore_signals(&old);
  return result;
}

__attribute__((noinline))
void *dlsym(void *handle, const char *name) {
  if (!foreign_init()) return NULL;
  sigset_t old;
  block_foreign_signals(&old);
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(saved_tls));
  void *real_func = __foreign.dlsym_real(handle, name);
  set_current_tls(saved_tls);
  restore_signals(&old);
#ifdef DLOPEN_DEBUG
  fprintf(stderr, "[DLSYM] %s -> %p\n", name, real_func);
#endif
  if (!real_func) return NULL;

  /* Wrap the function pointer so calling it from native (musl) code switches to
   * this thread's glibc TLS for the duration of the call. */
  void *wrapped = foreign_wrap(real_func);
  if (!wrapped) return NULL;

  /* Proc-addr getters need deeper treatment: the pointers THEY return must be
   * wrapped too, so hand out an interposing shim instead of the getter. */
  void *shim = interpose_proc_getter(name, wrapped);
  return shim ? shim : wrapped;
}

int dlclose(void *handle) {
  if (!foreign_init()) return -1;
  sigset_t old;
  block_foreign_signals(&old);
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(saved_tls));
  int result = __foreign.dlclose_real(handle);
  set_current_tls(saved_tls);
  restore_signals(&old);
  return result;
}

char *dlerror(void) {
  if (!foreign_init()) return dlerror_buf;
  sigset_t old;
  block_foreign_signals(&old);
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(saved_tls));
  char *e = __foreign.dlerror_real();
  set_current_tls(saved_tls);
  restore_signals(&old);
  return e ? dlerror_set(e) : NULL;
}

#endif /* __GLIBC__ */
