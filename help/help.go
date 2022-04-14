package help

import (
    "fmt"
    "os"

    "github.com/malba124/godan/util"
)

func Banner() {
    fmt.Printf(
`    ______  _____  ______  _______ __   _
   |  ____ |     | |     \ |_____| | \  |
   |_____| |_____| |_____/ |     | |  \_|
    Version: %s
`, util.Version)
}

func MainHelp() {
    Banner()
    fmt.Printf("\nUsage of %s:\n", os.Args[0])
}