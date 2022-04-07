package main

import (
  "fmt"
  "sort"
)

func main() {
  var chars = []string{"c", "a", "b"}
  sort.Strings(chars)
  fmt.Println("Strings: ", chars)

  var nums = []int{1,9, 3, 2, 7, 5}
  sort.Ints(nums)
  fmt.Println("Numbers: ", nums)

  s := sort.IntsAreSorted(nums)
  fmt.Println("Is number sorted?: ", s)
}
