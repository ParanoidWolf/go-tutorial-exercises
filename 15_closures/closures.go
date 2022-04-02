package main

import "fmt"

func return_anon() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  counter1 := return_anon()
  counter2 := return_anon()
  fmt.Println(counter1())
  fmt.Println(counter2())
  fmt.Println(counter1())
  fmt.Println(counter1())
  fmt.Println(counter2())
  fmt.Println(counter2())
}
