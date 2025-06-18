package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	// - `123` → int型の `i` に
	// - `1.234` → float64型の `f` に
	// - `1.0e4` (= 10000) → float64型の `g` に
	// - `test` → string型の `s` に
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
