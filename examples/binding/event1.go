package main

import (
    "fmt"
    "net/http"
    "github.com/stormasm/plum/binding"
)

type ContactForm struct {
    Email string `json:"email"`
    Message string `json:"message"`
}

func (cf *ContactForm) FieldMap() binding.FieldMap {
  return binding.FieldMap{
    &cf.Email:   "email",
    &cf.Message: "message",
  }
}

type Event1 struct {
  AccessToken  string `json:"access_token"`
  Dimension string `json:"dimension"`
  Key string `json:"key"`
  Value string `json:"value"`
  CreatedAt string `json:"created_at"`
  Interval []string `json:"interval"`
  Calculation []string `json:"calculation"`
}

func (ev *Event1) FieldMap() binding.FieldMap {
  return binding.FieldMap{
    &ev.AccessToken: "access_token",
    &ev.Dimension: "dimension",
    &ev.Key: "key",
    &ev.Value: "value",
    &ev.CreatedAt: "created_at",
    &ev.Interval: "interval",
    &ev.Calculation: "calculation",
  }
}

// Now your handlers can stay clean and simple.
func contacthandler(resp http.ResponseWriter, req *http.Request) {
    contactForm := new(ContactForm)
    errs := binding.Bind(req, contactForm)
    if errs.Handle(resp) {
        return
    }
    fmt.Println("c email = ", contactForm.Email)
    fmt.Println("c message = ", contactForm.Message)
}

func event1handler(resp http.ResponseWriter, req *http.Request) {
  event1 := new(Event1)
  errs := binding.Bind(req, event1)
  if errs.Handle(resp) {
    return
  }
  fmt.Println("access_token = ", event1.AccessToken)
  fmt.Println("dimension = ", event1.Dimension)
  fmt.Println("key = ", event1.Key)
  fmt.Println("value = ", event1.Value)
  fmt.Println("created_at = ", event1.CreatedAt)
  fmt.Println("interval = ", event1.Interval)
  fmt.Println("calculation = ", event1.Calculation)
}

func main() {
    http.HandleFunc("/contact", contacthandler)
    http.HandleFunc("/api/1.0/event", event1handler)
    fmt.Println("Listening on port 4567")
    http.ListenAndServe(":4567", nil)
}
