package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    stdin := bufio.NewReader(os.Stdin)

    readLine := func() error {
        line, err := stdin.ReadString('\n')
        fmt.Fprint(os.Stdout, os.ExpandEnv(line))
        return err
    }

    for readLine() == nil {}
}