package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("strinngs.Builder example\n"))
	fmt.Println(builder.String())
}
