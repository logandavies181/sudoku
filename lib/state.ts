import { cell } from "./board.ts"

class State {
  private cbs = new Array<(c: cell) => void>(81)

  subscribe(index: number, cb: (c: cell) => void) {
    this.cbs[index] = cb
  }

  publish(index: number, c: cell) {
    this.cbs[index](c)
  }

  publishAll(cs: cell[]) {
    for (let i = 0; i < cs.length; i++) {
      this.publish(i, cs[i])
    }
  }
}

export const PuzzleState = new State()
