#!/bin/sh
# Build ld64.lld as a wasm32-wasip1 module: native tblgen first, then
# the cross build with threads off and only the AArch64 target.
set -ex
W=$LLDWASM
cd $W/llvm-project

cmake -S llvm -B build-native -G Ninja \
    -DCMAKE_BUILD_TYPE=Release \
    -DLLVM_TARGETS_TO_BUILD=AArch64 \
    -DLLVM_INCLUDE_TESTS=OFF -DLLVM_INCLUDE_BENCHMARKS=OFF -DLLVM_INCLUDE_EXAMPLES=OFF
ninja -C build-native llvm-tblgen

cmake -S llvm -B build-wasm -G Ninja \
    -DCMAKE_TOOLCHAIN_FILE=$W/wasi-toolchain.cmake \
    -DUNIX=1 \
    -DCMAKE_BUILD_TYPE=MinSizeRel \
    -DLLVM_ENABLE_PROJECTS=lld \
    -DLLVM_TARGETS_TO_BUILD=AArch64 \
    -DLLVM_HOST_TRIPLE=wasm32-unknown-wasip1 \
    -DLLVM_DEFAULT_TARGET_TRIPLE=arm64-apple-ios \
    -DLLVM_ENABLE_THREADS=OFF \
    -DLLVM_ENABLE_PIC=OFF \
    -DLLVM_ENABLE_ZLIB=OFF -DLLVM_ENABLE_ZSTD=OFF \
    -DLLVM_ENABLE_LIBXML2=OFF -DLLVM_ENABLE_TERMINFO=OFF \
    -DLLVM_ENABLE_LIBEDIT=OFF \
    -DLLVM_INCLUDE_TESTS=OFF -DLLVM_INCLUDE_BENCHMARKS=OFF -DLLVM_INCLUDE_EXAMPLES=OFF \
    -DLLVM_NATIVE_TOOL_DIR=$W/llvm-project/build-native/bin
ninja -C build-wasm lld
ls -la build-wasm/bin/
