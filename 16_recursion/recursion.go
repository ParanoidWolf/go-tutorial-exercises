package main

import "fmt"

func fact(num int) int {
  if(num <= 0) {
    return 1
  }
  return num * fact(num - 1)
}

func main() {
  fmt.Println(fact(5))

  var compiler func(num int) int // Declared before use inside main

  compiler = func(num int) int {
    if(num <= 0) {
      return 1
    }
    return num * fact(num - 1)
  }

  fmt.Println(compiler(5))
}
