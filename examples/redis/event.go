package main

//import "fmt"
import "github.com/stormasm/plum/redisc"

func main() {
  dbnumber := "100"
  project := "1"
  dimension := "job-skills"
  key := "ruby"
  calculation := "sum"
  interval := "months"

  redisc.Get_calculated_data(dbnumber,project,dimension,key,calculation,interval)
  //fmt.Println("hashmap = ", hashmap)
}
