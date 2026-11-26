// patch.mjs dir: change the copies in dir of compiler.gd's loader
// (lib/wasm/wasm_gc*.js) and of the engine's (godot.js, from graphics.gd's
// web template) so that the page can run a program built for GOWASM=gc
// beside the engine. Each change replaces text that must be there exactly
// once: a loader that has moved on fails the build here, not in the page.
import fs from 'node:fs';

const dir = process.argv[2];
function patch(name, changes) {
  const path = `${dir}/${name}`;
  let text = fs.readFileSync(path, 'utf8');
  for (const [from, to] of changes) {
    if (text.split(from).length !== 2) throw new Error(`patch.mjs: ${name}: expected once: ${from}`);
    text = text.replace(from, () => to);
  }
  fs.writeFileSync(path, text);
}

// compiler.gd's loader makes the import object itself: the page adds the
// engine's functions to it (graphics.gd imports them from module "gd",
// which is the runtime's own too), and sees the instance when it is made.
// A compiler.gd that has options.imports and options.started needs no
// change; one from before them is given them.
if (!fs.readFileSync(`${dir}/wasm_gc_browser.js`, 'utf8').includes('options.started')) {
  patch('wasm_gc.js', [[
    '    wasi_snapshot_preview1: host.wasi.imports,\n  };\n',
    '    wasi_snapshot_preview1: host.wasi.imports,\n  };\n  for (const [m, fns] of Object.entries(host.imports || {})) imports[m] = Object.assign(imports[m] || {}, fns);\n',
  ]]);
  patch('wasm_gc_browser.js', [
    ['    global,\n  });', '    global,\n    imports: options.imports,\n  });'],
    ['start: (inst) => { instance = inst;', 'start: (inst) => { instance = inst; if (options.started) options.started(inst);'],
  ]);
}

// The engine's loader starts a Go program with wasm_exec.js (new Go()),
// from library.wasm. With window.gdGo, the page starts it instead (see
// run.html); with window.gdEngineModule, the engine is that module, which
// the page compiled once, not a download for every run.
patch('godot.js', [
  ['const go = new Go();', 'const go = window.gdGo ? window.gdGo() : new Go();'],
  ['WebAssembly.compileStreaming(fetch("library.wasm"))', '(go.library ? go.library() : WebAssembly.compileStreaming(fetch("library.wasm")))'],
  ['WebAssembly.instantiate(goModule, go.importObject)', '(go.instantiate ? go.instantiate(goModule) : WebAssembly.instantiate(goModule, go.importObject))'],
  ['var resume = exports.resume;', 'var resume = exports.resume || function () {};'],
  [
    "if (typeof (WebAssembly.instantiateStreaming) !== 'undefined') {",
    "if (window.gdEngineModule) {\n\t\t\t\t\tWebAssembly.instantiate(window.gdEngineModule, imports).then(function (instance) { onSuccess(instance, window.gdEngineModule); });\n\t\t\t\t} else if (typeof (WebAssembly.instantiateStreaming) !== 'undefined') {",
  ],
  [
    'loadPromise = preloader.loadPromise(`${loadPath}.wasm`, size, true);',
    'loadPromise = window.gdEngineModule ? Promise.resolve(new Response(new Uint8Array(0))) : preloader.loadPromise(`${loadPath}.wasm`, size, true);',
  ],
]);
