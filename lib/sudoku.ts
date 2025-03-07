// @ts-types="npm:@types/golang-wasm-exec"
import "./vendor/wasm_exec/wasm_exec.js"

const go = new Go()
const result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
go.run(result.instance)

// @ts-ignore _
export const NewPuzzle = (globalThis.NewPuzzle) as (targetClueCount: number, maxTries: number) => number[]
