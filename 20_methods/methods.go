package main

import "fmt"

type rect struct{
  width, height int
}

func (r *rect) area() int {
  return r.width * r.height
}

func (r rect) peri() int {
  return 2 * (r.width + r.height)
}

func main() {
  r := rect{5, 10}
  fmt.Println(r.area());
  fmt.Println(r.peri());
  rp := &r;
  fmt.Println(rp.area());
  fmt.Println(rp.peri());
}
