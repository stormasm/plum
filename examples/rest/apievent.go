package main

import (
  "fmt"
  "github.com/stormasm/goreq"
)

type Event1 struct {
    AccessToken  string `json:"access_token"`
    Dimension string `json:"dimension"`
    Key string `json:"key"`
    Value string `json:"value"`
    CreatedAt string `json:"created_at"`
    Interval []string `json:"interval"`
    Calculation []string `json:"calculation"`
}

func main() {

  item1 := Event1{ AccessToken: "104a5866-b844-4186-9322-19cacdcec298",
                  Dimension: "job-skills",
                  Key: "ruby",
                  Value: "18",
                  CreatedAt: "2014-09-28 16:33:31 -0700",
                  Interval: []string{"hours","weeks","months"},
                  Calculation: []string{"count","sum","average","standard_deviation","linear_regression"},
                  }

  fmt.Println(item1)

  res1, err1 := goreq.Request{
      Method: "POST",
      ContentType: "application/json",
      Uri: "http://localhost:4567/api/1.0/event",
      Body: item1,
  }.Do()

  fmt.Println(err1, res1)

}
