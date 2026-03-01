//go:build windows && cgo && !O0

package rodatacheck

/*
#include <windows.h>
#include <stdint.h>
#include <string.h>

static uintptr_t rodata_lo, rodata_hi;

static void rodatacheck_init(void) {
	HMODULE base = GetModuleHandle(NULL);
	if (!base)
		return;
	IMAGE_DOS_HEADER *dos = (IMAGE_DOS_HEADER *)base;
	IMAGE_NT_HEADERS *nt = (IMAGE_NT_HEADERS *)((char *)base + dos->e_lfanew);
	IMAGE_SECTION_HEADER *sec = IMAGE_FIRST_SECTION(nt);
	for (int i = 0; i < nt->FileHeader.NumberOfSections; i++, sec++) {
		if (memcmp(sec->Name, ".rdata", 6) == 0 ||
		    memcmp(sec->Name, ".rodata", 7) == 0) {
			uintptr_t lo = (uintptr_t)base + sec->VirtualAddress;
			uintptr_t hi = lo + sec->Misc.VirtualSize;
			if (rodata_lo == 0 || lo < rodata_lo)
				rodata_lo = lo;
			if (hi > rodata_hi)
				rodata_hi = hi;
		}
	}
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
