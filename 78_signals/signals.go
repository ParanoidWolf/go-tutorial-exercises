package main

import (
  "os/signal"
  "syscall"
  "os"
  "fmt"
)

func main() {
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  done := make(chan bool, 1)

  go func() {
    sig := <- sigs
    fmt.Println()
    fmt.Println(sig)
    done <- true
  }()

  fmt.Println("awaiting Signal")
  <- done
  fmt.Println("exiting")
}
