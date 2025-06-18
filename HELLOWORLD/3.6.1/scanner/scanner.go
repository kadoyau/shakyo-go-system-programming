package scanner

import (
	"bufio"
	"fmt"
	"strings"
)

func ScanLines(source string) {
	scanner := bufio.NewScanner(strings.NewReader(source))
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("Scanner: %#v\n", scanner.Text())
	}
}
