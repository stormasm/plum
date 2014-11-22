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

type RuleObserver struct {
	Account   string  `json:"account"`
	Project   string  `json:"project"`
	Dimension string  `json:"dimension"`
	Key       string  `json:"key"`
	Watch     string  `json:"watch"`
	Trigger   string  `json:"trigger"`
	Interval  string  `json:"interval"`
}

func (ro *RuleObserver) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&ro.Account:   "account",
		&ro.Project:   "project"
		&ro.Dimension: "dimension",
		&ro.Key:       "key",
		&ro.Watch:     "watch",
		&ro.Trigger:   "trigger",
		&ro.Interval:  "interval",
	}
}

type RuleComparator struct {
	Account     string `json:"account"`
	Project     string `json:"project"`
	Dimension   string `json:"dimension"`
	Key         string `json:"key"`
	Calculation string `json:"calculation"`
	Threshold   string `json:"threshold"`
	Operator    string `json:"operator"`
	Interval    string `json:"interval"`
}

func (rc *RuleComparator) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&rc.Account:     "account",
		&rc.Project:     "project"
		&rc.Dimension:   "dimension",
		&rc.Key:         "key",
		&rc.Calculation: "calculation",
		&rc.Threshold:   "threshold",
		&rc.Operator:    "operator",
		&rc.Interval:    "interval",
	}
}

func ruleObserverHandler(resp http.ResponseWriter, req *http.Request) {
	observer := new(RuleObserver)
	errs := binding.Bind(req, observer)
	if errs.Handle(resp) {
		return
	}
	fmt.Println("account = ", observer.Account)
	fmt.Println("project = ", observer.Project)
	fmt.Println("dimension = ", observer.Dimension)
	fmt.Println("key = ", observer.Key)
	fmt.Println("watch = ", observer.Watch)
	fmt.Println("trigger = ", observer.Trigger)
	fmt.Println("interval = ", observer.Interval)
}

func ruleComparatorHandler(resp http.ResponseWriter, req *http.Request) {
	comparator := new(RuleComparator)
	errs := binding.Bind(req, comparator)
	if errs.Handle(resp) {
		return
	}
	fmt.Println("account = ", observer.Account)
	fmt.Println("project = ", observer.Project)
	fmt.Println("dimension = ", observer.Dimension)
	fmt.Println("key = ", observer.Key)
	fmt.Println("calculation = ", observer.Watch)
	fmt.Println("threshold = ", observer.Trigger)
	fmt.Println("operator = ", comparator.Operator)
	fmt.Println("interval = ", observer.Interval)
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
	http.HandleFunc("/api/1.0/rule/comparator", ruleComparatorHandler)
	http.HandleFunc("/api/1.0/rule/observer", ruleObserverHandler)
	fmt.Println("Listening on port 4567")
	http.ListenAndServe(":4567", nil)
}
