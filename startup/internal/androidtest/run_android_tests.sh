#!/usr/bin/env bash
# Run the test suite on a connected android device/emulator via `gd test`.
#
# `gd test` (GOOS=android) builds the suite as a c-shared android library,
# exports + signs a debug APK, installs it, launches it, and reads the result
# back from the app's user-data dir (startup_android.go redirects the test
# output there). It exits non-zero if the suite fails — so this wrapper is thin.
#
# UNVERIFIED: this has not run on a real emulator yet (the dev host has no KVM).
# See the caveats on Android.Test in cmd/gd/internal/builder/android.go — expect
# to iterate on the user-dir/run-as path and arg passthrough on the first runs.
#
# Env: PROJECT_DIR (default <repo>/internal — tracked, has an x86_64 Android
#      preset), GD (prebuilt gd binary), GOARCH.
set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO="${REPO:-$(git -C "$HERE" rev-parse --show-toplevel)}"
PROJECT_DIR="${PROJECT_DIR:-$REPO/internal}"
GOARCH="${GOARCH:-amd64}"
GD="${GD:-}"

if [ -z "$GD" ]; then
  GD="$(mktemp -d)/gd"
  ( cd "$REPO/cmd/gd" && go build -o "$GD" . )
fi

echo ">> connected devices:"; adb devices || true
echo ">> gd test (android/$GOARCH) in $PROJECT_DIR"
cd "$PROJECT_DIR"
exec env GOOS=android GOARCH="$GOARCH" "$GD" test -v
