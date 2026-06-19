// The C in this package switches the CPU TLS register (%fs / tpidr_el0) in the
// middle of its public functions. With the stack protector enabled, the canary
// is loaded from one thread-control-block and checked against another after the
// switch, tripping __stack_chk_fail. The #cgo line below disables it, matching
// the previously working version.
package dlopen

// #cgo CFLAGS: -fno-stack-protector
import "C"
