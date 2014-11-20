package main

import (
	"fmt"
	"github.com/stormasm/plum/binding"
	"net/http"
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
func handler(resp http.ResponseWriter, req *http.Request) {
	contactForm := new(ContactForm)
	errs := binding.Bind(req, contactForm)
	if errs.Handle(resp) {
		return
	}
	fmt.Println("email = ", contactForm.Email)
	fmt.Println("message = ", contactForm.Message)
}

func main() {
	http.HandleFunc("/contact", handler)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
