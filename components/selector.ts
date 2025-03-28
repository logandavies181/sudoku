import { html } from "../html.ts"
import { globalState } from "../lib/state.ts"

export function Selector() {
  return html`
    <div class="flex flex-wrap justify-around min-w-full">
      <${SelectorBox}
        index="1"
        nomargin
      />
      <${SelectorBox}
        index="2"
        nomargin
      />
      <${SelectorBox}
        index="3"
        nomargin
      />
      <${SelectorBox}
        index="4"
        nomargin
      />
      <${SelectorBox}
        index="5"
        nomargin
      />
      <${SelectorBox}
        index="6"
        nomargin
      />
      <${SelectorBox} index="7" />
      <${SelectorBox} index="8" />
      <${SelectorBox} index="9" />
      ${OtherBoxFactory("e", () => {
        const newState = globalState.state
        newState.activeNum = 0
        globalState.publish(newState)
      })}
      ${OtherBoxFactory("p", () => {})} ${OtherBoxFactory("d", () => {})}
    </div>
  `
}

type SelectorBoxProps = {
  index: number
  nomargin: boolean
}

function boxClass(nomargin: boolean) {
  return `rounded-md m-2 aspect-square flex text-center justify-center items-center bg-white text-md min-w-1/9 ${!nomargin ? "mb-10" : ""}`
}

function SelectorBox(props: SelectorBoxProps) {
  const onClick = () => {
    const newState = globalState.state
    newState.activeNum = props.index | 0
    globalState.publish(newState)
  }
  return html`
    <div
      class=${boxClass(props.nomargin)}
      onClick=${onClick}
    >
      ${props.index}
    </div>
  `
}

function OtherBoxFactory(content: any, onClick: () => void) {
  return html`
    <div
      class=${boxClass(false)}
      onClick=${onClick}
    >
      ${content}
    </div>
  `
}
