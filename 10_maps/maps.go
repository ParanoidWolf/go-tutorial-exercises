package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["MDL15EC089"] = 44
  m["MDL15EC020"] = 11
  fmt.Println(m)
  delete(m, "MDL15EC020")
  fmt.Println(m)
  _, exists := m["MDL15EC089"]
  fmt.Println(exists)
  n := map[string]int{"MDL": 1, "CET": 2}
  fmt.Println(n)
}
