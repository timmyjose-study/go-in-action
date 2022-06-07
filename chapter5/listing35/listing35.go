package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// bytes.Buffer implements the io.Writer interface
	var b bytes.Buffer

	b.Write([]byte("Hello"))
	fmt.Fprintf(&b, ", ")
	fmt.Fprintf(&b, "world!")
	fmt.Fprintf(&b, "\n")

	io.Copy(os.Stdout, &b)
}