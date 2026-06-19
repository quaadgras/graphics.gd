// Minimal static file server that sets the cross-origin-isolation headers a
// Godot web export expects. Mirrors the headers gd's own dev server sends.
// Usage: node serve.mjs <dir> [port]
import { createServer } from "node:http";
import { readFile, stat } from "node:fs/promises";
import { join, normalize, extname } from "node:path";

const dir = process.argv[2] || ".";
const port = parseInt(process.argv[3] || "8080", 10);
const types = {
  ".html": "text/html", ".js": "text/javascript", ".mjs": "text/javascript",
  ".wasm": "application/wasm", ".pck": "application/octet-stream",
  ".png": "image/png", ".json": "application/json", ".css": "text/css",
};

createServer(async (req, res) => {
  res.setHeader("Cross-Origin-Embedder-Policy", "require-corp");
  res.setHeader("Cross-Origin-Opener-Policy", "same-origin");
  res.setHeader("Cross-Origin-Resource-Policy", "cross-origin");
  let p = decodeURIComponent((req.url || "/").split("?")[0]);
  if (p === "/" || p.endsWith("/")) p += "index.html";
  const file = join(dir, normalize(p).replace(/^(\.\.[/\\])+/, ""));
  try {
    if (!(await stat(file)).isFile()) throw new Error("not a file");
    res.setHeader("Content-Type", types[extname(file)] || "application/octet-stream");
    res.end(await readFile(file));
  } catch {
    res.statusCode = 404;
    res.end("not found");
  }
}).listen(port, () => console.log(`serving ${dir} on http://0.0.0.0:${port}`));
