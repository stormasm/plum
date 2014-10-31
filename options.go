package main

import (
	"fmt"

	flag "github.com/stormdock/plum/pkg/mflag"
)

var (
	i           int
	str, m      string
	b, b2, f, h bool
)

func init() {
	flag.Bool([]string{"#hp", "#-halp"}, false, "display the halp")
	flag.BoolVar(&f, []string{"f", "-forever"}, false, "forever")
	flag.BoolVar(&b, []string{"g", "#gil"}, false, "a simple bool")
	flag.BoolVar(&b2, []string{"#-bool"}, false, "a simple bool")
	flag.BoolVar(&h, []string{"h", "#help", "-help"}, false, "display the help")

	flag.IntVar(&i, []string{"integer", "number"}, -1, "a simple integer")

	flag.StringVar(&str, []string{"s", "#hidden", "-string"}, "", "a simple string") //-s -hidden and --string will work, but -hidden won't be in the usage
	flag.StringVar(&str, []string{"mode"}, "mode1", "set the mode\nmode1: use the mode1\nmode2: use the mode2\nmode3: use the mode3")

	flag.StringVar(&m, []string{"m"}, "visit-useragent", "dimension")

	flag.Parse()
}
func main() {
	if h {
		flag.PrintDefaults()
	} else {
		fmt.Printf("i: %d\n", i)
		fmt.Printf("f: %t\n", f)
		fmt.Printf("m: %s\n", m)

		fmt.Printf("s/#hidden/-string: %s\n", str)
		fmt.Printf("b: %t\n", b)
		fmt.Printf("-bool: %t\n", b2)
		fmt.Printf("s/#hidden/-string(via lookup): %s\n", flag.Lookup("s").Value.String())
		fmt.Printf("ARGS: %v\n", flag.Args())
	}
}
