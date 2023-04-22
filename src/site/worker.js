importScripts('wasm_exec.js');

const go = new Go();

(async () => {
  const response = await fetch('cs-story-forge.wasm');
  const wasmArrayBuffer = await response.arrayBuffer();
  const result = await WebAssembly.instantiate(wasmArrayBuffer, go.importObject);
  go.run(result.instance);

  self.addEventListener('message', (e) => {
    const { demoByteArray } = e.data;
    const story = Go_GenerateProompt(demoByteArray);
    self.postMessage({ story });
  });
})();
