// files.mjs: what the playground's worker (worker.mjs) and prebuild.mjs
// share: the go command's environment, and tar archives of files in
// memFS (lib/wasm/wasm_gc_fs.js).

// The environment of the go command: GOROOT at /goroot, $HOME (with
// the build cache) at /home, the program at /work.
export const env = {
  GOROOT: '/goroot',
  HOME: '/home',
  GOPATH: '/home/go',
  GOCACHE: '/home/cache',
  PWD: '/work',
  GOPROXY: 'off',
  GOTOOLCHAIN: 'local',
  GOTELEMETRY: 'off',
};

// untar writes the files of a ustar archive (bytes) into fs under dir,
// and returns how many there were. (Their times are now: the go command
// trims cache entries by their age.)
export function untar(fs, dir, tar) {
  const decoder = new TextDecoder();
  const text = (o, n) => {
    let end = o;
    while (end < o + n && tar[end] !== 0) end++;
    return decoder.decode(tar.subarray(o, end));
  };
  let n = 0;
  const made = new Set();
  for (let o = 0; o + 512 <= tar.length && tar[o] !== 0; ) {
    const name = [text(o + 345, 155), text(o, 100)].filter((s) => s).join('/');
    const size = parseInt(text(o + 124, 12).trim() || '0', 8);
    const kind = tar[o + 156] || 48;
    o += 512;
    if (kind === 48 /* '0' */) {
      const path = dir + '/' + name;
      const parent = path.slice(0, path.lastIndexOf('/'));
      if (!made.has(parent)) {
        fs.mkdirTree(parent);
        made.add(parent);
      }
      fs.writeFile(path, tar.slice(o, o + size));
      n++;
    }
    o += Math.ceil(size / 512) * 512;
  }
  return n;
}

// tar returns a ustar archive (bytes) of the files under dir in fs,
// named relative to dir.
export function tar(fs, dir) {
  const encoder = new TextEncoder();
  const parts = [];
  const header = (name, size) => {
    const h = new Uint8Array(512);
    let prefix = '';
    if (name.length > 100) {
      const cut = name.lastIndexOf('/', 155);
      prefix = name.slice(0, cut);
      name = name.slice(cut + 1);
    }
    const put = (o, s) => h.set(encoder.encode(s), o);
    put(0, name);
    put(100, '0000644\0');
    put(108, '0000000\0');
    put(116, '0000000\0');
    put(124, size.toString(8).padStart(11, '0') + '\0');
    put(136, '00000000000\0');
    put(148, '        ');
    h[156] = 48;
    put(257, 'ustar\0');
    put(263, '00');
    put(345, prefix);
    let sum = 0;
    for (const b of h) sum += b;
    put(148, sum.toString(8).padStart(6, '0') + '\0 ');
    return h;
  };
  const walk = (rel) => {
    for (const e of fs.readdir(dir + rel).sort()) {
      if (e === '.' || e === '..') continue;
      const path = rel + '/' + e;
      if (fs.isDir(fs.stat(dir + path).mode)) {
        walk(path);
        continue;
      }
      const data = fs.readFile(dir + path);
      parts.push(header(path.slice(1), data.length), data, new Uint8Array((512 - (data.length % 512)) % 512));
    }
  };
  walk('');
  parts.push(new Uint8Array(1024));
  const out = new Uint8Array(parts.reduce((n, p) => n + p.length, 0));
  let o = 0;
  for (const p of parts) {
    out.set(p, o);
    o += p.length;
  }
  return out;
}
