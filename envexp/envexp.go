package envexp

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func ExpandStream(reader io.Reader, writer io.Writer) error {
    lineScanner := bufio.NewScanner(reader)

    for lineScanner.Scan() {
        fmt.Fprintln(writer, os.ExpandEnv(lineScanner.Text()))
    }

    return lineScanner.Err()
}