/*-
 * Written by J.T. Conklin <jtc@netbsd.org>
 * Public domain.
 *
 *	$NetBSD: search.h,v 1.12 1999/02/22 10:34:28 christos Exp $
 * $FreeBSD: release/9.0.0/include/search.h 105250 2002-10-16 14:29:23Z robert $
 */

#pragma once

/**
 * @file search.h
 * @brief Queues, hash tables, trees, and linear array searches.
 */

#include <sys/cdefs.h>
#include <sys/types.h>

/** See hsearch()/hsearch_r(). */
typedef enum {
  FIND,
  ENTER
} ACTION;

/** See hsearch()/hsearch_r(). */
typedef struct entry {
  /** The string key. */
  char* _Nullable key;
  /** The associated data. */
  void* _Nullable data;
} ENTRY;

/**
 * Constants given to the twalk() visitor.
 * Note that the constant names are misleading.
 */
typedef enum {
  /**
   * If this is the first visit to a non-leaf node.
   * Use this for *preorder* traversal.
   */
  preorder,
  /**
   * If this is the second visit to a non-leaf node.
   * Use this for *inorder* traversal.
   */
  postorder,
  /**
   * If this is the third visit to a non-leaf node.
   * Use this for *postorder* traversal.
   */
  endorder,
  /** If this is the first and only visit to a leaf node. */
  leaf
} VISIT;

#if defined(__USE_BSD) || defined(__USE_GNU)
/** The hash table type for hcreate_r()/hdestroy_r()/hsearch_r(). */
struct hsearch_data {
  struct __hsearch* _Nullable __hsearch;
};
#endif

__BEGIN_DECLS

/**
 * [insque(3)](https://man7.org/linux/man-pages/man3/insque.3.html) inserts
 * an item in a queue (an intrusive doubly-linked list).
 */
void insque(void* _Nonnull __element, void* _Nullable __previous);

/**
 * [remque(3)](https://man7.org/linux/man-pages/man3/remque.3.html) removes
 * an item from a queue (an intrusive doubly-linked list).
 */
void remque(void* _Nonnull __element);

/**
 * [hcreate(3)](https://man7.org/linux/man-pages/man3/hcreate.3.html)
 * initializes the global hash table, with space for at least `__n` elements.
 *
 * See hcreate_r() if you need more than one hash table.
 *
 * Returns *non-zero* on success and returns 0 and sets `errno` on failure.
 *
 * Available since API level 28.
 */

#if __BIONIC_AVAILABILITY_GUARD(28)
int hcreate(size_t __n) __INTRODUCED_IN(28);

/**
 * [hdestroy(3)](https://man7.org/linux/man-pages/man3/hdestroy.3.html) destroys
 * the global hash table.
 *
 * See hdestroy_r() if you need more than one hash table.
 *
 * Available since API level 28.
 */
void hdestroy(void) __INTRODUCED_IN(28);

/**
 * [hsearch(3)](https://man7.org/linux/man-pages/man3/hsearch.3.html) finds or
 * inserts `__entry` in the global hash table, based on `__action`.
 *
 * See hsearch_r() if you need more than one hash table.
 *
 * Returns a pointer to the entry on success, and returns NULL and sets
 * `errno` on failure.
 *
 * Available since API level 28.
 */
ENTRY* _Nullable hsearch(ENTRY __entry, ACTION __action) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


#if defined(__USE_BSD) || defined(__USE_GNU)

/**
 * [hcreate_r(3)](https://man7.org/linux/man-pages/man3/hcreate_r.3.html)
 * initializes a hash table `__table` with space for at least `__n` elements.
 *
 * Returns *non-zero* on success and returns 0 and sets `errno` on failure.
 *
 * Available since API level 28.
 */

#if __BIONIC_AVAILABILITY_GUARD(28)
int hcreate_r(size_t __n, struct hsearch_data* _Nonnull __table) __INTRODUCED_IN(28);

/**
 * [hdestroy_r(3)](https://man7.org/linux/man-pages/man3/hdestroy_r.3.html) destroys
 * the hash table `__table`.
 *
 * Available since API level 28.
 */
void hdestroy_r(struct hsearch_data* _Nonnull __table) __INTRODUCED_IN(28);

/**
 * [hsearch_r(3)](https://man7.org/linux/man-pages/man3/hsearch_r.3.html) finds or
 * inserts `__entry` in the hash table `__table`, based on `__action`.
 *
 * Returns *non-zero* on success and returns 0 and sets `errno` on failure.
 * A pointer to the entry is returned in `*__result`.
 *
 * Available since API level 28.
 */
int hsearch_r(ENTRY __entry, ACTION __action, ENTRY* _Nullable * _Nonnull __result, struct hsearch_data* _Nonnull __table) __INTRODUCED_IN(28);
#endif /* __BIONIC_AVAILABILITY_GUARD(28) */


#endif

/**
 * [lfind(3)](https://man7.org/linux/man-pages/man3/lfind.3.html) brute-force
 * searches the unsorted array `__array` (of `__count` items each of size `__size`)
 * for `__key`, using `__comparator`.
 *
 * See bsearch() if you have a sorted array.
 *
 * Returns a pointer to the matching element on success, or NULL on failure.
 */
void* _Nullable lfind(const void* _Nonnull __key, const void* _Nonnull __array, size_t* _Nonnull __count, size_t __size, int (* _Nonnull __comparator)(const void* _Nonnull, const void* _Nonnull));

/**
 * [lsearch(3)](https://man7.org/linux/man-pages/man3/lsearch.3.html) brute-force
 * searches the unsorted array `__array` (of `__count` items each of size `__size`)
 * for `__key`, using `__comparator`.
 *
 * Unlike lfind(), on failure lsearch() will *insert* `__key` at the end of
 * `__array` and increment `*__count`.
 *
 * Returns a pointer to the matching element on success, or to the newly-added
 * element on failure.
 */
void* _Nonnull lsearch(const void* _Nonnull __key, void* _Nonnull __array, size_t* _Nonnull __count, size_t __size, int (* _Nonnull __comparator)(const void* _Nonnull, const void* _Nonnull));

/**
 * [tdelete(3)](https://man7.org/linux/man-pages/man3/tdelete.3.html) searches
 * for and removes an element in the tree `*__root_ptr`. The search is performed
 * using `__comparator`.
 *
 * Returns a pointer to the parent of the deleted node, or NULL on failure.
 */
void* _Nullable tdelete(const void* _Nonnull __key, void* _Nullable * _Nullable __root_ptr, int (* _Nonnull __comparator)(const void* _Nonnull, const void* _Nonnull));

/**
 * [tdestroy(3)](https://man7.org/linux/man-pages/man3/tdestroy.3.html) destroys
 * the hash table `__root` using `__free_fn` on each node.
 */
void tdestroy(void* _Nullable __root, void (* _Nullable __free_fn)(void* _Nullable));

/**
 * [tfind(3)](https://man7.org/linux/man-pages/man3/tfind.3.html) searches
 * for an element in the tree `*__root_ptr`. The search is performed using
 * `__comparator`.
 *
 * Returns a pointer to the matching node, or NULL on failure.
 */
void* _Nullable tfind(const void* _Nonnull __key, void* _Nullable const* _Nullable __root_ptr, int (* _Nonnull __comparator)(const void* _Nonnull, const void* _Nonnull));

/**
 * [tsearch(3)](https://man7.org/linux/man-pages/man3/tsearch.3.html) searches
 * for an element in the tree `*__root_ptr`. The search is performed using
 * `__comparator`.
 *
 * Unlike tfind(), on failure tsearch() will *insert* `__key` into the tree.
 *
 * Returns a pointer to the matching node, or to the newly-added node.
 */
void* _Nullable tsearch(const void* _Nonnull __key, void* _Nullable * _Nullable __root_ptr, int (* _Nonnull __comparator)(const void* _Nonnull, const void* _Nonnull));

/**
 * [twalk(3)](https://man7.org/linux/man-pages/man3/twalk.3.html) calls
 * `__visitor` on every node in the tree.
 */
void twalk(const void* _Nullable __root, void (* _Nullable __visitor)(const void* _Nullable, VISIT, int));

__END_DECLS
