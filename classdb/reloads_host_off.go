//go:build !reloads

package classdb

const registrationDisabled = false

func reloadsDeferRegistration(func()) {}

// ReloadsFallback is a no-op outside of a reloads host (-tags reloads):
// class registration is never disabled, so there is nothing to replay.
func ReloadsFallback() {}
