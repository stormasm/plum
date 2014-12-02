package main

import (
	"fmt"
	"github.com/stormasm/goreq"
	"net/url"
)

type Event1 struct {
	AccessToken string   `json:"access_token"`
	Dimension   string   `json:"dimension"`
	Key         string   `json:"key"`
	Value       string   `json:"value"`
	CreatedAt   string   `json:"created_at"`
	Interval    []string `json:"interval"`
	Calculation []string `json:"calculation"`
}

func main() {

	item1 := Event1{AccessToken: "104a5866-b844-4186-9322-19cacdcec297",
		Dimension:   "job-skills",
		Key:         "ruby",
		Value:       "18",
		CreatedAt:   "2014-09-28 16:33:31 -0700",
		Interval:    []string{"hours", "weeks", "months"},
		Calculation: []string{"count", "sum", "average", "standard_deviation", "linear_regression"},
	}

	fmt.Println(item1)

	res1, err1 := goreq.Request{
		Method:      "POST",
		ContentType: "application/json",
		Uri:         "http://localhost:4567/api/1.0/event",
		Body:        item1,
	}.Do()

	fmt.Println(err1, res1)

	item2 := url.Values{}
	item2.Set("access_token", "28037e39-456d-49e4-998a-17c48ce916aa")

	item3 := url.Values{}
	item3.Set("access_token", "15f32255-aaeb-4d2f-8988-26494bc4d58d")

	res2, err2 := goreq.Request{
		Method:      "GET",
		ContentType: "application/json",
		Uri:         "http://localhost:4567/api/1.0/event/job-skills/python",
		QueryString: item2,
	}.Do()

	fmt.Println(err2, res2)

	res3, err3 := goreq.Request{
		Method:      "GET",
		ContentType: "application/json",
		Uri:         "http://localhost:4567/api/1.0/event/job-skills/java/sum/weeks",
		QueryString: item3,
	}.Do()

	fmt.Println(err3, res3)
}
