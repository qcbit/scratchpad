package dependencyInjection

import (
	"bytes"
	"fmt"
)

// Greet writes a greeting to the writer.
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
