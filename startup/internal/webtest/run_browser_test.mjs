// Headless-browser test runner for graphics.gd web (js/wasm) builds.
//
// Loads a served Godot web export in headless Chromium, streams the page's
// console to stdout, and exits non-zero if the embedded `go test` binary fails.
//
// Go stdout is routed to console.log line-by-line by wasm_exec.js; a non-zero
// os.Exit() surfaces as console.warn("exit code:", N). We treat a "PASS"/"ok"
// line as success and "FAIL"/"--- FAIL"/"exit code:"/panic/exception as failure.
//
// Usage: node run_browser_test.mjs <url> [timeoutMs]
// Requires: a `chromium` (or CHROMIUM env) binary and Node >= 22 (global WebSocket/fetch).

import { spawn, spawnSync } from "node:child_process";
import { mkdtempSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";

const url = process.argv[2] || "http://localhost:8080";
const timeoutMs = parseInt(process.argv[3] || "120000", 10);

// Resolve a Chrome/Chromium binary: explicit CHROMIUM wins, otherwise probe the
// usual names (GitHub ubuntu-latest ships google-chrome; Debian ships chromium).
function resolveBrowser() {
  const candidates = [process.env.CHROMIUM, "google-chrome", "google-chrome-stable",
    "chromium", "chromium-browser"].filter(Boolean);
  for (const c of candidates) {
    if (c.includes("/")) return c; // explicit path
    if (spawnSync("which", [c]).status === 0) return c;
  }
  return candidates[0] || "chromium";
}
const chromium = resolveBrowser();
const port = 9222 + Math.floor((process.pid % 1000));
const profile = mkdtempSync(join(tmpdir(), "gdwebtest-"));

const flags = [
  "--headless=new",
  "--no-sandbox",
  "--disable-dev-shm-usage",
  // Godot web needs WebGL2; in headless CI there is no real GPU, so use the
  // software ANGLE/SwiftShader backend (and allow it when flagged unsafe).
  "--use-gl=angle",
  "--use-angle=swiftshader",
  "--enable-unsafe-swiftshader",
  "--mute-audio",
  // Chromium >=111 rejects the devtools WebSocket upgrade from non-browser
  // clients unless origins are explicitly allowed.
  "--remote-allow-origins=*",
  `--remote-debugging-port=${port}`,
  `--user-data-dir=${profile}`,
  url,
];

let proc, ws, done = false, timer;

function finish(code, reason) {
  if (done) return;
  done = true;
  console.log(`\n[runner] ${reason} -> exit ${code}`);
  clearTimeout(timer);
  try { ws && ws.close(); } catch {}
  try { proc && proc.kill("SIGKILL"); } catch {}
  try { rmSync(profile, { recursive: true, force: true }); } catch {}
  process.exit(code);
}

function classify(text) {
  for (const line of text.split("\n")) {
    const s = line.trim();
    if (s === "PASS" || /^ok\s/.test(s)) return { code: 0, reason: `test reported: ${s}` };
    if (s === "FAIL" || /^FAIL\s/.test(s) || s.startsWith("--- FAIL")) return { code: 1, reason: `test reported: ${s}` };
    if (s.startsWith("exit code:")) return { code: 1, reason: `non-zero ${s}` };
    if (/^panic:/.test(s) || /fatal error:/.test(s)) return { code: 1, reason: `runtime: ${s}` };
  }
  return null;
}

async function getJson(path) {
  const r = await fetch(`http://127.0.0.1:${port}${path}`);
  return r.json();
}

async function main() {
  proc = spawn(chromium, flags, { stdio: ["ignore", "inherit", "inherit"] });
  proc.on("exit", (c) => finish(c === 0 ? 0 : 70, `chromium exited (${c})`));

  // Chromium was launched pointing at `url`; wait for its devtools endpoint and
  // grab the existing page target (GET /json works on all versions).
  let target;
  for (let i = 0; i < 150 && !target; i++) {
    try {
      const list = await getJson(`/json/list`);
      target = (list || []).find((t) => t.type === "page" && t.webSocketDebuggerUrl);
    } catch { /* devtools not up yet */ }
    if (!target) await new Promise((r) => setTimeout(r, 200));
  }
  if (!target) finish(71, "could not find devtools page target");

  ws = new WebSocket(target.webSocketDebuggerUrl);
  let id = 0;
  const send = (method, params = {}) => ws.send(JSON.stringify({ id: ++id, method, params }));

  ws.addEventListener("open", () => {
    send("Runtime.enable");
    send("Log.enable");
    send("Page.enable");
    timer = setTimeout(() => finish(2, `timeout after ${timeoutMs}ms (no PASS/FAIL seen)`), timeoutMs);
  });

  ws.addEventListener("message", (ev) => {
    let msg;
    try { msg = JSON.parse(ev.data); } catch { return; }
    let text = null;
    if (msg.method === "Runtime.consoleAPICalled") {
      text = (msg.params.args || []).map((a) => a.value ?? a.description ?? "").join(" ");
      console.log(`[console.${msg.params.type}] ${text}`);
    } else if (msg.method === "Log.entryAdded") {
      text = msg.params.entry.text || "";
      console.log(`[log.${msg.params.entry.level}] ${text}`);
    } else if (msg.method === "Runtime.exceptionThrown") {
      const d = msg.params.exceptionDetails;
      text = (d.exception && (d.exception.description || d.exception.value)) || d.text || "exception";
      console.log(`[exception] ${text}`);
      return finish(1, `uncaught exception: ${String(text).split("\n")[0]}`);
    }
    if (text) {
      const verdict = classify(text);
      if (verdict) finish(verdict.code, verdict.reason);
    }
  });
  ws.addEventListener("error", (e) => finish(72, `websocket error: ${e?.message || e}`));
}

main().catch((e) => finish(73, `runner error: ${e?.stack || e}`));
