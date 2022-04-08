package main

import (
  "fmt"
  s "strings"
)

var p = fmt.Println

func main() {
  p("Contains: ", s.Contains("test", "st"))
  p("Count: ", s.Count("test", "t"))
  p("HasPrefix: ", s.HasPrefix("test", "te"))
  p("HasSuffix: ", s.HasSuffix("test", "st"))
  p("Index: ", s.Index("test", "s"))
  p("Join: ", s.Join([]string{"a", "b"}, "-"))
  p("Repeat: ", s.Repeat("ba", 3))
  p("Replace: ", s.Replace("Life is too short to find faults. Just fix those faults and move on!", "faults", "blames", -1))
  p("Replace: ", s.Replace("Life is too short to find faults. Just fix those faults and move on!", "faults", "blames", 1))
  p("Split: ", s.Split("triangle green solid shadow gradient", " "))
  p("ToLower: ", s.ToLower("CAPITAL"))
  p("ToUpper: ", s.ToUpper("lower"))
}
