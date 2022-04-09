package main

import (
  "fmt"
  "time"
)

func main() {
  p := fmt.Println

  t := time.Now()
  p(t.Format(time.RFC3339))

  t1, err := time.Parse(
    time.RFC3339,
    "2012-11-01T22:08:41+00:00",
  )
  p(t1)

  p(t.Format("3:04PM"))
  p(t.Format("Mon Jan _2 15:04:24 2012"))
  p(t.Format("2009-01-04T12:12:12.22222-07:00"))

  form := "2 24 PM"
  t2, err := time.Parse(form, "8 41 PM")
  p(t2)

  ansic := "Mon Jan _2 15:04:05 2006"
  _, err = time.Parse(ansic, "8:41PM")
  p(err)
}
