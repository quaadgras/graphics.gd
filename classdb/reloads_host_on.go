//go:build reloads

package classdb

// registrationDisabled: this binary is a reloads host (see
// graphics.gd/startup with -tags reloads) — class registration is
// performed exclusively by the hot-reloadable wasm guest build of the
// project, so the host's own Register calls become no-ops. They are
// recorded in order, so that [ReloadsFallback] can replay them if the
// host has to give up on hot reloading and run the project natively.
var registrationDisabled = true

var deferredRegistrations []func()

func reloadsDeferRegistration(register func()) {
	deferredRegistrations = append(deferredRegistrations, register)
}

// ReloadsFallback re-enables class registration in a reloads host and
// performs every Register call the project made while it was disabled,
// in the order they were made. The startup package calls it when hot
// reloading cannot be used (the project does not build for wasm, for
// example) so that the host binary — which is the whole project — still
// registers the project's classes, GoMainLoop included.
func ReloadsFallback() {
	registrationDisabled = false
	pending := deferredRegistrations
	deferredRegistrations = nil
	for _, register := range pending {
		register()
	}
}
