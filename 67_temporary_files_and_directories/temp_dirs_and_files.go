package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  f, err := os.CreateTemp("","sample")
  check(err)

  fmt.Println("temp file name", f.Name())

  defer os.RemoveAll(f.Name())

  _, err = f.Write([]byte{0, 1, 2, 3, 4, 5})
  check(err)

  dname, err := os.MkdirTemp("","sampleTmpDir")
  check(err)
  fmt.Println("Tmp dir name: ", dname)

  defer os.RemoveAll(dname)

  fname := filepath.Join(dname, "file1")
  check(os.WriteFile(fname, []byte{0, 1, 2}, 0666))
}
