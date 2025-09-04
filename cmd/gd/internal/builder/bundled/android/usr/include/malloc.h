/*
 * Copyright (C) 2012 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

/**
 * @file malloc.h
 * @brief Heap memory allocation.
 *
 * [Debugging Native Memory Use](https://source.android.com/devices/tech/debug/native-memory)
 * is the canonical source for documentation on Android's heap debugging
 * features.
 */

#include <sys/cdefs.h>
#include <stddef.h>
#include <stdio.h>

__BEGIN_DECLS

#define __BIONIC_ALLOC_SIZE(...) __attribute__((__alloc_size__(__VA_ARGS__)))

/**
 * [malloc(3)](https://man7.org/linux/man-pages/man3/malloc.3.html) allocates
 * memory on the heap.
 *
 * Returns a pointer to the allocated memory on success and returns a null
 * pointer and sets `errno` on failure.
 *
 * Note that Android (like most Unix systems) allows "overcommit". This
 * allows processes to allocate more memory than the system has, provided
 * they don't use it all. This works because only "dirty" pages that have
 * been written to actually require physical memory. In practice, this
 * means that it's rare to see memory allocation functions return a null
 * pointer, and that a non-null pointer does not mean that you actually
 * have all of the memory you asked for.
 *
 * Note also that the Linux Out Of Memory (OOM) killer behaves differently
 * for code run via `adb shell`. The assumption is that if you ran
 * something via `adb shell` you're a developer who actually wants the
 * device to do what you're asking it to do _even if_ that means killing
 * other processes. Obviously this is not the case for apps, which will
 * be killed in preference to killing other processes.
 */
__nodiscard void* _Nullable malloc(size_t __byte_count) __mallocfunc __BIONIC_ALLOC_SIZE(1);

/**
 * [calloc(3)](https://man7.org/linux/man-pages/man3/calloc.3.html) allocates
 * and clears memory on the heap.
 *
 * Returns a pointer to the allocated memory on success and returns a null
 * pointer and sets `errno` on failure (but see the notes for malloc()).
 */
__nodiscard void* _Nullable calloc(size_t __item_count, size_t __item_size) __mallocfunc __BIONIC_ALLOC_SIZE(1,2);

/**
 * [realloc(3)](https://man7.org/linux/man-pages/man3/realloc.3.html) resizes
 * allocated memory on the heap.
 *
 * Returns a pointer (which may be different from `__ptr`) to the resized
 * memory on success and returns a null pointer and sets `errno` on failure
 * (but see the notes for malloc()).
 */
__nodiscard void* _Nullable realloc(void* _Nullable __ptr, size_t __byte_count) __BIONIC_ALLOC_SIZE(2);

/**
 * [reallocarray(3)](https://man7.org/linux/man-pages/man3/reallocarray.3.html)
 * resizes allocated memory on the heap.
 *
 * Equivalent to `realloc(__ptr, __item_count * __item_size)` but fails if the
 * multiplication overflows.
 *
 * Returns a pointer (which may be different from `__ptr`) to the resized
 * memory on success and returns a null pointer and sets `errno` on failure
 * (but see the notes for malloc()).
 */
#if __ANDROID_API__ >= 29
__nodiscard void* _Nullable reallocarray(void* _Nullable __ptr, size_t __item_count, size_t __item_size) __BIONIC_ALLOC_SIZE(2, 3) __INTRODUCED_IN(29);
#elif defined(__ANDROID_UNAVAILABLE_SYMBOLS_ARE_WEAK__)
#include <errno.h>
static __inline __nodiscard void* _Nullable reallocarray(void* _Nullable __ptr, size_t __item_count, size_t __item_size) __BIONIC_ALLOC_SIZE(2, 3) {
  size_t __new_size;
  if (__builtin_mul_overflow(__item_count, __item_size, &__new_size)) {
    errno = ENOMEM;
    return NULL;
  }
  return realloc(__ptr, __new_size);
}
#endif

/**
 * [free(3)](https://man7.org/linux/man-pages/man3/free.3.html) deallocates
 * memory on the heap.
 */
void free(void* _Nullable __ptr);

/**
 * [memalign(3)](https://man7.org/linux/man-pages/man3/memalign.3.html) allocates
 * memory on the heap with the required alignment.
 *
 * Returns a pointer to the allocated memory on success and returns a null
 * pointer and sets `errno` on failure (but see the notes for malloc()).
 *
 * See also posix_memalign().
 */
__nodiscard void* _Nullable memalign(size_t __alignment, size_t __byte_count) __mallocfunc __BIONIC_ALLOC_SIZE(2);

/**
 * [malloc_usable_size(3)](https://man7.org/linux/man-pages/man3/malloc_usable_size.3.html)
 * returns the actual size of the given heap block.
 */
__nodiscard size_t malloc_usable_size(const void* _Nullable __ptr);

#define __MALLINFO_BODY \
  /** Total number of non-mmapped bytes currently allocated from OS. */ \
  size_t arena; \
  /** Number of free chunks. */ \
  size_t ordblks; \
  /** (Unused.) */ \
  size_t smblks; \
  /** (Unused.) */ \
  size_t hblks; \
  /** Total number of bytes in mmapped regions. */ \
  size_t hblkhd; \
  /** Maximum total allocated space; greater than total if trimming has occurred. */ \
  size_t usmblks; \
  /** (Unused.) */ \
  size_t fsmblks; \
  /** Total allocated space (normal or mmapped.) */ \
  size_t uordblks; \
  /** Total free space. */ \
  size_t fordblks; \
  /** Upper bound on number of bytes releasable by a trim operation. */ \
  size_t keepcost;

#ifndef STRUCT_MALLINFO_DECLARED
#define STRUCT_MALLINFO_DECLARED 1
struct mallinfo { __MALLINFO_BODY };
#endif

/**
 * [mallinfo(3)](https://man7.org/linux/man-pages/man3/mallinfo.3.html) returns
 * information about the current state of the heap. Note that mallinfo() is
 * inherently unreliable and consider using malloc_info() instead.
 */
struct mallinfo mallinfo(void);

/**
 * On Android the struct mallinfo and struct mallinfo2 are the same.
 */
struct mallinfo2 { __MALLINFO_BODY };

/**
 * [mallinfo2(3)](https://man7.org/linux/man-pages/man3/mallinfo2.3.html) returns
 * information about the current state of the heap. Note that mallinfo2() is
 * inherently unreliable and consider using malloc_info() instead.
 */
struct mallinfo2 mallinfo2(void) __RENAME(mallinfo);

/**
 * [malloc_info(3)](https://man7.org/linux/man-pages/man3/malloc_info.3.html)
 * writes information about the current state of the heap to the given stream.
 *
 * The XML structure for malloc_info() is as follows:
 * ```
 * <malloc version="jemalloc-1">
 *   <heap nr="INT">
 *     <allocated-large>INT</allocated-large>
 *     <allocated-huge>INT</allocated-huge>
 *     <allocated-bins>INT</allocated-bins>
 *     <bins-total>INT</bins-total>
 *     <bin nr="INT">
 *       <allocated>INT</allocated>
 *       <nmalloc>INT</nmalloc>
 *       <ndalloc>INT</ndalloc>
 *     </bin>
 *     <!-- more bins -->
 *   </heap>
 *   <!-- more heaps -->
 * </malloc>
 * ```
 *
 * Available since API level 23.
 */

#if __BIONIC_AVAILABILITY_GUARD(23)
int malloc_info(int __must_be_zero, FILE* _Nonnull __fp) __INTRODUCED_IN(23);
#endif /* __BIONIC_AVAILABILITY_GUARD(23) */


/**
 * mallopt() option to set the decay time. Valid values are -1, 0 and 1.
 *   -1 : Disable the releasing of unused pages. This value is available since
 *        API level 35.
 *    0 : Release the unused pages immediately.
 *    1 : Release the unused pages at a device-specific interval.
 *
 * Available since API level 27.
 */
#define M_DECAY_TIME (-100)
/**
 * mallopt() option to immediately purge any memory not in use. This
 * will release the memory back to the kernel. The value is ignored.
 *
 * Available since API level 28.
 */
#define M_PURGE (-101)
/**
 * mallopt() option to immediately purge all possible memory back to
 * the kernel. This call can take longer than a normal purge since it
 * examines everything. In some cases, it can take more than twice the
 * time of a M_PURGE call. The value is ignored.
 *
 * Available since API level 34.
 */
#define M_PURGE_ALL (-104)

/**
 * mallopt() option to tune the allocator's choice of memory tags to
 * make it more likely that a certain class of memory errors will be
 * detected. This is only relevant if MTE is enabled in this process
 * and ignored otherwise. The value argument should be one of the
 * M_MEMTAG_TUNING_* flags.
 * NOTE: This is only available in scudo.
 *
 * Available since API level 31.
 */
#define M_MEMTAG_TUNING (-102)

/**
 * When passed as a value of M_MEMTAG_TUNING mallopt() call, enables
 * deterministic detection of linear buffer overflow and underflow
 * bugs by assigning distinct tag values to adjacent allocations. This
 * mode has a slightly reduced chance to detect use-after-free bugs
 * because only half of the possible tag values are available for each
 * memory location.
 *
 * Please keep in mind that MTE can not detect overflow within the
 * same tag granule (16-byte aligned chunk), and can miss small
 * overflows even in this mode. Such overflow can not be the cause of
 * a memory corruption, because the memory within one granule is never
 * used for multiple allocations.
 */
#define M_MEMTAG_TUNING_BUFFER_OVERFLOW 0

/**
 * When passed as a value of M_MEMTAG_TUNING mallopt() call, enables
 * independently randomized tags for uniform ~93% probability of
 * detecting both spatial (buffer overflow) and temporal (use after
 * free) bugs.
 */
#define M_MEMTAG_TUNING_UAF 1

/**
 * mallopt() option for per-thread memory initialization tuning.
 * The value argument should be one of:
 * 1: Disable automatic heap initialization on this thread only.
 *    If memory tagging is enabled, disable as much as possible of the
 *    memory tagging initialization for this thread.
 * 0: Normal behavior.
 *
 * Available since API level 31.
 */
#define M_THREAD_DISABLE_MEM_INIT (-103)
/**
 * mallopt() option to set the maximum number of items in the secondary
 * cache of the scudo allocator.
 *
 * Available since API level 31.
 */
#define M_CACHE_COUNT_MAX (-200)
/**
 * mallopt() option to set the maximum size in bytes of a cacheable item in
 * the secondary cache of the scudo allocator.
 *
 * Available since API level 31.
 */
#define M_CACHE_SIZE_MAX (-201)
/**
 * mallopt() option to increase the maximum number of shared thread-specific
 * data structures that can be created. This number cannot be decreased,
 * only increased and only applies to the scudo allocator.
 *
 * Available since API level 31.
 */
#define M_TSDS_COUNT_MAX (-202)

/**
 * mallopt() option to decide whether heap memory is zero-initialized on
 * allocation across the whole process. May be called at any time, including
 * when multiple threads are running. An argument of zero indicates memory
 * should not be zero-initialized, any other value indicates to initialize heap
 * memory to zero.
 *
 * Note that this memory mitigation is only implemented in scudo and therefore
 * this will have no effect when using another allocator (such as jemalloc on
 * Android Go devices).
 *
 * Available since API level 31.
 */
#define M_BIONIC_ZERO_INIT (-203)

/**
 * mallopt() option to change the heap tagging state. May be called at any
 * time, including when multiple threads are running.
 * The value must be one of the M_HEAP_TAGGING_LEVEL_ constants.
 * NOTE: This is only available in scudo.
 *
 * Available since API level 31.
 */
#define M_BIONIC_SET_HEAP_TAGGING_LEVEL (-204)

/**
 * Constants for use with the M_BIONIC_SET_HEAP_TAGGING_LEVEL mallopt() option.
 */
enum HeapTaggingLevel {
  /**
   * Disable heap tagging and memory tag checks (if supported).
   * Heap tagging may not be re-enabled after being disabled.
   */
  M_HEAP_TAGGING_LEVEL_NONE = 0,
#define M_HEAP_TAGGING_LEVEL_NONE M_HEAP_TAGGING_LEVEL_NONE
  /**
   * Address-only tagging. Heap pointers have a non-zero tag in the
   * most significant ("top") byte which is checked in free(). Memory
   * accesses ignore the tag using arm64's Top Byte Ignore (TBI) feature.
   */
  M_HEAP_TAGGING_LEVEL_TBI = 1,
#define M_HEAP_TAGGING_LEVEL_TBI M_HEAP_TAGGING_LEVEL_TBI
  /**
   * Enable heap tagging and asynchronous memory tag checks (if supported).
   * Disable stack trace collection.
   */
  M_HEAP_TAGGING_LEVEL_ASYNC = 2,
#define M_HEAP_TAGGING_LEVEL_ASYNC M_HEAP_TAGGING_LEVEL_ASYNC
  /**
   * Enable heap tagging and synchronous memory tag checks (if supported).
   * Enable stack trace collection.
   */
  M_HEAP_TAGGING_LEVEL_SYNC = 3,
#define M_HEAP_TAGGING_LEVEL_SYNC M_HEAP_TAGGING_LEVEL_SYNC
};

/**
 * mallopt() option to print human readable statistics about the memory
 * allocator to the log. There is no format for this data, each allocator
 * can use a different format, and the data that is printed can
 * change at any time. This is expected to be used as a debugging aid.
 *
 * Available since API level 35.
 */
#define M_LOG_STATS (-205)

/**
 * [mallopt(3)](https://man7.org/linux/man-pages/man3/mallopt.3.html) modifies
 * heap behavior. Values of `__option` are the `M_` constants from this header.
 *
 * Returns 1 on success, 0 on error.
 *
 * Available since API level 26.
 */

#if __BIONIC_AVAILABILITY_GUARD(26)
int mallopt(int __option, int __value) __INTRODUCED_IN(26);
#endif /* __BIONIC_AVAILABILITY_GUARD(26) */


/**
 * [__malloc_hook(3)](https://man7.org/linux/man-pages/man3/__malloc_hook.3.html)
 * is called to implement malloc(). By default this points to the system's
 * implementation.
 *
 * Available since API level 28.
 *
 * See also: [extra documentation](https://android.googlesource.com/platform/bionic/+/main/libc/malloc_hooks/README.md)
 */

#if __BIONIC_AVAILABILITY_GUARD(28)
extern void* _Nonnull (*volatile _Nonnull __malloc_hook)(size_t __byte_count, const void* _Nonnull __caller) __INTRODUCED_IN(28);

/**
 * [__realloc_hook(3)](https://man7.org/linux/man-pages/man3/__realloc_hook.3.html)
 * is called to implement realloc(). By default this points to the system's
 * implementation.
 *
 * Available since API level 28.
 *
 * See also: [extra documentation](https://android.googlesource.com/platform/bionic/+/main/libc/malloc_hooks/README.md)
 */
extern void* _Nonnull (*volatile _Nonnull __realloc_hook)(void* _Nullable __ptr, size_t __byte_count, const void* _Nonnull __caller) __INTRODUCED_IN(28);

/**
 * [__free_hook(3)](https://man7.org/linux/man-pages/man3/__free_hook.3.html)
 * is called to implement free(). By default this points to the system's
 * implementation.
 *
 * Available since API level 28.
 *
 * See also: [extra documentation](https://android.googlesource.com/platform/bionic/+/main/libc/malloc_hooks/README.md)
 */
extern void (*volatile _Nonnull __free_hook)(void* _Nullable __ptr, const void* _Nonnull __caller) __INTRODUCED_IN(28);

/**
 * [__memalign_hook(3)](https://man7.org/linux/man-pages/man3/__memalign_hook.3.html)
 * is called to implement memalign(). By default this points to the system's
 * implementation.
 *
 * Available since API level 28.
 *
 * See also: [extra documentation](https://android.googlesource.com/platform/bionic/+/main/libc/malloc_hooks/README.md)
 */
extern void* _Nonnull (*volatile _Nonnull __memalign_hook)(size_t __alignment, size_t __byte_count, const void* _Nonnull __caller) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


__END_DECLS
