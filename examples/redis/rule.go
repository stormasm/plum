package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
  primarykey := redisc.Get_primary_key("100")
  fmt.Println("primarykey = ", primarykey)

  project := "232"
  eventype := "observer"
  rulekey := redisc.Build_rule_key(project, eventype, primarykey)
  fmt.Println("rulekey = ", rulekey)
}
