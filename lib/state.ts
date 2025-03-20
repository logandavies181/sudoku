import { cell } from "./cell.ts"

class GenericState<T> {
  private cbs = new Map<string, (t: T) => void>()
  state: T

  constructor(t: T) {
    this.state = t
  }

  subscribe(name: string, cb: (t: T) => void) {
    this.cbs.set(name, cb)
  }

  publish(t: T) {
    this.state = t
    this.cbs.forEach((cb) => cb(t))
  }
}

export type GlobalState = {
  activeNum: number
}

export const globalState = new GenericState<GlobalState>({
  activeNum: 1,
})
