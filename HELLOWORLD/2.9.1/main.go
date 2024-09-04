package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Fprintf(file, "%dや%s。現在時刻は%f", 123, "でー", 4.5)
}
