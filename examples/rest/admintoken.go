package main

import (
  "fmt"
  "github.com/stormasm/goreq"
)

type TokenForm1 struct {
    AccessToken  string `json:"access_token"`
    Token string `json:"token"`
    Account string `json:"account"`
    Project string `json:"project"`
}

type TokenForm2 struct {
    AccessToken  string `json:"access_token"`
    Account string `json:"account"`
    Project string `json:"project"`
}

func main() {

  item1 := TokenForm1{ AccessToken: "104a5866-b844-4186-9322-19cacdcec298",
                       Token: "704a5866-b844-4186-9322-99cacdcec299",
                       Account: "7",
                       Project: "9"}

  item2 := TokenForm2{ AccessToken: "104a5866-b844-4186-9322-19cacdcec298",
                       Account: "6",
                       Project: "8"}

  fmt.Println(item1)
  fmt.Println(item2)

  res1, err1 := goreq.Request{
      Method: "POST",
      ContentType: "json",
      Uri: "http://localhost:4567/api/1.0/admin/token",
      Body: item1,
  }.Do()
/*
  res2, err2 := goreq.Request{
      Method: "POST",
      ContentType: "application/json",
      Uri: "http://localhost:4567/api/1.0/admin/account",
      Body: item2,
  }.Do()
*/
  fmt.Println(err1, res1)
  //fmt.Println(err2, res2)

}
