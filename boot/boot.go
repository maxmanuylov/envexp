package boot

import (
    "github.com/maxmanuylov/envexp/envexp"
    "os"
)

func Run() {
    envexp.ExpandStream(os.Stdin, os.Stdout)
}