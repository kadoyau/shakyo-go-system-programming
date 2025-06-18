package main

import (
	"fmt"
	"helloworld/3.6.1/reader"
	"helloworld/3.6.1/scanner"
)

var source = `1 行め
2 行め
3 行め`

func main() {
	fmt.Println("Using Reader:")
	reader.ReadLines(source)
	fmt.Println("\nUsing Scanner:")
	scanner.ScanLines(source)
}
