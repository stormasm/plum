package redisc

import "fmt"
import "strconv"
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

func AddOneToString(value string) string {
	myint, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("problem converting string to integer")
		return "-1"
	}
	myint_plus1 := myint + 1
	plus1 := strconv.Itoa(myint_plus1)
	return plus1
}


func CreateDbNumber_from_accountid(account string) string {
	dbnumber := GetDbNumber_from_accountid(account)
	// First make sure it does not exist
	if dbnumber == "-1" {

		cfg := NewRedisConfig()
		connect_string := cfg.Connect_string()
		c, err := redis.Dial("tcp", connect_string)
		if err != nil {
			panic(err)
		}
		defer c.Close()

		tokencfg := NewTokenConfig()
		redis.String(c.Do("SELECT", tokencfg.Db_ap))

		nextdb, err := redis.String(c.Do("GET", tokencfg.Key_db_next))

		fmt.Println(nextdb)

		if err != nil {
			fmt.Println("nextdb does not exist", err)

			dbnumber = tokencfg.Db_start
			plus1 := AddOneToString(dbnumber)
			fmt.Println("nextdb = ", plus1)
			// store it in redis
			redis.String(c.Do("SET", tokencfg.Key_db_next, plus1))

		}	else {
			dbnumber = nextdb
			fmt.Println("nextdb already exists adding 1 to it")
			plus1 := AddOneToString(dbnumber)
			fmt.Println("nextdb = ", plus1)
			// store it in redis
			redis.String(c.Do("SET", tokencfg.Key_db_next, plus1))
		}
		fmt.Println("Before")
		redis.String(c.Do("HSET", tokencfg.Key_db_mapping, account, dbnumber))
		fmt.Println("After")
		//@redisc.hset(@key_db_mapping,account,db_number.to_s)
	}
	return dbnumber
}
