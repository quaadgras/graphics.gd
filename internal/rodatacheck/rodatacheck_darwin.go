//go:build darwin && cgo && !O0

package rodatacheck

/*
#include <mach-o/loader.h>
#include <mach-o/dyld.h>
#include <string.h>

static uintptr_t rodata_lo, rodata_hi;

static const char sentinel[] = "rodatacheck_sentinel";

static void rodatacheck_init(void) {
	uint32_t count = _dyld_image_count();
	for (uint32_t img = 0; img < count; img++) {
		const struct mach_header_64 *hdr = _dyld_get_image_header(img);
		if (!hdr || hdr->magic != MH_MAGIC_64)
			continue;
		intptr_t slide = _dyld_get_image_vmaddr_slide(img);

		// first pass: check if this image contains our sentinel
		int found = 0;
		const uint8_t *ptr = (const uint8_t *)(hdr + 1);
		for (uint32_t i = 0; i < hdr->ncmds; i++) {
			const struct load_command *lc = (const struct load_command *)ptr;
			if (lc->cmd == LC_SEGMENT_64) {
				const struct segment_command_64 *seg = (const struct segment_command_64 *)ptr;
				uintptr_t lo = (uintptr_t)seg->vmaddr + slide;
				uintptr_t hi = lo + seg->vmsize;
				if ((uintptr_t)sentinel >= lo && (uintptr_t)sentinel < hi) {
					found = 1;
					break;
				}
			}
			ptr += lc->cmdsize;
		}
		if (!found)
			continue;

		// second pass: collect __rodata and __cstring sections in __TEXT
		ptr = (const uint8_t *)(hdr + 1);
		for (uint32_t i = 0; i < hdr->ncmds; i++) {
			const struct load_command *lc = (const struct load_command *)ptr;
			if (lc->cmd == LC_SEGMENT_64) {
				const struct segment_command_64 *seg = (const struct segment_command_64 *)ptr;
				if (strncmp(seg->segname, "__TEXT", 6) == 0) {
					const struct section_64 *sect = (const struct section_64 *)(seg + 1);
					for (uint32_t j = 0; j < seg->nsects; j++) {
						if (strncmp(sect[j].sectname, "__rodata", 8) == 0 ||
						    strncmp(sect[j].sectname, "__cstring", 9) == 0) {
							uintptr_t lo = (uintptr_t)sect[j].addr + slide;
							uintptr_t hi = lo + sect[j].size;
							if (rodata_lo == 0 || lo < rodata_lo)
								rodata_lo = lo;
							if (hi > rodata_hi)
								rodata_hi = hi;
						}
					}
				}
			}
			ptr += lc->cmdsize;
		}
		return;
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
