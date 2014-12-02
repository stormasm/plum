package main

import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
	dbnumber := "100"
	project := "2"
	dimension := "job-skills"
	key := "java"

	calculation := "sum"
	interval := "weeks"

	calculation = "count"
	interval = "hours"

	result := redisc.Get_event_data(dbnumber, project, dimension, key)
	fmt.Println(result)
	result = redisc.Get_calculated_data(dbnumber, project, dimension, key, calculation, interval)
	fmt.Println(result)
}
