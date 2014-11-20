package main

import (
    "fmt"
    "net/http"
    "github.com/stormasm/plum/binding"
)

type ContactForm struct {
    Email   string `json:"email"`
    Message string `json:"message"`
}

// Then provide a field mapping (pointer receiver is vital)
func (cf *ContactForm) FieldMap() binding.FieldMap {
    return binding.FieldMap{
        &cf.Email:   "email",
        &cf.Message: "message",
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

func main() {
    http.HandleFunc("/contact", contacthandler)
    fmt.Println("Listening on port 4567")
    http.ListenAndServe(":4567", nil)
}
