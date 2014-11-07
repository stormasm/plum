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
  fmt.Println(tokencfg.Key_db_mapping)

  value, err := redis.String(c.Do("SELECT", tokencfg.Db_ap))
  fmt.Println(value)
  keys, err := redis.Strings(c.Do("KEYS", "*"))
  fmt.Println(keys)
  //get
  fmt.Println(tokencfg.Key_db_mapping)

  //world, err := redis.Values(c.Do("HGET", tokencfg.Key_db_mapping, account))

  world, err := redis.Int(c.Do("HGET", tokencfg.Key_db_mapping, account))

  if err != nil {
    fmt.Println("key not found", err)
  }
  fmt.Println(world)
  fmt.Println(account)

  fmt.Println("Bye")
  return "bye"
}

/*
func GetDbNumber_from_accountid(account string) {
    fmt.Println("Bye")
}

func GetDbNumber_from_accountid(account string) {
  cfg := NewRedisConfig()
  connect_string := cfg.Connect_string()
  c, err := redis.Dial("tcp", connect_string)
  if err != nil {
  panic(err)
  }
  defer c.Close()

  redis.String(c.Do("SELECT", "11"))
  //get
  world, err := redis.Values(c.Do("HGET", "hm:accountid:db", account))
  if err != nil {
  fmt.Println("key not found")
  }
  fmt.Println(world)
  fmt.Println("bye")
}
*/
