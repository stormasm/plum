package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
  dbnumber := redisc.GetDbNumber_from_accountid("3")
  fmt.Println(dbnumber)
}
