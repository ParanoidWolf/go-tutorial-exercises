package main

import "fmt"

func sum(num ...int) (int) {
  sum := 0
  for _,value := range num {
    sum += value
  }
  fmt.Println(sum)
  return sum
}

func main() {
  sum(1)
  sum(1, 2, 3, 4, 5)
  values := []int{1, 2, 4, 8, 16}
  sum(values...)
}
