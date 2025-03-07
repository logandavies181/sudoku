import { render } from "https://esm.sh/preact@10.25.3"

import { html } from "./html.ts"
import { Board } from "./lib/board.ts";

// if ("serviceWorker" in navigator) {
//   navigator.serviceWorker.register("sw.js", { scope: "/harmonies-planner/" })
// }

function App() {
  return html`
    <main class="flex justify-center min-w-full min-h-full">
      <${Board} />
    </main>
  `
}

render(html`<${App} />`, document.body)
