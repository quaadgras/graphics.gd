// Minimal Mach-O loader definitions (no SDK dependency).
#ifndef MACHO_LOADER_H
#define MACHO_LOADER_H

#include <stdint.h>

#define MH_MAGIC_64   0xfeedfacf
#define LC_SEGMENT_64 0x19

struct mach_header_64 {
	uint32_t magic;
	uint32_t cputype;
	uint32_t cpusubtype;
	uint32_t filetype;
	uint32_t ncmds;
	uint32_t sizeofcmds;
	uint32_t flags;
	uint32_t reserved;
};

struct load_command {
	uint32_t cmd;
	uint32_t cmdsize;
};

struct segment_command_64 {
	uint32_t cmd;
	uint32_t cmdsize;
	char     segname[16];
	uint64_t vmaddr;
	uint64_t vmsize;
	uint64_t fileoff;
	uint64_t filesize;
	uint32_t maxprot;
	uint32_t initprot;
	uint32_t nsects;
	uint32_t flags;
};

struct section_64 {
	char     sectname[16];
	char     segname[16];
	uint64_t addr;
	uint64_t size;
	uint32_t offset;
	uint32_t align;
	uint32_t reloff;
	uint32_t nreloc;
	uint32_t flags;
	uint32_t reserved1;
	uint32_t reserved2;
	uint32_t reserved3;
};

#endif
