package main

import "fmt"

func return_values() (int) {
  return 3
}

func main() {
  a := return_values()
  fmt.Println(a)
}
