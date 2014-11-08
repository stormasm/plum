package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
  //dbnumber := redisc.GetDbNumber_from_accountid("3")
  //fmt.Println(dbnumber)

  //nextvalue := redisc.AddOneToString("1010")
  //fmt.Println(nextvalue)

  dbnumber := redisc.CreateDbNumber_from_accountid("4")
  fmt.Println(dbnumber)
}
