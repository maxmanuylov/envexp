package envexp

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func ExpandStream(reader io.Reader, writer io.Writer) {
    bufReader := bufio.NewReader(reader)

    readLine := func() error {
        line, err := bufReader.ReadString('\n')
        fmt.Fprint(writer, os.ExpandEnv(line))
        return err
    }

    for readLine() == nil {}
}