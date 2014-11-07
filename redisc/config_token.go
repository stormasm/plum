package redisc

import (
	"fmt"
)

type TokenConfig struct {
	Db_zero        string
	Db_uuid        string
	Db_ap          string
	Db_start       string
	Key_db_next    string
	Key_db_mapping string
}

func (c *TokenConfig) Dbstart() string {
	db_start := fmt.Sprint(c.Db_start)
	return db_start
}

func NewTokenConfig() (*TokenConfig) {
	cfg := &TokenConfig{
		Db_zero:        "0",
		Db_uuid:        "10",
		Db_ap:          "11",
		Db_start:       "100",
		Key_db_next:    "nextdb",
		Key_db_mapping: "hm:accountid:db",
	}
	fmt.Printf("Initializing TokenConfig\n")
	return cfg
}
