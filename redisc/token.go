package redisc

import "fmt"
import "strconv"
import "strings"
import "github.com/garyburd/redigo/redis"

func Get_apkey_from_token(token string) string {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	tokencfg := NewTokenConfig()
	redis.String(c.Do("SELECT", tokencfg.Db_uuid))
	account, err := redis.String(c.Do("HGET", token, "account"))
	project, err := redis.String(c.Do("HGET", token, "project"))

	if err != nil {
		return "-1"
	}

	apkey := Get_apkey_from_account_project(account, project)
	return apkey
}

func Get_account_from_apkey(apkey string) string {
	s := strings.Split(apkey, ":")
	account := s[0]
	return account
}

func Get_project_from_apkey(apkey string) string {
	s := strings.Split(apkey, ":")
	project := s[1]
	return project
}

func Get_apkey_from_account_project(account, project string) string {
	apkey := fmt.Sprint(account, ":", project)
	return apkey
}

func GetDbNumber_from_accountid(account string) string {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	tokencfg := NewTokenConfig()
	redis.String(c.Do("SELECT", tokencfg.Db_dbnumber))
	dbnumber, err := redis.String(c.Do("HGET", tokencfg.Key_db_mapping, account))

	if err != nil {
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
	if dbnumber == "-1" {
		cfg := NewRedisConfig()
		connect_string := cfg.Connect_string()
		c, err := redis.Dial("tcp", connect_string)

		if err != nil {
			panic(err)
		}
		defer c.Close()

		tokencfg := NewTokenConfig()
		redis.String(c.Do("SELECT", tokencfg.Db_dbnumber))
		nextdb, err := redis.String(c.Do("GET", tokencfg.Key_db_next))

		if err != nil {
			dbnumber = tokencfg.Db_start
			plus1 := AddOneToString(dbnumber)
			redis.String(c.Do("SET", tokencfg.Key_db_next, plus1))
		} else {
			dbnumber = nextdb
			plus1 := AddOneToString(dbnumber)
			redis.String(c.Do("SET", tokencfg.Key_db_next, plus1))
		}

		redis.String(c.Do("HSET", tokencfg.Key_db_mapping, account, dbnumber))
	}
	return dbnumber
}

func Create_uuid_account_project(uuidin, account, project string) {
	apkey := Get_apkey_from_account_project(account, project)

	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)

	if err != nil {
		panic(err)
	}
	defer c.Close()

	tokencfg := NewTokenConfig()
	redis.String(c.Do("SELECT", tokencfg.Db_apkey))
	uuid, err := redis.String(c.Do("HGET", apkey, "uuid"))

	if err != nil {
		fmt.Println("account project key does not exist, creating a new uuid")

		redis.String(c.Do("SELECT", tokencfg.Db_uuid))
		redis.String(c.Do("HSET", uuidin, "account", account))
		redis.String(c.Do("HSET", uuidin, "project", project))

		redis.String(c.Do("SELECT", tokencfg.Db_apkey))
		redis.String(c.Do("HSET", apkey, "uuid", uuidin))
		CreateDbNumber_from_accountid(account)
	} else {
		fmt.Println("got uuid ", uuid)
	}
}
