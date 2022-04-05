package main

import "fmt"

func main() {
  jobs := make(chan int, 5)
  done := make(chan bool)

  go func() {
    for {
      j, more := <- jobs
      if more {
        fmt.Println("recieved ", j)
      } else {
        fmt.Println("recieved all jobs")
        done <- true
        return
      }
    }
  } ()

  for j := 0; j <= 3; j++ {
    jobs <- j
    fmt.Println("sent ", j)
  }
  close(jobs)
  fmt.Println("Sent all jobs")

  <- done
}
