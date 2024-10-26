export function isPrime(x: number): boolean {
  if (x < 2) return false
  if (x === 2) return true
  if (x % 2 === 0) return false

  const limit = Math.sqrt(x)
  for (let i = 3; i <= limit; i += 2) {
    if (x % i === 0) {
      return false
    }
  }
  return true
}
