package main

import (
  "fmt"
  "time"
)

func worker(id int, jobs <- chan int, result chan <- int) {
  for j := range jobs {
    fmt.Println("worker", id, "started job", j)
    time.Sleep(time.Second)
    fmt.Println("worker", id, "finished job", j)
    result <- 2 * j
  }
}

func main() {
  const numJobs = 5
  jobs := make(chan int, numJobs)
  results := make(chan int, numJobs)

  for i := 1; i <=3; i++ {
    go worker(i, jobs, results)
  }

  for j := 1; j <= numJobs; j++ {
    jobs <- j
  }
  close(jobs)

  for a := 1; a <= numJobs; a++ {
    <- results
  }
}
