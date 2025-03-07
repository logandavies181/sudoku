import { html } from "../html.ts"

export function Board() {
  return html`
    <div class="flex grow flex-wrap bg-white min-w-full" >
      <${Box} />
      <${Box} />
      <${Box} />
      <${Box} />
      <${Box} />
      <${Box} />
      <${Box} />
      <${Box} />
      <${Box} />
    </div>
  `
}

function Box() {
  return html`
    <div class="border-2 border-solid flex grow flex-wrap bg-white max-w-1/3 min-w-1/3" >
      <${Cell} />
      <${Cell} />
      <${Cell} />
      <${Cell} />
      <${Cell} />
      <${Cell} />
      <${Cell} />
      <${Cell} />
      <${Cell} />
    </div>
  `
}

function Cell() {
  return html`
    <div class="text-md border-1 border-solid flex text-center justify-center items-center aspect-square min-w-1/3" >
      ${Math.random() * 10 % 10 | 1}
    </div>
  `
}
