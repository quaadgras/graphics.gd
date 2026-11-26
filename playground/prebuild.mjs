// prebuild.mjs go.wasm goroot.tar.gz gd.tar.gz -cachedir dir ...: run the
// playground's go command under node as the page runs it (GOROOT at
// /goroot, graphics.gd's source at /gd, the cache at /home/cache, the
// program at /work), with the build cache in dir, which is read before
// and to which what the build added is written after (whether or not the
// build succeeded: a package can take a minute). The rest says what to
// build:
//
//	pkgs...                    those packages
//	-main main.go [-o out]     the program, as the page builds it
//	-deps file [-jobs n]       every package of file, which has a line for
//	                           each, the package and then its imports (as
//	                           go list -deps -f '{{.ImportPath}} {{join .Imports " "}}'
//	                           prints them)
//
// With -deps the packages are built a level at a time (a package's level is
// one more than the highest of its imports'), each level by n of these
// processes at once: nothing in a level imports anything else in it, so
// they share the cache directory. One process would do, but it takes half
// an hour for what a hello world imports, and the go command, which compiles
// in its own process here, does not give memory back between packages.
import fs from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import zlib from 'node:zlib';
import { spawn } from 'node:child_process';
import { fileURLToPath } from 'node:url';

const argv = process.argv.slice(2);
const [wasm, goroot, gd] = argv.splice(0, 3);
const opt = {};
while (argv[0] && argv[0].startsWith('-')) opt[argv.shift().slice(1)] = argv.shift();
if (!opt.cachedir) throw new Error('prebuild.mjs: -cachedir is required');
fs.mkdirSync(opt.cachedir, { recursive: true });

if (opt.deps) {
  // The levels.
  const imports = new Map();
  for (const line of fs.readFileSync(opt.deps, 'utf8').split('\n')) {
    const [pkg, ...deps] = line.trim().split(/\s+/);
    if (pkg && pkg !== 'unsafe') imports.set(pkg, deps.filter((d) => d !== 'unsafe' && d !== 'C'));
  }
  const level = new Map();
  const levelOf = (pkg) => {
    if (!imports.has(pkg)) return -1;
    if (!level.has(pkg)) level.set(pkg, 1 + Math.max(-1, ...imports.get(pkg).map(levelOf)));
    return level.get(pkg);
  };
  const levels = [];
  for (const pkg of imports.keys()) (levels[levelOf(pkg)] ||= []).push(pkg);
  const jobs = Number(opt.jobs) || Math.max(1, Math.min(8, os.cpus().length >> 1));
  const t0 = performance.now();
  const run = (pkgs) => new Promise((resolve, reject) => {
    const child = spawn(process.execPath, [...process.execArgv, fileURLToPath(import.meta.url), wasm, goroot, gd, '-cachedir', opt.cachedir, ...pkgs], { stdio: ['ignore', 'inherit', 'inherit'] });
    child.on('exit', (code) => (code === 0 ? resolve() : reject(new Error(`prebuild.mjs: exit ${code} building ${pkgs.join(' ')}`))));
  });
  for (const [i, pkgs] of levels.entries()) {
    const groups = Array.from({ length: Math.min(jobs, pkgs.length) }, () => []);
    pkgs.forEach((p, k) => groups[k % groups.length].push(p));
    console.error(`level ${i + 1} of ${levels.length}: ${pkgs.length} packages, ${((performance.now() - t0) / 1000).toFixed(0)} s so far`);
    await Promise.all(groups.map(run));
  }
  console.error(`prebuild: ${imports.size} packages in ${((performance.now() - t0) / 1000).toFixed(0)} s`);
  process.exit(0);
}

const { runGo } = await import('./wasm_gc_browser.js');
const { memFS } = await import('./wasm_gc_fs.js');
const { untar, env } = await import('./files.mjs');

const files = memFS();
untar(files, '/goroot', zlib.gunzipSync(fs.readFileSync(goroot)));
untar(files, '/gd', zlib.gunzipSync(fs.readFileSync(gd)));
// The cache, from the directory; had is what was there.
const had = new Set();
const load = (dir, to) => {
  files.mkdirTree(to);
  for (const e of fs.readdirSync(dir, { withFileTypes: true })) {
    if (e.isDirectory()) load(path.join(dir, e.name), `${to}/${e.name}`);
    else if (!e.name.includes('.tmp')) { files.writeFile(`${to}/${e.name}`, fs.readFileSync(path.join(dir, e.name))); had.add(`${to}/${e.name}`); }
  }
};
load(opt.cachedir, '/home/cache');
// save writes what the cache has gained (a file's name is a hash of what
// made it, so one that was there is the same; others may be writing the
// same directory: a file appears whole, by rename).
const save = (dir, from) => {
  fs.mkdirSync(dir, { recursive: true });
  for (const e of files.readdir(from)) {
    if (e === '.' || e === '..') continue;
    const p = `${from}/${e}`;
    if (files.isDir(files.stat(p).mode)) save(path.join(dir, e), p);
    else if (!had.has(p) || !/-[ad]$/.test(e)) {
      const tmp = path.join(dir, `${e}.tmp${process.pid}`);
      fs.writeFileSync(tmp, files.readFile(p));
      fs.renameSync(tmp, path.join(dir, e));
    }
  }
};
files.mkdirTree('/work');
files.writeFile('/work/go.mod', files.readFile('/gd/play.mod'));
if (opt.main) files.writeFile('/work/main.go', fs.readFileSync(opt.main));
// The page's global, but for its own fs, process and path (which
// runGo sets: node's own process is not the program's).
const own = {};
const global = new Proxy(globalThis, {
  get: (t, k) => (k in own ? own[k] : k === 'fs' || k === 'process' || k === 'path' ? undefined : Reflect.get(t, k)),
  set: (t, k, v) => { own[k] = v; return true; },
});
const args = opt.main ? ['go', 'build', '-v', '-o', '/work/main.wasm', '.'] : ['go', 'build', '-v', ...argv];
const t0 = performance.now();
const code = await runGo(fs.readFileSync(wasm), {
  args, env, fs: files, global,
  stdout: (l) => console.log(l), stderr: (l) => console.error(`${((performance.now() - t0) / 1000).toFixed(1)}s ${l}`),
});
save(opt.cachedir, '/home/cache');
if (opt.main || code !== 0) console.error(`go build: exit ${code} after ${Math.round(performance.now() - t0)} ms`);
if (code !== 0) process.exit(1);
if (opt.o) fs.writeFileSync(opt.o, files.readFile('/work/main.wasm'));
process.exit(0);
