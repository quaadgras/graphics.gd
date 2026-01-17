import { useState, useEffect } from "react";
import { Github, ExternalLink, Menu, X, Sun, Moon } from "lucide-react";
import { Highlight, themes } from "prism-react-renderer";

type Platform = "Windows" | "Linux" | "macOS" | "Android" | "iOS" | "Web";

const platformCommands: Record<Platform, { cmd: string; note?: string }> = {
  Windows: { cmd: "GOOS=windows gd build" },
  Linux: { cmd: "GOOS=linux gd build" },
  macOS: { cmd: "GOOS=darwin gd build" },
  Android: {
    cmd: "GOOS=android GOARCH=arm64 gd run",
    note: "launches the project on your USB-connected Android device",
  },
  iOS: {
    cmd: "GOOS=ios gd run",
    note: "scan the QR-code on any SideStore-enabled iOS device",
  },
  Web: {
    cmd: "GOOS=js GOARCH=wasm gd run",
    note: "visit the localhost URL to see your project running in the browser",
  },
};

const helloWorldCode = `package main

import (
    "graphics.gd/startup"
    "graphics.gd/classdb/Label"
    "graphics.gd/classdb/SceneTree"
)

func main() {
    startup.LoadingScene()

    hello := Label.New()
    hello.SetText("Hello, World!")
    SceneTree.Add(hello)

    startup.Scene()
}`;

const shaderCode = `type MyShader struct {
    CanvasItem.Shader[MyShader]

    Color rgba.Color \`gd:"color"\` // uniform
}

func (my MyShader) Material(frag CanvasItem.Fragment) CanvasItem.Material {
    return CanvasItem.Material{
        Color: my.Color,
    }
}`;

const extensionCode = `type Player struct {
  CharacterBody2D.Extension[Player]

  Speed    float64 \`gd:"speed" range:"0,500,10"\`
  JumpForce float64 \`gd:"jump_force"\`
}

func (p *Player) TakeDamage(amount int) {
    fmt.Println("Player took", amount, "damage")
}`;

const gdscriptCode = `@onready var player: Player = $Player

func _ready():
    player.speed = 200.0
    player.jump_force = 400.0

func _on_enemy_hit():
    player.take_damage(10)`;

function CodeBlock({
  code,
  filename,
  isDark,
}: {
  code: string;
  filename: string;
  isDark: boolean;
}) {
  return (
    <div className="bg-white dark:bg-zinc-900 rounded-lg border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm">
      <div className="flex items-center gap-2 px-4 py-2 bg-zinc-100 dark:bg-zinc-900 border-b border-zinc-200 dark:border-zinc-800">
        <span className="text-xs text-zinc-500 font-mono">{filename}</span>
      </div>
      <Highlight
        theme={isDark ? themes.nightOwl : themes.github}
        code={code}
        language="go"
      >
        {({ style, tokens, getLineProps, getTokenProps }) => (
          <pre
            className="p-4 text-sm overflow-x-auto"
            style={{ ...style, background: "transparent" }}
          >
            {tokens.map((line, i) => (
              <div key={i} {...getLineProps({ line })}>
                {line.map((token, key) => (
                  <span key={key} {...getTokenProps({ token })} />
                ))}
              </div>
            ))}
          </pre>
        )}
      </Highlight>
    </div>
  );
}

function App() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [selectedPlatform, setSelectedPlatform] = useState<Platform>("Windows");
  const [isDark, setIsDark] = useState(() => {
    if (typeof window !== "undefined") {
      const stored = localStorage.getItem("theme");
      if (stored) return stored === "dark";
      return window.matchMedia("(prefers-color-scheme: dark)").matches;
    }
    return true;
  });

  useEffect(() => {
    document.documentElement.classList.toggle("dark", isDark);
    localStorage.setItem("theme", isDark ? "dark" : "light");
  }, [isDark]);

  return (
    <div className="min-h-screen bg-zinc-50 dark:bg-zinc-950 text-zinc-900 dark:text-zinc-100 transition-colors">
      {/* Navigation */}
      <nav className="fixed top-0 w-full bg-zinc-50/95 dark:bg-zinc-950/95 backdrop-blur border-b border-zinc-200 dark:border-zinc-800 z-50 transition-colors">
        <div className="max-w-5xl mx-auto px-6">
          <div className="flex justify-between items-center h-14">
            <a
              href="#"
              className="font-mono text-lg font-bold text-zinc-900 dark:text-white"
            >
              graphics.gd
            </a>

            <div className="hidden md:flex items-center gap-8">
              <a
                href="https://the.graphics.gd/guide"
                target="_blank"
                rel="noopener noreferrer"
                className="text-sm text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-white transition-colors"
              >
                Documentation
              </a>
              <a
                href="https://github.com/quaadgras/graphics.gd/tree/samples"
                target="_blank"
                rel="noopener noreferrer"
                className="text-sm text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-white transition-colors"
              >
                Examples
              </a>
              <a
                href="https://github.com/quaadgras/graphics.gd"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center gap-2 text-sm text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-white transition-colors"
              >
                <Github className="h-4 w-4" />
                GitHub
              </a>
              <button
                onClick={() => setIsDark(!isDark)}
                className="p-2 text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-white transition-colors"
                aria-label="Toggle theme"
              >
                {isDark ? (
                  <Sun className="h-4 w-4" />
                ) : (
                  <Moon className="h-4 w-4" />
                )}
              </button>
            </div>

            <div className="flex items-center gap-2 md:hidden">
              <button
                onClick={() => setIsDark(!isDark)}
                className="p-2 text-zinc-600 dark:text-zinc-400"
                aria-label="Toggle theme"
              >
                {isDark ? (
                  <Sun className="h-5 w-5" />
                ) : (
                  <Moon className="h-5 w-5" />
                )}
              </button>
              <button
                onClick={() => setIsMenuOpen(!isMenuOpen)}
                className="p-2 text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-white"
              >
                {isMenuOpen ? (
                  <X className="h-5 w-5" />
                ) : (
                  <Menu className="h-5 w-5" />
                )}
              </button>
            </div>
          </div>
        </div>

        {isMenuOpen && (
          <div className="md:hidden border-t border-zinc-200 dark:border-zinc-800 bg-zinc-50 dark:bg-zinc-950">
            <div className="px-6 py-4 space-y-4">
              <a
                href="https://pkg.go.dev/graphics.gd"
                className="block text-sm text-zinc-600 dark:text-zinc-400"
              >
                pkg.go.dev
              </a>
              <a
                href="https://the.graphics.gd/guide"
                className="block text-sm text-zinc-600 dark:text-zinc-400"
              >
                Guide
              </a>
              <a
                href="https://github.com/quaadgras/graphics.gd/tree/samples"
                className="block text-sm text-zinc-600 dark:text-zinc-400"
              >
                Samples
              </a>
              <a
                href="https://github.com/quaadgras/graphics.gd"
                className="block text-sm text-zinc-600 dark:text-zinc-400"
              >
                GitHub
              </a>
            </div>
          </div>
        )}
      </nav>

      {/* Hero */}
      <section className="pt-32 pb-16 px-6">
        <div className="max-w-5xl mx-auto">
          <h1 className="font-mono text-4xl sm:text-5xl lg:text-6xl font-bold mb-6">
            <span className="text-cyan-600 dark:text-cyan-400">Go</span> +{" "}
            <span className="text-[#478CBF]">Godot</span>
          </h1>
          <p className="text-xl sm:text-2xl text-zinc-600 dark:text-zinc-400 max-w-2xl mb-8 leading-relaxed">
            The simplicity of Go + the full graphics and game development
            capabilities of Godot
          </p>

          <div className="flex flex-wrap gap-4 mb-12">
            <a
              href="https://the.graphics.gd/guide"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-2 bg-cyan-600 hover:bg-cyan-500 text-white px-5 py-2.5 rounded font-medium transition-colors"
            >
              Read the Guide
              <ExternalLink className="h-4 w-4" />
            </a>
            <a
              href="https://github.com/quaadgras/graphics.gd"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-2 border border-zinc-300 dark:border-zinc-700 hover:border-zinc-400 dark:hover:border-zinc-500 text-zinc-700 dark:text-zinc-300 px-5 py-2.5 rounded font-medium transition-colors"
            >
              <Github className="h-4 w-4" />
              View Source
            </a>
          </div>

          {/* Quick start */}
          <div className="bg-white dark:bg-zinc-900 rounded-lg border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm">
            <div className="flex items-center gap-2 px-4 py-2 bg-zinc-100 dark:bg-zinc-900 border-b border-zinc-200 dark:border-zinc-800">
              <div className="w-3 h-3 rounded-full bg-zinc-300 dark:bg-zinc-700"></div>
              <div className="w-3 h-3 rounded-full bg-zinc-300 dark:bg-zinc-700"></div>
              <div className="w-3 h-3 rounded-full bg-zinc-300 dark:bg-zinc-700"></div>
              <span className="ml-2 text-xs text-zinc-500 font-mono">
                terminal
              </span>
            </div>
            <div className="p-4 font-mono text-sm overflow-x-auto">
              <div className="text-green-600 dark:text-green-400">
                $ go install graphics.gd/cmd/gd@release
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Hello World */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold mb-4">
            Start with a single file
          </h2>
          <p className="text-zinc-600 dark:text-zinc-400 mb-8 max-w-2xl">
            No scaffolding needed. Create a{" "}
            <code className="bg-zinc-200 dark:bg-zinc-800 px-1.5 py-0.5 rounded text-sm">
              main.go
            </code>{" "}
            and use{" "}
            <code className="bg-zinc-200 dark:bg-zinc-800 px-1.5 py-0.5 rounded text-sm">
              gd run
            </code>
            to launch your project.
          </p>

          <CodeBlock code={helloWorldCode} filename="main.go" isDark={isDark} />
        </div>
      </section>

      {/* SDK-free cross-platform */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <div className="mb-8">
            <h2 className="text-2xl sm:text-3xl font-bold mb-4">
              Run Everywhere (no proprietary SDKs required)
            </h2>
            <p className="text-zinc-600 dark:text-zinc-400 max-w-2xl">
              Build native binaries for each platform, on any machine. Yes even
              for Android, iOS and macOS!
            </p>
          </div>

          {/* Interactive platform selector */}
          <div className="bg-white dark:bg-zinc-900 rounded-lg border border-zinc-200 dark:border-zinc-800 overflow-hidden shadow-sm">
            <div className="flex flex-wrap gap-1 p-2 bg-zinc-100 dark:bg-zinc-900 border-b border-zinc-200 dark:border-zinc-800">
              {(Object.keys(platformCommands) as Platform[]).map((platform) => (
                <button
                  key={platform}
                  onClick={() => setSelectedPlatform(platform)}
                  className={`px-3 py-1.5 rounded text-sm font-medium transition-colors ${
                    selectedPlatform === platform
                      ? "bg-cyan-600 text-white"
                      : "text-zinc-600 dark:text-zinc-400 hover:bg-zinc-200 dark:hover:bg-zinc-800"
                  }`}
                >
                  {platform}
                </button>
              ))}
            </div>
            <div className="p-4 font-mono text-sm">
              <div className="flex items-center gap-2">
                <span className="text-zinc-500">$</span>
                <span className="text-green-600 dark:text-green-400">
                  {platformCommands[selectedPlatform].cmd}
                </span>
              </div>
              {platformCommands[selectedPlatform].note && (
                <div className="mt-3 text-xs text-cyan-600 dark:text-cyan-400 bg-cyan-50 dark:bg-cyan-950/30 border border-cyan-200 dark:border-cyan-900 rounded px-3 py-2">
                  {platformCommands[selectedPlatform].note}
                </div>
              )}
            </div>
          </div>

          <p className="text-zinc-600 dark:text-zinc-400 mt-8 max-w-2xl">
            <i>
              You can also use the gd command line tool with existing GDScript
              projects to easily build your projects for different platforms.
            </i>
          </p>
        </div>
      </section>

      {/* Shaders in Go */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold mb-4">
            You can even write shaders in Go
          </h2>
          <p className="text-zinc-600 dark:text-zinc-400 mb-8 max-w-2xl">
            No GLSL required. Define shaders using Go's type system with full
            IDE support, type checking, and composition.
          </p>

          <CodeBlock code={shaderCode} filename="shader.go" isDark={isDark} />
        </div>
      </section>

      {/* GDScript Integration */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold mb-4">
            With Automatic Scripting Integration
          </h2>
          <p className="text-zinc-600 dark:text-zinc-400 mb-8 max-w-2xl">
            Define types in Go, use them in GDScript. Exported fields become
            properties, exported methods become callable — no boilerplate
            required.
          </p>

          <div className="grid lg:grid-cols-2 gap-4">
            <CodeBlock
              code={extensionCode}
              filename="player.go"
              isDark={isDark}
            />
            <CodeBlock code={gdscriptCode} filename="game.gd" isDark={isDark} />
          </div>
          <p className="text-zinc-600 dark:text-zinc-400 mt-8 max-w-2xl">
            <i>
              Perfect when you're still learning Go or you need a little native
              functionality in an existing GDScript project.
            </i>
          </p>
        </div>
      </section>

      {/* Why this exists */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold mb-8">
            Designed for Go, not just a wrapper
          </h2>

          <div className="grid sm:grid-cols-2 gap-6">
            <div className="space-y-4">
              <div>
                <h3 className="font-semibold text-zinc-900 dark:text-white mb-1">
                  Stronger Typing
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm">
                  RIDs, Callables, and Dictionaries are all distinctly typed.
                  Errors that other bindings catch at runtime, graphics.gd
                  catches at compile time.
                </p>
              </div>
              <div>
                <h3 className="font-semibold text-zinc-900 dark:text-white mb-1">
                  Documentation Included
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm">
                  The complete Godot API documentation on{" "}
                  <a
                    href="https://pkg.go.dev/graphics.gd"
                    className="text-cyan-600 dark:text-cyan-400 hover:underline"
                  >
                    pkg.go.dev
                  </a>{" "}
                  with all code-snippets, links and formatting ported from Godot
                  to Go.
                </p>
              </div>
              <div>
                <h3 className="font-semibold text-zinc-900 dark:text-white mb-1">
                  Fully Ported Math
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm">
                  Vector math, colors, and transforms have all been ported to
                  pure Go. That means no FFI overhead plus you can use them all
                  in any Go project.
                </p>
              </div>
            </div>

            <div className="space-y-4">
              <div>
                <h3 className="font-semibold text-zinc-900 dark:text-white mb-1">
                  Fast Compilation
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm">
                  After the initial cgo build, recompilation is fast. All of the
                  safety, speed and convieniance with none of the wait.
                </p>
              </div>
              <div>
                <h3 className="font-semibold text-zinc-900 dark:text-white mb-1">
                  Fully Cross Platform
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm">
                  Build and run on Android without Java nor the Google SDK.
                  Build for iOS/macOS without Xcode. Target any platform from
                  any host. All thanks to Go &{" "}
                  <a
                    href="https://ziglang.org/"
                    className="text-cyan-600 dark:text-cyan-400 hover:underline"
                  >
                    Zig
                  </a>
                  .
                </p>
              </div>
              <div>
                <h3 className="font-semibold text-zinc-900 dark:text-white mb-1">
                  Performance Tuning
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm">
                  Access specialised & advanced techniques for reducing
                  allocations, enabling garbage collection to be bypassed for
                  engine objects.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Project showcase */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold mb-8">
            Built with graphics.gd
          </h2>

          <a
            href="https://the.quetzal.community/aviary"
            target="_blank"
            rel="noopener noreferrer"
            className="block group"
          >
            <div className="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-lg overflow-hidden hover:border-zinc-300 dark:hover:border-zinc-700 transition-colors shadow-sm">
              <img
                src="https://github.com/user-attachments/assets/336e56f6-445b-42c9-bc9a-808f1931700c"
                alt="Aviary screenshot"
                className="w-full h-48 sm:h-64 object-cover"
              />
              <div className="p-4">
                <h3 className="font-semibold text-zinc-900 dark:text-white group-hover:text-cyan-600 dark:group-hover:text-cyan-400 transition-colors">
                  Aviary
                </h3>
                <p className="text-zinc-600 dark:text-zinc-400 text-sm mt-1">
                  A cooperative space and scene editor inspired by video games
                  from the RTS, Tycoon, and Simulation genres.
                </p>
              </div>
            </div>
          </a>
        </div>
      </section>

      {/* Get involved */}
      <section className="py-16 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-2xl sm:text-3xl font-bold mb-4">Get involved</h2>
          <p className="text-zinc-600 dark:text-zinc-400 mb-6 max-w-2xl">
            graphics.gd is shaped by real-world use. Try building something,{" "}
            <a
              href="https://github.com/quaadgras/graphics.gd/issues/new/choose"
              className="text-cyan-600 dark:text-cyan-400 hover:underline"
            >
              report rough edges
            </a>
            , improve the variant packages, write benchmarks, or just share what
            you're making. Every bit helps.
          </p>

          <div className="flex flex-wrap gap-4">
            <a
              href="https://github.com/quaadgras/graphics.gd/discussions"
              target="_blank"
              rel="noopener noreferrer"
              className="text-sm text-cyan-600 dark:text-cyan-400 hover:underline"
            >
              Discussions →
            </a>
            <a
              href="https://github.com/sponsors/Splizard"
              target="_blank"
              rel="noopener noreferrer"
              className="text-sm text-cyan-600 dark:text-cyan-400 hover:underline"
            >
              Sponsor →
            </a>
            <a
              href="https://x.com/splizard"
              target="_blank"
              rel="noopener noreferrer"
              className="text-sm text-cyan-600 dark:text-cyan-400 hover:underline"
            >
              @splizard →
            </a>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-8 px-6 border-t border-zinc-200 dark:border-zinc-800">
        <div className="max-w-5xl mx-auto flex flex-col sm:flex-row justify-between items-center gap-4">
          <p className="text-zinc-500 text-sm">MIT License</p>
          <div className="flex gap-6">
            <a
              href="https://the.graphics.gd/guide"
              className="text-zinc-500 text-sm hover:text-zinc-700 dark:hover:text-zinc-300"
            >
              Documentation
            </a>
            <a
              href="https://github.com/quaadgras/graphics.gd"
              className="text-zinc-500 text-sm hover:text-zinc-700 dark:hover:text-zinc-300"
            >
              GitHub
            </a>
          </div>
        </div>
      </footer>
    </div>
  );
}

export default App;
