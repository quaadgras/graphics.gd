# gd's SCons site file, for building an engine on a host its build system
# does not expect: nothing about a build against zig's toolchain depends on
# the host, but the engine's platform detection and SCons's defaults do.
import importlib.abc
import importlib.machinery
import importlib.util
import os
import re
import subprocess
import sys
import types

import SCons.Platform
import SCons.Tool
import SCons.Tool.default


def asked_by_linuxbsd():
    # the caller of the caller: frame 0 is us, 1 the lie, 2 who asked.
    name = sys._getframe(2).f_code.co_filename.replace("\\", "/")
    return name.endswith("/platform/linuxbsd/detect.py")


# The linux build refuses to configure unless the host is a posix system
# other than macOS, with pkg-config installed (platform/linuxbsd/detect.py
# can_build), even when the system's libraries are loaded at runtime (the
# use_sowrap default) and it needs neither. Its detection alone is told
# what it wants to hear, with sys.platform lied about only for it, since
# the standard library and SCons decide how to run things by it.
class OS(types.ModuleType):
    @property
    def name(self):
        return "posix" if asked_by_linuxbsd() else self.__dict__["name"]

    @property
    def system(self):
        real = self.__dict__["system"]

        def system(command):
            if asked_by_linuxbsd() and command.startswith("pkg-config --version"):
                return 0
            return real(command)

        return system


class Sys(types.ModuleType):
    @property
    def platform(self):
        return "linux" if asked_by_linuxbsd() else self.__dict__["platform"]


os.__class__ = OS
sys.__class__ = Sys


def target_platform():
    for arg in sys.argv:
        if arg.startswith("platform="):
            return arg[len("platform=") :]
    return ""


# A Windows host builds every other platform's engine with the toolchain
# SCons prefers on posix systems: the engine's build (which knows the
# compilers it wants) only asks for the "default" tools, which on Windows
# are Microsoft's, with their flags, file names and no idea what an
# Objective-C file is. Commands are run without cmd.exe (and its 8K limit
# on a command line), archives are made from response files, as the
# engine's own windows build does.
if os.name == "nt" and target_platform() not in ("", "windows"):
    tool_list = SCons.Tool.tool_list

    def posix_tool_list(platform, env):
        return tool_list("posix" if str(platform) == "win32" else platform, env)

    SCons.Tool.tool_list = posix_tool_list

    def spawn(sh, escape, cmd, args, env):
        if cmd == "del":  # response files are removed through the spawn.
            os.remove(args[1])
            return 0
        env = {str(key): str(value) for key, value in env.items()}
        stdin = None
        if "<" in args:  # the one shell feature the build uses, for ar's MRI scripts.
            args, stdin = args[: args.index("<")], open(args[args.index("<") + 1].strip('"'), "rb")
        try:
            return subprocess.call(cmd + " " + " ".join(args[1:]), env=env, shell=False, stdin=stdin)
        finally:
            if stdin:
                stdin.close()

    def response_file_arg(arg):
        # the response file is read as a GNU command line, where a
        # backslash escapes, so paths take the other separator.
        from SCons.Subst import quote_spaces

        return re.sub(r"\\([^\"'\\]|$)", r"/\1", quote_spaces(arg))

    generate = SCons.Tool.default.generate

    def posix_generate(env):
        generate(env)
        if env["PLATFORM"] != "win32":
            return
        env["OBJSUFFIX"] = ".o"
        env["PROGSUFFIX"] = ""
        env["LIBPREFIX"] = "lib"
        env["LIBSUFFIX"] = ".a"
        env["SPAWN"] = spawn
        # a case-insensitive file system makes .S (assembly to preprocess)
        # the same as .s, for which the assembler is run bare: the engine's
        # assembly is all for the compiler, as any other host would use.
        env["ASCOM"] = "$CC $ASFLAGS $ASPPFLAGS $CPPFLAGS $_CPPDEFFLAGS $_CPPINCFLAGS -c -o $TARGET $SOURCES"
        env["ARCOM"] = "${TEMPFILE('$AR $ARFLAGS $TARGET $SOURCES', '$ARCOMSTR')}"
        env["TEMPFILESUFFIX"] = ".rsp"
        env["TEMPFILEARGESCFUNC"] = response_file_arg

    SCons.Tool.default.generate = posix_generate

    # The Swift builder (platform_methods.setup_swift_builder) names the
    # other sources by joining the current path with "/" and removes the
    # one being compiled by its native path, which never match here: the
    # current path is made to join natively.
    class NativePath(str):
        def __add__(self, other):
            return NativePath(str.__add__(self, other).replace("/", os.path.sep))

    def native_swift_builder(module):
        setup_swift_builder = getattr(module, "setup_swift_builder", None)
        if setup_swift_builder is None:
            return

        def setup(env, *args, **kwargs):
            setup_swift_builder(env, *args, **kwargs)
            env["CURRENT_PATH"] = NativePath(env["CURRENT_PATH"])

        module.setup_swift_builder = setup

    def patched_spec(spec):
        if spec is None or spec.loader is None or spec.name != "platform_methods":
            return spec
        exec_module = spec.loader.exec_module

        def patched(module):
            exec_module(module)
            native_swift_builder(module)

        spec.loader.exec_module = patched
        return spec

    # the engine's SConstruct loads its helper modules by path.
    spec_from_file_location = importlib.util.spec_from_file_location

    def patched_spec_from_file_location(name, *args, **kwargs):
        return patched_spec(spec_from_file_location(name, *args, **kwargs))

    importlib.util.spec_from_file_location = patched_spec_from_file_location

    class PlatformMethods(importlib.abc.MetaPathFinder):
        def find_spec(self, name, path, target=None):
            if name != "platform_methods":
                return None
            return patched_spec(importlib.machinery.PathFinder.find_spec(name, path))

    sys.meta_path.insert(0, PlatformMethods())
