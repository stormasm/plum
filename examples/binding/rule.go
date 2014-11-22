package main

import (
	//"encoding/json"
	"fmt"
	"github.com/stormasm/plum/binding"
	//"github.com/stormasm/plum/rabbit"
	"github.com/stormasm/plum/redisc"
	//"log"
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
		&ro.Project:   "project",
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
		&rc.Project:     "project",
		&rc.Dimension:   "dimension",
		&rc.Key:         "key",
		&rc.Calculation: "calculation",
		&rc.Threshold:   "threshold",
		&rc.Operator:    "operator",
		&rc.Interval:    "interval",
	}
}

func (ro *RuleObserver) Process_observer() {
	fmt.Println("account = ", ro.Account)
	dbnumber := redisc.GetDbNumber_from_account(ro.Account)
	fmt.Println("dbnumber = ", dbnumber)
	pk := redisc.Get_primary_key(dbnumber)
	fmt.Println("primary key = ", pk)
	rulekey := redisc.Build_rule_key(ro.Project, "observer", pk)
	redisc.Process_set_and_interval_key(ro.Project,ro.Interval,rulekey)
}

func (rc *RuleComparator) Process_comparator() {
	fmt.Println("account = ", rc.Account)
	dbnumber := redisc.GetDbNumber_from_account(rc.Account)
	fmt.Println("dbnumber = ", dbnumber)
	pk := redisc.Get_primary_key(dbnumber)
	fmt.Println("primary key = ", pk)
	rulekey := redisc.Build_rule_key(rc.Project, "comparator", pk)
	redisc.Process_set_and_interval_key(rc.Project,rc.Interval,rulekey)
}

func ruleObserverHandler(resp http.ResponseWriter, req *http.Request) {
	observer := new(RuleObserver)
	errs := binding.Bind(req, observer)
	if errs.Handle(resp) {
		return
	}
	fmt.Println("--- observer ---")
	fmt.Println("account = ", observer.Account)
	fmt.Println("project = ", observer.Project)
	fmt.Println("dimension = ", observer.Dimension)
	fmt.Println("key = ", observer.Key)
	fmt.Println("watch = ", observer.Watch)
	fmt.Println("trigger = ", observer.Trigger)
	fmt.Println("interval = ", observer.Interval)

	observer.Process_observer()
}

func ruleComparatorHandler(resp http.ResponseWriter, req *http.Request) {
	comparator := new(RuleComparator)
	errs := binding.Bind(req, comparator)
	if errs.Handle(resp) {
		return
	}
	fmt.Println("--- comparator ---")
	fmt.Println("account = ", comparator.Account)
	fmt.Println("project = ", comparator.Project)
	fmt.Println("dimension = ", comparator.Dimension)
	fmt.Println("key = ", comparator.Key)
	fmt.Println("calculation = ", comparator.Calculation)
	fmt.Println("threshold = ", comparator.Threshold)
	fmt.Println("operator = ", comparator.Operator)
	fmt.Println("interval = ", comparator.Interval)

	comparator.Process_comparator()
}

func main() {
	http.HandleFunc("/api/1.0/rule/comparator", ruleComparatorHandler)
	http.HandleFunc("/api/1.0/rule/observer", ruleObserverHandler)
	fmt.Println("Listening on port 4567")
	http.ListenAndServe(":4567", nil)
}
