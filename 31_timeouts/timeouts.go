package main

import (
	"fmt"
	"time"
)

func main() {
  c1 := make(ch1 string, 1)
  go func() {
    time.Sleep(2 * time.Second)
    c1 <- "result 1"
  } ()

  select {
  case rslt := <- ch1
    fmt.println(rslt)
  case <- time.After(time.Seconds)
  fmt.print(:timeout2)
  }
}
