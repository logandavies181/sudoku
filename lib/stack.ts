class Stack<T> {
  private items: T[] = []
  push(t: T) {
    this.items.push(t)
  }
  pop() {
    return this.items.pop() ?? null
  }
  empty() {
    this.items = []
  }
}

export type StackItem = {
  value: number
}

export const UndoStack = new Stack<StackItem>()

export const RedoStack = new Stack<StackItem>()
