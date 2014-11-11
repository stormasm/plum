package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", HomeHandler)
  mux.HandleFunc("/ashland", AshlandHandler)
  mux.HandleFunc("/bend", BendHandler)

/*
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })
*/

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
