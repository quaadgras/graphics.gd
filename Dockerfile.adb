# Build static adb and fastboot binaries using Alpine Linux (musl)
# Based on nmeum/android-tools which provides CMake build system
FROM alpine:edge AS builder

# Install build dependencies (minimal - we build libs from source)
RUN apk add --no-cache \
    build-base \
    cmake \
    git \
    go \
    linux-headers \
    perl \
    samurai \
    autoconf \
    automake \
    libtool \
    pkgconfig

# Download android-tools source
ARG ANDROID_TOOLS_VERSION=35.0.2
RUN wget -q https://github.com/nmeum/android-tools/releases/download/${ANDROID_TOOLS_VERSION}/android-tools-${ANDROID_TOOLS_VERSION}.tar.xz \
    && tar -xf android-tools-${ANDROID_TOOLS_VERSION}.tar.xz \
    && mv android-tools-${ANDROID_TOOLS_VERSION} /android-tools

WORKDIR /android-tools

# Build with static linking
# Build zlib as static library from source
RUN git clone --depth 1 --branch v1.3.1 https://github.com/madler/zlib.git /zlib && \
    cd /zlib && \
    ./configure --static --prefix=/usr/local && \
    make -j$(nproc) && make install

# Build brotli as static library from source (all components)
RUN git clone --depth 1 --branch v1.1.0 https://github.com/google/brotli.git /brotli && \
    cd /brotli && mkdir build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release \
        -DBUILD_SHARED_LIBS=OFF \
        -DCMAKE_INSTALL_PREFIX=/usr/local \
        -DBROTLI_DISABLE_TESTS=ON && \
    make -j$(nproc) && make install && \
    # Create pkg-config files for brotli (CMake doesn't generate them)
    mkdir -p /usr/local/lib/pkgconfig && \
    printf 'prefix=/usr/local\nlibdir=${prefix}/lib\nincludedir=${prefix}/include\n\nName: libbrotlicommon\nDescription: Brotli common library\nVersion: 1.1.0\nLibs: -L${libdir} -lbrotlicommon\nCflags: -I${includedir}\n' > /usr/local/lib/pkgconfig/libbrotlicommon.pc && \
    printf 'prefix=/usr/local\nlibdir=${prefix}/lib\nincludedir=${prefix}/include\n\nName: libbrotlidec\nDescription: Brotli decoder library\nVersion: 1.1.0\nRequires: libbrotlicommon\nLibs: -L${libdir} -lbrotlidec\nCflags: -I${includedir}\n' > /usr/local/lib/pkgconfig/libbrotlidec.pc && \
    printf 'prefix=/usr/local\nlibdir=${prefix}/lib\nincludedir=${prefix}/include\n\nName: libbrotlienc\nDescription: Brotli encoder library\nVersion: 1.1.0\nRequires: libbrotlicommon\nLibs: -L${libdir} -lbrotlienc\nCflags: -I${includedir}\n' > /usr/local/lib/pkgconfig/libbrotlienc.pc

# Build lz4 as static library from source
RUN git clone --depth 1 --branch v1.10.0 https://github.com/lz4/lz4.git /lz4 && \
    cd /lz4 && \
    make -j$(nproc) BUILD_SHARED=no && \
    make install BUILD_SHARED=no PREFIX=/usr/local

# Build zstd as static library from source
RUN git clone --depth 1 --branch v1.5.6 https://github.com/facebook/zstd.git /zstd && \
    cd /zstd/build/cmake && \
    cmake . -DCMAKE_BUILD_TYPE=Release -DZSTD_BUILD_SHARED=OFF -DZSTD_BUILD_STATIC=ON -DZSTD_BUILD_PROGRAMS=OFF && \
    make -j$(nproc) && make install

# Build libusb as static library from source (use master for SuperSpeed Plus APIs)
# adb requires libusb_ssplus_* functions not in 1.0.27 release
RUN git clone --depth 1 https://github.com/libusb/libusb.git /libusb && \
    cd /libusb && \
    ./autogen.sh --prefix=/usr/local --enable-static --disable-shared --disable-udev && \
    make -j$(nproc) && make install

# Build googletest as static library from source
RUN git clone --depth 1 --branch v1.15.2 https://github.com/google/googletest.git /googletest && \
    cd /googletest && mkdir build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=OFF && \
    make -j$(nproc) && make install

# Build fmt as static library from source
RUN git clone --depth 1 --branch 11.0.2 https://github.com/fmtlib/fmt.git /fmt && \
    cd /fmt && mkdir build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DFMT_DOC=OFF -DFMT_TEST=OFF -DBUILD_SHARED_LIBS=OFF && \
    make -j$(nproc) && make install

# Build pcre2 as static library from source
RUN git clone --depth 1 --branch pcre2-10.44 https://github.com/PCRE2Project/pcre2.git /pcre2 && \
    cd /pcre2 && mkdir build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=OFF -DPCRE2_BUILD_PCRE2GREP=OFF -DPCRE2_BUILD_TESTS=OFF && \
    make -j$(nproc) && make install

# Build protobuf as static library from source (using bundled abseil)
RUN git clone --depth 1 --branch v25.6 --recurse-submodules https://github.com/protocolbuffers/protobuf.git /protobuf && \
    cd /protobuf && \
    cmake . -DCMAKE_BUILD_TYPE=Release \
        -DBUILD_SHARED_LIBS=OFF \
        -Dprotobuf_BUILD_TESTS=OFF \
        -Dprotobuf_ABSL_PROVIDER=module && \
    make -j$(nproc) && make install

ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:/usr/lib/pkgconfig

RUN mkdir build && cd build && \
    cmake .. \
        -G Ninja \
        -DCMAKE_BUILD_TYPE=Release \
        -DCMAKE_EXE_LINKER_FLAGS="-static -static-libgcc -static-libstdc++" \
        -DCMAKE_FIND_LIBRARY_SUFFIXES=".a" \
        -DBUILD_SHARED_LIBS=OFF && \
    ninja adb

# Verify static linking (static-pie on musl shows ld-musl but is actually static)
RUN file /android-tools/build/vendor/adb && \
    file /android-tools/build/vendor/adb | grep -q "static" && echo "adb is static"

# Strip binary
RUN strip /android-tools/build/vendor/adb

# Create output directory
RUN mkdir -p /output && \
    cp /android-tools/build/vendor/adb /output/adb.linux.amd64

# Final stage - export binaries
FROM scratch
COPY --from=builder /output /build

# Build and extract with:
#   docker build -f Dockerfile.adb --output=. .
