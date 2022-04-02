package main

import "fmt"

func main() {
  arr := []int{1, 2, 4, 8, 16}
  sum := 0
  for _, num := range arr {
    sum += num
  }
  fmt.Println(sum)
  for index, value := range arr {
    if value == 8 {
      fmt.Println(index, value)
    }
  }

  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
