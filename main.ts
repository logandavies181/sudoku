import { render } from "https://esm.sh/preact@10.25.3"

import { html } from "./html.ts"

import { NewPuzzle } from "./lib/sudoku.ts"

// if ("serviceWorker" in navigator) {
//   navigator.serviceWorker.register("sw.js", { scope: "/harmonies-planner/" })
// }

function App() {
  return html`
    <main class="flex items-center portrait:flex-col landscape:flex-row portrait:min-w-screen landscape:min-w-[50%]">
      Hello world ${NewPuzzle(40, 10_000)}
    </main>
  `
}

render(html`<${App} />`, document.body)
