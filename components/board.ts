import { html } from "../html.ts"
import { cell } from "../lib/cell.ts"
import { handleUpdate, newGame } from "../lib/game.ts"
import { puzzleState } from "../lib/game.ts"

import { useEffect, useState } from "https://esm.sh/preact@10.25.3/hooks"

function xPos(src: number): number {
  return src % 9
}

function yPos(src: number): number {
  return Math.floor(src / 9)
}

function boxIndex(src: number): number {
  const x = xPos(src)
  const y = yPos(src)

  return Math.floor(x / 3) + Math.floor(y / 3) * 3
}

export function Board() {
  const boxes = new Array<Array<number>>(9)
  for (let i = 0; i < 81; i++) {
    const bi = boxIndex(i)

    if (boxes[bi] == undefined) {
      boxes[bi] = [i]
      continue
    }
    boxes[bi].push(i)
  }
  useEffect(() => {
    newGame()
  }, [])

  return html`
    <div class="flex flex-wrap bg-white min-w-full">
      ${boxes.map((cells) => {
        return html`<${Box} cellIndexes=${cells} />`
      })}
    </div>
  `
}

type BoxProps = {
  cellIndexes: number[]
}

function Box(props: BoxProps) {
  return html`
    <div class="border-2 border-solid flex grow flex-wrap bg-white max-w-1/3 min-w-1/3">
      ${props.cellIndexes.map((i) => {
        return html`<${Cell} index=${i} />`
      })}
    </div>
  `
}

type CellProps = {
  index: number
}

function Cell(props: CellProps) {
  const [cell, setCell] = useState<cell>({ value: 0, initial: false })

  puzzleState.subscribe(props.index, (c) => {
    setCell(c)
  })

  const onClick = () => {
    handleUpdate(props.index)
  }

  let cellClass = "text-md border-1 border-solid flex text-center justify-center items-center aspect-square min-w-1/3"
  if (cell.initial) {
    cellClass += " bg-slate-300"
  }

  return html`
    <div
      class=${cellClass}
      onClick=${onClick}
    >
      ${cell.value ? cell.value : ""}
    </div>
  `
}
