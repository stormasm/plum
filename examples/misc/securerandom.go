package main

import (
  "fmt"
  sr "github.com/stormasm/securerandom"
)

func main() {
  b, _ := sr.Base64(10, true)
  fmt.Println(b)
  b, _ = sr.Hex(10)
  fmt.Println(b)
  b, _ = sr.Uuid()
  fmt.Println(b)
}
