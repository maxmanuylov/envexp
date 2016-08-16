package envexp

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func ExpandStream(reader io.Reader, writer io.Writer, expand bool) error {
    lineScanner := bufio.NewScanner(reader)

    for lineScanner.Scan() {
        if expand {
            fmt.Fprintln(writer, os.ExpandEnv(lineScanner.Text()))
        } else {
            fmt.Fprintln(writer, lineScanner.Text())
        }
    }

    return lineScanner.Err()
}