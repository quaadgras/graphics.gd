#!/usr/bin/env bash
# Run the test suite on a connected android device/emulator via `gd test`.
#
# `gd test` (GOOS=android) builds the suite as a c-shared android library,
# exports + signs a release APK, installs it, launches it, and reads the result
# from logcat (startup_android.go routes the test output there). It exits
# non-zero if the suite fails — so this wrapper is thin.
#
# Verified on a physical arm64 device (full suite, 46/46 PASS). The x86_64
# emulator path is exercised by .github/workflows/android.yml.
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
