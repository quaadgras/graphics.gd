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
(default: its working directory).

Type to talk to the agent. `!command` runs a shell command. `/help`
lists local commands — `/deploy` runs `GOOS=ios gd run`, which builds
the project, exports the IPA, serves it, and prints a SideStore QR
right into the scrollback (it is monospace half-block art, so it scans).

## Choosing an AI

The agent works with any of several backends — Anthropic (Claude), xAI
(Grok), Qwen/DashScope, OpenAI, or any OpenAI-compatible endpoint —
selected at runtime and remembered in `.harness/config.json`:

```
/provider              list providers and which one is active
/provider grok         switch to Grok (or anthropic / qwen / openai)
/provider custom https://host/v1 my-model   any OpenAI-compatible server
/model grok-4-fast     override the model for the current provider
/key <key>             set the API key for the current provider
```

Keys are also seeded from the environment on first run
(`ANTHROPIC_API_KEY`, `XAI_API_KEY`, `DASHSCOPE_API_KEY`,
`OPENAI_API_KEY`), and `GD_HARNESS_PROVIDER` / `GD_HARNESS_MODEL` set the
initial choice. Anthropic uses its Messages API; everyone else speaks
the OpenAI chat-completions API, so one format covers Grok, Qwen,
OpenAI, and local servers alike.

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
  tool calls → repeat. Pure `net/http`, no SDK. A neutral conversation
  is serialized by a `Provider` to either the Anthropic Messages or the
  OpenAI chat-completions wire format (`config.go`/`provider.go`), so
  Claude, Grok, Qwen, OpenAI and local servers all work. Editing tools:
  `read`, `write`, `edit` (literal, replace-all option, whitespace-
  mismatch hints), `replace` (regex/sed-like with capture backrefs),
  `grep` (project search), `ls`, `run`, `gd`.
- `internal/buildkit` — the in-process toolchain: `cmd/compile`,
  `cmd/link`, and `ld64.lld`, cross-compiled to wasm and run under
  wazero with zero process spawns. `Toolchain` exposes `Compile`,
  `LinkGo` (pure-Go internal link) and `LinkMachO` (the full game link
  against libgodot). Modules load from `GD_HARNESS_WASMTC_DIR` or a
  per-user cache; build them per `patches/` and `patches/lld-wasm/`.
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

1. **Go compilation in-process — proven.** `internal/buildkit` runs
   `cmd/compile` and `cmd/link`, cross-compiled to wasip1, as wazero
   modules inside this process: zero spawns, and a fresh module
   instance per invocation makes the (globally-stateful, single-shot)
   toolchain re-entrant and crash-isolated by construction. wazero's
   interpreter needs no JIT, so this is iOS-legal; cross-targeting is
   just `GOOS=ios GOARCH=arm64` in the guest env. The opt-in test
   (`GD_HARNESS_WASMTC_TEST=1`) builds and links a program this way
   twice on one runner and runs the result. Still to build: a mini
   cmd/go (module graph → importcfg → compile order) and a shipped
   pre-built export-data cache for std + graphics.gd targeting ios, so
   only user packages compile on-device.
2. **Pure-Go iOS binaries — proven on hardware.** Go's refusal to
   internally link ios/arm64 is policy, not physics: with the three
   toolchain patches in `internal/buildkit/patches/` (lift the
   external-linking gate, emit `LC_BUILD_VERSION` for iOS, give
   `_rt0_arm64_ios` a real body instead of stock Go's deliberate
   `UNDEF`), the wasm-hosted linker emits a signed, installable,
   *running* iPhone executable. Verified end-to-end on an iPhone 8:
   compile.wasm + link.wasm under wazero → .app → dev-sign →
   ideviceinstaller → ran and wrote its proof file on-device.
3. **Linking real games — proven.** A graphics.gd game links against
   the libgodot C++ archive, which internal linking can't do, so the
   full link needs `ld64.lld` — and lld builds to `wasm32-wasip1` and
   runs under wazero just like the Go toolchain, a fresh instance per
   link sidestepping its global-state re-entrancy problems. Verified:
   lld.wasm under wazero linked the full iOS game (MoltenVK + libgodot
   + libgo + faux-SDK frameworks, gd's exact ld64 args) to a 124MB
   arm64 Mach-O **byte-identical to the native linker except its
   LC_UUID**. Patches + recipe: `internal/buildkit/patches/lld-wasm/`.
   clang.wasm (future cgo tier) builds from the same tree.
4. **Shell tools in-process** — the built-in Go userland already covers
   the basics; the `ios_system` (BSD-3) dispatcher can add real ports,
   and git comes via go-git (pure Go).
5. **Install** — the built IPA is served on loopback and handed to
   SideStore (`sidestore://install?url=...`), which signs and installs
   on-device under the user's own Apple ID — the same handoff the
   remote kit uses today.
6. **Preview** — the graphics.gd web target in a WKWebView: wasm in
   WebKit is the one sanctioned JIT on iOS, so live preview is
   Store-legal even where running native output is not.

Measured on a desktop (wazero compiling runtime): hello-world compile
13s, link 3s. The iOS interpreter is slower; mitigations are per-package
incremental compiles, the wasm-module cache, and optionally borrowing
JIT via StikDebug. C compilation on-device is a later tier: libtcc for
plain C at near-zero size cost, embedded clang (the zig approach — or
clang-as-wasm) for everything including C++.
