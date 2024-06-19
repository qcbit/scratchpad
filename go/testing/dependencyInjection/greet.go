package dependencyInjection

import (
	"fmt"
	"io"
	"os"
)

// Greet writes a greeting to the writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Gophers")
}
