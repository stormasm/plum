package main

import (
	"encoding/json"
	"fmt"
	"github.com/stormasm/plum/binding"
	"github.com/stormasm/plum/rabbit"
	"github.com/stormasm/plum/redisc"
	"log"
	"net/http"
)

var (
	uri          string = "amqp://guest:guest@localhost:5672/"
	exchangeName string = "test.spnee.generic"
	exchangeType string = "fanout"
	routingKey   string = ""
	// body         string = "ralph in socorro"
	reliable bool = true
)

type ContactForm struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (cf *ContactForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&cf.Email:   "email",
		&cf.Message: "message",
	}
}

type Event1Customer struct {
	AccessToken string   `json:"access_token"`
	Dimension   string   `json:"dimension"`
	Key         string   `json:"key"`
	Value       string   `json:"value"`
	CreatedAt   string   `json:"created_at"`
	Interval    []string `json:"interval"`
	Calculation []string `json:"calculation"`
}

type Event1Storm struct {
	Account     string   `json:"account"`
	Project     string   `json:"project"`
	Dbnumber    string   `json:"dbnumber"`
	Dimension   string   `json:"dimension"`
	Key         string   `json:"key"`
	Value       string   `json:"value"`
	CreatedAt   string   `json:"created_at"`
	Interval    []string `json:"interval"`
	Calculation []string `json:"calculation"`
}

func (ev *Event1Customer) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&ev.AccessToken: "access_token",
		&ev.Dimension:   "dimension",
		&ev.Key:         "key",
		&ev.Value:       "value",
		&ev.CreatedAt:   "created_at",
		&ev.Interval:    "interval",
		&ev.Calculation: "calculation",
	}
}

func (evc *Event1Customer) Transform() *Event1Storm {
	evs := new(Event1Storm)
	access_token := evc.AccessToken
	apkey := redisc.Get_apkey_from_token(access_token)
	account := redisc.Get_account_from_apkey(apkey)
	project := redisc.Get_project_from_apkey(apkey)
	dbnumber := redisc.GetDbNumber_from_account(account)
	evs.Account = account
	evs.Project = project
	evs.Dbnumber = dbnumber
	evs.Dimension = evc.Dimension
	evs.Key = evc.Key
	evs.Value = evc.Value
	evs.CreatedAt = evc.CreatedAt
	evs.Interval = evc.Interval
	evs.Calculation = evc.Calculation
	return evs
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
	event1 := new(Event1Customer)
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

	evc := event1.Transform()
	fmt.Println("OOOOOOOOOOOOO")
	fmt.Println(evc.Account)
	fmt.Println(evc.Project)
	fmt.Println(evc.Dbnumber)
	fmt.Println(evc.Dimension)
	fmt.Println(evc.Key)
	fmt.Println(evc.Value)
	fmt.Println(evc.CreatedAt)
	fmt.Println(evc.Interval)
	fmt.Println(evc.Calculation)

	myjson, err := json.Marshal(evc)
	if err != nil {
		fmt.Println(err)
	} else {
		body := string(myjson)
		fmt.Println(body)
		if err := rabbit.Publish(uri, exchangeName, exchangeType, routingKey, body, reliable); err != nil {
			log.Fatalf("%s", err)
		}
		log.Printf("published %dB OK", len(body))

	}
}

func main() {
	http.HandleFunc("/contact", contacthandler)
	http.HandleFunc("/api/1.0/event", event1handler)
	fmt.Println("Listening on port 4567")
	http.ListenAndServe(":4567", nil)
}
