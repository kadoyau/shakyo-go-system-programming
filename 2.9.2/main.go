package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	header := []string{"ID", "Name"}
	writer.Write(header)
	writer.Flush()
	contents := []string{"1", "月ノ美兎"}
	writer.Write(contents)
	writer.Flush()
	contents2 := [][]string{{"2", "月の美兎"}, {"3", "樋口楓"}}
	writer.WriteAll(contents2)
	writer.Flush()
}
