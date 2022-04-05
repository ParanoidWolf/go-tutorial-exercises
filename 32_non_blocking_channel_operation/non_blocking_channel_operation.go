package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan string)

  //select makes the channel non-blocking with a default case
  select {
  case msg := <- messages:
    fmt.Println("Message:", msg)
  default:
    fmt.Println("No messages recieved")
  }

  msg := "hello"
  select {
  case signals <- msg:
    fmt.Println("Message:", msg)
  default:
    fmt.Println("No messages sent")
  }

  //a non blocking multi-way select
  select {
  case msg := <- messages:
    fmt.Println("Message:", msg)
  case signals <- msg:
    fmt.Println("Message:", msg)
  default:
    fmt.Println("No messages")
  }
}
