package reader

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadLines(source string) {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("Reader: %#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
