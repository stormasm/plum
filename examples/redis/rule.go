package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
  primarykey := redisc.Get_primary_key("100")
  fmt.Println("primarykey = ", primarykey)
}
