#!/usr/bin/env bash
# Build the js/wasm test binary, export the Godot web build, serve it, and run
# it in headless Chromium — exiting non-zero if the embedded `go test` fails.
#
# This is the single entrypoint used by CI and for local container verification.
#
# Env:
#   PROJECT_DIR    project dir to run `gd test` in        (default: <repo>/internal)
#   PORT           port gd's dev server listens on        (default: 8080)
#   TIMEOUT_MS     headless run timeout, ms               (default: 180000)
#   SERVE_TIMEOUT  seconds to wait for build+export+serve (default: 900)
#   GD             path to a prebuilt `gd` binary         (default: build from source)
#   CHROMIUM       chromium/chrome binary                 (default: chromium)
#
# Extra args are forwarded to `gd test` (e.g. -run TestFoo -v).
set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO="${REPO:-$(git -C "$HERE" rev-parse --show-toplevel)}"
PROJECT_DIR="${PROJECT_DIR:-$REPO/internal}"
PORT="${PORT:-8080}"
TIMEOUT_MS="${TIMEOUT_MS:-180000}"
SERVE_TIMEOUT="${SERVE_TIMEOUT:-900}"
GD="${GD:-}"

if [ -z "$GD" ]; then
  GD="$(mktemp -d)/gd"
  echo ">> building gd command from $REPO/cmd/gd"
  ( cd "$REPO/cmd/gd" && go build -o "$GD" . )
fi

LOG="$(mktemp)"
GD_PGID=""
cleanup() { [ -n "$GD_PGID" ] && kill -TERM -"$GD_PGID" 2>/dev/null || true; }
trap cleanup EXIT INT TERM

echo ">> exporting + serving web test build on :$PORT (project: $PROJECT_DIR)"
set -m   # run the background job in its own process group so we can kill the tree
( cd "$PROJECT_DIR" && exec env PORT="$PORT" GOOS=js "$GD" test "$@" ) >"$LOG" 2>&1 &
GD_PGID=$!
set +m

ok=""
for _ in $(seq 1 "$SERVE_TIMEOUT"); do
  if grep -q "serving wasm/js on http://localhost:$PORT" "$LOG"; then ok=1; break; fi
  if ! kill -0 "$GD_PGID" 2>/dev/null; then break; fi
  sleep 1
done

echo "----- gd output -----"; cat "$LOG"; echo "---------------------"
if [ -z "$ok" ]; then
  echo "!! gd dev server never came up"
  exit 1
fi

echo ">> server up; running headless browser test"
node "$HERE/run_browser_test.mjs" "http://localhost:$PORT" "$TIMEOUT_MS"
