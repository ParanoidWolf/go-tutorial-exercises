package main

import (
	"fmt"
	"time"
)

func main() {
  c1 := make(chan string, 1)
  go func() {
    time.Sleep(2 * time.Second)
    c1 <- "result 1"
  } ()

  select {
  case rslt := <- c1:
    fmt.Println(rslt)
  case <- time.After(1 * time.Second): // time.After returns current time through the channel after the delay
    fmt.Println("timeout 1")
  }

  c2 := make(chan string, 1)
  go func() {
    time.Sleep(2 * time.Second)
    c2 <- "result 2"
  } ()

  select {
  case rslt := <- c2:
    fmt.Println(rslt)
  case <- time.After(3 * time.Second):
    fmt.Println("timeout 2")
  }
}
