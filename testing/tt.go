package main

import (
	"fmt"
	"bytes"
)

func main() {
	var buf bytes.Buffer
	index(&buf, "richard")
	fmt.Printf("%s",buf.String())
}
