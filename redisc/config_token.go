package redisc

import (
	"fmt"
)

type TokenConfig struct {
	Db_uuid        string
	Db_apkey       string
	Db_dbnumber    string
	Db_admin       string
	Db_start       string
	Key_db_next    string
	Key_db_mapping string
}

func (c *TokenConfig) Dbstart() string {
	db_start := fmt.Sprint(c.Db_start)
	return db_start
}

func NewTokenConfig() *TokenConfig {
	cfg := &TokenConfig{
		Db_uuid:        "10",
		Db_apkey:       "11",
		Db_dbnumber:    "12",
		Db_admin:       "13",
		Db_start:       "100",
		Key_db_next:    "nextdb",
		Key_db_mapping: "hm:accountid:db",
	}
	return cfg
}
