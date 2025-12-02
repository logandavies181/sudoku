import { render } from "preact"

import { html } from "./html.ts"
import { Board } from "./components/board.ts";
import { Navbar } from "./components/navbar.ts";
import { Selector } from "./components/selector.ts";

if ("serviceWorker" in navigator) {
  navigator.serviceWorker.register("sw.js", { scope: "/" })
}

function App() {
  return html`
    <div class="touch-manipulation flex grow flex-col min-w-full min-h-full">
      <${Navbar} />
      <main class="flex grow flex-col justify-between min-w-full">
        <${Board} />
        <${Selector} />
      </main>
    </div>
  `
}

render(html`<${App} />`, document.body)
