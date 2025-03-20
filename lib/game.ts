import { cell } from "./cell.ts"
import { globalState } from "./state.ts"
import { NewPuzzle } from "./sudoku.ts"

export function handleUpdate(index: number) {
  const c = puzzleState.state[index]
  if (c.initial) {
    return
  }

  puzzleState.state[index].value = globalState.state.activeNum
  puzzleState.publish(index)
}

export function newGame() {
  const board = NewPuzzle(30, 10_000)

  const cells = new Array<cell>(board.length)
  for (let i = 0; i < cells.length; i++) {
    cells[i] = {
      value: board[i],
      initial: board[i] != 0,
    }
  }
  puzzleState.publishAll(cells)
}

class State {
  private cbs = new Array<(c: cell) => void>(81)

  public state = new Array<cell>(81)

  subscribe(index: number, cb: (c: cell) => void) {
    this.cbs[index] = cb
  }

  publish(index: number) {
    this.cbs[index](copyOf(this.state[index]))
  }

  publishAll(cs: cell[]) {
    const len = cs.length
    if (len != 81) {
      console.error(`cannot publishAll array of length: ${len}`)
      return
    }

    for (let i = 0; i < cs.length; i++) {
      this.state[i] = cs[i]
      this.publish(i)
    }
  }
}

export const puzzleState = new State()

function copyOf(c: cell): cell {
  return {
    value: c.value,
    initial: c.initial,
  }
}
