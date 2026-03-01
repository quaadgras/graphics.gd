//go:build linux && cgo && !O0

// Package rodatacheck detects whether a Go string is backed by the
// binary's read-only data section (i.e. a compile-time literal).
package rodatacheck

/*
#define _GNU_SOURCE
#include <link.h>
#include <stdint.h>
#include <string.h>

static uintptr_t rodata_lo, rodata_hi;

// sentinel lives in rodata of this compilation unit; we use its address
// to identify which ELF image is "ours".
static const char sentinel[] = "rodatacheck_sentinel";

static int phdr_callback(struct dl_phdr_info *info, size_t size, void *data) {
	// check whether this image contains our sentinel
	for (int i = 0; i < info->dlpi_phnum; i++) {
		const ElfW(Phdr) *ph = &info->dlpi_phdr[i];
		if (ph->p_type != PT_LOAD)
			continue;
		uintptr_t lo = info->dlpi_addr + ph->p_vaddr;
		uintptr_t hi = lo + ph->p_memsz;
		if ((uintptr_t)sentinel >= lo && (uintptr_t)sentinel < hi)
			goto found;
	}
	return 0; // not our image
found:
	// now collect all read-only PT_LOAD segments
	for (int i = 0; i < info->dlpi_phnum; i++) {
		const ElfW(Phdr) *ph = &info->dlpi_phdr[i];
		if (ph->p_type != PT_LOAD)
			continue;
		if (ph->p_flags & PF_W)
			continue; // skip writable segments
		uintptr_t lo = info->dlpi_addr + ph->p_vaddr;
		uintptr_t hi = lo + ph->p_memsz;
		if (rodata_lo == 0 || lo < rodata_lo)
			rodata_lo = lo;
		if (hi > rodata_hi)
			rodata_hi = hi;
	}
	return 1; // stop iteration
}

static void rodatacheck_init(void) {
	dl_iterate_phdr(phdr_callback, NULL);
}

static int rodatacheck_contains(uintptr_t p) {
	return p >= rodata_lo && p < rodata_hi;
}
*/
import "C"
import "unsafe"

func init() {
	C.rodatacheck_init()
}

// String reports whether s is backed by the binary's read-only data section.
func String(s string) bool {
	p := unsafe.StringData(s)
	if p == nil {
		return false
	}
	return C.rodatacheck_contains(C.uintptr_t(uintptr(unsafe.Pointer(p)))) != 0
}
