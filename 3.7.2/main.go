package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = io.ReadAll(teeReader)
	// けどバッファに残ってる
	fmt.Println(buffer.String())
}
