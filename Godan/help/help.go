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
    fmt.Printf(`
Usage of %s:

host (Displays information about specified host): 
    -ip <host-ip>
            Specify host ip to look up

api-info (Displays information about provided API key):
    NO local flags for api-info 

shell (Spawn godan in a shell-like enviorment)
    -q 
        Disables banner 

GLOBAL FLAGS:
    -v  (Default: false)
            Verbose output

    -c  (Default: true)
            Colorled output 

    -t <seconds> (Default: 10s)
            Define Request timeout

    -x <key>
            Use specified api key instead of looking for one
            in the config file       
`, os.Args[0])
}