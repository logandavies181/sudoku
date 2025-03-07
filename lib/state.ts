import { cell } from "./board.ts"

class State<T> {
  private cbs = new Map<string, (t: T) => void>()

  subscribe(name: string, cb: (t: T) => void) {
    this.cbs.set(name, cb)
  }

  publish(t: T) {
    for (const cb of this.cbs.values()) {
      cb(t)
    }
  }
}

export const PuzzleState = new State<cell[]>
