import { cell } from "../components/board.ts"

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

class GenericState<T> {
  private cbs = new Map<string, (t: T) => void>()

  subscribe(name: string, cb: (t: T) => void) {
    this.cbs.set(name, cb)
  }

  publish(t: T) {
    this.cbs.forEach((cb) => cb(t))
  }
}
