//go:build cgo

package ring

// #include <stdint.h>
// extern void gd_ring_flush(void *entries, uint32_t tail, uint32_t head, uint32_t *crash_index);
import "C"
import "unsafe"

func flush(entries unsafe.Pointer, tail, head uint32) {
	C.gd_ring_flush(entries, C.uint32_t(tail), C.uint32_t(head), (*C.uint32_t)(unsafe.Pointer(&CrashIndex)))
}
