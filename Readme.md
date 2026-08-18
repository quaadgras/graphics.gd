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

## The iPhone loop, today

The harness runs on iOS now (its own IPA installs via the usual
`GOOS=ios gd run` from any machine with gd). On-device it has agent
chat, file tools and a small built-in shell — and with a remote
configured it has builds too:

```
/key sk-ant-...
/remote user@fold:8022 /path/to/your/project
   (add the printed key to the remote's ~/.ssh/authorized_keys)
/deploy
```

`/remote` turns the harness into a thin client on that machine's
checkout: the agent's read/write/edit/ls/run all travel over SSH, and
`/deploy` runs `GOOS=ios gd run` there. When the served
`sidestore://install?url=...` link appears in the build output, the
harness opens it on-device — SideStore installs the freshly built app
onto the same iPhone you edited it from. Phone in your pocket builds,
phone in your hand ships. The in-process toolchain below will replace
the SSH hop without changing the interface.

## Design

Lightweight on purpose, in the spirit of pi: one file per concern, no
framework, no SDK.

- `internal/term` — the terminal widget: RichTextLabel scrollback +
  LineEdit prompt. Goroutines write into a channel; the frame loop
  drains it. Engine nodes are only touched on the main thread.
- `internal/agent` — the loop: send conversation → print text → run
  tool calls → repeat. Pure `net/http` against the Anthropic Messages
  API. Six tools: read, write, edit, ls, run, gd.
- `internal/toolchain` — a `Kit` interface (read, write, list, shell,
  gd) with three implementations:
  - **exec tier** (`!ios`): shells out to the `gd` command, which
    already handles every target including on-device Android (Termux)
    and iOS-over-SideStore.
  - **ios tier** (`ios`): a built-in pure-Go userland for the shell
    (iOS forbids process creation); building needs a remote or the
    roadmap's in-process tier.
  - **remote** (all platforms): the same operations over SSH against
    another machine's checkout — pure-Go `x/crypto/ssh`, generated
    ed25519 key, host key pinned on first connection.

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
