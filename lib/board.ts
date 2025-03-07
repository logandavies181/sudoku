import { html } from "../html.ts"
import { PuzzleState } from "./state.ts"
import { NewPuzzle } from "./sudoku.ts"

import { useEffect, useState } from "https://esm.sh/preact@10.25.3/hooks"

export type cell = {
  value: number
}

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
    const now = Date.now()
    const board = NewPuzzle(30, 10_000)
    console.log(Date.now() - now)

    const cells = new Array<cell>(board.length)
    for (let i = 0; i < cells.length; i++) {
      cells[i] = { value: board[i] }
    }
    PuzzleState.publish(cells)
  }, [])

  return html`
    <div class="flex grow flex-wrap bg-white min-w-full">
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
  const [cell, setCell] = useState({ value: 0 })

  PuzzleState.subscribe(`${props.index}`, (c) => {
    setCell(c[props.index])
  })

  return html`
    <div class="text-md border-1 border-solid flex text-center justify-center items-center aspect-square min-w-1/3">
      ${cell.value ? cell.value : ""}
    </div>
  `
}
