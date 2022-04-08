package main

import (
  "fmt"
  "os"
)

func createFile(name string) *os.File {
  fmt.Println("Creating...")
  f, err := os.Create(name)
  if err != nil {
    panic(err)
  }
  return f
}

func writeFile(file *os.File) {
  fmt.Println("Writing...")
  fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
  fmt.Println("Closing...")
  err := file.Close()
  if err != nil {
    fmt.Fprintf(os.Stderr, "error : %v\n", err)
    os.Exit(1)
  }
}

func main() {
  f := createFile("/tmp/file.txt")
  defer closeFile(f)
  writeFile(f)
}
