import { isPrime } from '.'
import { describe, expect, test } from 'vitest'

describe('isPrime', () => {
  test('1 is not a prime number', () => {
    expect(isPrime(1)).toBe(false)
  })

  test('2 is a prime number', () => {
    expect(isPrime(2)).toBe(true)
  })

  test('3 is a prime number', () => {
    expect(isPrime(3)).toBe(true)
  })

  test('4 is not a prime number', () => {
    expect(isPrime(4)).toBe(false)
  })

  test('17 is a prime number', () => {
    expect(isPrime(17)).toBe(true)
  })

  test('18 is not a prime number', () => {
    expect(isPrime(18)).toBe(false)
  })

  test('19 is a prime number', () => {
    expect(isPrime(19)).toBe(true)
  })

  test('104729 is a prime number', () => {
    expect(isPrime(104729)).toBe(true)
  })

  test('9999999967 is a prime number', () => {
    expect(isPrime(9999999967)).toBe(true)
  })

  test('9999999968 is not a prime number', () => {
    expect(isPrime(9999999968)).toBe(false)
  })

  test('9007199254740847 is a prime number', () => {
    expect(isPrime(9007199254740847)).toBe(true)
  })
})
