const wasm = require("./wasm_exec.js");
const fs = require("fs");
const go = new Go();

go.argv = process.argv.slice(2);
go.env = Object.assign({ TMPDIR: require("os").tmpdir() }, process.env);
go.exit = process.exit;

//run wasm
WebAssembly.instantiate(fs.readFileSync("./nodejs/lib.wasm"), go.importObject)
  .then((result) => {
    process.on("exit", (code) => {
      // Node.js exits if no event handler is pending
      if (code === 0 && !go.exited) {
        // deadlock, make Go print error and stack traces
        go._pendingEvent = { id: 0 };
        go._resume();
      }
    });
    return go.run(result.instance);
  })
  .catch((err) => {
    console.error(err);
    process.exit(1);
  });
