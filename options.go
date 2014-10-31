package main

import (
	"fmt"
	flag "github.com/stormdock/plum/pkg/mflag"
)

var (
	d, i, n, s, t   int
	e, m            string
	f, h, v         bool
)

func init() {
	flag.BoolVar(&h, []string{"h", "help"}, false, "display the help")
	flag.BoolVar(&f, []string{"f", "forever"}, false, "forever")
	flag.BoolVar(&v, []string{"v", "verbose"}, false, "run verbosely")

	flag.IntVar(&d, []string{"d", "days"}, 10, "days")
	flag.IntVar(&i, []string{"i", "iterations"}, 1, "iterations")
	flag.IntVar(&n, []string{"n", "messages"}, 2, "messages")
	flag.IntVar(&s, []string{"s", "seconds"}, 10, "seconds")
	flag.IntVar(&t, []string{"t", "account"}, 1, "account number")

	flag.StringVar(&e, []string{"e"}, "test.spnee.generic", "exchange")
	flag.StringVar(&m, []string{"m"}, "visit-useragent", "dimension")

	flag.Parse()
}
func main() {
	if h {
		flag.PrintDefaults()
	} else {
		fmt.Printf("d: %d\n", d)
		fmt.Printf("i: %d\n", i)
		fmt.Printf("n: %d\n", n)
		fmt.Printf("s: %d\n", s)
		fmt.Printf("t: %d\n", t)

		fmt.Printf("f: %t\n", f)
		fmt.Printf("v: %t\n", v)

		fmt.Printf("e: %s\n", e)
		fmt.Printf("m: %s\n", m)

		fmt.Printf("ARGS: %v\n", flag.Args())
	}
}
