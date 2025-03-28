import { cell } from "./cell.ts"
import { RedoStack, UndoStack } from "./stack.ts";
import { globalState } from "./state.ts"
import { CheckIfCanUpdateCell, NewPuzzle, UpdateCell } from "./sudoku.ts"

export function handleUpdate(index: number) {
  const c = puzzleState.state[index]
  if (c.initial) {
    return
  }

  const newNum = globalState.state.activeNum

  if (!CheckIfCanUpdateCell(index, newNum)) {
    console.log(`checked and couldn't update ${index} with ${c.value}`)
    console.log(newNum)
    console.log(typeof newNum)
    return
  }
  if (!UpdateCell(index, newNum)) {
    console.log(`got error updating ${index} with ${c.value}`)
    return
  }

  UndoStack.push({
    value: c.value,
  })
  RedoStack.empty()
  puzzleState.state[index].value = newNum
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
  UndoStack.empty()
  RedoStack.empty()
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
