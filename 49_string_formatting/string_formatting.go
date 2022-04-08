package main

import (
	"fmt"
	"os"
)

type point struct{
  x, y int
}

func main() {
  p := point{ x: 1, y: 2}
  fmt.Printf("struct1: %v\n", p)
  fmt.Printf("struct2: %+v\n", p)
  fmt.Printf("struct3: %#v\n", p)
  fmt.Printf("type: %T\n", p)
  fmt.Printf("bool: %t\n", true)
  fmt.Printf("int: %d\n", 123)
  fmt.Printf("bin: %b\n", 123)
  fmt.Printf("char: %c\n", 123)
  fmt.Printf("hex: %x\n", 123)
  fmt.Printf("float: %e\n", 123.001)
  fmt.Printf("float: %E\n", 123.001)
  fmt.Printf("string1: %s\n", "\"string\"")
  fmt.Printf("string2: %q\n", "\"string\"")
  fmt.Printf("string3: %x\n", "hex this thing")
  fmt.Printf("pointer: %p\n", &p)
  fmt.Printf("width1: |%6d|%6d|\n", 12, 1144)
  fmt.Printf("width2: |%6.2f|%6.2f|\n", 12.5, 114.4)
  fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 12.5, 114.4)
  fmt.Printf("width4: |%6s|%6s|\n", "Sample", "Text")
  s := fmt.Sprintf("Sprintf: a %s", "string")
  fmt.Println(s)
  fmt.Fprintf(os.Stderr, "IO: an %s \n", "error")
}
