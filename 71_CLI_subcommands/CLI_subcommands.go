package main

import (
  "fmt"
  "flag"
  "os"
)

func main() {
  fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
  fooEnable := fooCmd.Bool("enable", false, "enable")
  fooName := fooCmd.String("name", "", "name")

  barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
  barLevel := barCmd.Int("level", 0, "level")

  if len(os.Args) < 2 {
    fmt.Println("expected foo or bar subcommand")
    os.Exit(1)
  }

  switch(os.Args[1]) {
  case "foo":
    fooCmd.Parse(os.Args[2:])
    fmt.Println("foo subcommand")
    fmt.Println("Enable: ", *fooEnable)
    fmt.Println("Name: ", *fooName)
    fmt.Println("tail: ", fooCmd.Args())
  case "bar":
    barCmd.Parse(os.Args[2:])
    fmt.Println("bar subcommand")
    fmt.Println("Level: ", *barLevel)
    fmt.Println("tail: ", barCmd.Args())
  default:
    fmt.Println("expected foo or bar subcommand")
    os.Exit(1)
  }
}
