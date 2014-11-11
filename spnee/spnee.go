package spnee

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

type Message struct {
	Name string
	Body string
}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)

	mux.HandleFunc("/api/1.0/admin/token", AdminTokenHandler)
	mux.HandleFunc("/api/1.0/admin/account", AdminAccountHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

func AdminTokenHandler(res http.ResponseWriter, req *http.Request) {

	m := Message{"Admin", "Token"}
	b, err := json.Marshal(m)
	if err != nil {
    e := Message{"Error", "Message"}
    be, err1:= json.Marshal(e)
    fmt.Println(err1)
    myerr := string(be)
		fmt.Fprintf(res,myerr)
	} else {
		mystr := string(b)
		fmt.Fprintf(res, mystr)
	}
}

func AdminAccountHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to Admin Account!")
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to Spnee!")
}
