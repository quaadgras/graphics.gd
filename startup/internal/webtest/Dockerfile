# Glibc image mirroring a GitHub `ubuntu-latest` runner closely enough to run a
# Godot web export in headless Chromium. Used to verify the web-test flow before
# wiring up CI. SwiftShader (Chromium's software WebGL2) is glibc-only, so this
# must NOT be a musl base.
# Node 22 base (global WebSocket/fetch the runner relies on) + Debian's real
# (glibc) Chromium with its SwiftShader software-WebGL2 backend.
FROM docker.io/library/node:22-bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y --no-install-recommends \
        chromium ca-certificates fonts-liberation \
    && rm -rf /var/lib/apt/lists/*

ENV CHROMIUM=/usr/bin/chromium
WORKDIR /work
ENTRYPOINT []
