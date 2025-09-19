import React, { useState, useEffect } from "react";
import {
  Github,
  ExternalLink,
  Play,
  Download,
  Code2,
  Zap,
  Shield,
  Users,
  Heart,
  Menu,
  X,
  ChevronRight,
  Star,
  Coffee,
  Gamepad2,
  Twitter,
} from "lucide-react";

function App() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [activeSection, setActiveSection] = useState("");

  useEffect(() => {
    const handleScroll = () => {
      const sections = [
        "hero",
        "features",
        "code",
        "getting-started",
        "performance",
        "platforms",
        "contributing",
      ];
      const scrollPosition = window.scrollY + 100;

      for (const section of sections) {
        const element = document.getElementById(section);
        if (element) {
          const { offsetTop, offsetHeight } = element;
          if (
            scrollPosition >= offsetTop &&
            scrollPosition < offsetTop + offsetHeight
          ) {
            setActiveSection(section);
            break;
          }
        }
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  const scrollToSection = (sectionId: string) => {
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: "smooth" });
    }
    setIsMenuOpen(false);
  };

  const navItems = [
    { id: "hero", label: "Home" },
    { id: "features", label: "Features" },
    { id: "code", label: "Code" },
    { id: "getting-started", label: "Get Started" },
    { id: "performance", label: "Performance" },
    { id: "platforms", label: "Platforms" },
    { id: "contributing", label: "Contributing" },
  ];

  return (
    <div className="min-h-screen bg-white">
      {/* Navigation */}
      <nav className="fixed top-0 w-full bg-white/90 backdrop-blur-sm border-b border-gray-200 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-2">
              <Gamepad2 className="h-8 w-8 text-blue-600" />
              <span className="text-xl font-bold text-gray-900">
                graphics.gd
              </span>
            </div>

            {/* Desktop Navigation */}
            <div className="hidden md:flex items-center space-x-8">
              {navItems.map((item) => (
                <button
                  key={item.id}
                  onClick={() => scrollToSection(item.id)}
                  className={`text-sm font-medium transition-colors hover:text-blue-600 ${
                    activeSection === item.id
                      ? "text-blue-600"
                      : "text-gray-700"
                  }`}
                >
                  {item.label}
                </button>
              ))}
              <a
                href="https://github.com/quaadgras/graphics.gd"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center space-x-2 bg-gray-900 text-white px-4 py-2 rounded-lg hover:bg-gray-800 transition-colors"
              >
                <Github className="h-4 w-4" />
                <span>GitHub</span>
              </a>
            </div>

            {/* Mobile menu button */}
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="md:hidden p-2 rounded-lg hover:bg-gray-100"
            >
              {isMenuOpen ? (
                <X className="h-6 w-6" />
              ) : (
                <Menu className="h-6 w-6" />
              )}
            </button>
          </div>
        </div>

        {/* Mobile Navigation */}
        {isMenuOpen && (
          <div className="md:hidden bg-white border-t border-gray-200">
            <div className="px-4 py-2 space-y-2">
              {navItems.map((item) => (
                <button
                  key={item.id}
                  onClick={() => scrollToSection(item.id)}
                  className={`block w-full text-left px-3 py-2 rounded-lg text-sm font-medium transition-colors ${
                    activeSection === item.id
                      ? "bg-blue-50 text-blue-600"
                      : "text-gray-700 hover:bg-gray-50"
                  }`}
                >
                  {item.label}
                </button>
              ))}
              <a
                href="https://github.com/quaadgras/graphics.gd"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center space-x-2 bg-gray-900 text-white px-3 py-2 rounded-lg hover:bg-gray-800 transition-colors mt-4"
              >
                <Github className="h-4 w-4" />
                <span>GitHub</span>
              </a>
            </div>
          </div>
        )}
      </nav>

      {/* Hero Section */}
      <section
        id="hero"
        className="pt-16 bg-gradient-to-br from-blue-50 via-white to-teal-50"
      >
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20">
          <div className="text-center">
            <div className="flex justify-center mb-8">
              <div className="flex items-center space-x-4 bg-blue-100 px-6 py-3 rounded-full">
                <Star className="h-5 w-5 text-blue-600" />
                <span className="text-blue-800 font-medium">
                  Open Source & Actively Maintained
                </span>
              </div>
            </div>

            <h1 className="text-5xl sm:text-6xl lg:text-7xl font-bold text-gray-900 mb-6 leading-tight">
              graphics.<span className="text-blue-600">gd</span>
            </h1>

            <p className="text-xl sm:text-2xl text-gray-600 mb-8 max-w-4xl mx-auto leading-relaxed">
              A cross platform 2D/3D graphics runtime for Go suitable for
              building native mobile apps, gdextensions, multimedia
              applications, games and more!
            </p>

            <div className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
              <button
                onClick={() => scrollToSection("getting-started")}
                className="flex items-center justify-center space-x-2 bg-blue-600 text-white px-8 py-4 rounded-xl hover:bg-blue-700 transition-all duration-200 hover:shadow-lg"
              >
                <Play className="h-5 w-5" />
                <span className="font-semibold">Get Started</span>
              </button>

              <a
                href="https://the.graphics.gd/guide"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center justify-center space-x-2 border border-gray-300 text-gray-700 px-8 py-4 rounded-xl hover:bg-gray-50 transition-all duration-200 hover:shadow-lg"
              >
                <ExternalLink className="h-5 w-5" />
                <span className="font-semibold">Documentation</span>
              </a>
            </div>

            <div className="bg-white/60 backdrop-blur-sm rounded-2xl p-8 max-w-4xl mx-auto border border-gray-200 shadow-lg">
              <div className="flex items-center space-x-2 mb-4">
                <div className="flex space-x-2">
                  <div className="w-3 h-3 bg-red-400 rounded-full"></div>
                  <div className="w-3 h-3 bg-yellow-400 rounded-full"></div>
                  <div className="w-3 h-3 bg-green-400 rounded-full"></div>
                </div>
                <span className="text-gray-500 text-sm font-medium">
                  main.go
                </span>
              </div>
              <pre className="text-left text-sm text-gray-800 overflow-x-auto">
                <code>{`package main

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
}`}</code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section id="features" className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Why choose graphics.gd?
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              A thoughtfully designed toolkit that brings the power of Go to
              game development
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <div className="bg-gradient-to-br from-blue-50 to-blue-100 p-8 rounded-2xl">
              <Code2 className="h-12 w-12 text-blue-600 mb-6" />
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Write Shaders in Go
              </h3>
              <p className="text-gray-700 leading-relaxed">
                Express complex graphics logic using familiar Go syntax instead
                of traditional shader languages.
              </p>
            </div>

            <div className="bg-gradient-to-br from-teal-50 to-teal-100 p-8 rounded-2xl">
              <Shield className="h-12 w-12 text-teal-600 mb-6" />
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Type Safety First
              </h3>
              <p className="text-gray-700 leading-relaxed">
                RIDs, callables, and dictionary arguments are all distinctly
                typed, preventing common runtime errors.
              </p>
            </div>

            <div className="bg-gradient-to-br from-green-50 to-green-100 p-8 rounded-2xl">
              <Zap className="h-12 w-12 text-green-600 mb-6" />
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Performance Focused
              </h3>
              <p className="text-gray-700 leading-relaxed">
                Balanced performance and convenience with minimal allocations
                for frequently called functions.
              </p>
            </div>

            <div className="bg-gradient-to-br from-purple-50 to-purple-100 p-8 rounded-2xl">
              <Coffee className="h-12 w-12 text-purple-600 mb-6" />
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Fast Iteration
              </h3>
              <p className="text-gray-700 leading-relaxed">
                After the first build, recompile quickly with an experience
                similar to scripting languages.
              </p>
            </div>

            <div className="bg-gradient-to-br from-orange-50 to-orange-100 p-8 rounded-2xl">
              <Users className="h-12 w-12 text-orange-600 mb-6" />
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Cross-Platform
              </h3>
              <p className="text-gray-700 leading-relaxed">
                Cross-compile builds targeting Windows, Linux, and macOS from
                any development platform.
              </p>
            </div>

            <div className="bg-gradient-to-br from-pink-50 to-pink-100 p-8 rounded-2xl">
              <Heart className="h-12 w-12 text-pink-600 mb-6" />
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Pure Go Packages
              </h3>
              <p className="text-gray-700 leading-relaxed">
                General-purpose 'variant' packages that you can reuse in any Go
                project beyond game development.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Code Example Section */}
      <section id="code" className="py-20 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Simple. Elegant. Powerful.
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Just one file is all you need to get started with graphics and
              game development in Go
            </p>
          </div>

          <div className="bg-white rounded-2xl shadow-xl overflow-hidden">
            <div className="bg-gray-900 px-6 py-4 flex items-center justify-between">
              <div className="flex items-center space-x-3">
                <div className="flex space-x-2">
                  <div className="w-3 h-3 bg-red-400 rounded-full"></div>
                  <div className="w-3 h-3 bg-yellow-400 rounded-full"></div>
                  <div className="w-3 h-3 bg-green-400 rounded-full"></div>
                </div>
                <span className="text-gray-300 font-medium">main.go</span>
              </div>
              <div className="text-gray-400 text-sm">Hello World Example</div>
            </div>
            <div className="p-8">
              <pre className="text-sm leading-relaxed overflow-x-auto">
                <code className="text-gray-800">
                  {`// This file is all you need to start a project.
// Save it somewhere, install the 'gd' command and use 'gd run' to get started.
package main

import (
    "graphics.gd/startup"

    "graphics.gd/classdb/Control"
    "graphics.gd/classdb/GUI"
    "graphics.gd/classdb/Label"
    "graphics.gd/classdb/SceneTree"
)

func main() {
    startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
    hello := Label.New()
    hello.AsControl().SetAnchorsPreset(Control.PresetFullRect) // expand the label to take up the whole screen.
    hello.SetHorizontalAlignment(GUI.HorizontalAlignmentCenter)
    hello.SetVerticalAlignment(GUI.VerticalAlignmentCenter)
    hello.SetText("Hello, World!")
    SceneTree.Add(hello)
    startup.Scene() // starts up the scene and blocks until the engine shuts down.
}`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Getting Started Section */}
      <section id="getting-started" className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Getting Started
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Get up and running with graphics.gd in minutes
            </p>
          </div>

          <div className="max-w-4xl mx-auto">
            <div className="space-y-8">
              <div className="flex items-start space-x-6">
                <div className="flex-shrink-0 w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center">
                  <span className="text-blue-600 font-bold text-lg">1</span>
                </div>
                <div className="flex-1">
                  <h3 className="text-2xl font-bold text-gray-900 mb-3">
                    Install the GD Command
                  </h3>
                  <p className="text-gray-600 mb-4 leading-relaxed">
                    Install the drop-in replacement for the go command that
                    makes working with graphics projects seamless.
                  </p>
                  <div className="bg-gray-900 rounded-lg p-4">
                    <code className="text-green-400 font-mono">
                      go install graphics.gd/cmd/gd@master
                    </code>
                  </div>
                </div>
              </div>

              <div className="flex items-start space-x-6">
                <div className="flex-shrink-0 w-12 h-12 bg-teal-100 rounded-full flex items-center justify-center">
                  <span className="text-teal-600 font-bold text-lg">2</span>
                </div>
                <div className="flex-1">
                  <h3 className="text-2xl font-bold text-gray-900 mb-3">
                    Create Your Project
                  </h3>
                  <p className="text-gray-600 mb-4 leading-relaxed">
                    Start with a simple main.go file. The tool will
                    automatically set up the graphics environment for you.
                  </p>
                  <div className="bg-gray-900 rounded-lg p-4 space-y-2">
                    <code className="text-green-400 font-mono block">
                      gd run # Run your project
                    </code>
                    <code className="text-green-400 font-mono block">
                      gd test # Run tests
                    </code>
                    <code className="text-green-400 font-mono block">
                      gd # Open the Engine Editor
                    </code>
                  </div>
                </div>
              </div>

              <div className="flex items-start space-x-6">
                <div className="flex-shrink-0 w-12 h-12 bg-green-100 rounded-full flex items-center justify-center">
                  <span className="text-green-600 font-bold text-lg">3</span>
                </div>
                <div className="flex-1">
                  <h3 className="text-2xl font-bold text-gray-900 mb-3">
                    Start Building
                  </h3>
                  <p className="text-gray-600 mb-4 leading-relaxed">
                    On Linux and macOS, the engine will be downloaded
                    automatically. On Windows, you may need to set up CGO first.
                  </p>
                  <div className="flex flex-wrap gap-4">
                    <a
                      href="https://the.graphics.gd/guide"
                      target="_blank"
                      rel="noopener noreferrer"
                      className="inline-flex items-center space-x-2 bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors"
                    >
                      <ExternalLink className="h-4 w-4" />
                      <span>the.graphics.gd/guide</span>
                    </a>
                    <a
                      href="https://github.com/quaadgras/graphics.gd/tree/samples"
                      target="_blank"
                      rel="noopener noreferrer"
                      className="inline-flex items-center space-x-2 border border-gray-300 text-gray-700 px-6 py-3 rounded-lg hover:bg-gray-50 transition-colors"
                    >
                      <ExternalLink className="h-4 w-4" />
                      <span>explore the sample projects</span>
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Performance Section */}
      <section
        id="performance"
        className="py-20 bg-gradient-to-br from-gray-50 to-blue-50"
      >
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Built for Performance
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Designed from the ground up to minimize allocations and maximize
              throughput
            </p>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div>
              <h3 className="text-3xl font-bold text-gray-900 mb-6">
                Zero-Allocation Method Calls
              </h3>
              <div className="space-y-6">
                <div className="flex items-start space-x-4">
                  <div className="flex-shrink-0 w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                    <ChevronRight className="h-4 w-4 text-green-600" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900 mb-2">
                      Advanced Instance Calls
                    </h4>
                    <p className="text-gray-600">
                      Go → Engine method calls do not allocate in practice when
                      using Advanced instances.
                    </p>
                  </div>
                </div>

                <div className="flex items-start space-x-4">
                  <div className="flex-shrink-0 w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                    <ChevronRight className="h-4 w-4 text-green-600" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900 mb-2">
                      Virtual Method Overrides
                    </h4>
                    <p className="text-gray-600">
                      Ready, Process, and other virtual methods avoid
                      allocations entirely.
                    </p>
                  </div>
                </div>

                <div className="flex items-start space-x-4">
                  <div className="flex-shrink-0 w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                    <ChevronRight className="h-4 w-4 text-blue-600" />
                  </div>
                  <div>
                    <h4 className="font-semibold text-gray-900 mb-2">
                      Frame-Based Memory Management
                    </h4>
                    <p className="text-gray-600">
                      Automatic garbage collection ensures memory safety without
                      manual management.
                    </p>
                  </div>
                </div>
              </div>
            </div>

            <div className="bg-white rounded-2xl shadow-lg p-8">
              <h4 className="text-2xl font-bold text-gray-900 mb-6">
                Performance Tips
              </h4>
              <div className="space-y-4">
                <div className="bg-green-50 p-4 rounded-lg border-l-4 border-green-400">
                  <p className="text-green-800 font-medium">
                    ✓ Use Engine types where possible
                  </p>
                </div>
                <div className="bg-green-50 p-4 rounded-lg border-l-4 border-green-400">
                  <p className="text-green-800 font-medium">
                    ✓ Avoid heap allocations in hot paths
                  </p>
                </div>
                <div className="bg-green-50 p-4 rounded-lg border-l-4 border-green-400">
                  <p className="text-green-800 font-medium">
                    ✓ Prefer Advanced methods for fine control
                  </p>
                </div>
                <div className="bg-green-50 p-4 rounded-lg border-l-4 border-green-400">
                  <p className="text-green-800 font-medium">
                    ✓ Use servers and systems instead of lots of objects.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Platforms Section */}
      <section id="platforms" className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Supported Platforms
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Deploy your games and applications across all major platforms
            </p>
          </div>

          <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-8 gap-6">
            {[
              { name: "Windows", status: "stable" },
              { name: "Linux", status: "stable" },
              { name: "macOS", status: "stable" },
              { name: "Android", status: "stable" },
              { name: "iOS", status: "stable" },
              { name: "Web", status: "stable" },
            ].map((platform) => (
              <div key={platform.name} className="text-center">
                <div className="bg-gray-100 rounded-2xl p-6 hover:bg-gray-200 transition-colors">
                  <div className="text-3xl mb-3">
                    {platform.name === "Windows" && "🪟"}
                    {platform.name === "Linux" && "🐧"}
                    {platform.name === "macOS" && "🍎"}
                    {platform.name === "Android" && "🤖"}
                    {platform.name === "iOS" && "📱"}
                    {platform.name === "Web" && "🌐"}
                    {platform.name === "Steam Deck" && "🎮"}
                    {platform.name === "Meta Quest" && "🥽"}
                  </div>
                  <h3 className="font-semibold text-gray-900 mb-2">
                    {platform.name}
                  </h3>
                  <div
                    className={`inline-flex items-center px-3 py-1 rounded-full text-xs font-medium ${
                      platform.status === "stable"
                        ? "bg-green-100 text-green-800"
                        : platform.status === "experimental"
                          ? "bg-yellow-100 text-yellow-800"
                          : "bg-gray-100 text-gray-800"
                    }`}
                  >
                    {platform.status}
                  </div>
                </div>
              </div>
            ))}
          </div>

          <div className="mt-12 text-center">
            <div className="bg-blue-50 rounded-xl p-6 max-w-2xl mx-auto">
              <h3 className="text-xl font-bold text-blue-900 mb-3">
                Platform Restrictions
              </h3>
              <div className="text-blue-800 space-y-2">
                <p>• 64-bit only (arm64 & amd64)</p>
                <p>
                  • Not available for Video Game Consoles (possible in future)
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Contributing Section */}
      <section
        id="contributing"
        className="py-20 bg-gradient-to-br from-teal-50 to-blue-50"
      >
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">
              Join Our Community
            </h2>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Help make graphics.gd even better for everyone
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mb-12">
            <div className="bg-white rounded-2xl p-8 shadow-lg hover:shadow-xl transition-shadow">
              <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-6">
                <Play className="h-6 w-6 text-blue-600" />
              </div>
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Try It Out
              </h3>
              <p className="text-gray-600 mb-6 leading-relaxed">
                The best way to contribute is to try graphics.gd and let us know
                what works and what doesn't.
              </p>
              <a
                href="https://github.com/quaadgras/graphics.gd/discussions"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium"
              >
                <span>Join Discussions</span>
                <ChevronRight className="h-4 w-4" />
              </a>
            </div>

            <div className="bg-white rounded-2xl p-8 shadow-lg hover:shadow-xl transition-shadow">
              <div className="w-12 h-12 bg-teal-100 rounded-full flex items-center justify-center mb-6">
                <Code2 className="h-6 w-6 text-teal-600" />
              </div>
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Improve Variants
              </h3>
              <p className="text-gray-600 mb-6 leading-relaxed">
                Help optimize the general-purpose Variant packages with better
                functionality and test coverage.
              </p>
              <a
                href="https://github.com/quaadgras/graphics.gd"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center space-x-2 text-teal-600 hover:text-teal-700 font-medium"
              >
                <span>View Source</span>
                <ChevronRight className="h-4 w-4" />
              </a>
            </div>

            <div className="bg-white rounded-2xl p-8 shadow-lg hover:shadow-xl transition-shadow">
              <div className="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mb-6">
                <Shield className="h-6 w-6 text-green-600" />
              </div>
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Add Tests
              </h3>
              <p className="text-gray-600 mb-6 leading-relaxed">
                Contribute tests to ensure graphics.gd remains stable and covers
                all the functionality you need.
              </p>
              <div className="bg-gray-100 rounded-lg p-3">
                <code className="text-gray-800 text-sm">
                  cd internal && gd test
                </code>
              </div>
            </div>
          </div>

          <div className="text-center">
            <div className="bg-white/80 backdrop-blur-sm rounded-2xl p-8 max-w-2xl mx-auto border border-teal-200">
              <h3 className="text-2xl font-bold text-gray-900 mb-4">
                Support Development
              </h3>
              <p className="text-gray-600 mb-6">
                Help fund the project, motivate development, and prioritize
                issues.
              </p>
              <a
                href="https://github.com/sponsors/Splizard"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center space-x-2 bg-gradient-to-r from-pink-500 to-orange-500 text-white px-8 py-4 rounded-xl hover:from-pink-600 hover:to-orange-600 transition-all duration-200 hover:shadow-lg"
              >
                <Heart className="h-5 w-5" />
                <span className="font-semibold">Support graphics.gd</span>
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-gray-900 text-white py-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
            <div className="col-span-1 md:col-span-2">
              <div className="flex items-center space-x-2 mb-6">
                <Gamepad2 className="h-8 w-8 text-blue-400" />
                <span className="text-2xl font-bold">graphics.gd</span>
              </div>
              <p className="text-gray-300 mb-6 max-w-md leading-relaxed">
                A safe, performant way to work with graphics and game
                development in Go, building on top of Godot and other Open
                Source GDExtension hosts.
              </p>
              <div className="flex space-x-4">
                <a
                  href="https://github.com/quaadgras/graphics.gd"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-gray-400 hover:text-white transition-colors"
                >
                  <Github className="h-6 w-6" />
                </a>
                <a
                  href="https://x.com/splizard"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-gray-400 hover:text-white transition-colors"
                >
                  <Twitter className="h-6 w-6" />
                </a>
              </div>
            </div>

            <div>
              <h4 className="text-lg font-semibold mb-4">Resources</h4>
              <div className="space-y-3">
                <a
                  href="https://the.graphics.gd/guide"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-gray-300 hover:text-white transition-colors"
                >
                  Documentation
                </a>
                <a
                  href="https://github.com/quaadgras/graphics.gd/tree/samples"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-gray-300 hover:text-white transition-colors"
                >
                  Sample Projects
                </a>
                <a
                  href="https://docs.godotengine.org/en/latest/"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-gray-300 hover:text-white transition-colors"
                >
                  Godot Docs
                </a>
              </div>
            </div>

            <div>
              <h4 className="text-lg font-semibold mb-4">Community</h4>
              <div className="space-y-3">
                <a
                  href="https://github.com/quaadgras/graphics.gd/discussions"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-gray-300 hover:text-white transition-colors"
                >
                  Discussions
                </a>
                <a
                  href="https://github.com/quaadgras/graphics.gd"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-gray-300 hover:text-white transition-colors"
                >
                  GitHub
                </a>
                <a
                  href="https://github.com/sponsors/Splizard"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="block text-gray-300 hover:text-white transition-colors"
                >
                  Support
                </a>
              </div>
            </div>
          </div>

          <div className="border-t border-gray-800 mt-12 pt-8">
            <div className="flex flex-col md:flex-row justify-between items-center">
              <p className="text-gray-400 mb-4 md:mb-0">
                © 2025 Quentin Quaadgras. Licensed under MIT License.
              </p>
              <p className="text-gray-400 text-sm">
                Built with ❤️ for the Go and graphics/game development community
              </p>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

export default App;
