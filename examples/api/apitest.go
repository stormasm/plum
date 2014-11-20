package main

import (
  "fmt"
  "github.com/stormasm/goreq"
)

type TestForm1 struct {
    Red string `json:"red"`
    Yellow string `json:"yellow"`
}

type TestForm2 struct {
    Blue string `json:"blue"`
    Green string `json:"green"`
}

func main() {

  item1 := TestForm1{ Red: "Sam",
                      Yellow: "David"}

  item2 := TestForm2{ Blue: "Jon",
                      Green: "6"}

  fmt.Println(item1)
  fmt.Println(item2)

  res1, err1 := goreq.Request{
      Method: "POST",
      ContentType: "json",
      Uri: "http://localhost:4567/api/1.0/test1",
      Body: item1,
  }.Do()

  res2, err2 := goreq.Request{
      Method: "POST",
      ContentType: "application/json",
      Uri: "http://localhost:4567/api/1.0/test2",
      Body: item2,
  }.Do()

  fmt.Println(err1, res1)
  fmt.Println(err2, res2)

}
