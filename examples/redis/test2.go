package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
cfg := redisc.NewTokenConfig()
fmt.Println(cfg.Dbstart())
fmt.Println(cfg.Db_uuid)
}
