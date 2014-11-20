package main

import (
  "fmt"
  "github.com/stormasm/goreq"
)

type ContactForm struct {
    Email   string `json:"email"`
    Message string `json:"message"`
}

func contactRequest() {

  item := ContactForm{ Email: "f@g.edu", Message: "Hola Thursday" }

  res, err := goreq.Request{
      Method: "POST",
      ContentType: "application/json",
      Uri: "http://localhost:4567/contact",
      Body: item,
  }.Do()

  fmt.Println(err)
  fmt.Println(res)
}

func main() {
  contactRequest()
}
