package main

import "fmt"

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
  base
  str string
}

func main() {
  co := container{
    base: base{
      num: 1,
    },
    str: "some string",
  }
  fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
  fmt.Printf("Also, num: %v", co.base.num)
  fmt.Println("describle:", co.describe())

  type describer interface {
    describe() string
  }

  var d describer = co
  fmt.Println("describer:", d.describe())
}
