# graphics.gd harness

A coding agent in a terminal rendered by the engine itself. The harness
edits, builds and ships graphics.gd projects, and it is a graphics.gd
project — so everywhere the engine runs, the harness runs.

The long game: a full development loop **on-device**. The Android leg
already works today through the exec tier (Termux + gd). The iOS leg is
the roadmap below — an IDE on the iPhone that builds iPhone apps and
hands them to SideStore to install, no computer involved.

## Run it

```sh
go install graphics.gd/cmd/gd@release
ANTHROPIC_API_KEY=... gd run
```

The harness operates on the project named by `GD_HARNESS_PROJECT`
(default: its working directory). `GD_HARNESS_MODEL` overrides the
model (default `claude-sonnet-5`).

Type to talk to the agent. `!command` runs a shell command. `/help`
lists local commands — `/deploy` runs `GOOS=ios gd run`, which builds
the project, exports the IPA, serves it, and prints a SideStore QR
right into the scrollback (it is monospace half-block art, so it scans).

## Design

Lightweight on purpose, in the spirit of pi: one file per concern, no
framework, no SDK.

- `internal/term` — the terminal widget: RichTextLabel scrollback +
  LineEdit prompt. Goroutines write into a channel; the frame loop
  drains it. Engine nodes are only touched on the main thread.
- `internal/agent` — the loop: send conversation → print text → run
  tool calls → repeat. Pure `net/http` against the Anthropic Messages
  API. Six tools: read, write, edit, ls, run, gd.
- `internal/toolchain` — how things build, behind two build tags:
  - **exec tier** (`!ios`): shells out to the `gd` command, which
    already handles every target including on-device Android (Termux)
    and iOS-over-SideStore.
  - **ios tier** (`ios`): stubs. iOS forbids process creation, so this
    tier must run everything in-process — see the roadmap.

## Roadmap: the iOS tier

iOS allows neither `exec` nor JIT, but it allows everything this
actually needs:

1. **Go compilation in-process** — compiler.gd already runs compile and
   link inside one process; embed it with a pre-populated build cache
   for std and graphics.gd so only user packages compile on-device.
2. **Linking in-process** — `ld64.lld` via `lld::safeLldMain` (LLVM
   supports being embedded; one link per build stays inside its
   envelope). The engine archive + gd.c are release-constant, so they
   pre-link on desktop with `ld -r` into one relocatable blob.
3. **Shell tools in-process** — the `ios_system` (BSD-3) dispatcher:
   commands are frameworks, `exec` becomes dlopen + call main on a
   thread. git via go-git (pure Go) or libgit2.
4. **Install** — the built IPA is served on loopback and handed to
   SideStore (`sidestore://install?url=...`), which signs and installs
   on-device under the user's own Apple ID.
5. **Preview** — the graphics.gd web target in a WKWebView: wasm in
   WebKit is the one sanctioned JIT on iOS, so live preview is
   Store-legal even where running native output is not.

C compilation on-device is a later tier: libtcc for plain C at near-zero
size cost, embedded clang (the zig approach) for everything including
C++.
