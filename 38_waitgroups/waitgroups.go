package main

import (
  "fmt"
  "sync"
  "time"
)

func worker(id int) {
  fmt.Printf("Worker %d starting\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Worker %d done\n", id)
}

func main() {
  var wg sync.WaitGroup

  for i := 1; i <= 5; i++ {
    wg.Add(1)

    //Avoid reuse of the same i variable for each goroutine closure.
    //It shares the same variable across all the goroutunes and thus the values will be modified while goroutine is executed
    i := i

    go func() {
      defer wg.Done()
      worker(i)
    }()
  }

  //waits till the worker group counter reaches 0 after each go routine finishes
  wg.Wait()
}
