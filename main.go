package main

import (
	"flag"
	"os"

	"github.com/malba124/godan/util"
    "github.com/malba124/godan/printer"
    "github.com/malba124/godan/shodan"
    "github.com/malba124/godan/help"
)

type Host struct {
    ip *string
    color *bool
    verbose *bool
}

type ApiInfo struct {
    verbose *bool
    color *bool
}

func main() {

    if len(os.Args) < 2 {
        help.MainHelp()
        os.Exit(1)
    }

	if util.NotAnArgument("-h", os.Args) && util.NotAnArgument("--help", os.Args) {
        key, err := util.LoadConfig("config.json")
        if err != nil {
            printer.Error("Config file is invalid or not found!", nil)
        } else {
            client := shodan.NewClient(key)

            hostCmd := flag.NewFlagSet("host", flag.ExitOnError)
            hIp := hostCmd.String("ip", "", "Host ip")
            hVerbose := hostCmd.Bool("v", false, "Verbose mode")
            hColor := hostCmd.Bool("c", true, "Disable colored output")
            hTimeout := hostCmd.Int("t", 10, "Specify request timeout")

            apiInfoCmd := flag.NewFlagSet("api-info", flag.ExitOnError)
            aiVerbose := apiInfoCmd.Bool("v", false, "Verbose mode")
            aiColor := apiInfoCmd.Bool("c", true, "Disable colored output")
            aiTimeout := apiInfoCmd.Int("t", 10, "Specify request timeout")

            switch os.Args[1] {
            case "host":
                hostCmd.Parse(os.Args[2:])
                client.Timeout = *hTimeout
                var h Host = Host{
                        ip: hIp,
                        color: hColor,
                        verbose: hVerbose,
                    }
                printer.Verbose("Executing Host()", h.verbose)
                client.Host(*h.ip, h.color)                    
            case "api-info":
                apiInfoCmd.Parse(os.Args[2:])
                client.Timeout = *aiTimeout
                var ai ApiInfo = ApiInfo {
                    verbose: aiVerbose,
                    color: aiColor,
                }
                printer.Verbose("Executing ApiInfo()", ai.verbose)
                client.ApiInfo(ai.color)
            default:
                help.MainHelp()
                printer.Error("Missing arguments", nil)
                os.Exit(1)
            }
        }
    } else {
 		help.MainHelp()
    }
}
