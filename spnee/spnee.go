package spnee

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
)

func Start() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", HomeHandler)
  mux.HandleFunc("/ashland", AshlandHandler)
  mux.HandleFunc("/bend", BendHandler)

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome to Corvallis!")
}

func AshlandHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome to Ashland!")
}

func BendHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome to Bend!")
}
