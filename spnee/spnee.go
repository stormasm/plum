package spnee

import (
	"fmt"
	"github.com/stormasm/mux"
	"github.com/stormasm/plum/handle"
	"net/http"
)

func Start() {

	http.HandleFunc("/api/1.0/admin/token", handle.MyToken1)
	http.HandleFunc("/api/1.0/admin/account", handle.MyToken2)

	r := mux.NewRouter()
	r.HandleFunc("/api/1.0/event", handle.Event1)
	r.HandleFunc("/api/1.0/event/{dimension}/{key}", handle.Event_data)
	r.HandleFunc("/api/1.0/event/{dimension}/{key}/{calculation}/{interval}", handle.Calculated_data)

	http.Handle("/", r)

	http.HandleFunc("/api/1.0/rule/comparator", handle.MyRuleComparator)
	http.HandleFunc("/api/1.0/rule/observer", handle.MyRuleObserver)
	fmt.Println("Listening on port 4567")
	http.ListenAndServe(":4567", nil)
}
