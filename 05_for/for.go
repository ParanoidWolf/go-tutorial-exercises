package main

import "fmt"

func main() {
  i := 3
  for i < 5 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 0; j < 5; j++ {
    fmt.Println(j)
  }

  for k := 0; k < 5; k++ {
    if k == 3 {
      continue
    }
    fmt.Println(k)
  }
}
