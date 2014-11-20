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

  apkey := redisc.Get_apkey_from_token("3339efca-5e99-4ea9-9cff-2075136e04bf")
  fmt.Println(apkey)

}
