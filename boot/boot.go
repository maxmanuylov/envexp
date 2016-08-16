package boot

import (
    "github.com/maxmanuylov/envexp/envexp"
    "github.com/maxmanuylov/utils/application"
    "os"
)

func Run() {
    if err := envexp.ExpandStream(os.Stdin, os.Stdout, true); err != nil {
        application.Exit(err.Error())
    }
}