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

/* get calculated data

def get_calculated_data(db_number,project,dimension,key,calculation,interval)
@redisc.select db_number
hashkey = build_hash_key(project,dimension,key,calculation,interval)
hmap = @redisc.hgetall(hashkey)
end

get event data

def build_primary_key(project,dimension,key,primarykey)
project + ":" + dimension + ":" + key + ':' + primarykey
end

def build_set_key(project,dimension,key)
project + ":" + dimension + ":" + key + ":set"
end

def build_hash_key(project,dimension,key,calculation,interval)
"hash:" + project + ":" + dimension + ":" + key + ':' + calculation + ":" + interval
end


def get_event_data(db_number,project,dimension,key)
myary = []
@redisc.select db_number
keyset = build_set_key(project,dimension,key)
primary_keys = @redisc.smembers keyset
primary_keys.each do |primary_key|
keyprimary = build_primary_key(project,dimension,key,primary_key)
hmap = @redisc.hgetall(keyprimary)
myary.push(hmap)
end
myary
end

*/

/*
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

func Process_set_key(project,rulekey string) {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	values := []interface{}{"set:", project, ":rules"}
	setkey := fmt.Sprintf("%s%s%s", values...)
	redis.String(c.Do("SADD", setkey, rulekey))
}

func Process_interval_key(project,interval,rulekey string) {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	values := []interface{}{"set:", project, ":rules:", interval}
	intervalkey := fmt.Sprintf("%s%s%s%s", values...)
	redis.String(c.Do("SADD", intervalkey, rulekey))
}

func Process_set_and_interval_key(project,interval,rulekey string) {
	Process_set_key(project,rulekey)
	Process_interval_key(project,interval,rulekey)
}

func Set_rule_key(dbnumber,key,field,value string) {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	redis.String(c.Do("SELECT", dbnumber))
	redis.String(c.Do("HSET", key, field, value))
}
*/
