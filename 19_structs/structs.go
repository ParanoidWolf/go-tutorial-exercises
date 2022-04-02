package main

import "fmt"

type person struct {
  name string
  age int
}

func newPerson(name string, age int) *person {
  p := person{name: name, age: age}
  return &p
}

func main() {
  fmt.Println(newPerson("Steve", 22))
  fmt.Println(newPerson("RDJ", 24))
  fmt.Println(newPerson("Bruce", 20))
}
