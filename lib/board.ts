import { html } from "../html.ts"
import { NewPuzzle } from "./sudoku.ts";

import { useEffect, useState } from "https://esm.sh/preact@10.25.3/hooks"

type cell = {
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

	return Math.floor(x/3) + Math.floor(y/3)*3
}

function newBoxes(): Array<Array<cell>> {
  const b = new Array<Array<cell>>(9)
  for (let i = 0; i < b.length; i++) {
    b[i] = []
  }
  return b
}

export function Board() {
  const initialBoxes = newBoxes()
  const initialCells = []
  for (let i = 0; i < initialBoxes.length; i++) {
    initialCells.push({ value: 0 })
  }
  for (let i = 0; i < initialBoxes.length; i++) {
    initialBoxes[i] = initialCells
  }

  const [boxes, setBoxes] = useState<cell[][]>(initialBoxes);

  useEffect(() => {
    (async () => {
      const now = Date.now()
      const board = NewPuzzle(30, 10_000)
      console.log(Date.now() - now)

      const nb = newBoxes()
      board.forEach((v, i) => {
	nb[boxIndex(i)].push({ value: v })
      })

      setBoxes(nb)
      console.log("loaded")
    })()
  }, [])

  return html`
    <div class="flex grow flex-wrap bg-white min-w-full">
      ${boxes.map((cells) => {
        return html`<${Box} cells=${cells} />`
      })}
    </div>
  `
}

type BoxProps = {
  cells: cell[]
}

function Box(props: BoxProps) {
  return html`
    <div class="border-2 border-solid flex grow flex-wrap bg-white max-w-1/3 min-w-1/3">
      ${props.cells.map((cell) => {
        return html`<${Cell} cell=${cell} />`
      })}
    </div>
  `
}

type CellProps = {
  cell: cell
}

function Cell(props: CellProps) {
  return html`
    <div class="text-md border-1 border-solid flex text-center justify-center items-center aspect-square min-w-1/3">
      ${props.cell.value ? props.cell.value : ""}
    </div>
  `
}
