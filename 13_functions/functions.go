package main

import "fmt"

func return_values() (int, int) {
  return 3, 9
}

func main() {
  a,b := return_values()
  fmt.Println(a, b)
}
