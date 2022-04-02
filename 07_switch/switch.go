package main

import (
  "fmt"
  "time"
) 

func main() {
  i := 6
  switch i {
    case 0:
      fmt.Println("You got 0")
    case 1:
      fmt.Println("You got 1")
    case 2:
      fmt.Println("You got 2")
    case 3:
      fmt.Println("You got 3")
    case 4:
      fmt.Println("You got 4")
    case 5:
      fmt.Println("You got 5")
    case 6:
      fmt.Println("You got 6")
  }
  fmt.Println(time.Now().Hour())
  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("Boolean found")
    case int:
      fmt.Println("Integer found")
    }
  }
}
