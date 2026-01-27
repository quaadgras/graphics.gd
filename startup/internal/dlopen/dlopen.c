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
#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/wait.h>
#include <unistd.h>

#define RTLD_LOCAL 0
#define RTLD_LAZY  1
#define RTLD_NOW   2
#define RTLD_GLOBAL 256

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
  "  __asm__ volatile(\"mov %%fs:0, %0\" : \"=r\"(tls));\n" \
  "  return tls;\n" \
  "}\n" \
  "\n" \
  "static void *pool_thread(void *arg) {\n" \
  "  int idx = (int)(intptr_t)arg;\n" \
  "  pthread_mutex_lock(&__tls_pool.lock);\n" \
  "  __tls_pool.tls_ptrs[idx] = get_tls();\n" \
  "  __tls_pool.tramp_ctxs[idx] = &__tramp_ctx;\n" \
  "  __tls_pool.count++;\n" \
  "  if (__tls_pool.count == TLS_POOL_SIZE) sem_post(&__tls_pool.ready);\n" \
  "  pthread_mutex_unlock(&__tls_pool.lock);\n" \
  "  sem_wait(&__tls_pool.shutdown); /* sleep forever */\n" \
  "  return NULL;\n" \
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
  "  for (int i = 1; i < TLS_POOL_SIZE; i++) {\n" \
  "    pthread_t t;\n" \
  "    pthread_create(&t, &attr, pool_thread, (void*)(intptr_t)i);\n" \
  "    pthread_detach(t);\n" \
  "  }\n" \
  "  pthread_attr_destroy(&attr);\n" \
  "  /* Wait for all pool threads to be ready */\n" \
  "  while (__tls_pool.count < TLS_POOL_SIZE) {\n" \
  "    sem_wait(&__tls_pool.ready);\n" \
  "  }\n" \
  "  return ((int (*)(void *))addr)((void *[]){\n" \
  "      dlopen,\n" \
  "      dlsym,\n" \
  "      dlclose,\n" \
  "      dlerror,\n" \
  "      &__tls_pool,\n" \
  "  });\n" \
  "}\n"

struct Loaded {
  char *base;
  char *entry;
  Elf64_Ehdr eh;
  Elf64_Phdr ph[25];
};

static pthread_mutex_t dlopen_mutex = PTHREAD_MUTEX_INITIALIZER;
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
} __foreign;

/* Thread-local slot assignment - must be separate for _Thread_local */
static _Thread_local int my_slot = -1;

/* Get the TLS slot for the current thread, assigning one if needed.
 * This function is called from generated stubs BEFORE TLS switching.
 * It MUST be called with native TLS active (which is the normal case when entering a stub).
 */
int get_thread_slot(void) {
  if (my_slot < 0) {
    my_slot = atomic_fetch_add(&__foreign.next_slot, 1);
    if (my_slot >= TLS_POOL_SIZE) {
      /* Pool exhausted - this is fatal but we can't do much */
      my_slot = 0;  /* Fall back to slot 0 - may cause corruption but won't crash immediately */
    }
  }
  return my_slot;
}

/* Forward declarations for assembly functions */
extern void *foreign_tramp(void);  // Assembly trampoline for TLS switching

/* Helper functions */
static void __dlopen_lock(void) { pthread_mutex_lock(&dlopen_mutex); }
static void __dlopen_unlock(void) { pthread_mutex_unlock(&dlopen_mutex); }

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
  *--sp = 0;                              /* AT_SECURE value */
  *--sp = 23;                             /* AT_SECURE key */
  *--sp = (long)random_bytes;             /* AT_RANDOM value (pointer to 16 random bytes) */
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

static void *foreign_alloc(size_t n) {
  static char *block = NULL;
  static size_t used = 0;
  __dlopen_lock();
  if (!block || used + n > 65536) {
    block = mmap(NULL, 65536, PROT_READ | PROT_WRITE | PROT_EXEC,
                 MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
    if (block == MAP_FAILED) block = NULL;
    used = 4;
  }
  void *p = block ? block + used : NULL;
  if (block) used += (n + 7) & ~7ULL;
  __dlopen_unlock();
  return p;
}

/* Generate a trampoline stub for calling a foreign function through TLS switching.
 *
 * The stub:
 * 1. Saves argument registers (including %r9, the 6th arg)
 * 2. Calls get_thread_slot() to get this thread's TLS pool slot
 * 3. Restores ALL argument registers (including %r9)
 * 4. Pushes slot onto stack for trampoline
 * 5. Loads real_func into %rax
 * 6. Jumps to foreign_tramp (which pops slot from stack)
 *
 * IMPORTANT: We must preserve %r9 because it's the 6th argument register!
 */
__attribute__((noinline))
static void *foreign_wrap(void *real_func) {
  if (!real_func) return NULL;

#ifdef DLOPEN_DEBUG
  fprintf(stderr, "[TRAMP] wrapping function at %p, foreign_tramp at %p\n",
          real_func, (void*)foreign_tramp);
#endif

#ifdef __x86_64__
  /*
   * Generate code that:
   * 1. Allocates stack frame and saves args
   * 2. Calls get_thread_slot()
   * 3. Saves slot to %r10 (caller-saved, not an arg register)
   * 4. Restores ALL args including %r9
   * 5. Deallocates stack frame
   * 6. Pushes slot onto stack (for trampoline to read)
   * 7. Loads real_func and jumps to trampoline
   *
   * Stack alignment: on entry %rsp ≡ 8 (mod 16) after caller's CALL.
   * We sub $56 to make %rsp ≡ 0 (mod 16), so our CALL to get_thread_slot
   * will result in %rsp ≡ 8 on entry to that function (correct for ABI).
   * After restoring and pushing slot, stack is: [slot] [ret_addr] with %rsp ≡ 0 (mod 16).
   */
  unsigned char *stub = foreign_alloc(112);
  if (!stub) return NULL;

  int i = 0;

  /* sub $56, %rsp - allocate stack frame */
  stub[i++] = 0x48; stub[i++] = 0x83; stub[i++] = 0xec; stub[i++] = 56;

  /* Save argument registers to stack */
  /* mov %rdi, 0(%rsp) */
  stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0x3c; stub[i++] = 0x24;
  /* mov %rsi, 8(%rsp) */
  stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0x74; stub[i++] = 0x24; stub[i++] = 0x08;
  /* mov %rdx, 16(%rsp) */
  stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0x54; stub[i++] = 0x24; stub[i++] = 0x10;
  /* mov %rcx, 24(%rsp) */
  stub[i++] = 0x48; stub[i++] = 0x89; stub[i++] = 0x4c; stub[i++] = 0x24; stub[i++] = 0x18;
  /* mov %r8, 32(%rsp) */
  stub[i++] = 0x4c; stub[i++] = 0x89; stub[i++] = 0x44; stub[i++] = 0x24; stub[i++] = 0x20;
  /* mov %r9, 40(%rsp) */
  stub[i++] = 0x4c; stub[i++] = 0x89; stub[i++] = 0x4c; stub[i++] = 0x24; stub[i++] = 0x28;

  /* movabs $get_thread_slot, %rax */
  stub[i++] = 0x48; stub[i++] = 0xb8;
  WRITE64LE(stub + i, (uintptr_t)get_thread_slot);
  i += 8;

  /* call *%rax */
  stub[i++] = 0xff; stub[i++] = 0xd0;

  /* mov %eax, %r10d - save slot in r10 (caller-saved, not arg register) */
  stub[i++] = 0x41; stub[i++] = 0x89; stub[i++] = 0xc2;

  /* Restore ALL argument registers from stack (including %r9!) */
  /* mov 0(%rsp), %rdi */
  stub[i++] = 0x48; stub[i++] = 0x8b; stub[i++] = 0x3c; stub[i++] = 0x24;
  /* mov 8(%rsp), %rsi */
  stub[i++] = 0x48; stub[i++] = 0x8b; stub[i++] = 0x74; stub[i++] = 0x24; stub[i++] = 0x08;
  /* mov 16(%rsp), %rdx */
  stub[i++] = 0x48; stub[i++] = 0x8b; stub[i++] = 0x54; stub[i++] = 0x24; stub[i++] = 0x10;
  /* mov 24(%rsp), %rcx */
  stub[i++] = 0x48; stub[i++] = 0x8b; stub[i++] = 0x4c; stub[i++] = 0x24; stub[i++] = 0x18;
  /* mov 32(%rsp), %r8 */
  stub[i++] = 0x4c; stub[i++] = 0x8b; stub[i++] = 0x44; stub[i++] = 0x24; stub[i++] = 0x20;
  /* mov 40(%rsp), %r9 - RESTORE the 6th argument! */
  stub[i++] = 0x4c; stub[i++] = 0x8b; stub[i++] = 0x4c; stub[i++] = 0x24; stub[i++] = 0x28;

  /* add $56, %rsp - restore stack */
  stub[i++] = 0x48; stub[i++] = 0x83; stub[i++] = 0xc4; stub[i++] = 56;

  /* push %r10 - push slot onto stack for trampoline */
  stub[i++] = 0x41; stub[i++] = 0x52;

  /* movabs $real_func, %rax */
  stub[i++] = 0x48; stub[i++] = 0xb8;
  WRITE64LE(stub + i, (uintptr_t)real_func);
  i += 8;

  /* movabs $foreign_tramp, %r11 */
  stub[i++] = 0x49; stub[i++] = 0xbb;
  WRITE64LE(stub + i, (uintptr_t)foreign_tramp);
  i += 8;

  /* jmp *%r11 */
  stub[i++] = 0x41; stub[i++] = 0xff; stub[i++] = 0xe3;

#ifdef DLOPEN_DEBUG
  fprintf(stderr, "[TRAMP] generated stub at %p, size=%d (first bytes: %02x %02x %02x)\n",
          (void*)stub, i, stub[0], stub[1], stub[2]);
#endif

  return stub;

#elif defined(__aarch64__)
  /*
   * For aarch64, we need to:
   * 1. Save x0-x7 (args) and x30 (lr)
   * 2. Call get_thread_slot
   * 3. Save result to x9
   * 4. Restore x0-x7
   * 5. Load real_func to x8 and jump to trampoline
   */
  unsigned char *stub = foreign_alloc(128);
  if (!stub) return NULL;

  int i = 0;

  /* stp x29, x30, [sp, #-80]! - allocate frame and save fp/lr */
  WRITE32LE(stub + i, 0xa9b57bfd); i += 4;
  /* stp x0, x1, [sp, #16] */
  WRITE32LE(stub + i, 0xa9010fe0); i += 4;
  /* stp x2, x3, [sp, #32] */
  WRITE32LE(stub + i, 0xa90217e2); i += 4;
  /* stp x4, x5, [sp, #48] */
  WRITE32LE(stub + i, 0xa9031fe4); i += 4;
  /* stp x6, x7, [sp, #64] */
  WRITE32LE(stub + i, 0xa90427e6); i += 4;

  /* ldr x16, .Lget_slot - load get_thread_slot address */
  WRITE32LE(stub + i, 0x58000110); i += 4;  /* ldr x16, [pc, #32] - adjusted later */
  /* blr x16 */
  WRITE32LE(stub + i, 0xd63f0200); i += 4;
  /* mov x9, x0 - save slot */
  WRITE32LE(stub + i, 0xaa0003e9); i += 4;

  /* Restore args */
  /* ldp x0, x1, [sp, #16] */
  WRITE32LE(stub + i, 0xa9410fe0); i += 4;
  /* ldp x2, x3, [sp, #32] */
  WRITE32LE(stub + i, 0xa94217e2); i += 4;
  /* ldp x4, x5, [sp, #48] */
  WRITE32LE(stub + i, 0xa9431fe4); i += 4;
  /* ldp x6, x7, [sp, #64] */
  WRITE32LE(stub + i, 0xa94427e6); i += 4;
  /* ldp x29, x30, [sp], #80 - restore and deallocate */
  WRITE32LE(stub + i, 0xa8c57bfd); i += 4;

  /* ldr x8, .Lreal_func */
  WRITE32LE(stub + i, 0x58000068); i += 4;  /* ldr x8, [pc, #12] - adjusted later */
  /* ldr x16, .Ltramp */
  WRITE32LE(stub + i, 0x58000070); i += 4;  /* ldr x16, [pc, #12] - adjusted later */
  /* br x16 */
  WRITE32LE(stub + i, 0xd61f0200); i += 4;

  /* Align to 8 bytes for literal pool */
  while (i & 7) { WRITE32LE(stub + i, 0xd503201f); i += 4; }  /* nop padding */

  /* Literal pool - fix up the ldr offsets above */
  int pool_start = i;
  /* .Lget_slot: */
  WRITE64LE(stub + i, (uintptr_t)get_thread_slot); i += 8;
  /* .Lreal_func: */
  WRITE64LE(stub + i, (uintptr_t)real_func); i += 8;
  /* .Ltramp: */
  WRITE64LE(stub + i, (uintptr_t)foreign_tramp); i += 8;

  /* Fix up ldr instruction offsets (they load from PC + imm*4) */
  /* The ldr x16 for get_thread_slot is at offset 20, pool is at pool_start */
  /* ldr x16, [pc, #offset] where offset = (pool_start - 20) */
  /* Encoding: 0x58 | (imm19 << 5) | Rt, imm19 = offset/4 */
  int off1 = (pool_start - 20) / 4;
  WRITE32LE(stub + 20, 0x58000010 | (off1 << 5));

  /* ldr x8 for real_func is at offset 52, real_func is at pool_start + 8 */
  int off2 = (pool_start + 8 - 52) / 4;
  WRITE32LE(stub + 52, 0x58000008 | (off2 << 5));

  /* ldr x16 for tramp is at offset 56, tramp is at pool_start + 16 */
  int off3 = (pool_start + 16 - 56) / 4;
  WRITE32LE(stub + 56, 0x58000010 | (off3 << 5));

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

/* Get the foreign TLS for the current thread from the pool */
static void *get_thread_foreign_tls(void) {
  int slot = get_thread_slot();
  if (__foreign.pool && slot >= 0 && slot < TLS_POOL_SIZE) {
    return __foreign.pool->tls_ptrs[slot];
  }
  /* Fallback to main thread's TLS */
  return __foreign.foreign_tls;
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
  set_current_tls(get_thread_foreign_tls());
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
  set_current_tls(get_thread_foreign_tls());

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

  return stub;
#else
#error "unsupported architecture"
#endif
}

__attribute__((noinline))
void *dlsym(void *handle, const char *name) {
  if (!foreign_init()) return NULL;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls());
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
  set_current_tls(get_thread_foreign_tls());
  int result = __foreign.dlclose_real(handle);
  set_current_tls(saved_tls);
  return result;
}

char *dlerror(void) {
  if (!foreign_init()) return dlerror_buf;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls());
  char *e = __foreign.dlerror_real();
  set_current_tls(saved_tls);
  return e ? dlerror_set(e) : NULL;
}

/* Get raw function pointer without wrapping (for use with manual TLS switching) */
__attribute__((noinline))
void *dlsym_raw(void *handle, const char *name) {
  if (!foreign_init()) return NULL;
  void *saved_tls = get_current_tls();
  set_current_tls(get_thread_foreign_tls());
  void *real_func = __foreign.dlsym_real(handle, name);
  set_current_tls(saved_tls);
  return real_func;
}

/* Switch to foreign TLS (call before using foreign libraries extensively) */
void dlopen_set_foreign_tls(void) {
  if (!foreign_init()) return;
  set_current_tls(get_thread_foreign_tls());
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
