//go:build cgo

package ring

// #include <stdint.h>
// extern void gd_ring_flush(void *entries, uint32_t tail, uint32_t head);
import "C"
import "unsafe"

func flush(entries unsafe.Pointer, tail, head uint32) {
	C.gd_ring_flush(entries, C.uint32_t(tail), C.uint32_t(head))
}
