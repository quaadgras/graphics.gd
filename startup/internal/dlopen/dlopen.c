/* dlopen.c - Cosmopolitan-style dlopen port for static musl binaries
 *
 * Supports loading both musl-built and glibc-built shared libraries on Linux.
 * Uses the helper executable trick to borrow the system's dynamic loader.
 * Implements TLS switching to allow foreign libraries to use their own TLS.
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

#define HELPER \
  "#define _GNU_SOURCE\n" \
  "#include <dlfcn.h>\n" \
  "#include <stdio.h>\n" \
  "#include <stdlib.h>\n" \
  "#include <pthread.h>\n" \
  "#include <semaphore.h>\n" \
  "#include <stdint.h>\n" \
  "\n" \
  "#define TLS_POOL_SIZE 64\n" \
  "\n" \
  "/* TLS stack for nested trampoline calls (per-thread). */\n" \
  "__thread struct {\n" \
  "  long sp;\n" \
  "  void *stack[32];\n" \
  "} __tramp_ctx;\n" \
  "\n" \
  "/* TLS pool: array of glibc TLS pointers, one per pool thread. */\n" \
  "struct tls_pool {\n" \
  "  void *tls_ptrs[TLS_POOL_SIZE];      /* glibc TLS pointers */\n" \
  "  void *tramp_ctxs[TLS_POOL_SIZE];    /* per-thread __tramp_ctx addresses */\n" \
  "  sem_t ready;                         /* signaled when all threads ready */\n" \
  "  sem_t shutdown;                      /* signaled to shut down threads */\n" \
  "  int count;                           /* number of threads initialized */\n" \
  "  pthread_mutex_t lock;\n" \
  "} __tls_pool;\n" \
  "\n" \
  "static void *get_tls(void) {\n" \
  "  void *tls;\n" \
  "#if defined(__x86_64__)\n" \
  "  __asm__ volatile(\"mov %%fs:0, %0\" : \"=r\"(tls));\n" \
  "#elif defined(__aarch64__)\n" \
  "  __asm__ volatile(\"mrs %0, tpidr_el0\" : \"=r\"(tls));\n" \
  "#else\n" \
  "#error \"unsupported architecture\"\n" \
  "#endif\n" \
  "  return tls;\n" \
  "}\n" \
  "\n" \
  "static void *pool_thread(void *arg) {\n" \
  "  int idx = (int)(intptr_t)arg;\n" \
  "  pthread_mutex_lock(&__tls_pool.lock);\n" \
  "  __tls_pool.tls_ptrs[idx] = get_tls();\n" \
  "  __tls_pool.tramp_ctxs[idx] = &__tramp_ctx;\n" \
  "  __tls_pool.count++;\n" \
  "  pthread_mutex_unlock(&__tls_pool.lock);\n" \
  "  sem_post(&__tls_pool.ready); /* signal this thread recorded its TLS */\n" \
  "  sem_wait(&__tls_pool.shutdown); /* sleep forever */\n" \
  "  return NULL;\n" \
  "}\n" \
  "\n" \
  "/* On-demand glibc TCB factory: spawn a parked glibc thread and return its\n" \
  "   TLS pointer, so the musl side can hand fresh glibc TCBs to threads beyond\n" \
  "   the pre-spawned pool (no fixed cap). The musl side serializes calls, so the\n" \
  "   statics below need no locking. */\n" \
  "static sem_t __tcb_ready;\n" \
  "static void *__tcb_captured;\n" \
  "static void *tcb_thread(void *arg) {\n" \
  "  (void)arg;\n" \
  "  __tcb_captured = get_tls();\n" \
  "  sem_post(&__tcb_ready);\n" \
  "  sem_wait(&__tls_pool.shutdown); /* park forever */\n" \
  "  return NULL;\n" \
  "}\n" \
  "void *glibc_tcb_create(void) {\n" \
  "  pthread_attr_t attr;\n" \
  "  pthread_attr_init(&attr);\n" \
  "  pthread_attr_setstacksize(&attr, 16384);\n" \
  "  sem_init(&__tcb_ready, 0, 0);\n" \
  "  pthread_t t;\n" \
  "  int rc = pthread_create(&t, &attr, tcb_thread, NULL);\n" \
  "  pthread_attr_destroy(&attr);\n" \
  "  if (rc != 0) return NULL;\n" \
  "  pthread_detach(t);\n" \
  "  sem_wait(&__tcb_ready);\n" \
  "  return __tcb_captured;\n" \
  "}\n" \
  "\n" \
  "int main(int argc, char **argv, char **envp) {\n" \
  "  char *ep;\n" \
  "  long addr;\n" \
  "  if (argc != 2) {\n" \
  "    fprintf(stderr, \"%s: not intended to be run directly\\n\", argv[0]);\n" \
  "    return 1;\n" \
  "  }\n" \
  "  addr = strtol(argv[1], &ep, 10);\n" \
  "  if (*ep) {\n" \
  "    fprintf(stderr, \"%s: invalid function address\\n\", argv[0]);\n" \
  "    return 2;\n" \
  "  }\n" \
  "  /* Initialize TLS pool */\n" \
  "  sem_init(&__tls_pool.ready, 0, 0);\n" \
  "  sem_init(&__tls_pool.shutdown, 0, 0);\n" \
  "  pthread_mutex_init(&__tls_pool.lock, NULL);\n" \
  "  __tls_pool.count = 0;\n" \
  "  /* Slot 0 is for main thread */\n" \
  "  __tls_pool.tls_ptrs[0] = get_tls();\n" \
  "  __tls_pool.tramp_ctxs[0] = &__tramp_ctx;\n" \
  "  __tls_pool.count = 1;\n" \
  "  /* Create pool threads */\n" \
  "  pthread_attr_t attr;\n" \
  "  pthread_attr_init(&attr);\n" \
  "  pthread_attr_setstacksize(&attr, 16384); /* minimal stack */\n" \
  "  int created = 0;\n" \
  "  for (int i = 1; i < TLS_POOL_SIZE; i++) {\n" \
  "    pthread_t t;\n" \
  "    if (pthread_create(&t, &attr, pool_thread, (void*)(intptr_t)i) == 0) {\n" \
  "      pthread_detach(t);\n" \
  "      created++;\n" \
  "    }\n" \
  "  }\n" \
  "  pthread_attr_destroy(&attr);\n" \
  "  /* Wait only for the threads that were actually created, so a failed\n" \
  "     pthread_create (e.g. RLIMIT_NPROC in a container) cannot hang us\n" \
  "     forever. Unfilled slots keep NULL TLS pointers; the musl side's\n" \
  "     get_thread_slot() aborts cleanly if one is ever assigned. */\n" \
  "  for (int i = 0; i < created; i++) sem_wait(&__tls_pool.ready);\n" \
  "  if (created < TLS_POOL_SIZE - 1)\n" \
  "    fprintf(stderr, \"dlopen helper: only %d of %d TLS pool threads \"\n" \
  "                    \"created; foreign calls limited to %d threads\\n\",\n" \
  "            created, TLS_POOL_SIZE - 1, created + 1);\n" \
  "  return ((int (*)(void *))addr)((void *[]){\n" \
  "      dlopen,\n" \
  "      dlsym,\n" \
  "      dlclose,\n" \
  "      dlerror,\n" \
  "      &__tls_pool,\n" \
  "      glibc_tcb_create,\n" \
  "  });\n" \
  "}\n"

struct Loaded {
  char *base;
  char *entry;
  Elf64_Ehdr eh;
  Elf64_Phdr ph[25];
};

static _Thread_local char dlerror_buf[128];

/* TLS pool structure (must match helper's struct tls_pool) */
struct tls_pool {
  void *tls_ptrs[TLS_POOL_SIZE];      /* glibc TLS pointers */
  void *tramp_ctxs[TLS_POOL_SIZE];    /* per-thread __tramp_ctx addresses */
  /* sem_t ready, shutdown - we don't need these after init */
  /* int count, mutex - we don't need these after init */
};

/*
 * __foreign struct layout - offsets must match foreign_tramp.S:
 *   offset 0:  atomic_uint once
 *   offset 8:  void *pool (TLS pool pointer)
 *   offset 16: void *foreign_tls (main thread's, for compat)
 *   offset 24: void *native_tls (main thread's)
 *   offset 32: bool is_supported
 */
struct {
  atomic_uint once;
  struct tls_pool *pool;  // offset 8: TLS pool for multi-threaded support
  void *foreign_tls;      // offset 16: foreign TLS pointer (main thread's, for compat)
  void *native_tls;       // offset 24: native TLS pointer (main thread's)
  bool is_supported;      // offset 32
  void *(*dlopen_real)(const char *, int);
  void *(*dlsym_real)(void *, const char *);
  int (*dlclose_real)(void *);
  char *(*dlerror_real)(void);
  void *(*dlopen)(const char *, int);
  void *(*dlsym)(void *, const char *);
  int (*dlclose)(void *);
  char *(*dlerror)(void);
  jmp_buf jb;
  /* Thread-to-slot mapping */
  atomic_int next_slot;
  /* Factory (in the helper's glibc context) that spawns a parked glibc thread
   * and returns its TLS pointer, used to grow past the pre-spawned pool. */
  void *(*glibc_tcb_create)(void);
} __foreign;

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

static size_t my_strlcat(char *dst, const char *src, size_t siz) {
  size_t len = strlen(dst);
  if (len < siz) my_strlcpy(dst + len, src, siz - len);
  return len + strlen(src);
}

static const char *get_tmp_dir(void) {
  const char *t = getenv("TMPDIR");
  if (!t || !*t) t = getenv("HOME");
  if (!t || !*t) t = ".";
  return t;
}

static int timespec_cmp(const struct timespec *a, const struct timespec *b) {
  if (a->tv_sec != b->tv_sec) return a->tv_sec > b->tv_sec ? 1 : -1;
  if (a->tv_nsec != b->tv_nsec) return a->tv_nsec > b->tv_nsec ? 1 : -1;
  return 0;
}

static int is_file_newer_than(const char *path, const char *other) {
  struct stat st1, st2;
  if (stat(path, &st1)) return -1;
  if (stat(other, &st2)) return errno == ENOENT ? 2 : -1;
  return timespec_cmp(&st1.st_mtim, &st2.st_mtim) > 0;
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

static bool elf_load(struct Loaded *l, const char *file, long pagesz,
                     char *interp_path, size_t interp_size) {
  int fd = open(file, O_RDONLY | O_CLOEXEC);
  if (fd == -1) return false;

  if (pread(fd, &l->eh, sizeof(l->eh), 0) != sizeof(l->eh) ||
      !is_elf64(&l->eh) ||
      l->eh.e_phnum > sizeof(l->ph)/sizeof(l->ph[0]) ||
      l->eh.e_machine != get_host_elf_machine()) {
    close(fd);
    errno = ENOEXEC;
    return false;
  }

  if (pread(fd, l->ph, l->eh.e_phnum * sizeof(l->ph[0]), l->eh.e_phoff) !=
      (ssize_t)(l->eh.e_phnum * sizeof(l->ph[0]))) {
    close(fd);
    return false;
  }

  l->base = elf_map(fd, &l->eh, l->ph, pagesz, interp_path, interp_size);
  close(fd);
  if (l->base == MAP_FAILED) return false;

  l->entry = l->base + l->eh.e_entry;
  return true;
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

static void elf_exec(const char *file, char **envp) {
  long pagesz = sysconf(_SC_PAGESIZE);
  if (pagesz <= 0) pagesz = 4096;

  struct Loaded prog;
  char interp_path[256] = {0};
  if (!elf_load(&prog, file, pagesz, interp_path, sizeof(interp_path))) return;

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
 * calls the real function, and restores the caller's TLS. Keeping the logic in
 * one hand-written assembly routine (foreign_tramp.S) instead of per-stub
 * generated bytes is both simpler and avoids the previous stub's bugs (it never
 * saved the FP argument registers around its helper call, and it read a
 * _Thread_local under possibly-foreign TLS). All argument registers, the
 * variadic count (%al / none on aarch64) and the struct-return pointer flow
 * through untouched. */
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

static bool foreign_compile(char exe[PATH_MAX]) {
  my_strlcpy(exe, get_tmp_dir(), PATH_MAX);
  my_strlcat(exe, "/.musl_dlopen_helper", PATH_MAX);
  if (mkdir(exe, 0755) && errno != EEXIST) return false;
  my_strlcat(exe, "/helper", PATH_MAX);

  switch (is_file_newer_than(get_program_executable_name(), exe)) {
    case 0: return true;
    case 1: case 2: break;
    default: return false;
  }

  char src[PATH_MAX];
  my_strlcpy(src, exe, PATH_MAX);
  my_strlcat(src, ".c", PATH_MAX);

  int fd = open(src, O_WRONLY | O_CREAT | O_TRUNC, 0600);
  if (fd == -1) return false;
  if (write(fd, HELPER, sizeof(HELPER)-1) != sizeof(HELPER)-1) {
    close(fd); unlink(src); return false;
  }
  close(fd);

  char tmp[PATH_MAX];
  my_strlcpy(tmp, exe, PATH_MAX);
  my_strlcat(tmp, ".tmpXXXXXX", PATH_MAX);
  int tmpfd = mkstemp(tmp);
  if (tmpfd == -1) { unlink(src); return false; }
  close(tmpfd);

  char *args[] = {"cc", "-pie", "-fPIC", src, "-o", tmp, "-ldl", NULL};
  pid_t pid;
  int status;
  if (posix_spawnp(&pid, "cc", NULL, NULL, args, environ) != 0) {
    unlink(tmp); unlink(src); return false;
  }
  waitpid(pid, &status, 0);
  unlink(src);
  if (status != 0) { unlink(tmp); return false; }
  if (rename(tmp, exe) == -1) { unlink(tmp); return false; }
  return true;
}

static void foreign_setup(void) {
  char exe[PATH_MAX];
  if (!foreign_compile(exe)) {
    dlerror_set("Failed to compile dlopen helper");
    return;
  }

  /* Save our native TLS before executing the helper (it will change TLS) */
  __foreign.native_tls = get_current_tls();

  if (setjmp(__foreign.jb) == 0) {
    elf_exec(exe, environ);
    dlerror_set("Failed to execute dlopen helper");
    return;
  }

  /* Restore our native TLS - the helper's interpreter changed it */
  set_current_tls(__foreign.native_tls);

  /* Sanity check: make sure we captured the foreign TLS */
  if (!__foreign.foreign_tls) {
    dlerror_set("Failed to capture foreign TLS pointer");
    return;
  }

  __foreign.is_supported = true;
}

static pthread_once_t foreign_once_control = PTHREAD_ONCE_INIT;
static void foreign_once(void) { foreign_setup(); }

static bool foreign_init(void) {
  pthread_once(&foreign_once_control, foreign_once);
  return __foreign.is_supported;
}

/* Public dlfcn API
 *
 * These functions save and restore the current TLS rather than unconditionally
 * switching to native TLS at exit. This is critical for callbacks: if a foreign
 * function calls back into code that calls dlsym(), we must preserve foreign
 * TLS context rather than corrupting it by switching to native.
 */
__attribute__((noinline))
void *dlopen(const char *path, int mode) {
  if (!foreign_init()) return NULL;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(get_current_tls()));
  void *result = __foreign.dlopen_real(path, mode);
  set_current_tls(saved_tls);
  return result;
}

/*
 * Check if a function name is a "GetProcAddress" style function that returns
 * function pointers. These need special handling - we must wrap the returned
 * function pointer, not just the GetProcAddress function itself.
 */
static bool is_procaddr_function(const char *name) {
  return strcmp(name, "glXGetProcAddressARB") == 0 ||
         strcmp(name, "glXGetProcAddress") == 0 ||
         strcmp(name, "eglGetProcAddress") == 0 ||
         strcmp(name, "wlEglGetProcAddress") == 0 ||
         strcmp(name, "vkGetInstanceProcAddr") == 0 ||
         strcmp(name, "vkGetDeviceProcAddr") == 0 ||
         strcmp(name, "SDL_GL_GetProcAddress") == 0;
}

/*
 * Wrapper for GetProcAddress-style functions.
 * This wraps the returned function pointer so it switches TLS when called.
 *
 * We generate a trampoline that:
 * 1. Calls the real GetProcAddress with TLS switching
 * 2. Wraps the returned function pointer
 *
 * Since we can't easily generate code that does this, we use a different
 * approach: we return a stub that captures the real function and wraps
 * its return value. This requires generating custom code for each lookup.
 */
typedef void *(*procaddr_fn)(const char *);
typedef void *(*procaddr_fn2)(void *, const char *);

/* Storage for GetProcAddress wrappers - we need to track the real function */
struct procaddr_wrapper {
  void *real_func;
  bool has_handle;  /* true if function takes (handle, name), false if just (name) */
};

#define MAX_PROCADDR_WRAPPERS 16
static struct procaddr_wrapper procaddr_wrappers[MAX_PROCADDR_WRAPPERS];
static int procaddr_wrapper_count = 0;

/* Forward declaration */
static void *create_procaddr_wrapper(void *real_func, bool has_handle);

/*
 * Generic wrapper that calls a GetProcAddress function and wraps the result.
 * The wrapper index is encoded in the stub.
 */
static void *call_procaddr_and_wrap(int idx, void *handle, const char *name) {
  if (idx < 0 || idx >= procaddr_wrapper_count) return NULL;
  struct procaddr_wrapper *w = &procaddr_wrappers[idx];

  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(get_current_tls()));

  void *func;
  if (w->has_handle) {
    func = ((procaddr_fn2)w->real_func)(handle, name);
  } else {
    func = ((procaddr_fn)w->real_func)(name);
  }

  set_current_tls(saved_tls);

  if (!func) return NULL;

#ifdef DLOPEN_DEBUG
  fprintf(stderr, "[PROCADDR] wrapper %d(\"%s\") -> %p", idx, name, func);
#endif

  /*
   * If the returned function is ALSO a procaddr function, wrap it specially.
   * This handles cases like glXGetProcAddressARB("glXGetProcAddress").
   */
  if (is_procaddr_function(name)) {
    bool has_handle = (strcmp(name, "vkGetInstanceProcAddr") == 0 ||
                       strcmp(name, "vkGetDeviceProcAddr") == 0);
#ifdef DLOPEN_DEBUG
    fprintf(stderr, " (procaddr, has_handle=%d)\n", has_handle);
#endif
    return create_procaddr_wrapper(func, has_handle);
  }

#ifdef DLOPEN_DEBUG
  fprintf(stderr, " (wrapping)\n");
#endif
  return foreign_wrap(func);
}

/* Generate wrapper stub for GetProcAddress-style function */
static void *create_procaddr_wrapper(void *real_func, bool has_handle) {
  if (procaddr_wrapper_count >= MAX_PROCADDR_WRAPPERS) return NULL;

  int idx = procaddr_wrapper_count++;
  procaddr_wrappers[idx].real_func = real_func;
  procaddr_wrappers[idx].has_handle = has_handle;

#ifdef __x86_64__
  /*
   * Generate a stub that calls call_procaddr_and_wrap(idx, handle, name)
   * For has_handle=false: rdi=name, we pass (idx, NULL, name)
   * For has_handle=true: rdi=handle, rsi=name, we pass (idx, handle, name)
   *
   * Code:
   *   mov %rsi, %rdx        ; name -> arg3 (or for !has_handle: mov %rdi, %rdx)
   *   mov %rdi, %rsi        ; handle -> arg2 (or for !has_handle: xor %esi, %esi)
   *   mov $idx, %edi        ; idx -> arg1
   *   movabs $call_procaddr_and_wrap, %rax
   *   jmp *%rax
   */
  unsigned char *stub = foreign_alloc(32);
  if (!stub) return NULL;

  int i = 0;
  if (has_handle) {
    /* mov %rsi, %rdx */
    stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0xf2;
    /* mov %rdi, %rsi */
    stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0xfe;
  } else {
    /* mov %rdi, %rdx (name is in rdi for single-arg version) */
    stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0xfa;
    /* xor %esi, %esi (handle = NULL) */
    stub[i++] = 0x31; stub[i++] = 0xf6;
  }
  /* mov $idx, %edi */
  stub[i++] = 0xbf;
  WRITE32LE(stub + i, idx);
  i += 4;
  /* movabs $call_procaddr_and_wrap, %rax */
  stub[i++] = 0x48; stub[i++] = 0xb8;
  WRITE64LE(stub + i, (uintptr_t)call_procaddr_and_wrap);
  i += 8;
  /* jmp *%rax */
  stub[i++] = 0xff; stub[i++] = 0xe0;

  if (!foreign_seal(stub, i)) return NULL;
  return stub;
#elif defined(__aarch64__)
  /* Similar for aarch64 - adjust register shuffling */
  unsigned char *stub = foreign_alloc(48);
  if (!stub) return NULL;

  /* For aarch64, args are in x0, x1, x2...
   * We need: x0=idx, x1=handle (or NULL), x2=name
   */
  int i = 0;
  if (has_handle) {
    /* mov x2, x1 (name -> x2) */
    WRITE32LE(stub + i, 0xaa0103e2); i += 4;
    /* mov x1, x0 (handle -> x1) */
    WRITE32LE(stub + i, 0xaa0003e1); i += 4;
  } else {
    /* mov x2, x0 (name -> x2) */
    WRITE32LE(stub + i, 0xaa0003e2); i += 4;
    /* mov x1, xzr (handle = NULL) */
    WRITE32LE(stub + i, 0xaa1f03e1); i += 4;
  }
  /* mov x0, #idx */
  WRITE32LE(stub + i, 0xd2800000 | (idx << 5)); i += 4;
  /* ldr x16, [pc, #12] */
  WRITE32LE(stub + i, 0x58000070); i += 4;
  /* br x16 */
  WRITE32LE(stub + i, 0xd61f0200); i += 4;
  /* nop (padding) */
  WRITE32LE(stub + i, 0xd503201f); i += 4;
  /* literal: call_procaddr_and_wrap address */
  WRITE64LE(stub + i, (uintptr_t)call_procaddr_and_wrap);

  /* seal i+8 bytes to cover the trailing 64-bit literal as well */
  if (!foreign_seal(stub, i + 8)) return NULL;
  return stub;
#else
#error "unsupported architecture"
#endif
}

__attribute__((noinline))
void *dlsym(void *handle, const char *name) {
  if (!foreign_init()) return NULL;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(get_current_tls()));
  void *real_func = __foreign.dlsym_real(handle, name);
  set_current_tls(saved_tls);
#ifdef DLOPEN_DEBUG
  fprintf(stderr, "[DLSYM] %s -> %p\n", name, real_func);
#endif
  if (!real_func) return NULL;

  /*
   * Special handling for GetProcAddress-style functions.
   * These return function pointers that also need to be wrapped.
   */
  if (is_procaddr_function(name)) {
    /* Determine if it takes a handle argument */
    bool has_handle = (strcmp(name, "vkGetInstanceProcAddr") == 0 ||
                       strcmp(name, "vkGetDeviceProcAddr") == 0);
#ifdef DLOPEN_DEBUG
    fprintf(stderr, "[DLSYM] %s is a procaddr function (has_handle=%d)\n", name, has_handle);
#endif
    return create_procaddr_wrapper(real_func, has_handle);
  }

  /* Wrap the function pointer with TLS switching trampoline */
  return foreign_wrap(real_func);
}

int dlclose(void *handle) {
  if (!foreign_init()) return -1;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(get_current_tls()));
  int result = __foreign.dlclose_real(handle);
  set_current_tls(saved_tls);
  return result;
}

char *dlerror(void) {
  if (!foreign_init()) return dlerror_buf;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(get_current_tls()));
  char *e = __foreign.dlerror_real();
  set_current_tls(saved_tls);
  return e ? dlerror_set(e) : NULL;
}

/* Get raw function pointer without wrapping (for use with manual TLS switching) */
__attribute__((noinline))
void *dlsym_raw(void *handle, const char *name) {
  if (!foreign_init()) return NULL;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls(get_current_tls()));
  void *real_func = __foreign.dlsym_real(handle, name);
  set_current_tls(saved_tls);
  return real_func;
}

/* Switch to foreign TLS (call before using foreign libraries extensively) */
void dlopen_set_foreign_tls(void) {
  if (!foreign_init()) return;
  set_current_tls(get_thread_foreign_tls(get_current_tls()));
}

/* Switch back to native TLS (call when done with foreign libraries) */
void dlopen_set_native_tls(void) {
  if (!foreign_init()) return;
  set_current_tls(__foreign.native_tls);
}

/*
 * Callback TLS recovery functions.
 *
 * Use these when a foreign library calls back into native code that needs
 * native TLS. The callback MUST call dlopen_callback_exit() before returning
 * to restore foreign TLS, otherwise the foreign library will crash.
 *
 * Usage:
 *   void my_callback(void *data) {
 *       void *saved = dlopen_callback_enter();
 *       // ... native code runs with native TLS ...
 *       dlopen_callback_exit(saved);
 *   }
 */

/* Enter callback: save current TLS, switch to native TLS. Returns saved TLS. */
void *dlopen_callback_enter(void) {
  if (!__foreign.is_supported) return NULL;
  void *saved = get_current_tls();
  set_current_tls(__foreign.native_tls);
  return saved;
}

/* Exit callback: restore the TLS that was saved at callback entry. */
void dlopen_callback_exit(void *saved_tls) {
  if (!__foreign.is_supported || !saved_tls) return;
  set_current_tls(saved_tls);
}

#endif /* __GLIBC__ */
