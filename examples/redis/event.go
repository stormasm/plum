package main

//import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
  dbnumber := "100"
  project := "2"
  dimension := "job-skills"
  key := "java"
/*
  calculation := "sum"
  interval := "weeks"

  redisc.Get_calculated_data(dbnumber,project,dimension,key,calculation,interval)

  calculation = "count"
  interval = "hours"

  redisc.Get_calculated_data(dbnumber,project,dimension,key,calculation,interval)
*/
  redisc.Get_event_data(dbnumber,project,dimension,key)
}
