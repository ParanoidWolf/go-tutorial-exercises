package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func check (e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  err := os.Mkdir("Subdir", 0755)
  check(err)

  defer os.RemoveAll("Subdir")

  createEmptyFile := func(name string) {
    d := []byte("")
    check(os.WriteFile(name, d, 0644))
  }

  createEmptyFile("Subdir/file1")

  check(os.MkdirAll("Subdir/parent/child", 0755))

  c, err := os.ReadDir(".")
  check(err)

  fmt.Println("Listing subdir/parent/child")
  for _, entry := range c {
      fmt.Println(" ", entry.Name(), entry.IsDir())
  }
  fmt.Println("Visiting subdir")
  err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}
