package redisc

import "fmt"
import "github.com/garyburd/redigo/redis"

func Get_primary_key(dbnumber string) string {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	tokencfg := NewTokenConfig()
	redis.String(c.Do("SELECT", dbnumber))
	redis.String(c.Do("INCR", tokencfg.Key_rule_primary_key))
	primarykey, err := redis.String(c.Do("GET", tokencfg.Key_rule_primary_key))

	if err != nil {
		return "-1"
	}
	return primarykey
}

func Build_rule_key(project, eventype, primarykey string) string {
	values := []interface{}{"hash:", project, ":", eventype, ":rule:", primarykey}
	rulekey := fmt.Sprintf("%s%s%s%s%s%s", values...)
	return rulekey
}
