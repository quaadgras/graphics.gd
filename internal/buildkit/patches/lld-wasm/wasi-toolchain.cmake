# Cross-compile to wasm32-wasip1 using the host clang with the
# wasi-sdk sysroot (the sdk's own binaries are glibc-linked and do not
# run on this musl host). The shim resource dir maps clang 21's headers
# together with the sdk's wasm builtins archive.
set(LLDWASM $LLDWASM)
set(WASI_SDK ${LLDWASM}/wasi-sdk-25.0-x86_64-linux)

list(APPEND CMAKE_MODULE_PATH ${WASI_SDK}/share/cmake)
set(CMAKE_SYSTEM_NAME WASI)
set(CMAKE_SYSTEM_VERSION 1)
set(CMAKE_SYSTEM_PROCESSOR wasm32)

set(CMAKE_C_COMPILER /usr/bin/clang)
set(CMAKE_CXX_COMPILER /usr/bin/clang++)
set(CMAKE_ASM_COMPILER /usr/bin/clang)
set(triple wasm32-wasip1)
set(CMAKE_C_COMPILER_TARGET ${triple})
set(CMAKE_CXX_COMPILER_TARGET ${triple})
set(CMAKE_ASM_COMPILER_TARGET ${triple})
set(CMAKE_SYSROOT ${WASI_SDK}/share/wasi-sysroot)

add_compile_options(
    -isystem${LLDWASM}/shim-include
    -resource-dir=${LLDWASM}/resource
    -D_WASI_EMULATED_MMAN
    -D_WASI_EMULATED_SIGNAL
    -D_WASI_EMULATED_PROCESS_CLOCKS
    -D_WASI_EMULATED_GETPID)
add_link_options(
    -resource-dir=${LLDWASM}/resource
    -lwasi-emulated-mman
    -lwasi-emulated-signal
    -lwasi-emulated-process-clocks
    -lwasi-emulated-getpid
    -Wl,-z,stack-size=8388608
    -Wl,--max-memory=4294967296)

set(CMAKE_FIND_ROOT_PATH_MODE_PROGRAM NEVER)
set(CMAKE_FIND_ROOT_PATH_MODE_LIBRARY ONLY)
set(CMAKE_FIND_ROOT_PATH_MODE_INCLUDE ONLY)
set(CMAKE_FIND_ROOT_PATH_MODE_PACKAGE ONLY)

# wasm-ld has no rpath concept; stop cmake from emitting -rpath-link.
set(CMAKE_SKIP_RPATH ON)
set(CMAKE_EXECUTABLE_RUNTIME_C_FLAG "")
set(CMAKE_EXECUTABLE_RUNTIME_CXX_FLAG "")
set(CMAKE_SHARED_LIBRARY_RUNTIME_C_FLAG "")
set(CMAKE_SHARED_LIBRARY_RUNTIME_CXX_FLAG "")

add_link_options(${LLDWASM}/wasi-compat.o)
