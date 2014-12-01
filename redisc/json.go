package redisc

import "encoding/json"
import "fmt"

// This function takes in a String array which is the result
// of an hgetall return function reply from redigo
// sa stands for string array
// For now this function is lower case because it is not
// called outside of the redisc package
// meaning it is directly tied to redis functionality for now
// and so it should live in this package as a local function

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
