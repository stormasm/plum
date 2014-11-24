package redisc

import "fmt"
import "github.com/garyburd/redigo/redis"


func Build_hash_key(project, dimension, key, calculation, interval string) string {
	values := []interface{}{"hash:", project, ":", dimension, ":", key, ":", calculation, ":", interval}
	hashkey := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", values...)
	return hashkey
}

func Build_primary_key(project, dimension, key, primarykey string) string {
	values := []interface{}{project, ":", dimension, ":", key, ":", primarykey}
	myprimarykey := fmt.Sprintf("%s%s%s%s%s%s%s", values...)
	return myprimarykey
}

func Build_set_key(project, dimension, key string) string {
	values := []interface{}{project, ":", dimension, ":", key, ":set"}
	setkey := fmt.Sprintf("%s%s%s%s%s%s", values...)
	return setkey
}

func Get_calculated_data(dbnumber,project,dimension,key,calculation,interval string) {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	redis.String(c.Do("SELECT", dbnumber))
	hashkey := Build_hash_key(project,dimension,key,calculation,interval)
	fmt.Println(dbnumber, " ", hashkey)
	values, err := redis.Values(c.Do("HGETALL", hashkey))

	if err != nil {
		panic(err)
	}

	var hmap []struct {
		Date  string
		Value string
	}
	if err := redis.ScanSlice(values, &hmap); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", hmap)
	// Output:
	// [{Earthbound 1} {Beat 4} {Red 5}]
}

func Get_event_data(dbnumber,project,dimension,key string) {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	redis.String(c.Do("SELECT", dbnumber))
	setkey := Build_set_key(project,dimension,key)

	primarykeys, err := redis.Strings(c.Do("SMEMBERS", setkey))

	if err != nil {
		// return "", err
		fmt.Println(err)
	}
/*
	if len(primarykeys) != 1 {
		return "", redis.ErrNil
	}
*/
	fmt.Printf("%v\n", primarykeys)
}
