package spnee

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
)

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
    fmt.Fprintf(res, "Welcome to Admin Token!")
}

func AdminAccountHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome to Admin Account!")
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome to Spnee!")
}
