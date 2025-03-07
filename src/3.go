func(n int) {
  if n < 2 {
    return n
  }
  return n * func(n-1)
}