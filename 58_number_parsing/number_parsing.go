package main

import (
  "fmt"
  "strconv"
)

func main() {
  f, _ := strconv.ParseFloat("3.14", 64)
  fmt.Println(f)
  i, _ := strconv.ParseInt("314", 0, 64)
  fmt.Println(i)
  d, _ := strconv.ParseInt("0x1c8", 0, 64)
  fmt.Println(d)
  u, _ := strconv.ParseUint("64", 0, 64)
  fmt.Println(u)
  k, _ := strconv.Atoi("64")
  fmt.Println(k)
  _, e := strconv.Atoi("Duh")
  fmt.Println(e)
}
