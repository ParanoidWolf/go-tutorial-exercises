package main

import (
  "fmt"
  "math/rand"
  "sync/atomic"
  "time"
)

type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
  key int
  value int
  resp chan bool
}

func main() {
  var readOps uint32
  var writeOps uint32

  reads := make(chan readOp)
  writes := make(chan writeOp)

  go func() {
    var state = make(map[int]int)
    for {
      select {
      case read := <- reads:
        read.resp <- state[read.key]
      case write := <- writes:
        state[write.key] = write.value
        write.resp <- true
      }
    }
  }()

  for r := 0; r < 100; r++ {
    go func() {
      for {
        read := readOp{
          key: rand.Intn(5),
          resp: make(chan int),
        }
        reads <- read
        <- read.resp
        atomic.AddUint32(&readOps, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  for w := 0; w < 10; w++ {
    go func() {
      for {
        write := writeOp{
          key: rand.Intn(5),
          value: rand.Intn(100),
          resp: make(chan bool),
        }
        writes <- write
        <- write.resp
        atomic.AddUint32(&writeOps, 1)
        time.Sleep(time.Millisecond)
      }
    }()
  }

  time.Sleep(time.Second)
  readOpsFinal := atomic.LoadUint32(&readOps)
  fmt.Println("readOps: ", readOpsFinal)
  writeOpsFinal := atomic.LoadUint32(&writeOps)
  fmt.Println("writeOps: ", writeOpsFinal)
}
