package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello we are all happy and doing well")

	buf := make([]byte, 20)
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
}
