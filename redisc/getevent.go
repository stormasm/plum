package redisc

import "encoding/json"
import "fmt"
import "strconv"
import "github.com/garyburd/redigo/redis"

func getpairs(sa []string, args ...string) ([]string, error) {

	mymap := make(map[string]string)

	for i := range args {
		switch {
			case i%2 == 0:
			mymap[args[i]] = args[i+1]
			default:
			//
		}
	}

	str, err := json.Marshal(mymap)
	if err != nil {
		return nil, fmt.Errorf("getpairs error encoding JSON")
	}

	myjson := string(str)

	sa = append(sa, myjson)
	sa = append(sa, ",")
	return sa, nil
}

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
	strings, err := redis.Strings(c.Do("HGETALL", hashkey))

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(strings)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

// This function is currently not being used
// It is here for reference only as a way to get
// data via Values instead of Strings
func Get_calculated_data_values(dbnumber,project,dimension,key,calculation,interval string) {
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

func Get_event_data(dbnumber,project,dimension,key string) string {
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
		fmt.Println(err)
		return("Get_event_data redis error getting primarykeys")
	}

	if len(primarykeys) < 1 {
		return "No primary keys"
	}

	fmt.Printf("%v\n", primarykeys)

	sa := make([]string, 0)

	for pk := range primarykeys {
		fmt.Println(pk)
		pkstr := strconv.Itoa(pk)
		hashkey := Build_primary_key(project, dimension, key, pkstr)
		hstrings, err := redis.Strings(c.Do("HGETALL", hashkey))

		if err != nil {
			fmt.Println(err)
			return("Get_event_data redis error getting hashkey")
		}

		fmt.Println(hashkey)
		fmt.Println(hstrings)
		sa, err = getpairs(sa, hstrings...)
	}

	result := fmt.Sprintf("%s", sa)
	return result
}
