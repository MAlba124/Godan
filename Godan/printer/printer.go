package printer

import (
    "fmt"

    "github.com/malba124/godan/color"
)

const WHITESPACE string = "                                                      "

func Print(str string) {
    fmt.Printf("%s", str)
}

func Error(error string, colored *bool) {

    if *colored || colored == nil {
        fmt.Printf("[ %sERROR%s ] %s\n", color.CRed, color.CReset, error)
    } else {
        fmt.Printf("[ ERROR ] %s\n", error)
    }
}

func Debug(str string, debug bool) {
    if debug {
        fmt.Printf("[ DEBUG ]%s\n", str)
    }
}

func Verbose(str string, verbose *bool) {
    if *verbose {
        fmt.Printf("[ VERBOSE ] %s\n", str)
    }
}
