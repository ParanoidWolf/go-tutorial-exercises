package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
  Page int
  Fruits []string
}

type response2 struct {
  Page int `json:"page"`
  Fruits []string `json:"fruits"`
}

func main() {
  bolB, _ := json.Marshal(true)
  fmt.Println(string(bolB))

  intB, _ := json.Marshal(1)
  fmt.Println(string(intB))

  fltB, _ := json.Marshal(3.14)
  fmt.Println(string(fltB))

  strB, _ := json.Marshal("hello world")
  fmt.Println(string(strB))

  slcB, _ := json.Marshal([]string{"hello", "world"})
  fmt.Println(string(slcB))

  mapB, _ := json.Marshal(map[string]int{"map":5, "paper":1})
  fmt.Println(string(mapB))

  res1D := response1{
    Page: 1,
    Fruits: []string{"apple", "peach", "pear"},
  }
  res1B, _ := json.Marshal(res1D)
  fmt.Println(string(res1B))

  res2D := response1{
    Page: 1,
    Fruits: []string{"apple", "peach", "pear"},
  }
  res2B, _ := json.Marshal(res2D)
  fmt.Println(string(res2B))

  byt := []byte(`{"num":3.14,"strs":["a","b"]}`)

  var dat map[string]interface{}

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fmt.Println(dat)

  num := dat["num"].(float64)
  fmt.Println(num)

  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  fmt.Println(str1)

  str := `{"page":1, "fruits":["apple", "peach"]}`
  res := response2{}
  json.Unmarshal([]byte(str), &res)
  fmt.Println(res)
  fmt.Println(res.Fruits[0])

  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple":5, "lettuce":7}
  enc.Encode(d)
}
