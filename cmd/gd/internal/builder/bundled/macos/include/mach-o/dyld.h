// Minimal dyld definitions (no SDK dependency).
#ifndef MACHO_DYLD_H
#define MACHO_DYLD_H

#include <stdint.h>
#include "loader.h"

extern uint32_t                    _dyld_image_count(void);
extern const struct mach_header_64 *_dyld_get_image_header(uint32_t);
extern intptr_t                    _dyld_get_image_vmaddr_slide(uint32_t);

#endif
