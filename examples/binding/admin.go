package main

import (
	"fmt"
	"github.com/stormasm/plum/binding"
	"github.com/stormasm/plum/redisc"
	"net/http"
)

type Token1 struct {
	AccessToken  string `json:"access_token"`
	Token string `json:"token"`
	Account string `json:"account"`
	Project string `json:"project"`
}

type Token2 struct {
	AccessToken  string `json:"access_token"`
	Account string `json:"account"`
	Project string `json:"project"`
}

func (t1 *Token1) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t1.AccessToken: "access_token",
		&t1.Token: "token",
		&t1.Account: "account",
		&t1.Project: "project",
	}
}

func (t2 *Token2) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t2.AccessToken: "access_token",
		&t2.Account: "account",
		&t2.Project: "project",
	}
}

func token1Handler(resp http.ResponseWriter, req *http.Request) {
	t1 := new(Token1)
	errs := binding.Bind(req, t1)
	if errs.Handle(resp) {
		return
	}
	fmt.Println("access_token = ", t1.AccessToken)
	fmt.Println("token = ", t1.Token)
	fmt.Println("account = ", t1.Account)
	fmt.Println("project = ", t1.Project)

	mybool := redisc.Authenticate_admin(t1.AccessToken)
	if(mybool) {
		redisc.Create_uuid_account_project(t1.Token,t1.Account,t1.Project)
	}
}

func main() {
	http.HandleFunc("/api/1.0/admin/token", token1Handler)
	//http.HandleFunc("/api/1.0/admin/account", adminAccountHandler)
	fmt.Println("Listening on port 4567")
	http.ListenAndServe(":4567", nil)
}
