package handle

import (
	"fmt"
	"github.com/stormasm/mux"
	"net/http"
)

func StartHandlers() {

	http.HandleFunc("/api/1.0/admin/token", MyToken1)
	http.HandleFunc("/api/1.0/admin/account", MyToken2)

	r := mux.NewRouter()
	r.HandleFunc("/api/1.0/event", Event1)
	r.HandleFunc("/api/1.0/event/{dimension}/{key}", Event_data)
	r.HandleFunc("/api/1.0/event/{dimension}/{key}/{calculation}/{interval}", Calculated_data)

	http.Handle("/", r)

	http.HandleFunc("/api/1.0/rule/comparator", MyRuleComparator)
	http.HandleFunc("/api/1.0/rule/observer", MyRuleObserver)
	fmt.Println("Listening on port 4567")
	http.ListenAndServe(":4567", nil)
}
