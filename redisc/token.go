package redisc

import "fmt"
import "github.com/garyburd/redigo/redis"

func GetDbNumber_from_accountid(account string) string {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	tokencfg := NewTokenConfig()
	redis.String(c.Do("SELECT", tokencfg.Db_ap))
	dbnumber, err := redis.String(c.Do("HGET", tokencfg.Key_db_mapping, account))

	if err != nil {
		fmt.Println("dbnumber not found", err)
		return "-1"
	}

	return dbnumber
}
