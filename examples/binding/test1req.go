package main

import (
	"fmt"
	"github.com/stormasm/goreq"
)

type ContactForm struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {

	item := ContactForm{Email: "d@e.edu", Message: "Monday Tuesday"}

	res, err := goreq.Request{
		Method:      "POST",
		ContentType: "application/json",
		Uri:         "http://localhost:3000/contact",
		Body:        item,
	}.Do()

	fmt.Println(err)
	fmt.Println(res)
}
