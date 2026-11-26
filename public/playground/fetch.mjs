// fetch.mjs: the playground's large files, for the worker (worker.mjs) and
// the page (index.html). Their names have a hash of what is in them, and
// they are kept in Cache Storage: a browser's HTTP cache may not keep a
// response of tens of megabytes (Chrome keeps an entry of at most an
// eighth of its cache, a private window very little), however long its
// headers allow.

const store = (async () => {
  try {
    return await caches.open('graphics-gd-playground');
  } catch {
    return null; // (No Cache Storage: not a secure context.)
  }
})();

// keepOnly deletes what the store has but names.
export async function keepOnly(names) {
  const s = await store;
  if (!s) return;
  const want = new Set(names.map((n) => new URL(n, location.href).href));
  for (const req of await s.keys()) if (!want.has(req.url)) await s.delete(req);
}

// fetchBody returns the stream of the body of url, from the store or the
// network (keeping it), reporting progress(label, done, total).
export async function fetchBody(url, what, progress) {
  const s = await store;
  let res = s && (await s.match(url));
  const cached = !!res;
  if (!res) {
    res = await fetch(url);
    if (!res.ok) throw new Error(`${url}: ${res.status} ${res.statusText}`);
    if (s) s.put(url, res.clone()).catch(() => {}); // (Full or not allowed: it is downloaded again next time.)
  }
  const label = (cached ? 'loading ' : 'downloading ') + what + (cached ? ' (cached)' : '');
  const total = Number(res.headers.get('content-length')) || 0;
  let done = 0;
  return res.body.pipeThrough(new TransformStream({
    transform(chunk, c) {
      done += chunk.length;
      if (progress) progress(label, done, total);
      c.enqueue(chunk);
    },
  }));
}

export const gunzip = (stream) => stream.pipeThrough(new DecompressionStream('gzip'));
export const bytes = async (stream) => new Uint8Array(await new Response(stream).arrayBuffer());

// loadManifest returns manifest.json, which says where the large files are.
// (build.sh writes it last: a site with none is still being built.)
export async function loadManifest() {
  const res = await fetch('manifest.json', { cache: 'no-cache' });
  if (!res.ok) throw new Error(`manifest.json: ${res.status} (the playground is not built, or is being rebuilt)`);
  return res.json();
}
