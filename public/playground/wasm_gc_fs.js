// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// wasm_gc_fs.js gives a GOWASM=gc program (GOOS=wasip1) files where
// there is no system to ask for them: in a browser. It does so the way
// Emscripten does, and through Emscripten's interface:
//
//   - wasiFS(fs, mem) makes the file calls of wasi_snapshot_preview1
//     (what package os and syscall use) from an object with the
//     interface of Emscripten's FS: open, read, write, llseek, close,
//     stat, readdir, mkdir, rename... It can be the FS of an Emscripten
//     program on the same page (Module.FS), and the Go program then has
//     that program's files, whatever they are mounted on.
//
//   - memFS() is such an object for when there is no Emscripten: files
//     in memory, as Emscripten's MEMFS, with /tmp, /home/web_user and
//     /dev/null. fs.persist(dir) keeps a directory in IndexedDB as
//     Emscripten's IDBFS mounted there does (the same database, store
//     and records), so the two read what the other wrote.
//
// wasm_gc_browser.js uses both: see runGo's options fs, persist and
// files.

const S_IFMT = 0o170000, S_IFDIR = 0o040000, S_IFREG = 0o100000, S_IFLNK = 0o120000, S_IFCHR = 0o020000;
const isDir = (mode) => (mode & S_IFMT) === S_IFDIR;
const isFile = (mode) => (mode & S_IFMT) === S_IFREG;
const isLink = (mode) => (mode & S_IFMT) === S_IFLNK;

// The flags of open, as Emscripten's libc has them.
const O_ACCMODE = 3, O_WRONLY = 1, O_RDWR = 2, O_CREAT = 0o100, O_EXCL = 0o200, O_TRUNC = 0o1000, O_APPEND = 0o2000,
  O_DIRECTORY = 0o200000, O_NOFOLLOW = 0o400000;

// The errors, which Emscripten numbers as WASI does.
const EBADF = 8, EBUSY = 10, EEXIST = 20, EINVAL = 28, EISDIR = 31, ELOOP = 32, ENOENT = 44, ENOTDIR = 54,
  ENOTEMPTY = 55, ENOTSUP = 58, EPERM = 63, ENOTCAPABLE = 76;

const encoder = new TextEncoder(), decoder = new TextDecoder();

// wasiFS returns the calls of wasi_snapshot_preview1 for files, made
// with fs (the interface of Emscripten's FS). mem returns the buffer of
// the program's memory. The root of fs is the one directory the program
// is given (descriptor 3); descriptors 0 to 2 are not its business.
//
// Beyond Emscripten's interface it uses, if fs has them: link(old, new);
// utimens(path, atime, mtime, follow) and futimens(fd, atime, mtime),
// with times in nanoseconds (null: unchanged); and atimeNs, mtimeNs and
// ctimeNs in what stat returns. (Emscripten keeps milliseconds.)
export function wasiFS(fs, mem) {
  const fds = new Map(); // descriptor -> { stream, dir, list (what fd_readdir reads) }
  const open = (stream) => {
    let fd = 3;
    while (fds.has(fd)) fd++;
    fds.set(fd, { stream, dir: isDir(stream.node.mode), list: null });
    return fd;
  };
  open(fs.open('/', O_DIRECTORY));

  const view = () => new DataView(mem());
  const heap = () => new Uint8Array(mem());
  const fail = (errno) => { throw { errno }; };
  // (An error of fs has its number: FS.ErrnoError.)
  const call = (f) => (...a) => {
    try {
      return f(...a) || 0;
    } catch (e) {
      if (e && typeof e.errno === 'number') return e.errno;
      throw e;
    }
  };
  const get = (fd) => fds.get(fd) || fail(EBADF);
  const dirOf = (fd) => {
    const e = get(fd);
    if (!e.dir) fail(ENOTDIR);
    // (Where the directory is now, if it was renamed.)
    return fs.getPath ? fs.getPath(e.stream.node) : e.stream.path;
  };
  // path returns the path in fs of the n bytes at p, a path relative to
  // the directory fd, which it must not leave.
  const path = (fd, p, n) => {
    const rel = decoder.decode(new Uint8Array(mem(), p, n));
    if (rel === '') fail(ENOENT);
    if (rel.includes('\0')) fail(EINVAL);
    if (rel.startsWith('/')) fail(ENOTCAPABLE);
    let depth = 0;
    for (const part of rel.split('/')) {
      if (part === '..') {
        if (--depth < 0) fail(ENOTCAPABLE);
      } else if (part !== '' && part !== '.') {
        depth++;
      }
    }
    const dir = dirOf(fd);
    return (dir === '/' ? '' : dir) + '/' + rel;
  };
  const lstat = (file) => (fs.lstat ? fs.lstat(file) : fs.stat(file, true));
  const exists = (file) => {
    try {
      lstat(file);
      return true;
    } catch (e) {
      if (e && typeof e.errno === 'number') return false;
      throw e;
    }
  };
  const fstat = (stream) => (fs.fstat ? fs.fstat(stream.fd) : stream.node.node_ops.getattr(stream.node));
  const filetype = (mode) => (isDir(mode) ? 3 : isFile(mode) ? 4 : isLink(mode) ? 7 : (mode & S_IFMT) === S_IFCHR ? 2 : 0);
  const ns = (st, name) => st[name + 'Ns'] ?? BigInt(Math.round(+st[name])) * 1000000n;
  const putStat = (p, st) => {
    const v = view();
    v.setBigUint64(p, BigInt(st.dev || 0), true);
    v.setBigUint64(p + 8, BigInt(st.ino || 0), true);
    v.setUint8(p + 16, filetype(st.mode));
    v.setBigUint64(p + 24, BigInt(st.nlink || 1), true);
    v.setBigUint64(p + 32, BigInt(st.size), true);
    v.setBigUint64(p + 40, ns(st, 'atime'), true);
    v.setBigUint64(p + 48, ns(st, 'mtime'), true);
    v.setBigUint64(p + 56, ns(st, 'ctime'), true);
  };
  // times returns the times, in nanoseconds, that the flags fst ask for
  // (ATIM, ATIM_NOW, MTIM, MTIM_NOW): null for one that stays.
  const times = (atim, mtim, fst) => {
    const now = BigInt(Date.now()) * 1000000n;
    return [fst & 2 ? now : fst & 1 ? atim : null, fst & 8 ? now : fst & 4 ? mtim : null];
  };
  const setTimes = (file, follow, a, m) => {
    if (fs.utimens) return fs.utimens(file, a, m, follow);
    const st = follow ? fs.stat(file) : lstat(file);
    if (isLink(st.mode)) return; // (Emscripten has no way to set a link's.)
    const ms = (t, old) => (t === null ? +old : Number(t / 1000000n));
    fs.utime(file, ms(a, st.atime), ms(m, st.mtime));
  };
  const iovecs = (iovs, n) => {
    const v = view(), list = [];
    for (let i = 0; i < n; i++) list.push([v.getUint32(iovs + 8 * i, true), v.getUint32(iovs + 8 * i + 4, true)]);
    return list;
  };
  const read = (fd, iovs, n, at, nreadp) => {
    const e = get(fd);
    if (e.dir) fail(EISDIR);
    let total = 0;
    for (const [p, l] of iovecs(iovs, n)) {
      const k = fs.read(e.stream, heap(), p, l, at === undefined ? undefined : at + total);
      total += k;
      if (k < l) break;
    }
    view().setUint32(nreadp, total, true);
  };
  const write = (fd, iovs, n, at, nwrittenp) => {
    const e = get(fd);
    if (e.dir) fail(EISDIR);
    let total = 0;
    for (const [p, l] of iovecs(iovs, n)) {
      if (l === 0) continue;
      const k = fs.write(e.stream, heap(), p, l, at === undefined ? undefined : at + total);
      total += k;
      if (k < l) break;
    }
    view().setUint32(nwrittenp, total, true);
  };

  const calls = {
    fd_prestat_get: (fd, p) => {
      if (fd !== 3 || !fds.has(3)) fail(EBADF);
      view().setUint8(p, 0); // a directory
      view().setUint32(p + 4, 1, true);
    },
    fd_prestat_dir_name: (fd, p, len) => {
      if (fd !== 3 || !fds.has(3)) fail(EBADF);
      if (len > 0) heap()[p] = 47; // "/"
    },
    path_open: (dirfd, lookup, p, n, oflags, base, inheriting, fdflags, fdp) => {
      const file = path(dirfd, p, n);
      // (The rights FD_READ and FD_WRITE say how it is opened.)
      const reads = base & 2n, writes = base & 64n;
      let flags = writes ? (reads ? O_RDWR : O_WRONLY) : 0;
      if (oflags & 1) flags |= O_CREAT;
      if (oflags & 2) flags |= O_DIRECTORY;
      if (oflags & 4) flags |= O_EXCL;
      if (oflags & 8) flags |= O_TRUNC;
      if (fdflags & 1) flags |= O_APPEND;
      if (!(lookup & 1)) flags |= O_NOFOLLOW;
      // (A new file is not made through a link that is there: not every
      // fs refuses to.)
      if ((oflags & 1) && (oflags & 4) && exists(file)) fail(EEXIST);
      view().setUint32(fdp, open(fs.open(file, flags, 0o666)), true);
    },
    fd_close: (fd) => {
      fs.close(get(fd).stream);
      fds.delete(fd);
    },
    fd_read: (fd, iovs, n, nreadp) => read(fd, iovs, n, undefined, nreadp),
    fd_pread: (fd, iovs, n, offset, nreadp) => read(fd, iovs, n, Number(offset), nreadp),
    fd_write: (fd, iovs, n, nwrittenp) => write(fd, iovs, n, undefined, nwrittenp),
    fd_pwrite: (fd, iovs, n, offset, nwrittenp) => write(fd, iovs, n, Number(offset), nwrittenp),
    fd_seek: (fd, offset, whence, newp) => {
      const e = get(fd);
      let at = 0;
      if (e.dir) {
        e.list = null; // (read again from the start)
      } else {
        at = fs.llseek(e.stream, Number(offset), whence);
      }
      view().setBigUint64(newp, BigInt(at), true);
    },
    fd_tell: (fd, p) => {
      const e = get(fd);
      view().setBigUint64(p, BigInt(e.dir ? 0 : fs.llseek(e.stream, 0, 1)), true);
    },
    fd_fdstat_get: (fd, statp) => {
      const e = get(fd), v = view();
      v.setUint8(statp, filetype(e.stream.node.mode));
      v.setUint16(statp + 2, e.stream.flags & O_APPEND ? 1 : 0, true);
      v.setBigUint64(statp + 8, 0xffffffffffffffffn, true);
      v.setBigUint64(statp + 16, 0xffffffffffffffffn, true);
    },
    fd_fdstat_set_flags: (fd, flags) => {
      const s = get(fd).stream;
      s.flags = (s.flags & ~O_APPEND) | (flags & 1 ? O_APPEND : 0);
    },
    fd_fdstat_set_rights: (fd) => { get(fd); },
    fd_sync: (fd) => { get(fd); },
    fd_datasync: (fd) => { get(fd); },
    fd_advise: (fd) => { get(fd); },
    fd_allocate: (fd) => { get(fd); },
    fd_filestat_get: (fd, p) => putStat(p, fstat(get(fd).stream)),
    fd_filestat_set_size: (fd, size) => {
      const e = get(fd);
      if (e.dir) fail(EISDIR);
      fs.ftruncate(e.stream.fd, Number(size));
    },
    fd_filestat_set_times: (fd, atim, mtim, fst) => {
      const s = get(fd).stream, [a, m] = times(atim, mtim, fst);
      if (fs.futimens) return fs.futimens(s.fd, a, m);
      setTimes(fs.getPath ? fs.getPath(s.node) : s.path, true, a, m);
    },
    // An entry is where the next one is (a cookie), the file's number,
    // the length of its name, its type, and the name; the last may be
    // cut short by the end of the buffer, which is how the program
    // knows there is more.
    fd_readdir: (fd, buf, len, cookie, usedp) => {
      const e = get(fd), dir = dirOf(fd);
      if (cookie === 0n || !e.list) {
        e.list = [];
        for (const name of fs.readdir(dir)) {
          if (name === '.' || name === '..') continue;
          let st;
          try {
            st = lstat((dir === '/' ? '' : dir) + '/' + name);
          } catch {
            continue;
          }
          e.list.push({ name: encoder.encode(name), ino: st.ino || 0, type: filetype(st.mode) });
        }
      }
      let used = 0;
      for (let i = Number(cookie); i < e.list.length && used < len; i++) {
        const d = e.list[i], rec = new Uint8Array(24 + d.name.length), v = new DataView(rec.buffer);
        v.setBigUint64(0, BigInt(i + 1), true);
        v.setBigUint64(8, BigInt(d.ino), true);
        v.setUint32(16, d.name.length, true);
        v.setUint8(20, d.type);
        rec.set(d.name, 24);
        const k = Math.min(rec.length, len - used);
        heap().set(rec.subarray(0, k), buf + used);
        used += k;
      }
      view().setUint32(usedp, used, true);
    },
    path_create_directory: (fd, p, n) => { fs.mkdir(path(fd, p, n), 0o777); },
    path_remove_directory: (fd, p, n) => {
      const dir = path(fd, p, n);
      if (/(^|\/)\.\/*$/.test(dir)) fail(EINVAL); // (not ".")
      fs.rmdir(dir);
    },
    path_unlink_file: (fd, p, n) => { fs.unlink(path(fd, p, n)); },
    path_rename: (fd, p, n, newfd, p2, n2) => { fs.rename(path(fd, p, n), path(newfd, p2, n2)); },
    path_symlink: (tp, tn, fd, p, n) => {
      const target = decoder.decode(new Uint8Array(mem(), tp, tn));
      fs.symlink(target, path(fd, p, n));
    },
    path_readlink: (fd, p, n, buf, len, usedp) => {
      const target = encoder.encode(fs.readlink(path(fd, p, n)));
      const k = Math.min(target.length, len);
      heap().set(target.subarray(0, k), buf);
      view().setUint32(usedp, k, true);
    },
    path_link: (fd, flags, p, n, newfd, p2, n2) => {
      if (!fs.link) fail(ENOTSUP); // (Emscripten's files have one name.)
      fs.link(path(fd, p, n), path(newfd, p2, n2));
    },
    path_filestat_get: (fd, flags, p, n, statp) => {
      const file = path(fd, p, n);
      // (A path that ends in a slash is of a directory, through a link
      // if need be: not every fs sees the slash.)
      const dir = file.endsWith('/');
      const st = flags & 1 || dir ? fs.stat(file) : lstat(file);
      if (dir && !isDir(st.mode)) fail(ENOTDIR);
      putStat(statp, st);
    },
    path_filestat_set_times: (fd, flags, p, n, atim, mtim, fst) => {
      setTimes(path(fd, p, n), !!(flags & 1), ...times(atim, mtim, fst));
    },
  };
  return Object.fromEntries(Object.entries(calls).map(([name, f]) => [name, call(f)]));
}

// memFS returns files in memory with the interface of Emscripten's FS
// (the part of it wasiFS uses, and writeFile, readFile, mkdirTree and
// chmod), for a page that has no Emscripten. Paths are absolute. There
// are no users and no permissions: every file can be read and written.
//
// fs.persist(dir) loads the directory dir from IndexedDB, and keeps
// what is written there afterwards (a moment after each change, and
// when fs.sync() is called, which returns a promise): the database is
// the one Emscripten's IDBFS has for a file system mounted at dir.
export function memFS() {
  class ErrnoError extends Error {
    constructor(errno) {
      super('file system error ' + errno);
      this.name = 'ErrnoError';
      this.errno = errno;
    }
  }
  const fail = (errno) => { throw new ErrnoError(errno); };
  const now = () => BigInt(Date.now()) * 1000000n;
  let lastIno = 0, lastFd = 2;
  const streams = new Map();

  // A node is a directory (contents: a Map of its entries), a file
  // (contents: its bytes, the first size of them), a link (contents:
  // its target) or a device (dev: how it is read and written).
  const node = (parent, name, mode) => {
    const t = now();
    const n = { id: ++lastIno, parent, name, mode, nlink: 1, atime: t, mtime: t, ctime: t, contents: null, size: 0, dev: null };
    if (isDir(mode)) n.contents = new Map();
    if (isFile(mode)) n.contents = new Uint8Array(0);
    if (parent) {
      parent.contents.set(name, n);
      parent.mtime = parent.ctime = t;
    }
    changed();
    return n;
  };
  let root = null;
  const getPath = (n) => {
    const parts = [];
    for (; n !== root; n = n.parent) parts.unshift(n.name);
    return '/' + parts.join('/');
  };
  // walk finds what path names: the node (null if there is none), and
  // the directory and the name it has or would have there. Links are
  // followed, but for one at the end of the path if follow is false.
  const walk = (path, follow, depth = 0) => {
    if (typeof path !== 'string' || path === '') fail(ENOENT);
    if (depth > 40) fail(ELOOP);
    const parts = path.split('/').filter((p) => p !== '');
    const mustDir = path.endsWith('/') && parts.length > 0;
    let n = root, parent = root, name = '';
    for (let i = 0; i < parts.length; i++) {
      const part = parts[i];
      if (!isDir(n.mode)) fail(ENOTDIR);
      if (part === '.') continue;
      if (part === '..') {
        n = n.parent;
        parent = n.parent;
        name = n.name;
        continue;
      }
      parent = n;
      name = part;
      const last = i === parts.length - 1;
      const next = parent.contents.get(part);
      if (!next) {
        if (last) return { parent, name, node: null };
        fail(ENOENT);
      }
      if (isLink(next.mode) && (!last || follow || mustDir)) {
        if (next.contents === '') fail(ENOENT);
        const rest = parts.slice(i + 1).join('/');
        const from = next.contents.startsWith('/') ? '' : getPath(parent);
        return walk(from + '/' + next.contents + (rest ? '/' + rest : '') + (mustDir ? '/' : ''), follow, depth + 1);
      }
      n = next;
    }
    if (mustDir && !isDir(n.mode)) fail(ENOTDIR);
    return { parent, name, node: n };
  };
  const find = (path, follow = true) => walk(path, follow).node || fail(ENOENT);
  const resize = (n, len) => {
    if (len < 0) fail(EINVAL);
    if (len > n.contents.length) {
      const bytes = new Uint8Array(Math.max(len, Math.min(2 * n.contents.length, n.contents.length + (16 << 20))));
      bytes.set(n.contents.subarray(0, n.size));
      n.contents = bytes;
    } else if (len < n.size) {
      n.contents.fill(0, len, n.size);
    }
    n.size = len;
  };
  const attr = (n) => ({
    dev: 1, ino: n.id, mode: n.mode, nlink: n.nlink, uid: 0, gid: 0, rdev: 0,
    size: isFile(n.mode) ? n.size : isLink(n.mode) ? encoder.encode(n.contents).length : isDir(n.mode) ? 4096 : 0,
    atime: new Date(Number(n.atime / 1000000n)), mtime: new Date(Number(n.mtime / 1000000n)), ctime: new Date(Number(n.ctime / 1000000n)),
    atimeNs: n.atime, mtimeNs: n.mtime, ctimeNs: n.ctime,
    blksize: 4096, blocks: Math.ceil(n.size / 512),
  });
  const stream = (fd) => streams.get(fd) || fail(EBADF);
  const setTimes = (n, a, m) => {
    if (a !== null) n.atime = a;
    if (m !== null) n.mtime = m;
    n.ctime = now();
    changed();
  };

  const fs = {
    ErrnoError,
    isDir, isFile, isLink,
    getPath,
    getStream: (fd) => streams.get(fd),
    open(path, flags = 0, mode = 0o666) {
      const access = flags & O_ACCMODE, excl = (flags & O_CREAT) && (flags & O_EXCL);
      let { parent, name, node: n } = walk(path, !(flags & O_NOFOLLOW) && !excl);
      if (n) {
        if (excl) fail(EEXIST);
        if (isLink(n.mode)) fail(ELOOP);
        if (isDir(n.mode)) {
          if (access !== 0 || (flags & O_TRUNC)) fail(EISDIR);
        } else if (flags & O_DIRECTORY) {
          fail(ENOTDIR);
        }
      } else {
        if (!(flags & O_CREAT)) fail(ENOENT);
        if (path.endsWith('/')) fail(EISDIR);
        n = node(parent, name, S_IFREG | (mode & 0o7777));
      }
      if ((flags & O_TRUNC) && isFile(n.mode) && n.size > 0) {
        resize(n, 0);
        n.mtime = n.ctime = now();
        changed();
      }
      const s = { fd: ++lastFd, node: n, flags, position: 0, path };
      streams.set(s.fd, s);
      return s;
    },
    close(s) {
      streams.delete(s.fd);
    },
    read(s, buffer, offset, length, position) {
      const n = s.node;
      if ((s.flags & O_ACCMODE) === O_WRONLY) fail(EBADF);
      if (isDir(n.mode)) fail(EISDIR);
      if (n.dev) return n.dev.read(buffer, offset, length);
      const at = position ?? s.position;
      const k = Math.max(0, Math.min(length, n.size - at));
      buffer.set(n.contents.subarray(at, at + k), offset);
      if (position === undefined) s.position += k;
      return k;
    },
    write(s, buffer, offset, length, position) {
      const n = s.node;
      if ((s.flags & O_ACCMODE) === 0) fail(EBADF);
      if (isDir(n.mode)) fail(EISDIR);
      if (n.dev) return n.dev.write(buffer, offset, length);
      const at = s.flags & O_APPEND ? n.size : position ?? s.position;
      if (at + length > n.size) resize(n, at + length);
      n.contents.set(buffer.subarray(offset, offset + length), at);
      n.mtime = n.ctime = now();
      if (position === undefined) s.position = at + length;
      changed();
      return length;
    },
    llseek(s, offset, whence) {
      const at = offset + (whence === 1 ? s.position : whence === 2 ? s.node.size : 0);
      if (whence < 0 || whence > 2 || at < 0) fail(EINVAL);
      return (s.position = at);
    },
    stat: (path, dontFollow) => attr(find(path, !dontFollow)),
    lstat: (path) => attr(find(path, false)),
    fstat: (fd) => attr(stream(fd).node),
    readdir(path) {
      const n = find(path);
      if (!isDir(n.mode)) fail(ENOTDIR);
      return ['.', '..', ...n.contents.keys()];
    },
    mkdir(path, mode = 0o777) {
      const { parent, name, node: n } = walk(path, false);
      if (n) fail(EEXIST);
      return node(parent, name, S_IFDIR | (mode & 0o7777));
    },
    mkdirTree(path, mode = 0o777) {
      let dir = '';
      for (const part of path.split('/').filter((p) => p !== '')) {
        dir += '/' + part;
        if (!walk(dir, true).node) fs.mkdir(dir, mode);
      }
    },
    rmdir(path) {
      const { parent, name, node: n } = walk(path, false);
      if (!n) fail(ENOENT);
      if (!isDir(n.mode)) fail(ENOTDIR);
      if (n === root) fail(EBUSY);
      if (n.contents.size > 0) fail(ENOTEMPTY);
      parent.contents.delete(name);
      parent.mtime = parent.ctime = now();
      n.nlink = 0;
      changed();
    },
    unlink(path) {
      const { parent, name, node: n } = walk(path, false);
      if (!n) fail(ENOENT);
      if (isDir(n.mode)) fail(EISDIR);
      parent.contents.delete(name);
      parent.mtime = parent.ctime = now();
      n.nlink--;
      changed();
    },
    rename(from, to) {
      const a = walk(from, false), b = walk(to, false);
      if (!a.node) fail(ENOENT);
      if (a.node === b.node) return;
      if (isDir(a.node.mode)) {
        // (Not into itself.)
        for (let d = b.parent; ; d = d.parent) {
          if (d === a.node) fail(EINVAL);
          if (d === root) break;
        }
        if (b.node) {
          if (!isDir(b.node.mode)) fail(ENOTDIR);
          if (b.node.contents.size > 0) fail(ENOTEMPTY);
        }
      } else {
        if (b.node && isDir(b.node.mode)) fail(EISDIR);
        if (to.endsWith('/')) fail(ENOTDIR);
      }
      if (a.node === root) fail(EBUSY);
      if (b.node) b.node.nlink--;
      a.parent.contents.delete(a.name);
      b.parent.contents.set(b.name, a.node);
      a.node.parent = b.parent;
      a.node.name = b.name;
      a.parent.mtime = a.parent.ctime = b.parent.mtime = b.parent.ctime = a.node.ctime = now();
      changed();
    },
    symlink(target, path) {
      const { parent, name, node: n } = walk(path, false);
      if (n) fail(EEXIST);
      if (path.endsWith('/')) fail(ENOENT);
      node(parent, name, S_IFLNK | 0o777).contents = target;
    },
    readlink(path) {
      const n = find(path, false);
      if (!isLink(n.mode)) fail(EINVAL);
      return n.contents;
    },
    link(from, to) {
      const n = find(from, false), b = walk(to, false);
      if (isDir(n.mode)) fail(EPERM);
      if (b.node) fail(EEXIST);
      if (to.endsWith('/')) fail(ENOENT);
      b.parent.contents.set(b.name, n);
      b.parent.mtime = b.parent.ctime = n.ctime = now();
      n.nlink++;
      changed();
    },
    truncate(path, len) {
      const n = find(path);
      if (isDir(n.mode)) fail(EISDIR);
      if (!isFile(n.mode)) fail(EINVAL);
      resize(n, len);
      n.mtime = n.ctime = now();
      changed();
    },
    ftruncate(fd, len) {
      const s = stream(fd);
      if ((s.flags & O_ACCMODE) === 0 || !isFile(s.node.mode)) fail(EINVAL);
      resize(s.node, len);
      s.node.mtime = s.node.ctime = now();
      changed();
    },
    // (Times in milliseconds, as Emscripten's.)
    utime(path, atime, mtime) {
      const ns = (ms) => BigInt(Math.round(+ms)) * 1000000n;
      setTimes(find(path), ns(atime), ns(mtime));
    },
    utimens: (path, a, m, follow = true) => setTimes(find(path, follow), a, m),
    futimens: (fd, a, m) => setTimes(stream(fd).node, a, m),
    chmod(path, mode, dontFollow) {
      const n = find(path, !dontFollow);
      n.mode = (n.mode & S_IFMT) | (mode & 0o7777);
      n.ctime = now();
      changed();
    },
    writeFile(path, data) {
      const s = fs.open(path, O_WRONLY | O_CREAT | O_TRUNC);
      const bytes = typeof data === 'string' ? encoder.encode(data) : data;
      fs.write(s, bytes, 0, bytes.length);
      fs.close(s);
    },
    readFile(path, opts = {}) {
      const n = find(path);
      if (!isFile(n.mode)) fail(EISDIR);
      const bytes = n.contents.slice(0, n.size);
      return opts.encoding === 'utf8' ? decoder.decode(bytes) : bytes;
    },
    // device makes a file that is read and written by the functions of
    // dev: read(buffer, offset, length) and write(buffer, offset,
    // length), which return how many bytes.
    device(path, dev) {
      const { parent, name, node: n } = walk(path, false);
      if (n) fail(EEXIST);
      node(parent, name, S_IFCHR | 0o666).dev = dev;
    },
    persist, sync,
  };

  // The directories kept in IndexedDB: for each its database, as
  // Emscripten's IDBFS makes it (libidbfs.js): one named as the
  // directory, with a store FILE_DATA of records { timestamp, mode,
  // contents } by path, for the directories and the files below it
  // (links are not kept), and an index of them by timestamp.
  const kept = new Map();
  let timer = 0, syncing = Promise.resolve();
  function changed() {
    if (kept.size > 0 && !timer) timer = setTimeout(sync, 100);
  }
  const request = (req) => new Promise((done, failed) => {
    req.onsuccess = () => done(req.result);
    req.onerror = () => failed(req.error);
  });
  async function persist(dir) {
    dir = '/' + dir.split('/').filter((p) => p !== '').join('/');
    if (kept.has(dir)) return;
    const req = indexedDB.open(dir, 21);
    req.onupgradeneeded = () => {
      const db = req.result;
      const store = db.objectStoreNames.contains('FILE_DATA') ? req.transaction.objectStore('FILE_DATA') : db.createObjectStore('FILE_DATA');
      if (!store.indexNames.contains('timestamp')) store.createIndex('timestamp', 'timestamp', { unique: false });
    };
    const db = await request(req);
    const store = db.transaction(['FILE_DATA'], 'readonly').objectStore('FILE_DATA');
    const [paths, entries] = await Promise.all([request(store.getAllKeys()), request(store.getAll())]);
    fs.mkdirTree(dir);
    // (A directory comes before what is in it.)
    for (const i of paths.map((_, i) => i).sort((a, b) => (paths[a] < paths[b] ? -1 : 1))) {
      const path = paths[i], entry = entries[i];
      if (isDir(entry.mode)) {
        fs.mkdirTree(path);
      } else if (isFile(entry.mode)) {
        fs.mkdirTree(path.slice(0, path.lastIndexOf('/')));
        fs.writeFile(path, entry.contents);
      } else {
        continue;
      }
      fs.chmod(path, entry.mode);
      fs.utime(path, entry.timestamp, entry.timestamp);
    }
    kept.set(dir, db);
  }
  // sync writes what changed in the directories kept: the records whose
  // timestamp is not the file's, and no record for what is gone.
  function sync() {
    clearTimeout(timer);
    timer = 0;
    return (syncing = syncing.catch(() => {}).then(async () => {
      for (const [dir, db] of kept) {
        const local = new Map();
        const list = (path, n) => {
          for (const [name, child] of n.contents) {
            if (!isDir(child.mode) && !isFile(child.mode)) continue;
            local.set(path + '/' + name, child);
            if (isDir(child.mode)) list(path + '/' + name, child);
          }
        };
        const top = walk(dir, true).node;
        if (top && isDir(top.mode)) list(dir, top);
        const tx = db.transaction(['FILE_DATA'], 'readwrite'), store = tx.objectStore('FILE_DATA');
        const finished = new Promise((done, failed) => {
          tx.oncomplete = done;
          tx.onerror = tx.onabort = () => failed(tx.error);
        });
        const index = store.index('timestamp');
        const [paths, stamps] = await Promise.all([request(index.getAllKeys()), request(index.getAll())]);
        const remote = new Map(paths.map((path, i) => [path, stamps[i].timestamp]));
        for (const [path, n] of local) {
          const timestamp = new Date(Number(n.mtime / 1000000n));
          if (remote.get(path)?.getTime() === timestamp.getTime()) continue;
          const entry = { timestamp, mode: n.mode };
          if (isFile(n.mode)) entry.contents = n.contents.slice(0, n.size);
          store.put(entry, path);
        }
        for (const path of remote.keys()) {
          if (!local.has(path)) store.delete(path);
        }
        await finished;
      }
    }));
  }

  root = node(null, '', S_IFDIR | 0o777);
  root.parent = root;
  // As an Emscripten program finds it.
  fs.mkdir('/tmp');
  fs.mkdirTree('/home/web_user');
  fs.mkdir('/dev');
  fs.device('/dev/null', { read: () => 0, write: (buffer, offset, length) => length });
  fs.device('/dev/urandom', {
    read: (buffer, offset, length) => {
      for (let i = 0; i < length; i += 65536) crypto.getRandomValues(buffer.subarray(offset + i, offset + Math.min(length, i + 65536)));
      return length;
    },
    write: (buffer, offset, length) => length,
  });
  return fs;
}

// The names of the errors, by their numbers (WASI's, which are
// Emscripten's), as node gives them in an error's code.
const errnoNames = ('SUCCESS 2BIG ACCES ADDRINUSE ADDRNOTAVAIL AFNOSUPPORT AGAIN ALREADY BADF BADMSG BUSY CANCELED CHILD ' +
  'CONNABORTED CONNREFUSED CONNRESET DEADLK DESTADDRREQ DOM DQUOT EXIST FAULT FBIG HOSTUNREACH IDRM ILSEQ INPROGRESS INTR ' +
  'INVAL IO ISCONN ISDIR LOOP MFILE MLINK MSGSIZE MULTIHOP NAMETOOLONG NETDOWN NETRESET NETUNREACH NFILE NOBUFS NODEV NOENT ' +
  'NOEXEC NOLCK NOLINK NOMEM NOMSG NOPROTOOPT NOSPC NOSYS NOTCONN NOTDIR NOTEMPTY NOTRECOVERABLE NOTSOCK NOTSUP NOTTY NXIO ' +
  'OVERFLOW OWNERDEAD PERM PIPE PROTO PROTONOSUPPORT PROTOTYPE RANGE ROFS SPIPE SRCH STALE TIMEDOUT TXTBSY XDEV NOTCAPABLE').split(' ').map((n) => 'E' + n);

// withKind gives what stat returns the methods of node's Stats.
export function withKind(st) {
  const t = st.mode & S_IFMT;
  st.isDirectory = () => t === S_IFDIR;
  st.isFile = () => t === S_IFREG;
  st.isSymbolicLink = () => t === S_IFLNK;
  st.isCharacterDevice = () => t === S_IFCHR;
  st.isBlockDevice = () => t === 0o060000;
  st.isFIFO = () => t === 0o010000;
  st.isSocket = () => t === 0o140000;
  return st;
}

// resolvePath is node's path.resolve(...paths), from the directory cwd.
export function resolvePath(cwd, ...paths) {
  let p = cwd;
  for (const q of paths) p = q.startsWith('/') ? q : p + '/' + q;
  const parts = [];
  for (const part of p.split('/')) {
    if (part === '' || part === '.') continue;
    if (part === '..') parts.pop();
    else parts.push(part);
  }
  return '/' + parts.join('/');
}

// nodeFS returns the globals fs, process and path that a GOOS=js program
// uses for files (src/syscall/fs_js.go: node's fs, with callbacks), made with
// fs (the interface of Emscripten's FS, as wasiFS takes), so that a
// program sees the same files whether it is built for GOOS=js or
// wasip1. Descriptors 0 to 2 are not fs's: standard input is empty,
// and what is written to 1 and 2 goes to write(fd, bytes). The callbacks
// are called before the functions return: a goroutine waits for none.
// The working directory is cwd (process.chdir changes it).
export function nodeFS(fs, write, cwd = '/') {
  const S_IFCHR_RW = S_IFCHR | 0o666;
  const error = (e) => {
    if (e && typeof e.errno === 'number') {
      const code = errnoNames[e.errno] || 'EIO';
      return Object.assign(new Error(code), { code });
    }
    return e;
  };
  // (A function of fs, with node's callback last.)
  const call = (f) => (...args) => {
    const cb = args.pop();
    let r;
    try {
      r = f(...args);
    } catch (e) {
      cb(error(e));
      return;
    }
    cb(null, r);
  };
  const abs = (path) => {
    const p = path.startsWith('/') ? path : cwd + (cwd.endsWith('/') ? '' : '/') + path;
    // (As node: . and .. are resolved, but for a path that ends in a
    // slash, which says it is a directory.)
    const parts = [];
    for (const part of p.split('/')) {
      if (part === '' || part === '.') continue;
      if (part === '..') parts.pop();
      else parts.push(part);
    }
    return '/' + parts.join('/') + (p.endsWith('/') && parts.length > 0 ? '/' : '');
  };
  const stats = (a) => withKind({
    dev: a.dev, ino: a.ino, mode: a.mode, nlink: a.nlink, uid: a.uid, gid: a.gid, rdev: a.rdev, size: a.size,
    blksize: a.blksize, blocks: a.blocks,
    atimeMs: a.atimeNs !== undefined ? Number(a.atimeNs) / 1e6 : a.atime.getTime(),
    mtimeMs: a.mtimeNs !== undefined ? Number(a.mtimeNs) / 1e6 : a.mtime.getTime(),
    ctimeMs: a.ctimeNs !== undefined ? Number(a.ctimeNs) / 1e6 : a.ctime.getTime(),
  });
  const std = (fd) => fd >= 0 && fd <= 2;
  const stream = (fd) => fs.getStream(fd) || fail(EBADF);
  const fail = (errno) => { throw { errno }; };
  const ms = (s) => (s instanceof Date ? s.getTime() : s * 1000);
  const nodeFs = {
    constants: { O_WRONLY, O_RDWR, O_CREAT, O_TRUNC, O_APPEND, O_EXCL, O_DIRECTORY },
    open: call((path, flags, mode) => fs.open(abs(path), flags, mode).fd),
    close: call((fd) => { if (!std(fd)) fs.close(stream(fd)); }),
    read: call((fd, buffer, offset, length, position) => {
      if (fd === 0) return 0;
      if (std(fd)) fail(EBADF);
      return fs.read(stream(fd), buffer, offset, length, position ?? undefined);
    }),
    write: call((fd, buffer, offset, length, position) => {
      if (fd === 1 || fd === 2) {
        write(fd, buffer.slice(offset, offset + length));
        return length;
      }
      if (std(fd)) fail(EBADF);
      return fs.write(stream(fd), buffer, offset, length, position ?? undefined);
    }),
    // (wasm_exec.js's page has it: a program may write with it.)
    writeSync(fd, buffer) {
      if (fd === 1 || fd === 2) {
        write(fd, buffer.slice());
        return buffer.length;
      }
      return fs.write(stream(fd), buffer, 0, buffer.length);
    },
    fstat: call((fd) => {
      if (std(fd)) {
        const t = new Date();
        return stats({ dev: 0, ino: 0, mode: S_IFCHR_RW, nlink: 1, uid: 0, gid: 0, rdev: 0, size: 0, blksize: 4096, blocks: 0, atime: t, mtime: t, ctime: t });
      }
      return stats(fs.fstat(fd));
    }),
    stat: call((path) => stats(fs.stat(abs(path)))),
    lstat: call((path) => stats(fs.lstat(abs(path)))),
    readdir: call((path) => fs.readdir(abs(path)).filter((n) => n !== '.' && n !== '..')),
    mkdir: call((path, perm) => { fs.mkdir(abs(path), perm); }),
    rmdir: call((path) => { fs.rmdir(abs(path)); }),
    unlink: call((path) => { fs.unlink(abs(path)); }),
    rename: call((from, to) => { fs.rename(abs(from), abs(to)); }),
    link: call((from, to) => { if (!fs.link) fail(ENOTSUP); fs.link(abs(from), abs(to)); }),
    symlink: call((target, path) => { fs.symlink(target, abs(path)); }),
    readlink: call((path) => fs.readlink(abs(path))),
    chmod: call((path, mode) => { fs.chmod(abs(path), mode); }),
    fchmod: call((fd, mode) => { fs.fchmod ? fs.fchmod(fd, mode) : fs.chmod(stream(fd).path, mode); }),
    // (No users.)
    chown: call((path) => { fs.stat(abs(path)); }),
    lchown: call((path) => { fs.lstat(abs(path)); }),
    fchown: call((fd) => { stream(fd); }),
    utimes: call((path, atime, mtime) => { fs.utime(abs(path), ms(atime), ms(mtime)); }),
    truncate: call((path, len) => { fs.truncate(abs(path), len); }),
    ftruncate: call((fd, len) => { fs.ftruncate(fd, len); }),
    fsync: call((fd) => { if (!std(fd)) stream(fd); }),
  };
  const process = {
    pid: -1, ppid: -1,
    getuid: () => -1, getgid: () => -1, geteuid: () => -1, getegid: () => -1, getgroups: () => [],
    umask: () => 0o22,
    cwd: () => cwd,
    chdir(path) {
      const dir = abs(path).replace(/(.)\/$/, '$1');
      let a;
      try {
        a = fs.stat(dir);
      } catch (e) {
        throw error(e);
      }
      if (!isDir(a.mode)) throw error({ errno: ENOTDIR });
      cwd = dir;
    },
  };
  return { fs: nodeFs, process, path: { resolve: (...paths) => resolvePath(cwd, ...paths) } };
}
