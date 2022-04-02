package main

import (
  "fmt"
  "math"
)

type geometry interface {
  area() float32
  peri() float32
}

type rect struct {
  width, height float32
}

type circle struct {
  radius float32
}

func (r rect) area() float32 {
  return r.height * r.width
}

func (r rect) peri() float32 {
  return 2 * (r.width + r.height)
}

func (r circle) area() float32 {
  return math.Pi * r.radius * r.radius
}

func (r circle) peri() float32 {
  return 2 * math.Pi * r.radius
}

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.peri())
}

func main() {
  r := rect{ width: 5, height: 10}
  c := circle{ radius: 5}

  measure(r)
  measure(c)
}
