package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	index(&buf, "richard")
	fmt.Printf("%s", buf.String())
}
