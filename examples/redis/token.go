package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
	//dbnumber := redisc.GetDbNumber_from_accountid("3")
	//fmt.Println(dbnumber)

	//nextvalue := redisc.AddOneToString("1010")
	//fmt.Println(nextvalue)

	//dbnumber := redisc.CreateDbNumber_from_accountid("1")
	//fmt.Println(dbnumber)

	//fmt.Println("Hi")
	//redisc.Generate_token()

	//apkey := redisc.Get_apkey_from_token("3339efca-5e99-4ea9-9cff-2075136e04bf")
	//fmt.Println(apkey)

	//account := redisc.Get_account_from_apkey(apkey)
	//fmt.Println("account = ", account)

	//project := redisc.Get_project_from_apkey(apkey)
	//fmt.Println("project = ", project)

	// bad token
	// mybool := redisc.Authenticate_admin("3339efca-5e99-4ea9-9cff-2075136e04bf")
	// good token
	mybool := redisc.Authenticate_admin("104a5866-b844-4186-9322-19cacdcec298")
	fmt.Println("authenticate = ", mybool)
}
