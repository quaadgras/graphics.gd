// worker.mjs: the playground's go command (compiler.gd's, for GOOS=js, as
// build.sh makes it) in a worker of index.html, so that the page stays live
// while it compiles. Its files are in memory (memFS): GOROOT at /goroot;
// graphics.gd and the modules it needs at /gd, which the program's go.mod
// replaces them with; the program at /work; $HOME at /home, kept in
// IndexedDB, with the build cache, which starts as the one build.sh made.
//
// From the page: {run: source}. To the page: {line, fd} (fd 2 is the go
// command's output, 0 the playground's notes), {status, done?, total?}
// while loading, {ready}, {build: 'building' or how it went, failed?},
// {compiling: package}, {wasm: bytes} for a program that built, {finished}.
//
// The program does not run here: it runs beside the engine, in the page
// (run.html).
import { runGo } from './wasm_gc_browser.js';
import { memFS } from './wasm_gc_fs.js';
import { env, untar } from './files.mjs';
import { fetchBody, gunzip, bytes, keepOnly, loadManifest } from './fetch.mjs';

const post = (m, transfer) => postMessage(m, transfer || []);
const note = (line) => post({ line, fd: 0 });
const ms = (t0) => (performance.now() - t0 < 10000 ? Math.round(performance.now() - t0) + ' ms' : ((performance.now() - t0) / 1000).toFixed(1) + ' s');
const progress = (status, done, total) => post({ status, done, total });

// removeTree removes dir and what is in it.
function removeTree(fs, dir) {
  for (const e of fs.readdir(dir)) {
    if (e === '.' || e === '..') continue;
    const p = dir + '/' + e;
    if (fs.isDir(fs.stat(p).mode)) removeTree(fs, p);
    else fs.unlink(p);
  }
  fs.rmdir(dir);
}

const ready = (async () => {
  const t0 = performance.now();
  const manifest = await loadManifest();
  keepOnly([manifest.go, manifest.goroot, manifest.gd, manifest.cache, manifest.engine]).catch(() => {});
  const goP = (async () => WebAssembly.compileStreaming(new Response(gunzip(await fetchBody(manifest.go, 'the go command', progress)), { headers: { 'Content-Type': 'application/wasm' } })))();
  const gorootP = (async () => bytes(gunzip(await fetchBody(manifest.goroot, 'the standard library', progress))))();
  const gdP = (async () => bytes(gunzip(await fetchBody(manifest.gd, 'graphics.gd', progress))))();
  const fs = memFS();
  fs.mkdirTree('/home');
  await fs.persist('/home');
  // The cache of another version of the go command or of graphics.gd is
  // of no use.
  let version = '';
  try { version = fs.readFile('/home/version', { encoding: 'utf8' }); } catch {}
  const want = [manifest.go, manifest.gd, manifest.cache].join(' ');
  const fresh = version !== want;
  if (fresh) {
    try { removeTree(fs, '/home/cache'); } catch {}
    fs.mkdirTree('/home/cache');
    untar(fs, '/home/cache', await bytes(gunzip(await fetchBody(manifest.cache, 'the build cache', progress))));
    fs.writeFile('/home/version', want);
  }
  untar(fs, '/goroot', await gorootP);
  untar(fs, '/gd', await gdP);
  fs.mkdirTree('/work');
  fs.writeFile('/work/go.mod', fs.readFile('/gd/play.mod'));
  post({ status: 'compiling the go command' });
  const go = await goP;
  note(`[go command ready in ${ms(t0)}${fresh ? '' : ', build cache from an earlier visit'}]`);
  post({ ready: manifest.version });
  return { go, fs };
})();
ready.catch((e) => {
  post({ line: 'failed to load the go command: ' + ((e && e.message) || e), fd: 2 });
  post({ status: 'failed' });
});

onmessage = async ({ data }) => {
  try {
    const { go, fs } = await ready;
    fs.writeFile('/work/main.go', data.run);
    try { fs.unlink('/work/main.wasm'); } catch {}
    const t0 = performance.now();
    post({ build: 'building' });
    // -v prints each package as its compilation starts: a package that
    // is not in the cache takes seconds, and the page says which.
    const code = await runGo(go, {
      args: ['go', 'build', '-v', '-o', '/work/main.wasm', '.'], env, fs,
      stdout: (line) => post({ line, fd: 2 }),
      stderr: (line) => (/^[\w.\/-]+$/.test(line) ? post({ compiling: line }) : post({ line, fd: 2 })),
    });
    fs.sync().catch(() => {});
    if (code === 0) {
      const bin = fs.readFile('/work/main.wasm');
      post({ build: `built in ${ms(t0)} · ${(bin.length / 1e6).toFixed(1)} MB` });
      const copy = bin.slice();
      post({ wasm: copy.buffer }, [copy.buffer]);
    } else {
      post({ build: `build failed after ${ms(t0)}`, failed: true });
    }
  } catch (e) {
    // (Safari's stack does not have the message.)
    const stack = (e && e.stack) || '';
    post({ line: stack.includes(String(e)) ? stack : String(e) + (stack ? '\n' + stack : ''), fd: 2 });
    post({ build: 'failed', failed: true });
  }
  post({ finished: true });
};
