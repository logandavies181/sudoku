export function xPos(src: number): number {
  return src % 9
}

export function yPos(src: number): number {
  return Math.floor(src / 9)
}

export function boxIndex(src: number): number {
  const x = xPos(src)
  const y = yPos(src)

  return Math.floor(x / 3) + Math.floor(y / 3) * 3
}

export function canSee(src: number, dst: number): boolean {
  return xPos(src) == xPos(dst) || yPos(src) == yPos(dst) || boxIndex(src) == boxIndex(dst)
}
