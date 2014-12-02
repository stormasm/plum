package main

import "fmt"
import "github.com/garyburd/redigo/redis"
import "github.com/stormasm/plum/redisc"

func main() {
	cfg := redisc.NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//set
	c.Do("SET", "michael", "angerman")

	//get
	world, err := redis.String(c.Do("GET", "michael"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
	//ENDINIT OMIT
}
